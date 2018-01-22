package model

import ()

// AaRank represents the root objects of Alternate Abilities
// swagger:response
type AaRank struct {

	//aa_ranks
	ID             int64 `json:"ID" db:"id"`                           //`id` int(10) unsigned NOT NULL,
	UpperHotkeySid int64 `json:"upperHotkeySid" db:"upper_hotkey_sid"` //`upper_hotkey_sid` int(10) NOT NULL DEFAULT '-1',
	LowerHotkeySid int64 `json:"lowerHotkeySid" db:"lower_hotkey_sid"` //`lower_hotkey_sid` int(10) NOT NULL DEFAULT '-1',
	TitleSid       int64 `json:"titleSid" db:"title_sid"`              //`title_sid` int(10) NOT NULL DEFAULT '-1',
	DescSid        int64 `json:"descSid" db:"desc_sid"`                //`desc_sid` int(10) NOT NULL DEFAULT '-1',
	Cost           int64 `json:"cost" db:"cost"`                       //`cost` int(10) NOT NULL DEFAULT '1',
	LevelReq       int64 `json:"levelReq" db:"level_req"`              //`level_req` int(10) NOT NULL DEFAULT '51',
	Spell          int64 `json:"spell" db:"spell"`                     //`spell` int(10) NOT NULL DEFAULT '-1',
	SpellType      int64 `json:"spellType" db:"spell_type"`            //`spell_type` int(10) NOT NULL DEFAULT '0',
	RecastTime     int64 `json:"recastTime" db:"recast_time"`          //`recast_time` int(10) NOT NULL DEFAULT '0',
	Expansion      int64 `json:"expansion" db:"expansion"`             //`expansion` int(10) NOT NULL DEFAULT '0',
	PrevID         int64 `json:"prevID" db:"prev_id"`                  //`prev_id` int(10) NOT NULL DEFAULT '-1',
	NextID         int64 `json:"nextID" db:"next_id"`                  //`next_id` int(10) NOT NULL DEFAULT '-1',
}
