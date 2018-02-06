package model

// SpawnEvents is an array of SpawnEvent
// swagger:model
type SpawnEvents []*SpawnEvent

// SpawnEvent represents the spawn_events table
// swagger:model
type SpawnEvent struct {
	ID            int64  `json:"ID,omitempty" db:"id"`                  //`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	ZoneShortName string `json:"zoneShortName,omitempty" db:"zone"`     //`zone` varchar(32) DEFAULT NULL,
	ConditionID   int64  `json:"conditionID,omitempty" db:"cond_id"`    //`cond_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
	Name          string `json:"name,omitempty" db:"name"`              //`name` varchar(255) NOT NULL DEFAULT '',
	Period        int64  `json:"period,omitempty" db:"period"`          //`period` int(10) unsigned NOT NULL DEFAULT '0',
	NextMinute    int64  `json:"nextMinute,omitempty" db:"next_minute"` //`next_minute` tinyint(3) unsigned NOT NULL DEFAULT '0',
	NextHour      int64  `json:"nextHour,omitempty" db:"next_hour"`     //`next_hour` tinyint(3) unsigned NOT NULL DEFAULT '0',
	NextDay       int64  `json:"nextDay,omitempty" db:"next_day"`       //`next_day` tinyint(3) unsigned NOT NULL DEFAULT '0',
	NextMonth     int64  `json:"nextMonth,omitempty" db:"next_month"`   //`next_month` tinyint(3) unsigned NOT NULL DEFAULT '0',
	NextYear      int64  `json:"nextYear,omitempty" db:"next_year"`     //`next_year` int(10) unsigned NOT NULL DEFAULT '0',
	Enabled       int64  `json:"enabled,omitempty" db:"enabled"`        //`enabled` tinyint(4) NOT NULL DEFAULT '1',
	Action        int64  `json:"action,omitempty" db:"action"`          //`action` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Argument      int64  `json:"argument,omitempty" db:"argument"`      //`argument` mediumint(9) NOT NULL DEFAULT '0',
	Strict        int64  `json:"strict,omitempty" db:"strict"`          //`strict` tinyint(4) NOT NULL DEFAULT '0',
}
