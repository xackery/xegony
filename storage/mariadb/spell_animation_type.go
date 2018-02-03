package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSpellAnimationType will grab data from storage
func (s *Storage) GetSpellAnimationType(spellAnimationType *model.SpellAnimationType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSpellAnimationType will grab data from storage
func (s *Storage) CreateSpellAnimationType(spellAnimationType *model.SpellAnimationType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellAnimationType will grab data from storage
func (s *Storage) ListSpellAnimationType(page *model.Page) (spellAnimationTypes []*model.SpellAnimationType, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellAnimationTypeTotalCount will grab data from storage
func (s *Storage) ListSpellAnimationTypeTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellAnimationTypeBySearch will grab data from storage
func (s *Storage) ListSpellAnimationTypeBySearch(page *model.Page, spellAnimationType *model.SpellAnimationType) (spellAnimationTypes []*model.SpellAnimationType, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellAnimationTypeBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellAnimationTypeBySearchTotalCount(spellAnimationType *model.SpellAnimationType) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSpellAnimationType will grab data from storage
func (s *Storage) EditSpellAnimationType(spellAnimationType *model.SpellAnimationType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSpellAnimationType will grab data from storage
func (s *Storage) DeleteSpellAnimationType(spellAnimationType *model.SpellAnimationType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
