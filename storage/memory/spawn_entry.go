package memory

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSpawnEntry will grab data from storage
func (s *Storage) GetSpawnEntry(spawn *model.Spawn, spawnEntry *model.SpawnEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSpawnEntry will grab data from storage
func (s *Storage) CreateSpawnEntry(spawn *model.Spawn, spawnEntry *model.SpawnEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpawnEntry will grab data from storage
func (s *Storage) ListSpawnEntry(page *model.Page, spawn *model.Spawn) (spawnEntrys []*model.SpawnEntry, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpawnEntryTotalCount will grab data from storage
func (s *Storage) ListSpawnEntryTotalCount(spawn *model.Spawn) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpawnEntryBySearch will grab data from storage
func (s *Storage) ListSpawnEntryBySearch(page *model.Page, spawn *model.Spawn, spawnEntry *model.SpawnEntry) (spawnEntrys []*model.SpawnEntry, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpawnEntryBySearchTotalCount will grab data from storage
func (s *Storage) ListSpawnEntryBySearchTotalCount(spawn *model.Spawn, spawnEntry *model.SpawnEntry) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSpawnEntry will grab data from storage
func (s *Storage) EditSpawnEntry(spawn *model.Spawn, spawnEntry *model.SpawnEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSpawnEntry will grab data from storage
func (s *Storage) DeleteSpawnEntry(spawn *model.Spawn, spawnEntry *model.SpawnEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
