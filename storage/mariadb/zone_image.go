package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetZoneImage will grab data from storage
func (s *Storage) GetZoneImage(zoneImage *model.ZoneImage) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateZoneImage will grab data from storage
func (s *Storage) CreateZoneImage(zoneImage *model.ZoneImage) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneImage will grab data from storage
func (s *Storage) ListZoneImage(page *model.Page) (zoneImages []*model.ZoneImage, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneImageTotalCount will grab data from storage
func (s *Storage) ListZoneImageTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneImageByBit will grab data from storage
func (s *Storage) ListZoneImageByBit(page *model.Page, zoneImage *model.ZoneImage) (zoneImages []*model.ZoneImage, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneImageByBitTotalCount will grab data from storage
func (s *Storage) ListZoneImageByBitTotalCount(zoneImage *model.ZoneImage) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneImageBySearch will grab data from storage
func (s *Storage) ListZoneImageBySearch(page *model.Page, zoneImage *model.ZoneImage) (zoneImages []*model.ZoneImage, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneImageBySearchTotalCount will grab data from storage
func (s *Storage) ListZoneImageBySearchTotalCount(zoneImage *model.ZoneImage) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditZoneImage will grab data from storage
func (s *Storage) EditZoneImage(zoneImage *model.ZoneImage) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteZoneImage will grab data from storage
func (s *Storage) DeleteZoneImage(zoneImage *model.ZoneImage) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
