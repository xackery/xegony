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
	spellTravelTypeLock = sync.RWMutex{}
)

//GetSpellTravelType will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetSpellTravelType(spellTravelType *model.SpellTravelType) (err error) {
	spellTravelTypeLock.Lock()
	defer spellTravelTypeLock.Unlock()
	spellTravelTypesDatabase, err := s.readSpellTravelTypeFile()
	if err != nil {
		return
	}
	for _, tmpSpellTravelType := range spellTravelTypesDatabase {
		if tmpSpellTravelType.ID == spellTravelType.ID {
			*spellTravelType = *tmpSpellTravelType
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSpellTravelType will grab data from storage
func (s *Storage) CreateSpellTravelType(spellTravelType *model.SpellTravelType) (err error) {
	spellTravelTypeLock.Lock()
	defer spellTravelTypeLock.Unlock()
	spellTravelTypesDatabase, err := s.readSpellTravelTypeFile()
	if err != nil {
		return
	}
	for _, tmpSpellTravelType := range spellTravelTypesDatabase {
		if tmpSpellTravelType.ID == spellTravelType.ID {
			err = fmt.Errorf("spellTravelType already exists")
			return
		}
	}
	spellTravelTypesDatabase = append(spellTravelTypesDatabase, spellTravelType)
	err = s.writeSpellTravelTypeFile(spellTravelTypesDatabase)
	if err != nil {
		return
	}
	return
}

//ListSpellTravelType will grab data from storage
func (s *Storage) ListSpellTravelType(page *model.Page) (spellTravelTypes []*model.SpellTravelType, err error) {
	spellTravelTypeLock.Lock()
	defer spellTravelTypeLock.Unlock()
	spellTravelTypesDatabase, err := s.readSpellTravelTypeFile()
	if err != nil {
		return
	}

	spellTravelTypes = make([]*model.SpellTravelType, len(spellTravelTypesDatabase))

	spellTravelTypes = spellTravelTypesDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellTravelTypes, func(i, j int) bool {
			return spellTravelTypes[i].Name < spellTravelTypes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(spellTravelTypes))
		}
	*/
	return
}

//ListSpellTravelTypeTotalCount will grab data from storage
func (s *Storage) ListSpellTravelTypeTotalCount() (count int64, err error) {
	spellTravelTypeLock.Lock()
	defer spellTravelTypeLock.Unlock()
	spellTravelTypesDatabase, err := s.readSpellTravelTypeFile()
	if err != nil {
		return
	}
	count = int64(len(spellTravelTypesDatabase))
	return
}

//ListSpellTravelTypeBySearch will grab data from storage
func (s *Storage) ListSpellTravelTypeBySearch(page *model.Page, spellTravelType *model.SpellTravelType) (spellTravelTypes []*model.SpellTravelType, err error) {
	spellTravelTypeLock.Lock()
	defer spellTravelTypeLock.Unlock()
	spellTravelTypesDatabase, err := s.readSpellTravelTypeFile()
	if err != nil {
		return
	}
	if len(spellTravelType.Name) > 0 {
		for i := range spellTravelTypesDatabase {
			if strings.Contains(spellTravelTypesDatabase[i].Name, spellTravelType.Name) {
				spellTravelTypes = append(spellTravelTypes, spellTravelTypesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellTravelTypes, func(i, j int) bool {
			return spellTravelTypes[i].Name < spellTravelTypes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(spellTravelTypes))
	//}
	return
}

//ListSpellTravelTypeBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellTravelTypeBySearchTotalCount(spellTravelType *model.SpellTravelType) (count int64, err error) {
	spellTravelTypeLock.Lock()
	defer spellTravelTypeLock.Unlock()
	spellTravelTypesDatabase, err := s.readSpellTravelTypeFile()
	if err != nil {
		return
	}

	spellTravelTypes := []*model.SpellTravelType{}
	if len(spellTravelType.Name) > 0 {
		for i := range spellTravelTypesDatabase {
			if strings.Contains(spellTravelTypesDatabase[i].Name, spellTravelType.Name) {
				spellTravelTypes = append(spellTravelTypes, spellTravelTypesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(spellTravelTypes))
	return
}

//EditSpellTravelType will grab data from storage
func (s *Storage) EditSpellTravelType(spellTravelType *model.SpellTravelType) (err error) {
	spellTravelTypeLock.Lock()
	defer spellTravelTypeLock.Unlock()
	spellTravelTypesDatabase, err := s.readSpellTravelTypeFile()
	if err != nil {
		return
	}

	for i := range spellTravelTypesDatabase {
		if spellTravelTypesDatabase[i].ID == spellTravelType.ID {
			*spellTravelTypesDatabase[i] = *spellTravelType
			err = s.writeSpellTravelTypeFile(spellTravelTypesDatabase)
			if err != nil {
				return
			}
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSpellTravelType will grab data from storage
func (s *Storage) DeleteSpellTravelType(spellTravelType *model.SpellTravelType) (err error) {
	spellTravelTypeLock.Lock()
	defer spellTravelTypeLock.Unlock()
	spellTravelTypesDatabase, err := s.readSpellTravelTypeFile()
	if err != nil {
		return
	}
	indexToDelete := 0
	for i := range spellTravelTypesDatabase {
		if spellTravelTypesDatabase[i].ID == spellTravelType.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	spellTravelTypesDatabase[len(spellTravelTypesDatabase)-1], spellTravelTypesDatabase[indexToDelete] = spellTravelTypesDatabase[indexToDelete], spellTravelTypesDatabase[len(spellTravelTypesDatabase)-1]
	spellTravelTypesDatabase = spellTravelTypesDatabase[:len(spellTravelTypesDatabase)-1]
	err = s.writeSpellTravelTypeFile(spellTravelTypesDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readSpellTravelTypeFile() (spellTravelTypes []*model.SpellTravelType, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			spellTravelTypes = loadSpellTravelTypeDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeSpellTravelTypeFile(spellTravelTypes)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default spellTravelType data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &spellTravelTypes)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeSpellTravelTypeFile(spellTravelTypes []*model.SpellTravelType) (err error) {

	bData, err := yaml.Marshal(spellTravelTypes)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal spellTravelTypes")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
