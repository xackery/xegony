package model

import ()

//Base is an everquest base data table representation
// swagger:model
type Base struct {
	Level   int64   `json:"level" db:"level"`      //`level` int(10) unsigned NOT NULL,
	Class   int64   `json:"class" db:"class"`      //`class` int(10) unsigned NOT NULL,
	Hp      float64 `json:"hp" db:"hp"`            //`hp` double NOT NULL,
	Mana    float64 `json:"mana" db:"mana"`        //`mana` double NOT NULL,
	End     float64 `json:"end" db:"end"`          //`end` double NOT NULL,
	Unk1    float64 `json:"unk1" db:"unk1"`        //`unk1` double NOT NULL,
	Unk2    float64 `json:"unk2" db:"unk2"`        //`unk2` double NOT NULL,
	HpFac   float64 `json:"hpFac" db:"hp_fac"`     //`hp_fac` double NOT NULL,
	ManaFac float64 `json:"manaFac" db:"mana_fac"` //`mana_fac` double NOT NULL,
	EndFac  float64 `json:"endFac" db:"end_fac"`   //`end_fac` double NOT NULL,
}
