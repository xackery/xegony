package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	spellEffectTypesDatabase = []*model.SpellEffectType{}
	spellEffectTypeLock      = sync.RWMutex{}
)

//GetSpellEffectType will grab data from storage
func (s *Storage) GetSpellEffectType(spellEffectType *model.SpellEffectType) (err error) {
	spellEffectTypeLock.RLock()
	defer spellEffectTypeLock.RUnlock()
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
	for _, tmpSpellEffectType := range spellEffectTypesDatabase {
		if tmpSpellEffectType.ID == spellEffectType.ID {
			err = fmt.Errorf("spellEffectType already exists")
			return
		}
	}
	spellEffectTypesDatabase = append(spellEffectTypesDatabase, spellEffectType)
	return
}

//ListSpellEffectType will grab data from storage
func (s *Storage) ListSpellEffectType(page *model.Page) (spellEffectTypes []*model.SpellEffectType, err error) {
	spellEffectTypeLock.RLock()
	defer spellEffectTypeLock.RUnlock()

	spellEffectTypes = make([]*model.SpellEffectType, len(spellEffectTypesDatabase))

	spellEffectTypes = spellEffectTypesDatabase

	switch page.OrderBy {
	case "id":
		sort.Slice(spellEffectTypes, func(i, j int) bool {
			return spellEffectTypes[i].ID < spellEffectTypes[j].ID
		})
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
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(spellEffectTypesDatabase))
	return
}

//ListSpellEffectTypeBySearch will grab data from storage
func (s *Storage) ListSpellEffectTypeBySearch(page *model.Page, spellEffectType *model.SpellEffectType) (spellEffectTypes []*model.SpellEffectType, err error) {
	spellEffectTypeLock.RLock()
	defer spellEffectTypeLock.RUnlock()

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
	case "id":
		sort.Slice(spellEffectTypes, func(i, j int) bool {
			return spellEffectTypes[i].ID < spellEffectTypes[j].ID
		})
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
	spellEffectTypeLock.RLock()
	defer spellEffectTypeLock.RUnlock()

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
	for i := range spellEffectTypesDatabase {
		if spellEffectTypesDatabase[i].ID == spellEffectType.ID {
			*spellEffectTypesDatabase[i] = *spellEffectType
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
	return
}
