package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	spellTravelTypesDatabase = []*model.SpellTravelType{}
	spellTravelTypeLock      = sync.RWMutex{}
)

//GetSpellTravelType will grab data from storage
func (s *Storage) GetSpellTravelType(spellTravelType *model.SpellTravelType) (err error) {
	spellTravelTypeLock.RLock()
	defer spellTravelTypeLock.RUnlock()
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
	for _, tmpSpellTravelType := range spellTravelTypesDatabase {
		if tmpSpellTravelType.ID == spellTravelType.ID {
			err = fmt.Errorf("spellTravelType already exists")
			return
		}
	}
	spellTravelTypesDatabase = append(spellTravelTypesDatabase, spellTravelType)
	return
}

//ListSpellTravelType will grab data from storage
func (s *Storage) ListSpellTravelType(page *model.Page) (spellTravelTypes []*model.SpellTravelType, err error) {
	spellTravelTypeLock.RLock()
	defer spellTravelTypeLock.RUnlock()

	spellTravelTypes = make([]*model.SpellTravelType, len(spellTravelTypesDatabase))

	spellTravelTypes = spellTravelTypesDatabase

	switch page.OrderBy {
	case "id":
		sort.Slice(spellTravelTypes, func(i, j int) bool {
			return spellTravelTypes[i].ID < spellTravelTypes[j].ID
		})
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
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(spellTravelTypesDatabase))
	return
}

//ListSpellTravelTypeBySearch will grab data from storage
func (s *Storage) ListSpellTravelTypeBySearch(page *model.Page, spellTravelType *model.SpellTravelType) (spellTravelTypes []*model.SpellTravelType, err error) {
	spellTravelTypeLock.RLock()
	defer spellTravelTypeLock.RUnlock()

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
	case "id":
		sort.Slice(spellTravelTypes, func(i, j int) bool {
			return spellTravelTypes[i].ID < spellTravelTypes[j].ID
		})
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
	spellTravelTypeLock.RLock()
	defer spellTravelTypeLock.RUnlock()

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
	for i := range spellTravelTypesDatabase {
		if spellTravelTypesDatabase[i].ID == spellTravelType.ID {
			*spellTravelTypesDatabase[i] = *spellTravelType
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
	return
}
