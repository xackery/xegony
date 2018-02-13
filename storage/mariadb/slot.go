package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSlot will grab data from storage
func (s *Storage) GetSlot(slot *model.Slot) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSlot will grab data from storage
func (s *Storage) CreateSlot(slot *model.Slot) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSlot will grab data from storage
func (s *Storage) ListSlot(page *model.Page) (slots []*model.Slot, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSlotTotalCount will grab data from storage
func (s *Storage) ListSlotTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSlotByBit will grab data from storage
func (s *Storage) ListSlotByBit(page *model.Page, slot *model.Slot) (slots []*model.Slot, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSlotByBitTotalCount will grab data from storage
func (s *Storage) ListSlotByBitTotalCount(slot *model.Slot) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSlotBySearch will grab data from storage
func (s *Storage) ListSlotBySearch(page *model.Page, slot *model.Slot) (slots []*model.Slot, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSlotBySearchTotalCount will grab data from storage
func (s *Storage) ListSlotBySearchTotalCount(slot *model.Slot) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSlot will grab data from storage
func (s *Storage) EditSlot(slot *model.Slot) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSlot will grab data from storage
func (s *Storage) DeleteSlot(slot *model.Slot) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
