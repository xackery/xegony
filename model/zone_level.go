package model

import ()

//ZoneLevel is a cache table used to get levels a zone is available for
// swagger:response
type ZoneLevel struct {
	ZoneID     int64   `json:"zoneID" db:"zone_id"`
	Levels     int64   `json:"levels" db:"levels"`
	MapAspect  float64 `json:"mapAspect" db:"map_aspect"`
	MapXOffset float64 `json:"mapXOffset" db:"map_x_offset"`
	MapYOffset float64 `json:"mapYOffset" db:"map_y_offset"`
}
