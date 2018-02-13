package memory

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetLootEntry will grab data from storage
func (s *Storage) GetLootEntry(loot *model.Loot, lootEntry *model.LootEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateLootEntry will grab data from storage
func (s *Storage) CreateLootEntry(loot *model.Loot, lootEntry *model.LootEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootEntry will grab data from storage
func (s *Storage) ListLootEntry(page *model.Page, loot *model.Loot) (lootEntrys []*model.LootEntry, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootEntryTotalCount will grab data from storage
func (s *Storage) ListLootEntryTotalCount(loot *model.Loot) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootEntryBySearch will grab data from storage
func (s *Storage) ListLootEntryBySearch(page *model.Page, loot *model.Loot, lootEntry *model.LootEntry) (lootEntrys []*model.LootEntry, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootEntryBySearchTotalCount will grab data from storage
func (s *Storage) ListLootEntryBySearchTotalCount(loot *model.Loot, lootEntry *model.LootEntry) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditLootEntry will grab data from storage
func (s *Storage) EditLootEntry(loot *model.Loot, lootEntry *model.LootEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteLootEntry will grab data from storage
func (s *Storage) DeleteLootEntry(loot *model.Loot, lootEntry *model.LootEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
