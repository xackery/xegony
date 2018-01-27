package model

import ()

// Spawns is an array of Spawn
// swagger:model
type Spawns []*Spawn

// Spawn represents the spawn_group table
// swagger:model
type Spawn struct {
	ID           int64   `json:"id" db:"id"`
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
