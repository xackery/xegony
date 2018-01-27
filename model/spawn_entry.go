package model

import (
	"database/sql"
)

// SpawnEntrys is an array of SpawnEntry
// swagger:model
type SpawnEntrys []*SpawnEntry

// SpawnEntry represents the spawn2 table
// swagger:model
type SpawnEntry struct {
	Zone *Zone `json:"zone"`
	//XScaled is used for scaling positions on the map
	XScaled float64 `json:"xScaled"`
	//YScaled is used for scaling positions on the map
	YScaled float64 `json:"yScaled"`
	//ZScaled is used for scaling positions on the map
	ZScaled float64 `json:"zScaled"`

	ID            int64          `json:"id" db:"id"`                   //`id` int(11) NOT NULL AUTO_INCREMENT,
	SpawnID       int64          `json:"spawnID" db:"spawngroupID"`    //`spawngroupID` int(11) NOT NULL DEFAULT '0',
	ZoneShortName sql.NullString `json:"zoneShortName" db:"zone"`      //`zone` varchar(32) DEFAULT NULL,
	Version       int64          `json:"version" db:"version"`         //`version` smallint(5) unsigned NOT NULL DEFAULT '0',
	X             float64        `json:"x" db:"x"`                     //`x` float(14,6) NOT NULL DEFAULT '0.000000',
	Y             float64        `json:"y" db:"y"`                     //`y` float(14,6) NOT NULL DEFAULT '0.000000',
	Z             float64        `json:"z" db:"z"`                     //`z` float(14,6) NOT NULL DEFAULT '0.000000',
	Heading       float64        `json:"heading" db:"heading"`         //`heading` float(14,6) NOT NULL DEFAULT '0.000000',
	RespawnTime   int64          `json:"respawntime" db:"respawntime"` //`respawntime` int(11) NOT NULL DEFAULT '0',
	Variance      int64          `json:"variance" db:"variance"`       //`variance` int(11) NOT NULL DEFAULT '0',
	Pathgrid      int64          `json:"pathgrid" db:"pathgrid"`       //`pathgrid` int(10) NOT NULL DEFAULT '0',
	Condition     int64          `json:"Condition" db:"_condition"`    //`_condition` mediumint(8) unsigned NOT NULL DEFAULT '0',
	CondValue     int64          `json:"condValue" db:"cond_value"`    //`cond_value` mediumint(9) NOT NULL DEFAULT '1',
	Enabled       int64          `json:"enabled" db:"enabled"`         //`enabled` tinyint(3) unsigned NOT NULL DEFAULT '1',
	Animation     int64          `json:"animation" db:"animation"`     //`animation` tinyint(3) unsigned NOT NULL DEFAULT '0',
}
