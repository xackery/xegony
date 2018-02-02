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
	zoneExpansionLock = sync.RWMutex{}
)

//GetZoneExpansion will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error) {
	zoneExpansionLock.Lock()
	defer zoneExpansionLock.Unlock()

	zoneExpansionsDatabase, err := s.readZoneExpansionFile()
	if err != nil {
		return
	}
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

	zoneExpansionsDatabase, err := s.readZoneExpansionFile()
	if err != nil {
		return
	}
	for _, tmpZoneExpansion := range zoneExpansionsDatabase {
		if tmpZoneExpansion.ID == zoneExpansion.ID {
			err = fmt.Errorf("zoneExpansion already exists")
			return
		}
	}
	zoneExpansionsDatabase = append(zoneExpansionsDatabase, zoneExpansion)
	err = s.writeZoneExpansionFile(zoneExpansionsDatabase)
	if err != nil {
		return
	}
	return
}

//ListZoneExpansion will grab data from storage
func (s *Storage) ListZoneExpansion(page *model.Page) (zoneExpansions []*model.ZoneExpansion, err error) {
	zoneExpansionLock.Lock()
	defer zoneExpansionLock.Unlock()

	zoneExpansionsDatabase, err := s.readZoneExpansionFile()
	if err != nil {
		return
	}

	zoneExpansions = make([]*model.ZoneExpansion, len(zoneExpansionsDatabase))

	zoneExpansions = zoneExpansionsDatabase

	if page.OrderBy == "" {
		page.OrderBy = "short_name"
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

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(zoneExpansions))
		}
	*/
	return
}

//ListZoneExpansionTotalCount will grab data from storage
func (s *Storage) ListZoneExpansionTotalCount() (count int64, err error) {
	zoneExpansionLock.Lock()
	defer zoneExpansionLock.Unlock()

	zoneExpansionsDatabase, err := s.readZoneExpansionFile()
	if err != nil {
		return
	}
	count = int64(len(zoneExpansionsDatabase))
	return
}

//ListZoneExpansionBySearch will grab data from storage
func (s *Storage) ListZoneExpansionBySearch(page *model.Page, zoneExpansion *model.ZoneExpansion) (zoneExpansions []*model.ZoneExpansion, err error) {
	zoneExpansionLock.Lock()
	defer zoneExpansionLock.Unlock()

	zoneExpansionsDatabase, err := s.readZoneExpansionFile()
	if err != nil {
		return
	}
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
	zoneExpansionLock.Lock()
	defer zoneExpansionLock.Unlock()

	zoneExpansionsDatabase, err := s.readZoneExpansionFile()
	if err != nil {
		return
	}

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

	zoneExpansionsDatabase, err := s.readZoneExpansionFile()
	if err != nil {
		return
	}

	for i := range zoneExpansionsDatabase {
		if zoneExpansionsDatabase[i].ID == zoneExpansion.ID {
			*zoneExpansionsDatabase[i] = *zoneExpansion
			err = s.writeZoneExpansionFile(zoneExpansionsDatabase)
			if err != nil {
				return
			}
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

	zoneExpansionsDatabase, err := s.readZoneExpansionFile()
	if err != nil {
		return
	}
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
	err = s.writeZoneExpansionFile(zoneExpansionsDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readZoneExpansionFile() (zoneExpansions []*model.ZoneExpansion, err error) {

	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			zoneExpansions = loadZoneExpansionDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeZoneExpansionFile(zoneExpansions)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default zoneExpansion data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &zoneExpansions)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}
	return
}

func (s *Storage) writeZoneExpansionFile(zoneExpansions []*model.ZoneExpansion) (err error) {

	bData, err := yaml.Marshal(zoneExpansions)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal zoneExpansions")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
