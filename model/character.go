package model

import ()

// Characters is an array of character
// swagger:model
type Characters []*Character

//Character holds data about players inside Everquest, it primarily uses character_data table
// swagger:model
type Character struct {
	Zone  *Zone  `json:"zone,omitempty"`
	Race  *Race  `json:"race,omitempty"`
	Class *Class `json:"class,omitempty"`

	/*Account   *Account `json:"account,omitempty"`
	Base      *Base    `json:"base,omitempty"`
	Inventory []*Item  `json:"inventory,omitempty"`
	Zone      *Zone    `json:"zone,omitempty"`

	AASpent   int64    `json:"aaSpent,omitempty"`
	TotalHP   int64    `json:"totalHP,omitempty"`
	TotalMana int64    `json:"totalMana,omitempty"`*/

	ID                    int64   `json:"ID,omitempty" db:"id"`                                         //`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	AccountID             int64   `json:"accountID,omitempty" db:"account_id"`                          //`account_id` int(11) NOT NULL DEFAULT '0',
	Name                  string  `json:"name,omitempty" db:"name"`                                     //`name` varchar(64) NOT NULL DEFAULT '',
	LastName              string  `json:"lastName,omitempty" db:"last_name"`                            //`last_name` varchar(64) NOT NULL DEFAULT '',
	Title                 string  `json:"title,omitempty" db:"title"`                                   //`title` varchar(32) NOT NULL DEFAULT '',
	Suffix                string  `json:"suffix,omitempty" db:"suffix"`                                 //`suffix` varchar(32) NOT NULL DEFAULT '',
	ZoneID                int64   `json:"zoneID,omitempty" db:"zone_id"`                                //`zone_id` int(11) unsigned NOT NULL DEFAULT '0',
	ZoneInstance          int64   `json:"zoneInstance,omitempty" db:"zone_instance"`                    //`zone_instance` int(11) unsigned NOT NULL DEFAULT '0',
	Y                     float64 `json:"y,omitempty" db:"y"`                                           //`y` float NOT NULL DEFAULT '0',
	X                     float64 `json:"x,omitempty" db:"x"`                                           //`x` float NOT NULL DEFAULT '0',
	Z                     float64 `json:"z,omitempty" db:"z"`                                           //`z` float NOT NULL DEFAULT '0',
	Heading               float64 `json:"heading,omitempty" db:"heading"`                               //`heading` float NOT NULL DEFAULT '0',
	Gender                int64   `json:"gender,omitempty" db:"gender"`                                 //`gender` tinyint(11) unsigned NOT NULL DEFAULT '0',
	RaceID                int64   `json:"raceID,omitempty" db:"race"`                                   //`race` smallint(11) unsigned NOT NULL DEFAULT '0',
	ClassID               int64   `json:"classID,omitempty" db:"class"`                                 //`class` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Level                 int64   `json:"level,omitempty" db:"level"`                                   //`level` int(11) unsigned NOT NULL DEFAULT '0',
	Deity                 int64   `json:"deity,omitempty" db:"deity"`                                   //`deity` int(11) unsigned NOT NULL DEFAULT '0',
	Birthday              int64   `json:"birthday,omitempty" db:"birthday"`                             //`birthday` int(11) unsigned NOT NULL DEFAULT '0',
	LastLogin             int64   `json:"lastLogin,omitempty" db:"last_login"`                          //`last_login` int(11) unsigned NOT NULL DEFAULT '0',
	TimePlayed            int64   `json:"timePlayed,omitempty" db:"time_played"`                        //`time_played` int(11) unsigned NOT NULL DEFAULT '0',
	Level2                int64   `json:"level2,omitempty" db:"level2"`                                 //`level2` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Anon                  int64   `json:"anon,omitempty" db:"anon"`                                     //`anon` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Gm                    int64   `json:"gm,omitempty" db:"gm"`                                         //`gm` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Face                  int64   `json:"face,omitempty" db:"face"`                                     //`face` int(11) unsigned NOT NULL DEFAULT '0',
	HairColor             int64   `json:"hairColor,omitempty" db:"hair_color"`                          //`hair_color` tinyint(11) unsigned NOT NULL DEFAULT '0',
	HairStyle             int64   `json:"hairStyle,omitempty" db:"hair_style"`                          //`hair_style` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Beard                 int64   `json:"beard,omitempty" db:"beard"`                                   //`beard` tinyint(11) unsigned NOT NULL DEFAULT '0',
	BeardColor            int64   `json:"beardColor,omitempty" db:"beard_color"`                        //`beard_color` tinyint(11) unsigned NOT NULL DEFAULT '0',
	EyeColor1             int64   `json:"eyeColor1,omitempty" db:"eye_color_1"`                         //`eye_color_1` tinyint(11) unsigned NOT NULL DEFAULT '0',
	EyeColor2             int64   `json:"eyeColor2,omitempty" db:"eye_color_2"`                         //`eye_color_2` tinyint(11) unsigned NOT NULL DEFAULT '0',
	DrakkinHeritage       int64   `json:"drakkinHeritage,omitempty" db:"drakkin_heritage"`              //`drakkin_heritage` int(11) unsigned NOT NULL DEFAULT '0',
	DrakkinTattoo         int64   `json:"drakkinTattoo,omitempty" db:"drakkin_tattoo"`                  //`drakkin_tattoo` int(11) unsigned NOT NULL DEFAULT '0',
	DrakkinDetails        int64   `json:"drakkinDetails,omitempty" db:"drakkin_details"`                //`drakkin_details` int(11) unsigned NOT NULL DEFAULT '0',
	AbilityTimeSeconds    int64   `json:"abilityTimeSeconds,omitempty" db:"ability_time_seconds"`       //`ability_time_seconds` tinyint(11) unsigned NOT NULL DEFAULT '0',
	AbilityNumber         int64   `json:"abilityNumber,omitempty" db:"ability_number"`                  //`ability_number` tinyint(11) unsigned NOT NULL DEFAULT '0',
	AbilityTimeMinutes    int64   `json:"abilityTimeMinutes,omitempty" db:"ability_time_minutes"`       //`ability_time_minutes` tinyint(11) unsigned NOT NULL DEFAULT '0',
	AbilityTimeHours      int64   `json:"abilityTimeHours,omitempty" db:"ability_time_hours"`           //`ability_time_hours` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Exp                   int64   `json:"exp,omitempty" db:"exp"`                                       //`exp` int(11) unsigned NOT NULL DEFAULT '0',
	AaPointsSpent         int64   `json:"aaPointsSpent,omitempty" db:"aa_points_spent"`                 //`aa_points_spent` int(11) unsigned NOT NULL DEFAULT '0',
	AaExp                 int64   `json:"aaExp,omitempty" db:"aa_exp"`                                  //`aa_exp` int(11) unsigned NOT NULL DEFAULT '0',
	AaPoints              int64   `json:"aaPoints,omitempty" db:"aa_points"`                            //`aa_points` int(11) unsigned NOT NULL DEFAULT '0',
	GroupLeadershipExp    int64   `json:"groupLeadershipExp,omitempty" db:"group_leadership_exp"`       //`group_leadership_exp` int(11) unsigned NOT NULL DEFAULT '0',
	RaidLeadershipExp     int64   `json:"raidLeadershipExp,omitempty" db:"raid_leadership_exp"`         //`raid_leadership_exp` int(11) unsigned NOT NULL DEFAULT '0',
	GroupLeadershipPoints int64   `json:"groupLeadershipPoints,omitempty" db:"group_leadership_points"` //`group_leadership_points` int(11) unsigned NOT NULL DEFAULT '0',
	RaidLeadershipPoints  int64   `json:"raidLeadershipPoints,omitempty" db:"raid_leadership_points"`   //`raid_leadership_points` int(11) unsigned NOT NULL DEFAULT '0',
	Points                int64   `json:"points,omitempty" db:"points"`                                 //`points` int(11) unsigned NOT NULL DEFAULT '0',
	CurHp                 int64   `json:"curHp,omitempty" db:"cur_hp"`                                  //`cur_hp` int(11) unsigned NOT NULL DEFAULT '0',
	Mana                  int64   `json:"mana,omitempty" db:"mana"`                                     //`mana` int(11) unsigned NOT NULL DEFAULT '0',
	Endurance             int64   `json:"endurance,omitempty" db:"endurance"`                           //`endurance` int(11) unsigned NOT NULL DEFAULT '0',
	Intoxication          int64   `json:"intoxication,omitempty" db:"intoxication"`                     //`intoxication` int(11) unsigned NOT NULL DEFAULT '0',
	Str                   int64   `json:"str,omitempty" db:"str"`                                       //`str` int(11) unsigned NOT NULL DEFAULT '0',
	Sta                   int64   `json:"sta,omitempty" db:"sta"`                                       //`sta` int(11) unsigned NOT NULL DEFAULT '0',
	Cha                   int64   `json:"cha,omitempty" db:"cha"`                                       //`cha` int(11) unsigned NOT NULL DEFAULT '0',
	Dex                   int64   `json:"dex,omitempty" db:"dex"`                                       //`dex` int(11) unsigned NOT NULL DEFAULT '0',
	Int                   int64   `json:"int,omitempty" db:"int"`                                       //`int` int(11) unsigned NOT NULL DEFAULT '0',
	Agi                   int64   `json:"agi,omitempty" db:"agi"`                                       //`agi` int(11) unsigned NOT NULL DEFAULT '0',
	Wis                   int64   `json:"wis,omitempty" db:"wis"`                                       //`wis` int(11) unsigned NOT NULL DEFAULT '0',
	ZoneChangeCount       int64   `json:"zoneChangeCount,omitempty" db:"zone_change_count"`             //`zone_change_count` int(11) unsigned NOT NULL DEFAULT '0',
	Toxicity              int64   `json:"toxicity,omitempty" db:"toxicity"`                             //`toxicity` int(11) unsigned NOT NULL DEFAULT '0',
	HungerLevel           int64   `json:"hungerLevel,omitempty" db:"hunger_level"`                      //`hunger_level` int(11) unsigned NOT NULL DEFAULT '0',
	ThirstLevel           int64   `json:"thirstLevel,omitempty" db:"thirst_level"`                      //`thirst_level` int(11) unsigned NOT NULL DEFAULT '0',
	AbilityUp             int64   `json:"abilityUp,omitempty" db:"ability_up"`                          //`ability_up` int(11) unsigned NOT NULL DEFAULT '0',
	LdonPointsGuk         int64   `json:"ldonPointsGuk,omitempty" db:"ldon_points_guk"`                 //`ldon_points_guk` int(11) unsigned NOT NULL DEFAULT '0',
	LdonPointsMir         int64   `json:"ldonPointsMir,omitempty" db:"ldon_points_mir"`                 //`ldon_points_mir` int(11) unsigned NOT NULL DEFAULT '0',
	LdonPointsMmc         int64   `json:"ldonPointsMmc,omitempty" db:"ldon_points_mmc"`                 //`ldon_points_mmc` int(11) unsigned NOT NULL DEFAULT '0',
	LdonPointsRuj         int64   `json:"ldonPointsRuj,omitempty" db:"ldon_points_ruj"`                 //`ldon_points_ruj` int(11) unsigned NOT NULL DEFAULT '0',
	LdonPointsTak         int64   `json:"ldonPointsTak,omitempty" db:"ldon_points_tak"`                 //`ldon_points_tak` int(11) unsigned NOT NULL DEFAULT '0',
	LdonPointsAvailable   int64   `json:"ldonPointsAvailable,omitempty" db:"ldon_points_available"`     //`ldon_points_available` int(11) unsigned NOT NULL DEFAULT '0',
	TributeTimeRemaining  int64   `json:"tributeTimeRemaining,omitempty" db:"tribute_time_remaining"`   //`tribute_time_remaining` int(11) unsigned NOT NULL DEFAULT '0',
	CareerTributePoints   int64   `json:"careerTributePoints,omitempty" db:"career_tribute_points"`     //`career_tribute_points` int(11) unsigned NOT NULL DEFAULT '0',
	TributePoints         int64   `json:"tributePoints,omitempty" db:"tribute_points"`                  //`tribute_points` int(11) unsigned NOT NULL DEFAULT '0',
	TributeActive         int64   `json:"tributeActive,omitempty" db:"tribute_active"`                  //`tribute_active` int(11) unsigned NOT NULL DEFAULT '0',
	PvpStatus             int64   `json:"pvpStatus,omitempty" db:"pvp_status"`                          //`pvp_status` tinyint(11) unsigned NOT NULL DEFAULT '0',
	PvpKills              int64   `json:"pvpKills,omitempty" db:"pvp_kills"`                            //`pvp_kills` int(11) unsigned NOT NULL DEFAULT '0',
	PvpDeaths             int64   `json:"pvpDeaths,omitempty" db:"pvp_deaths"`                          //`pvp_deaths` int(11) unsigned NOT NULL DEFAULT '0',
	PvpCurrentPoints      int64   `json:"pvpCurrentPoints,omitempty" db:"pvp_current_points"`           //`pvp_current_points` int(11) unsigned NOT NULL DEFAULT '0',
	PvpCareerPoints       int64   `json:"pvpCareerPoints,omitempty" db:"pvp_career_points"`             //`pvp_career_points` int(11) unsigned NOT NULL DEFAULT '0',
	PvpBestKillStreak     int64   `json:"pvpBestKillStreak,omitempty" db:"pvp_best_kill_streak"`        //`pvp_best_kill_streak` int(11) unsigned NOT NULL DEFAULT '0',
	PvpWorstDeathStreak   int64   `json:"pvpWorstDeathStreak,omitempty" db:"pvp_worst_death_streak"`    //`pvp_worst_death_streak` int(11) unsigned NOT NULL DEFAULT '0',
	PvpCurrentKillStreak  int64   `json:"pvpCurrentKillStreak,omitempty" db:"pvp_current_kill_streak"`  //`pvp_current_kill_streak` int(11) unsigned NOT NULL DEFAULT '0',
	Pvp2                  int64   `json:"pvp2,omitempty" db:"pvp2"`                                     //`pvp2` int(11) unsigned NOT NULL DEFAULT '0',
	PvpType               int64   `json:"pvpType,omitempty" db:"pvp_type"`                              //`pvp_type` int(11) unsigned NOT NULL DEFAULT '0',
	ShowHelm              int64   `json:"showHelm,omitempty" db:"show_helm"`                            //`show_helm` int(11) unsigned NOT NULL DEFAULT '0',
	GroupAutoConsent      int64   `json:"groupAutoConsent,omitempty" db:"group_auto_consent"`           //`group_auto_consent` tinyint(11) unsigned NOT NULL DEFAULT '0',
	RaidAutoConsent       int64   `json:"raidAutoConsent,omitempty" db:"raid_auto_consent"`             //`raid_auto_consent` tinyint(11) unsigned NOT NULL DEFAULT '0',
	GuildAutoConsent      int64   `json:"guildAutoConsent,omitempty" db:"guild_auto_consent"`           //`guild_auto_consent` tinyint(11) unsigned NOT NULL DEFAULT '0',
	LeadershipExpOn       int64   `json:"leadershipExpOn,omitempty" db:"leadership_exp_on"`             //`leadership_exp_on` tinyint(11) unsigned NOT NULL DEFAULT '0',
	Resttimer             int64   `json:"RestTimer,omitempty" db:"RestTimer"`                           //`RestTimer` int(11) unsigned NOT NULL DEFAULT '0',
	AirRemaining          int64   `json:"airRemaining,omitempty" db:"air_remaining"`                    //`air_remaining` int(11) unsigned NOT NULL DEFAULT '0',
	AutosplitEnabled      int64   `json:"autosplitEnabled,omitempty" db:"autosplit_enabled"`            //`autosplit_enabled` int(11) unsigned NOT NULL DEFAULT '0',
	Lfp                   int64   `json:"lfp,omitempty" db:"lfp"`                                       //`lfp` tinyint(1) unsigned NOT NULL DEFAULT '0',
	Lfg                   int64   `json:"lfg,omitempty" db:"lfg"`                                       //`lfg` tinyint(1) unsigned NOT NULL DEFAULT '0',
	Mailkey               string  `json:"mailkey,omitempty" db:"mailkey"`                               //`mailkey` char(16) NOT NULL DEFAULT '',
	Xtargets              int64   `json:"xtargets,omitempty" db:"xtargets"`                             //`xtargets` tinyint(3) unsigned NOT NULL DEFAULT '5',
	Firstlogon            int64   `json:"firstlogon,omitempty" db:"firstlogon"`                         //`firstlogon` tinyint(3) NOT NULL DEFAULT '0',
	EAaEffects            int64   `json:"eAaEffects,omitempty" db:"e_aa_effects"`                       //`e_aa_effects` int(11) unsigned NOT NULL DEFAULT '0',
	EPercentToAa          int64   `json:"ePercentToAa,omitempty" db:"e_percent_to_aa"`                  //`e_percent_to_aa` int(11) unsigned NOT NULL DEFAULT '0',
	EExpendedAaSpent      int64   `json:"eExpendedAaSpent,omitempty" db:"e_expended_aa_spent"`          //`e_expended_aa_spent` int(11) unsigned NOT NULL DEFAULT '0',
	AaPointsSpentOld      int64   `json:"aaPointsSpentOld,omitempty" db:"aa_points_spent_old"`          //`aa_points_spent_old` int(11) unsigned NOT NULL DEFAULT '0',
	AaPointsOld           int64   `json:"aaPointsOld,omitempty" db:"aa_points_old"`                     //`aa_points_old` int(11) unsigned NOT NULL DEFAULT '0',
	ELastInvsnapshot      int64   `json:"eLastInvsnapshot,omitempty" db:"e_last_invsnapshot"`           //`e_last_invsnapshot` int(11) unsigned NOT NULL DEFAULT '0',
}
