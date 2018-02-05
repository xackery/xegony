package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	spawnsDatabase = []*model.Spawn{}
	spawnLock      = sync.RWMutex{}
)

//GetSpawn will grab data from storage
func (s *Storage) GetSpawn(spawn *model.Spawn) (err error) {
	spawnLock.RLock()
	defer spawnLock.RUnlock()
	for _, tmpSpawn := range spawnsDatabase {
		if tmpSpawn.ID == spawn.ID {
			*spawn = *tmpSpawn
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSpawn will grab data from storage
func (s *Storage) CreateSpawn(spawn *model.Spawn) (err error) {
	spawnLock.Lock()
	defer spawnLock.Unlock()
	for _, tmpSpawn := range spawnsDatabase {
		if tmpSpawn.ID == spawn.ID {
			err = fmt.Errorf("spawn already exists")
			return
		}
	}
	spawnsDatabase = append(spawnsDatabase, spawn)
	return
}

//ListSpawn will grab data from storage
func (s *Storage) ListSpawn(page *model.Page) (spawns []*model.Spawn, err error) {
	spawnLock.RLock()
	defer spawnLock.RUnlock()

	spawns = make([]*model.Spawn, len(spawnsDatabase))

	spawns = spawnsDatabase

	switch page.OrderBy {
	case "short_name":
		sort.Slice(spawns, func(i, j int) bool {
			return spawns[i].Name < spawns[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(spawns))
		}
	*/
	return
}

//ListSpawnTotalCount will grab data from storage
func (s *Storage) ListSpawnTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(spawnsDatabase))
	return
}

//ListSpawnBySearch will grab data from storage
func (s *Storage) ListSpawnBySearch(page *model.Page, spawn *model.Spawn) (spawns []*model.Spawn, err error) {
	spawnLock.RLock()
	defer spawnLock.RUnlock()

	if len(spawn.Name) > 0 {
		for i := range spawnsDatabase {
			if strings.Contains(spawnsDatabase[i].Name, spawn.Name) {
				spawns = append(spawns, spawnsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "short_name":
		sort.Slice(spawns, func(i, j int) bool {
			return spawns[i].Name < spawns[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(spawns))
	//}
	return
}

//ListSpawnBySearchTotalCount will grab data from storage
func (s *Storage) ListSpawnBySearchTotalCount(spawn *model.Spawn) (count int64, err error) {
	spawnLock.RLock()
	defer spawnLock.RUnlock()

	spawns := []*model.Spawn{}
	if len(spawn.Name) > 0 {
		for i := range spawnsDatabase {
			if strings.Contains(spawnsDatabase[i].Name, spawn.Name) {
				spawns = append(spawns, spawnsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(spawns))
	return
}

//EditSpawn will grab data from storage
func (s *Storage) EditSpawn(spawn *model.Spawn) (err error) {
	spawnLock.Lock()
	defer spawnLock.Unlock()
	for i := range spawnsDatabase {
		if spawnsDatabase[i].ID == spawn.ID {
			*spawnsDatabase[i] = *spawn
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSpawn will grab data from storage
func (s *Storage) DeleteSpawn(spawn *model.Spawn) (err error) {
	spawnLock.Lock()
	defer spawnLock.Unlock()
	indexToDelete := 0
	for i := range spawnsDatabase {
		if spawnsDatabase[i].ID == spawn.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	spawnsDatabase[len(spawnsDatabase)-1], spawnsDatabase[indexToDelete] = spawnsDatabase[indexToDelete], spawnsDatabase[len(spawnsDatabase)-1]
	spawnsDatabase = spawnsDatabase[:len(spawnsDatabase)-1]
	return
}
