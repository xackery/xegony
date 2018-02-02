package google

import (
	"io"
	alog "log"
	"os"
)

//Storage implements the storage interface
type Storage struct {
	log    *alog.Logger
	logErr *alog.Logger
}

// New creates a new storage entry
func New(token string, secret string, w io.Writer, wErr io.Writer) (db *Storage, err error) {
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

	return
}

//InsertTestData will grab data from storage
func (s *Storage) InsertTestData() (err error) {

	return
}

//DropTables will grab data from storage
func (s *Storage) DropTables() (err error) {

	return
}

//VerifyTables will grab data from storage
func (s *Storage) VerifyTables() (err error) {

	return
}
