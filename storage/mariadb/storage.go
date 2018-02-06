package mariadb

import (
	"fmt"
	"io"
	alog "log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

//Storage implements the storage interface
type Storage struct {
	db     *sqlx.DB
	log    *alog.Logger
	logErr *alog.Logger
}

// New creates a new storage entry
func New(config string, w io.Writer, wErr io.Writer) (db *Storage, err error) {
	if w == nil {
		w = os.Stdout
	}
	if wErr == nil {
		wErr = os.Stderr
	}
	db = &Storage{
		log:    alog.New(w, "SQL: ", 0),
		logErr: alog.New(w, "SQLErr: ", 0),
	}

	if config == "" {
		err = fmt.Errorf("Invalid configuration passed (empty)")
		return
	}
	if db.db, err = sqlx.Open("mysql", config); err != nil {
		return
	}
	return
}

//InsertTestData will grab data from storage
func (s *Storage) InsertTestData() (err error) {
	_, err = s.db.Exec(`INSERT INTO user (id, display_name, email, password)
	   VALUES
	   	(1, 'Test', 'test@test.com', '$2a$10$YV0PiWDMiuXL4e77.jv8leD3NpDCk.v41aXPn7Yyi7fBWwBa0XzzC');`)
	if err != nil {
		err = errors.Wrap(err, "failed to insert user data")
		return
	}

	_, err = s.db.Exec(`INSERT INTO account (id, name, status)
	   VALUES
	   	(1, 'Shin', 200);`)
	if err != nil {
		err = errors.Wrap(err, "failed to insert account data")
		return
	}
	return
}

//DropTables will grab data from storage
func (s *Storage) DropTables() (err error) {
	_, err = s.db.Exec(`SET FOREIGN_KEY_CHECKS = 0`)
	if err != nil {
		return
	}
	_, err = s.db.Exec(`drop table if exists character_data`)
	if err != nil {
		return
	}
	tables := []string{
		"account",
		"activities",
		"base",
		"bazaar",
		"character_data",
		"faction",
		"forum",
		"goal",
		"item",
		"lootdrop",
		"lootdropentry",
		"loottable",
		"loottablentry",
		"spells_new",
		"npc",
		"npcloot",
		"post",
		"tasks",
		"topic",
		"user",
		"zone",
	}
	for _, table := range tables {
		_, err = s.db.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS %s`, table))
		if err != nil {
			errors.Wrap(err, fmt.Sprintf("Failed to delete %s", table))
			return
		}
	}
	_, err = s.db.Exec(`SET FOREIGN_KEY_CHECKS = 1`)
	if err != nil {
		return
	}
	return
}

//VerifyTables will grab data from storage
func (s *Storage) VerifyTables() (err error) {
	type TableCheck struct {
		Func func() (err error)
		Name string
	}

	tables := []TableCheck{
		{
			Func: s.createTableAccount,
			Name: "Account",
		},
		{
			Func: s.createTableCharacter,
			Name: "Character",
		},
		{
			Func: s.createTableRule,
			Name: "Rule",
		},
		{
			Func: s.createTableRuleEntry,
			Name: "RuleEntry",
		},
		{
			Func: s.createTableUser,
			Name: "User",
		},
		{
			Func: s.createTableVariable,
			Name: "Variable",
		},
		{
			Func: s.createTableZone,
			Name: "Zone",
		},
	}

	for _, table := range tables {
		err = table.Func()
		if err != nil && !isExistErr(err) {
			return
		}
		if err == nil {
			s.log.Println("Created table for", table.Name)
		}
	}
	err = nil
	return
}

func isExistErr(err error) bool {
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == 1050 {
			return true
		}
	}
	return false
}
