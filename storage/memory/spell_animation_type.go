package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	spellAnimationTypesDatabase = []*model.SpellAnimationType{}
	spellAnimationTypeLock      = sync.RWMutex{}
)

//GetSpellAnimationType will grab data from storage
func (s *Storage) GetSpellAnimationType(spellAnimationType *model.SpellAnimationType) (err error) {
	spellAnimationTypeLock.RLock()
	defer spellAnimationTypeLock.RUnlock()
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
	for _, tmpSpellAnimationType := range spellAnimationTypesDatabase {
		if tmpSpellAnimationType.ID == spellAnimationType.ID {
			err = fmt.Errorf("spellAnimationType already exists")
			return
		}
	}
	spellAnimationTypesDatabase = append(spellAnimationTypesDatabase, spellAnimationType)
	return
}

//ListSpellAnimationType will grab data from storage
func (s *Storage) ListSpellAnimationType(page *model.Page) (spellAnimationTypes []*model.SpellAnimationType, err error) {
	spellAnimationTypeLock.RLock()
	defer spellAnimationTypeLock.RUnlock()

	spellAnimationTypes = make([]*model.SpellAnimationType, len(spellAnimationTypesDatabase))

	spellAnimationTypes = spellAnimationTypesDatabase

	switch page.OrderBy {
	case "id":
		sort.Slice(spellAnimationTypes, func(i, j int) bool {
			return spellAnimationTypes[i].ID < spellAnimationTypes[j].ID
		})
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
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(spellAnimationTypesDatabase))
	return
}

//ListSpellAnimationTypeBySearch will grab data from storage
func (s *Storage) ListSpellAnimationTypeBySearch(page *model.Page, spellAnimationType *model.SpellAnimationType) (spellAnimationTypes []*model.SpellAnimationType, err error) {
	spellAnimationTypeLock.RLock()
	defer spellAnimationTypeLock.RUnlock()

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
	case "id":
		sort.Slice(spellAnimationTypes, func(i, j int) bool {
			return spellAnimationTypes[i].ID < spellAnimationTypes[j].ID
		})
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
	spellAnimationTypeLock.RLock()
	defer spellAnimationTypeLock.RUnlock()

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
	for i := range spellAnimationTypesDatabase {
		if spellAnimationTypesDatabase[i].ID == spellAnimationType.ID {
			*spellAnimationTypesDatabase[i] = *spellAnimationType
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
	return
}
