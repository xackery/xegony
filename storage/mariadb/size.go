package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSize will grab data from storage
func (s *Storage) GetSize(size *model.Size) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSize will grab data from storage
func (s *Storage) CreateSize(size *model.Size) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSize will grab data from storage
func (s *Storage) ListSize(page *model.Page) (sizes []*model.Size, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSizeTotalCount will grab data from storage
func (s *Storage) ListSizeTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSizeByBit will grab data from storage
func (s *Storage) ListSizeByBit(page *model.Page, size *model.Size) (sizes []*model.Size, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSizeByBitTotalCount will grab data from storage
func (s *Storage) ListSizeByBitTotalCount(size *model.Size) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSizeBySearch will grab data from storage
func (s *Storage) ListSizeBySearch(page *model.Page, size *model.Size) (sizes []*model.Size, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSizeBySearchTotalCount will grab data from storage
func (s *Storage) ListSizeBySearchTotalCount(size *model.Size) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSize will grab data from storage
func (s *Storage) EditSize(size *model.Size) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSize will grab data from storage
func (s *Storage) DeleteSize(size *model.Size) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
