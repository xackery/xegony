package model

import (
	"database/sql"
	"time"
)

// Accounts is an array of account
// swagger:model
type Accounts []*Account

// Account ties characters together, and represents login information to everquest
// swagger:model
type Account struct {

	//ID of user
	//required: false
	//example: 74887
	ID int64 `json:"ID,omitempty" db:"id"` //`id` int(11) NOT NULL AUTO_INCREMENT,
	//example: xackery
	//required: true
	Name string `json:"name,omitempty" db:"name"` //`name` varchar(30) NOT NULL DEFAULT '',
	//example: Shin
	Charname string `json:"charname,omitempty" db:"charname"` //`charname` varchar(64) NOT NULL DEFAULT '',
	//example: 1
	Sharedplat int64 `json:"sharedplat,omitempty" db:"sharedplat"` //`sharedplat` int(11) unsigned NOT NULL DEFAULT '0',
	//example: miniloginPassword
	Password string `json:"password,omitempty" db:"password"` //`password` varchar(50) NOT NULL DEFAULT '',
	//example: 250
	Status int64 `json:"status,omitempty" db:"status"` //`status` int(5) NOT NULL DEFAULT '0',
	//example: 0
	LsaccountID sql.NullInt64 `json:"lsaccountID,omitempty" db:"lsaccount_id"` //`lsaccount_id` int(11) unsigned DEFAULT NULL,
	//example: 0
	Gmspeed int64 `json:"gmspeed,omitempty" db:"gmspeed"` //`gmspeed` tinyint(3) unsigned NOT NULL DEFAULT '0',
	//example: 0
	Revoked int64 `json:"revoked,omitempty" db:"revoked"` //`revoked` tinyint(3) unsigned NOT NULL DEFAULT '0',
	//example: 0
	Karma int64 `json:"karma,omitempty" db:"karma"` //`karma` int(5) unsigned NOT NULL DEFAULT '0',
	//example: 127.0.0.1
	MiniloginIP string `json:"miniloginIp,omitempty" db:"minilogin_ip"` //`minilogin_ip` varchar(32) NOT NULL DEFAULT '',
	//example: 0
	Hideme int64 `json:"hideme,omitempty" db:"hideme"` //`hideme` tinyint(4) NOT NULL DEFAULT '0',
	//example: 0
	Rulesflag int64 `json:"rulesflag,omitempty" db:"rulesflag"` //`rulesflag` tinyint(1) unsigned NOT NULL DEFAULT '0',
	//example: 0001-01-01T00:00:00Z
	Suspendeduntil time.Time `json:"suspendeduntil,omitempty" db:"suspendeduntil"` //`suspendeduntil` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
	//example: 1451818675
	TimeCreation int64 `json:"timeCreation,omitempty" db:"time_creation"` //`time_creation` int(10) unsigned NOT NULL DEFAULT '0',
	//example: 0
	Expansion int64 `json:"expansion,omitempty" db:"expansion"` //`expansion` tinyint(4) NOT NULL DEFAULT '0',
	//exampe:
	BanReason sql.NullString `json:"banReason,omitempty" db:"ban_reason"` //`ban_reason` text,
	//example:
	SuspendReason sql.NullString `json:"suspendReason,omitempty" db:"suspend_reason"` //`suspend_reason` text,
}
