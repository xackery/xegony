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

//CreateTestDatabase creates a new test database
func (s *Storage) CreateTestDatabase() (err error) {
	_, err = s.db.Exec(`DROP DATABASE eqemu_test;`)
	if err != nil {
		err = errors.Wrap(err, "failed to drop database")
		return
	}

	_, err = s.db.Exec(`CREATE DATABASE eqemu_test;`)
	if err != nil {
		err = errors.Wrap(err, "failed to create database")
		return
	}
	return
}

//InsertTestData will insert test data to database
func (s *Storage) InsertTestData() (err error) {
	type InsertFunc struct {
		Func func() (err error)
		Name string
	}

	inserts := []InsertFunc{
		{
			Func: s.insertTestAccount,
			Name: "Account",
		},
		{
			Func: s.insertTestCharacter,
			Name: "Character",
		},
		{
			Func: s.insertTestForum,
			Name: "Forum",
		},
		{
			Func: s.insertTestNpc,
			Name: "Npc",
		},
		{
			Func: s.insertTestRule,
			Name: "Rule",
		},
		{
			Func: s.insertTestRuleEntry,
			Name: "RuleEntry",
		},
		{
			Func: s.insertTestUser,
			Name: "User",
		},
		{
			Func: s.insertTestUserAccount,
			Name: "UserAccount",
		},
		/*{
			Func: s.insertTestUserLink,
			Name: "UserLink",
		},*/
		{
			Func: s.insertTestVariable,
			Name: "Variable",
		},
		{
			Func: s.insertTestZone,
			Name: "Zone",
		},
	}

	s.log.Println("inserting", len(inserts))
	for _, insert := range inserts {
		err = insert.Func()
		if err != nil {
			return
		}
	}
	err = nil
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
			Func: s.createTableForum,
			Name: "Forum",
		},
		{
			Func: s.createTableNpc,
			Name: "Npc",
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
			Func: s.createTableUserAccount,
			Name: "UserAccount",
		},
		{
			Func: s.createTableUserLink,
			Name: "UserLink",
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
