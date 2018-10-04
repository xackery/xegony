package mariadb

import (
	"fmt"

	//required for sqlx
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Storage implements mariadb
type Storage struct {
	db *sqlx.DB
}

// New returns a new storage implementation
func New() (s *Storage, err error) {
	s = &Storage{}
	return
}

// Connect implements storage
func (s *Storage) Connect(host string, username string, password string, database string) (err error) {
	s.db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s", username, password, host, database))
	if err != nil {
		err = errors.Wrap(err, "failed to connect")
		return
	}

	return
}

// Close implements storage
func (s *Storage) Close() (err error) {
	return
}
