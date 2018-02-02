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
	Animation           *SpellAnimation      `json:"animation,omitempty"`
	BuffDurationFormula *SpellFormula        `json:"buffDurationFormula,omitempty"`
	CastingAnimation    *SpellAnimation      `json:"castingAnimation,omitempty"`
	Components          []*SpellComponent    `json:"components,omitempty"`
	Deitys              []*SpellDeity        `json:"deities,omitempty"`
	DBStr               *DBStr               `json:"dBStr,omitempty"`
	DBStrEffect1        *DBStr               `json:"dBStrEffect1,omitempty"`
	DBStrEffect2        *DBStr               `json:"dBStrEffect2,omitempty"`
	EffectCategory      *SpellEffectCategory `json:"effectCategory,omitempty"`
	Effects             []*SpellEffect       `json:"effects,omitempty"`
	EnvironmentType     *EnvironmentType     `json:"environmentType,omitempty"`
	Group               *SpellGroup          `json:"group"`
	Icon                *SpellIcon           `json:"spellIcon,omitempty"`
	LightType           *LightType           `json:"lightType,omitempty"`
	NumHitsType         *SpellNumHitsType    `json:"numHitsType,omitempty"`
	OldIcon             *SpellOldIcon        `json:"spellOldIcon,omitempty"`
	Reagents            []*SpellReagent      `json:"reagents,omitempty"`
	RecourseLinkSpell   *Spell               `json:"recourseLinkSpell,omitempty"`
	ResistType          *ResistType          `json:"resistType,omitempty"`
	Skill               *Skill               `json:"skill,omitempty"`
	TargetAnimation     *SpellAnimation      `json:"targetAnimation,omitempty"`
	TargetType          *SpellTargetType     `json:"targetType,omitempty"`
	TeleportZone        *Zone                `json:"teleportZone"`
	Nimbus              *SpellNimbus         `json:"nimbus"`

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
}

//DescriptionName returns a summary of this spell's attributes
func (s *Spell) DescriptionName() string {
	return "Description"
}
