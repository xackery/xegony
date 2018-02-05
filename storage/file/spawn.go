package file

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSpawn will grab data from storage
func (s *Storage) GetSpawn(spawn *model.Spawn) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSpawn will grab data from storage
func (s *Storage) CreateSpawn(spawn *model.Spawn) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpawn will grab data from storage
func (s *Storage) ListSpawn(page *model.Page) (spawns []*model.Spawn, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpawnTotalCount will grab data from storage
func (s *Storage) ListSpawnTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpawnBySearch will grab data from storage
func (s *Storage) ListSpawnBySearch(page *model.Page, spawn *model.Spawn) (spawns []*model.Spawn, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpawnBySearchTotalCount will grab data from storage
func (s *Storage) ListSpawnBySearchTotalCount(spawn *model.Spawn) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSpawn will grab data from storage
func (s *Storage) EditSpawn(spawn *model.Spawn) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSpawn will grab data from storage
func (s *Storage) DeleteSpawn(spawn *model.Spawn) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
