package model

import ()

// SpawnRespawn represents the respawn_times table
// swagger:model
type SpawnRespawn struct {
	ID         int64 `json:"ID" db:"id"`                  //`id` int(11) NOT NULL DEFAULT '0',
	Start      int64 `json:"start" db:"start"`            //`start` int(11) NOT NULL DEFAULT '0',
	Duration   int64 `json:"duration" db:"duration"`      //`duration` int(11) NOT NULL DEFAULT '0',
	InstanceID int64 `json:"instanceID" db:"instance_id"` //`instance_id` smallint(6) NOT NULL DEFAULT '0',
}
