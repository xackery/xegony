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
	spellEffectFormulaLock = sync.RWMutex{}
)

//GetSpellEffectFormula will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula) (err error) {
	spellEffectFormulaLock.Lock()
	defer spellEffectFormulaLock.Unlock()
	spellEffectFormulasDatabase, err := s.readSpellEffectFormulaFile()
	if err != nil {
		return
	}
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
	spellEffectFormulasDatabase, err := s.readSpellEffectFormulaFile()
	if err != nil {
		return
	}
	for _, tmpSpellEffectFormula := range spellEffectFormulasDatabase {
		if tmpSpellEffectFormula.ID == spellEffectFormula.ID {
			err = fmt.Errorf("spellEffectFormula already exists")
			return
		}
	}
	spellEffectFormulasDatabase = append(spellEffectFormulasDatabase, spellEffectFormula)
	err = s.writeSpellEffectFormulaFile(spellEffectFormulasDatabase)
	if err != nil {
		return
	}
	return
}

//ListSpellEffectFormula will grab data from storage
func (s *Storage) ListSpellEffectFormula(page *model.Page) (spellEffectFormulas []*model.SpellEffectFormula, err error) {
	spellEffectFormulaLock.Lock()
	defer spellEffectFormulaLock.Unlock()
	spellEffectFormulasDatabase, err := s.readSpellEffectFormulaFile()
	if err != nil {
		return
	}

	spellEffectFormulas = make([]*model.SpellEffectFormula, len(spellEffectFormulasDatabase))

	spellEffectFormulas = spellEffectFormulasDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
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

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(spellEffectFormulas))
		}
	*/
	return
}

//ListSpellEffectFormulaTotalCount will grab data from storage
func (s *Storage) ListSpellEffectFormulaTotalCount() (count int64, err error) {
	spellEffectFormulaLock.Lock()
	defer spellEffectFormulaLock.Unlock()
	spellEffectFormulasDatabase, err := s.readSpellEffectFormulaFile()
	if err != nil {
		return
	}
	count = int64(len(spellEffectFormulasDatabase))
	return
}

//ListSpellEffectFormulaBySearch will grab data from storage
func (s *Storage) ListSpellEffectFormulaBySearch(page *model.Page, spellEffectFormula *model.SpellEffectFormula) (spellEffectFormulas []*model.SpellEffectFormula, err error) {
	spellEffectFormulaLock.Lock()
	defer spellEffectFormulaLock.Unlock()
	spellEffectFormulasDatabase, err := s.readSpellEffectFormulaFile()
	if err != nil {
		return
	}
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
	spellEffectFormulaLock.Lock()
	defer spellEffectFormulaLock.Unlock()
	spellEffectFormulasDatabase, err := s.readSpellEffectFormulaFile()
	if err != nil {
		return
	}

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
	spellEffectFormulasDatabase, err := s.readSpellEffectFormulaFile()
	if err != nil {
		return
	}

	for i := range spellEffectFormulasDatabase {
		if spellEffectFormulasDatabase[i].ID == spellEffectFormula.ID {
			*spellEffectFormulasDatabase[i] = *spellEffectFormula
			err = s.writeSpellEffectFormulaFile(spellEffectFormulasDatabase)
			if err != nil {
				return
			}
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
	spellEffectFormulasDatabase, err := s.readSpellEffectFormulaFile()
	if err != nil {
		return
	}
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
	err = s.writeSpellEffectFormulaFile(spellEffectFormulasDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readSpellEffectFormulaFile() (spellEffectFormulas []*model.SpellEffectFormula, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			spellEffectFormulas = loadSpellEffectFormulaDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeSpellEffectFormulaFile(spellEffectFormulas)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default spellEffectFormula data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &spellEffectFormulas)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeSpellEffectFormulaFile(spellEffectFormulas []*model.SpellEffectFormula) (err error) {

	bData, err := yaml.Marshal(spellEffectFormulas)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal spellEffectFormulas")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
