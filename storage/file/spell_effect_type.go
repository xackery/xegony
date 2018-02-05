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
	spellEffectTypeLock = sync.RWMutex{}
)

//GetSpellEffectType will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetSpellEffectType(spellEffectType *model.SpellEffectType) (err error) {
	spellEffectTypeLock.Lock()
	defer spellEffectTypeLock.Unlock()
	spellEffectTypesDatabase, err := s.readSpellEffectTypeFile()
	if err != nil {
		return
	}
	for _, tmpSpellEffectType := range spellEffectTypesDatabase {
		if tmpSpellEffectType.ID == spellEffectType.ID {
			*spellEffectType = *tmpSpellEffectType
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSpellEffectType will grab data from storage
func (s *Storage) CreateSpellEffectType(spellEffectType *model.SpellEffectType) (err error) {
	spellEffectTypeLock.Lock()
	defer spellEffectTypeLock.Unlock()
	spellEffectTypesDatabase, err := s.readSpellEffectTypeFile()
	if err != nil {
		return
	}
	for _, tmpSpellEffectType := range spellEffectTypesDatabase {
		if tmpSpellEffectType.ID == spellEffectType.ID {
			err = fmt.Errorf("spellEffectType already exists")
			return
		}
	}
	spellEffectTypesDatabase = append(spellEffectTypesDatabase, spellEffectType)
	err = s.writeSpellEffectTypeFile(spellEffectTypesDatabase)
	if err != nil {
		return
	}
	return
}

//ListSpellEffectType will grab data from storage
func (s *Storage) ListSpellEffectType(page *model.Page) (spellEffectTypes []*model.SpellEffectType, err error) {
	spellEffectTypeLock.Lock()
	defer spellEffectTypeLock.Unlock()
	spellEffectTypesDatabase, err := s.readSpellEffectTypeFile()
	if err != nil {
		return
	}

	spellEffectTypes = make([]*model.SpellEffectType, len(spellEffectTypesDatabase))

	spellEffectTypes = spellEffectTypesDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellEffectTypes, func(i, j int) bool {
			return spellEffectTypes[i].Name < spellEffectTypes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(spellEffectTypes))
		}
	*/
	return
}

//ListSpellEffectTypeTotalCount will grab data from storage
func (s *Storage) ListSpellEffectTypeTotalCount() (count int64, err error) {
	spellEffectTypeLock.Lock()
	defer spellEffectTypeLock.Unlock()
	spellEffectTypesDatabase, err := s.readSpellEffectTypeFile()
	if err != nil {
		return
	}
	count = int64(len(spellEffectTypesDatabase))
	return
}

//ListSpellEffectTypeBySearch will grab data from storage
func (s *Storage) ListSpellEffectTypeBySearch(page *model.Page, spellEffectType *model.SpellEffectType) (spellEffectTypes []*model.SpellEffectType, err error) {
	spellEffectTypeLock.Lock()
	defer spellEffectTypeLock.Unlock()
	spellEffectTypesDatabase, err := s.readSpellEffectTypeFile()
	if err != nil {
		return
	}
	if len(spellEffectType.Name) > 0 {
		for i := range spellEffectTypesDatabase {
			if strings.Contains(spellEffectTypesDatabase[i].Name, spellEffectType.Name) {
				spellEffectTypes = append(spellEffectTypes, spellEffectTypesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellEffectTypes, func(i, j int) bool {
			return spellEffectTypes[i].Name < spellEffectTypes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(spellEffectTypes))
	//}
	return
}

//ListSpellEffectTypeBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellEffectTypeBySearchTotalCount(spellEffectType *model.SpellEffectType) (count int64, err error) {
	spellEffectTypeLock.Lock()
	defer spellEffectTypeLock.Unlock()
	spellEffectTypesDatabase, err := s.readSpellEffectTypeFile()
	if err != nil {
		return
	}

	spellEffectTypes := []*model.SpellEffectType{}
	if len(spellEffectType.Name) > 0 {
		for i := range spellEffectTypesDatabase {
			if strings.Contains(spellEffectTypesDatabase[i].Name, spellEffectType.Name) {
				spellEffectTypes = append(spellEffectTypes, spellEffectTypesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(spellEffectTypes))
	return
}

//EditSpellEffectType will grab data from storage
func (s *Storage) EditSpellEffectType(spellEffectType *model.SpellEffectType) (err error) {
	spellEffectTypeLock.Lock()
	defer spellEffectTypeLock.Unlock()
	spellEffectTypesDatabase, err := s.readSpellEffectTypeFile()
	if err != nil {
		return
	}

	for i := range spellEffectTypesDatabase {
		if spellEffectTypesDatabase[i].ID == spellEffectType.ID {
			*spellEffectTypesDatabase[i] = *spellEffectType
			err = s.writeSpellEffectTypeFile(spellEffectTypesDatabase)
			if err != nil {
				return
			}
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSpellEffectType will grab data from storage
func (s *Storage) DeleteSpellEffectType(spellEffectType *model.SpellEffectType) (err error) {
	spellEffectTypeLock.Lock()
	defer spellEffectTypeLock.Unlock()
	spellEffectTypesDatabase, err := s.readSpellEffectTypeFile()
	if err != nil {
		return
	}
	indexToDelete := 0
	for i := range spellEffectTypesDatabase {
		if spellEffectTypesDatabase[i].ID == spellEffectType.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	spellEffectTypesDatabase[len(spellEffectTypesDatabase)-1], spellEffectTypesDatabase[indexToDelete] = spellEffectTypesDatabase[indexToDelete], spellEffectTypesDatabase[len(spellEffectTypesDatabase)-1]
	spellEffectTypesDatabase = spellEffectTypesDatabase[:len(spellEffectTypesDatabase)-1]
	err = s.writeSpellEffectTypeFile(spellEffectTypesDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readSpellEffectTypeFile() (spellEffectTypes []*model.SpellEffectType, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			spellEffectTypes = loadSpellEffectTypeDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeSpellEffectTypeFile(spellEffectTypes)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default spellEffectType data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &spellEffectTypes)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeSpellEffectTypeFile(spellEffectTypes []*model.SpellEffectType) (err error) {

	bData, err := yaml.Marshal(spellEffectTypes)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal spellEffectTypes")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
