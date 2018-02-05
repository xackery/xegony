package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	spellEffectFormulasDatabase = []*model.SpellEffectFormula{}
	spellEffectFormulaLock      = sync.RWMutex{}
)

//GetSpellEffectFormula will grab data from storage
func (s *Storage) GetSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula) (err error) {
	spellEffectFormulaLock.RLock()
	defer spellEffectFormulaLock.RUnlock()
	for _, tmpSpellEffectFormula := range spellEffectFormulasDatabase {
		if tmpSpellEffectFormula.ID == spellEffectFormula.ID {
			*spellEffectFormula = *tmpSpellEffectFormula
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSpellEffectFormula will grab data from storage
func (s *Storage) CreateSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula) (err error) {
	spellEffectFormulaLock.Lock()
	defer spellEffectFormulaLock.Unlock()
	for _, tmpSpellEffectFormula := range spellEffectFormulasDatabase {
		if tmpSpellEffectFormula.ID == spellEffectFormula.ID {
			err = fmt.Errorf("spellEffectFormula already exists")
			return
		}
	}
	spellEffectFormulasDatabase = append(spellEffectFormulasDatabase, spellEffectFormula)
	return
}

//ListSpellEffectFormula will grab data from storage
func (s *Storage) ListSpellEffectFormula(page *model.Page) (spellEffectFormulas []*model.SpellEffectFormula, err error) {
	spellEffectFormulaLock.RLock()
	defer spellEffectFormulaLock.RUnlock()

	spellEffectFormulas = make([]*model.SpellEffectFormula, len(spellEffectFormulasDatabase))

	spellEffectFormulas = spellEffectFormulasDatabase

	switch page.OrderBy {
	case "name":
		sort.Slice(spellEffectFormulas, func(i, j int) bool {
			return spellEffectFormulas[i].Name < spellEffectFormulas[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(spellEffectFormulas))
		}
	*/
	return
}

//ListSpellEffectFormulaTotalCount will grab data from storage
func (s *Storage) ListSpellEffectFormulaTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(spellEffectFormulasDatabase))
	return
}

//ListSpellEffectFormulaBySearch will grab data from storage
func (s *Storage) ListSpellEffectFormulaBySearch(page *model.Page, spellEffectFormula *model.SpellEffectFormula) (spellEffectFormulas []*model.SpellEffectFormula, err error) {
	spellEffectFormulaLock.RLock()
	defer spellEffectFormulaLock.RUnlock()

	if len(spellEffectFormula.Name) > 0 {
		for i := range spellEffectFormulasDatabase {
			if strings.Contains(spellEffectFormulasDatabase[i].Name, spellEffectFormula.Name) {
				spellEffectFormulas = append(spellEffectFormulas, spellEffectFormulasDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellEffectFormulas, func(i, j int) bool {
			return spellEffectFormulas[i].Name < spellEffectFormulas[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(spellEffectFormulas))
	//}
	return
}

//ListSpellEffectFormulaBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellEffectFormulaBySearchTotalCount(spellEffectFormula *model.SpellEffectFormula) (count int64, err error) {
	spellEffectFormulaLock.RLock()
	defer spellEffectFormulaLock.RUnlock()

	spellEffectFormulas := []*model.SpellEffectFormula{}
	if len(spellEffectFormula.Name) > 0 {
		for i := range spellEffectFormulasDatabase {
			if strings.Contains(spellEffectFormulasDatabase[i].Name, spellEffectFormula.Name) {
				spellEffectFormulas = append(spellEffectFormulas, spellEffectFormulasDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(spellEffectFormulas))
	return
}

//EditSpellEffectFormula will grab data from storage
func (s *Storage) EditSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula) (err error) {
	spellEffectFormulaLock.Lock()
	defer spellEffectFormulaLock.Unlock()
	for i := range spellEffectFormulasDatabase {
		if spellEffectFormulasDatabase[i].ID == spellEffectFormula.ID {
			*spellEffectFormulasDatabase[i] = *spellEffectFormula
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSpellEffectFormula will grab data from storage
func (s *Storage) DeleteSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula) (err error) {
	spellEffectFormulaLock.Lock()
	defer spellEffectFormulaLock.Unlock()
	indexToDelete := 0
	for i := range spellEffectFormulasDatabase {
		if spellEffectFormulasDatabase[i].ID == spellEffectFormula.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	spellEffectFormulasDatabase[len(spellEffectFormulasDatabase)-1], spellEffectFormulasDatabase[indexToDelete] = spellEffectFormulasDatabase[indexToDelete], spellEffectFormulasDatabase[len(spellEffectFormulasDatabase)-1]
	spellEffectFormulasDatabase = spellEffectFormulasDatabase[:len(spellEffectFormulasDatabase)-1]
	return
}
