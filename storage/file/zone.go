package file

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetZone will grab data from storage
func (s *Storage) GetZone(zone *model.Zone) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateZone will grab data from storage
func (s *Storage) CreateZone(zone *model.Zone) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZone will grab data from storage
func (s *Storage) ListZone(page *model.Page) (zones []*model.Zone, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneTotalCount will grab data from storage
func (s *Storage) ListZoneTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneBySearch will grab data from storage
func (s *Storage) ListZoneBySearch(page *model.Page, zone *model.Zone) (zones []*model.Zone, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListZoneBySearchTotalCount will grab data from storage
func (s *Storage) ListZoneBySearchTotalCount(zone *model.Zone) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditZone will grab data from storage
func (s *Storage) EditZone(zone *model.Zone) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteZone will grab data from storage
func (s *Storage) DeleteZone(zone *model.Zone) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
