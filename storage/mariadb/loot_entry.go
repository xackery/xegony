package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	lootEntryTable = "loottable_entries"

	lootEntryFields = "loottable_id, lootdrop_id, multiplier, droplimit, mindrop, probability"
	lootEntryBinds  = ":loottable_id, :lootdrop_id, :multiplier, :droplimit, :mindrop, :probability,"
)

//GetLootEntry will grab data from storage
func (s *Storage) GetLootEntry(loot *model.Loot, lootEntry *model.LootEntry) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE loottable_id = ? AND lootdrop_id = ?", lootEntryFields, lootEntryTable)
	err = s.db.Get(lootEntry, query, loot.ID, lootEntry.LootDropID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateLootEntry will grab data from storage
func (s *Storage) CreateLootEntry(loot *model.Loot, lootEntry *model.LootEntry) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", lootEntryTable, lootEntryFields, lootEntryBinds)
	_, err = s.db.NamedExec(query, lootEntry)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListLootEntry will grab data from storage
func (s *Storage) ListLootEntry(page *model.Page, loot *model.Loot) (lootEntrys []*model.LootEntry, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "lootdrop_id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE loottable_id = ? ORDER BY %s LIMIT %d OFFSET %d", lootEntryFields, lootEntryTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query, loot.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		lootEntry := model.LootEntry{}
		if err = rows.StructScan(&lootEntry); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		lootEntrys = append(lootEntrys, &lootEntry)
	}
	return
}

//ListLootEntryTotalCount will grab data from storage
func (s *Storage) ListLootEntryTotalCount(loot *model.Loot) (count int64, err error) {
	query := fmt.Sprintf("SELECT count(loottable_id) FROM %s WHERE loottable_id = ?", lootEntryTable)
	err = s.db.Get(&count, query, loot.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListLootEntryBySearch will grab data from storage
func (s *Storage) ListLootEntryBySearch(page *model.Page, loot *model.Loot, lootEntry *model.LootEntry) (lootEntrys []*model.LootEntry, err error) {

	field := ""

	if lootEntry.LootID > 0 {
		field += `loottable_id = :loottable_id OR`
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]
	lootEntry.LootID = loot.ID

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", lootEntryFields, lootEntryTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, lootEntry)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		lootEntry := model.LootEntry{}
		if err = rows.StructScan(&lootEntry); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		lootEntrys = append(lootEntrys, &lootEntry)
	}
	return
}

//ListLootEntryBySearchTotalCount will grab data from storage
func (s *Storage) ListLootEntryBySearchTotalCount(loot *model.Loot, lootEntry *model.LootEntry) (count int64, err error) {
	field := ""
	if lootEntry.LootID > 0 {
		field += `loottable_id = :loottable_id OR`
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	lootEntry.LootID = loot.ID
	query := fmt.Sprintf("SELECT count(loottable_id) FROM %s WHERE %s", lootEntryTable, field)

	rows, err := s.db.NamedQuery(query, lootEntry)
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

//EditLootEntry will grab data from storage
func (s *Storage) EditLootEntry(loot *model.Loot, lootEntry *model.LootEntry) (err error) {

	prevLootEntry := &model.LootEntry{
		LootID:     lootEntry.LootID,
		LootDropID: lootEntry.LootDropID,
	}
	err = s.GetLootEntry(loot, prevLootEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous lootEntry")
		return
	}

	field := ""

	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE loottable_id = :loottable_id AND id = :id", lootEntryTable, field)
	result, err := s.db.NamedExec(query, lootEntry)
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

//DeleteLootEntry will grab data from storage
func (s *Storage) DeleteLootEntry(loot *model.Loot, lootEntry *model.LootEntry) (err error) {

	query := fmt.Sprintf("DELETE FROM %s WHERE loottable_id = ? AND lootdrop_id = ?", lootEntryTable)
	result, err := s.db.Exec(query, loot.ID, lootEntry.LootDropID)
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

//createTableLootEntry will grab data from storage
func (s *Storage) createTableLootEntry() (err error) {
	_, err = s.db.Exec(`
		CREATE TABLE loottable_entries (
			loottable_id int(11) unsigned NOT NULL DEFAULT '0',
			lootdrop_id int(11) unsigned NOT NULL DEFAULT '0',
			multiplier tinyint(2) unsigned NOT NULL DEFAULT '1',
			droplimit tinyint(2) unsigned NOT NULL DEFAULT '0',
			mindrop tinyint(2) unsigned NOT NULL DEFAULT '0',
			probability float NOT NULL DEFAULT '100',
			PRIMARY KEY (loottable_id,lootdrop_id)
		  ) ENGINE=MyISAM DEFAULT CHARSET=latin1;
		`)
	if err != nil {
		return
	}
	return
}
