package model

import ()

//Mail represents mail inside everquest
// swagger:model
type Mail struct {
	ID          int64  `json:"ID" db:"msgid"`            //`msgid` int(10) unsigned NOT NULL AUTO_INCREMENT,
	CharacterID int64  `json:"characterID" db:"charid"`  //`charid` int(10) unsigned NOT NULL DEFAULT '0',
	Timestamp   int64  `json:"timestamp" db:"timestamp"` //`timestamp` int(11) NOT NULL DEFAULT '0',
	From        string `json:"from" db:"from"`           //`from` varchar(100) NOT NULL DEFAULT '',
	Subject     string `json:"subject" db:"subject"`     //`subject` varchar(200) NOT NULL DEFAULT '',
	Body        string `json:"body" db:"body"`           //`body` text NOT NULL,
	To          string `json:"to" db:"to"`               //`to` text NOT NULL,
	Status      int64  `json:"status" db:"status"`       //`status` tinyint(4) NOT NULL DEFAULT '0',
}
