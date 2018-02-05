package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSpellTravelType will grab data from storage
func (s *Storage) GetSpellTravelType(spellTravelType *model.SpellTravelType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSpellTravelType will grab data from storage
func (s *Storage) CreateSpellTravelType(spellTravelType *model.SpellTravelType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellTravelType will grab data from storage
func (s *Storage) ListSpellTravelType(page *model.Page) (spellTravelTypes []*model.SpellTravelType, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellTravelTypeTotalCount will grab data from storage
func (s *Storage) ListSpellTravelTypeTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellTravelTypeBySearch will grab data from storage
func (s *Storage) ListSpellTravelTypeBySearch(page *model.Page, spellTravelType *model.SpellTravelType) (spellTravelTypes []*model.SpellTravelType, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSpellTravelTypeBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellTravelTypeBySearchTotalCount(spellTravelType *model.SpellTravelType) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSpellTravelType will grab data from storage
func (s *Storage) EditSpellTravelType(spellTravelType *model.SpellTravelType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSpellTravelType will grab data from storage
func (s *Storage) DeleteSpellTravelType(spellTravelType *model.SpellTravelType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
