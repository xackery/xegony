package model

// Spawns is an array of Spawn
// swagger:model
type Spawns []*Spawn

// Spawn represents the spawn_group table
// swagger:model
type Spawn struct {
	Entrys []*SpawnEntry `json:"entries,omitempty"`
	Npcs   []*SpawnNpc   `json:"npcs,omitempty"`

	ID           int64   `json:"ID,omitempty" db:"id"`
	Name         string  `json:"name,omitempty" db:"name"`                  //`name` varchar(50) NOT NULL DEFAULT '',
	Limit        int64   `json:"limit,omitempty" db:"spawn_limit"`          //`spawn_limit` tinyint(4) NOT NULL DEFAULT '0',
	Distance     float64 `json:"distance,omitempty" db:"dist"`              //`dist` float NOT NULL DEFAULT '0',
	MaxX         float64 `json:"maxX,omitempty" db:"max_x"`                 //`max_x` float NOT NULL DEFAULT '0',
	MinX         float64 `json:"minX,omitempty" db:"min_x"`                 //`min_x` float NOT NULL DEFAULT '0',
	MaxY         float64 `json:"maxY,omitempty" db:"max_y"`                 //`max_y` float NOT NULL DEFAULT '0',
	MinY         float64 `json:"minY,omitempty" db:"min_y"`                 //`min_y` float NOT NULL DEFAULT '0',
	Delay        int64   `json:"delay,omitempty" db:"delay"`                //`delay` int(11) NOT NULL DEFAULT '45000',
	MinimumDelay int64   `json:"minimumDelay,omitempty" db:"mindelay"`      //`mindelay` int(11) NOT NULL DEFAULT '15000',
	Despawn      int64   `json:"despawn,omitempty" db:"despawn"`            //`despawn` tinyint(3) NOT NULL DEFAULT '0',
	DespawnTimer int64   `json:"despawnTimer,omitempty" db:"despawn_timer"` //`despawn_timer` int(11) NOT NULL DEFAULT '100',

}
