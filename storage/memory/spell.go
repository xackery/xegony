package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	spellsDatabase = []*model.Spell{}
	spellLock      = sync.RWMutex{}
)

//GetSpell will grab data from storage
func (s *Storage) GetSpell(spell *model.Spell) (err error) {
	spellLock.RLock()
	defer spellLock.RUnlock()
	for _, tmpSpell := range spellsDatabase {
		if tmpSpell.ID == spell.ID {
			*spell = *tmpSpell
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSpell will grab data from storage
func (s *Storage) CreateSpell(spell *model.Spell) (err error) {
	spellLock.Lock()
	defer spellLock.Unlock()
	for _, tmpSpell := range spellsDatabase {
		if tmpSpell.ID == spell.ID {
			err = fmt.Errorf("spell already exists")
			return
		}
	}
	spellsDatabase = append(spellsDatabase, spell)
	return
}

//ListSpell will grab data from storage
func (s *Storage) ListSpell(page *model.Page) (spells []*model.Spell, err error) {
	spellLock.RLock()
	defer spellLock.RUnlock()

	spells = make([]*model.Spell, len(spellsDatabase))

	spells = spellsDatabase

	switch page.OrderBy {
	case "short_name":
		sort.Slice(spells, func(i, j int) bool {
			return spells[i].Name.String < spells[j].Name.String
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(spells))
		}
	*/
	return
}

//ListSpellTotalCount will grab data from storage
func (s *Storage) ListSpellTotalCount() (count int64, err error) {
	count = int64(len(spellsDatabase))
	return
}

//ListSpellBySearch will grab data from storage
func (s *Storage) ListSpellBySearch(page *model.Page, spell *model.Spell) (spells []*model.Spell, err error) {
	spellLock.RLock()
	defer spellLock.RUnlock()

	if len(spell.Name.String) > 0 {
		for i := range spellsDatabase {
			if strings.Contains(spellsDatabase[i].Name.String, spell.Name.String) {
				spells = append(spells, spellsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "short_name":
		sort.Slice(spells, func(i, j int) bool {
			return spells[i].Name.String < spells[j].Name.String
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(spells))
	//}
	return
}

//ListSpellBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellBySearchTotalCount(spell *model.Spell) (count int64, err error) {
	spellLock.RLock()
	defer spellLock.RUnlock()

	spells := []*model.Spell{}
	if len(spell.Name.String) > 0 {
		for i := range spellsDatabase {
			if strings.Contains(spellsDatabase[i].Name.String, spell.Name.String) {
				spells = append(spells, spellsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(spells))
	return
}

//EditSpell will grab data from storage
func (s *Storage) EditSpell(spell *model.Spell) (err error) {
	spellLock.Lock()
	defer spellLock.Unlock()
	for i := range spellsDatabase {
		if spellsDatabase[i].ID == spell.ID {
			*spellsDatabase[i] = *spell
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSpell will grab data from storage
func (s *Storage) DeleteSpell(spell *model.Spell) (err error) {
	spellLock.Lock()
	defer spellLock.Unlock()
	indexToDelete := 0
	for i := range spellsDatabase {
		if spellsDatabase[i].ID == spell.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	spellsDatabase[len(spellsDatabase)-1], spellsDatabase[indexToDelete] = spellsDatabase[indexToDelete], spellsDatabase[len(spellsDatabase)-1]
	spellsDatabase = spellsDatabase[:len(spellsDatabase)-1]
	return
}
