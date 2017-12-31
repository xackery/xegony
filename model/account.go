package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/xeipuuv/gojsonschema"
)

//Account holds together characters inside Everquest's design
type Account struct {
	Id             int64         `json:"id" db:"id"`                         //`id` int(11) NOT NULL AUTO_INCREMENT,
	Name           string        `json:"name" db:"name"`                     //`name` varchar(30) NOT NULL DEFAULT '',
	Charname       string        `json:"charname" db:"charname"`             //`charname` varchar(64) NOT NULL DEFAULT '',
	Sharedplat     int64         `json:"sharedplat" db:"sharedplat"`         //`sharedplat` int(11) unsigned NOT NULL DEFAULT '0',
	Password       string        `json:"password" db:"password"`             //`password` varchar(50) NOT NULL DEFAULT '',
	Status         int64         `json:"status" db:"status"`                 //`status` int(5) NOT NULL DEFAULT '0',
	LsaccountID    sql.NullInt64 `json:"lsaccountID" db:"lsaccount_id"`      //`lsaccount_id` int(11) unsigned DEFAULT NULL,
	Gmspeed        int64         `json:"gmspeed" db:"gmspeed"`               //`gmspeed` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Revoked        int64         `json:"revoked" db:"revoked"`               //`revoked` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Karma          int64         `json:"karma" db:"karma"`                   //`karma` int(5) unsigned NOT NULL DEFAULT '0',
	MiniloginIp    string        `json:"miniloginIp" db:"minilogin_ip"`      //`minilogin_ip` varchar(32) NOT NULL DEFAULT '',
	Hideme         int64         `json:"hideme" db:"hideme"`                 //`hideme` tinyint(4) NOT NULL DEFAULT '0',
	Rulesflag      int64         `json:"rulesflag" db:"rulesflag"`           //`rulesflag` tinyint(1) unsigned NOT NULL DEFAULT '0',
	Suspendeduntil time.Time     `json:"suspendeduntil" db:"suspendeduntil"` //`suspendeduntil` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
	TimeCreation   int64         `json:"timeCreation" db:"time_creation"`    //`time_creation` int(10) unsigned NOT NULL DEFAULT '0',
	Expansion      int64         `json:"expansion" db:"expansion"`           //`expansion` tinyint(4) NOT NULL DEFAULT '0',
	BanReason      string        `json:"banReason" db:"ban_reason"`          //`ban_reason` text,
	SuspendReason  string        `json:"suspendReason" db:"suspend_reason"`  //`suspend_reason` text,
}

func (c *Account) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]Schema)
	var field string
	var prop Schema
	for _, field = range requiredFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	jsRef := gojsonschema.NewGoLoader(s)
	schema, err = gojsonschema.NewSchema(jsRef)
	if err != nil {
		return
	}
	return
}

func (c *Account) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "status":
		prop.Type = "integer"
		prop.Minimum = 1
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
