package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

var (
	raceLock = sync.RWMutex{}
)

//GetRace will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetRace(race *model.Race) (err error) {
	raceLock.Lock()
	defer raceLock.Unlock()
	racesDatabase, err := s.readRaceFile()
	if err != nil {
		return
	}
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
	racesDatabase, err := s.readRaceFile()
	if err != nil {
		return
	}
	for _, tmpRace := range racesDatabase {
		if tmpRace.ID == race.ID {
			err = fmt.Errorf("race already exists")
			return
		}
	}
	racesDatabase = append(racesDatabase, race)
	err = s.writeRaceFile(racesDatabase)
	if err != nil {
		return
	}
	return
}

//ListRace will grab data from storage
func (s *Storage) ListRace(page *model.Page) (races []*model.Race, err error) {
	raceLock.Lock()
	defer raceLock.Unlock()
	racesDatabase, err := s.readRaceFile()
	if err != nil {
		return
	}

	races = make([]*model.Race, len(racesDatabase))

	races = racesDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
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

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(races))
		}
	*/
	return
}

//ListRaceTotalCount will grab data from storage
func (s *Storage) ListRaceTotalCount() (count int64, err error) {
	raceLock.Lock()
	defer raceLock.Unlock()
	racesDatabase, err := s.readRaceFile()
	if err != nil {
		return
	}
	count = int64(len(racesDatabase))
	return
}

//ListRaceByBit will grab data from storage
func (s *Storage) ListRaceByBit(page *model.Page, race *model.Race) (races []*model.Race, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRaceByBitTotalCount will grab data from storage
func (s *Storage) ListRaceByBitTotalCount(race *model.Race) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRaceBySearch will grab data from storage
func (s *Storage) ListRaceBySearch(page *model.Page, race *model.Race) (races []*model.Race, err error) {
	raceLock.Lock()
	defer raceLock.Unlock()
	racesDatabase, err := s.readRaceFile()
	if err != nil {
		return
	}
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
	raceLock.Lock()
	defer raceLock.Unlock()
	racesDatabase, err := s.readRaceFile()
	if err != nil {
		return
	}

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
	racesDatabase, err := s.readRaceFile()
	if err != nil {
		return
	}

	for i := range racesDatabase {
		if racesDatabase[i].ID == race.ID {
			*racesDatabase[i] = *race
			err = s.writeRaceFile(racesDatabase)
			if err != nil {
				return
			}
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
	racesDatabase, err := s.readRaceFile()
	if err != nil {
		return
	}
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
	err = s.writeRaceFile(racesDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readRaceFile() (races []*model.Race, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			races = loadRaceDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeRaceFile(races)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default race data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &races)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeRaceFile(races []*model.Race) (err error) {

	bData, err := yaml.Marshal(races)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal races")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
