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
	spellTargetTypeLock = sync.RWMutex{}
)

//GetSpellTargetType will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetSpellTargetType(spellTargetType *model.SpellTargetType) (err error) {
	spellTargetTypeLock.Lock()
	defer spellTargetTypeLock.Unlock()
	spellTargetTypesDatabase, err := s.readSpellTargetTypeFile()
	if err != nil {
		return
	}
	for _, tmpSpellTargetType := range spellTargetTypesDatabase {
		if tmpSpellTargetType.ID == spellTargetType.ID {
			*spellTargetType = *tmpSpellTargetType
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSpellTargetType will grab data from storage
func (s *Storage) CreateSpellTargetType(spellTargetType *model.SpellTargetType) (err error) {
	spellTargetTypeLock.Lock()
	defer spellTargetTypeLock.Unlock()
	spellTargetTypesDatabase, err := s.readSpellTargetTypeFile()
	if err != nil {
		return
	}
	for _, tmpSpellTargetType := range spellTargetTypesDatabase {
		if tmpSpellTargetType.ID == spellTargetType.ID {
			err = fmt.Errorf("spellTargetType already exists")
			return
		}
	}
	spellTargetTypesDatabase = append(spellTargetTypesDatabase, spellTargetType)
	err = s.writeSpellTargetTypeFile(spellTargetTypesDatabase)
	if err != nil {
		return
	}
	return
}

//ListSpellTargetType will grab data from storage
func (s *Storage) ListSpellTargetType(page *model.Page) (spellTargetTypes []*model.SpellTargetType, err error) {
	spellTargetTypeLock.Lock()
	defer spellTargetTypeLock.Unlock()
	spellTargetTypesDatabase, err := s.readSpellTargetTypeFile()
	if err != nil {
		return
	}

	spellTargetTypes = make([]*model.SpellTargetType, len(spellTargetTypesDatabase))

	spellTargetTypes = spellTargetTypesDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellTargetTypes, func(i, j int) bool {
			return spellTargetTypes[i].Name < spellTargetTypes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(spellTargetTypes))
		}
	*/
	return
}

//ListSpellTargetTypeTotalCount will grab data from storage
func (s *Storage) ListSpellTargetTypeTotalCount() (count int64, err error) {
	spellTargetTypeLock.Lock()
	defer spellTargetTypeLock.Unlock()
	spellTargetTypesDatabase, err := s.readSpellTargetTypeFile()
	if err != nil {
		return
	}
	count = int64(len(spellTargetTypesDatabase))
	return
}

//ListSpellTargetTypeBySearch will grab data from storage
func (s *Storage) ListSpellTargetTypeBySearch(page *model.Page, spellTargetType *model.SpellTargetType) (spellTargetTypes []*model.SpellTargetType, err error) {
	spellTargetTypeLock.Lock()
	defer spellTargetTypeLock.Unlock()
	spellTargetTypesDatabase, err := s.readSpellTargetTypeFile()
	if err != nil {
		return
	}
	if len(spellTargetType.Name) > 0 {
		for i := range spellTargetTypesDatabase {
			if strings.Contains(spellTargetTypesDatabase[i].Name, spellTargetType.Name) {
				spellTargetTypes = append(spellTargetTypes, spellTargetTypesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellTargetTypes, func(i, j int) bool {
			return spellTargetTypes[i].Name < spellTargetTypes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(spellTargetTypes))
	//}
	return
}

//ListSpellTargetTypeBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellTargetTypeBySearchTotalCount(spellTargetType *model.SpellTargetType) (count int64, err error) {
	spellTargetTypeLock.Lock()
	defer spellTargetTypeLock.Unlock()
	spellTargetTypesDatabase, err := s.readSpellTargetTypeFile()
	if err != nil {
		return
	}

	spellTargetTypes := []*model.SpellTargetType{}
	if len(spellTargetType.Name) > 0 {
		for i := range spellTargetTypesDatabase {
			if strings.Contains(spellTargetTypesDatabase[i].Name, spellTargetType.Name) {
				spellTargetTypes = append(spellTargetTypes, spellTargetTypesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(spellTargetTypes))
	return
}

//EditSpellTargetType will grab data from storage
func (s *Storage) EditSpellTargetType(spellTargetType *model.SpellTargetType) (err error) {
	spellTargetTypeLock.Lock()
	defer spellTargetTypeLock.Unlock()
	spellTargetTypesDatabase, err := s.readSpellTargetTypeFile()
	if err != nil {
		return
	}

	for i := range spellTargetTypesDatabase {
		if spellTargetTypesDatabase[i].ID == spellTargetType.ID {
			*spellTargetTypesDatabase[i] = *spellTargetType
			err = s.writeSpellTargetTypeFile(spellTargetTypesDatabase)
			if err != nil {
				return
			}
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSpellTargetType will grab data from storage
func (s *Storage) DeleteSpellTargetType(spellTargetType *model.SpellTargetType) (err error) {
	spellTargetTypeLock.Lock()
	defer spellTargetTypeLock.Unlock()
	spellTargetTypesDatabase, err := s.readSpellTargetTypeFile()
	if err != nil {
		return
	}
	indexToDelete := 0
	for i := range spellTargetTypesDatabase {
		if spellTargetTypesDatabase[i].ID == spellTargetType.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	spellTargetTypesDatabase[len(spellTargetTypesDatabase)-1], spellTargetTypesDatabase[indexToDelete] = spellTargetTypesDatabase[indexToDelete], spellTargetTypesDatabase[len(spellTargetTypesDatabase)-1]
	spellTargetTypesDatabase = spellTargetTypesDatabase[:len(spellTargetTypesDatabase)-1]
	err = s.writeSpellTargetTypeFile(spellTargetTypesDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readSpellTargetTypeFile() (spellTargetTypes []*model.SpellTargetType, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			spellTargetTypes = loadSpellTargetTypeDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeSpellTargetTypeFile(spellTargetTypes)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default spellTargetType data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &spellTargetTypes)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeSpellTargetTypeFile(spellTargetTypes []*model.SpellTargetType) (err error) {

	bData, err := yaml.Marshal(spellTargetTypes)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal spellTargetTypes")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
