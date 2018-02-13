package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	lootDropTable  = "lootdrop"
	lootDropFields = "id, name"
	lootDropBinds  = ":id, :name"
)

//GetLootDrop will grab data from storage
func (s *Storage) GetLootDrop(lootDrop *model.LootDrop) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", lootDropFields, lootDropTable)
	err = s.db.Get(lootDrop, query, lootDrop.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateLootDrop will grab data from storage
func (s *Storage) CreateLootDrop(lootDrop *model.LootDrop) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", lootDropTable, lootDropFields, lootDropBinds)
	result, err := s.db.NamedExec(query, lootDrop)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	lootDropID, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	lootDrop.ID = lootDropID
	return
}

//ListLootDrop will grab data from storage
func (s *Storage) ListLootDrop(page *model.Page) (lootDrops []*model.LootDrop, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", lootDropFields, lootDropTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		lootDrop := model.LootDrop{}
		if err = rows.StructScan(&lootDrop); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		lootDrops = append(lootDrops, &lootDrop)
	}
	return
}

//ListLootDropTotalCount will grab data from storage
func (s *Storage) ListLootDropTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(id) FROM %s", lootDropTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListLootDropBySearch will grab data from storage
func (s *Storage) ListLootDropBySearch(page *model.Page, lootDrop *model.LootDrop) (lootDrops []*model.LootDrop, err error) {

	field := ""

	if len(lootDrop.Name) > 0 {
		field += `name LIKE :name OR`
		lootDrop.Name = fmt.Sprintf("%%%s%%", lootDrop.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", lootDropFields, lootDropTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, lootDrop)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		lootDrop := model.LootDrop{}
		if err = rows.StructScan(&lootDrop); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		lootDrops = append(lootDrops, &lootDrop)
	}
	return
}

//ListLootDropBySearchTotalCount will grab data from storage
func (s *Storage) ListLootDropBySearchTotalCount(lootDrop *model.LootDrop) (count int64, err error) {
	field := ""
	if len(lootDrop.Name) > 0 {
		field += `name LIKE :name OR`
		lootDrop.Name = fmt.Sprintf("%%%s%%", lootDrop.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", lootDropTable, field)

	rows, err := s.db.NamedQuery(query, lootDrop)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
	}
	return
}

//EditLootDrop will grab data from storage
func (s *Storage) EditLootDrop(lootDrop *model.LootDrop) (err error) {

	prevLootDrop := &model.LootDrop{
		ID: lootDrop.ID,
	}
	err = s.GetLootDrop(prevLootDrop)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous lootDrop")
		return
	}

	field := ""
	if len(lootDrop.Name) > 0 && prevLootDrop.Name != lootDrop.Name {
		field += "name = :name, "
	}
	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", lootDropTable, field)
	result, err := s.db.NamedExec(query, lootDrop)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//DeleteLootDrop will grab data from storage
func (s *Storage) DeleteLootDrop(lootDrop *model.LootDrop) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", lootDropTable)
	result, err := s.db.Exec(query, lootDrop.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//createTableLootDrop will grab data from storage
func (s *Storage) createTableLootDrop() (err error) {
	_, err = s.db.Exec(`
		CREATE TABLE lootdrop (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			name varchar(255) NOT NULL DEFAULT '',
			PRIMARY KEY (id)
		  ) ENGINE=MyISAM AUTO_INCREMENT=165863 DEFAULT CHARSET=latin1 PACK_KEYS=0;
		`)
	if err != nil {
		return
	}
	return
}
