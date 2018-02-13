package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	lootDropsDatabase = []*model.LootDrop{}
	lootDropLock      = sync.RWMutex{}
)

//GetLootDrop will grab data from storage
func (s *Storage) GetLootDrop(lootDrop *model.LootDrop) (err error) {
	lootDropLock.RLock()
	defer lootDropLock.RUnlock()
	for _, tmpLootDrop := range lootDropsDatabase {
		if tmpLootDrop.ID == lootDrop.ID {
			*lootDrop = *tmpLootDrop
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateLootDrop will grab data from storage
func (s *Storage) CreateLootDrop(lootDrop *model.LootDrop) (err error) {
	lootDropLock.Lock()
	defer lootDropLock.Unlock()
	for _, tmpLootDrop := range lootDropsDatabase {
		if tmpLootDrop.ID == lootDrop.ID {
			err = fmt.Errorf("lootDrop already exists")
			return
		}
	}
	lootDropsDatabase = append(lootDropsDatabase, lootDrop)
	return
}

//ListLootDrop will grab data from storage
func (s *Storage) ListLootDrop(page *model.Page) (lootDrops []*model.LootDrop, err error) {
	lootDropLock.RLock()
	defer lootDropLock.RUnlock()

	lootDrops = make([]*model.LootDrop, len(lootDropsDatabase))

	lootDrops = lootDropsDatabase

	switch page.OrderBy {
	case "short_name":
		sort.Slice(lootDrops, func(i, j int) bool {
			return lootDrops[i].Name < lootDrops[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(lootDrops))
		}
	*/
	return
}

//ListLootDropTotalCount will grab data from storage
func (s *Storage) ListLootDropTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(lootDropsDatabase))
	return
}

//ListLootDropBySearch will grab data from storage
func (s *Storage) ListLootDropBySearch(page *model.Page, lootDrop *model.LootDrop) (lootDrops []*model.LootDrop, err error) {
	lootDropLock.RLock()
	defer lootDropLock.RUnlock()

	if len(lootDrop.Name) > 0 {
		for i := range lootDropsDatabase {
			if strings.Contains(lootDropsDatabase[i].Name, lootDrop.Name) {
				lootDrops = append(lootDrops, lootDropsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "short_name":
		sort.Slice(lootDrops, func(i, j int) bool {
			return lootDrops[i].Name < lootDrops[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(lootDrops))
	//}
	return
}

//ListLootDropBySearchTotalCount will grab data from storage
func (s *Storage) ListLootDropBySearchTotalCount(lootDrop *model.LootDrop) (count int64, err error) {
	lootDropLock.RLock()
	defer lootDropLock.RUnlock()

	lootDrops := []*model.LootDrop{}
	if len(lootDrop.Name) > 0 {
		for i := range lootDropsDatabase {
			if strings.Contains(lootDropsDatabase[i].Name, lootDrop.Name) {
				lootDrops = append(lootDrops, lootDropsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(lootDrops))
	return
}

//EditLootDrop will grab data from storage
func (s *Storage) EditLootDrop(lootDrop *model.LootDrop) (err error) {
	lootDropLock.Lock()
	defer lootDropLock.Unlock()
	for i := range lootDropsDatabase {
		if lootDropsDatabase[i].ID == lootDrop.ID {
			*lootDropsDatabase[i] = *lootDrop
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteLootDrop will grab data from storage
func (s *Storage) DeleteLootDrop(lootDrop *model.LootDrop) (err error) {
	lootDropLock.Lock()
	defer lootDropLock.Unlock()
	indexToDelete := 0
	for i := range lootDropsDatabase {
		if lootDropsDatabase[i].ID == lootDrop.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	lootDropsDatabase[len(lootDropsDatabase)-1], lootDropsDatabase[indexToDelete] = lootDropsDatabase[indexToDelete], lootDropsDatabase[len(lootDropsDatabase)-1]
	lootDropsDatabase = lootDropsDatabase[:len(lootDropsDatabase)-1]
	return
}
