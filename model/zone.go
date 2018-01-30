package model

import (
	"database/sql"
)

// Zones is an array of zone
// swagger:model
type Zones []*Zone

// Zone represents the zone table, Everquest is split into zones
// swagger:model
type Zone struct {
	//Levels   int64   `json:"levels,omitempty"`
	//Modifier float64 `json:"modifier,omitempty"`

	ShortName         sql.NullString `json:"shortName,omitempty" db:"short_name"`                  //short_name` varchar(32) DEFAULT NULL,
	ID                int64          `json:"ID,omitempty" db:"id"`                                 //id` int(10) NOT NULL AUTO_INCREMENT,
	FileName          sql.NullString `json:"fileName,omitempty" db:"file_name"`                    //file_name` varchar(16) DEFAULT NULL,
	LongName          string         `json:"longName,omitempty" db:"long_name"`                    //long_name` text NOT NULL,
	MapFileName       sql.NullString `json:"mapFileName,omitempty" db:"map_file_name"`             //map_file_name` varchar(100) DEFAULT NULL,
	SafeX             float64        `json:"safeX,omitempty" db:"safe_x"`                          //safe_x` float NOT NULL DEFAULT '0',
	SafeY             float64        `json:"safeY,omitempty" db:"safe_y"`                          //safe_y` float NOT NULL DEFAULT '0',
	SafeZ             float64        `json:"safeZ,omitempty" db:"safe_z"`                          //safe_z` float NOT NULL DEFAULT '0',
	GraveyardID       float64        `json:"graveyardID,omitempty" db:"graveyard_id"`              //graveyard_id` float NOT NULL DEFAULT '0',
	MinLevel          int64          `json:"minLevel,omitempty" db:"min_level"`                    //min_level` tinyint(3) unsigned NOT NULL DEFAULT '0',
	MinStatus         int64          `json:"minStatus,omitempty" db:"min_status"`                  //min_status` tinyint(3) unsigned NOT NULL DEFAULT '0',
	ZoneIDNumber      int64          `json:"zoneIDNumber,omitempty" db:"zoneidnumber"`             //zoneidnumber` int(4) NOT NULL DEFAULT '0',
	Version           int64          `json:"version,omitempty" db:"version"`                       //version` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Timezone          int64          `json:"timezone,omitempty" db:"timezone"`                     //timezone` int(5) NOT NULL DEFAULT '0',
	MaxClients        int64          `json:"maxClients,omitempty" db:"maxclients"`                 //maxclients` int(5) NOT NULL DEFAULT '0',
	Ruleset           int64          `json:"ruleset,omitempty" db:"ruleset"`                       //ruleset` int(10) unsigned NOT NULL DEFAULT '0',
	Note              sql.NullString `json:"note,omitempty" db:"note"`                             //note` varchar(80) DEFAULT NULL,
	Underworld        float64        `json:"underworld,omitempty" db:"underworld"`                 //underworld` float NOT NULL DEFAULT '0',
	MinClip           float64        `json:"MinClip,omitempty" db:"minclip"`                       //minClip` float NOT NULL DEFAULT '450',
	MaxClip           float64        `json:"MaxClip,omitempty" db:"maxclip"`                       //maxClip` float NOT NULL DEFAULT '450',
	FogMinClip        float64        `json:"fogMinClip,omitempty" db:"fog_minclip"`                //fog_minClip` float NOT NULL DEFAULT '450',
	FogMaxClip        float64        `json:"fogMaxClip,omitempty" db:"fog_maxclip"`                //fog_maxClip` float NOT NULL DEFAULT '450',
	FogBlue           int64          `json:"fogBlue,omitempty" db:"fog_blue"`                      //fog_blue` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogRed            int64          `json:"fogRed,omitempty" db:"fog_red"`                        //fog_red` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen          int64          `json:"fogGreen,omitempty" db:"fog_green"`                    //fog_green` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Sky               int64          `json:"sky,omitempty" db:"sky"`                               //sky` tinyint(3) unsigned NOT NULL DEFAULT '1',
	ZType             int64          `json:"zType,omitempty" db:"ztype"`                           //ztype` tinyint(3) unsigned NOT NULL DEFAULT '1',
	ZoneExpMultiplier float64        `json:"zoneExpMultiplier,omitempty" db:"zone_exp_multiplier"` //zone_exp_multiplier` decimal(6,2) NOT NULL DEFAULT '0.00',
	WalkSpeed         float64        `json:"walkSpeed,omitempty" db:"walkspeed"`                   //walkspeed` float NOT NULL DEFAULT '0.4',
	TimeType          int64          `json:"timeType,omitempty" db:"time_type"`                    //time_type` tinyint(3) unsigned NOT NULL DEFAULT '2',
	FogRed1           int64          `json:"fogRed1,omitempty" db:"fog_red1"`                      //fog_red1` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen1         int64          `json:"fogGreen1,omitempty" db:"fog_green1"`                  //fog_green1` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogBlue1          int64          `json:"fogBlue1,omitempty" db:"fog_blue1"`                    //fog_blue1` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogMinClip1       float64        `json:"fogMinClip1,omitempty" db:"fog_minclip1"`              //fog_minClip1` float NOT NULL DEFAULT '450',
	FogMaxClip1       float64        `json:"fogMaxClip1,omitempty" db:"fog_maxclip1"`              //fog_maxClip1` float NOT NULL DEFAULT '450',
	FogRed2           int64          `json:"fogRed2,omitempty" db:"fog_red2"`                      //fog_red2` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen2         int64          `json:"fogGreen2,omitempty" db:"fog_green2"`                  //fog_green2` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogBlue2          int64          `json:"fogBlue2,omitempty" db:"fog_blue2"`                    //fog_blue2` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogMinClip2       float64        `json:"fogMinClip2,omitempty" db:"fog_minclip2"`              //fog_minClip2` float NOT NULL DEFAULT '450',
	FogMaxClip2       float64        `json:"fogMaxClip2,omitempty" db:"fog_maxclip2"`              //fog_maxClip2` float NOT NULL DEFAULT '450',
	FogRed3           int64          `json:"fogRed3,omitempty" db:"fog_red3"`                      //fog_red3` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen3         int64          `json:"fogGreen3,omitempty" db:"fog_green3"`                  //fog_green3` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogBlue3          int64          `json:"fogBlue3,omitempty" db:"fog_blue3"`                    //fog_blue3` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogMinClip3       float64        `json:"fogMinClip3,omitempty" db:"fog_minclip3"`              //fog_minClip3` float NOT NULL DEFAULT '450',
	FogMaxClip3       float64        `json:"fogMaxClip3,omitempty" db:"fog_maxclip3"`              //fog_maxClip3` float NOT NULL DEFAULT '450',
	FogRed4           int64          `json:"fogRed4,omitempty" db:"fog_red4"`                      //fog_red4` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogGreen4         int64          `json:"fogGreen4,omitempty" db:"fog_green4"`                  //fog_green4` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogBlue4          int64          `json:"fogBlue4,omitempty" db:"fog_blue4"`                    //fog_blue4` tinyint(3) unsigned NOT NULL DEFAULT '0',
	FogMinClip4       float64        `json:"fogMinClip4,omitempty" db:"fog_minclip4"`              //fog_minClip4` float NOT NULL DEFAULT '450',
	FogMaxClip4       float64        `json:"fogMaxClip4,omitempty" db:"fog_maxclip4"`              //fog_maxClip4` float NOT NULL DEFAULT '450',
	FogDensity        float64        `json:"fogDensity,omitempty" db:"fog_density"`                //fog_density` float NOT NULL DEFAULT '0',
	FlagNeeded        string         `json:"flagNeeded,omitempty" db:"flag_needed"`                //flag_needed` varchar(128) NOT NULL DEFAULT '',
	CanBind           int64          `json:"canBind,omitempty" db:"canbind"`                       //canbind` tinyint(4) NOT NULL DEFAULT '1',
	CanCombat         int64          `json:"canCombat,omitempty" db:"cancombat"`                   //cancombat` tinyint(4) NOT NULL DEFAULT '1',
	CanLevitate       int64          `json:"canLevitate,omitempty" db:"canlevitate"`               //canlevitate` tinyint(4) NOT NULL DEFAULT '1',
	CastOutdoor       int64          `json:"castOutdoor,omitempty" db:"castoutdoor"`               //castoutdoor` tinyint(4) NOT NULL DEFAULT '1',
	HotZone           int64          `json:"hotZone,omitempty" db:"hotzone"`                       //hotzone` tinyint(3) unsigned NOT NULL DEFAULT '0',
	InstType          int64          `json:"instType,omitempty" db:"insttype"`                     //insttype` tinyint(1) unsigned zerofill NOT NULL DEFAULT '0',
	ShutdownDelay     int64          `json:"shutdownDelay,omitempty" db:"shutdowndelay"`           //shutdowndelay` bigint(16) unsigned NOT NULL DEFAULT '5000',
	PeqZone           int64          `json:"peqZone,omitempty" db:"peqzone"`                       //peqzone` tinyint(4) NOT NULL DEFAULT '1',
	Expansion         int64          `json:"expansion,omitempty" db:"expansion"`                   //expansion` tinyint(3) NOT NULL DEFAULT '0',
	SuspendBuffs      int64          `json:"suspendBuffs,omitempty" db:"suspendbuffs"`             //suspendbuffs` tinyint(1) unsigned NOT NULL DEFAULT '0',
	RainChance1       int64          `json:"rainChance1,omitempty" db:"rain_chance1"`              //rain_chance1` int(4) NOT NULL DEFAULT '0',
	RainChance2       int64          `json:"rainChance2,omitempty" db:"rain_chance2"`              //rain_chance2` int(4) NOT NULL DEFAULT '0',
	RainChance3       int64          `json:"rainChance3,omitempty" db:"rain_chance3"`              //rain_chance3` int(4) NOT NULL DEFAULT '0',
	RainChance4       int64          `json:"rainChance4,omitempty" db:"rain_chance4"`              //rain_chance4` int(4) NOT NULL DEFAULT '0',
	RainDuration1     int64          `json:"rainDuration1,omitempty" db:"rain_duration1"`          //rain_duration1` int(4) NOT NULL DEFAULT '0',
	RainDuration2     int64          `json:"rainDuration2,omitempty" db:"rain_duration2"`          //rain_duration2` int(4) NOT NULL DEFAULT '0',
	RainDuration3     int64          `json:"rainDuration3,omitempty" db:"rain_duration3"`          //rain_duration3` int(4) NOT NULL DEFAULT '0',
	RainDuration4     int64          `json:"rainDuration4,omitempty" db:"rain_duration4"`          //rain_duration4` int(4) NOT NULL DEFAULT '0',
	SnowChance1       int64          `json:"snowChance1,omitempty" db:"snow_chance1"`              //snow_chance1` int(4) NOT NULL DEFAULT '0',
	SnowChance2       int64          `json:"snowChance2,omitempty" db:"snow_chance2"`              //snow_chance2` int(4) NOT NULL DEFAULT '0',
	SnowChance3       int64          `json:"snowChance3,omitempty" db:"snow_chance3"`              //snow_chance3` int(4) NOT NULL DEFAULT '0',
	SnowChance4       int64          `json:"snowChance4,omitempty" db:"snow_chance4"`              //snow_chance4` int(4) NOT NULL DEFAULT '0',
	SnowDuration1     int64          `json:"snowDuration1,omitempty" db:"snow_duration1"`          //snow_duration1` int(4) NOT NULL DEFAULT '0',
	SnowDuration2     int64          `json:"snowDuration2,omitempty" db:"snow_duration2"`          //snow_duration2` int(4) NOT NULL DEFAULT '0',
	SnowDuration3     int64          `json:"snowDuration3,omitempty" db:"snow_duration3"`          //snow_duration3` int(4) NOT NULL DEFAULT '0',
	SnowDuration4     int64          `json:"snowDuration4,omitempty" db:"snow_duration4"`          //snow_duration4` int(4) NOT NULL DEFAULT '0',
	Gravity           float64        `json:"gravity,omitempty" db:"gravity"`                       //gravity` float NOT NULL DEFAULT '0.4',
	Type              int64          `json:"type,omitempty" db:"type"`                             //type` int(3) NOT NULL DEFAULT '0',
	Skylock           int64          `json:"skylock,omitempty" db:"skylock"`                       //skylock` tinyint(4) NOT NULL DEFAULT '0',
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
