package memory

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetLootDropEntry will grab data from storage
func (s *Storage) GetLootDropEntry(lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateLootDropEntry will grab data from storage
func (s *Storage) CreateLootDropEntry(lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootDropEntry will grab data from storage
func (s *Storage) ListLootDropEntry(page *model.Page, lootDrop *model.LootDrop) (lootDropEntrys []*model.LootDropEntry, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootDropEntryTotalCount will grab data from storage
func (s *Storage) ListLootDropEntryTotalCount(lootDrop *model.LootDrop) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootDropEntryBySearch will grab data from storage
func (s *Storage) ListLootDropEntryBySearch(page *model.Page, lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry) (lootDropEntrys []*model.LootDropEntry, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootDropEntryBySearchTotalCount will grab data from storage
func (s *Storage) ListLootDropEntryBySearchTotalCount(lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditLootDropEntry will grab data from storage
func (s *Storage) EditLootDropEntry(lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteLootDropEntry will grab data from storage
func (s *Storage) DeleteLootDropEntry(lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
