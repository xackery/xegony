package model

import (
	"database/sql"
)

// Spells is an array of Spell
// swagger:model
type Spells []*Spell

//Spell represents items inside everquest
// swagger:model
type Spell struct {
	Animation           *SpellAnimation       `json:"animation,omitempty"`
	BuffDurationFormula *SpellDurationFormula `json:"buffDurationFormula,omitempty"`
	CastingAnimation    *SpellAnimation       `json:"castingAnimation,omitempty"`
	Components          []*SpellComponent     `json:"components,omitempty"`
	DBStr               *DBStr                `json:"dBStr,omitempty"`
	DBStrEffect1        *DBStr                `json:"dBStrEffect1,omitempty"`
	DBStrEffect2        *DBStr                `json:"dBStrEffect2,omitempty"`
	EffectCategory      *SpellEffectCategory  `json:"effectCategory,omitempty"`
	EnvironmentType     *EnvironmentType      `json:"environmentType,omitempty"`
	Group               *SpellGroup           `json:"group,omitempty"`
	Icon                *SpellIcon            `json:"spellIcon,omitempty"`
	LightType           *LightType            `json:"lightType,omitempty"`
	NumHitsType         *SpellNumHitsType     `json:"numHitsType,omitempty"`
	OldIcon             *SpellOldIcon         `json:"spellOldIcon,omitempty"`
	Reagents            []*SpellReagent       `json:"reagents,omitempty"`
	RecourseLinkSpell   *Spell                `json:"recourseLinkSpell,omitempty"`
	ResistType          *ResistType           `json:"resistType,omitempty"`
	Skill               *Skill                `json:"skill,omitempty"`
	TargetAnimation     *SpellAnimation       `json:"targetAnimation,omitempty"`
	TargetType          *SpellTargetType      `json:"targetType,omitempty"`
	TeleportZone        *Zone                 `json:"teleportZone,omitempty"`
	TravelType          *SpellTravelType      `json:"travelType,omitempty"`
	Nimbus              *SpellNimbus          `json:"nimbus,omitempty"`
	Formula1            *SpellEffectFormula   `json:"formula1,omitempty"`
	Formula2            *SpellEffectFormula   `json:"formula2,omitempty"`
	Formula3            *SpellEffectFormula   `json:"formula3,omitempty"`
	Formula4            *SpellEffectFormula   `json:"formula4,omitempty"`
	Formula5            *SpellEffectFormula   `json:"formula5,omitempty"`
	Formula6            *SpellEffectFormula   `json:"formula6,omitempty"`
	Formula7            *SpellEffectFormula   `json:"formula7,omitempty"`
	Formula8            *SpellEffectFormula   `json:"formula8,omitempty"`
	Formula9            *SpellEffectFormula   `json:"formula9,omitempty"`
	Formula10           *SpellEffectFormula   `json:"formula10,omitempty"`
	Formula11           *SpellEffectFormula   `json:"formula11,omitempty"`
	Formula12           *SpellEffectFormula   `json:"formula12,omitempty"`
	Effect1             *SpellEffectType      `json:"effect1,omitempty"`
	Effect2             *SpellEffectType      `json:"effect2,omitempty"`
	Effect3             *SpellEffectType      `json:"effect3,omitempty"`
	Effect4             *SpellEffectType      `json:"effect4,omitempty"`
	Effect5             *SpellEffectType      `json:"effect5,omitempty"`
	Effect6             *SpellEffectType      `json:"effect6,omitempty"`
	Effect7             *SpellEffectType      `json:"effect7,omitempty"`
	Effect8             *SpellEffectType      `json:"effect8,omitempty"`
	Effect9             *SpellEffectType      `json:"effect9,omitempty"`
	Effect10            *SpellEffectType      `json:"effect10,omitempty"`
	Effect11            *SpellEffectType      `json:"effect11,omitempty"`
	Effect12            *SpellEffectType      `json:"effect12,omitempty"`
	Deity0              *Deity                `json:"deity0,omitempty"`
	Deity1              *Deity                `json:"deity1,omitempty"`
	Deity2              *Deity                `json:"deity2,omitempty"`
	Deity3              *Deity                `json:"deity3,omitempty"`
	Deity4              *Deity                `json:"deity4,omitempty"`
	Deity5              *Deity                `json:"deity5,omitempty"`
	Deity6              *Deity                `json:"deity6,omitempty"`
	Deity7              *Deity                `json:"deity7,omitempty"`
	Deity8              *Deity                `json:"deity8,omitempty"`
	Deity9              *Deity                `json:"deity9,omitempty"`
	Deity10             *Deity                `json:"deity10,omitempty"`
	Deity11             *Deity                `json:"deity11,omitempty"`
	Deity12             *Deity                `json:"deity12,omitempty"`
	Deity13             *Deity                `json:"deity13,omitempty"`
	Deity14             *Deity                `json:"deity14,omitempty"`
	Deity15             *Deity                `json:"deity15,omitempty"`
	Deity16             *Deity                `json:"deity16,omitempty"`

	ID                    int64          `json:"ID,omitempty" db:"id"`                                     //`id` int(11) NOT NULL DEFAULT '0',
	Name                  sql.NullString `json:"name,omitempty" db:"name"`                                 //`name` varchar(64) DEFAULT NULL,
	Player1               sql.NullString `json:"player1,omitempty" db:"player_1"`                          //`player_1` varchar(64) DEFAULT 'BLUE_TRAIL',
	TeleportZoneShortName sql.NullString `json:"teleportZoneShortName,omitempty" db:"teleport_zone"`       //`teleport_zone` varchar(64) DEFAULT NULL,
	YouCast               sql.NullString `json:"youCast,omitempty" db:"you_cast"`                          //`you_cast` varchar(120) DEFAULT NULL,
	OtherCasts            sql.NullString `json:"otherCasts,omitempty" db:"other_casts"`                    //`other_casts` varchar(120) DEFAULT NULL,
	CastOnYou             sql.NullString `json:"castOnYou,omitempty" db:"cast_on_you"`                     //`cast_on_you` varchar(120) DEFAULT NULL,
	CastOnOther           sql.NullString `json:"castOnOther,omitempty" db:"cast_on_other"`                 //`cast_on_other` varchar(120) DEFAULT NULL,
	SpellFades            sql.NullString `json:"spellFades,omitempty" db:"spell_fades"`                    //`spell_fades` varchar(120) DEFAULT NULL,
	Range                 int64          `json:"range,omitempty" db:"range"`                               //`range` int(11) NOT NULL DEFAULT '100',
	AoeRange              int64          `json:"aoeRange,omitempty" db:"aoerange"`                         //`aoerange` int(11) NOT NULL DEFAULT '0',
	PushBack              int64          `json:"pushBack,omitempty" db:"pushback"`                         //`pushback` int(11) NOT NULL DEFAULT '0',
	PushUp                int64          `json:"pushUp,omitempty" db:"pushup"`                             //`pushup` int(11) NOT NULL DEFAULT '0',
	CastTime              int64          `json:"castTime,omitempty" db:"cast_time"`                        //`cast_time` int(11) NOT NULL DEFAULT '0',
	RecoveryTime          int64          `json:"recoveryTime,omitempty" db:"recovery_time"`                //`recovery_time` int(11) NOT NULL DEFAULT '0',
	RecastTime            int64          `json:"recastTime,omitempty" db:"recast_time"`                    //`recast_time` int(11) NOT NULL DEFAULT '0',
	BuffDurationFormulaID int64          `json:"buffDurationFormulaID,omitempty" db:"buffdurationformula"` //`buffdurationformula` int(11) NOT NULL DEFAULT '7',
	BuffDuration          int64          `json:"buffduration,omitempty" db:"buffduration"`                 //`buffduration` int(11) NOT NULL DEFAULT '65',
	AEDuration            int64          `json:"aEDuration,omitempty" db:"AEDuration"`                     //`AEDuration` int(11) NOT NULL DEFAULT '0',
	Mana                  int64          `json:"mana,omitempty" db:"mana"`                                 //`mana` int(11) NOT NULL DEFAULT '0',
	UnusedIconID          int64          `json:"unusuedIconID,omitempty" db:"icon"`                        //`icon` int(11) NOT NULL DEFAULT '0',
	OldIconID             int64          `json:"oldIconID,omitempty" db:"memicon"`                         //`memicon` int(11) NOT NULL DEFAULT '0',
	LightTypeID           int64          `json:"lightTypeID,omitempty" db:"LightType"`                     //`LightType` int(11) NOT NULL DEFAULT '0',
	EffectCategoryID      int64          `json:"effectCategoryID,omitempty" db:"goodEffect"`               //`goodEffect` int(11) NOT NULL DEFAULT '0',
	Activated             int64          `json:"activated,omitempty" db:"Activated"`                       //`Activated` int(11) NOT NULL DEFAULT '0',
	ResistTypeID          int64          `json:"resistTypeID,omitempty" db:"resisttype"`                   //`resisttype` int(11) NOT NULL DEFAULT '0',
	TargetTypeID          int64          `json:"targetTypeID,omitempty" db:"targettype"`                   //`targettype` int(11) NOT NULL DEFAULT '2',
	FizzleDifficulty      int64          `json:"fizzleDifficulty,omitempty" db:"basediff"`                 //`basediff` int(11) NOT NULL DEFAULT '0',
	SkillID               int64          `json:"skillID,omitempty" db:"skill"`                             //`skill` int(11) NOT NULL DEFAULT '98',
	Zonetype              int64          `json:"zoneType,omitempty" db:"zonetype"`                         //`zonetype` int(11) NOT NULL DEFAULT '-1',
	EnvironmentTypeID     int64          `json:"environmentTypeID,omitempty" db:"EnvironmentType"`         //`EnvironmentType` int(11) NOT NULL DEFAULT '0',
	TimeOfDay             int64          `json:"timeOfDay,omitempty" db:"TimeOfDay"`                       //`TimeOfDay` int(11) NOT NULL DEFAULT '0',
	CastingAnimationID    int64          `json:"castingAnimationID,omitempty" db:"CastingAnim"`            //`CastingAnim` int(11) NOT NULL DEFAULT '44',
	TargetAnimationID     int64          `json:"targetAnimationID,omitempty" db:"TargetAnim"`              //`TargetAnim` int(11) NOT NULL DEFAULT '13',
	TravelTypeID          int64          `json:"travelTypeID,omitempty" db:"TravelType"`                   //`TravelType` int(11) NOT NULL DEFAULT '0',
	SpellAffectID         int64          `json:"spellAffectID,omitempty" db:"SpellAffectIndex"`            //`SpellAffectIndex` int(11) NOT NULL DEFAULT '-1',
	DisallowSit           int64          `json:"disallowSit,omitempty" db:"disallow_sit"`                  //`disallow_sit` int(11) NOT NULL DEFAULT '0',
	Field142              int64          `json:"field142,omitempty" db:"field142"`                         //`field142` int(11) NOT NULL DEFAULT '100',
	Field143              int64          `json:"field143,omitempty" db:"field143"`                         //`field143` int(11) NOT NULL DEFAULT '0',
	IconID                int64          `json:"icon,omitempty" db:"new_icon"`                             //`new_icon` int(11) NOT NULL DEFAULT '161',
	AnimationID           int64          `json:"animationID,omitempty" db:"spellanim"`                     //`spellanim` int(11) NOT NULL DEFAULT '0',
	Uninterruptable       int64          `json:"uninterruptable,omitempty" db:"uninterruptable"`           //`uninterruptable` int(11) NOT NULL DEFAULT '0',
	ResistDifference      int64          `json:"resistDifference,omitempty" db:"ResistDiff"`               //`ResistDiff` int(11) NOT NULL DEFAULT '-150',
	DotStackingExempt     int64          `json:"dotStackingExempt,omitempty" db:"dot_stacking_exempt"`     //`dot_stacking_exempt` int(11) NOT NULL DEFAULT '0',
	CanDelete             int64          `json:"canDelete,omitempty" db:"deleteable"`                      //`deleteable` int(11) NOT NULL DEFAULT '0',
	RecourseLinkSpellID   int64          `json:"recourseLinkSpellID,omitempty" db:"RecourseLink"`          //`RecourseLink` int(11) NOT NULL DEFAULT '0',
	NoPartialResist       int64          `json:"noPartialResist,omitempty" db:"no_partial_resist"`         //`no_partial_resist` int(11) NOT NULL DEFAULT '0',
	Field152              int64          `json:"field152,omitempty" db:"field152"`                         //`field152` int(11) NOT NULL DEFAULT '0',
	Field153              int64          `json:"field153,omitempty" db:"field153"`                         //`field153` int(11) NOT NULL DEFAULT '0',
	UseShortBuffBox       int64          `json:"useShortBuffBox,omitempty" db:"short_buff_box"`            //`short_buff_box` int(11) NOT NULL DEFAULT '-1',
	DBStrID               int64          `json:"dBStrID,omitempty" db:"descnum"`                           //`descnum` int(11) NOT NULL DEFAULT '0',
	DBStrTypeID           int64          `json:"dBStrTypeID,omitempty" db:"typedescnum"`                   //`typedescnum` int(11) DEFAULT NULL,
	DBStrEffectID         int64          `json:"dBStrEffectID,omitempty" db:"effectdescnum"`               //`effectdescnum` int(11) DEFAULT NULL,
	DBStrEffect2ID        int64          `json:"dBStrEffect2ID,omitempty" db:"effectdescnum2"`             //`effectdescnum2` int(11) NOT NULL DEFAULT '0',
	NpcNoLos              int64          `json:"npcNoLos,omitempty" db:"npc_no_los"`                       //`npc_no_los` int(11) NOT NULL DEFAULT '0',
	Field160              int64          `json:"field160,omitempty" db:"field160"`                         //`field160` int(11) NOT NULL DEFAULT '0',
	Reflectable           int64          `json:"reflectable,omitempty" db:"reflectable"`                   //`reflectable` int(11) NOT NULL DEFAULT '0',
	BonusHate             int64          `json:"bonusHate,omitempty" db:"bonushate"`                       //`bonushate` int(11) NOT NULL DEFAULT '0',
	Field163              int64          `json:"field163,omitempty" db:"field163"`                         //`field163` int(11) NOT NULL DEFAULT '100',
	Field164              int64          `json:"field164,omitempty" db:"field164"`                         //`field164` int(11) NOT NULL DEFAULT '-150',
	LdonTrap              int64          `json:"ldonTrap,omitempty" db:"ldon_trap"`                        //`ldon_trap` int(11) NOT NULL DEFAULT '0',
	Endurcost             int64          `json:"endurCost,omitempty" db:"EndurCost"`                       //`EndurCost` int(11) NOT NULL DEFAULT '0',
	Endurtimerindex       int64          `json:"endurTimerIndex,omitempty" db:"EndurTimerIndex"`           //`EndurTimerIndex` int(11) NOT NULL DEFAULT '0',
	Isdiscipline          int64          `json:"isDiscipline,omitempty" db:"IsDiscipline"`                 //`IsDiscipline` int(11) NOT NULL DEFAULT '0',
	Field169              int64          `json:"field169,omitempty" db:"field169"`                         //`field169` int(11) NOT NULL DEFAULT '0',
	Field170              int64          `json:"field170,omitempty" db:"field170"`                         //`field170` int(11) NOT NULL DEFAULT '0',
	Field171              int64          `json:"field171,omitempty" db:"field171"`                         //`field171` int(11) NOT NULL DEFAULT '0',
	Field172              int64          `json:"field172,omitempty" db:"field172"`                         //`field172` int(11) NOT NULL DEFAULT '0',
	Hateadded             int64          `json:"hateAdded,omitempty" db:"HateAdded"`                       //`HateAdded` int(11) NOT NULL DEFAULT '0',
	Endurupkeep           int64          `json:"endurUpkeep,omitempty" db:"EndurUpkeep"`                   //`EndurUpkeep` int(11) NOT NULL DEFAULT '0',
	NumHitsTypeID         int64          `json:"numHitsTypeID,omitempty" db:"numhitstype"`                 //`numhitstype` int(11) NOT NULL DEFAULT '0',
	NumHits               int64          `json:"numHits,omitempty" db:"numhits"`                           //`numhits` int(11) NOT NULL DEFAULT '0',
	PvpResistBase         int64          `json:"pvpResistBase,omitempty" db:"pvpresistbase"`               //`pvpresistbase` int(11) NOT NULL DEFAULT '-150',
	PvpResistCalc         int64          `json:"pvpResistCalc,omitempty" db:"pvpresistcalc"`               //`pvpresistcalc` int(11) NOT NULL DEFAULT '100',
	PvpResistCap          int64          `json:"pvpResistCap,omitempty" db:"pvpresistcap"`                 //`pvpresistcap` int(11) NOT NULL DEFAULT '-150',
	SpellCategory         int64          `json:"spellCategory,omitempty" db:"spell_category"`              //`spell_category` int(11) NOT NULL DEFAULT '-99',
	Field181              int64          `json:"field181,omitempty" db:"field181"`                         //`field181` int(11) NOT NULL DEFAULT '7',
	Field182              int64          `json:"field182,omitempty" db:"field182"`                         //`field182` int(11) NOT NULL DEFAULT '65',
	PcNpcOnlyFlag         int64          `json:"pcNpcOnlyFlag,omitempty" db:"pcnpc_only_flag"`             //`pcnpc_only_flag` int(11) DEFAULT '0',
	CastNotStanding       int64          `json:"castNotStanding,omitempty" db:"cast_not_standing"`         //`cast_not_standing` int(11) DEFAULT '0',
	CanMgb                int64          `json:"canMgb,omitempty" db:"can_mgb"`                            //`can_mgb` int(11) NOT NULL DEFAULT '0',
	NoDispell             int64          `json:"noDispell,omitempty" db:"nodispell"`                       //`nodispell` int(11) NOT NULL DEFAULT '-1',
	NpcCategory           int64          `json:"npcCategory,omitempty" db:"npc_category"`                  //`npc_category` int(11) NOT NULL DEFAULT '0',
	NpcUsefulness         int64          `json:"npcUsefulness,omitempty" db:"npc_usefulness"`              //`npc_usefulness` int(11) NOT NULL DEFAULT '0',
	MinResist             int64          `json:"minResist,omitempty" db:"MinResist"`                       //`MinResist` int(11) NOT NULL DEFAULT '0',
	MaxResist             int64          `json:"maxResist,omitempty" db:"MaxResist"`                       //`MaxResist` int(11) NOT NULL DEFAULT '0',
	ViralTargets          int64          `json:"viralTargets,omitempty" db:"viral_targets"`                //`viral_targets` int(11) NOT NULL DEFAULT '0',
	ViralTimer            int64          `json:"viralTimer,omitempty" db:"viral_timer"`                    //`viral_timer` int(11) NOT NULL DEFAULT '0',
	NimbusEffectID        int64          `json:"nimbusEffectID,omitempty" db:"nimbuseffect"`               //`nimbuseffect` int(11) DEFAULT '0',
	Conestartangle        int64          `json:"coneStartAngle,omitempty" db:"ConeStartAngle"`             //`ConeStartAngle` int(11) NOT NULL DEFAULT '0',
	Conestopangle         int64          `json:"coneStopAngle,omitempty" db:"ConeStopAngle"`               //`ConeStopAngle` int(11) NOT NULL DEFAULT '0',
	Sneaking              int64          `json:"sneaking,omitempty" db:"sneaking"`                         //`sneaking` int(11) NOT NULL DEFAULT '0',
	NotExtendable         int64          `json:"notExtendable,omitempty" db:"not_extendable"`              //`not_extendable` int(11) NOT NULL DEFAULT '0',
	Field198              int64          `json:"field198,omitempty" db:"field198"`                         //`field198` int(11) NOT NULL DEFAULT '0',
	Field199              int64          `json:"field199,omitempty" db:"field199"`                         //`field199` int(11) NOT NULL DEFAULT '1',
	Suspendable           sql.NullInt64  `json:"suspendable,omitempty" db:"suspendable"`                   //`suspendable` int(11) DEFAULT '0',
	ViralRange            int64          `json:"viralRange,omitempty" db:"viral_range"`                    //`viral_range` int(11) NOT NULL DEFAULT '0',
	Songcap               sql.NullInt64  `json:"songcap,omitempty" db:"songcap"`                           //`songcap` int(11) DEFAULT '0',
	Field203              sql.NullInt64  `json:"field203,omitempty" db:"field203"`                         //`field203` int(11) DEFAULT '0',
	Field204              sql.NullInt64  `json:"field204,omitempty" db:"field204"`                         //`field204` int(11) DEFAULT '0',
	NoBlock               int64          `json:"noBlock,omitempty" db:"no_block"`                          //`no_block` int(11) NOT NULL DEFAULT '0',
	Field206              sql.NullInt64  `json:"field206,omitempty" db:"field206"`                         //`field206` int(11) DEFAULT '-1',
	SpellGroupID          sql.NullInt64  `json:"spellGroup,omitempty" db:"spellgroup"`                     //`spellgroup` int(11) DEFAULT '0',
	Rank                  int64          `json:"rank,omitempty" db:"rank"`                                 //`rank` int(11) NOT NULL DEFAULT '0',
	Field209              sql.NullInt64  `json:"field209,omitempty" db:"field209"`                         //`field209` int(11) DEFAULT '0',
	Field210              sql.NullInt64  `json:"field210,omitempty" db:"field210"`                         //`field210` int(11) DEFAULT '1',
	Castrestriction       int64          `json:"castRestriction,omitempty" db:"CastRestriction"`           //`CastRestriction` int(11) NOT NULL DEFAULT '0',
	Allowrest             sql.NullInt64  `json:"allowrest,omitempty" db:"allowrest"`                       //`allowrest` int(11) DEFAULT '0',
	Incombat              int64          `json:"inCombat,omitempty" db:"InCombat"`                         //`InCombat` int(11) NOT NULL DEFAULT '0',
	Outofcombat           int64          `json:"outofCombat,omitempty" db:"OutofCombat"`                   //`OutofCombat` int(11) NOT NULL DEFAULT '0',
	Field215              sql.NullInt64  `json:"field215,omitempty" db:"field215"`                         //`field215` int(11) DEFAULT '0',
	Field216              sql.NullInt64  `json:"field216,omitempty" db:"field216"`                         //`field216` int(11) DEFAULT '0',
	Field217              sql.NullInt64  `json:"field217,omitempty" db:"field217"`                         //`field217` int(11) DEFAULT '0',
	Aemaxtargets          int64          `json:"aemaxtargets,omitempty" db:"aemaxtargets"`                 //`aemaxtargets` int(11) NOT NULL DEFAULT '0',
	Maxtargets            sql.NullInt64  `json:"maxtargets,omitempty" db:"maxtargets"`                     //`maxtargets` int(11) DEFAULT '0',
	Field220              sql.NullInt64  `json:"field220,omitempty" db:"field220"`                         //`field220` int(11) DEFAULT '0',
	Field221              sql.NullInt64  `json:"field221,omitempty" db:"field221"`                         //`field221` int(11) DEFAULT '0',
	Field222              sql.NullInt64  `json:"field222,omitempty" db:"field222"`                         //`field222` int(11) DEFAULT '0',
	Field223              sql.NullInt64  `json:"field223,omitempty" db:"field223"`                         //`field223` int(11) DEFAULT '0',
	Persistdeath          sql.NullInt64  `json:"persistdeath,omitempty" db:"persistdeath"`                 //`persistdeath` int(11) DEFAULT '0',
	Field225              int64          `json:"field225,omitempty" db:"field225"`                         //`field225` int(11) NOT NULL DEFAULT '0',
	Field226              int64          `json:"field226,omitempty" db:"field226"`                         //`field226` int(11) NOT NULL DEFAULT '0',
	MinDist               float64        `json:"minDist,omitempty" db:"min_dist"`                          //`min_dist` float NOT NULL DEFAULT '0',
	MinDistMod            float64        `json:"minDistMod,omitempty" db:"min_dist_mod"`                   //`min_dist_mod` float NOT NULL DEFAULT '0',
	MaxDist               float64        `json:"maxDist,omitempty" db:"max_dist"`                          //`max_dist` float NOT NULL DEFAULT '0',
	MaxDistMod            float64        `json:"maxDistMod,omitempty" db:"max_dist_mod"`                   //`max_dist_mod` float NOT NULL DEFAULT '0',
	MinRange              int64          `json:"minRange,omitempty" db:"min_range"`                        //`min_range` int(11) NOT NULL DEFAULT '0',
	Field232              int64          `json:"field232,omitempty" db:"field232"`                         //`field232` int(11) NOT NULL DEFAULT '0',
	Field233              int64          `json:"field233,omitempty" db:"field233"`                         //`field233` int(11) NOT NULL DEFAULT '0',
	Field234              int64          `json:"field234,omitempty" db:"field234"`                         //`field234` int(11) NOT NULL DEFAULT '0',
	Field235              int64          `json:"field235,omitempty" db:"field235"`                         //`field235` int(11) NOT NULL DEFAULT '0',
	Field236              int64          `json:"field236,omitempty" db:"field236"`                         //`field236` int(11) NOT NULL DEFAULT '0',

	FormulaID1  int64 `json:"formulaID1,omitempty" db:"formula1"`               //`formula1` int(11) NOT NULL DEFAULT '100',
	Base1       int64 `json:"baseValue1,omitempty" db:"effect_base_value1"`     //`effect_base_value1` int(11) NOT NULL DEFAULT '100',
	Limit1      int64 `json:"limitValue1,omitempty" db:"effect_limit_value1"`   //`effect_limit_value1` int(11) NOT NULL DEFAULT '0',
	Max1        int64 `json:"max1,omitempty" db:"max1"`                         //`max1` int(11) NOT NULL DEFAULT '0',
	EffectID1   int64 `json:"effectID1,omitempty" db:"effectid1"`               //`effectid1` int(11) NOT NULL DEFAULT '254',
	FormulaID2  int64 `json:"formulaID2,omitempty" db:"formula2"`               //`formula2` int(11) NOT NULL DEFAULT '100',
	Base2       int64 `json:"baseValue2,omitempty" db:"effect_base_value2"`     //`effect_base_value2` int(11) NOT NULL DEFAULT '100',
	Limit2      int64 `json:"limitValue2,omitempty" db:"effect_limit_value2"`   //`effect_limit_value2` int(11) NOT NULL DEFAULT '0',
	Max2        int64 `json:"max2,omitempty" db:"max2"`                         //`max2` int(11) NOT NULL DEFAULT '0',
	EffectID2   int64 `json:"effectID2,omitempty" db:"effectid2"`               //`effectid2` int(11) NOT NULL DEFAULT '254',
	FormulaID3  int64 `json:"formulaID3,omitempty" db:"formula3"`               //`formula3` int(11) NOT NULL DEFAULT '100',
	Base3       int64 `json:"baseValue3,omitempty" db:"effect_base_value3"`     //`effect_base_value3` int(11) NOT NULL DEFAULT '100',
	Limit3      int64 `json:"limitValue3,omitempty" db:"effect_limit_value3"`   //`effect_limit_value3` int(11) NOT NULL DEFAULT '0',
	Max3        int64 `json:"max3,omitempty" db:"max3"`                         //`max3` int(11) NOT NULL DEFAULT '0',
	EffectID3   int64 `json:"effectID3,omitempty" db:"effectid3"`               //`effectid3` int(11) NOT NULL DEFAULT '254',
	FormulaID4  int64 `json:"formulaID4,omitempty" db:"formula4"`               //`formula4` int(11) NOT NULL DEFAULT '100',
	Base4       int64 `json:"baseValue4,omitempty" db:"effect_base_value4"`     //`effect_base_value4` int(11) NOT NULL DEFAULT '100',
	Limit4      int64 `json:"limitValue4,omitempty" db:"effect_limit_value4"`   //`effect_limit_value4` int(11) NOT NULL DEFAULT '0',
	Max4        int64 `json:"max4,omitempty" db:"max4"`                         //`max4` int(11) NOT NULL DEFAULT '0',
	EffectID4   int64 `json:"effectID4,omitempty" db:"effectid4"`               //`effectid4` int(11) NOT NULL DEFAULT '254',
	FormulaID5  int64 `json:"formulaID5,omitempty" db:"formula5"`               //`formula5` int(11) NOT NULL DEFAULT '100',
	Base5       int64 `json:"baseValue5,omitempty" db:"effect_base_value5"`     //`effect_base_value5` int(11) NOT NULL DEFAULT '100',
	Limit5      int64 `json:"limitValue5,omitempty" db:"effect_limit_value5"`   //`effect_limit_value5` int(11) NOT NULL DEFAULT '0',
	Max5        int64 `json:"max5,omitempty" db:"max5"`                         //`max5` int(11) NOT NULL DEFAULT '0',
	EffectID5   int64 `json:"effectID5,omitempty" db:"effectid5"`               //`effectid5` int(11) NOT NULL DEFAULT '254',
	FormulaID6  int64 `json:"formulaID6,omitempty" db:"formula6"`               //`formula6` int(11) NOT NULL DEFAULT '100',
	Base6       int64 `json:"baseValue6,omitempty" db:"effect_base_value6"`     //`effect_base_value6` int(11) NOT NULL DEFAULT '100',
	Limit6      int64 `json:"limitValue6,omitempty" db:"effect_limit_value6"`   //`effect_limit_value6` int(11) NOT NULL DEFAULT '0',
	Max6        int64 `json:"max6,omitempty" db:"max6"`                         //`max6` int(11) NOT NULL DEFAULT '0',
	EffectID6   int64 `json:"effectID6,omitempty" db:"effectid6"`               //`effectid6` int(11) NOT NULL DEFAULT '254',
	FormulaID7  int64 `json:"formulaID7,omitempty" db:"formula7"`               //`formula7` int(11) NOT NULL DEFAULT '100',
	Base7       int64 `json:"baseValue7,omitempty" db:"effect_base_value7"`     //`effect_base_value7` int(11) NOT NULL DEFAULT '100',
	Limit7      int64 `json:"limitValue7,omitempty" db:"effect_limit_value7"`   //`effect_limit_value7` int(11) NOT NULL DEFAULT '0',
	Max7        int64 `json:"max7,omitempty" db:"max7"`                         //`max7` int(11) NOT NULL DEFAULT '0',
	EffectID7   int64 `json:"effectID7,omitempty" db:"effectid7"`               //`effectid7` int(11) NOT NULL DEFAULT '254',
	FormulaID8  int64 `json:"formulaID8,omitempty" db:"formula8"`               //`formula8` int(11) NOT NULL DEFAULT '100',
	Base8       int64 `json:"baseValue8,omitempty" db:"effect_base_value8"`     //`effect_base_value8` int(11) NOT NULL DEFAULT '100',
	Limit8      int64 `json:"limitValue8,omitempty" db:"effect_limit_value8"`   //`effect_limit_value8` int(11) NOT NULL DEFAULT '0',
	Max8        int64 `json:"max8,omitempty" db:"max8"`                         //`max8` int(11) NOT NULL DEFAULT '0',
	EffectID8   int64 `json:"effectID8,omitempty" db:"effectid8"`               //`effectid8` int(11) NOT NULL DEFAULT '254',
	FormulaID9  int64 `json:"formulaID9,omitempty" db:"formula9"`               //`formula9` int(11) NOT NULL DEFAULT '100',
	Base9       int64 `json:"baseValue9,omitempty" db:"effect_base_value9"`     //`effect_base_value9` int(11) NOT NULL DEFAULT '100',
	Limit9      int64 `json:"limitValue9,omitempty" db:"effect_limit_value9"`   //`effect_limit_value9` int(11) NOT NULL DEFAULT '0',
	Max9        int64 `json:"max9,omitempty" db:"max9"`                         //`max9` int(11) NOT NULL DEFAULT '0',
	EffectID9   int64 `json:"effectID9,omitempty" db:"effectid9"`               //`effectid9` int(11) NOT NULL DEFAULT '254',
	FormulaID10 int64 `json:"formulaID10,omitempty" db:"formula10"`             //`formula10` int(11) NOT NULL DEFAULT '100',
	Base10      int64 `json:"baseValue10,omitempty" db:"effect_base_value10"`   //`effect_base_value10` int(11) NOT NULL DEFAULT '100',
	Limit10     int64 `json:"limitValue10,omitempty" db:"effect_limit_value10"` //`effect_limit_value10` int(11) NOT NULL DEFAULT '0',
	Max10       int64 `json:"max10,omitempty" db:"max10"`                       //`max10` int(11) NOT NULL DEFAULT '0',
	EffectID10  int64 `json:"effectID10,omitempty" db:"effectid10"`             //`effectid10` int(11) NOT NULL DEFAULT '254',
	FormulaID11 int64 `json:"formulaID11,omitempty" db:"formula11"`             //`formula11` int(11) NOT NULL DEFAULT '100',
	Base11      int64 `json:"baseValue11,omitempty" db:"effect_base_value11"`   //`effect_base_value11` int(11) NOT NULL DEFAULT '100',
	Limit11     int64 `json:"limitValue11,omitempty" db:"effect_limit_value11"` //`effect_limit_value11` int(11) NOT NULL DEFAULT '0',
	Max11       int64 `json:"max11,omitempty" db:"max11"`                       //`max11` int(11) NOT NULL DEFAULT '0',
	EffectID11  int64 `json:"effectID11,omitempty" db:"effectid11"`             //`effectid11` int(11) NOT NULL DEFAULT '254',
	FormulaID12 int64 `json:"formulaID12,omitempty" db:"formula12"`             //`formula12` int(11) NOT NULL DEFAULT '100',
	Base12      int64 `json:"baseValue12,omitempty" db:"effect_base_value12"`   //`effect_base_value12` int(11) NOT NULL DEFAULT '100',
	Limit12     int64 `json:"limitValue12,omitempty" db:"effect_limit_value12"` //`effect_limit_value12` int(11) NOT NULL DEFAULT '0',
	Max12       int64 `json:"max12,omitempty" db:"max12"`                       //`max12` int(11) NOT NULL DEFAULT '0',
	EffectID12  int64 `json:"effectID12,omitempty" db:"effectid12"`             //`effectid12` int(11) NOT NULL DEFAULT '254',

	DeityID0  int64 `json:"deityID0,omitempty" db:"deities0"`   //`deities0` int(11) NOT NULL DEFAULT '0',
	DeityID1  int64 `json:"deityID1,omitempty" db:"deities1"`   //`deities1` int(11) NOT NULL DEFAULT '0',
	DeityID2  int64 `json:"deityID2,omitempty" db:"deities2"`   //`deities2` int(11) NOT NULL DEFAULT '0',
	DeityID3  int64 `json:"deityID3,omitempty" db:"deities3"`   //`deities3` int(11) NOT NULL DEFAULT '0',
	DeityID4  int64 `json:"deityID4,omitempty" db:"deities4"`   //`deities4` int(11) NOT NULL DEFAULT '0',
	DeityID5  int64 `json:"deityID5,omitempty" db:"deities5"`   //`deities5` int(11) NOT NULL DEFAULT '0',
	DeityID6  int64 `json:"deityID6,omitempty" db:"deities6"`   //`deities6` int(11) NOT NULL DEFAULT '0',
	DeityID7  int64 `json:"deityID7,omitempty" db:"deities7"`   //`deities7` int(11) NOT NULL DEFAULT '0',
	DeityID8  int64 `json:"deityID8,omitempty" db:"deities8"`   //`deities8` int(11) NOT NULL DEFAULT '0',
	DeityID9  int64 `json:"deityID9,omitempty" db:"deities9"`   //`deities9` int(11) NOT NULL DEFAULT '0',
	DeityID10 int64 `json:"deityID10,omitempty" db:"deities10"` //`deities10` int(11) NOT NULL DEFAULT '0',
	DeityID11 int64 `json:"deityID11,omitempty" db:"deities11"` //`deities11` int(11) NOT NULL DEFAULT '0',
	DeityID12 int64 `json:"deityID12,omitempty" db:"deities12"` //`deities12` int(12) NOT NULL DEFAULT '0',
	DeityID13 int64 `json:"deityID13,omitempty" db:"deities13"` //`deities13` int(11) NOT NULL DEFAULT '0',
	DeityID14 int64 `json:"deityID14,omitempty" db:"deities14"` //`deities14` int(11) NOT NULL DEFAULT '0',
	DeityID15 int64 `json:"deityID15,omitempty" db:"deities15"` //`deities15` int(11) NOT NULL DEFAULT '0',
	DeityID16 int64 `json:"deityID16,omitempty" db:"deities16"` //`deities16` int(11) NOT NULL DEFAULT '0',

}

//DescriptionName returns a summary of this spell's attributes
func (s *Spell) DescriptionName() string {
	return "Description"
}
