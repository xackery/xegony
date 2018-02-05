package memory

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSpawnNpc will grab data from storage
func (s *Storage) GetSpawnNpc(spawn *model.Spawn, spawnNpc *model.SpawnNpc) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSpawnNpc will grab data from storage
func (s *Storage) CreateSpawnNpc(spawn *model.Spawn, spawnNpc *model.SpawnNpc) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpawnNpc will grab data from storage
func (s *Storage) ListSpawnNpc(page *model.Page, spawn *model.Spawn) (spawnNpcs []*model.SpawnNpc, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpawnNpcTotalCount will grab data from storage
func (s *Storage) ListSpawnNpcTotalCount(spawn *model.Spawn) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpawnNpcBySearch will grab data from storage
func (s *Storage) ListSpawnNpcBySearch(page *model.Page, spawn *model.Spawn, spawnNpc *model.SpawnNpc) (spawnNpcs []*model.SpawnNpc, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpawnNpcBySearchTotalCount will grab data from storage
func (s *Storage) ListSpawnNpcBySearchTotalCount(spawn *model.Spawn, spawnNpc *model.SpawnNpc) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSpawnNpc will grab data from storage
func (s *Storage) EditSpawnNpc(spawn *model.Spawn, spawnNpc *model.SpawnNpc) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSpawnNpc will grab data from storage
func (s *Storage) DeleteSpawnNpc(spawn *model.Spawn, spawnNpc *model.SpawnNpc) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
