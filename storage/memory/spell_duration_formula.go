package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	spellDurationFormulasDatabase = []*model.SpellDurationFormula{}
	spellDurationFormulaLock      = sync.RWMutex{}
)

//GetSpellDurationFormula will grab data from storage
func (s *Storage) GetSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula) (err error) {
	spellDurationFormulaLock.RLock()
	defer spellDurationFormulaLock.RUnlock()
	for _, tmpSpellDurationFormula := range spellDurationFormulasDatabase {
		if tmpSpellDurationFormula.ID == spellDurationFormula.ID {
			*spellDurationFormula = *tmpSpellDurationFormula
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSpellDurationFormula will grab data from storage
func (s *Storage) CreateSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula) (err error) {
	spellDurationFormulaLock.Lock()
	defer spellDurationFormulaLock.Unlock()
	for _, tmpSpellDurationFormula := range spellDurationFormulasDatabase {
		if tmpSpellDurationFormula.ID == spellDurationFormula.ID {
			err = fmt.Errorf("spellDurationFormula already exists")
			return
		}
	}
	spellDurationFormulasDatabase = append(spellDurationFormulasDatabase, spellDurationFormula)
	return
}

//ListSpellDurationFormula will grab data from storage
func (s *Storage) ListSpellDurationFormula(page *model.Page) (spellDurationFormulas []*model.SpellDurationFormula, err error) {
	spellDurationFormulaLock.RLock()
	defer spellDurationFormulaLock.RUnlock()

	spellDurationFormulas = make([]*model.SpellDurationFormula, len(spellDurationFormulasDatabase))

	spellDurationFormulas = spellDurationFormulasDatabase

	switch page.OrderBy {
	case "name":
		sort.Slice(spellDurationFormulas, func(i, j int) bool {
			return spellDurationFormulas[i].Name < spellDurationFormulas[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(spellDurationFormulas))
		}
	*/
	return
}

//ListSpellDurationFormulaTotalCount will grab data from storage
func (s *Storage) ListSpellDurationFormulaTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(spellDurationFormulasDatabase))
	return
}

//ListSpellDurationFormulaBySearch will grab data from storage
func (s *Storage) ListSpellDurationFormulaBySearch(page *model.Page, spellDurationFormula *model.SpellDurationFormula) (spellDurationFormulas []*model.SpellDurationFormula, err error) {
	spellDurationFormulaLock.RLock()
	defer spellDurationFormulaLock.RUnlock()

	if len(spellDurationFormula.Name) > 0 {
		for i := range spellDurationFormulasDatabase {
			if strings.Contains(spellDurationFormulasDatabase[i].Name, spellDurationFormula.Name) {
				spellDurationFormulas = append(spellDurationFormulas, spellDurationFormulasDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellDurationFormulas, func(i, j int) bool {
			return spellDurationFormulas[i].Name < spellDurationFormulas[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(spellDurationFormulas))
	//}
	return
}

//ListSpellDurationFormulaBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellDurationFormulaBySearchTotalCount(spellDurationFormula *model.SpellDurationFormula) (count int64, err error) {
	spellDurationFormulaLock.RLock()
	defer spellDurationFormulaLock.RUnlock()

	spellDurationFormulas := []*model.SpellDurationFormula{}
	if len(spellDurationFormula.Name) > 0 {
		for i := range spellDurationFormulasDatabase {
			if strings.Contains(spellDurationFormulasDatabase[i].Name, spellDurationFormula.Name) {
				spellDurationFormulas = append(spellDurationFormulas, spellDurationFormulasDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(spellDurationFormulas))
	return
}

//EditSpellDurationFormula will grab data from storage
func (s *Storage) EditSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula) (err error) {
	spellDurationFormulaLock.Lock()
	defer spellDurationFormulaLock.Unlock()
	for i := range spellDurationFormulasDatabase {
		if spellDurationFormulasDatabase[i].ID == spellDurationFormula.ID {
			*spellDurationFormulasDatabase[i] = *spellDurationFormula
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSpellDurationFormula will grab data from storage
func (s *Storage) DeleteSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula) (err error) {
	spellDurationFormulaLock.Lock()
	defer spellDurationFormulaLock.Unlock()
	indexToDelete := 0
	for i := range spellDurationFormulasDatabase {
		if spellDurationFormulasDatabase[i].ID == spellDurationFormula.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	spellDurationFormulasDatabase[len(spellDurationFormulasDatabase)-1], spellDurationFormulasDatabase[indexToDelete] = spellDurationFormulasDatabase[indexToDelete], spellDurationFormulasDatabase[len(spellDurationFormulasDatabase)-1]
	spellDurationFormulasDatabase = spellDurationFormulasDatabase[:len(spellDurationFormulasDatabase)-1]
	return
}
