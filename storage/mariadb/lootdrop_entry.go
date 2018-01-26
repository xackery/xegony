package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	lootDropEntrySets   = `lootdrop_id=:lootdrop_id, item_id=:item_id, item_charges=:item_charges, equip_item=:equip_item, chance=:chance, disabled_chance=:disabled_chance, minlevel=:minlevel, maxlevel=:maxlevel, multiplier=:multiplier`
	lootDropEntryFields = `lootdrop_id, item_id, item_charges, equip_item, chance, disabled_chance, minlevel, maxlevel, multiplier`
	lootDropEntryBinds  = `:lootdrop_id, :item_id, :item_charges, :equip_item, :chance, :disabled_chance, :minlevel, :maxlevel, :multiplier,`
)

//GetLootDropEntry will grab data from storage
func (s *Storage) GetLootDropEntry(lootDropEntry *model.LootDropEntry) (err error) {
	err = s.db.Get(lootDropEntry, fmt.Sprintf("SELECT %s FROM lootdrop_entries WHERE lootdrop_id = ? AND item_id = ?", lootDropEntryFields), lootDropEntry.LootdropID, lootDropEntry.ItemID)
	if err != nil {
		return
	}
	return
}

//CreateLootDropEntry will grab data from storage
func (s *Storage) CreateLootDropEntry(lootDropEntry *model.LootDropEntry) (err error) {
	if lootDropEntry == nil {
		err = fmt.Errorf("Must provide lootDropEntry")
		return
	}

	_, err = s.db.NamedExec(fmt.Sprintf(`INSERT INTO lootdrop_entries(%s)
		VALUES (%s)`, lootDropEntryFields, lootDropEntryBinds), lootDropEntry)
	if err != nil {
		return
	}

	return
}

//ListLootDropEntryByLootDrop will grab data from storage
func (s *Storage) ListLootDropEntryByLootDrop(lootDrop *model.LootDrop) (lootDropEntrys []*model.LootDropEntry, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM lootdrop_entries WHERE lootdrop_id = ?`, lootDropEntryFields), lootDrop.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		lootDropEntry := model.LootDropEntry{}
		if err = rows.StructScan(&lootDropEntry); err != nil {
			return
		}
		lootDropEntrys = append(lootDropEntrys, &lootDropEntry)
	}
	return
}

//EditLootDropEntry will grab data from storage
func (s *Storage) EditLootDropEntry(lootDropEntry *model.LootDropEntry) (err error) {
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE lootdrop_entries SET %s WHERE lootdrop_id = :lootdrop_id AND item_id = :item_id`, lootDropEntrySets), lootDropEntry)
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

//DeleteLootDropEntry will grab data from storage
func (s *Storage) DeleteLootDropEntry(lootDropEntry *model.LootDropEntry) (err error) {
	result, err := s.db.Exec(`DELETE FROM lootdrop_entries WHERE lootdrop_id = ? AND item_id = ?`, lootDropEntry.LootdropID, lootDropEntry.ItemID)
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
