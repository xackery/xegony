package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	lootTableEntrySets   = `loottable_id=:loottable_id, lootdrop_id=:lootdrop_id, multiplier=:multiplier, droplimit=:droplimit, mindrop=:mindrop, probability=:probability`
	lootTableEntryFields = `loottable_id, lootdrop_id, multiplier, droplimit, mindrop, probability`
	lootTableEntryBinds  = `:loottable_id, :lootdrop_id, :multiplier, :droplimit, :mindrop, :probability`
)

func (s *Storage) GetLootTableEntry(lootTableID int64, lootDropID int64) (lootTableEntry *model.LootTableEntry, err error) {
	lootTableEntry = &model.LootTableEntry{}
	err = s.db.Get(lootTableEntry, fmt.Sprintf("SELECT %s FROM loottable_entries WHERE loottable_id = ? AND lootdrop_id = ?", lootTableEntryFields), lootTableID, lootDropID)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateLootTableEntry(lootTableEntry *model.LootTableEntry) (err error) {
	if lootTableEntry == nil {
		err = fmt.Errorf("Must provide lootTableEntry")
		return
	}

	_, err = s.db.NamedExec(fmt.Sprintf(`INSERT INTO loottable_entries(%s)
		VALUES (%s)`, lootTableEntryFields, lootTableEntryBinds), lootTableEntry)
	if err != nil {
		return
	}

	return
}

func (s *Storage) ListLootTableEntry(lootTableID int64) (lootTableEntrys []*model.LootTableEntry, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM loottable_entries WHERE loottable_id = ?`, lootTableEntryFields), lootTableID)
	if err != nil {
		return
	}

	for rows.Next() {
		lootTableEntry := model.LootTableEntry{}
		if err = rows.StructScan(&lootTableEntry); err != nil {
			return
		}
		lootTableEntrys = append(lootTableEntrys, &lootTableEntry)
	}
	return
}

func (s *Storage) EditLootTableEntry(lootTableID int64, lootDropID int64, lootTableEntry *model.LootTableEntry) (err error) {

	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE loottable_entries SET %s WHERE loottable_id = :loottable_id AND lootdrop_id = :lootdrop_id`, lootTableEntrySets), lootTableEntry)
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

func (s *Storage) DeleteLootTableEntry(lootTableID int64, lootDropID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM loottable_entries WHERE loottable_id = ?  and lootdrop_id = ?`, lootTableID, lootDropID)
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
