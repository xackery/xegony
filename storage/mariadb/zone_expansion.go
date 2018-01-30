package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetZoneExpansion will grab data from storage
func (s *Storage) GetZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateZoneExpansion will grab data from storage
func (s *Storage) CreateZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneExpansion will grab data from storage
func (s *Storage) ListZoneExpansion(page *model.Page) (zoneExpansions []*model.ZoneExpansion, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneExpansionTotalCount will grab data from storage
func (s *Storage) ListZoneExpansionTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneExpansionBySearch will grab data from storage
func (s *Storage) ListZoneExpansionBySearch(page *model.Page, zoneExpansion *model.ZoneExpansion) (zoneExpansions []*model.ZoneExpansion, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneExpansionBySearchTotalCount will grab data from storage
func (s *Storage) ListZoneExpansionBySearchTotalCount(zoneExpansion *model.ZoneExpansion) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditZoneExpansion will grab data from storage
func (s *Storage) EditZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteZoneExpansion will grab data from storage
func (s *Storage) DeleteZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
