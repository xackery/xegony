package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	zoneSets   = `zone.short_name=:short_name, zone.file_name=:file_name, zone.long_name=:long_name, zone.map_file_name=:map_file_name, zone.safe_x=:safe_x, zone.safe_y=:safe_y, zone.safe_z=:safe_z, zone.graveyard_id=:graveyard_id, zone.min_level=:min_level, zone.min_status=:min_status, zone.zoneidnumber=:zoneidnumber, zone.version=:version, zone.timezone=:timezone, zone.maxclients=:maxclients, zone.ruleset=:ruleset, zone.note=:note, zone.underworld=:underworld, zone.minclip=:minclip, zone.maxclip=:maxclip, zone.fog_minclip=:fog_minclip, zone.fog_maxclip=:fog_maxclip, zone.fog_blue=:fog_blue, zone.fog_red=:fog_red, zone.fog_green=:fog_green, zone.sky=:sky, zone.ztype=:ztype, zone.zone_exp_multiplier=:zone_exp_multiplier, zone.walkspeed=:walkspeed, zone.time_type=:time_type, zone.fog_red1=:fog_red1, zone.fog_green1=:fog_green1, zone.fog_blue1=:fog_blue1, zone.fog_minclip1=:fog_minclip1, zone.fog_maxclip1=:fog_maxclip1, zone.fog_red2=:fog_red2, zone.fog_green2=:fog_green2, zone.fog_blue2=:fog_blue2, zone.fog_minclip2=:fog_minclip2, zone.fog_maxclip2=:fog_maxclip2, zone.fog_red3=:fog_red3, zone.fog_green3=:fog_green3, zone.fog_blue3=:fog_blue3, zone.fog_minclip3=:fog_minclip3, zone.fog_maxclip3=:fog_maxclip3, zone.fog_red4=:fog_red4, zone.fog_green4=:fog_green4, zone.fog_blue4=:fog_blue4, zone.fog_minclip4=:fog_minclip4, zone.fog_maxclip4=:fog_maxclip4, zone.fog_density=:fog_density, zone.flag_needed=:flag_needed, zone.canbind=:canbind, zone.cancombat=:cancombat, zone.canlevitate=:canlevitate, zone.castoutdoor=:castoutdoor, zone.hotzone=:hotzone, zone.insttype=:insttype, zone.shutdowndelay=:shutdowndelay, zone.peqzone=:peqzone, zone.expansion=:expansion, zone.suspendbuffs=:suspendbuffs, zone.rain_chance1=:rain_chance1, zone.rain_chance2=:rain_chance2, zone.rain_chance3=:rain_chance3, zone.rain_chance4=:rain_chance4, zone.rain_duration1=:rain_duration1, zone.rain_duration2=:rain_duration2, zone.rain_duration3=:rain_duration3, zone.rain_duration4=:rain_duration4, zone.snow_chance1=:snow_chance1, zone.snow_chance2=:snow_chance2, zone.snow_chance3=:snow_chance3, zone.snow_chance4=:snow_chance4, zone.snow_duration1=:snow_duration1, zone.snow_duration2=:snow_duration2, zone.snow_duration3=:snow_duration3, zone.snow_duration4=:snow_duration4, zone.gravity=:gravity, zone.type=:type, zone.skylock=:skylock`
	zoneFields = `zone.short_name, zone.file_name, zone.long_name, zone.map_file_name, zone.safe_x, zone.safe_y, zone.safe_z, zone.graveyard_id, zone.min_level, zone.min_status, zone.zoneidnumber, zone.version, zone.timezone, zone.maxclients, zone.ruleset, zone.note, zone.underworld, zone.minclip, zone.maxclip, zone.fog_minclip, zone.fog_maxclip, zone.fog_blue, zone.fog_red, zone.fog_green, zone.sky, zone.ztype, zone.zone_exp_multiplier, zone.walkspeed, zone.time_type, zone.fog_red1, zone.fog_green1, zone.fog_blue1, zone.fog_minclip1, zone.fog_maxclip1, zone.fog_red2, zone.fog_green2, zone.fog_blue2, zone.fog_minclip2, zone.fog_maxclip2, zone.fog_red3, zone.fog_green3, zone.fog_blue3, zone.fog_minclip3, zone.fog_maxclip3, zone.fog_red4, zone.fog_green4, zone.fog_blue4, zone.fog_minclip4, zone.fog_maxclip4, zone.fog_density, zone.flag_needed, zone.canbind, zone.cancombat, zone.canlevitate, zone.castoutdoor, zone.hotzone, zone.insttype, zone.shutdowndelay, zone.peqzone, zone.expansion, zone.suspendbuffs, zone.rain_chance1, zone.rain_chance2, zone.rain_chance3, zone.rain_chance4, zone.rain_duration1, zone.rain_duration2, zone.rain_duration3, zone.rain_duration4, zone.snow_chance1, zone.snow_chance2, zone.snow_chance3, zone.snow_chance4, zone.snow_duration1, zone.snow_duration2, zone.snow_duration3, zone.snow_duration4, zone.gravity, zone.type, zone.skylock`
	zoneBinds  = `:short_name, :file_name, :long_name, :map_file_name, :safe_x, :safe_y, :safe_z, :graveyard_id, :min_level, :min_status, :zoneidnumber, :version, :timezone, :maxclients, :ruleset, :note, :underworld, :minclip, :maxclip, :fog_minclip, :fog_maxclip, :fog_blue, :fog_red, :fog_green, :sky, :ztype, :zone_exp_multiplier, :walkspeed, :time_type, :fog_red1, :fog_green1, :fog_blue1, :fog_minclip1, :fog_maxclip1, :fog_red2, :fog_green2, :fog_blue2, :fog_minclip2, :fog_maxclip2, :fog_red3, :fog_green3, :fog_blue3, :fog_minclip3, :fog_maxclip3, :fog_red4, :fog_green4, :fog_blue4, :fog_minclip4, :fog_maxclip4, :fog_density, :flag_needed, :canbind, :cancombat, :canlevitate, :castoutdoor, :hotzone, :insttype, :shutdowndelay, :peqzone, :expansion, :suspendbuffs, :rain_chance1, :rain_chance2, :rain_chance3, :rain_chance4, :rain_duration1, :rain_duration2, :rain_duration3, :rain_duration4, :snow_chance1, :snow_chance2, :snow_chance3, :snow_chance4, :snow_duration1, :snow_duration2, :snow_duration3, :snow_duration4, :gravity, :type, :skylock`
)

//GetZone will grab data from storage
func (s *Storage) GetZone(zoneID int64) (zone *model.Zone, err error) {
	zone = &model.Zone{}
	err = s.db.Get(zone, fmt.Sprintf(`SELECT zone_level_cache.levels, zone.id, %s 
		FROM zone 
		INNER JOIN zone_level_cache.zone_id = zone.zoneidnumber 
		WHERE zoneidnumber = ?`, zoneFields), zoneID)
	if err != nil {
		return
	}
	return
}

//CreateZone will grab data from storage
func (s *Storage) CreateZone(zone *model.Zone) (err error) {
	if zone == nil {
		err = fmt.Errorf("Must provide zone")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO zone(%s)
		VALUES (%s)`, zoneFields, zoneBinds), zone)
	if err != nil {
		return
	}
	zoneID, err := result.LastInsertId()
	if err != nil {
		return
	}
	zone.ID = zoneID
	return
}

//ListZone will grab data from storage
func (s *Storage) ListZone() (zones []*model.Zone, err error) {
	query := fmt.Sprintf(`SELECT zone.id, %s 
		FROM zone ORDER BY long_name ASC`, zoneFields)
	rows, err := s.db.Queryx(query)
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

//ListZoneByHotzone will grab data from storage
func (s *Storage) ListZoneByHotzone() (zones []*model.Zone, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT zone.id, %s FROM zone 
		WHERE hotzone = 1 ORDER BY zone_exp_multiplier DESC`, zoneFields))
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

//EditZone will grab data from storage
func (s *Storage) EditZone(zoneID int64, zone *model.Zone) (err error) {
	zone.ZoneIDNumber = zoneID
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE zone SET %s WHERE zoneidnumber = :zoneidnumber`, zoneSets), zone)
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

//DeleteZone will grab data from storage
func (s *Storage) DeleteZone(zoneID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM zone WHERE zoneidnumber = ?`, zoneID)
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

//createTableZone will grab data from storage
func (s *Storage) createTableZone() (err error) {
	_, err = s.db.Exec(`CREATE TABLE zone (
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
