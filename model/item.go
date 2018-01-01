package model

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

//Item represents items inside everquest
type Item struct {
	//inventory added details
	CharId            int64  `json:"charid" db:"charid"`                         //`charid` int(11) unsigned NOT NULL DEFAULT '0',
	SlotID            int64  `json:"slotid" db:"slotid"`                         //`slotid` mediumint(7) unsigned NOT NULL DEFAULT '0',
	ItemId            int64  `json:"itemid" db:"itemid"`                         //`itemid` int(11) unsigned DEFAULT '0',
	Charges           int64  `json:"charges" db:"charges"`                       //`charges` smallint(3) unsigned DEFAULT '0',
	InvColor          int64  `json:"invcolor" db:"invcolor"`                     //`color` int(11) unsigned NOT NULL DEFAULT '0',
	Augslot1          int64  `json:"augslot1" db:"augslot1"`                     //`augslot1` mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot2          int64  `json:"augslot2" db:"augslot2"`                     //`augslot2` mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot3          int64  `json:"augslot3" db:"augslot3"`                     //`augslot3` mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot4          int64  `json:"augslot4" db:"augslot4"`                     //`augslot4` mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot5          int64  `json:"augslot5" db:"augslot5"`                     //`augslot5` mediumint(7) unsigned DEFAULT '0',
	Augslot6          int64  `json:"augslot6" db:"augslot6"`                     //`augslot6` mediumint(7) NOT NULL DEFAULT '0',
	Instnodrop        int64  `json:"instnodrop" db:"instnodrop"`                 //`instnodrop` tinyint(1) unsigned NOT NULL DEFAULT '0',
	CustomData        string `json:"customData" db:"custom_data"`                //`custom_data` text,
	Ornamenticon      int64  `json:"ornamenticon" db:"ornamenticon"`             //`ornamenticon` int(11) unsigned NOT NULL DEFAULT '0',
	Ornamentidfile    int64  `json:"ornamentidfile" db:"ornamentidfile"`         //`ornamentidfile` int(11) unsigned NOT NULL DEFAULT '0',
	OrnamentHeroModel int64  `json:"ornamentHeroModel" db:"ornament_hero_model"` //`ornament_hero_model` int(11) NOT NULL DEFAULT '0',

	Character *Character `json:"character"`

	//normal items db
	Id                  int64          `json:"id" db:"id"`                                   //
	Name                string         `json:"name" db:"name"`                               //
	Aagi                int64          `json:"aagi" db:"aagi"`                               //
	Ac                  int64          `json:"ac" db:"ac"`                                   //
	Accuracy            int64          `json:"accuracy" db:"accuracy"`                       //
	Acha                int64          `json:"acha" db:"acha"`                               //
	Adex                int64          `json:"adex" db:"adex"`                               //
	Aint                int64          `json:"aint" db:"aint"`                               //
	Artifactflag        int64          `json:"artifactflag" db:"artifactflag"`               //
	Asta                int64          `json:"asta" db:"asta"`                               //
	Astr                int64          `json:"astr" db:"astr"`                               //
	Attack              int64          `json:"attack" db:"attack"`                           //
	Augrestrict         int64          `json:"augrestrict" db:"augrestrict"`                 //
	Augslot1type        int64          `json:"augslot1type" db:"augslot1type"`               //
	Augslot1visible     sql.NullInt64  `json:"augslot1visible" db:"augslot1visible"`         //
	Augslot2type        int64          `json:"augslot2type" db:"augslot2type"`               //
	Augslot2visible     sql.NullInt64  `json:"augslot2visible" db:"augslot2visible"`         //
	Augslot3type        int64          `json:"augslot3type" db:"augslot3type"`               //
	Augslot3visible     sql.NullInt64  `json:"augslot3visible" db:"augslot3visible"`         //
	Augslot4type        int64          `json:"augslot4type" db:"augslot4type"`               //
	Augslot4visible     sql.NullInt64  `json:"augslot4visible" db:"augslot4visible"`         //
	Augslot5type        int64          `json:"augslot5type" db:"augslot5type"`               //
	Augslot5visible     sql.NullInt64  `json:"augslot5visible" db:"augslot5visible"`         //
	Augslot6type        int64          `json:"augslot6type" db:"augslot6type"`               //
	Augslot6visible     int64          `json:"augslot6visible" db:"augslot6visible"`         //
	Augtype             int64          `json:"augtype" db:"augtype"`                         //
	Avoidance           int64          `json:"avoidance" db:"avoidance"`                     //
	Awis                int64          `json:"awis" db:"awis"`                               //
	Bagsize             int64          `json:"bagsize" db:"bagsize"`                         //
	Bagslots            int64          `json:"bagslots" db:"bagslots"`                       //
	Bagtype             int64          `json:"bagtype" db:"bagtype"`                         //
	Bagwr               int64          `json:"bagwr" db:"bagwr"`                             //
	Banedmgamt          int64          `json:"banedmgamt" db:"banedmgamt"`                   //
	Banedmgraceamt      int64          `json:"banedmgraceamt" db:"banedmgraceamt"`           //
	Banedmgbody         int64          `json:"banedmgbody" db:"banedmgbody"`                 //
	Banedmgrace         int64          `json:"banedmgrace" db:"banedmgrace"`                 //
	Bardtype            int64          `json:"bardtype" db:"bardtype"`                       //
	Bardvalue           int64          `json:"bardvalue" db:"bardvalue"`                     //
	Book                int64          `json:"book" db:"book"`                               //
	Casttime            int64          `json:"casttime" db:"casttime"`                       //
	Casttime2           int64          `json:"casttime_" db:"casttime_"`                     //
	Charmfile           string         `json:"charmfile" db:"charmfile"`                     //
	Charmfileid         string         `json:"charmfileid" db:"charmfileid"`                 //
	Classes             int64          `json:"classes" db:"classes"`                         //
	Color               int64          `json:"color" db:"color"`                             //
	Combateffects       string         `json:"combateffects" db:"combateffects"`             //
	Extradmgskill       int64          `json:"extradmgskill" db:"extradmgskill"`             //
	Extradmgamt         int64          `json:"extradmgamt" db:"extradmgamt"`                 //
	Price               int64          `json:"price" db:"price"`                             //
	Cr                  int64          `json:"cr" db:"cr"`                                   //
	Damage              int64          `json:"damage" db:"damage"`                           //
	Damageshield        int64          `json:"damageshield" db:"damageshield"`               //
	Deity               int64          `json:"deity" db:"deity"`                             //
	Delay               int64          `json:"delay" db:"delay"`                             //
	Augdistiller        int64          `json:"augdistiller" db:"augdistiller"`               //
	Dotshielding        int64          `json:"dotshielding" db:"dotshielding"`               //
	Dr                  int64          `json:"dr" db:"dr"`                                   //
	Clicktype           int64          `json:"clicktype" db:"clicktype"`                     //
	Clicklevel2         int64          `json:"clicklevel2" db:"clicklevel2"`                 //
	Elemdmgtype         int64          `json:"elemdmgtype" db:"elemdmgtype"`                 //
	Elemdmgamt          int64          `json:"elemdmgamt" db:"elemdmgamt"`                   //
	Endur               int64          `json:"endur" db:"endur"`                             //
	Factionamt1         int64          `json:"factionamt1" db:"factionamt1"`                 //
	Factionamt2         int64          `json:"factionamt2" db:"factionamt2"`                 //
	Factionamt3         int64          `json:"factionamt3" db:"factionamt3"`                 //
	Factionamt4         int64          `json:"factionamt4" db:"factionamt4"`                 //
	Factionmod1         int64          `json:"factionmod1" db:"factionmod1"`                 //
	Factionmod2         int64          `json:"factionmod2" db:"factionmod2"`                 //
	Factionmod3         int64          `json:"factionmod3" db:"factionmod3"`                 //
	Factionmod4         int64          `json:"factionmod4" db:"factionmod4"`                 //
	Filename            string         `json:"filename" db:"filename"`                       //
	Focuseffect         int64          `json:"focuseffect" db:"focuseffect"`                 //
	Fr                  int64          `json:"fr" db:"fr"`                                   //
	Fvnodrop            int64          `json:"fvnodrop" db:"fvnodrop"`                       //
	Haste               int64          `json:"haste" db:"haste"`                             //
	Clicklevel          int64          `json:"clicklevel" db:"clicklevel"`                   //
	Hp                  int64          `json:"hp" db:"hp"`                                   //
	Regen               int64          `json:"regen" db:"regen"`                             //
	Icon                int64          `json:"icon" db:"icon"`                               //
	Idfile              string         `json:"idfile" db:"idfile"`                           //
	Itemclass           int64          `json:"itemclass" db:"itemclass"`                     //
	Itemtype            int64          `json:"itemtype" db:"itemtype"`                       //
	Ldonprice           int64          `json:"ldonprice" db:"ldonprice"`                     //
	Ldontheme           int64          `json:"ldontheme" db:"ldontheme"`                     //
	Ldonsold            int64          `json:"ldonsold" db:"ldonsold"`                       //
	Light               int64          `json:"light" db:"light"`                             //
	Lore                string         `json:"lore" db:"lore"`                               //
	Loregroup           int64          `json:"loregroup" db:"loregroup"`                     //
	Magic               int64          `json:"magic" db:"magic"`                             //
	Mana                int64          `json:"mana" db:"mana"`                               //
	Manaregen           int64          `json:"manaregen" db:"manaregen"`                     //
	Enduranceregen      int64          `json:"enduranceregen" db:"enduranceregen"`           //
	Material            int64          `json:"material" db:"material"`                       //
	Herosforgemodel     int64          `json:"herosforgemodel" db:"herosforgemodel"`         //
	Maxcharges          int64          `json:"maxcharges" db:"maxcharges"`                   //
	Mr                  int64          `json:"mr" db:"mr"`                                   //
	Nodrop              int64          `json:"nodrop" db:"nodrop"`                           //
	Norent              int64          `json:"norent" db:"norent"`                           //
	Pendingloreflag     int64          `json:"pendingloreflag" db:"pendingloreflag"`         //
	Pr                  int64          `json:"pr" db:"pr"`                                   //
	Procrate            int64          `json:"procrate" db:"procrate"`                       //
	Races               int64          `json:"races" db:"races"`                             //
	Range               int64          `json:"range" db:"range"`                             //
	Reclevel            int64          `json:"reclevel" db:"reclevel"`                       //
	Recskill            int64          `json:"recskill" db:"recskill"`                       //
	Reqlevel            int64          `json:"reqlevel" db:"reqlevel"`                       //
	Sellrate            float64        `json:"sellrate" db:"sellrate"`                       //
	Shielding           int64          `json:"shielding" db:"shielding"`                     //
	Size                int64          `json:"size" db:"size"`                               //
	Skillmodtype        int64          `json:"skillmodtype" db:"skillmodtype"`               //
	Skillmodvalue       int64          `json:"skillmodvalue" db:"skillmodvalue"`             //
	Slots               int64          `json:"slots" db:"slots"`                             //
	Clickeffect         int64          `json:"clickeffect" db:"clickeffect"`                 //
	Spellshield         int64          `json:"spellshield" db:"spellshield"`                 //
	Strikethrough       int64          `json:"strikethrough" db:"strikethrough"`             //
	Stunresist          int64          `json:"stunresist" db:"stunresist"`                   //
	Summonedflag        int64          `json:"summonedflag" db:"summonedflag"`               //
	Tradeskills         int64          `json:"tradeskills" db:"tradeskills"`                 //
	Favor               int64          `json:"favor" db:"favor"`                             //
	Weight              int64          `json:"weight" db:"weight"`                           //
	Unk012              int64          `json:"UNK012" db:"UNK012"`                           //
	Unk013              int64          `json:"UNK013" db:"UNK013"`                           //
	Benefitflag         int64          `json:"benefitflag" db:"benefitflag"`                 //
	Unk054              int64          `json:"UNK054" db:"UNK054"`                           //
	Unk059              int64          `json:"UNK059" db:"UNK059"`                           //
	Booktype            int64          `json:"booktype" db:"booktype"`                       //
	Recastdelay         int64          `json:"recastdelay" db:"recastdelay"`                 //
	Recasttype          int64          `json:"recasttype" db:"recasttype"`                   //
	Guildfavor          int64          `json:"guildfavor" db:"guildfavor"`                   //
	Unk123              int64          `json:"UNK123" db:"UNK123"`                           //
	Unk124              int64          `json:"UNK124" db:"UNK124"`                           //
	Attuneable          int64          `json:"attuneable" db:"attuneable"`                   //
	Nopet               int64          `json:"nopet" db:"nopet"`                             //
	Updated             mysql.NullTime `json:"updated" db:"updated"`                         //
	Comment             string         `json:"comment" db:"comment"`                         //
	Unk127              int64          `json:"UNK127" db:"UNK127"`                           //
	Pointtype           int64          `json:"pointtype" db:"pointtype"`                     //
	Potionbelt          int64          `json:"potionbelt" db:"potionbelt"`                   //
	Potionbeltslots     int64          `json:"potionbeltslots" db:"potionbeltslots"`         //
	Stacksize           int64          `json:"stacksize" db:"stacksize"`                     //
	Notransfer          int64          `json:"notransfer" db:"notransfer"`                   //
	Stackable           int64          `json:"stackable" db:"stackable"`                     //
	Unk134              string         `json:"UNK134" db:"UNK134"`                           //
	Unk137              int64          `json:"UNK137" db:"UNK137"`                           //
	Proceffect          int64          `json:"proceffect" db:"proceffect"`                   //
	Proctype            int64          `json:"proctype" db:"proctype"`                       //
	Proclevel2          int64          `json:"proclevel2" db:"proclevel2"`                   //
	Proclevel           int64          `json:"proclevel" db:"proclevel"`                     //
	Unk142              int64          `json:"UNK142" db:"UNK142"`                           //
	Worneffect          int64          `json:"worneffect" db:"worneffect"`                   //
	Worntype            int64          `json:"worntype" db:"worntype"`                       //
	Wornlevel2          int64          `json:"wornlevel2" db:"wornlevel2"`                   //
	Wornlevel           int64          `json:"wornlevel" db:"wornlevel"`                     //
	Unk147              int64          `json:"UNK147" db:"UNK147"`                           //
	Focustype           int64          `json:"focustype" db:"focustype"`                     //
	Focuslevel2         int64          `json:"focuslevel2" db:"focuslevel2"`                 //
	Focuslevel          int64          `json:"focuslevel" db:"focuslevel"`                   //
	Unk152              int64          `json:"UNK152" db:"UNK152"`                           //
	Scrolleffect        int64          `json:"scrolleffect" db:"scrolleffect"`               //
	Scrolltype          int64          `json:"scrolltype" db:"scrolltype"`                   //
	Scrolllevel2        int64          `json:"scrolllevel2" db:"scrolllevel2"`               //
	Scrolllevel         int64          `json:"scrolllevel" db:"scrolllevel"`                 //
	Unk157              int64          `json:"UNK157" db:"UNK157"`                           //
	Serialized          mysql.NullTime `json:"serialized" db:"serialized"`                   //
	Verified            mysql.NullTime `json:"verified" db:"verified"`                       //
	Serialization       sql.NullString `json:"serialization" db:"serialization"`             //
	Source              string         `json:"source" db:"source"`                           //
	Unk033              int64          `json:"UNK033" db:"UNK033"`                           //
	Lorefile            string         `json:"lorefile" db:"lorefile"`                       //
	Unk014              int64          `json:"UNK014" db:"UNK014"`                           //
	Svcorruption        int64          `json:"svcorruption" db:"svcorruption"`               //
	Skillmodmax         int64          `json:"skillmodmax" db:"skillmodmax"`                 //
	Unk060              int64          `json:"UNK060" db:"UNK060"`                           //
	Augslot1unk2        int64          `json:"augslot1unk2" db:"augslot1unk2"`               //
	Augslot2unk2        int64          `json:"augslot2unk2" db:"augslot2unk2"`               //
	Augslot3unk2        int64          `json:"augslot3unk2" db:"augslot3unk2"`               //
	Augslot4unk2        int64          `json:"augslot4unk2" db:"augslot4unk2"`               //
	Augslot5unk2        int64          `json:"augslot5unk2" db:"augslot5unk2"`               //
	Augslot6unk2        int64          `json:"augslot6unk2" db:"augslot6unk2"`               //
	Unk120              int64          `json:"UNK120" db:"UNK120"`                           //
	Unk121              int64          `json:"UNK121" db:"UNK121"`                           //
	Questitemflag       int64          `json:"questitemflag" db:"questitemflag"`             //
	Unk132              sql.NullString `json:"UNK132" db:"UNK132"`                           //
	Clickunk5           int64          `json:"clickunk5" db:"clickunk5"`                     //
	Clickunk6           string         `json:"clickunk6" db:"clickunk6"`                     //
	Clickunk7           int64          `json:"clickunk7" db:"clickunk7"`                     //
	Procunk1            int64          `json:"procunk1" db:"procunk1"`                       //
	Procunk2            int64          `json:"procunk2" db:"procunk2"`                       //
	Procunk3            int64          `json:"procunk3" db:"procunk3"`                       //
	Procunk4            int64          `json:"procunk4" db:"procunk4"`                       //
	Procunk6            string         `json:"procunk6" db:"procunk6"`                       //
	Procunk7            int64          `json:"procunk7" db:"procunk7"`                       //
	Wornunk1            int64          `json:"wornunk1" db:"wornunk1"`                       //
	Wornunk2            int64          `json:"wornunk2" db:"wornunk2"`                       //
	Wornunk3            int64          `json:"wornunk3" db:"wornunk3"`                       //
	Wornunk4            int64          `json:"wornunk4" db:"wornunk4"`                       //
	Wornunk5            int64          `json:"wornunk5" db:"wornunk5"`                       //
	Wornunk6            string         `json:"wornunk6" db:"wornunk6"`                       //
	Wornunk7            int64          `json:"wornunk7" db:"wornunk7"`                       //
	Focusunk1           int64          `json:"focusunk1" db:"focusunk1"`                     //
	Focusunk2           int64          `json:"focusunk2" db:"focusunk2"`                     //
	Focusunk3           int64          `json:"focusunk3" db:"focusunk3"`                     //
	Focusunk4           int64          `json:"focusunk4" db:"focusunk4"`                     //
	Focusunk5           int64          `json:"focusunk5" db:"focusunk5"`                     //
	Focusunk6           string         `json:"focusunk6" db:"focusunk6"`                     //
	Focusunk7           int64          `json:"focusunk7" db:"focusunk7"`                     //
	Scrollunk1          int64          `json:"scrollunk1" db:"scrollunk1"`                   //
	Scrollunk2          int64          `json:"scrollunk2" db:"scrollunk2"`                   //
	Scrollunk3          int64          `json:"scrollunk3" db:"scrollunk3"`                   //
	Scrollunk4          int64          `json:"scrollunk4" db:"scrollunk4"`                   //
	Scrollunk5          int64          `json:"scrollunk5" db:"scrollunk5"`                   //
	Scrollunk6          string         `json:"scrollunk6" db:"scrollunk6"`                   //
	Scrollunk7          int64          `json:"scrollunk7" db:"scrollunk7"`                   //
	Unk193              int64          `json:"UNK193" db:"UNK193"`                           //
	Purity              int64          `json:"purity" db:"purity"`                           //
	Evoitem             int64          `json:"evoitem" db:"evoitem"`                         //
	Evoid               int64          `json:"evoid" db:"evoid"`                             //
	Evolvinglevel       int64          `json:"evolvinglevel" db:"evolvinglevel"`             //
	Evomax              int64          `json:"evomax" db:"evomax"`                           //
	Clickname           string         `json:"clickname" db:"clickname"`                     //
	Procname            string         `json:"procname" db:"procname"`                       //
	Wornname            string         `json:"wornname" db:"wornname"`                       //
	Focusname           string         `json:"focusname" db:"focusname"`                     //
	Scrollname          string         `json:"scrollname" db:"scrollname"`                   //
	Dsmitigation        int64          `json:"dsmitigation" db:"dsmitigation"`               //
	HeroicStr           int64          `json:"heroicStr" db:"heroic_str"`                    //
	HeroicInt           int64          `json:"heroicInt" db:"heroic_int"`                    //
	HeroicWis           int64          `json:"heroicWis" db:"heroic_wis"`                    //
	HeroicAgi           int64          `json:"heroicAgi" db:"heroic_agi"`                    //
	HeroicDex           int64          `json:"heroicDex" db:"heroic_dex"`                    //
	HeroicSta           int64          `json:"heroicSta" db:"heroic_sta"`                    //
	HeroicCha           int64          `json:"heroicCha" db:"heroic_cha"`                    //
	HeroicPr            int64          `json:"heroicPr" db:"heroic_pr"`                      //
	HeroicDr            int64          `json:"heroicDr" db:"heroic_dr"`                      //
	HeroicFr            int64          `json:"heroicFr" db:"heroic_fr"`                      //
	HeroicCr            int64          `json:"heroicCr" db:"heroic_cr"`                      //
	HeroicMr            int64          `json:"heroicMr" db:"heroic_mr"`                      //
	HeroicSvcorrup      int64          `json:"heroicSvcorrup" db:"heroic_svcorrup"`          //
	Healamt             int64          `json:"healamt" db:"healamt"`                         //
	Spelldmg            int64          `json:"spelldmg" db:"spelldmg"`                       //
	Clairvoyance        int64          `json:"clairvoyance" db:"clairvoyance"`               //
	Backstabdmg         int64          `json:"backstabdmg" db:"backstabdmg"`                 //
	Created             string         `json:"created" db:"created"`                         //
	Elitematerial       int64          `json:"elitematerial" db:"elitematerial"`             //
	Ldonsellbackrate    int64          `json:"ldonsellbackrate" db:"ldonsellbackrate"`       //
	Scriptfileid        int64          `json:"scriptfileid" db:"scriptfileid"`               //
	Expendablearrow     int64          `json:"expendablearrow" db:"expendablearrow"`         //
	Powersourcecapacity int64          `json:"powersourcecapacity" db:"powersourcecapacity"` //
	Bardeffect          int64          `json:"bardeffect" db:"bardeffect"`                   //
	Bardeffecttype      int64          `json:"bardeffecttype" db:"bardeffecttype"`           //
	Bardlevel2          int64          `json:"bardlevel2" db:"bardlevel2"`                   //
	Bardlevel           int64          `json:"bardlevel" db:"bardlevel"`                     //
	Bardunk1            int64          `json:"bardunk1" db:"bardunk1"`                       //
	Bardunk2            int64          `json:"bardunk2" db:"bardunk2"`                       //
	Bardunk3            int64          `json:"bardunk3" db:"bardunk3"`                       //
	Bardunk4            int64          `json:"bardunk4" db:"bardunk4"`                       //
	Bardunk5            int64          `json:"bardunk5" db:"bardunk5"`                       //
	Bardname            string         `json:"bardname" db:"bardname"`                       //
	Bardunk7            int64          `json:"bardunk7" db:"bardunk7"`                       //
	Unk214              int64          `json:"UNK214" db:"UNK214"`                           //
	Unk219              int64          `json:"UNK219" db:"UNK219"`                           //
	Unk220              int64          `json:"UNK220" db:"UNK220"`                           //
	Unk221              int64          `json:"UNK221" db:"UNK221"`                           //
	Heirloom            int64          `json:"heirloom" db:"heirloom"`                       //
	Unk223              int64          `json:"UNK223" db:"UNK223"`                           //
	Unk224              int64          `json:"UNK224" db:"UNK224"`                           //
	Unk225              int64          `json:"UNK225" db:"UNK225"`                           //
	Unk226              int64          `json:"UNK226" db:"UNK226"`                           //
	Unk227              int64          `json:"UNK227" db:"UNK227"`                           //
	Unk228              int64          `json:"UNK228" db:"UNK228"`                           //
	Unk229              int64          `json:"UNK229" db:"UNK229"`                           //
	Unk230              int64          `json:"UNK230" db:"UNK230"`                           //
	Unk231              int64          `json:"UNK231" db:"UNK231"`                           //
	Unk232              int64          `json:"UNK232" db:"UNK232"`                           //
	Unk233              int64          `json:"UNK233" db:"UNK233"`                           //
	Unk234              int64          `json:"UNK234" db:"UNK234"`                           //
	Placeable           int64          `json:"placeable" db:"placeable"`                     //
	Unk236              int64          `json:"UNK236" db:"UNK236"`                           //
	Unk237              int64          `json:"UNK237" db:"UNK237"`                           //
	Unk238              int64          `json:"UNK238" db:"UNK238"`                           //
	Unk239              int64          `json:"UNK239" db:"UNK239"`                           //
	Unk240              int64          `json:"UNK240" db:"UNK240"`                           //
	Unk241              int64          `json:"UNK241" db:"UNK241"`                           //
	Epicitem            int64          `json:"epicitem" db:"epicitem"`                       //
}

func (c *Item) SizeName() string {
	return "TINY"
}

func (c *Item) SlotList() string {
	slots := ""
	if c.Slots&1 == 1 {
		slots += "CHARM "
	}
	if c.Slots&4 == 4 {
		slots += "HEAD "
	}
	if c.Slots&8 == 8 {
		slots += "FACE "
	}
	if c.Slots&18 == 18 {
		slots += "EARS "
	}
	if c.Slots&32 == 32 {
		slots += "NECK "
	}
	if c.Slots&64 == 64 {
		slots += "SHOULDER "
	}
	if c.Slots&128 == 128 {
		slots += "ARMS "
	}
	if c.Slots&256 == 256 {
		slots += "BACK "
	}
	if c.Slots&1536 == 1536 {
		slots += "BRACERS "
	}
	if c.Slots&2048 == 2048 {
		slots += "RANGE "
	}
	if c.Slots&4096 == 4096 {
		slots += "HANDS "
	}
	if c.Slots&8192 == 8192 {
		slots += "PRIMARY "
	}
	if c.Slots&16384 == 16384 {
		slots += "SECONDARY "
	}
	if c.Slots&98304 == 98304 {
		slots += "RINGS "
	}
	if c.Slots&131072 == 131072 {
		slots += "CHEST "
	}
	if c.Slots&262144 == 262144 {
		slots += "LEGS "
	}
	if c.Slots&524288 == 524288 {
		slots += "FEET "
	}
	if c.Slots&1048576 == 1048576 {
		slots += "WAIST "
	}
	if c.Slots&2097152 == 2097152 {
		slots += "AMMO "
	}
	if c.Slots&4194304 == 4194304 {
		slots += "POWERSOURCE "
	}
	if len(slots) > 0 {
		slots = slots[0 : len(slots)-1]
	}
	return slots
}

func (c *Item) ClassList() string {
	classes := ""
	if c.Classes == 65535 {
		return "ALL"
	}
	if c.Classes&1 == 1 {
		classes += "WAR "
	}
	if c.Classes&2 == 2 {
		classes += "CLR "
	}
	if c.Classes&4 == 4 {
		classes += "PAL "
	}
	if c.Classes&8 == 8 {
		classes += "RNG "
	}
	if c.Classes&16 == 16 {
		classes += "SHM "
	}
	if c.Classes&32 == 32 {
		classes += "DRU "
	}
	if c.Classes&64 == 64 {
		classes += "MNK "
	}
	if c.Classes&128 == 128 {
		classes += "BRD "
	}
	if c.Classes&256 == 256 {
		classes += "ROG "
	}
	if c.Classes&512 == 512 {
		classes += "SHD "
	}
	if c.Classes&1024 == 1024 {
		classes += "NEC "
	}
	if c.Classes&2048 == 2048 {
		classes += "WIZ "
	}
	if c.Classes&4096 == 4096 {
		classes += "MAG "
	}
	if c.Classes&8192 == 8192 {
		classes += "ENC "
	}
	if c.Classes&16384 == 16384 {
		classes += "BST "
	}
	if c.Classes&32768 == 32768 {
		classes += "BER "
	}
	if len(classes) > 0 {
		classes = classes[0 : len(classes)-1]
	}
	if len(classes) == 0 {
		classes = "NONE"
	}
	return classes
}

func (c *Item) RaceList() string {
	races := ""
	if c.Races == 65535 {
		return "ALL"
	}
	if c.Races&1 == 1 {
		races += "HUM "
	}
	if c.Races&2 == 2 {
		races += "BAR "
	}
	if c.Races&4 == 4 {
		races += "ERU "
	}
	if c.Races&8 == 8 {
		races += "WEF "
	}
	if c.Races&16 == 16 {
		races += "HEF "
	}
	if c.Races&32 == 32 {
		races += "DEF "
	}
	if c.Races&64 == 64 {
		races += "HLF "
	}
	if c.Races&128 == 128 {
		races += "DWF "
	}
	if c.Races&256 == 256 {
		races += "TRL "
	}
	if c.Races&512 == 512 {
		races += "OGR "
	}
	if c.Races&1024 == 1024 {
		races += "HLF "
	}
	if c.Races&2048 == 2048 {
		races += "GNM "
	}
	if c.Races&4096 == 4096 {
		races += "IKS "
	}
	if c.Races&8192 == 8192 {
		races += "VAH "
	}
	if c.Races&16384 == 16384 {
		races += "FRO "
	}
	if c.Races&32768 == 32768 {
		races += "DRA "
	}
	if len(races) > 0 {
		races = races[0 : len(races)-1]
	}
	if len(races) == 0 {
		races = "NONE"
	}
	return races
}

func (c *Item) ItemtypeName() string {
	switch c.Itemtype {
	case 0:
		return "1 Hand Slash"
	case 1:
		return "2 Hand Slash"
	case 2:
		return "Piercing"
	case 3:
		return "1 Hand Blunt"
	case 4:
		return "2 Hand Blunt"
	case 5:
		return "Archery"
	case 6:
		return "Unused (6)"
	case 7:
		return "Throwing"
	case 8:
		return "Shield"
	case 9:
		return "Unused"
	case 10:
		return "Armor"
	case 11:
		return "Tradeskills"
	case 12:
		return "Lock Picking"
	case 13:
		return "Unused (13)"
	case 14:
		return "Food"
	case 15:
		return "Drink"
	case 16:
		return "Light Source"
	case 17:
		return "Common Inventory Item"
	case 18:
		return "Bind Wound"
	case 19:
		return "Thrown Casting Items (Explosive potions etc)"
	case 20:
		return "Spells / Song Sheets"
	case 21:
		return "Potions"
	case 22:
		return "Fletched Arrows?..."
	case 23:
		return "Wind Instruments"
	case 24:
		return "Stringed Instruments"
	case 25:
		return "Brass Instruments"
	case 26:
		return "Drum Instruments"
	case 27:
		return "Ammo (In most cases, Arrows)"
	case 28:
		return "Unused"
	case 29:
		return "Jewlery Items (As far as I can tell)"
	case 30:
		return "Unused"
	case 31:
		return "Usually Readable Notes and Scrolls"
	case 32:
		return "Usually Readable Books"
	case 33:
		return "Keys"
	case 34:
		return "Odd Items (Not sure what they are for)"
	case 35:
		return "2H Pierce"
	case 36:
		return "Fishing Poles"
	case 37:
		return "Fishing Bait"
	case 38:
		return "Alcoholic Beverages"
	case 39:
		return "More Keys"
	case 40:
		return "Compasses"
	case 41:
		return "Unused"
	case 42:
		return "Poisens"
	case 43:
		return "Unused"
	case 44:
		return "Unused"
	case 45:
		return "H2H (Hand to Hand)"
	case 46:
		return "Unused"
	case 47:
		return "Unused"
	case 48:
		return "Unused"
	case 49:
		return "Unused"
	case 50:
		return "Unused"
	case 51:
		return "Unused"
	case 52:
		return "Charms"
	case 53:
		return "Dyes"
	case 54:
		return "Augments"
	case 55:
		return "Augment Solvents"
	case 56:
		return "Augment Distillers"
	case 58:
		return "Fellowship Banner Materials"
	case 60:
		return "Cultural Armor Manuals"
	case 63:
		return "Currency"
	}
	return "Unknown"
}

func (c *Item) ItemtypeIcon() string {
	switch c.Itemtype {
	case 0:
		return "xa-crossed-swords"
	case 1:
		return "xa-croc-sword"
	case 2:
		return "xa-plain-dagger"
	case 3:
		return "xa-flat-hammer"
	case 4:
		return "xa-gavel"
	case 5:
		return "xa-crossbow"
	case 6:
		return "xa-help"
	case 7:
		return "xa-hammer-drop"
	case 8:
		return "xa-fire-shield"
	case 9:
		return "xa-help"
	case 10:
		return "xa-vest"
	case 11:
		return "xa-archery-target" //Involves Tradeskills (Not sure how)
	case 12:
		return "xa-key"
	case 13:
		return "xa-help"
	case 14:
		return "xa-apple"
	case 15:
		return "xa-brandy-bottle"
	case 16:
		return "xa-light-bulb"
	case 17:
		return "xa-shovel"
	case 18:
		return "xa-health"
	case 19:
		return "xa-bottled-bolt"
	case 20:
		return "xa-scroll-unfurled"
	case 21:
		return "xa-flask"
	case 22:
		return "xa-arrow-flights"
	case 23:
		return "xa-ocarina"
	case 24:
		return "xa-ocarina"
	case 25:
		return "xa-ocarina"
	case 26:
		return "xa-ocarina"
	case 27:
		return "xa-broadhead-arrow"
	case 28:
		return "xa-help"
	case 29:
		return "xa-explosion"
	case 30:
		return "xa-help"
	case 31:
		return "xa-book" //Usually Readable Notes and Scrolls"
	case 32:
		return "xa-book" //Usually Readable Books"
	case 33:
		return "xa-key"
	case 34:
		return "xa-vail" //Odd Items (Not sure what they are for)"
	case 35:
		return "xa-relic-blade" //2hp
	case 36:
		return "xa-fish"
	case 37:
		return "xa-venomous-snake"
	case 38:
		return "xa-beer"
	case 39:
		return "xa-key"
	case 40:
		return "xa-compass"
	case 41:
		return "xa-help"
	case 42:
		return "xa-bottle-vapors"
	case 43:
		return "xa-help"
	case 44:
		return "xa-help"
	case 45:
		return "xa-hand"
	case 46:
		return "xa-help"
	case 47:
		return "xa-help"
	case 48:
		return "xa-help"
	case 49:
		return "xa-help"
	case 50:
		return "xa-help"
	case 51:
		return "xa-help"
	case 52:
		return "xa-sapphire"
	case 53:
		return "xa-round-bottome-flask"
	case 54:
		return "xa-bubbling-potion"
	case 55:
		return "xa-corked-tube"
	case 56:
		return "xa-corked-tube"
	case 58:
		return "xa-castle-flag"
	case 60:
		return "xa-book"
	case 63:
		return "xa-sapphire"
	}
	return "xa-help"
}

func (c *Item) SlotsFirstName() string {
	switch {
	case c.Slots&8192 == 8192:
		if c.Slots > 8192 {
			return "Primary+"
		}
		return "Primary"
	case c.Slots&16384 == 16384:
		if c.Slots > 16384 {
			return "Secondary+"
		}
		return "Secondary"
	case c.Slots&2048 == 2048:
		if c.Slots > 2048 {
			return "Range+"
		}
		return "Range"
	case c.Slots&1 == 1:
		if c.Slots > 1 {
			return "Charm+"
		}
		return "Charm"
	case c.Slots&4 == 4:
		if c.Slots > 4 {
			return "Head+"
		}
		return "Head"
	case c.Slots&8 == 8:
		if c.Slots > 8 {
			return "Face+"
		}
		return "Face"
	case c.Slots&18 == 18:
		if c.Slots > 18 {
			return "Ears+"
		}
		return "Ears"
	case c.Slots&32 == 32:
		if c.Slots > 32 {
			return "Neck+"
		}
		return "Neck"
	case c.Slots&64 == 64:
		if c.Slots > 64 {
			return "Shoulder+"
		}
		return "Shoulder"
	case c.Slots&128 == 128:
		if c.Slots > 128 {
			return "Arms+"
		}
		return "Arms"
	case c.Slots&256 == 256:
		if c.Slots > 256 {
			return "Back+"
		}
		return "Back"
	case c.Slots&1536 == 1536:
		if c.Slots > 1536 {
			return "Bracers+"
		}
		return "Bracers"
	case c.Slots&4096 == 4096:
		if c.Slots > 4096 {
			return "Hands+"
		}
		return "Hands"
	case c.Slots&98304 == 98304:
		if c.Slots > 98304 {
			return "Rings+"
		}
		return "Rings"
	case c.Slots&131072 == 131072:
		if c.Slots > 131072 {
			return "Chest+"
		}
		return "Chest"
	case c.Slots&262144 == 262144:
		if c.Slots > 262144 {
			return "Legs+"
		}
		return "Legs"
	case c.Slots&524288 == 524288:
		if c.Slots > 524288 {
			return "Feet+"
		}
		return "Feet"
	case c.Slots&1048576 == 1048576:
		if c.Slots > 1048576 {
			return "Waist+"
		}
		return "Waist"
	case c.Slots&2097152 == 2097152:
		if c.Slots > 2097152 {
			return "Ammo+"
		}
		return "Ammo"
	case c.Slots&4194304 == 4194304:
		if c.Slots > 4194304 {
			return "Powersource+"
		}
		return "Powersource"
	}
	return "None"
}

func (c *Item) SlotName() string {
	s := c.SlotID
	switch {
	case s == 0:
		return "Charm"
	case s == 1:
		return "Left Ear"
	case s == 2:
		return "Head"
	case s == 3:
		return "Face"
	case s == 4:
		return "Right Ear"
	case s == 5:
		return "Neck"
	case s == 6:
		return "Shoulder"
	case s == 7:
		return "Arms"
	case s == 8:
		return "Back"
	case s == 9:
		return "Left Bracer"
	case s == 10:
		return "Right Bracer"
	case s == 11:
		return "Range"
	case s == 12:
		return "Hands"
	case s == 13:
		return "Primary"
	case s == 14:
		return "Secondary"
	case s == 15:
		return "Left Ring"
	case s == 16:
		return "Right Ring"
	case s == 17:
		return "Chest"
	case s == 18:
		return "Legs"
	case s == 19:
		return "Feet"
	case s == 20:
		return "Waist"
	case s == 21:
		return "Ammo"
	case s == 22:
		return "TopLeft Inventory"
	case s <= 271 && s >= 262:
		return "TopLeft Bag"
	case s == 23:
		return "TopRight Inventory"
	case s <= 281 && s >= 272:
		return "TopRight Bag"
	case s == 24:
		return "TopLeft, One Down Inventory"
	case s <= 291 && s >= 282:
		return "TopLeft, One Down Bag"
	case s == 25:
		return "TopRight, One Down Inventory"
	case s <= 301 && s >= 292:
		return "TopRight, One Down Bag"
	case s == 26:
		return "BottomLeft, Two Up Inventory"
	case s <= 311 && s >= 302:
		return "BottomLeft, Two Up Bag"
	case s == 27:
		return "BottomRight, Two Up Inventory"
	case s <= 321 && s >= 312:
		return "BottomRight, Two Up Bag"
	case s == 28:
		return "BottomLeft Inventory"
	case s <= 331 && s >= 322:
		return "BottomLeft Bag"
	case s == 29:
		return "BottomRight Inventory"
	case s <= 341 && s >= 332:
		return "BottomRight Bag"
	case s == 30:
		return "Cursor"
	case s >= 2000 && s <= 2271:
		return "Bank"
	case s >= 400 && s <= 404:
		return "Tribute"
	case s >= 2500 && s <= 2551:
		return "Shared Bank"
	default:
		return fmt.Sprintf("Unknown (%d)", c.SlotID)
	}
}
