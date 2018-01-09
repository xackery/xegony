package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

var (
	bazaarFields = `accountid, itemid, price`
	bazaarSets   = `accountid=:accountid, itemid=:itemid, price=:price`
	bazaarBinds  = `:accountid, :itemid, :price`
)

//GetBazaar will grab data from storage
func (s *Storage) GetBazaar(bazaarID int64) (bazaar *model.Bazaar, err error) {
	bazaar = &model.Bazaar{}
	query := fmt.Sprintf(`SELECT id, %s FROM bazaar WHERE id = ?`, bazaarFields)
	err = s.db.Get(bazaar, query, bazaarID)
	if err != nil {
		return
	}
	return
}

//CreateBazaar will grab data from storage
func (s *Storage) CreateBazaar(bazaar *model.Bazaar) (err error) {
	if bazaar == nil {
		err = fmt.Errorf("Must provide bazaar")
		return
	}

	query := fmt.Sprintf(`INSERT INTO bazaar(%s)
		VALUES (%s)`, bazaarFields, bazaarBinds)
	result, err := s.db.NamedExec(query, bazaar)
	if err != nil {
		return
	}
	bazaarID, err := result.LastInsertId()
	if err != nil {
		return
	}
	bazaar.ID = bazaarID
	return
}

//ListBazaar will grab data from storage
func (s *Storage) ListBazaar() (bazaars []*model.Bazaar, err error) {
	query := fmt.Sprintf(`SELECT id, %s FROM bazaar ORDER BY id DESC`, bazaarFields)
	rows, err := s.db.Queryx(query)
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

//EditBazaar will grab data from storage
func (s *Storage) EditBazaar(bazaarID int64, bazaar *model.Bazaar) (err error) {
	bazaar.ID = bazaarID
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE bazaar SET %s WHERE id = :id`, bazaarSets), bazaar)
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

//DeleteBazaar will grab data from storage
func (s *Storage) DeleteBazaar(bazaarID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM bazaar WHERE id = ?`, bazaarID)
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

//createTableBazaar will grab data from storage
func (s *Storage) createTableBazaar() (err error) {
	_, err = s.db.Exec(`CREATE TABLE if NOT EXISTS bazaar (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  accountid int(11) unsigned NOT NULL,
  itemid int(11) unsigned NOT NULL,
  price int(11) unsigned NOT NULL,
  createdate timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	return
}
