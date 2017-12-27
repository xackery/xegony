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
