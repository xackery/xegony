package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSpellEffectFormula will grab data from storage
func (s *Storage) GetSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSpellEffectFormula will grab data from storage
func (s *Storage) CreateSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellEffectFormula will grab data from storage
func (s *Storage) ListSpellEffectFormula(page *model.Page) (spellEffectFormulas []*model.SpellEffectFormula, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellEffectFormulaTotalCount will grab data from storage
func (s *Storage) ListSpellEffectFormulaTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellEffectFormulaBySearch will grab data from storage
func (s *Storage) ListSpellEffectFormulaBySearch(page *model.Page, spellEffectFormula *model.SpellEffectFormula) (spellEffectFormulas []*model.SpellEffectFormula, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellEffectFormulaBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellEffectFormulaBySearchTotalCount(spellEffectFormula *model.SpellEffectFormula) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSpellEffectFormula will grab data from storage
func (s *Storage) EditSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSpellEffectFormula will grab data from storage
func (s *Storage) DeleteSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
