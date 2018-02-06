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
	oauthTypeLock = sync.RWMutex{}
)

//GetOauthType will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetOauthType(oauthType *model.OauthType) (err error) {
	oauthTypeLock.Lock()
	defer oauthTypeLock.Unlock()
	oauthTypesDatabase, err := s.readOauthTypeFile()
	if err != nil {
		return
	}
	for _, tmpOauthType := range oauthTypesDatabase {
		if tmpOauthType.ID == oauthType.ID {
			*oauthType = *tmpOauthType
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateOauthType will grab data from storage
func (s *Storage) CreateOauthType(oauthType *model.OauthType) (err error) {
	oauthTypeLock.Lock()
	defer oauthTypeLock.Unlock()
	oauthTypesDatabase, err := s.readOauthTypeFile()
	if err != nil {
		return
	}
	for _, tmpOauthType := range oauthTypesDatabase {
		if tmpOauthType.ID == oauthType.ID {
			err = fmt.Errorf("oauthType already exists")
			return
		}
	}
	oauthTypesDatabase = append(oauthTypesDatabase, oauthType)
	err = s.writeOauthTypeFile(oauthTypesDatabase)
	if err != nil {
		return
	}
	return
}

//ListOauthType will grab data from storage
func (s *Storage) ListOauthType(page *model.Page) (oauthTypes []*model.OauthType, err error) {
	oauthTypeLock.Lock()
	defer oauthTypeLock.Unlock()
	oauthTypesDatabase, err := s.readOauthTypeFile()
	if err != nil {
		return
	}

	oauthTypes = make([]*model.OauthType, len(oauthTypesDatabase))

	oauthTypes = oauthTypesDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(oauthTypes, func(i, j int) bool {
			return oauthTypes[i].Name < oauthTypes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(oauthTypes))
		}
	*/
	return
}

//ListOauthTypeTotalCount will grab data from storage
func (s *Storage) ListOauthTypeTotalCount() (count int64, err error) {
	oauthTypeLock.Lock()
	defer oauthTypeLock.Unlock()
	oauthTypesDatabase, err := s.readOauthTypeFile()
	if err != nil {
		return
	}
	count = int64(len(oauthTypesDatabase))
	return
}

//ListOauthTypeBySearch will grab data from storage
func (s *Storage) ListOauthTypeBySearch(page *model.Page, oauthType *model.OauthType) (oauthTypes []*model.OauthType, err error) {
	oauthTypeLock.Lock()
	defer oauthTypeLock.Unlock()
	oauthTypesDatabase, err := s.readOauthTypeFile()
	if err != nil {
		return
	}
	if len(oauthType.Name) > 0 {
		for i := range oauthTypesDatabase {
			if strings.Contains(oauthTypesDatabase[i].Name, oauthType.Name) {
				oauthTypes = append(oauthTypes, oauthTypesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(oauthTypes, func(i, j int) bool {
			return oauthTypes[i].Name < oauthTypes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(oauthTypes))
	//}
	return
}

//ListOauthTypeBySearchTotalCount will grab data from storage
func (s *Storage) ListOauthTypeBySearchTotalCount(oauthType *model.OauthType) (count int64, err error) {
	oauthTypeLock.Lock()
	defer oauthTypeLock.Unlock()
	oauthTypesDatabase, err := s.readOauthTypeFile()
	if err != nil {
		return
	}

	oauthTypes := []*model.OauthType{}
	if len(oauthType.Name) > 0 {
		for i := range oauthTypesDatabase {
			if strings.Contains(oauthTypesDatabase[i].Name, oauthType.Name) {
				oauthTypes = append(oauthTypes, oauthTypesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(oauthTypes))
	return
}

//EditOauthType will grab data from storage
func (s *Storage) EditOauthType(oauthType *model.OauthType) (err error) {
	oauthTypeLock.Lock()
	defer oauthTypeLock.Unlock()
	oauthTypesDatabase, err := s.readOauthTypeFile()
	if err != nil {
		return
	}

	for i := range oauthTypesDatabase {
		if oauthTypesDatabase[i].ID == oauthType.ID {
			*oauthTypesDatabase[i] = *oauthType
			err = s.writeOauthTypeFile(oauthTypesDatabase)
			if err != nil {
				return
			}
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteOauthType will grab data from storage
func (s *Storage) DeleteOauthType(oauthType *model.OauthType) (err error) {
	oauthTypeLock.Lock()
	defer oauthTypeLock.Unlock()
	oauthTypesDatabase, err := s.readOauthTypeFile()
	if err != nil {
		return
	}
	indexToDelete := 0
	for i := range oauthTypesDatabase {
		if oauthTypesDatabase[i].ID == oauthType.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	oauthTypesDatabase[len(oauthTypesDatabase)-1], oauthTypesDatabase[indexToDelete] = oauthTypesDatabase[indexToDelete], oauthTypesDatabase[len(oauthTypesDatabase)-1]
	oauthTypesDatabase = oauthTypesDatabase[:len(oauthTypesDatabase)-1]
	err = s.writeOauthTypeFile(oauthTypesDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readOauthTypeFile() (oauthTypes []*model.OauthType, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			oauthTypes = loadOauthTypeDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeOauthTypeFile(oauthTypes)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default oauthType data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &oauthTypes)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeOauthTypeFile(oauthTypes []*model.OauthType) (err error) {

	bData, err := yaml.Marshal(oauthTypes)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal oauthTypes")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
