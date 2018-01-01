package model

import ()

//ZoneLevel is a cache table used to get levels a zone is available for
type ZoneLevel struct {
	ZoneID int64 `json:"zoneID" db:"zone_id"`
	Levels int64 `json:"levels" db:"levels"`
}
