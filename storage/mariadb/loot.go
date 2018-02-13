package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	lootTable  = "loottable"
	lootFields = "id, name, mincash, maxcash, avgcoin, done"
	lootBinds  = ":id, :name, :mincash, :maxcash, :avgcoin, :done"
)

//GetLoot will grab data from storage
func (s *Storage) GetLoot(loot *model.Loot) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", lootFields, lootTable)
	err = s.db.Get(loot, query, loot.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateLoot will grab data from storage
func (s *Storage) CreateLoot(loot *model.Loot) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", lootTable, lootFields, lootBinds)
	result, err := s.db.NamedExec(query, loot)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	lootID, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	loot.ID = lootID
	return
}

//ListLoot will grab data from storage
func (s *Storage) ListLoot(page *model.Page) (loots []*model.Loot, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", lootFields, lootTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		loot := model.Loot{}
		if err = rows.StructScan(&loot); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		loots = append(loots, &loot)
	}
	return
}

//ListLootTotalCount will grab data from storage
func (s *Storage) ListLootTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(id) FROM %s", lootTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListLootBySearch will grab data from storage
func (s *Storage) ListLootBySearch(page *model.Page, loot *model.Loot) (loots []*model.Loot, err error) {

	field := ""

	if len(loot.Name) > 0 {
		field += `name LIKE :name OR`
		loot.Name = fmt.Sprintf("%%%s%%", loot.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", lootFields, lootTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, loot)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		loot := model.Loot{}
		if err = rows.StructScan(&loot); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		loots = append(loots, &loot)
	}
	return
}

//ListLootBySearchTotalCount will grab data from storage
func (s *Storage) ListLootBySearchTotalCount(loot *model.Loot) (count int64, err error) {
	field := ""
	if len(loot.Name) > 0 {
		field += `name LIKE :name OR`
		loot.Name = fmt.Sprintf("%%%s%%", loot.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", lootTable, field)

	rows, err := s.db.NamedQuery(query, loot)
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

//EditLoot will grab data from storage
func (s *Storage) EditLoot(loot *model.Loot) (err error) {

	prevLoot := &model.Loot{
		ID: loot.ID,
	}
	err = s.GetLoot(prevLoot)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous loot")
		return
	}

	field := ""
	if len(loot.Name) > 0 && prevLoot.Name != loot.Name {
		field += "name = :name, "
	}
	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", lootTable, field)
	result, err := s.db.NamedExec(query, loot)
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

//DeleteLoot will grab data from storage
func (s *Storage) DeleteLoot(loot *model.Loot) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", lootTable)
	result, err := s.db.Exec(query, loot.ID)
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

//createTableLoot will grab data from storage
func (s *Storage) createTableLoot() (err error) {
	_, err = s.db.Exec(`
	CREATE TABLE loottable (
		id int(11) unsigned NOT NULL AUTO_INCREMENT,
		name varchar(255) NOT NULL DEFAULT '',
		mincash int(11) unsigned NOT NULL DEFAULT '0',
		maxcash int(11) unsigned NOT NULL DEFAULT '0',
		avgcoin int(10) unsigned NOT NULL DEFAULT '0',
		done tinyint(3) NOT NULL DEFAULT '0',
		PRIMARY KEY (id)
		) ENGINE=MyISAM AUTO_INCREMENT=105101 DEFAULT CHARSET=latin1 PACK_KEYS=0;`)
	if err != nil {
		return
	}
	return
}
