package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

func (s *Storage) GetZone(zoneId int64) (zone *model.Zone, err error) {
	zone = &model.Zone{}
	err = s.db.Get(zone, "SELECT id, short_name, long_name, min_status, zoneidnumber FROM zone WHERE zoneidnumber = ?", zoneId)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateZone(zone *model.Zone) (err error) {
	if zone == nil {
		err = fmt.Errorf("Must provide zone")
		return
	}

	result, err := s.db.NamedExec(`INSERT INTO zone(short_name, long_name, min_status, zoneidnumber)
		VALUES (:short_name, :long_name, :min_status, :zoneidnumber)`, zone)
	if err != nil {
		return
	}
	zoneId, err := result.LastInsertId()
	if err != nil {
		return
	}
	zone.Id = zoneId
	return
}

func (s *Storage) ListZone() (zones []*model.Zone, err error) {
	rows, err := s.db.Queryx(`SELECT id, short_name, long_name, min_status, zoneidnumber FROM zone ORDER BY long_name DESC`)
	if err != nil {
		return
	}

	for rows.Next() {
		zone := model.Zone{}
		if err = rows.StructScan(&zone); err != nil {
			return
		}
		zones = append(zones, &zone)
	}
	return
}

func (s *Storage) EditZone(zoneId int64, zone *model.Zone) (err error) {
	zone.ZoneIdNumber = zoneId
	result, err := s.db.NamedExec(`UPDATE zone SET short_name=:short_name, long_name=:long_name, min_status=:min_status, zoneidnumber=:zoneidnumber WHERE zoneidnumber = :zoneidnumber`, zone)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}

func (s *Storage) DeleteZone(zoneId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM zone WHERE zoneidnumber = ?`, zoneId)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}
