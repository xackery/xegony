package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	zoneLevelSets   = `zone_id=:zone_id, levels=:levels`
	zoneLevelFields = `zone_id, levels`
	zoneLevelBinds  = `:zone_id, :levels`
)

//GetZoneLevel will grab data from storage
func (s *Storage) GetZoneLevel(zoneID int64) (zoneLevel *model.ZoneLevel, err error) {
	zoneLevel = &model.ZoneLevel{}
	err = s.db.Get(zoneLevel, fmt.Sprintf(`SELECT %s FROM zone_level_cache
	INNER JOIN zone ON zone.zoneidnumber = zone_level_cache.zone_id
	WHERE zone_level_cache.zone_id = ?`, zoneLevelFields), zoneID)
	if err != nil {
		return
	}
	return
}

//CreateZoneLevel will grab data from storage
func (s *Storage) CreateZoneLevel(zoneLevel *model.ZoneLevel) (err error) {
	if zoneLevel == nil {
		err = fmt.Errorf("Must provide zoneLevel")
		return
	}

	_, err = s.db.NamedExec(fmt.Sprintf(`INSERT INTO zone_level_cache(%s)
		VALUES (%s)`, zoneLevelFields, zoneLevelBinds), zoneLevel)
	if err != nil {
		return
	}
	return
}

//ListZoneLevel will grab data from storage
func (s *Storage) ListZoneLevel() (zoneLevels []*model.ZoneLevel, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM zone_level_cache
	INNER JOIN zone ON zone.zoneidnumber = zone_level_cache.zone_id`, zoneLevelFields))
	if err != nil {
		return
	}

	for rows.Next() {
		zoneLevel := model.ZoneLevel{}
		if err = rows.StructScan(&zoneLevel); err != nil {
			return
		}
		zoneLevels = append(zoneLevels, &zoneLevel)
	}
	return
}

//EditZoneLevel will grab data from storage
func (s *Storage) EditZoneLevel(zoneID int64, zoneLevel *model.ZoneLevel) (err error) {
	zoneLevel.ZoneID = zoneID
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE zone_level_cache SET %s WHERE zone_id = :zone_id`, zoneLevelSets), zoneLevel)
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

//TruncateZoneLevel will grab data from storage
func (s *Storage) TruncateZoneLevel() (err error) {
	_, err = s.db.Exec(`TRUNCATE zone_level_cache`)
	if err != nil {
		return
	}
	return
}

//DeleteZoneLevel will grab data from storage
func (s *Storage) DeleteZoneLevel(zoneID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM zone_level_cache WHERE zone_id = ?`, zoneID)
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

//createTableZoneLevel will grab data from storage
func (s *Storage) createTableZoneLevel() (err error) {
	_, err = s.db.Exec(`CREATE TABLE zone_level_cache (
  zone_id int(11) unsigned NOT NULL,
  levels int(11) unsigned NOT NULL,
  UNIQUE KEY zone_id (zone_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	return
}
