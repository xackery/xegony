package model

import (
	"database/sql"
	"sort"
)

//Zone represents the zone table, Everquest is split into zones
// swagger:response
type Zone struct {
	Levels   int64   `json:"levels" db:"levels"`
	Modifier float64 `json:"modifier"`

	ShortName         sql.NullString `json:"shortName" db:"short_name"`                  //short_name` varchar(32) DEFAULT NULL,
	ID                int64          `json:"id" db:"id"`                                 //id` int(10) NOT NULL AUTO_INCREMENT,
	FileName          sql.NullString `json:"fileName" db:"file_name"`                    //file_name` varchar(16) DEFAULT NULL,
	LongName          string         `json:"longName" db:"long_name"`                    //long_name` text NOT NULL,
	MapFileName       sql.NullString `json:"mapFileName" db:"map_file_name"`             //map_file_name` varchar(100) DEFAULT NULL,
	SafeX             float64        `json:"safeX" db:"safe_x"`                          //safe_x` float NOT NULL DEFAULT '0',
	SafeY             float64        `json:"safeY" db:"safe_y"`                          //safe_y` float NOT NULL DEFAULT '0',
	SafeZ             float64        `json:"safeZ" db:"safe_z"`                          //safe_z` float NOT NULL DEFAULT '0',
	GraveyardID       float64        `json:"graveyardID" db:"graveyard_id"`              //graveyard_id` float NOT NULL DEFAULT '0',
	MinLevel          int64          `json:"minLevel" db:"min_level"`                    //min_level` tinyint(3) unsigned NOT NULL DEFAULT '0',
	MinStatus         int64          `json:"minStatus" db:"min_status"`                  //min_status` tinyint(3) unsigned NOT NULL DEFAULT '0',
	ZoneIDNumber      int64          `json:"zoneidnumber" db:"zoneidnumber"`             //zoneidnumber` int(4) NOT NULL DEFAULT '0',
	Version           int64          `json:"version" db:"version"`                       //version` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Timezone          int64          `json:"timezone" db:"timezone"`                     //timezone` int(5) NOT NULL DEFAULT '0',
	MaxClients        int64          `json:"maxclients" db:"maxclients"`                 //maxclients` int(5) NOT NULL DEFAULT '0',
	Ruleset           int64          `json:"ruleset" db:"ruleset"`                       //ruleset` int(10) unsigned NOT NULL DEFAULT '0',
	Note              sql.NullString `json:"note" db:"note"`                             //note` varchar(80) DEFAULT NULL,
	Underworld        float64        `json:"underworld" db:"underworld"`                 //underworld` float NOT NULL DEFAULT '0',
	Minclip           float64        `json:"minclip" db:"minclip"`                       //minclip` float NOT NULL DEFAULT '450',
	Maxclip           float64        `json:"maxclip" db:"maxclip"`                       //maxclip` float NOT NULL DEFAULT '450',
	FogMinclip        float64        `json:"fogMinclip" db:"fog_minclip"`                //fog_minclip` float NOT NULL DEFAULT '450',
	FogMaxclip        float64        `json:"fogMaxclip" db:"fog_maxclip"`                //fog_maxclip` float NOT NULL DEFAULT '450',
	FogBlue           int64          `json:"fogBlue" db:"fog_blue"`                      //fog_blue` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogRed            int64          `json:"fogRed" db:"fog_red"`                        //fog_red` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen          int64          `json:"fogGreen" db:"fog_green"`                    //fog_green` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Sky               int64          `json:"sky" db:"sky"`                               //sky` tinyint(3) unsigned NOT NULL DEFAULT '1',
	Ztype             int64          `json:"ztype" db:"ztype"`                           //ztype` tinyint(3) unsigned NOT NULL DEFAULT '1',
	ZoneExpMultiplier float64        `json:"zoneExpMultiplier" db:"zone_exp_multiplier"` //zone_exp_multiplier` decimal(6,2) NOT NULL DEFAULT '0.00',
	Walkspeed         float64        `json:"walkspeed" db:"walkspeed"`                   //walkspeed` float NOT NULL DEFAULT '0.4',
	TimeType          int64          `json:"timeType" db:"time_type"`                    //time_type` tinyint(3) unsigned NOT NULL DEFAULT '2',
	FogRed1           int64          `json:"fogRed1" db:"fog_red1"`                      //fog_red1` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen1         int64          `json:"fogGreen1" db:"fog_green1"`                  //fog_green1` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogBlue1          int64          `json:"fogBlue1" db:"fog_blue1"`                    //fog_blue1` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogMinclip1       float64        `json:"fogMinclip1" db:"fog_minclip1"`              //fog_minclip1` float NOT NULL DEFAULT '450',
	FogMaxclip1       float64        `json:"fogMaxclip1" db:"fog_maxclip1"`              //fog_maxclip1` float NOT NULL DEFAULT '450',
	FogRed2           int64          `json:"fogRed2" db:"fog_red2"`                      //fog_red2` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen2         int64          `json:"fogGreen2" db:"fog_green2"`                  //fog_green2` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogBlue2          int64          `json:"fogBlue2" db:"fog_blue2"`                    //fog_blue2` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogMinclip2       float64        `json:"fogMinclip2" db:"fog_minclip2"`              //fog_minclip2` float NOT NULL DEFAULT '450',
	FogMaxclip2       float64        `json:"fogMaxclip2" db:"fog_maxclip2"`              //fog_maxclip2` float NOT NULL DEFAULT '450',
	FogRed3           int64          `json:"fogRed3" db:"fog_red3"`                      //fog_red3` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen3         int64          `json:"fogGreen3" db:"fog_green3"`                  //fog_green3` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogBlue3          int64          `json:"fogBlue3" db:"fog_blue3"`                    //fog_blue3` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogMinclip3       float64        `json:"fogMinclip3" db:"fog_minclip3"`              //fog_minclip3` float NOT NULL DEFAULT '450',
	FogMaxclip3       float64        `json:"fogMaxclip3" db:"fog_maxclip3"`              //fog_maxclip3` float NOT NULL DEFAULT '450',
	FogRed4           int64          `json:"fogRed4" db:"fog_red4"`                      //fog_red4` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen4         int64          `json:"fogGreen4" db:"fog_green4"`                  //fog_green4` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogBlue4          int64          `json:"fogBlue4" db:"fog_blue4"`                    //fog_blue4` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogMinclip4       float64        `json:"fogMinclip4" db:"fog_minclip4"`              //fog_minclip4` float NOT NULL DEFAULT '450',
	FogMaxclip4       float64        `json:"fogMaxclip4" db:"fog_maxclip4"`              //fog_maxclip4` float NOT NULL DEFAULT '450',
	FogDensity        float64        `json:"fogDensity" db:"fog_density"`                //fog_density` float NOT NULL DEFAULT '0',
	FlagNeeded        string         `json:"flagNeeded" db:"flag_needed"`                //flag_needed` varchar(128) NOT NULL DEFAULT '',
	Canbind           int64          `json:"canbind" db:"canbind"`                       //canbind` tinyint(4) NOT NULL DEFAULT '1',
	Cancombat         int64          `json:"cancombat" db:"cancombat"`                   //cancombat` tinyint(4) NOT NULL DEFAULT '1',
	Canlevitate       int64          `json:"canlevitate" db:"canlevitate"`               //canlevitate` tinyint(4) NOT NULL DEFAULT '1',
	Castoutdoor       int64          `json:"castoutdoor" db:"castoutdoor"`               //castoutdoor` tinyint(4) NOT NULL DEFAULT '1',
	Hotzone           int64          `json:"hotzone" db:"hotzone"`                       //hotzone` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Insttype          int64          `json:"insttype" db:"insttype"`                     //insttype` tinyint(1) unsigned zerofill NOT NULL DEFAULT '0',
	Shutdowndelay     int64          `json:"shutdowndelay" db:"shutdowndelay"`           //shutdowndelay` bigint(16) unsigned NOT NULL DEFAULT '5000',
	Peqzone           int64          `json:"peqzone" db:"peqzone"`                       //peqzone` tinyint(4) NOT NULL DEFAULT '1',
	Expansion         int64          `json:"expansion" db:"expansion"`                   //expansion` tinyint(3) NOT NULL DEFAULT '0',
	Suspendbuffs      int64          `json:"suspendbuffs" db:"suspendbuffs"`             //suspendbuffs` tinyint(1) unsigned NOT NULL DEFAULT '0',
	RainChance1       int64          `json:"rainChance1" db:"rain_chance1"`              //rain_chance1` int(4) NOT NULL DEFAULT '0',
	RainChance2       int64          `json:"rainChance2" db:"rain_chance2"`              //rain_chance2` int(4) NOT NULL DEFAULT '0',
	RainChance3       int64          `json:"rainChance3" db:"rain_chance3"`              //rain_chance3` int(4) NOT NULL DEFAULT '0',
	RainChance4       int64          `json:"rainChance4" db:"rain_chance4"`              //rain_chance4` int(4) NOT NULL DEFAULT '0',
	RainDuration1     int64          `json:"rainDuration1" db:"rain_duration1"`          //rain_duration1` int(4) NOT NULL DEFAULT '0',
	RainDuration2     int64          `json:"rainDuration2" db:"rain_duration2"`          //rain_duration2` int(4) NOT NULL DEFAULT '0',
	RainDuration3     int64          `json:"rainDuration3" db:"rain_duration3"`          //rain_duration3` int(4) NOT NULL DEFAULT '0',
	RainDuration4     int64          `json:"rainDuration4" db:"rain_duration4"`          //rain_duration4` int(4) NOT NULL DEFAULT '0',
	SnowChance1       int64          `json:"snowChance1" db:"snow_chance1"`              //snow_chance1` int(4) NOT NULL DEFAULT '0',
	SnowChance2       int64          `json:"snowChance2" db:"snow_chance2"`              //snow_chance2` int(4) NOT NULL DEFAULT '0',
	SnowChance3       int64          `json:"snowChance3" db:"snow_chance3"`              //snow_chance3` int(4) NOT NULL DEFAULT '0',
	SnowChance4       int64          `json:"snowChance4" db:"snow_chance4"`              //snow_chance4` int(4) NOT NULL DEFAULT '0',
	SnowDuration1     int64          `json:"snowDuration1" db:"snow_duration1"`          //snow_duration1` int(4) NOT NULL DEFAULT '0',
	SnowDuration2     int64          `json:"snowDuration2" db:"snow_duration2"`          //snow_duration2` int(4) NOT NULL DEFAULT '0',
	SnowDuration3     int64          `json:"snowDuration3" db:"snow_duration3"`          //snow_duration3` int(4) NOT NULL DEFAULT '0',
	SnowDuration4     int64          `json:"snowDuration4" db:"snow_duration4"`          //snow_duration4` int(4) NOT NULL DEFAULT '0',
	Gravity           float64        `json:"gravity" db:"gravity"`                       //gravity` float NOT NULL DEFAULT '0.4',
	Type              int64          `json:"type" db:"type"`                             //type` int(3) NOT NULL DEFAULT '0',
	Skylock           int64          `json:"skylock" db:"skylock"`                       //skylock` tinyint(4) NOT NULL DEFAULT '0',
}

//ZoneBy is used for sorting zones based on a context
type ZoneBy func(z1, z2 *Zone) bool

//Sort implements the sort interface
func (by ZoneBy) Sort(zones []Zone) {
	zs := &zoneSorter{
		zones: zones,
		by:    by,
	}
	sort.Sort(zs)
}

type zoneSorter struct {
	zones []Zone
	by    func(z1, z2 *Zone) bool
}

//Len implements the sort interface
func (s *zoneSorter) Len() int {
	return len(s.zones)
}

//Swap implements the sort interface
func (s *zoneSorter) Swap(i, j int) {
	s.zones[i], s.zones[j] = s.zones[j], s.zones[i]
}

//Less implements the sort interface
func (s *zoneSorter) Less(i, j int) bool {
	return s.by(&s.zones[i], &s.zones[j])
}

//GetMinStatusName converts the MinStatus field to the human readable name
func (c *Zone) GetMinStatusName() string {
	switch {
	case c.MinStatus >= 200:
		return "Admin"
	case c.MinStatus >= 100:
		return "Guide"
	}
	return ""
}

//ExpansionName returns the human readable form of an expansion id
func (c *Zone) ExpansionName() string {
	switch c.Expansion {
	case -1:
		return "Unknown"
	case 1:
		return "Classic"
	case 2:
		return "Ruins of Kunark"
	case 3:
		return "Scars of Velious"
	case 4:
		return "Shadows of Luclin"
	case 5:
		return "Planes of Power"
	case 6:
		return "Legacy of Ykesha"
	case 7:
		return "Lost Dungeons of Norrath"
	case 8:
		return "Gates of Discord"
	case 9:
		return "Omens of War"
	case 10:
		return "Dragons of Norrath"
	case 11:
		return "Depths of Darkhallow"
	case 12:
		return "Prophecy of Ro"
	case 13:
		return "Serpent's Spine"
	case 14:
		return "The Buried Sea"
	case 15:
		return "Secrets of Faydwer"
	case 16:
		return "Seeds of Destruction"
	case 17:
		return "Underfoot"
	case 18:
		return "House of Thule"
	case 19:
		return "Veil of Alaris"
	case 20:
		return "Rain of Fear"
	case 21:
		return "Call of the Forsaken"
	case 22:
		return "The Darkened Sea"
	case 23:
		return "The Broken Mirror"
	case 24:
		return "Empires of Kunark"
	case 25:
		return "Ring of Scale"
	}
	return "Unknown"
}
