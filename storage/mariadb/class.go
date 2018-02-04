package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetClass will grab data from storage
func (s *Storage) GetClass(class *model.Class) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateClass will grab data from storage
func (s *Storage) CreateClass(class *model.Class) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListClass will grab data from storage
func (s *Storage) ListClass(page *model.Page) (classs []*model.Class, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListClassTotalCount will grab data from storage
func (s *Storage) ListClassTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListClassByBit will grab data from storage
func (s *Storage) ListClassByBit(page *model.Page, class *model.Class) (classs []*model.Class, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListClassByBitTotalCount will grab data from storage
func (s *Storage) ListClassByBitTotalCount(class *model.Class) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListClassBySearch will grab data from storage
func (s *Storage) ListClassBySearch(page *model.Page, class *model.Class) (classs []*model.Class, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListClassBySearchTotalCount will grab data from storage
func (s *Storage) ListClassBySearchTotalCount(class *model.Class) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditClass will grab data from storage
func (s *Storage) EditClass(class *model.Class) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteClass will grab data from storage
func (s *Storage) DeleteClass(class *model.Class) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
