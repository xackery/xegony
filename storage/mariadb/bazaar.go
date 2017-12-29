package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

func (s *Storage) CreateTableBazaar() (err error) {
	return
	_, err = s.db.Exec(`CREATE TABLE if NOT EXISTS forum (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(32) NOT NULL DEFAULT '',
  owner_id int(11) unsigned NOT NULL,
  description varchar(128) NOT NULL DEFAULT '',
  last_modified timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  create_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  icon varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	return
}

func (s *Storage) GetBazaar(bazaarId int64) (bazaar *model.Bazaar, err error) {
	bazaar = &model.Bazaar{}
	err = s.db.Get(bazaar, "SELECT id, name, itemid FROM bazaar WHERE id = ?", bazaarId)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateBazaar(bazaar *model.Bazaar) (err error) {
	if bazaar == nil {
		err = fmt.Errorf("Must provide bazaar")
		return
	}

	result, err := s.db.NamedExec(`INSERT INTO bazaar(name, itemid)
		VALUES (:name, :itemid)`, bazaar)
	if err != nil {
		return
	}
	bazaarId, err := result.LastInsertId()
	if err != nil {
		return
	}
	bazaar.Id = bazaarId
	return
}

func (s *Storage) ListBazaar() (bazaars []*model.Bazaar, err error) {
	rows, err := s.db.Queryx(`SELECT id, name, itemid FROM bazaar ORDER BY id DESC`)
	if err != nil {
		return
	}

	for rows.Next() {
		bazaar := model.Bazaar{}
		if err = rows.StructScan(&bazaar); err != nil {
			return
		}
		bazaars = append(bazaars, &bazaar)
	}
	return
}

func (s *Storage) EditBazaar(bazaarId int64, bazaar *model.Bazaar) (err error) {
	bazaar.Id = bazaarId
	result, err := s.db.NamedExec(`UPDATE bazaar SET name=:name, itemid=:itemid WHERE id = :id`, bazaar)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}

func (s *Storage) DeleteBazaar(bazaarId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM bazaar WHERE id = ?`, bazaarId)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}
