package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	racesDatabase = []*model.Race{}
	raceLock      = sync.RWMutex{}
)

//GetRace will grab data from storage
func (s *Storage) GetRace(race *model.Race) (err error) {
	raceLock.RLock()
	defer raceLock.RUnlock()
	for _, tmpRace := range racesDatabase {
		if tmpRace.ID == race.ID {
			*race = *tmpRace
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateRace will grab data from storage
func (s *Storage) CreateRace(race *model.Race) (err error) {
	raceLock.Lock()
	defer raceLock.Unlock()
	for _, tmpRace := range racesDatabase {
		if tmpRace.ID == race.ID {
			err = fmt.Errorf("race already exists")
			return
		}
	}
	racesDatabase = append(racesDatabase, race)
	return
}

//ListRace will grab data from storage
func (s *Storage) ListRace(page *model.Page) (races []*model.Race, err error) {
	raceLock.RLock()
	defer raceLock.RUnlock()

	races = make([]*model.Race, len(racesDatabase))

	races = racesDatabase

	switch page.OrderBy {
	case "name":
		sort.Slice(races, func(i, j int) bool {
			return races[i].Name < races[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(races))
		}
	*/
	return
}

//ListRaceTotalCount will grab data from storage
func (s *Storage) ListRaceTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(racesDatabase))
	return
}

//ListRaceByBit will grab data from storage
func (s *Storage) ListRaceByBit(page *model.Page, race *model.Race) (races []*model.Race, err error) {
	raceLock.RLock()
	defer raceLock.RUnlock()

	for i := range racesDatabase {
		if racesDatabase[i].Bit < 1 || race.Bit < 1 {
			continue
		}

		if race.Bit&racesDatabase[i].Bit == racesDatabase[i].Bit {
			races = append(races, racesDatabase[i])
		}
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(races, func(i, j int) bool {
			return races[i].Name < races[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(races))
	//}
	return
}

//ListRaceByBitTotalCount will grab data from storage
func (s *Storage) ListRaceByBitTotalCount(race *model.Race) (count int64, err error) {
	raceLock.RLock()
	defer raceLock.RUnlock()

	races := []*model.Race{}
	for i := range racesDatabase {
		if race.Bit < 1 || racesDatabase[i].Bit < 1 {
			continue
		}
		if race.Bit&racesDatabase[i].Bit == racesDatabase[i].Bit {
			races = append(races, racesDatabase[i])
		}
	}

	count = int64(len(races))
	return
}

//ListRaceBySearch will grab data from storage
func (s *Storage) ListRaceBySearch(page *model.Page, race *model.Race) (races []*model.Race, err error) {
	raceLock.RLock()
	defer raceLock.RUnlock()

	if len(race.Name) > 0 {
		for i := range racesDatabase {
			if strings.Contains(racesDatabase[i].Name, race.Name) {
				races = append(races, racesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(races, func(i, j int) bool {
			return races[i].Name < races[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(races))
	//}
	return
}

//ListRaceBySearchTotalCount will grab data from storage
func (s *Storage) ListRaceBySearchTotalCount(race *model.Race) (count int64, err error) {
	raceLock.RLock()
	defer raceLock.RUnlock()

	races := []*model.Race{}
	if len(race.Name) > 0 {
		for i := range racesDatabase {
			if strings.Contains(racesDatabase[i].Name, race.Name) {
				races = append(races, racesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(races))
	return
}

//EditRace will grab data from storage
func (s *Storage) EditRace(race *model.Race) (err error) {
	raceLock.Lock()
	defer raceLock.Unlock()
	for i := range racesDatabase {
		if racesDatabase[i].ID == race.ID {
			*racesDatabase[i] = *race
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteRace will grab data from storage
func (s *Storage) DeleteRace(race *model.Race) (err error) {
	raceLock.Lock()
	defer raceLock.Unlock()
	indexToDelete := 0
	for i := range racesDatabase {
		if racesDatabase[i].ID == race.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	racesDatabase[len(racesDatabase)-1], racesDatabase[indexToDelete] = racesDatabase[indexToDelete], racesDatabase[len(racesDatabase)-1]
	racesDatabase = racesDatabase[:len(racesDatabase)-1]
	return
}
