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
	spellAnimationTypeLock = sync.RWMutex{}
)

//GetSpellAnimationType will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetSpellAnimationType(spellAnimationType *model.SpellAnimationType) (err error) {
	spellAnimationTypeLock.Lock()
	defer spellAnimationTypeLock.Unlock()
	spellAnimationTypesDatabase, err := s.readSpellAnimationTypeFile()
	if err != nil {
		return
	}
	for _, tmpSpellAnimationType := range spellAnimationTypesDatabase {
		if tmpSpellAnimationType.ID == spellAnimationType.ID {
			*spellAnimationType = *tmpSpellAnimationType
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSpellAnimationType will grab data from storage
func (s *Storage) CreateSpellAnimationType(spellAnimationType *model.SpellAnimationType) (err error) {
	spellAnimationTypeLock.Lock()
	defer spellAnimationTypeLock.Unlock()
	spellAnimationTypesDatabase, err := s.readSpellAnimationTypeFile()
	if err != nil {
		return
	}
	for _, tmpSpellAnimationType := range spellAnimationTypesDatabase {
		if tmpSpellAnimationType.ID == spellAnimationType.ID {
			err = fmt.Errorf("spellAnimationType already exists")
			return
		}
	}
	spellAnimationTypesDatabase = append(spellAnimationTypesDatabase, spellAnimationType)
	err = s.writeSpellAnimationTypeFile(spellAnimationTypesDatabase)
	if err != nil {
		return
	}
	return
}

//ListSpellAnimationType will grab data from storage
func (s *Storage) ListSpellAnimationType(page *model.Page) (spellAnimationTypes []*model.SpellAnimationType, err error) {
	spellAnimationTypeLock.Lock()
	defer spellAnimationTypeLock.Unlock()
	spellAnimationTypesDatabase, err := s.readSpellAnimationTypeFile()
	if err != nil {
		return
	}

	spellAnimationTypes = make([]*model.SpellAnimationType, len(spellAnimationTypesDatabase))

	spellAnimationTypes = spellAnimationTypesDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellAnimationTypes, func(i, j int) bool {
			return spellAnimationTypes[i].Name < spellAnimationTypes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(spellAnimationTypes))
		}
	*/
	return
}

//ListSpellAnimationTypeTotalCount will grab data from storage
func (s *Storage) ListSpellAnimationTypeTotalCount() (count int64, err error) {
	spellAnimationTypeLock.Lock()
	defer spellAnimationTypeLock.Unlock()
	spellAnimationTypesDatabase, err := s.readSpellAnimationTypeFile()
	if err != nil {
		return
	}
	count = int64(len(spellAnimationTypesDatabase))
	return
}

//ListSpellAnimationTypeBySearch will grab data from storage
func (s *Storage) ListSpellAnimationTypeBySearch(page *model.Page, spellAnimationType *model.SpellAnimationType) (spellAnimationTypes []*model.SpellAnimationType, err error) {
	spellAnimationTypeLock.Lock()
	defer spellAnimationTypeLock.Unlock()
	spellAnimationTypesDatabase, err := s.readSpellAnimationTypeFile()
	if err != nil {
		return
	}
	if len(spellAnimationType.Name) > 0 {
		for i := range spellAnimationTypesDatabase {
			if strings.Contains(spellAnimationTypesDatabase[i].Name, spellAnimationType.Name) {
				spellAnimationTypes = append(spellAnimationTypes, spellAnimationTypesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellAnimationTypes, func(i, j int) bool {
			return spellAnimationTypes[i].Name < spellAnimationTypes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(spellAnimationTypes))
	//}
	return
}

//ListSpellAnimationTypeBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellAnimationTypeBySearchTotalCount(spellAnimationType *model.SpellAnimationType) (count int64, err error) {
	spellAnimationTypeLock.Lock()
	defer spellAnimationTypeLock.Unlock()
	spellAnimationTypesDatabase, err := s.readSpellAnimationTypeFile()
	if err != nil {
		return
	}

	spellAnimationTypes := []*model.SpellAnimationType{}
	if len(spellAnimationType.Name) > 0 {
		for i := range spellAnimationTypesDatabase {
			if strings.Contains(spellAnimationTypesDatabase[i].Name, spellAnimationType.Name) {
				spellAnimationTypes = append(spellAnimationTypes, spellAnimationTypesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(spellAnimationTypes))
	return
}

//EditSpellAnimationType will grab data from storage
func (s *Storage) EditSpellAnimationType(spellAnimationType *model.SpellAnimationType) (err error) {
	spellAnimationTypeLock.Lock()
	defer spellAnimationTypeLock.Unlock()
	spellAnimationTypesDatabase, err := s.readSpellAnimationTypeFile()
	if err != nil {
		return
	}

	for i := range spellAnimationTypesDatabase {
		if spellAnimationTypesDatabase[i].ID == spellAnimationType.ID {
			*spellAnimationTypesDatabase[i] = *spellAnimationType
			err = s.writeSpellAnimationTypeFile(spellAnimationTypesDatabase)
			if err != nil {
				return
			}
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSpellAnimationType will grab data from storage
func (s *Storage) DeleteSpellAnimationType(spellAnimationType *model.SpellAnimationType) (err error) {
	spellAnimationTypeLock.Lock()
	defer spellAnimationTypeLock.Unlock()
	spellAnimationTypesDatabase, err := s.readSpellAnimationTypeFile()
	if err != nil {
		return
	}
	indexToDelete := 0
	for i := range spellAnimationTypesDatabase {
		if spellAnimationTypesDatabase[i].ID == spellAnimationType.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	spellAnimationTypesDatabase[len(spellAnimationTypesDatabase)-1], spellAnimationTypesDatabase[indexToDelete] = spellAnimationTypesDatabase[indexToDelete], spellAnimationTypesDatabase[len(spellAnimationTypesDatabase)-1]
	spellAnimationTypesDatabase = spellAnimationTypesDatabase[:len(spellAnimationTypesDatabase)-1]
	err = s.writeSpellAnimationTypeFile(spellAnimationTypesDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readSpellAnimationTypeFile() (spellAnimationTypes []*model.SpellAnimationType, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			spellAnimationTypes = loadSpellAnimationTypeDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeSpellAnimationTypeFile(spellAnimationTypes)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default spellAnimationType data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &spellAnimationTypes)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeSpellAnimationTypeFile(spellAnimationTypes []*model.SpellAnimationType) (err error) {

	bData, err := yaml.Marshal(spellAnimationTypes)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal spellAnimationTypes")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
