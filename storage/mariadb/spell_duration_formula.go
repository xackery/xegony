package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSpellDurationFormula will grab data from storage
func (s *Storage) GetSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSpellDurationFormula will grab data from storage
func (s *Storage) CreateSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellDurationFormula will grab data from storage
func (s *Storage) ListSpellDurationFormula(page *model.Page) (spellDurationFormulas []*model.SpellDurationFormula, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellDurationFormulaTotalCount will grab data from storage
func (s *Storage) ListSpellDurationFormulaTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellDurationFormulaBySearch will grab data from storage
func (s *Storage) ListSpellDurationFormulaBySearch(page *model.Page, spellDurationFormula *model.SpellDurationFormula) (spellDurationFormulas []*model.SpellDurationFormula, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellDurationFormulaBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellDurationFormulaBySearchTotalCount(spellDurationFormula *model.SpellDurationFormula) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSpellDurationFormula will grab data from storage
func (s *Storage) EditSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSpellDurationFormula will grab data from storage
func (s *Storage) DeleteSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
