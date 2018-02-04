package model

import ()

// ItemInventory represents inventory inside everquest
type ItemInventory struct {
	Slot     *Slot `json:"slot,omitempty"`
	AugSlot1 *Slot `json:"augSlot1,omitempty"`
	AugSlot2 *Slot `json:"augSlot2,omitempty"`
	AugSlot3 *Slot `json:"augSlot3,omitempty"`
	AugSlot4 *Slot `json:"augSlot4,omitempty"`
	AugSlot5 *Slot `json:"augSlot5,omitempty"`
	AugSlot6 *Slot `json:"augSlot6,omitempty"`

	CharacterID       int64  `json:"characterID" db:"charid"`                    //`charid` int(11) unsigned NOT NULL DEFAULT '0',
	SlotID            int64  `json:"slotID" db:"slotid"`                         //`slotid` mediumint(7) unsigned NOT NULL DEFAULT '0',
	ItemID            int64  `json:"itemID" db:"itemid"`                         //`itemid` int(11) unsigned DEFAULT '0',
	ChargeCount       int64  `json:"chargeCount" db:"charges"`                   //`charges` smallint(3) unsigned DEFAULT '0',
	Color             int64  `json:"color" db:"invcolor"`                        //`color` int(11) unsigned NOT NULL DEFAULT '0',
	AugSlot1ID        int64  `json:"augSlot1ID" db:"augslot1"`                   //`augslot1` mediumint(7) unsigned NOT NULL DEFAULT '0',
	AugSlot2ID        int64  `json:"augSlot2ID" db:"augslot2"`                   //`augslot2` mediumint(7) unsigned NOT NULL DEFAULT '0',
	AugSlot3ID        int64  `json:"augSlot3ID" db:"augslot3"`                   //`augslot3` mediumint(7) unsigned NOT NULL DEFAULT '0',
	AugSlot4ID        int64  `json:"augSlot4ID" db:"augslot4"`                   //`augslot4` mediumint(7) unsigned NOT NULL DEFAULT '0',
	AugSlot5ID        int64  `json:"augSlot5ID" db:"augslot5"`                   //`augslot5` mediumint(7) unsigned DEFAULT '0',
	AugSlot6ID        int64  `json:"augSlot6ID" db:"augslot6"`                   //`augslot6` mediumint(7) NOT NULL DEFAULT '0',
	InstNoDrop        int64  `json:"instNoDrop" db:"instNoDrop"`                 //`instnodrop` tinyint(1) unsigned NOT NULL DEFAULT '0',
	CustomData        string `json:"customData" db:"custom_data"`                //`custom_data` text,
	OrnamentIcon      int64  `json:"ornamentIcon" db:"ornamenticon"`             //`ornamenticon` int(11) unsigned NOT NULL DEFAULT '0',
	OrnamentIDFile    int64  `json:"ornamentIDFile" db:"ornamentidfile"`         //`ornamentidfile` int(11) unsigned NOT NULL DEFAULT '0',
	OrnamentHeroModel int64  `json:"ornamentHeroModel" db:"ornament_hero_model"` //`ornament_hero_model` int(11) NOT NULL DEFAULT '0',
}
