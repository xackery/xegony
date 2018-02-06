package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	zonesDatabase = []*model.Zone{}
	zoneLock      = sync.RWMutex{}
)

//GetZoneByShortName will grab data from storage
func (s *Storage) GetZoneByShortName(zone *model.Zone) (err error) {
	zoneLock.RLock()
	defer zoneLock.RUnlock()
	for _, tmpZone := range zonesDatabase {
		if tmpZone.ShortName.String == zone.ShortName.String {
			*zone = *tmpZone
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//GetZone will grab data from storage
func (s *Storage) GetZone(zone *model.Zone) (err error) {
	zoneLock.RLock()
	defer zoneLock.RUnlock()
	for _, tmpZone := range zonesDatabase {
		if tmpZone.ID == zone.ID {
			*zone = *tmpZone
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateZone will grab data from storage
func (s *Storage) CreateZone(zone *model.Zone) (err error) {
	zoneLock.Lock()
	defer zoneLock.Unlock()
	for _, tmpZone := range zonesDatabase {
		if tmpZone.ID == zone.ID {
			err = fmt.Errorf("zone already exists")
			return
		}
	}
	zonesDatabase = append(zonesDatabase, zone)
	return
}

//ListZone will grab data from storage
func (s *Storage) ListZone(page *model.Page) (zones []*model.Zone, err error) {
	zoneLock.RLock()
	defer zoneLock.RUnlock()

	zones = make([]*model.Zone, len(zonesDatabase))

	zones = zonesDatabase

	switch page.OrderBy {
	case "short_name":
		sort.Slice(zones, func(i, j int) bool {
			return zones[i].ShortName.String < zones[j].ShortName.String
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(zones))
		}
	*/
	return
}

//ListZoneTotalCount will grab data from storage
func (s *Storage) ListZoneTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(zonesDatabase))
	return
}

//ListZoneBySearch will grab data from storage
func (s *Storage) ListZoneBySearch(page *model.Page, zone *model.Zone) (zones []*model.Zone, err error) {
	zoneLock.RLock()
	defer zoneLock.RUnlock()

	if len(zone.ShortName.String) > 0 {
		for i := range zonesDatabase {
			if strings.Contains(zonesDatabase[i].ShortName.String, zone.ShortName.String) {
				zones = append(zones, zonesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "short_name":
		sort.Slice(zones, func(i, j int) bool {
			return zones[i].ShortName.String < zones[j].ShortName.String
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(zones))
	//}
	return
}

//ListZoneBySearchTotalCount will grab data from storage
func (s *Storage) ListZoneBySearchTotalCount(zone *model.Zone) (count int64, err error) {
	zoneLock.RLock()
	defer zoneLock.RUnlock()

	zones := []*model.Zone{}
	if len(zone.ShortName.String) > 0 {
		for i := range zonesDatabase {
			if strings.Contains(zonesDatabase[i].ShortName.String, zone.ShortName.String) {
				zones = append(zones, zonesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(zones))
	return
}

//EditZone will grab data from storage
func (s *Storage) EditZone(zone *model.Zone) (err error) {
	zoneLock.Lock()
	defer zoneLock.Unlock()
	for i := range zonesDatabase {
		if zonesDatabase[i].ID == zone.ID {
			*zonesDatabase[i] = *zone
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteZone will grab data from storage
func (s *Storage) DeleteZone(zone *model.Zone) (err error) {
	zoneLock.Lock()
	defer zoneLock.Unlock()
	indexToDelete := 0
	for i := range zonesDatabase {
		if zonesDatabase[i].ID == zone.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	zonesDatabase[len(zonesDatabase)-1], zonesDatabase[indexToDelete] = zonesDatabase[indexToDelete], zonesDatabase[len(zonesDatabase)-1]
	zonesDatabase = zonesDatabase[:len(zonesDatabase)-1]
	return
}
