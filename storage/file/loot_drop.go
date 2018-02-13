package file

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetLootDrop will grab data from storage
func (s *Storage) GetLootDrop(lootDrop *model.LootDrop) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateLootDrop will grab data from storage
func (s *Storage) CreateLootDrop(lootDrop *model.LootDrop) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootDrop will grab data from storage
func (s *Storage) ListLootDrop(page *model.Page) (lootDrops []*model.LootDrop, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootDropTotalCount will grab data from storage
func (s *Storage) ListLootDropTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootDropBySearch will grab data from storage
func (s *Storage) ListLootDropBySearch(page *model.Page, lootDrop *model.LootDrop) (lootDrops []*model.LootDrop, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootDropBySearchTotalCount will grab data from storage
func (s *Storage) ListLootDropBySearchTotalCount(lootDrop *model.LootDrop) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditLootDrop will grab data from storage
func (s *Storage) EditLootDrop(lootDrop *model.LootDrop) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteLootDrop will grab data from storage
func (s *Storage) DeleteLootDrop(lootDrop *model.LootDrop) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
