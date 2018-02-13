package file

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetLoot will grab data from storage
func (s *Storage) GetLoot(loot *model.Loot) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateLoot will grab data from storage
func (s *Storage) CreateLoot(loot *model.Loot) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLoot will grab data from storage
func (s *Storage) ListLoot(page *model.Page) (loots []*model.Loot, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootTotalCount will grab data from storage
func (s *Storage) ListLootTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootBySearch will grab data from storage
func (s *Storage) ListLootBySearch(page *model.Page, loot *model.Loot) (loots []*model.Loot, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListLootBySearchTotalCount will grab data from storage
func (s *Storage) ListLootBySearchTotalCount(loot *model.Loot) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditLoot will grab data from storage
func (s *Storage) EditLoot(loot *model.Loot) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteLoot will grab data from storage
func (s *Storage) DeleteLoot(loot *model.Loot) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
