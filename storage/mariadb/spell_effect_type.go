package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSpellEffectType will grab data from storage
func (s *Storage) GetSpellEffectType(spellEffectType *model.SpellEffectType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSpellEffectType will grab data from storage
func (s *Storage) CreateSpellEffectType(spellEffectType *model.SpellEffectType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellEffectType will grab data from storage
func (s *Storage) ListSpellEffectType(page *model.Page) (spellEffectTypes []*model.SpellEffectType, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellEffectTypeTotalCount will grab data from storage
func (s *Storage) ListSpellEffectTypeTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellEffectTypeBySearch will grab data from storage
func (s *Storage) ListSpellEffectTypeBySearch(page *model.Page, spellEffectType *model.SpellEffectType) (spellEffectTypes []*model.SpellEffectType, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellEffectTypeBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellEffectTypeBySearchTotalCount(spellEffectType *model.SpellEffectType) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSpellEffectType will grab data from storage
func (s *Storage) EditSpellEffectType(spellEffectType *model.SpellEffectType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSpellEffectType will grab data from storage
func (s *Storage) DeleteSpellEffectType(spellEffectType *model.SpellEffectType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
