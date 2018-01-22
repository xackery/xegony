package model

import (
	"database/sql"
)

//SharedBank represents the sharedbank table, used for npcs
// swagger:response
type SharedBank struct {
	AccountID  int64          `json:"acctid" db:"acctid"`         //`acctid` int(11) unsigned DEFAULT '0',
	SlotID     int64          `json:"slotid" db:"slotid"`         //`slotid` mediumint(7) unsigned DEFAULT '0',
	ItemID     int64          `json:"itemid" db:"itemid"`         //`itemid` int(11) unsigned DEFAULT '0',
	Charges    int64          `json:"charges" db:"charges"`       //`charges` smallint(3) unsigned DEFAULT '0',
	Augslot1   int64          `json:"augslot1" db:"augslot1"`     //`augslot1` mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot2   int64          `json:"augslot2" db:"augslot2"`     //`augslot2` mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot3   int64          `json:"augslot3" db:"augslot3"`     //`augslot3` mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot4   int64          `json:"augslot4" db:"augslot4"`     //`augslot4` mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot5   int64          `json:"augslot5" db:"augslot5"`     //`augslot5` mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot6   int64          `json:"augslot6" db:"augslot6"`     //`augslot6` mediumint(7) NOT NULL DEFAULT '0',
	CustomData sql.NullString `json:"customData" db:"customData"` //`custom_data` text,
}
