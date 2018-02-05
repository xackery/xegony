package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	spellTargetTypesDatabase = []*model.SpellTargetType{}
	spellTargetTypeLock      = sync.RWMutex{}
)

//GetSpellTargetType will grab data from storage
func (s *Storage) GetSpellTargetType(spellTargetType *model.SpellTargetType) (err error) {
	spellTargetTypeLock.RLock()
	defer spellTargetTypeLock.RUnlock()
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
	for _, tmpSpellTargetType := range spellTargetTypesDatabase {
		if tmpSpellTargetType.ID == spellTargetType.ID {
			err = fmt.Errorf("spellTargetType already exists")
			return
		}
	}
	spellTargetTypesDatabase = append(spellTargetTypesDatabase, spellTargetType)
	return
}

//ListSpellTargetType will grab data from storage
func (s *Storage) ListSpellTargetType(page *model.Page) (spellTargetTypes []*model.SpellTargetType, err error) {
	spellTargetTypeLock.RLock()
	defer spellTargetTypeLock.RUnlock()

	spellTargetTypes = make([]*model.SpellTargetType, len(spellTargetTypesDatabase))

	spellTargetTypes = spellTargetTypesDatabase

	switch page.OrderBy {
	case "id":
		sort.Slice(spellTargetTypes, func(i, j int) bool {
			return spellTargetTypes[i].ID < spellTargetTypes[j].ID
		})
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
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(spellTargetTypesDatabase))
	return
}

//ListSpellTargetTypeBySearch will grab data from storage
func (s *Storage) ListSpellTargetTypeBySearch(page *model.Page, spellTargetType *model.SpellTargetType) (spellTargetTypes []*model.SpellTargetType, err error) {
	spellTargetTypeLock.RLock()
	defer spellTargetTypeLock.RUnlock()

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
	case "id":
		sort.Slice(spellTargetTypes, func(i, j int) bool {
			return spellTargetTypes[i].ID < spellTargetTypes[j].ID
		})
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
	spellTargetTypeLock.RLock()
	defer spellTargetTypeLock.RUnlock()

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
	for i := range spellTargetTypesDatabase {
		if spellTargetTypesDatabase[i].ID == spellTargetType.ID {
			*spellTargetTypesDatabase[i] = *spellTargetType
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
	return
}
