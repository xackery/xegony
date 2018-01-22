package model

import (
	"time"
)

//CharacterGraph holds data about players performance, it primarily uses character_graph table
// swagger:model
type CharacterGraph struct {
	ID           int64     `json:"id" db:"id"`                      //`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	CharacterID  int64     `json:"characterId" db:"character_id"`   //`character_id` int(11) DEFAULT NULL,
	CreateDate   time.Time `json:"createDate" db:"create_date"`     //`create_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	Experience   int64     `json:"experience" db:"experience"`      //`experience` int(10) unsigned NOT NULL DEFAULT '0',
	AAExperience int64     `json:"aaExperience" db:"aa_experience"` //`aa_experience` int(10) unsigned NOT NULL DEFAULT '0',
}
