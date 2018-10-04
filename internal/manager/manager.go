package manager

import (
	"github.com/pkg/errors"
	"github.com/xackery/xegony/internal/storage"
	"github.com/xackery/xegony/internal/storage/mariadb"
)

// Manager implements a manager instance
type Manager struct {
	db storage.Storager
}

// New creates a new Manager instance
func New() (m *Manager, err error) {
	m = &Manager{}
	m.db, err = mariadb.New()
	if err != nil {
		err = errors.Wrap(err, "failed to create mariadb")
		return
	}
	err = m.db.Connect("127.0.0.1:3306", "root", "", "eqemu")
	if err != nil {
		err = errors.Wrap(err, "failed to connect to mariadb")
		return
	}
	go m.pump()
	return
}
