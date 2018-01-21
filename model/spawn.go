package model

import (
	"database/sql"
)

//Spawn represents the spawn2 table, used for npcs
type Spawn struct {
	XScaled float64 `json:"xScaled"`
	YScaled float64 `json:"yScaled"`
	ZScaled float64 `json:"zScaled"`

	ID           int64          `json:"id" db:"spawn2id"`               //`id` int(11) NOT NULL AUTO_INCREMENT,
	SpawngroupID int64          `json:"spawngroupID" db:"spawngroupID"` //`spawngroupID` int(11) NOT NULL DEFAULT '0',
	Zone         sql.NullString `json:"zone" db:"zone"`                 //`zone` varchar(32) DEFAULT NULL,
	Version      int64          `json:"version" db:"version"`           //`version` smallint(5) unsigned NOT NULL DEFAULT '0',
	X            float64        `json:"x" db:"x"`                       //`x` float(14,6) NOT NULL DEFAULT '0.000000',
	Y            float64        `json:"y" db:"y"`                       //`y` float(14,6) NOT NULL DEFAULT '0.000000',
	Z            float64        `json:"z" db:"z"`                       //`z` float(14,6) NOT NULL DEFAULT '0.000000',
	Heading      float64        `json:"heading" db:"heading"`           //`heading` float(14,6) NOT NULL DEFAULT '0.000000',
	RespawnTime  int64          `json:"respawntime" db:"respawntime"`   //`respawntime` int(11) NOT NULL DEFAULT '0',
	Variance     int64          `json:"variance" db:"variance"`         //`variance` int(11) NOT NULL DEFAULT '0',
	Pathgrid     int64          `json:"pathgrid" db:"pathgrid"`         //`pathgrid` int(10) NOT NULL DEFAULT '0',
	Condition    int64          `json:"Condition" db:"_condition"`      //`_condition` mediumint(8) unsigned NOT NULL DEFAULT '0',
	CondValue    int64          `json:"condValue" db:"cond_value"`      //`cond_value` mediumint(9) NOT NULL DEFAULT '1',
	Enabled      int64          `json:"enabled" db:"enabled"`           //`enabled` tinyint(3) unsigned NOT NULL DEFAULT '1',
	Animation    int64          `json:"animation" db:"animation"`       //`animation` tinyint(3) unsigned NOT NULL DEFAULT '0',

	Name         string  `json:"name" db:"name"`                  //`name` varchar(50) NOT NULL DEFAULT '',
	SpawnLimit   int64   `json:"spawnLimit" db:"spawn_limit"`     //`spawn_limit` tinyint(4) NOT NULL DEFAULT '0',
	Dist         float64 `json:"dist" db:"dist"`                  //`dist` float NOT NULL DEFAULT '0',
	MaxX         float64 `json:"maxX" db:"max_x"`                 //`max_x` float NOT NULL DEFAULT '0',
	MinX         float64 `json:"minX" db:"min_x"`                 //`min_x` float NOT NULL DEFAULT '0',
	MaxY         float64 `json:"maxY" db:"max_y"`                 //`max_y` float NOT NULL DEFAULT '0',
	MinY         float64 `json:"minY" db:"min_y"`                 //`min_y` float NOT NULL DEFAULT '0',
	Delay        int64   `json:"delay" db:"delay"`                //`delay` int(11) NOT NULL DEFAULT '45000',
	Mindelay     int64   `json:"mindelay" db:"mindelay"`          //`mindelay` int(11) NOT NULL DEFAULT '15000',
	Despawn      int64   `json:"despawn" db:"despawn"`            //`despawn` tinyint(3) NOT NULL DEFAULT '0',
	DespawnTimer int64   `json:"despawnTimer" db:"despawn_timer"` //`despawn_timer` int(11) NOT NULL DEFAULT '100',
}
