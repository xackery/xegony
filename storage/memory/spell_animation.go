package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	spellAnimationsDatabase = []*model.SpellAnimation{}
	spellAnimationLock      = sync.RWMutex{}
)

//GetSpellAnimation will grab data from storage
func (s *Storage) GetSpellAnimation(spellAnimation *model.SpellAnimation) (err error) {
	spellAnimationLock.RLock()
	defer spellAnimationLock.RUnlock()
	for _, tmpSpellAnimation := range spellAnimationsDatabase {
		if tmpSpellAnimation.ID == spellAnimation.ID {
			*spellAnimation = *tmpSpellAnimation
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSpellAnimation will grab data from storage
func (s *Storage) CreateSpellAnimation(spellAnimation *model.SpellAnimation) (err error) {
	spellAnimationLock.Lock()
	defer spellAnimationLock.Unlock()
	for _, tmpSpellAnimation := range spellAnimationsDatabase {
		if tmpSpellAnimation.ID == spellAnimation.ID {
			err = fmt.Errorf("spellAnimation already exists")
			return
		}
	}
	spellAnimationsDatabase = append(spellAnimationsDatabase, spellAnimation)
	return
}

//ListSpellAnimation will grab data from storage
func (s *Storage) ListSpellAnimation(page *model.Page) (spellAnimations []*model.SpellAnimation, err error) {
	spellAnimationLock.RLock()
	defer spellAnimationLock.RUnlock()

	spellAnimations = make([]*model.SpellAnimation, len(spellAnimationsDatabase))

	spellAnimations = spellAnimationsDatabase

	switch page.OrderBy {
	case "id":
		sort.Slice(spellAnimations, func(i, j int) bool {
			return spellAnimations[i].ID < spellAnimations[j].ID
		})
	case "name":
		sort.Slice(spellAnimations, func(i, j int) bool {
			return spellAnimations[i].Name < spellAnimations[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(spellAnimations))
		}
	*/
	return
}

//ListSpellAnimationTotalCount will grab data from storage
func (s *Storage) ListSpellAnimationTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(spellAnimationsDatabase))
	return
}

//ListSpellAnimationBySearch will grab data from storage
func (s *Storage) ListSpellAnimationBySearch(page *model.Page, spellAnimation *model.SpellAnimation) (spellAnimations []*model.SpellAnimation, err error) {
	spellAnimationLock.RLock()
	defer spellAnimationLock.RUnlock()

	if len(spellAnimation.Name) > 0 {
		for i := range spellAnimationsDatabase {
			if strings.Contains(spellAnimationsDatabase[i].Name, spellAnimation.Name) {
				spellAnimations = append(spellAnimations, spellAnimationsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "id":
		sort.Slice(spellAnimations, func(i, j int) bool {
			return spellAnimations[i].ID < spellAnimations[j].ID
		})
	case "name":
		sort.Slice(spellAnimations, func(i, j int) bool {
			return spellAnimations[i].Name < spellAnimations[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(spellAnimations))
	//}
	return
}

//ListSpellAnimationBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellAnimationBySearchTotalCount(spellAnimation *model.SpellAnimation) (count int64, err error) {
	spellAnimationLock.RLock()
	defer spellAnimationLock.RUnlock()

	spellAnimations := []*model.SpellAnimation{}
	if len(spellAnimation.Name) > 0 {
		for i := range spellAnimationsDatabase {
			if strings.Contains(spellAnimationsDatabase[i].Name, spellAnimation.Name) {
				spellAnimations = append(spellAnimations, spellAnimationsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(spellAnimations))
	return
}

//EditSpellAnimation will grab data from storage
func (s *Storage) EditSpellAnimation(spellAnimation *model.SpellAnimation) (err error) {
	spellAnimationLock.Lock()
	defer spellAnimationLock.Unlock()
	for i := range spellAnimationsDatabase {
		if spellAnimationsDatabase[i].ID == spellAnimation.ID {
			*spellAnimationsDatabase[i] = *spellAnimation
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSpellAnimation will grab data from storage
func (s *Storage) DeleteSpellAnimation(spellAnimation *model.SpellAnimation) (err error) {
	spellAnimationLock.Lock()
	defer spellAnimationLock.Unlock()
	indexToDelete := 0
	for i := range spellAnimationsDatabase {
		if spellAnimationsDatabase[i].ID == spellAnimation.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	spellAnimationsDatabase[len(spellAnimationsDatabase)-1], spellAnimationsDatabase[indexToDelete] = spellAnimationsDatabase[indexToDelete], spellAnimationsDatabase[len(spellAnimationsDatabase)-1]
	spellAnimationsDatabase = spellAnimationsDatabase[:len(spellAnimationsDatabase)-1]
	return
}
