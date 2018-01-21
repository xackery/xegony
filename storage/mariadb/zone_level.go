package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	zoneLevelSets   = `zone_id=:zone_id, levels=:levels, map_aspect=:map_aspect, map_x_offset=:map_x_offset, map_y_offset=:map_y_offset`
	zoneLevelFields = `zone_id, levels, map_aspect, map_x_offset, map_y_offset`
	zoneLevelBinds  = `:zone_id, :levels, :map_aspect, :map_x_offset, :map_y_offset`
)

//GetZoneLevel will grab data from storage
func (s *Storage) GetZoneLevel(zoneLevel *model.ZoneLevel) (err error) {
	zoneLevel = &model.ZoneLevel{}
	err = s.db.Get(zoneLevel, fmt.Sprintf(`SELECT %s FROM zone_level_cache
	INNER JOIN zone ON zone.zoneidnumber = zone_level_cache.zone_id
	WHERE zone_level_cache.zone_id = ?`, zoneLevelFields), zoneLevel.ZoneID)
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
func (s *Storage) EditZoneLevel(zoneLevel *model.ZoneLevel) (err error) {
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
func (s *Storage) DeleteZoneLevel(zoneLevel *model.ZoneLevel) (err error) {
	result, err := s.db.Exec(`DELETE FROM zone_level_cache WHERE zone_id = ?`, zoneLevel.ZoneID)
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
  map_aspect float unsigned NOT NULL DEFAULT '1',
  map_x_offset float NOT NULL DEFAULT '0',
  map_y_offset float NOT NULL DEFAULT '0',
  UNIQUE KEY zone_id (zone_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	return
}
