package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	lootsDatabase = []*model.Loot{}
	lootLock      = sync.RWMutex{}
)

//GetLoot will grab data from storage
func (s *Storage) GetLoot(loot *model.Loot) (err error) {
	lootLock.RLock()
	defer lootLock.RUnlock()
	for _, tmpLoot := range lootsDatabase {
		if tmpLoot.ID == loot.ID {
			*loot = *tmpLoot
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateLoot will grab data from storage
func (s *Storage) CreateLoot(loot *model.Loot) (err error) {
	lootLock.Lock()
	defer lootLock.Unlock()
	for _, tmpLoot := range lootsDatabase {
		if tmpLoot.ID == loot.ID {
			err = fmt.Errorf("loot already exists")
			return
		}
	}
	lootsDatabase = append(lootsDatabase, loot)
	return
}

//ListLoot will grab data from storage
func (s *Storage) ListLoot(page *model.Page) (loots []*model.Loot, err error) {
	lootLock.RLock()
	defer lootLock.RUnlock()

	loots = make([]*model.Loot, len(lootsDatabase))

	loots = lootsDatabase

	switch page.OrderBy {
	case "short_name":
		sort.Slice(loots, func(i, j int) bool {
			return loots[i].Name < loots[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(loots))
		}
	*/
	return
}

//ListLootTotalCount will grab data from storage
func (s *Storage) ListLootTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(lootsDatabase))
	return
}

//ListLootBySearch will grab data from storage
func (s *Storage) ListLootBySearch(page *model.Page, loot *model.Loot) (loots []*model.Loot, err error) {
	lootLock.RLock()
	defer lootLock.RUnlock()

	if len(loot.Name) > 0 {
		for i := range lootsDatabase {
			if strings.Contains(lootsDatabase[i].Name, loot.Name) {
				loots = append(loots, lootsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "short_name":
		sort.Slice(loots, func(i, j int) bool {
			return loots[i].Name < loots[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(loots))
	//}
	return
}

//ListLootBySearchTotalCount will grab data from storage
func (s *Storage) ListLootBySearchTotalCount(loot *model.Loot) (count int64, err error) {
	lootLock.RLock()
	defer lootLock.RUnlock()

	loots := []*model.Loot{}
	if len(loot.Name) > 0 {
		for i := range lootsDatabase {
			if strings.Contains(lootsDatabase[i].Name, loot.Name) {
				loots = append(loots, lootsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(loots))
	return
}

//EditLoot will grab data from storage
func (s *Storage) EditLoot(loot *model.Loot) (err error) {
	lootLock.Lock()
	defer lootLock.Unlock()
	for i := range lootsDatabase {
		if lootsDatabase[i].ID == loot.ID {
			*lootsDatabase[i] = *loot
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteLoot will grab data from storage
func (s *Storage) DeleteLoot(loot *model.Loot) (err error) {
	lootLock.Lock()
	defer lootLock.Unlock()
	indexToDelete := 0
	for i := range lootsDatabase {
		if lootsDatabase[i].ID == loot.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	lootsDatabase[len(lootsDatabase)-1], lootsDatabase[indexToDelete] = lootsDatabase[indexToDelete], lootsDatabase[len(lootsDatabase)-1]
	lootsDatabase = lootsDatabase[:len(lootsDatabase)-1]
	return
}
