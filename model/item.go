package model

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// Items contains an array of item
// swagger:model
type Items []*Item

// Item represents items inside everquest
// swagger:model
type Item struct {
	Classs     Classs `json:"classs,omitempty"`
	Races      Races  `json:"races,omitempty"`
	ProcEffect *Spell `json:"procEffect,omitempty`
	//Reference is used when listing from npc
	/*Reference  string        `json:"reference,omitempty"`
	SizeName   string        `json:"sizeName,omitempty"`
	SlotList   string        `json:"slotList,omitempty"`
	Slot       *Slot         `json:"slot,omitempty"`
	Category   *ItemCategory `json:"itemCategory,omitempty"`
	StyleColor string        `json:"styleColor,omitempty"`
	*/

	//ElementalDamageType
	//BardType

	//BagSize
	//AugType
	//AugRestriction
	//AugSlot
	//Size

	//normal items db
	ID                      int64          `json:"id,omitempty" db:"id"`                                   //
	Name                    string         `json:"name,omitempty" db:"name"`                               //
	Agility                 int64          `json:"agility,omitempty" db:"aagi"`                            //
	ArmorClass              int64          `json:"armorClass,omitempty" db:"ac"`                           //
	Accuracy                int64          `json:"accuracy,omitempty" db:"accuracy"`                       //
	Charisma                int64          `json:"charisma,omitempty" db:"acha"`                           //
	Dexterity               int64          `json:"dexterity,omitempty" db:"adex"`                          //
	Intelligence            int64          `json:"intelligence,omitempty" db:"aint"`                       //
	ArtifactFlag            int64          `json:"artifactFlag,omitempty" db:"artifactflag"`               //
	Stamina                 int64          `json:"stamina,omitempty" db:"asta"`                            //
	Strength                int64          `json:"strength,omitempty" db:"astr"`                           //
	Attack                  int64          `json:"attack,omitempty" db:"attack"`                           //
	AugRestrictionBit       int64          `json:"augRestrictionBit,omitempty" db:"augrestrict"`           //
	AugSlot1TypeID          int64          `json:"augSlot1Type,omitempty" db:"augslot1type"`               //
	AugSlot1Visible         sql.NullInt64  `json:"augSlot1Visible,omitempty" db:"augslot1visible"`         //
	AugSlot2TypeID          int64          `json:"augSlot2Type,omitempty" db:"augslot2type"`               //
	AugSlot2Visible         sql.NullInt64  `json:"augSlot2Visible,omitempty" db:"augslot2visible"`         //
	AugSlot3TypeID          int64          `json:"augSlot3Type,omitempty" db:"augslot3type"`               //
	AugSlot3Visible         sql.NullInt64  `json:"augSlot3Visible,omitempty" db:"augslot3visible"`         //
	AugSlot4TypeID          int64          `json:"augSlot4Type,omitempty" db:"augslot4type"`               //
	AugSlot4Visible         sql.NullInt64  `json:"augSlot4Visible,omitempty" db:"augslot4visible"`         //
	AugSlot5TypeID          int64          `json:"augSlot5Type,omitempty" db:"augslot5type"`               //
	AugSlot5Visible         sql.NullInt64  `json:"augSlot5Visible,omitempty" db:"augslot5visible"`         //
	AugSlot6TypeID          int64          `json:"augSlot6Type,omitempty" db:"augslot6type"`               //
	AugSlot6Visible         int64          `json:"augSlot6Visible,omitempty" db:"augslot6visible"`         //
	AugTypeID               int64          `json:"augTypeID,omitempty" db:"augtype"`                       //
	Avoidance               int64          `json:"avoidance,omitempty" db:"avoidance"`                     //
	Wisdom                  int64          `json:"wisdom,omitempty" db:"awis"`                             //
	BagSizeID               int64          `json:"bagSizeID,omitempty" db:"bagsize"`                       //
	BagSlotCount            int64          `json:"bagSlotCount,omitempty" db:"bagslots"`                   //
	BagTypeID               int64          `json:"bagTypeID,omitempty" db:"bagtype"`                       //
	BagWeightReduction      int64          `json:"bagWeightReduction,omitempty" db:"bagwr"`                //
	BaneDamageBodyAmount    int64          `json:"baneDamageBodyAmount,omitempty" db:"banedmgamt"`         //
	BaneDamageRaceAmonut    int64          `json:"baneDamageRaceAmount,omitempty" db:"banedmgraceamt"`     //
	BaneDamageBodyID        int64          `json:"baneDamageBodyID,omitempty" db:"banedmgbody"`            //
	BaneDamageRaceID        int64          `json:"baneDamageRaceID,omitempty" db:"banedmgrace"`            //
	BardTypeID              int64          `json:"bardTypeID,omitempty" db:"bardtype"`                     //
	BardValue               int64          `json:"bardValue,omitempty" db:"bardvalue"`                     //
	Book                    int64          `json:"book,omitempty" db:"book"`                               //
	CastTime                int64          `json:"castTime,omitempty" db:"casttime"`                       //
	CastTime2               int64          `json:"castTime2,omitempty" db:"casttime_"`                     //
	CharmFile               string         `json:"charmFile,omitempty" db:"charmfile"`                     //
	CharmFileID             string         `json:"charmFileID,omitempty" db:"charmfileid"`                 //
	ClassBit                int64          `json:"classBit,omitempty" db:"classes"`                        //
	Color                   int64          `json:"color,omitempty" db:"color"`                             //
	CombatEffectAmount      string         `json:"combatEffectAmount,omitempty" db:"combateffects"`        //
	ExtraDamageSkill        int64          `json:"extraDamageSkill,omitempty" db:"extradmgskill"`          //
	ExtraDamageAmount       int64          `json:"extraDamageAmount,omitempty" db:"extradmgamt"`           //
	Price                   int64          `json:"price,omitempty" db:"price"`                             //
	ColdResistance          int64          `json:"coldResistance,omitempty" db:"cr"`                       //
	Damage                  int64          `json:"damage,omitempty" db:"damage"`                           //
	DamageShieldAmount      int64          `json:"damageShieldAmount,omitempty" db:"damageshield"`         //
	Deity                   int64          `json:"deity,omitempty" db:"deity"`                             //
	Delay                   int64          `json:"delay,omitempty" db:"delay"`                             //
	AugDistiller            int64          `json:"augDistiller,omitempty" db:"augdistiller"`               //
	DamageOverTimeShielding int64          `json:"damageOverTimeShielding,omitempty" db:"dotshielding"`    //
	DiseaseResistance       int64          `json:"diseaseResistance,omitempty" db:"dr"`                    //
	ClickType               int64          `json:"clickType,omitempty" db:"clicktype"`                     //
	ClickLevel2             int64          `json:"clickLevel2,omitempty" db:"clicklevel2"`                 //
	ElementalDamageTypeID   int64          `json:"elementalDamageTypeID,omitempty" db:"elemdmgtype"`       //
	ElementalDamageAmount   int64          `json:"elementalDamageAmount,omitempty" db:"elemdmgamt"`        //
	Endurance               int64          `json:"endurance,omitempty" db:"endur"`                         //
	FactionAmount1          int64          `json:"factionAmount1,omitempty" db:"factionamt1"`              //
	FactionAmount2          int64          `json:"factionAmount2,omitempty" db:"factionamt2"`              //
	FactionAmount3          int64          `json:"factionAmount3,omitempty" db:"factionamt3"`              //
	FactionAmount4          int64          `json:"factionAmount4,omitempty" db:"factionamt4"`              //
	FactionModifer1         int64          `json:"factionmod1,omitempty" db:"factionmod1"`                 //
	FactionModifier2        int64          `json:"factionmod2,omitempty" db:"factionmod2"`                 //
	FactionModifer3         int64          `json:"factionmod3,omitempty" db:"factionmod3"`                 //
	FactionModifer4         int64          `json:"factionmod4,omitempty" db:"factionmod4"`                 //
	FileName                string         `json:"fileName,omitempty" db:"filename"`                       //
	FocusEffect             int64          `json:"focusEffect,omitempty" db:"focuseffect"`                 //
	FireResistance          int64          `json:"fireResistance,omitempty" db:"fr"`                       //
	FirionaVieNoDrop        int64          `json:"firionaVieNoDrop,omitempty" db:"fvnodrop"`               //
	Haste                   int64          `json:"haste,omitempty" db:"haste"`                             //
	ClickLevel              int64          `json:"clickLevel,omitempty" db:"clicklevel"`                   //
	Hitpoint                int64          `json:"hitpoint,omitempty" db:"hp"`                             //
	Regen                   int64          `json:"regen,omitempty" db:"regen"`                             //
	Icon                    int64          `json:"icon,omitempty" db:"icon"`                               //
	IDFile                  string         `json:"IDFile,omitempty" db:"idfile"`                           //
	ItemClass               int64          `json:"itemClass,omitempty" db:"itemclass"`                     //
	ItemType                int64          `json:"itemType,omitempty" db:"itemtype"`                       //
	LdonPrice               int64          `json:"ldonPrice,omitempty" db:"ldonprice"`                     //
	LdonTheme               int64          `json:"ldonTheme,omitempty" db:"ldontheme"`                     //
	LdonSold                int64          `json:"ldonSold,omitempty" db:"ldonsold"`                       //
	Light                   int64          `json:"light,omitempty" db:"light"`                             //
	Lore                    string         `json:"lore,omitempty" db:"lore"`                               //
	LoreGroup               int64          `json:"loregroup,omitempty" db:"loregroup"`                     //
	Magic                   int64          `json:"magic,omitempty" db:"magic"`                             //
	Mana                    int64          `json:"mana,omitempty" db:"mana"`                               //
	ManaRegen               int64          `json:"manaRegen,omitempty" db:"manaregen"`                     //
	EnduranceRegen          int64          `json:"enduranceRegen,omitempty" db:"enduranceregen"`           //
	Material                int64          `json:"material,omitempty" db:"material"`                       //
	HeroForgeModel          int64          `json:"heroForgeModel,omitempty" db:"herosforgemodel"`          //
	MaxCharges              int64          `json:"maxcharges,omitempty" db:"maxcharges"`                   //
	MagicResistance         int64          `json:"magicResistance,omitempty" db:"mr"`                      //
	NoDrop                  int64          `json:"noDrop,omitempty" db:"nodrop"`                           //
	NoRent                  int64          `json:"noRent,omitempty" db:"norent"`                           //
	PendingLoreFlag         int64          `json:"pendingLoreFlag,omitempty" db:"pendingloreflag"`         //
	PoisonResistance        int64          `json:"poisonResistance,omitempty" db:"pr"`                     //
	ProcRate                int64          `json:"procRate,omitempty" db:"procrate"`                       //
	RaceBit                 int64          `json:"raceBit,omitempty" db:"races"`                           //
	Range                   int64          `json:"range,omitempty" db:"range"`                             //
	RecLevel                int64          `json:"recLevel,omitempty" db:"reclevel"`                       //
	RecSkill                int64          `json:"recSkill,omitempty" db:"recskill"`                       //
	ReqLevel                int64          `json:"reqLevel,omitempty" db:"reqlevel"`                       //
	SellRate                float64        `json:"sellRate,omitempty" db:"sellrate"`                       //
	Shielding               int64          `json:"shielding,omitempty" db:"shielding"`                     //
	SizeID                  int64          `json:"sizeID,omitempty" db:"size"`                             //
	SkillModifierType       int64          `json:"skillModifierType,omitempty" db:"skillmodtype"`          //
	SkillModifierValue      int64          `json:"skillModifierValue,omitempty" db:"skillmodvalue"`        //
	SlotBit                 int64          `json:"slotBit,omitempty" db:"slots"`                           //
	ClickEffect             int64          `json:"clickEffect,omitempty" db:"clickeffect"`                 //
	SpellShield             int64          `json:"spellShield,omitempty" db:"spellshield"`                 //
	StrikeThrough           int64          `json:"strikeThrough,omitempty" db:"strikethrough"`             //
	StunResist              int64          `json:"stunResist,omitempty" db:"stunresist"`                   //
	SummonedFlag            int64          `json:"summonedFlag,omitempty" db:"summonedflag"`               //
	TradeSkills             int64          `json:"tradeSkills,omitempty" db:"tradeskills"`                 //
	Favor                   int64          `json:"favor,omitempty" db:"favor"`                             //
	Weight                  int64          `json:"weight,omitempty" db:"weight"`                           //
	Unk012                  int64          `json:"UNK012,omitempty" db:"UNK012"`                           //
	Unk013                  int64          `json:"UNK013,omitempty" db:"UNK013"`                           //
	BenefitFlag             int64          `json:"benefitFlag,omitempty" db:"benefitflag"`                 //
	Unk054                  int64          `json:"UNK054,omitempty" db:"UNK054"`                           //
	Unk059                  int64          `json:"UNK059,omitempty" db:"UNK059"`                           //
	BookType                int64          `json:"bookType,omitempty" db:"booktype"`                       //
	RecastDelay             int64          `json:"recastDelay,omitempty" db:"recastdelay"`                 //
	RecastType              int64          `json:"recastType,omitempty" db:"recasttype"`                   //
	GuildFavor              int64          `json:"guildFavor,omitempty" db:"guildfavor"`                   //
	Unk123                  int64          `json:"UNK123,omitempty" db:"UNK123"`                           //
	Unk124                  int64          `json:"UNK124,omitempty" db:"UNK124"`                           //
	Attuneable              int64          `json:"attuneable,omitempty" db:"attuneable"`                   //
	NoPet                   int64          `json:"noPet,omitempty" db:"nopet"`                             //
	Updated                 mysql.NullTime `json:"updated,omitempty" db:"updated"`                         //
	Comment                 string         `json:"comment,omitempty" db:"comment"`                         //
	Unk127                  int64          `json:"UNK127,omitempty" db:"UNK127"`                           //
	PointType               int64          `json:"pointType,omitempty" db:"pointtype"`                     //
	PotionBelt              int64          `json:"potionBelt,omitempty" db:"potionbelt"`                   //
	PotionBeltSlots         int64          `json:"potionBeltSlots,omitempty" db:"potionbeltslots"`         //
	StackSize               int64          `json:"stackSize,omitempty" db:"stacksize"`                     //
	NoTransfer              int64          `json:"noTransfer,omitempty" db:"notransfer"`                   //
	Stackable               int64          `json:"stackable,omitempty" db:"stackable"`                     //
	Unk134                  string         `json:"UNK134,omitempty" db:"UNK134"`                           //
	Unk137                  int64          `json:"UNK137,omitempty" db:"UNK137"`                           //
	ProcEffectSpellID       int64          `json:"procEffectSpellID,omitempty" db:"proceffect"`            //
	ProcType                int64          `json:"procType,omitempty" db:"proctype"`                       //
	ProcLevel2              int64          `json:"procLevel2,omitempty" db:"proclevel2"`                   //
	ProcLevel               int64          `json:"procLevel,omitempty" db:"proclevel"`                     //
	Unk142                  int64          `json:"UNK142,omitempty" db:"UNK142"`                           //
	WornEffect              int64          `json:"wornEffect,omitempty" db:"worneffect"`                   //
	WornType                int64          `json:"wornType,omitempty" db:"worntype"`                       //
	WornLevel2              int64          `json:"wornLevel2,omitempty" db:"wornlevel2"`                   //
	WornLevel               int64          `json:"wornLevel,omitempty" db:"wornlevel"`                     //
	Unk147                  int64          `json:"UNK147,omitempty" db:"UNK147"`                           //
	FocusType               int64          `json:"focusType,omitempty" db:"focustype"`                     //
	FocusLevel2             int64          `json:"focusLevel2,omitempty" db:"focuslevel2"`                 //
	FocusLevel              int64          `json:"focusLevel,omitempty" db:"focuslevel"`                   //
	Unk152                  int64          `json:"UNK152,omitempty" db:"UNK152"`                           //
	ScrollEffect            int64          `json:"scrollEffect,omitempty" db:"scrolleffect"`               //
	ScrollType              int64          `json:"scrollType,omitempty" db:"scrolltype"`                   //
	ScrollLevel2            int64          `json:"scrollLevel2,omitempty" db:"scrolllevel2"`               //
	ScrollLevel             int64          `json:"scrollLevel,omitempty" db:"scrolllevel"`                 //
	Unk157                  int64          `json:"UNK157,omitempty" db:"UNK157"`                           //
	Serialized              mysql.NullTime `json:"serialized,omitempty" db:"serialized"`                   //
	Verified                mysql.NullTime `json:"verified,omitempty" db:"verified"`                       //
	Serialization           sql.NullString `json:"serialization,omitempty" db:"serialization"`             //
	Source                  string         `json:"source,omitempty" db:"source"`                           //
	Unk033                  int64          `json:"UNK033,omitempty" db:"UNK033"`                           //
	LoreFile                string         `json:"loreFile,omitempty" db:"lorefile"`                       //
	Unk014                  int64          `json:"UNK014,omitempty" db:"UNK014"`                           //
	SaveCorruption          int64          `json:"saveCorruption,omitempty" db:"svcorruption"`             //
	SkillModifierMax        int64          `json:"skillModifierMax,omitempty" db:"skillmodmax"`            //
	Unk060                  int64          `json:"UNK060,omitempty" db:"UNK060"`                           //
	AugSlot1Unk2            int64          `json:"augSlot1Unk2,omitempty" db:"augslot1unk2"`               //
	AugSlot2Unk2            int64          `json:"augSlot2Unk2,omitempty" db:"augslot2unk2"`               //
	AugSlot3Unk2            int64          `json:"augSlot3Unk2,omitempty" db:"augslot3unk2"`               //
	AugSlot4Unk2            int64          `json:"augSlot4Unk2,omitempty" db:"augslot4unk2"`               //
	AugSlot5Unk2            int64          `json:"augSlot5Unk2,omitempty" db:"augslot5unk2"`               //
	AugSlot6Unk2            int64          `json:"augSlot6Unk2,omitempty" db:"augslot6unk2"`               //
	Unk120                  int64          `json:"UNK120,omitempty" db:"UNK120"`                           //
	Unk121                  int64          `json:"UNK121,omitempty" db:"UNK121"`                           //
	QuestItemFlag           int64          `json:"questItemFlag,omitempty" db:"questitemflag"`             //
	Unk132                  sql.NullString `json:"UNK132,omitempty" db:"UNK132"`                           //
	ClickUnk5               int64          `json:"clickUnk5,omitempty" db:"clickunk5"`                     //
	ClickUnk6               string         `json:"clickUnk6,omitempty" db:"clickunk6"`                     //
	ClickUnk7               int64          `json:"clickUnk7,omitempty" db:"clickunk7"`                     //
	ProcUnk1                int64          `json:"procUnk1,omitempty" db:"procunk1"`                       //
	ProcUnk2                int64          `json:"procUnk2,omitempty" db:"procunk2"`                       //
	ProcUnk3                int64          `json:"procUnk3,omitempty" db:"procunk3"`                       //
	ProcUnk4                int64          `json:"procUnk4,omitempty" db:"procunk4"`                       //
	ProcUnk6                string         `json:"procUnk6,omitempty" db:"procunk6"`                       //
	ProcUnk7                int64          `json:"procUnk7,omitempty" db:"procunk7"`                       //
	WornUnk1                int64          `json:"wornUnk1,omitempty" db:"wornunk1"`                       //
	WornUnk2                int64          `json:"wornUnk2,omitempty" db:"wornunk2"`                       //
	WornUnk3                int64          `json:"wornUnk3,omitempty" db:"wornunk3"`                       //
	WornUnk4                int64          `json:"wornUnk4,omitempty" db:"wornunk4"`                       //
	WornUnk5                int64          `json:"wornUnk5,omitempty" db:"wornunk5"`                       //
	WornUnk6                string         `json:"wornUnk6,omitempty" db:"wornunk6"`                       //
	WornUnk7                int64          `json:"wornUnk7,omitempty" db:"wornunk7"`                       //
	FocusUnk1               int64          `json:"focusUnk1,omitempty" db:"focusunk1"`                     //
	FocusUnk2               int64          `json:"focusUnk2,omitempty" db:"focusunk2"`                     //
	FocusUnk3               int64          `json:"focusUnk3,omitempty" db:"focusunk3"`                     //
	FocusUnk4               int64          `json:"focusUnk4,omitempty" db:"focusunk4"`                     //
	FocusUnk5               int64          `json:"focusUnk5,omitempty" db:"focusunk5"`                     //
	FocusUnk6               string         `json:"focusUnk6,omitempty" db:"focusunk6"`                     //
	FocusUnk7               int64          `json:"focusUnk7,omitempty" db:"focusunk7"`                     //
	ScrollUnk1              int64          `json:"scrollUnk1,omitempty" db:"scrollunk1"`                   //
	ScrollUnk2              int64          `json:"scrollUnk2,omitempty" db:"scrollunk2"`                   //
	ScrollUnk3              int64          `json:"scrollUnk3,omitempty" db:"scrollunk3"`                   //
	ScrollUnk4              int64          `json:"scrollUnk4,omitempty" db:"scrollunk4"`                   //
	ScrollUnk5              int64          `json:"scrollUnk5,omitempty" db:"scrollunk5"`                   //
	ScrollUnk6              string         `json:"scrollUnk6,omitempty" db:"scrollunk6"`                   //
	ScrollUnk7              int64          `json:"scrollUnk7,omitempty" db:"scrollunk7"`                   //
	Unk193                  int64          `json:"UNK193,omitempty" db:"UNK193"`                           //
	Purity                  int64          `json:"purity,omitempty" db:"purity"`                           //
	EvolvingItemID          int64          `json:"evolvingItemID,omitempty" db:"evoitem"`                  //
	EvolvingID              int64          `json:"evolvingID,omitempty" db:"evoid"`                        //
	Evolvinglevel           int64          `json:"evolvingLevel,omitempty" db:"evolvinglevel"`             //
	EvolvingMax             int64          `json:"evovlingMax,omitempty" db:"evomax"`                      //
	ClickName               string         `json:"clickName,omitempty" db:"clickname"`                     //
	ProcName                string         `json:"procName,omitempty" db:"procname"`                       //
	WornName                string         `json:"wornName,omitempty" db:"wornname"`                       //
	FocusName               string         `json:"focusName,omitempty" db:"focusname"`                     //
	ScrollName              string         `json:"scrollName,omitempty" db:"scrollname"`                   //
	DamageShieldMitigation  int64          `json:"damageShieldMitigation,omitempty" db:"dsmitigation"`     //
	HeroicStrength          int64          `json:"heroicStrength,omitempty" db:"heroic_str"`               //
	HeroicIntelligence      int64          `json:"heroicIntelligence,omitempty" db:"heroic_int"`           //
	HeroicWisdom            int64          `json:"heroicWisdom,omitempty" db:"heroic_wis"`                 //
	HeroicAgility           int64          `json:"heroicAgility,omitempty" db:"heroic_agi"`                //
	HeroicDexterity         int64          `json:"heroicDexterity,omitempty" db:"heroic_dex"`              //
	HeroicStamina           int64          `json:"heroicStamina,omitempty" db:"heroic_sta"`                //
	HeroicCharisma          int64          `json:"heroicCharisma,omitempty" db:"heroic_cha"`               //
	HeroicPoisonResistance  int64          `json:"heroicPoisonResistance,omitempty" db:"heroic_pr"`        //
	HeroicDiseaseResistance int64          `json:"heroicDiseaseResistance,omitempty" db:"heroic_dr"`       //
	HeroicFireResistance    int64          `json:"heroicFireResistance,omitempty" db:"heroic_fr"`          //
	HeroicColdResistance    int64          `json:"heroicColdResistance,omitempty" db:"heroic_cr"`          //
	HeroicMagicResistance   int64          `json:"heroicMagicResistance,omitempty" db:"heroic_mr"`         //
	HeroicSvcorrup          int64          `json:"heroicSvcorrup,omitempty" db:"heroic_svcorrup"`          //
	HealAmount              int64          `json:"healAmount,omitempty" db:"healamt"`                      //
	SpellDamage             int64          `json:"spellDamage,omitempty" db:"spelldmg"`                    //
	Clairvoyance            int64          `json:"clairvoyance,omitempty" db:"clairvoyance"`               //
	BackstabDamage          int64          `json:"backstabDamage,omitempty" db:"backstabdmg"`              //
	Created                 string         `json:"created,omitempty" db:"created"`                         //
	EliteMaterial           int64          `json:"eliteMaterial,omitempty" db:"elitematerial"`             //
	LdonSellBackRate        int64          `json:"ldonSellBackRate,omitempty" db:"ldonsellbackrate"`       //
	ScriptFileid            int64          `json:"scriptFileID,omitempty" db:"scriptfileid"`               //
	ExpendableArrow         int64          `json:"expendableArrow,omitempty" db:"expendablearrow"`         //
	PowerSourceCapacity     int64          `json:"powerSourceCapacity,omitempty" db:"powersourcecapacity"` //
	BardEffect              int64          `json:"bardEffect,omitempty" db:"bardeffect"`                   //
	BardEffectType          int64          `json:"bardEffectType,omitempty" db:"bardeffecttype"`           //
	BardLevel2              int64          `json:"bardLevel2,omitempty" db:"bardlevel2"`                   //
	BardLevel               int64          `json:"bardLevel,omitempty" db:"bardlevel"`                     //
	BardUnk1                int64          `json:"bardUnk1,omitempty" db:"bardunk1"`                       //
	BardUnk2                int64          `json:"bardUnk2,omitempty" db:"bardunk2"`                       //
	BardUnk3                int64          `json:"bardUnk3,omitempty" db:"bardunk3"`                       //
	BardUnk4                int64          `json:"bardUnk4,omitempty" db:"bardunk4"`                       //
	BardUnk5                int64          `json:"bardUnk5,omitempty" db:"bardunk5"`                       //
	BardName                string         `json:"bardName,omitempty" db:"bardname"`                       //
	BardUnk7                int64          `json:"bardUnk7,omitempty" db:"bardunk7"`                       //
	Unk214                  int64          `json:"UNK214,omitempty" db:"UNK214"`                           //
	Unk219                  int64          `json:"UNK219,omitempty" db:"UNK219"`                           //
	Unk220                  int64          `json:"UNK220,omitempty" db:"UNK220"`                           //
	Unk221                  int64          `json:"UNK221,omitempty" db:"UNK221"`                           //
	Heirloom                int64          `json:"heirloom,omitempty" db:"heirloom"`                       //
	Unk223                  int64          `json:"UNK223,omitempty" db:"UNK223"`                           //
	Unk224                  int64          `json:"UNK224,omitempty" db:"UNK224"`                           //
	Unk225                  int64          `json:"UNK225,omitempty" db:"UNK225"`                           //
	Unk226                  int64          `json:"UNK226,omitempty" db:"UNK226"`                           //
	Unk227                  int64          `json:"UNK227,omitempty" db:"UNK227"`                           //
	Unk228                  int64          `json:"UNK228,omitempty" db:"UNK228"`                           //
	Unk229                  int64          `json:"UNK229,omitempty" db:"UNK229"`                           //
	Unk230                  int64          `json:"UNK230,omitempty" db:"UNK230"`                           //
	Unk231                  int64          `json:"UNK231,omitempty" db:"UNK231"`                           //
	Unk232                  int64          `json:"UNK232,omitempty" db:"UNK232"`                           //
	Unk233                  int64          `json:"UNK233,omitempty" db:"UNK233"`                           //
	Unk234                  int64          `json:"UNK234,omitempty" db:"UNK234"`                           //
	Placeable               int64          `json:"placeable,omitempty" db:"placeable"`                     //
	Unk236                  int64          `json:"UNK236,omitempty" db:"UNK236"`                           //
	Unk237                  int64          `json:"UNK237,omitempty" db:"UNK237"`                           //
	Unk238                  int64          `json:"UNK238,omitempty" db:"UNK238"`                           //
	Unk239                  int64          `json:"UNK239,omitempty" db:"UNK239"`                           //
	Unk240                  int64          `json:"UNK240,omitempty" db:"UNK240"`                           //
	Unk241                  int64          `json:"UNK241,omitempty" db:"UNK241"`                           //
	EpicItem                int64          `json:"epicItem,omitempty" db:"epicitem"`                       //
}
