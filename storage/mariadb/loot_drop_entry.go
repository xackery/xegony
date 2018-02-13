package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	lootDropEntryTable = "lootdrop_entries"

	lootDropEntryFields = "lootdrop_id, item_id, item_charges, equip_item, chance, disabled_chance, minlevel, maxlevel, multiplier"
	lootDropEntryBinds  = ":lootdrop_id, :item_id, :item_charges, :equip_item, :chance, :disabled_chance, :minlevel, :maxlevel, :multiplier"
)

//GetLootDropEntry will grab data from storage
func (s *Storage) GetLootDropEntry(lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE lootdrop_id = ? AND item_id = ?", lootDropEntryFields, lootDropEntryTable)
	err = s.db.Get(lootDropEntry, query, lootDrop.ID, lootDropEntry.ItemID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateLootDropEntry will grab data from storage
func (s *Storage) CreateLootDropEntry(lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", lootDropEntryTable, lootDropEntryFields, lootDropEntryBinds)
	_, err = s.db.NamedExec(query, lootDropEntry)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListLootDropEntry will grab data from storage
func (s *Storage) ListLootDropEntry(page *model.Page, lootDrop *model.LootDrop) (lootDropEntrys []*model.LootDropEntry, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "item_id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE lootdrop_id = ? ORDER BY %s LIMIT %d OFFSET %d", lootDropEntryFields, lootDropEntryTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query, lootDrop.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		lootDropEntry := model.LootDropEntry{}
		if err = rows.StructScan(&lootDropEntry); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		lootDropEntrys = append(lootDropEntrys, &lootDropEntry)
	}
	return
}

//ListLootDropEntryTotalCount will grab data from storage
func (s *Storage) ListLootDropEntryTotalCount(lootDrop *model.LootDrop) (count int64, err error) {
	query := fmt.Sprintf("SELECT count(lootdrop_id) FROM %s WHERE lootdrop_id = ?", lootDropEntryTable)
	err = s.db.Get(&count, query, lootDrop.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListLootDropEntryBySearch will grab data from storage
func (s *Storage) ListLootDropEntryBySearch(page *model.Page, lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry) (lootDropEntrys []*model.LootDropEntry, err error) {

	field := ""

	if lootDropEntry.LootDropID > 0 {
		field += `lootdrop_id = :lootdrop_id OR`
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]
	lootDropEntry.LootDropID = lootDrop.ID

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", lootDropEntryFields, lootDropEntryTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, lootDropEntry)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		lootDropEntry := model.LootDropEntry{}
		if err = rows.StructScan(&lootDropEntry); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		lootDropEntrys = append(lootDropEntrys, &lootDropEntry)
	}
	return
}

//ListLootDropEntryBySearchTotalCount will grab data from storage
func (s *Storage) ListLootDropEntryBySearchTotalCount(lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry) (count int64, err error) {
	field := ""
	if lootDropEntry.LootDropID > 0 {
		field += `lootdrop_id = :lootdrop_id OR`
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	lootDropEntry.LootDropID = lootDrop.ID
	query := fmt.Sprintf("SELECT count(lootdrop_id) FROM %s WHERE %s", lootDropEntryTable, field)

	rows, err := s.db.NamedQuery(query, lootDropEntry)
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

//EditLootDropEntry will grab data from storage
func (s *Storage) EditLootDropEntry(lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry) (err error) {

	prevLootDropEntry := &model.LootDropEntry{
		LootDropID: lootDropEntry.LootDropID,
		ItemID:     lootDropEntry.ItemID,
	}
	err = s.GetLootDropEntry(lootDrop, prevLootDropEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous lootDropEntry")
		return
	}

	field := ""

	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE lootdrop_id = :lootdrop_id AND id = :id", lootDropEntryTable, field)
	result, err := s.db.NamedExec(query, lootDropEntry)
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

//DeleteLootDropEntry will grab data from storage
func (s *Storage) DeleteLootDropEntry(lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry) (err error) {

	query := fmt.Sprintf("DELETE FROM %s WHERE lootdrop_id = ? AND lootdrop_id = ?", lootDropEntryTable)
	result, err := s.db.Exec(query, lootDrop.ID, lootDropEntry.LootDropID)
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

//createTableLootDropEntry will grab data from storage
func (s *Storage) createTableLootDropEntry() (err error) {
	_, err = s.db.Exec(`
		CREATE TABLE lootdrop_entries (
			lootdrop_id int(11) unsigned NOT NULL DEFAULT '0',
			item_id int(11) NOT NULL DEFAULT '0',
			item_charges smallint(2) unsigned NOT NULL DEFAULT '1',
			equip_item tinyint(2) unsigned NOT NULL DEFAULT '0',
			chance float NOT NULL DEFAULT '1',
			disabled_chance float NOT NULL DEFAULT '0',
			minlevel tinyint(3) NOT NULL DEFAULT '0',
			maxlevel tinyint(3) NOT NULL DEFAULT '127',
			multiplier tinyint(2) unsigned NOT NULL DEFAULT '1',
			PRIMARY KEY (lootdrop_id,item_id)
		  ) ENGINE=MyISAM DEFAULT CHARSET=latin1;
		`)
	if err != nil {
		return
	}
	return
}
