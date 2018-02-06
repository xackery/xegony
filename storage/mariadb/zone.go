package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	zoneTable  = "zone"
	zoneFields = "short_name, file_name, long_name, map_file_name, safe_x, safe_y, safe_z, graveyard_id, min_level, min_status, zoneidnumber, version, timezone, maxclients, ruleset, note, underworld, minclip, maxclip, fog_minclip, fog_maxclip, fog_blue, fog_red, fog_green, sky, ztype, zone_exp_multiplier, walkspeed, time_type, fog_red1, fog_green1, fog_blue1, fog_minclip1, fog_maxclip1, fog_red2, fog_green2, fog_blue2, fog_minclip2, fog_maxclip2, fog_red3, fog_green3, fog_blue3, fog_minclip3, fog_maxclip3, fog_red4, fog_green4, fog_blue4, fog_minclip4, fog_maxclip4, fog_density, flag_needed, canbind, cancombat, canlevitate, castoutdoor, hotzone, insttype, shutdowndelay, peqzone, expansion, suspendbuffs, rain_chance1, rain_chance2, rain_chance3, rain_chance4, rain_duration1, rain_duration2, rain_duration3, rain_duration4, snow_chance1, snow_chance2, snow_chance3, snow_chance4, snow_duration1, snow_duration2, snow_duration3, snow_duration4, gravity, type, skylock"
	zoneBinds  = ":short_name, :file_name, :long_name, :map_file_name, :safe_x, :safe_y, :safe_z, :graveyard_id, :min_level, :min_status, :zoneidnumber, :version, :timezone, :maxclients, :ruleset, :note, :underworld, :minclip, :maxclip, :fog_minclip, :fog_maxclip, :fog_blue, :fog_red, :fog_green, :sky, :ztype, :zone_exp_multiplier, :walkspeed, :time_type, :fog_red1, :fog_green1, :fog_blue1, :fog_minclip1, :fog_maxclip1, :fog_red2, :fog_green2, :fog_blue2, :fog_minclip2, :fog_maxclip2, :fog_red3, :fog_green3, :fog_blue3, :fog_minclip3, :fog_maxclip3, :fog_red4, :fog_green4, :fog_blue4, :fog_minclip4, :fog_maxclip4, :fog_density, :flag_needed, :canbind, :cancombat, :canlevitate, :castoutdoor, :hotzone, :insttype, :shutdowndelay, :peqzone, :expansion, :suspendbuffs, :rain_chance1, :rain_chance2, :rain_chance3, :rain_chance4, :rain_duration1, :rain_duration2, :rain_duration3, :rain_duration4, :snow_chance1, :snow_chance2, :snow_chance3, :snow_chance4, :snow_duration1, :snow_duration2, :snow_duration3, :snow_duration4, :gravity, :type, :skylock"
)

//GetZoneByShortName will grab data from storage
func (s *Storage) GetZoneByShortName(zone *model.Zone) (err error) {
	query := fmt.Sprintf("SELECT id, %s FROM %s WHERE short_name = ?", zoneFields, zoneTable)
	err = s.db.Get(zone, query, zone.ShortName.String)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//GetZone will grab data from storage
func (s *Storage) GetZone(zone *model.Zone) (err error) {
	query := fmt.Sprintf("SELECT id, %s FROM %s WHERE id = ?", zoneFields, zoneTable)
	err = s.db.Get(zone, query, zone.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateZone will grab data from storage
func (s *Storage) CreateZone(zone *model.Zone) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", zoneTable, zoneFields, zoneBinds)
	result, err := s.db.NamedExec(query, zone)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	zoneID, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	zone.ID = zoneID
	return
}

//ListZone will grab data from storage
func (s *Storage) ListZone(page *model.Page) (zones []*model.Zone, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "zoneidnumber"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT id, %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", zoneFields, zoneTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		zone := model.Zone{}
		if err = rows.StructScan(&zone); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		zones = append(zones, &zone)
	}
	return
}

//ListZoneTotalCount will grab data from storage
func (s *Storage) ListZoneTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(id) FROM %s", zoneTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListZoneBySearch will grab data from storage
func (s *Storage) ListZoneBySearch(page *model.Page, zone *model.Zone) (zones []*model.Zone, err error) {

	field := ""

	if len(zone.ShortName.String) > 0 {
		field += `short_name LIKE :short_name OR`
		zone.ShortName.String = fmt.Sprintf("%%%s%%", zone.ShortName.String)
		zone.ShortName.Valid = true
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT id, %s FROM %s WHERE %s LIMIT %d OFFSET %d", zoneFields, zoneTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, zone)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		zone := model.Zone{}
		if err = rows.StructScan(&zone); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		zones = append(zones, &zone)
	}
	return
}

//ListZoneBySearchTotalCount will grab data from storage
func (s *Storage) ListZoneBySearchTotalCount(zone *model.Zone) (count int64, err error) {
	field := ""
	if len(zone.ShortName.String) > 0 {
		field += `short_name LIKE :short_name OR`
		zone.ShortName.String = fmt.Sprintf("%%%s%%", zone.ShortName.String)
		zone.ShortName.Valid = true
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", zoneTable, field)

	rows, err := s.db.NamedQuery(query, zone)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
	}
	return
}

//EditZone will grab data from storage
func (s *Storage) EditZone(zone *model.Zone) (err error) {

	prevZone := &model.Zone{
		ID: zone.ID,
	}
	err = s.GetZone(prevZone)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous zone")
		return
	}

	field := ""
	/*	if len(zone.Name) > 0 && prevZone.Name != zone.Name {
			field += "name = :name, "
		}

		if len(zone.Charname) > 0 && prevZone.Charname != zone.Charname {
			field += "charname = :charname, "
		}

		if len(zone.Name) > 0 && prevZone.Name != zone.Name {
			field += "name=:name, "
		}
		if len(zone.Charname) > 0 && prevZone.Charname != zone.Charname {
			field += "charname=:charname, "
		}
		if zone.Sharedplat > 0 && prevZone.Sharedplat != zone.Sharedplat {
			field += "sharedplat=:sharedplat, "
		}
		if len(zone.Password) > 0 && prevZone.Password != zone.Password {
			field += "password=:password, "
		}
		if zone.Status > 0 && prevZone.Status != zone.Status {
			field += "status=:status, "
		}
		if zone.LszoneID.Int64 > 0 && prevZone.LszoneID != zone.LszoneID {
			field += "lszone_id=:lszone_id, "
		}
		if zone.Gmspeed > 0 && prevZone.Gmspeed != zone.Gmspeed {
			field += "gmspeed=:gmspeed, "
		}
		if zone.Revoked > 0 && prevZone.Revoked != zone.Revoked {
			field += "revoked=:revoked, "
		}
		if zone.Karma > 0 && prevZone.Karma != zone.Karma {
			field += "karma=:karma, "
		}
		if len(zone.MiniloginIP) > 0 && prevZone.MiniloginIP != zone.MiniloginIP {
			field += "minilogin_ip=:minilogin_ip, "
		}
		if zone.Hideme > 0 && prevZone.Hideme != zone.Hideme {
			field += "hideme=:hideme, "
		}
		if zone.Rulesflag > 0 && prevZone.Rulesflag != zone.Rulesflag {
			field += "rulesflag=:rulesflag, "
		}
		if !zone.Suspendeduntil.IsZero() && prevZone.Suspendeduntil != zone.Suspendeduntil {
			field += "suspendeduntil=:suspendeduntil, "
		}
		if zone.TimeCreation > 0 && prevZone.TimeCreation != zone.TimeCreation {
			field += "time_creation=:time_creation, "
		}
		if zone.Expansion > 0 && prevZone.Expansion != zone.Expansion {
			field += "expansion=:expansion, "
		}
		if len(zone.BanReason.String) > 0 && prevZone.BanReason != zone.BanReason {
			field += "ban_reason=:ban_reason, "
		}
		if len(zone.SuspendReason.String) > 0 && prevZone.SuspendReason != zone.SuspendReason {
			field += "suspend_reason=:suspend_reason, "
		}
	*/
	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", zoneTable, field)
	result, err := s.db.NamedExec(query, zone)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//DeleteZone will grab data from storage
func (s *Storage) DeleteZone(zone *model.Zone) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", zoneTable)
	result, err := s.db.Exec(query, zone.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//createTableZone will grab data from storage
func (s *Storage) createTableZone() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE zone (
  short_name varchar(32) DEFAULT NULL,
  id int(10) NOT NULL AUTO_INCREMENT,
  file_name varchar(16) DEFAULT NULL,
  long_name text NOT NULL,
  map_file_name varchar(100) DEFAULT NULL,
  safe_x float NOT NULL DEFAULT '0',
  safe_y float NOT NULL DEFAULT '0',
  safe_z float NOT NULL DEFAULT '0',
  graveyard_id float NOT NULL DEFAULT '0',
  min_level tinyint(3) unsigned NOT NULL DEFAULT '0',
  min_status tinyint(3) unsigned NOT NULL DEFAULT '0',
  zoneidnumber int(4) NOT NULL DEFAULT '0',
  version tinyint(3) unsigned NOT NULL DEFAULT '0',
  timezone int(5) NOT NULL DEFAULT '0',
  maxclients int(5) NOT NULL DEFAULT '0',
  ruleset int(10) unsigned NOT NULL DEFAULT '0',
  note varchar(80) DEFAULT NULL,
  underworld float NOT NULL DEFAULT '0',
  minclip float NOT NULL DEFAULT '450',
  maxclip float NOT NULL DEFAULT '450',
  fog_minclip float NOT NULL DEFAULT '450',
  fog_maxclip float NOT NULL DEFAULT '450',
  fog_blue tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_red tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_green tinyint(3) unsigned NOT NULL DEFAULT '0',
  sky tinyint(3) unsigned NOT NULL DEFAULT '1',
  ztype tinyint(3) unsigned NOT NULL DEFAULT '1',
  zone_exp_multiplier decimal(6,2) NOT NULL DEFAULT '0.00',
  walkspeed float NOT NULL DEFAULT '0.4',
  time_type tinyint(3) unsigned NOT NULL DEFAULT '2',
  fog_red1 tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_green1 tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_blue1 tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_minclip1 float NOT NULL DEFAULT '450',
  fog_maxclip1 float NOT NULL DEFAULT '450',
  fog_red2 tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_green2 tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_blue2 tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_minclip2 float NOT NULL DEFAULT '450',
  fog_maxclip2 float NOT NULL DEFAULT '450',
  fog_red3 tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_green3 tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_blue3 tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_minclip3 float NOT NULL DEFAULT '450',
  fog_maxclip3 float NOT NULL DEFAULT '450',
  fog_red4 tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_green4 tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_blue4 tinyint(3) unsigned NOT NULL DEFAULT '0',
  fog_minclip4 float NOT NULL DEFAULT '450',
  fog_maxclip4 float NOT NULL DEFAULT '450',
  fog_density float NOT NULL DEFAULT '0',
  flag_needed varchar(128) NOT NULL DEFAULT '',
  canbind tinyint(4) NOT NULL DEFAULT '1',
  cancombat tinyint(4) NOT NULL DEFAULT '1',
  canlevitate tinyint(4) NOT NULL DEFAULT '1',
  castoutdoor tinyint(4) NOT NULL DEFAULT '1',
  hotzone tinyint(3) unsigned NOT NULL DEFAULT '0',
  insttype tinyint(1) unsigned zerofill NOT NULL DEFAULT '0',
  shutdowndelay bigint(16) unsigned NOT NULL DEFAULT '5000',
  peqzone tinyint(4) NOT NULL DEFAULT '1',
  expansion tinyint(3) NOT NULL DEFAULT '0',
  suspendbuffs tinyint(1) unsigned NOT NULL DEFAULT '0',
  rain_chance1 int(4) NOT NULL DEFAULT '0',
  rain_chance2 int(4) NOT NULL DEFAULT '0',
  rain_chance3 int(4) NOT NULL DEFAULT '0',
  rain_chance4 int(4) NOT NULL DEFAULT '0',
  rain_duration1 int(4) NOT NULL DEFAULT '0',
  rain_duration2 int(4) NOT NULL DEFAULT '0',
  rain_duration3 int(4) NOT NULL DEFAULT '0',
  rain_duration4 int(4) NOT NULL DEFAULT '0',
  snow_chance1 int(4) NOT NULL DEFAULT '0',
  snow_chance2 int(4) NOT NULL DEFAULT '0',
  snow_chance3 int(4) NOT NULL DEFAULT '0',
  snow_chance4 int(4) NOT NULL DEFAULT '0',
  snow_duration1 int(4) NOT NULL DEFAULT '0',
  snow_duration2 int(4) NOT NULL DEFAULT '0',
  snow_duration3 int(4) NOT NULL DEFAULT '0',
  snow_duration4 int(4) NOT NULL DEFAULT '0',
  gravity float NOT NULL DEFAULT '0.4',
  type int(3) NOT NULL DEFAULT '0',
  skylock tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (id),
  KEY zoneidnumber (zoneidnumber),
  KEY zonename (short_name)
) ENGINE=InnoDB AUTO_INCREMENT=5895 DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
