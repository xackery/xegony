package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	zoneExpansionsDatabase = []*model.ZoneExpansion{}
	zoneExpansionLock      = sync.RWMutex{}
)

//GetZoneExpansion will grab data from storage
func (s *Storage) GetZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error) {
	zoneExpansionLock.RLock()
	defer zoneExpansionLock.RUnlock()
	for _, tmpZoneExpansion := range zoneExpansionsDatabase {
		if tmpZoneExpansion.ID == zoneExpansion.ID {
			*zoneExpansion = *tmpZoneExpansion
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateZoneExpansion will grab data from storage
func (s *Storage) CreateZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error) {
	zoneExpansionLock.Lock()
	defer zoneExpansionLock.Unlock()
	for _, tmpZoneExpansion := range zoneExpansionsDatabase {
		if tmpZoneExpansion.ID == zoneExpansion.ID {
			err = fmt.Errorf("zoneExpansion already exists")
			return
		}
	}
	zoneExpansionsDatabase = append(zoneExpansionsDatabase, zoneExpansion)
	return
}

//ListZoneExpansion will grab data from storage
func (s *Storage) ListZoneExpansion(page *model.Page) (zoneExpansions []*model.ZoneExpansion, err error) {
	zoneExpansionLock.RLock()
	defer zoneExpansionLock.RUnlock()

	zoneExpansions = make([]*model.ZoneExpansion, len(zoneExpansionsDatabase))

	zoneExpansions = zoneExpansionsDatabase

	if len(page.OrderBy) == 0 {
		page.OrderBy = "id"
	}
	switch page.OrderBy {
	case "short_name":
		sort.Slice(zoneExpansions, func(i, j int) bool {
			return zoneExpansions[i].ShortName < zoneExpansions[j].ShortName
		})
	case "id":
		sort.Slice(zoneExpansions, func(i, j int) bool {
			return zoneExpansions[i].ID < zoneExpansions[j].ID
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(zoneExpansions))
		}
	*/
	return
}

//ListZoneExpansionTotalCount will grab data from storage
func (s *Storage) ListZoneExpansionTotalCount() (count int64, err error) {
	count = int64(len(zoneExpansionsDatabase))
	return
}

//ListZoneExpansionBySearch will grab data from storage
func (s *Storage) ListZoneExpansionBySearch(page *model.Page, zoneExpansion *model.ZoneExpansion) (zoneExpansions []*model.ZoneExpansion, err error) {
	zoneExpansionLock.RLock()
	defer zoneExpansionLock.RUnlock()

	if len(zoneExpansion.ShortName) > 0 {
		for i := range zoneExpansionsDatabase {
			if strings.Contains(zoneExpansionsDatabase[i].ShortName, zoneExpansion.ShortName) {
				zoneExpansions = append(zoneExpansions, zoneExpansionsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "short_name":
		sort.Slice(zoneExpansions, func(i, j int) bool {
			return zoneExpansions[i].ShortName < zoneExpansions[j].ShortName
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(zoneExpansions))
	//}
	return
}

//ListZoneExpansionBySearchTotalCount will grab data from storage
func (s *Storage) ListZoneExpansionBySearchTotalCount(zoneExpansion *model.ZoneExpansion) (count int64, err error) {
	zoneExpansionLock.RLock()
	defer zoneExpansionLock.RUnlock()

	zoneExpansions := []*model.ZoneExpansion{}
	if len(zoneExpansion.ShortName) > 0 {
		for i := range zoneExpansionsDatabase {
			if strings.Contains(zoneExpansionsDatabase[i].ShortName, zoneExpansion.ShortName) {
				zoneExpansions = append(zoneExpansions, zoneExpansionsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(zoneExpansions))
	return
}

//EditZoneExpansion will grab data from storage
func (s *Storage) EditZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error) {
	zoneExpansionLock.Lock()
	defer zoneExpansionLock.Unlock()
	for i := range zoneExpansionsDatabase {
		if zoneExpansionsDatabase[i].ID == zoneExpansion.ID {
			*zoneExpansionsDatabase[i] = *zoneExpansion
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteZoneExpansion will grab data from storage
func (s *Storage) DeleteZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error) {
	zoneExpansionLock.Lock()
	defer zoneExpansionLock.Unlock()
	indexToDelete := 0
	for i := range zoneExpansionsDatabase {
		if zoneExpansionsDatabase[i].ID == zoneExpansion.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	zoneExpansionsDatabase[len(zoneExpansionsDatabase)-1], zoneExpansionsDatabase[indexToDelete] = zoneExpansionsDatabase[indexToDelete], zoneExpansionsDatabase[len(zoneExpansionsDatabase)-1]
	zoneExpansionsDatabase = zoneExpansionsDatabase[:len(zoneExpansionsDatabase)-1]
	return
}
