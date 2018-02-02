package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSpell will grab data from storage
func (s *Storage) GetSpell(spell *model.Spell) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSpell will grab data from storage
func (s *Storage) CreateSpell(spell *model.Spell) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpell will grab data from storage
func (s *Storage) ListSpell(page *model.Page) (spells []*model.Spell, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellTotalCount will grab data from storage
func (s *Storage) ListSpellTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellBySearch will grab data from storage
func (s *Storage) ListSpellBySearch(page *model.Page, spell *model.Spell) (spells []*model.Spell, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellBySearchTotalCount(spell *model.Spell) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSpell will grab data from storage
func (s *Storage) EditSpell(spell *model.Spell) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSpell will grab data from storage
func (s *Storage) DeleteSpell(spell *model.Spell) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
