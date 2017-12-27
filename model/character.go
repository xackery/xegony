package model

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

type Character struct {
	Id                    int64   `json:"id" db:"id"`                                         //`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	AccountId             int64   `json:"accountId" db:"account_id"`                          //`account_id` int(11) NOT NULL DEFAULT '0',
	Name                  string  `json:"name" db:"name"`                                     //`name` varchar(64) NOT NULL DEFAULT '',
	LastName              string  `json:"lastName" db:"last_name"`                            //`last_name` varchar(64) NOT NULL DEFAULT '',
	Title                 string  `json:"title" db:"title"`                                   //`title` varchar(32) NOT NULL DEFAULT '',
	Suffix                string  `json:"suffix" db:"suffix"`                                 //`suffix` varchar(32) NOT NULL DEFAULT '',
	ZoneId                int64   `json:"zoneId" db:"zone_id"`                                //`zone_id` int(11) unsigned NOT NULL DEFAULT '0',
	ZoneInstance          int64   `json:"zoneInstance" db:"zone_instance"`                    //`zone_instance` int(11) unsigned NOT NULL DEFAULT '0',
	Y                     float64 `json:"y" db:"y"`                                           //`y` float NOT NULL DEFAULT '0',
	X                     float64 `json:"x" db:"x"`                                           //`x` float NOT NULL DEFAULT '0',
	Z                     float64 `json:"z" db:"z"`                                           //`z` float NOT NULL DEFAULT '0',
	Heading               float64 `json:"heading" db:"heading"`                               //`heading` float NOT NULL DEFAULT '0',
	Gender                int64   `json:"gender" db:"gender"`                                 //`gender` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Race                  int64   `json:"race" db:"race"`                                     //`race` smallint(11) unsigned NOT NULL DEFAULT '0',
	Class                 int64   `json:"class" db:"class"`                                   //`class` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Level                 int64   `json:"level" db:"level"`                                   //`level` int(11) unsigned NOT NULL DEFAULT '0',
	Deity                 int64   `json:"deity" db:"deity"`                                   //`deity` int(11) unsigned NOT NULL DEFAULT '0',
	Birthday              int64   `json:"birthday" db:"birthday"`                             //`birthday` int(11) unsigned NOT NULL DEFAULT '0',
	LastLogin             int64   `json:"lastLogin" db:"last_login"`                          //`last_login` int(11) unsigned NOT NULL DEFAULT '0',
	TimePlayed            int64   `json:"timePlayed" db:"time_played"`                        //`time_played` int(11) unsigned NOT NULL DEFAULT '0',
	Level2                int64   `json:"level2" db:"level2"`                                 //`level2` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Anon                  int64   `json:"anon" db:"anon"`                                     //`anon` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Gm                    int64   `json:"gm" db:"gm"`                                         //`gm` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Face                  int64   `json:"face" db:"face"`                                     //`face` int(11) unsigned NOT NULL DEFAULT '0',
	HairColor             int64   `json:"hairColor" db:"hair_color"`                          //`hair_color` tinyint(11) unsigned NOT NULL DEFAULT '0',
	HairStyle             int64   `json:"hairStyle" db:"hair_style"`                          //`hair_style` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Beard                 int64   `json:"beard" db:"beard"`                                   //`beard` tinyint(11) unsigned NOT NULL DEFAULT '0',
	BeardColor            int64   `json:"beardColor" db:"beard_color"`                        //`beard_color` tinyint(11) unsigned NOT NULL DEFAULT '0',
	EyeColor1             int64   `json:"eyeColor1" db:"eye_color_1"`                         //`eye_color_1` tinyint(11) unsigned NOT NULL DEFAULT '0',
	EyeColor2             int64   `json:"eyeColor2" db:"eye_color_2"`                         //`eye_color_2` tinyint(11) unsigned NOT NULL DEFAULT '0',
	DrakkinHeritage       int64   `json:"drakkinHeritage" db:"drakkin_heritage"`              //`drakkin_heritage` int(11) unsigned NOT NULL DEFAULT '0',
	DrakkinTattoo         int64   `json:"drakkinTattoo" db:"drakkin_tattoo"`                  //`drakkin_tattoo` int(11) unsigned NOT NULL DEFAULT '0',
	DrakkinDetails        int64   `json:"drakkinDetails" db:"drakkin_details"`                //`drakkin_details` int(11) unsigned NOT NULL DEFAULT '0',
	AbilityTimeSeconds    int64   `json:"abilityTimeSeconds" db:"ability_time_seconds"`       //`ability_time_seconds` tinyint(11) unsigned NOT NULL DEFAULT '0',
	AbilityNumber         int64   `json:"abilityNumber" db:"ability_number"`                  //`ability_number` tinyint(11) unsigned NOT NULL DEFAULT '0',
	AbilityTimeMinutes    int64   `json:"abilityTimeMinutes" db:"ability_time_minutes"`       //`ability_time_minutes` tinyint(11) unsigned NOT NULL DEFAULT '0',
	AbilityTimeHours      int64   `json:"abilityTimeHours" db:"ability_time_hours"`           //`ability_time_hours` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Exp                   int64   `json:"exp" db:"exp"`                                       //`exp` int(11) unsigned NOT NULL DEFAULT '0',
	AaPointsSpent         int64   `json:"aaPointsSpent" db:"aa_points_spent"`                 //`aa_points_spent` int(11) unsigned NOT NULL DEFAULT '0',
	AaExp                 int64   `json:"aaExp" db:"aa_exp"`                                  //`aa_exp` int(11) unsigned NOT NULL DEFAULT '0',
	AaPoints              int64   `json:"aaPoints" db:"aa_points"`                            //`aa_points` int(11) unsigned NOT NULL DEFAULT '0',
	GroupLeadershipExp    int64   `json:"groupLeadershipExp" db:"group_leadership_exp"`       //`group_leadership_exp` int(11) unsigned NOT NULL DEFAULT '0',
	RaidLeadershipExp     int64   `json:"raidLeadershipExp" db:"raid_leadership_exp"`         //`raid_leadership_exp` int(11) unsigned NOT NULL DEFAULT '0',
	GroupLeadershipPoints int64   `json:"groupLeadershipPoints" db:"group_leadership_points"` //`group_leadership_points` int(11) unsigned NOT NULL DEFAULT '0',
	RaidLeadershipPoints  int64   `json:"raidLeadershipPoints" db:"raid_leadership_points"`   //`raid_leadership_points` int(11) unsigned NOT NULL DEFAULT '0',
	Points                int64   `json:"points" db:"points"`                                 //`points` int(11) unsigned NOT NULL DEFAULT '0',
	CurHp                 int64   `json:"curHp" db:"cur_hp"`                                  //`cur_hp` int(11) unsigned NOT NULL DEFAULT '0',
	Mana                  int64   `json:"mana" db:"mana"`                                     //`mana` int(11) unsigned NOT NULL DEFAULT '0',
	Endurance             int64   `json:"endurance" db:"endurance"`                           //`endurance` int(11) unsigned NOT NULL DEFAULT '0',
	Intoxication          int64   `json:"intoxication" db:"intoxication"`                     //`intoxication` int(11) unsigned NOT NULL DEFAULT '0',
	Str                   int64   `json:"str" db:"str"`                                       //`str` int(11) unsigned NOT NULL DEFAULT '0',
	Sta                   int64   `json:"sta" db:"sta"`                                       //`sta` int(11) unsigned NOT NULL DEFAULT '0',
	Cha                   int64   `json:"cha" db:"cha"`                                       //`cha` int(11) unsigned NOT NULL DEFAULT '0',
	Dex                   int64   `json:"dex" db:"dex"`                                       //`dex` int(11) unsigned NOT NULL DEFAULT '0',
	Int                   int64   `json:"int" db:"int"`                                       //`int` int(11) unsigned NOT NULL DEFAULT '0',
	Agi                   int64   `json:"agi" db:"agi"`                                       //`agi` int(11) unsigned NOT NULL DEFAULT '0',
	Wis                   int64   `json:"wis" db:"wis"`                                       //`wis` int(11) unsigned NOT NULL DEFAULT '0',
	ZoneChangeCount       int64   `json:"zoneChangeCount" db:"zone_change_count"`             //`zone_change_count` int(11) unsigned NOT NULL DEFAULT '0',
	Toxicity              int64   `json:"toxicity" db:"toxicity"`                             //`toxicity` int(11) unsigned NOT NULL DEFAULT '0',
	HungerLevel           int64   `json:"hungerLevel" db:"hunger_level"`                      //`hunger_level` int(11) unsigned NOT NULL DEFAULT '0',
	ThirstLevel           int64   `json:"thirstLevel" db:"thirst_level"`                      //`thirst_level` int(11) unsigned NOT NULL DEFAULT '0',
	AbilityUp             int64   `json:"abilityUp" db:"ability_up"`                          //`ability_up` int(11) unsigned NOT NULL DEFAULT '0',
	LdonPointsGuk         int64   `json:"ldonPointsGuk" db:"ldon_points_guk"`                 //`ldon_points_guk` int(11) unsigned NOT NULL DEFAULT '0',
	LdonPointsMir         int64   `json:"ldonPointsMir" db:"ldon_points_mir"`                 //`ldon_points_mir` int(11) unsigned NOT NULL DEFAULT '0',
	LdonPointsMmc         int64   `json:"ldonPointsMmc" db:"ldon_points_mmc"`                 //`ldon_points_mmc` int(11) unsigned NOT NULL DEFAULT '0',
	LdonPointsRuj         int64   `json:"ldonPointsRuj" db:"ldon_points_ruj"`                 //`ldon_points_ruj` int(11) unsigned NOT NULL DEFAULT '0',
	LdonPointsTak         int64   `json:"ldonPointsTak" db:"ldon_points_tak"`                 //`ldon_points_tak` int(11) unsigned NOT NULL DEFAULT '0',
	LdonPointsAvailable   int64   `json:"ldonPointsAvailable" db:"ldon_points_available"`     //`ldon_points_available` int(11) unsigned NOT NULL DEFAULT '0',
	TributeTimeRemaining  int64   `json:"tributeTimeRemaining" db:"tribute_time_remaining"`   //`tribute_time_remaining` int(11) unsigned NOT NULL DEFAULT '0',
	CareerTributePoints   int64   `json:"careerTributePoints" db:"career_tribute_points"`     //`career_tribute_points` int(11) unsigned NOT NULL DEFAULT '0',
	TributePoints         int64   `json:"tributePoints" db:"tribute_points"`                  //`tribute_points` int(11) unsigned NOT NULL DEFAULT '0',
	TributeActive         int64   `json:"tributeActive" db:"tribute_active"`                  //`tribute_active` int(11) unsigned NOT NULL DEFAULT '0',
	PvpStatus             int64   `json:"pvpStatus" db:"pvp_status"`                          //`pvp_status` tinyint(11) unsigned NOT NULL DEFAULT '0',
	PvpKills              int64   `json:"pvpKills" db:"pvp_kills"`                            //`pvp_kills` int(11) unsigned NOT NULL DEFAULT '0',
	PvpDeaths             int64   `json:"pvpDeaths" db:"pvp_deaths"`                          //`pvp_deaths` int(11) unsigned NOT NULL DEFAULT '0',
	PvpCurrentPoints      int64   `json:"pvpCurrentPoints" db:"pvp_current_points"`           //`pvp_current_points` int(11) unsigned NOT NULL DEFAULT '0',
	PvpCareerPoints       int64   `json:"pvpCareerPoints" db:"pvp_career_points"`             //`pvp_career_points` int(11) unsigned NOT NULL DEFAULT '0',
	PvpBestKillStreak     int64   `json:"pvpBestKillStreak" db:"pvp_best_kill_streak"`        //`pvp_best_kill_streak` int(11) unsigned NOT NULL DEFAULT '0',
	PvpWorstDeathStreak   int64   `json:"pvpWorstDeathStreak" db:"pvp_worst_death_streak"`    //`pvp_worst_death_streak` int(11) unsigned NOT NULL DEFAULT '0',
	PvpCurrentKillStreak  int64   `json:"pvpCurrentKillStreak" db:"pvp_current_kill_streak"`  //`pvp_current_kill_streak` int(11) unsigned NOT NULL DEFAULT '0',
	Pvp2                  int64   `json:"pvp2" db:"pvp2"`                                     //`pvp2` int(11) unsigned NOT NULL DEFAULT '0',
	PvpType               int64   `json:"pvpType" db:"pvp_type"`                              //`pvp_type` int(11) unsigned NOT NULL DEFAULT '0',
	ShowHelm              int64   `json:"showHelm" db:"show_helm"`                            //`show_helm` int(11) unsigned NOT NULL DEFAULT '0',
	GroupAutoConsent      int64   `json:"groupAutoConsent" db:"group_auto_consent"`           //`group_auto_consent` tinyint(11) unsigned NOT NULL DEFAULT '0',
	RaidAutoConsent       int64   `json:"raidAutoConsent" db:"raid_auto_consent"`             //`raid_auto_consent` tinyint(11) unsigned NOT NULL DEFAULT '0',
	GuildAutoConsent      int64   `json:"guildAutoConsent" db:"guild_auto_consent"`           //`guild_auto_consent` tinyint(11) unsigned NOT NULL DEFAULT '0',
	LeadershipExpOn       int64   `json:"leadershipExpOn" db:"leadership_exp_on"`             //`leadership_exp_on` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Resttimer             int64   `json:"RestTimer" db:"RestTimer"`                           //`RestTimer` int(11) unsigned NOT NULL DEFAULT '0',
	AirRemaining          int64   `json:"airRemaining" db:"air_remaining"`                    //`air_remaining` int(11) unsigned NOT NULL DEFAULT '0',
	AutosplitEnabled      int64   `json:"autosplitEnabled" db:"autosplit_enabled"`            //`autosplit_enabled` int(11) unsigned NOT NULL DEFAULT '0',
	Lfp                   int64   `json:"lfp" db:"lfp"`                                       //`lfp` tinyint(1) unsigned NOT NULL DEFAULT '0',
	Lfg                   int64   `json:"lfg" db:"lfg"`                                       //`lfg` tinyint(1) unsigned NOT NULL DEFAULT '0',
	Mailkey               string  `json:"mailkey" db:"mailkey"`                               //`mailkey` char(16) NOT NULL DEFAULT '',
	Xtargets              int64   `json:"xtargets" db:"xtargets"`                             //`xtargets` tinyint(3) unsigned NOT NULL DEFAULT '5',
	Firstlogon            int64   `json:"firstlogon" db:"firstlogon"`                         //`firstlogon` tinyint(3) NOT NULL DEFAULT '0',
	EAaEffects            int64   `json:"eAaEffects" db:"e_aa_effects"`                       //`e_aa_effects` int(11) unsigned NOT NULL DEFAULT '0',
	EPercentToAa          int64   `json:"ePercentToAa" db:"e_percent_to_aa"`                  //`e_percent_to_aa` int(11) unsigned NOT NULL DEFAULT '0',
	EExpendedAaSpent      int64   `json:"eExpendedAaSpent" db:"e_expended_aa_spent"`          //`e_expended_aa_spent` int(11) unsigned NOT NULL DEFAULT '0',
	AaPointsSpentOld      int64   `json:"aaPointsSpentOld" db:"aa_points_spent_old"`          //`aa_points_spent_old` int(11) unsigned NOT NULL DEFAULT '0',
	AaPointsOld           int64   `json:"aaPointsOld" db:"aa_points_old"`                     //`aa_points_old` int(11) unsigned NOT NULL DEFAULT '0',
	ELastInvsnapshot      int64   `json:"eLastInvsnapshot" db:"e_last_invsnapshot"`           //`e_last_invsnapshot` int(11) unsigned NOT NULL DEFAULT '0',
}

func (c *Character) RaceIcon() string {
	switch c.Race {
	case 1:
		return "ra-player" //human
	case 2:
		return "ra-fox" //barbarian
	case 3:
		return "ra-book" //erudite
	case 4:
		return "ra-pine-tree" //woodelf
	case 5:
		return "ra-tesla" //helf
	case 6:
		return "ra-bleeding-eye" //delf
	case 7:
		return "ra-aware" //halfelf
	case 8:
		return "ra-beer" //dwarf
	case 9:
		return "ra-bird-mask" //troll
	case 10:
		return "ra-muscle-fat" //ogre
	case 11:
		return "ra-footprint" //halfling
	case 12:
		return "ra-gears" //gnome
	case 128:
		return "ra-gecko" //iksar
	case 130:
		return "ra-lion" //vahshir
	case 26: //may be wrong
		return "ra-water-drop" //froglok
	case 522:
		return "ra-wyvern" //drakkin
	}
	return "ra-help"
}

func (c *Character) ClassIcon() string {
	switch c.Class {
	case 1:
		return "ra-shield" //warrior
	case 2:
		return "ra-ankh" //cleric
	case 3:
		return "ra-fireball-sword" //paladin
	case 4:
		return "ra-arrow-cluster" //ranger
	case 5:
		return "ra-bat-sword" //shd
	case 6:
		return "ra-leaf" //druid
	case 7:
		return "ra-hand-emblem" //Monk
	case 8:
		return "ra-ocarina" //Bard
	case 9:
		return "ra-hood" //rogue
	case 10:
		return "ra-incense" //shaman
	case 11:
		return "ra-skull" //necro
	case 12:
		return "ra-fire" //wiz
	case 13:
		return "ra-burning-book" //magician
	case 14:
		return "ra-crystal-ball" //enchanter
	case 15:
		return "ra-pawprint" //beastlord
	case 16:
		return "ra-axe" //ber
	}
	return "ra-help"
}
func (c *Character) AA() int64 {
	return 0
}

func (c *Character) TotalHP() int64 {
	hp := c.CurHp
	return hp
}

func (c *Character) TotalMana() int64 {
	mana := c.Mana
	return mana
}

func (c *Character) ATK() int64 {
	atk := c.Dex
	return atk
}

func (c *Character) AC() int64 {
	ac := c.Agi
	return ac
}

func (c *Character) HPRegen() int64 {
	return 0
}

func (c *Character) ManaRegen() int64 {
	return 0
}

func (c *Character) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *Character) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "accountId":
		prop.Type = "integer"
		prop.Minimum = 1
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

func (c *Character) ClassName() string {
	switch c.Class {
	case 1:
		return "Warrior"
	case 2:
		return "Cleric"
	case 3:
		return "Paladin"
	case 4:
		return "Ranger"
	case 5:
		return "Shadowknight"
	case 6:
		return "Druid"
	case 7:
		return "Monk"
	case 8:
		return "Bard"
	case 9:
		return "Rogue"
	case 10:
		return "Shaman"
	case 11:
		return "Necromancer"
	case 12:
		return "Wizard"
	case 13:
		return "Magician"
	case 14:
		return "Enchanter"
	case 15:
		return "Beastlord"
	case 16:
		return "Berserker"
	}
	return "Unknown"
}

func (c *Character) RaceName() string {
	switch c.Race {
	case 1:
		return "Human"
	case 2:
		return "Barbarian"
	case 3:
		return "Erudite"
	case 4:
		return "Wood Elf"
	case 5:
		return "High Elf"
	case 6:
		return "Dark Elf"
	case 7:
		return "Half Elf"
	case 8:
		return "Dwarf"
	case 9:
		return "Troll"
	case 10:
		return "Ogre"
	case 11:
		return "Halfling"
	case 12:
		return "Gnome"
	case 13:
		return "Aviak"
	case 14:
		return "Werewolf"
	case 15:
		return "Brownie"
	case 16:
		return "Centaur"
	case 17:
		return "Golem"
	case 18:
		return "Giant"
	case 19:
		return "Trakanon"
	case 20:
		return "Venril Sathir"
	case 21:
		return "Evil Eye"
	case 22:
		return "Beetle"
	case 23:
		return "Kerran"
	case 24:
		return "Fish"
	case 25:
		return "Fairy"
	case 26:
		return "Froglok"
	case 27:
		return "Froglok"
	case 28:
		return "Fungusman"
	case 29:
		return "Gargoyle"
	case 30:
		return "Gasbag"
	case 31:
		return "Gelatinous Cube"
	case 32:
		return "Ghost"
	case 33:
		return "Ghoul"
	case 34:
		return "Bat"
	case 35:
		return "Eel"
	case 36:
		return "Rat"
	case 37:
		return "Snake"
	case 38:
		return "Spider"
	case 39:
		return "Gnoll"
	case 40:
		return "Goblin"
	case 41:
		return "Gorilla"
	case 42:
		return "Wolf"
	case 43:
		return "Bear"
	case 44:
		return "Guard"
	case 45:
		return "Demi Lich"
	case 46:
		return "Imp"
	case 47:
		return "Griffin"
	case 48:
		return "Kobold"
	case 49:
		return "Dragon"
	case 50:
		return "Lion"
	case 51:
		return "Lizard Man"
	case 52:
		return "Mimic"
	case 53:
		return "Minotaur"
	case 54:
		return "Orc"
	case 55:
		return "Beggar"
	case 56:
		return "Pixie"
	case 57:
		return "Drachnid"
	case 58:
		return "Solusek Ro"
	case 59:
		return "Goblin"
	case 60:
		return "Skeleton"
	case 61:
		return "Shark"
	case 62:
		return "Tunare"
	case 63:
		return "Tiger"
	case 64:
		return "Treant"
	case 65:
		return "Vampire"
	case 66:
		return "Rallos Zek"
	case 67:
		return "Human"
	case 68:
		return "Tentacle Terror"
	case 69:
		return "Will-O-Wisp"
	case 70:
		return "Zombie"
	case 71:
		return "Human"
	case 72:
		return "Ship"
	case 73:
		return "Launch"
	case 74:
		return "Piranha"
	case 75:
		return "Elemental"
	case 76:
		return "Puma"
	case 77:
		return "Dark Elf"
	case 78:
		return "Erudite"
	case 79:
		return "Bixie"
	case 80:
		return "Reanimated Hand"
	case 81:
		return "Halfling"
	case 82:
		return "Scarecrow"
	case 83:
		return "Skunk"
	case 84:
		return "Snake Elemental"
	case 85:
		return "Spectre"
	case 86:
		return "Sphinx"
	case 87:
		return "Armadillo"
	case 88:
		return "Clockwork Gnome"
	case 89:
		return "Drake"
	case 90:
		return "Barbarian"
	case 91:
		return "Alligator"
	case 92:
		return "Troll"
	case 93:
		return "Ogre"
	case 94:
		return "Dwarf"
	case 95:
		return "Cazic Thule"
	case 96:
		return "Cockatrice"
	case 97:
		return "Daisy Man"
	case 98:
		return "Vampire"
	case 99:
		return "Amygdalan"
	case 100:
		return "Dervish"
	case 101:
		return "Efreeti"
	case 102:
		return "Tadpole"
	case 103:
		return "Kedge"
	case 104:
		return "Leech"
	case 105:
		return "Swordfish"
	case 106:
		return "Guard"
	case 107:
		return "Mammoth"
	case 108:
		return "Eye"
	case 109:
		return "Wasp"
	case 110:
		return "Mermaid"
	case 111:
		return "Harpy"
	case 112:
		return "Guard"
	case 113:
		return "Drixie"
	case 114:
		return "Ghost Ship"
	case 115:
		return "Clam"
	case 116:
		return "Seahorse"
	case 117:
		return "Ghost"
	case 118:
		return "Ghost"
	case 119:
		return "Sabertooth"
	case 120:
		return "Wolf"
	case 121:
		return "Gorgon"
	case 122:
		return "Dragon"
	case 123:
		return "Innoruuk"
	case 124:
		return "Unicorn"
	case 125:
		return "Pegasus"
	case 126:
		return "Djinn"
	case 127:
		return "Invisible Man"
	case 128:
		return "Iksar"
	case 129:
		return "Scorpion"
	case 130:
		return "Vah Shir"
	case 131:
		return "Sarnak"
	case 132:
		return "Draglock"
	case 133:
		return "Drolvarg"
	case 134:
		return "Mosquito"
	case 135:
		return "Rhinoceros"
	case 136:
		return "Xalgoz"
	case 137:
		return "Goblin"
	case 138:
		return "Yeti"
	case 139:
		return "Iksar"
	case 140:
		return "Giant"
	case 141:
		return "Boat"
	case 142:
		return "Object"
	case 143:
		return "Tree"
	case 144:
		return "Burynai"
	case 145:
		return "Goo"
	case 146:
		return "Sarnak Spirit"
	case 147:
		return "Iksar Spirit"
	case 148:
		return "Fish"
	case 149:
		return "Scorpion"
	case 150:
		return "Erollisi"
	case 151:
		return "Tribunal"
	case 152:
		return "Bertoxxulous"
	case 153:
		return "Bristlebane"
	case 154:
		return "Fay Drake"
	case 155:
		return "Undead Sarnak"
	case 156:
		return "Ratman"
	case 157:
		return "Wyvern"
	case 158:
		return "Wurm"
	case 159:
		return "Devourer"
	case 160:
		return "Iksar Golem"
	case 161:
		return "Undead Iksar"
	case 162:
		return "ManEating Plant"
	case 163:
		return "Raptor"
	case 164:
		return "Sarnak Golem"
	case 165:
		return "Dragon"
	case 166:
		return "Animated Hand"
	case 167:
		return "Succulent"
	case 168:
		return "Holgresh"
	case 169:
		return "Brontotherium"
	case 170:
		return "Snow Dervish"
	case 171:
		return "Dire Wolf"
	case 172:
		return "Manticore"
	case 173:
		return "Totem"
	case 174:
		return "Ice Spectre"
	case 175:
		return "Enchanted Armor"
	case 176:
		return "Snow Rabbit"
	case 177:
		return "Walrus"
	case 178:
		return "Geonid"
	case 181:
		return "Yakkar"
	case 182:
		return "Faun"
	case 183:
		return "Coldain"
	case 184:
		return "Dragon"
	case 185:
		return "Hag"
	case 186:
		return "Hippogriff"
	case 187:
		return "Siren"
	case 188:
		return "Giant"
	case 189:
		return "Giant"
	case 190:
		return "Othmir"
	case 191:
		return "Ulthork"
	case 192:
		return "Dragon"
	case 193:
		return "Abhorrent"
	case 194:
		return "Sea Turtle"
	case 195:
		return "Dragon"
	case 196:
		return "Dragon"
	case 197:
		return "Ronnie Test"
	case 198:
		return "Dragon"
	case 199:
		return "Shik'Nar"
	case 200:
		return "Rockhopper"
	case 201:
		return "Underbulk"
	case 202:
		return "Grimling"
	case 203:
		return "Worm"
	case 204:
		return "Evan Test"
	case 205:
		return "Shadel"
	case 206:
		return "Owlbear"
	case 207:
		return "Rhino Beetle"
	case 208:
		return "Vampire"
	case 209:
		return "Earth Elemental"
	case 210:
		return "Air Elemental"
	case 211:
		return "Water Elemental"
	case 212:
		return "Fire Elemental"
	case 213:
		return "Wetfang Minnow"
	case 214:
		return "Thought Horror"
	case 215:
		return "Tegi"
	case 216:
		return "Horse"
	case 217:
		return "Shissar"
	case 218:
		return "Fungal Fiend"
	case 219:
		return "Vampire"
	case 220:
		return "Stonegrabber"
	case 221:
		return "Scarlet Cheetah"
	case 222:
		return "Zelniak"
	case 223:
		return "Lightcrawler"
	case 224:
		return "Shade"
	case 225:
		return "Sunflower"
	case 226:
		return "Sun Revenant"
	case 227:
		return "Shrieker"
	case 228:
		return "Galorian"
	case 229:
		return "Netherbian"
	case 230:
		return "Akheva"
	case 231:
		return "Grieg Veneficus"
	case 232:
		return "Sonic Wolf"
	case 233:
		return "Ground Shaker"
	case 234:
		return "Vah Shir Skeleton"
	case 235:
		return "Wretch"
	case 236:
		return "Seru"
	case 237:
		return "Recuso"
	case 238:
		return "Vah Shir"
	case 239:
		return "Guard"
	case 240:
		return "Teleport Man"
	case 241:
		return "Werewolf"
	case 242:
		return "Nymph"
	case 243:
		return "Dryad"
	case 244:
		return "Treant"
	case 245:
		return "Fly"
	case 246:
		return "Tarew Marr"
	case 247:
		return "Solusek Ro"
	case 248:
		return "Clockwork Golem"
	case 249:
		return "Clockwork Brain"
	case 250:
		return "Banshee"
	case 251:
		return "Guard of Justice"
	case 252:
		return "Mini POM"
	case 253:
		return "Diseased Fiend"
	case 254:
		return "Solusek Ro Guard"
	case 255:
		return "Bertoxxulous"
	case 256:
		return "The Tribunal"
	case 257:
		return "Terris Thule"
	case 258:
		return "Vegerog"
	case 259:
		return "Crocodile"
	case 260:
		return "Bat"
	case 261:
		return "Hraquis"
	case 262:
		return "Tranquilion"
	case 263:
		return "Tin Soldier"
	case 264:
		return "Nightmare Wraith"
	case 265:
		return "Malarian"
	case 266:
		return "Knight of Pestilence"
	case 267:
		return "Lepertoloth"
	case 268:
		return "Bubonian"
	case 269:
		return "Bubonian Underling"
	case 270:
		return "Pusling"
	case 271:
		return "Water Mephit"
	case 272:
		return "Stormrider"
	case 273:
		return "Junk Beast"
	case 274:
		return "Broken Clockwork"
	case 275:
		return "Giant Clockwork"
	case 276:
		return "Clockwork Beetle"
	case 277:
		return "Nightmare Goblin"
	case 278:
		return "Karana"
	case 279:
		return "Blood Raven"
	case 280:
		return "Nightmare Gargoyle"
	case 281:
		return "Mouth of Insanity"
	case 282:
		return "Skeletal Horse"
	case 283:
		return "Saryrn"
	case 284:
		return "Fennin Ro"
	case 285:
		return "Tormentor"
	case 286:
		return "Soul Devourer"
	case 287:
		return "Nightmare"
	case 288:
		return "Rallos Zek"
	case 289:
		return "Vallon Zek"
	case 290:
		return "Tallon Zek"
	case 291:
		return "Air Mephit"
	case 292:
		return "Earth Mephit"
	case 293:
		return "Fire Mephit"
	case 294:
		return "Nightmare Mephit"
	case 295:
		return "Zebuxoruk"
	case 296:
		return "Mithaniel Marr"
	case 297:
		return "Undead Knight"
	case 298:
		return "The Rathe"
	case 299:
		return "Xegony"
	case 300:
		return "Fiend"
	case 301:
		return "Test Object"
	case 302:
		return "Crab"
	case 303:
		return "Phoenix"
	case 304:
		return "Dragon"
	case 305:
		return "Bear"
	case 306:
		return "Giant"
	case 307:
		return "Giant"
	case 308:
		return "Giant"
	case 309:
		return "Giant"
	case 310:
		return "Giant"
	case 311:
		return "Giant"
	case 312:
		return "Giant"
	case 313:
		return "War Wraith"
	case 314:
		return "Wrulon"
	case 315:
		return "Kraken"
	case 316:
		return "Poison Frog"
	case 317:
		return "Nilborien"
	case 318:
		return "Valorian"
	case 319:
		return "War Boar"
	case 320:
		return "Efreeti"
	case 321:
		return "War Boar"
	case 322:
		return "Valorian"
	case 323:
		return "Animated Armor"
	case 324:
		return "Undead Footman"
	case 325:
		return "Rallos Zek Minion"
	case 326:
		return "Arachnid"
	case 327:
		return "Crystal Spider"
	case 328:
		return "Zebuxoruk's Cage"
	case 329:
		return "Bastion of Thunder Portal"
	case 330:
		return "Froglok"
	case 331:
		return "Troll"
	case 332:
		return "Troll"
	case 333:
		return "Troll"
	case 334:
		return "Ghost"
	case 335:
		return "Pirate"
	case 336:
		return "Pirate"
	case 337:
		return "Pirate"
	case 338:
		return "Pirate"
	case 339:
		return "Pirate"
	case 340:
		return "Pirate"
	case 341:
		return "Pirate"
	case 342:
		return "Pirate"
	case 343:
		return "Frog"
	case 344:
		return "Troll Zombie"
	case 345:
		return "Luggald"
	case 346:
		return "Luggald"
	case 347:
		return "Luggalds"
	case 348:
		return "Drogmore"
	case 349:
		return "Froglok Skeleton"
	case 350:
		return "Undead Froglok"
	case 351:
		return "Knight of Hate"
	case 352:
		return "Arcanist of Hate"
	case 353:
		return "Veksar"
	case 354:
		return "Veksar"
	case 355:
		return "Veksar"
	case 356:
		return "Chokidai"
	case 357:
		return "Undead Chokidai"
	case 358:
		return "Undead Veksar"
	case 359:
		return "Vampire"
	case 360:
		return "Vampire"
	case 361:
		return "Rujarkian Orc"
	case 362:
		return "Bone Golem"
	case 363:
		return "Synarcana"
	case 364:
		return "Sand Elf"
	case 365:
		return "Vampire"
	case 366:
		return "Rujarkian Orc"
	case 367:
		return "Skeleton"
	case 368:
		return "Mummy"
	case 369:
		return "Goblin"
	case 370:
		return "Insect"
	case 371:
		return "Froglok Ghost"
	case 372:
		return "Dervish"
	case 373:
		return "Shade"
	case 374:
		return "Golem"
	case 375:
		return "Evil Eye"
	case 376:
		return "Box"
	case 377:
		return "Barrel"
	case 378:
		return "Chest"
	case 379:
		return "Vase"
	case 380:
		return "Table"
	case 381:
		return "Weapon Rack"
	case 382:
		return "Coffin"
	case 383:
		return "Bones"
	case 384:
		return "Jokester"
	case 385:
		return "Nihil"
	case 386:
		return "Trusik"
	case 387:
		return "Stone Worker"
	case 388:
		return "Hynid"
	case 389:
		return "Turepta"
	case 390:
		return "Cragbeast"
	case 391:
		return "Stonemite"
	case 392:
		return "Ukun"
	case 393:
		return "Ixt"
	case 394:
		return "Ikaav"
	case 395:
		return "Aneuk"
	case 396:
		return "Kyv"
	case 397:
		return "Noc"
	case 398:
		return "Ra`tuk"
	case 399:
		return "Taneth"
	case 400:
		return "Huvul"
	case 401:
		return "Mutna"
	case 402:
		return "Mastruq"
	case 403:
		return "Taelosian"
	case 404:
		return "Discord Ship"
	case 405:
		return "Stone Worker"
	case 406:
		return "Mata Muram"
	case 407:
		return "Lightning Warrior"
	case 408:
		return "Succubus"
	case 409:
		return "Bazu"
	case 410:
		return "Feran"
	case 411:
		return "Pyrilen"
	case 412:
		return "Chimera"
	case 413:
		return "Dragorn"
	case 414:
		return "Murkglider"
	case 415:
		return "Rat"
	case 416:
		return "Bat"
	case 417:
		return "Gelidran"
	case 418:
		return "Discordling"
	case 419:
		return "Girplan"
	case 420:
		return "Minotaur"
	case 421:
		return "Dragorn Box"
	case 422:
		return "Runed Orb"
	case 423:
		return "Dragon Bones"
	case 424:
		return "Muramite Armor Pile"
	case 425:
		return "Crystal Shard"
	case 426:
		return "Portal"
	case 427:
		return "Coin Purse"
	case 428:
		return "Rock Pile"
	case 429:
		return "Murkglider Egg Sack"
	case 430:
		return "Drake"
	case 431:
		return "Dervish"
	case 432:
		return "Drake"
	case 433:
		return "Goblin"
	case 434:
		return "Kirin"
	case 435:
		return "Dragon"
	case 436:
		return "Basilisk"
	case 437:
		return "Dragon"
	case 438:
		return "Dragon"
	case 439:
		return "Puma"
	case 440:
		return "Spider"
	case 441:
		return "Spider Queen"
	case 442:
		return "Animated Statue"
	case 445:
		return "Dragon Egg"
	case 446:
		return "Dragon Statue"
	case 447:
		return "Lava Rock"
	case 448:
		return "Animated Statue"
	case 449:
		return "Spider Egg Sack"
	case 450:
		return "Lava Spider"
	case 451:
		return "Lava Spider Queen"
	case 452:
		return "Dragon"
	case 453:
		return "Giant"
	case 454:
		return "Werewolf"
	case 455:
		return "Kobold"
	case 456:
		return "Sporali"
	case 457:
		return "Gnomework"
	case 458:
		return "Orc"
	case 459:
		return "Corathus"
	case 460:
		return "Coral"
	case 461:
		return "Drachnid"
	case 462:
		return "Drachnid Cocoon"
	case 463:
		return "Fungus Patch"
	case 464:
		return "Gargoyle"
	case 465:
		return "Witheran"
	case 466:
		return "Dark Lord"
	case 467:
		return "Shiliskin"
	case 468:
		return "Snake"
	case 469:
		return "Evil Eye"
	case 470:
		return "Minotaur"
	case 471:
		return "Zombie"
	case 472:
		return "Clockwork Boar"
	case 473:
		return "Fairy"
	case 474:
		return "Witheran"
	case 475:
		return "Air Elemental"
	case 476:
		return "Earth Elemental"
	case 477:
		return "Fire Elemental"
	case 478:
		return "Water Elemental"
	case 479:
		return "Alligator"
	case 480:
		return "Bear"
	case 481:
		return "Scaled Wolf"
	case 482:
		return "Wolf"
	case 483:
		return "Spirit Wolf"
	case 484:
		return "Skeleton"
	case 485:
		return "Spectre"
	case 486:
		return "Bolvirk"
	case 487:
		return "Banshee"
	case 488:
		return "Banshee"
	case 489:
		return "Elddar"
	case 490:
		return "Forest Giant"
	case 491:
		return "Bone Golem"
	case 492:
		return "Horse"
	case 493:
		return "Pegasus"
	case 494:
		return "Shambling Mound"
	case 495:
		return "Scrykin"
	case 496:
		return "Treant"
	case 497:
		return "Vampire"
	case 498:
		return "Ayonae Ro"
	case 499:
		return "Sullon Zek"
	case 500:
		return "Banner"
	case 501:
		return "Flag"
	case 502:
		return "Rowboat"
	case 503:
		return "Bear Trap"
	case 504:
		return "Clockwork Bomb"
	case 505:
		return "Dynamite Keg"
	case 506:
		return "Pressure Plate"
	case 507:
		return "Puffer Spore"
	case 508:
		return "Stone Ring"
	case 509:
		return "Root Tentacle"
	case 510:
		return "Runic Symbol"
	case 511:
		return "Saltpetter Bomb"
	case 512:
		return "Floating Skull"
	case 513:
		return "Spike Trap"
	case 514:
		return "Totem"
	case 515:
		return "Web"
	case 516:
		return "Wicker Basket"
	case 517:
		return "Nightmare/Unicorn"
	case 518:
		return "Horse"
	case 519:
		return "Nightmare/Unicorn"
	case 520:
		return "Bixie"
	case 521:
		return "Centaur"
	case 522:
		return "Drakkin"
	case 523:
		return "Giant"
	case 524:
		return "Gnoll"
	case 525:
		return "Griffin"
	case 526:
		return "Giant Shade"
	case 527:
		return "Harpy"
	case 528:
		return "Mammoth"
	case 529:
		return "Satyr"
	case 530:
		return "Dragon"
	case 531:
		return "Dragon"
	case 532:
		return "Dyn'Leth"
	case 533:
		return "Boat"
	case 534:
		return "Weapon Rack"
	case 535:
		return "Armor Rack"
	case 536:
		return "Honey Pot"
	case 537:
		return "Jum Jum Bucket"
	case 538:
		return "Toolbox"
	case 539:
		return "Stone Jug"
	case 540:
		return "Small Plant"
	case 541:
		return "Medium Plant"
	case 542:
		return "Tall Plant"
	case 543:
		return "Wine Cask"
	case 544:
		return "Elven Boat"
	case 545:
		return "Gnomish Boat"
	case 546:
		return "Barrel Barge Ship"
	case 547:
		return "Goo"
	case 548:
		return "Goo"
	case 549:
		return "Goo"
	case 550:
		return "Merchant Ship"
	case 551:
		return "Pirate Ship"
	case 552:
		return "Ghost Ship"
	case 553:
		return "Banner"
	case 554:
		return "Banner"
	case 555:
		return "Banner"
	case 556:
		return "Banner"
	case 557:
		return "Banner"
	case 558:
		return "Aviak"
	case 559:
		return "Beetle"
	case 560:
		return "Gorilla"
	case 561:
		return "Kedge"
	case 562:
		return "Kerran"
	case 563:
		return "Shissar"
	case 564:
		return "Siren"
	case 565:
		return "Sphinx"
	case 566:
		return "Human"
	case 567:
		return "Campfire"
	case 568:
		return "Brownie"
	case 569:
		return "Dragon"
	case 570:
		return "Exoskeleton"
	case 571:
		return "Ghoul"
	case 572:
		return "Clockwork Guardian"
	case 573:
		return "Mantrap"
	case 574:
		return "Minotaur"
	case 575:
		return "Scarecrow"
	case 576:
		return "Shade"
	case 577:
		return "Rotocopter"
	case 578:
		return "Tentacle Terror"
	case 579:
		return "Wereorc"
	case 580:
		return "Worg"
	case 581:
		return "Wyvern"
	case 582:
		return "Chimera"
	case 583:
		return "Kirin"
	case 584:
		return "Puma"
	case 585:
		return "Boulder"
	case 586:
		return "Banner"
	case 587:
		return "Elven Ghost"
	case 588:
		return "Human Ghost"
	case 589:
		return "Chest"
	case 590:
		return "Chest"
	case 591:
		return "Crystal"
	case 592:
		return "Coffin"
	case 593:
		return "Guardian CPU"
	case 594:
		return "Worg"
	case 595:
		return "Mansion"
	case 596:
		return "Floating Island"
	case 597:
		return "Cragslither"
	case 598:
		return "Wrulon"
	case 599:
		return "Spell Particle 1"
	case 600:
		return "Invisible Man of Zomm"
	case 601:
		return "Robocopter of Zomm"
	case 602:
		return "Burynai"
	case 603:
		return "Frog"
	case 604:
		return "Dracolich"
	case 605:
		return "Iksar Ghost"
	case 606:
		return "Iksar Skeleton"
	case 607:
		return "Mephit"
	case 608:
		return "Muddite"
	case 609:
		return "Raptor"
	case 610:
		return "Sarnak"
	case 611:
		return "Scorpion"
	case 612:
		return "Tsetsian"
	case 613:
		return "Wurm"
	case 614:
		return "Nekhon"
	case 615:
		return "Hydra Crystal"
	case 616:
		return "Crystal Sphere"
	case 617:
		return "Gnoll"
	case 618:
		return "Sokokar"
	case 619:
		return "Stone Pylon"
	case 620:
		return "Demon Vulture"
	case 621:
		return "Wagon"
	case 622:
		return "God of Discord"
	case 623:
		return "Feran Mount"
	case 624:
		return "Ogre NPC Male"
	case 625:
		return "Sokokar Mount"
	case 626:
		return "Giant"
	case 627:
		return "Sokokar"
	case 628:
		return "10th Anniversary Banner"
	case 629:
		return "10th Anniversary Cake"
	case 630:
		return "Wine Cask"
	case 631:
		return "Hydra Mount"
	case 632:
		return "Hydra NPC"
	case 633:
		return "Wedding Flowers"
	case 634:
		return "Wedding Arbor"
	case 635:
		return "Wedding Altar"
	case 636:
		return "Powder Keg"
	case 637:
		return "Apexus"
	case 638:
		return "Bellikos"
	case 639:
		return "Brell's First Creation"
	case 640:
		return "Brell"
	case 641:
		return "Crystalskin Ambuloid"
	case 642:
		return "Cliknar Queen"
	case 643:
		return "Cliknar Soldier"
	case 644:
		return "Cliknar Worker"
	case 645:
		return "Coldain"
	case 646:
		return "Coldain"
	case 647:
		return "Crystalskin Sessiloid"
	case 648:
		return "Genari"
	case 649:
		return "Gigyn"
	case 650:
		return "Greken Young Adult"
	case 651:
		return "Greken Young"
	case 652:
		return "Cliknar Mount"
	case 653:
		return "Telmira"
	case 654:
		return "Spider Mount"
	case 655:
		return "Bear Mount"
	case 656:
		return "Rat Mount Mystery Race"
	case 657:
		return "Sessiloid Mount"
	case 658:
		return "Morell Thule"
	case 659:
		return "Marionette"
	case 660:
		return "Book Dervish"
	case 661:
		return "Topiary Lion"
	case 662:
		return "Rotdog"
	case 663:
		return "Amygdalan"
	case 664:
		return "Sandman"
	case 665:
		return "Grandfather Clock"
	case 666:
		return "Gingerbread Man"
	case 667:
		return "Royal Guard"
	case 668:
		return "Rabbit"
	case 669:
		return "Blind Dreamer"
	case 670:
		return "Cazic Thule"
	case 671:
		return "Topiary Lion Mount"
	case 672:
		return "Rot Dog Mount"
	case 673:
		return "Goral Mount"
	case 674:
		return "Selyrah Mount"
	case 675:
		return "Sclera Mount"
	case 676:
		return "Braxi Mount"
	case 677:
		return "Kangon Mount"
	case 678:
		return "Erudite"
	case 679:
		return "Wurm Mount"
	case 680:
		return "Raptor Mount"
	case 681:
		return "Invisible Man"
	case 682:
		return "Whirligig"
	case 683:
		return "Gnomish Balloon"
	case 684:
		return "Gnomish Rocket Pack"
	case 685:
		return "Gnomish Hovering Transport"
	case 686:
		return "Selyrah"
	case 687:
		return "Goral"
	case 688:
		return "Braxi"
	case 689:
		return "Kangon"
	case 690:
		return "Invisible Man"
	case 691:
		return "Floating Tower"
	case 692:
		return "Explosive Cart"
	case 693:
		return "Blimp Ship"
	case 694:
		return "Tumbleweed"
	case 695:
		return "Alaran"
	case 696:
		return "Swinetor"
	case 697:
		return "Triumvirate"
	case 698:
		return "Hadal"
	case 699:
		return "Hovering Platform"
	case 700:
		return "Parasitic Scavenger"
	case 701:
		return "Grendlaen"
	case 702:
		return "Ship in a Bottle"
	case 703:
		return "Alaran Sentry Stone"
	case 704:
		return "Dervish"
	case 705:
		return "Regeneration Pool"
	case 706:
		return "Teleportation Stand"
	case 707:
		return "Relic Case"
	case 708:
		return "Alaran Ghost"
	case 709:
		return "Skystrider"
	case 710:
		return "Water Spout"
	case 711:
		return "Aviak Pull Along"
	case 712:
		return "Gelatinous Cube"
	case 713:
		return "Cat"
	case 714:
		return "Elk Head"
	case 715:
		return "Holgresh"
	case 716:
		return "Beetle"
	case 717:
		return "Vine Maw"
	case 718:
		return "Ratman"
	case 719:
		return "Fallen Knight"
	case 720:
		return "Flying Carpet"
	case 721:
		return "Carrier Hand"
	case 722:
		return "Akheva"
	case 723:
		return "Servant of Shadow"
	case 724:
		return "Luclin"
	}
	return "Unknown"
}
