package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSpellTargetType will grab data from storage
func (s *Storage) GetSpellTargetType(spellTargetType *model.SpellTargetType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSpellTargetType will grab data from storage
func (s *Storage) CreateSpellTargetType(spellTargetType *model.SpellTargetType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellTargetType will grab data from storage
func (s *Storage) ListSpellTargetType(page *model.Page) (spellTargetTypes []*model.SpellTargetType, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellTargetTypeTotalCount will grab data from storage
func (s *Storage) ListSpellTargetTypeTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellTargetTypeBySearch will grab data from storage
func (s *Storage) ListSpellTargetTypeBySearch(page *model.Page, spellTargetType *model.SpellTargetType) (spellTargetTypes []*model.SpellTargetType, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellTargetTypeBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellTargetTypeBySearchTotalCount(spellTargetType *model.SpellTargetType) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSpellTargetType will grab data from storage
func (s *Storage) EditSpellTargetType(spellTargetType *model.SpellTargetType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSpellTargetType will grab data from storage
func (s *Storage) DeleteSpellTargetType(spellTargetType *model.SpellTargetType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
