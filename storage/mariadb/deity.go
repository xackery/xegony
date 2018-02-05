package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetDeity will grab data from storage
func (s *Storage) GetDeity(deity *model.Deity) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//GetDeityBySpell will grab data from storage
func (s *Storage) GetDeityBySpell(spell *model.Spell, deity *model.Deity) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateDeity will grab data from storage
func (s *Storage) CreateDeity(deity *model.Deity) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListDeity will grab data from storage
func (s *Storage) ListDeity(page *model.Page) (deitys []*model.Deity, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListDeityTotalCount will grab data from storage
func (s *Storage) ListDeityTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListDeityByBit will grab data from storage
func (s *Storage) ListDeityByBit(page *model.Page, deity *model.Deity) (deitys []*model.Deity, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListDeityByBitTotalCount will grab data from storage
func (s *Storage) ListDeityByBitTotalCount(deity *model.Deity) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListDeityBySearch will grab data from storage
func (s *Storage) ListDeityBySearch(page *model.Page, deity *model.Deity) (deitys []*model.Deity, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListDeityBySearchTotalCount will grab data from storage
func (s *Storage) ListDeityBySearchTotalCount(deity *model.Deity) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditDeity will grab data from storage
func (s *Storage) EditDeity(deity *model.Deity) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteDeity will grab data from storage
func (s *Storage) DeleteDeity(deity *model.Deity) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
