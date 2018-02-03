package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSpellAnimation will grab data from storage
func (s *Storage) GetSpellAnimation(spellAnimation *model.SpellAnimation) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSpellAnimation will grab data from storage
func (s *Storage) CreateSpellAnimation(spellAnimation *model.SpellAnimation) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellAnimation will grab data from storage
func (s *Storage) ListSpellAnimation(page *model.Page) (spellAnimations []*model.SpellAnimation, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellAnimationTotalCount will grab data from storage
func (s *Storage) ListSpellAnimationTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellAnimationBySearch will grab data from storage
func (s *Storage) ListSpellAnimationBySearch(page *model.Page, spellAnimation *model.SpellAnimation) (spellAnimations []*model.SpellAnimation, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellAnimationBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellAnimationBySearchTotalCount(spellAnimation *model.SpellAnimation) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSpellAnimation will grab data from storage
func (s *Storage) EditSpellAnimation(spellAnimation *model.SpellAnimation) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSpellAnimation will grab data from storage
func (s *Storage) DeleteSpellAnimation(spellAnimation *model.SpellAnimation) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
