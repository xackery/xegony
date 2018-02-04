package memory

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetItem will grab data from storage
func (s *Storage) GetItem(item *model.Item) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateItem will grab data from storage
func (s *Storage) CreateItem(item *model.Item) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListItem will grab data from storage
func (s *Storage) ListItem(page *model.Page) (items []*model.Item, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListItemTotalCount will grab data from storage
func (s *Storage) ListItemTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListItemBySearch will grab data from storage
func (s *Storage) ListItemBySearch(page *model.Page, item *model.Item) (items []*model.Item, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListItemBySearchTotalCount will grab data from storage
func (s *Storage) ListItemBySearchTotalCount(item *model.Item) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditItem will grab data from storage
func (s *Storage) EditItem(item *model.Item) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteItem will grab data from storage
func (s *Storage) DeleteItem(item *model.Item) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
