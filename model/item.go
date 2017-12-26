package model

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"

	"github.com/xeipuuv/gojsonschema"
)

type Item struct {
	//inventory added details
	SlotId              int64          `json:"slotid" db:"slotid"`
	Charges             sql.NullInt64  `json:"charges" db:"charges"`
	InvColor            int64          `json:"invColor" db:"invcolor"`
	Augslot1            int64          `json:"augslot1" db:"augslot1"`
	Augslot2            int64          `json:"augslot2"`            // mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot3            int64          `json:"augslot3"`            // mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot4            int64          `json:"augslot4"`            // mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot5            int64          `json:"augslot5"`            // mediumint(7) unsigned DEFAULT '0',
	Augslot6            int64          `json:"augslot6"`            // mediumint(7) NOT NULL DEFAULT '0',
	Instnodrop          int64          `json:"instnodrop"`          // tinyint(1) unsigned NOT NULL DEFAULT '0',
	Custom_data         sql.NullString `json:"custom_data"`         // text,
	Ornamenticon        int64          `json:"ornamenticon"`        // int(11) unsigned NOT NULL DEFAULT '0',
	Ornamentidfile      int64          `json:"ornamentidfile"`      // int(11) unsigned NOT NULL DEFAULT '0',
	Ornament_hero_model int64          `json:"ornament_hero_model"` // int(11) NOT NULL DEFAULT '0',

	//normal items db
	Id                  int64          `json:"id" db:"id"`
	Name                string         `json:"name" db:"name"`
	Aagi                int64          `json:"aagi" db:"aagi"`
	Ac                  int64          `json:"ac" db:"ac"`
	Accuracy            int64          `json:"accuracy" db:"accuracy"`
	Acha                int64          `json:"acha" db:"acha"`
	Adex                int64          `json:"adex" db:"adex"`
	Aint                int64          `json:"aint" db:"aint"`
	Artifactflag        int64          `json:"artifactflag" db:"artifactflag"`
	Asta                int64          `json:"asta" db:"asta"`
	Astr                int64          `json:"astr" db:"astr"`
	Attack              int64          `json:"attack" db:"attack"`
	Augrestrict         int64          `json:"augrestrict" db:"augrestrict"`
	Augslot1type        int64          `json:"augslot1type" db:"augslot1type"`
	Augslot1visible     sql.NullInt64  `json:"augslot1visible" db:"augslot1visible"`
	Augslot2type        int64          `json:"augslot2type" db:"augslot2type"`
	Augslot2visible     sql.NullInt64  `json:"augslot2visible" db:"augslot2visible"`
	Augslot3type        int64          `json:"augslot3type" db:"augslot3type"`
	Augslot3visible     sql.NullInt64  `json:"augslot3visible" db:"augslot3visible"`
	Augslot4type        int64          `json:"augslot4type" db:"augslot4type"`
	Augslot4visible     sql.NullInt64  `json:"augslot4visible" db:"augslot4visible"`
	Augslot5type        int64          `json:"augslot5type" db:"augslot5type"`
	Augslot5visible     sql.NullInt64  `json:"augslot5visible" db:"augslot5visible"`
	Augslot6type        int64          `json:"augslot6type" db:"augslot6type"`
	Augslot6visible     int64          `json:"augslot6visible" db:"augslot6visible"`
	Augtype             int64          `json:"augtype" db:"augtype"`
	Avoidance           int64          `json:"avoidance" db:"avoidance"`
	Awis                int64          `json:"awis" db:"awis"`
	Bagsize             int64          `json:"bagsize" db:"bagsize"`
	Bagslots            int64          `json:"bagslots" db:"bagslots"`
	Bagtype             int64          `json:"bagtype" db:"bagtype"`
	Bagwr               int64          `json:"bagwr" db:"bagwr"`
	Banedmgamt          int64          `json:"banedmgamt" db:"banedmgamt"`
	Banedmgraceamt      int64          `json:"banedmgraceamt" db:"banedmgraceamt"`
	Banedmgbody         int64          `json:"banedmgbody" db:"banedmgbody"`
	Banedmgrace         int64          `json:"banedmgrace" db:"banedmgrace"`
	Bardtype            int64          `json:"bardtype" db:"bardtype"`
	Bardvalue           int64          `json:"bardvalue" db:"bardvalue"`
	Book                int64          `json:"book" db:"book"`
	Casttime            int64          `json:"casttime" db:"casttime"`
	Casttime_           int64          `json:"casttime_" db:"casttime_"`
	Charmfile           string         `json:"charmfile" db:"charmfile"`
	Charmfileid         string         `json:"charmfileid" db:"charmfileid"`
	Classes             int64          `json:"classes" db:"classes"`
	Color               int64          `json:"color" db:"color"`
	Combateffects       string         `json:"combateffects" db:"combateffects"`
	Extradmgskill       int64          `json:"extradmgskill" db:"extradmgskill"`
	Extradmgamt         int64          `json:"extradmgamt" db:"extradmgamt"`
	Price               int64          `json:"price" db:"price"`
	Cr                  int64          `json:"cr" db:"cr"`
	Damage              int64          `json:"damage" db:"damage"`
	Damageshield        int64          `json:"damageshield" db:"damageshield"`
	Deity               int64          `json:"deity" db:"deity"`
	Delay               int64          `json:"delay" db:"delay"`
	Augdistiller        int64          `json:"augdistiller" db:"augdistiller"`
	Dotshielding        int64          `json:"dotshielding" db:"dotshielding"`
	Dr                  int64          `json:"dr" db:"dr"`
	Clicktype           int64          `json:"clicktype" db:"clicktype"`
	Clicklevel2         int64          `json:"clicklevel2" db:"clicklevel2"`
	Elemdmgtype         int64          `json:"elemdmgtype" db:"elemdmgtype"`
	Elemdmgamt          int64          `json:"elemdmgamt" db:"elemdmgamt"`
	Endur               int64          `json:"endur" db:"endur"`
	Factionamt1         int64          `json:"factionamt1" db:"factionamt1"`
	Factionamt2         int64          `json:"factionamt2" db:"factionamt2"`
	Factionamt3         int64          `json:"factionamt3" db:"factionamt3"`
	Factionamt4         int64          `json:"factionamt4" db:"factionamt4"`
	Factionmod1         int64          `json:"factionmod1" db:"factionmod1"`
	Factionmod2         int64          `json:"factionmod2" db:"factionmod2"`
	Factionmod3         int64          `json:"factionmod3" db:"factionmod3"`
	Factionmod4         int64          `json:"factionmod4" db:"factionmod4"`
	Filename            string         `json:"filename" db:"filename"`
	Focuseffect         int64          `json:"focuseffect" db:"focuseffect"`
	Fr                  int64          `json:"fr" db:"fr"`
	Fvnodrop            int64          `json:"fvnodrop" db:"fvnodrop"`
	Haste               int64          `json:"haste" db:"haste"`
	Clicklevel          int64          `json:"clicklevel" db:"clicklevel"`
	Hp                  int64          `json:"hp" db:"hp"`
	Regen               int64          `json:"regen" db:"regen"`
	Icon                int64          `json:"icon" db:"icon"`
	Idfile              string         `json:"idfile" db:"idfile"`
	Itemclass           int64          `json:"itemclass" db:"itemclass"`
	Itemtype            int64          `json:"itemtype" db:"itemtype"`
	Ldonprice           int64          `json:"ldonprice" db:"ldonprice"`
	Ldontheme           int64          `json:"ldontheme" db:"ldontheme"`
	Ldonsold            int64          `json:"ldonsold" db:"ldonsold"`
	Light               int64          `json:"light" db:"light"`
	Lore                string         `json:"lore" db:"lore"`
	Loregroup           int64          `json:"loregroup" db:"loregroup"`
	Magic               int64          `json:"magic" db:"magic"`
	Mana                int64          `json:"mana" db:"mana"`
	Manaregen           int64          `json:"manaregen" db:"manaregen"`
	Enduranceregen      int64          `json:"enduranceregen" db:"enduranceregen"`
	Material            int64          `json:"material" db:"material"`
	Herosforgemodel     int64          `json:"herosforgemodel" db:"herosforgemodel"`
	Maxcharges          int64          `json:"maxcharges" db:"maxcharges"`
	Mr                  int64          `json:"mr" db:"mr"`
	Nodrop              int64          `json:"nodrop" db:"nodrop"`
	Norent              int64          `json:"norent" db:"norent"`
	Pendingloreflag     int64          `json:"pendingloreflag" db:"pendingloreflag"`
	Pr                  int64          `json:"pr" db:"pr"`
	Procrate            int64          `json:"procrate" db:"procrate"`
	Races               int64          `json:"races" db:"races"`
	Range               int64          `json:"range" db:"range"`
	Reclevel            int64          `json:"reclevel" db:"reclevel"`
	Recskill            int64          `json:"recskill" db:"recskill"`
	Reqlevel            int64          `json:"reqlevel" db:"reqlevel"`
	Sellrate            float64        `json:"sellrate" db:"sellrate"`
	Shielding           int64          `json:"shielding" db:"shielding"`
	Size                int64          `json:"size" db:"size"`
	Skillmodtype        int64          `json:"skillmodtype" db:"skillmodtype"`
	Skillmodvalue       int64          `json:"skillmodvalue" db:"skillmodvalue"`
	Slots               int64          `json:"slots" db:"slots"`
	Clickeffect         int64          `json:"clickeffect" db:"clickeffect"`
	Spellshield         int64          `json:"spellshield" db:"spellshield"`
	Strikethrough       int64          `json:"strikethrough" db:"strikethrough"`
	Stunresist          int64          `json:"stunresist" db:"stunresist"`
	Summonedflag        int64          `json:"summonedflag" db:"summonedflag"`
	Tradeskills         int64          `json:"tradeskills" db:"tradeskills"`
	Favor               int64          `json:"favor" db:"favor"`
	Weight              int64          `json:"weight" db:"weight"`
	Unk012              int64          `json:"UNK012" db:"UNK012"`
	Unk013              int64          `json:"UNK013" db:"UNK013"`
	Benefitflag         int64          `json:"benefitflag" db:"benefitflag"`
	Unk054              int64          `json:"UNK054" db:"UNK054"`
	Unk059              int64          `json:"UNK059" db:"UNK059"`
	Booktype            int64          `json:"booktype" db:"booktype"`
	Recastdelay         int64          `json:"recastdelay" db:"recastdelay"`
	Recasttype          int64          `json:"recasttype" db:"recasttype"`
	Guildfavor          int64          `json:"guildfavor" db:"guildfavor"`
	Unk123              int64          `json:"UNK123" db:"UNK123"`
	Unk124              int64          `json:"UNK124" db:"UNK124"`
	Attuneable          int64          `json:"attuneable" db:"attuneable"`
	Nopet               int64          `json:"nopet" db:"nopet"`
	Updated             mysql.NullTime `json:"updated" db:"updated"`
	Comment             string         `json:"comment" db:"comment"`
	Unk127              int64          `json:"UNK127" db:"UNK127"`
	Pointtype           int64          `json:"pointtype" db:"pointtype"`
	Potionbelt          int64          `json:"potionbelt" db:"potionbelt"`
	Potionbeltslots     int64          `json:"potionbeltslots" db:"potionbeltslots"`
	Stacksize           int64          `json:"stacksize" db:"stacksize"`
	Notransfer          int64          `json:"notransfer" db:"notransfer"`
	Stackable           int64          `json:"stackable" db:"stackable"`
	Unk134              string         `json:"UNK134" db:"UNK134"`
	Unk137              int64          `json:"UNK137" db:"UNK137"`
	Proceffect          int64          `json:"proceffect" db:"proceffect"`
	Proctype            int64          `json:"proctype" db:"proctype"`
	Proclevel2          int64          `json:"proclevel2" db:"proclevel2"`
	Proclevel           int64          `json:"proclevel" db:"proclevel"`
	Unk142              int64          `json:"UNK142" db:"UNK142"`
	Worneffect          int64          `json:"worneffect" db:"worneffect"`
	Worntype            int64          `json:"worntype" db:"worntype"`
	Wornlevel2          int64          `json:"wornlevel2" db:"wornlevel2"`
	Wornlevel           int64          `json:"wornlevel" db:"wornlevel"`
	Unk147              int64          `json:"UNK147" db:"UNK147"`
	Focustype           int64          `json:"focustype" db:"focustype"`
	Focuslevel2         int64          `json:"focuslevel2" db:"focuslevel2"`
	Focuslevel          int64          `json:"focuslevel" db:"focuslevel"`
	Unk152              int64          `json:"UNK152" db:"UNK152"`
	Scrolleffect        int64          `json:"scrolleffect" db:"scrolleffect"`
	Scrolltype          int64          `json:"scrolltype" db:"scrolltype"`
	Scrolllevel2        int64          `json:"scrolllevel2" db:"scrolllevel2"`
	Scrolllevel         int64          `json:"scrolllevel" db:"scrolllevel"`
	Unk157              int64          `json:"UNK157" db:"UNK157"`
	Serialized          mysql.NullTime `json:"serialized" db:"serialized"`
	Verified            mysql.NullTime `json:"verified" db:"verified"`
	Serialization       sql.NullString `json:"serialization" db:"serialization"`
	Source              string         `json:"source" db:"source"`
	Unk033              int64          `json:"UNK033" db:"UNK033"`
	Lorefile            string         `json:"lorefile" db:"lorefile"`
	Unk014              int64          `json:"UNK014" db:"UNK014"`
	Svcorruption        int64          `json:"svcorruption" db:"svcorruption"`
	Skillmodmax         int64          `json:"skillmodmax" db:"skillmodmax"`
	Unk060              int64          `json:"UNK060" db:"UNK060"`
	Augslot1unk2        int64          `json:"augslot1unk2" db:"augslot1unk2"`
	Augslot2unk2        int64          `json:"augslot2unk2" db:"augslot2unk2"`
	Augslot3unk2        int64          `json:"augslot3unk2" db:"augslot3unk2"`
	Augslot4unk2        int64          `json:"augslot4unk2" db:"augslot4unk2"`
	Augslot5unk2        int64          `json:"augslot5unk2" db:"augslot5unk2"`
	Augslot6unk2        int64          `json:"augslot6unk2" db:"augslot6unk2"`
	Unk120              int64          `json:"UNK120" db:"UNK120"`
	Unk121              int64          `json:"UNK121" db:"UNK121"`
	Questitemflag       int64          `json:"questitemflag" db:"questitemflag"`
	Unk132              string         `json:"UNK132" db:"UNK132"`
	Clickunk5           int64          `json:"clickunk5" db:"clickunk5"`
	Clickunk6           string         `json:"clickunk6" db:"clickunk6"`
	Clickunk7           int64          `json:"clickunk7" db:"clickunk7"`
	Procunk1            int64          `json:"procunk1" db:"procunk1"`
	Procunk2            int64          `json:"procunk2" db:"procunk2"`
	Procunk3            int64          `json:"procunk3" db:"procunk3"`
	Procunk4            int64          `json:"procunk4" db:"procunk4"`
	Procunk6            string         `json:"procunk6" db:"procunk6"`
	Procunk7            int64          `json:"procunk7" db:"procunk7"`
	Wornunk1            int64          `json:"wornunk1" db:"wornunk1"`
	Wornunk2            int64          `json:"wornunk2" db:"wornunk2"`
	Wornunk3            int64          `json:"wornunk3" db:"wornunk3"`
	Wornunk4            int64          `json:"wornunk4" db:"wornunk4"`
	Wornunk5            int64          `json:"wornunk5" db:"wornunk5"`
	Wornunk6            string         `json:"wornunk6" db:"wornunk6"`
	Wornunk7            int64          `json:"wornunk7" db:"wornunk7"`
	Focusunk1           int64          `json:"focusunk1" db:"focusunk1"`
	Focusunk2           int64          `json:"focusunk2" db:"focusunk2"`
	Focusunk3           int64          `json:"focusunk3" db:"focusunk3"`
	Focusunk4           int64          `json:"focusunk4" db:"focusunk4"`
	Focusunk5           int64          `json:"focusunk5" db:"focusunk5"`
	Focusunk6           string         `json:"focusunk6" db:"focusunk6"`
	Focusunk7           int64          `json:"focusunk7" db:"focusunk7"`
	Scrollunk1          int64          `json:"scrollunk1" db:"scrollunk1"`
	Scrollunk2          int64          `json:"scrollunk2" db:"scrollunk2"`
	Scrollunk3          int64          `json:"scrollunk3" db:"scrollunk3"`
	Scrollunk4          int64          `json:"scrollunk4" db:"scrollunk4"`
	Scrollunk5          int64          `json:"scrollunk5" db:"scrollunk5"`
	Scrollunk6          string         `json:"scrollunk6" db:"scrollunk6"`
	Scrollunk7          int64          `json:"scrollunk7" db:"scrollunk7"`
	Unk193              int64          `json:"UNK193" db:"UNK193"`
	Purity              int64          `json:"purity" db:"purity"`
	Evoitem             int64          `json:"evoitem" db:"evoitem"`
	Evoid               int64          `json:"evoid" db:"evoid"`
	Evolvinglevel       int64          `json:"evolvinglevel" db:"evolvinglevel"`
	Evomax              int64          `json:"evomax" db:"evomax"`
	Clickname           string         `json:"clickname" db:"clickname"`
	Procname            string         `json:"procname" db:"procname"`
	Wornname            string         `json:"wornname" db:"wornname"`
	Focusname           string         `json:"focusname" db:"focusname"`
	Scrollname          string         `json:"scrollname" db:"scrollname"`
	Dsmitigation        int64          `json:"dsmitigation" db:"dsmitigation"`
	Heroic_str          int64          `json:"heroic_str" db:"heroic_str"`
	Heroic_int          int64          `json:"heroic_int" db:"heroic_int"`
	Heroic_wis          int64          `json:"heroic_wis" db:"heroic_wis"`
	Heroic_agi          int64          `json:"heroic_agi" db:"heroic_agi"`
	Heroic_dex          int64          `json:"heroic_dex" db:"heroic_dex"`
	Heroic_sta          int64          `json:"heroic_sta" db:"heroic_sta"`
	Heroic_cha          int64          `json:"heroic_cha" db:"heroic_cha"`
	Heroic_pr           int64          `json:"heroic_pr" db:"heroic_pr"`
	Heroic_dr           int64          `json:"heroic_dr" db:"heroic_dr"`
	Heroic_fr           int64          `json:"heroic_fr" db:"heroic_fr"`
	Heroic_cr           int64          `json:"heroic_cr" db:"heroic_cr"`
	Heroic_mr           int64          `json:"heroic_mr" db:"heroic_mr"`
	Heroic_svcorrup     int64          `json:"heroic_svcorrup" db:"heroic_svcorrup"`
	Healamt             int64          `json:"healamt" db:"healamt"`
	Spelldmg            int64          `json:"spelldmg" db:"spelldmg"`
	Clairvoyance        int64          `json:"clairvoyance" db:"clairvoyance"`
	Backstabdmg         int64          `json:"backstabdmg" db:"backstabdmg"`
	Created             string         `json:"created" db:"created"`
	Elitematerial       int64          `json:"elitematerial" db:"elitematerial"`
	Ldonsellbackrate    int64          `json:"ldonsellbackrate" db:"ldonsellbackrate"`
	Scriptfileid        int64          `json:"scriptfileid" db:"scriptfileid"`
	Expendablearrow     int64          `json:"expendablearrow" db:"expendablearrow"`
	Powersourcecapacity int64          `json:"powersourcecapacity" db:"powersourcecapacity"`
	Bardeffect          int64          `json:"bardeffect" db:"bardeffect"`
	Bardeffecttype      int64          `json:"bardeffecttype" db:"bardeffecttype"`
	Bardlevel2          int64          `json:"bardlevel2" db:"bardlevel2"`
	Bardlevel           int64          `json:"bardlevel" db:"bardlevel"`
	Bardunk1            int64          `json:"bardunk1" db:"bardunk1"`
	Bardunk2            int64          `json:"bardunk2" db:"bardunk2"`
	Bardunk3            int64          `json:"bardunk3" db:"bardunk3"`
	Bardunk4            int64          `json:"bardunk4" db:"bardunk4"`
	Bardunk5            int64          `json:"bardunk5" db:"bardunk5"`
	Bardname            string         `json:"bardname" db:"bardname"`
	Bardunk7            int64          `json:"bardunk7" db:"bardunk7"`
	Unk214              int64          `json:"UNK214" db:"UNK214"`
	Unk219              int64          `json:"UNK219" db:"UNK219"`
	Unk220              int64          `json:"UNK220" db:"UNK220"`
	Unk221              int64          `json:"UNK221" db:"UNK221"`
	Heirloom            int64          `json:"heirloom" db:"heirloom"`
	Unk223              int64          `json:"UNK223" db:"UNK223"`
	Unk224              int64          `json:"UNK224" db:"UNK224"`
	Unk225              int64          `json:"UNK225" db:"UNK225"`
	Unk226              int64          `json:"UNK226" db:"UNK226"`
	Unk227              int64          `json:"UNK227" db:"UNK227"`
	Unk228              int64          `json:"UNK228" db:"UNK228"`
	Unk229              int64          `json:"UNK229" db:"UNK229"`
	Unk230              int64          `json:"UNK230" db:"UNK230"`
	Unk231              int64          `json:"UNK231" db:"UNK231"`
	Unk232              int64          `json:"UNK232" db:"UNK232"`
	Unk233              int64          `json:"UNK233" db:"UNK233"`
	Unk234              int64          `json:"UNK234" db:"UNK234"`
	Placeable           int64          `json:"placeable" db:"placeable"`
	Unk236              int64          `json:"UNK236" db:"UNK236"`
	Unk237              int64          `json:"UNK237" db:"UNK237"`
	Unk238              int64          `json:"UNK238" db:"UNK238"`
	Unk239              int64          `json:"UNK239" db:"UNK239"`
	Unk240              int64          `json:"UNK240" db:"UNK240"`
	Unk241              int64          `json:"UNK241" db:"UNK241"`
	Epicitem            int64          `json:"epicitem" db:"epicitem"`
}

func (c *Item) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *Item) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "zoneId":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
