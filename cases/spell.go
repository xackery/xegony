package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//SpellRepository handles SpellRepository cases and is a gateway to storage
type SpellRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *SpellRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *SpellRepository) Get(spellID int64) (spell *model.Spell, err error) {
	if spellID == 0 {
		err = fmt.Errorf("Invalid Spell ID")
		return
	}
	spell, err = c.stor.GetSpell(spellID)
	if err = c.prepare(spell); err != nil {
		return
	}
	return
}

//Search handles logic
func (c *SpellRepository) Search(search string) (spells []*model.Spell, err error) {
	spells, err = c.stor.SearchSpell(search)
	if err != nil {
		return
	}
	for _, spell := range spells {
		if err = c.prepare(spell); err != nil {
			return
		}
	}
	return
}

//Create handles logic
func (c *SpellRepository) Create(spell *model.Spell) (err error) {
	if spell == nil {
		err = fmt.Errorf("Empty spell")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	spell.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(spell))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}
	err = c.stor.CreateSpell(spell)
	if err != nil {
		return
	}
	if err = c.prepare(spell); err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *SpellRepository) Edit(spellID int64, spell *model.Spell) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spell))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}

	err = c.stor.EditSpell(spellID, spell)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *SpellRepository) Delete(spellID int64) (err error) {
	err = c.stor.DeleteSpell(spellID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *SpellRepository) List(pageSize int64, pageNumber int64) (spells []*model.Spell, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	spells, err = c.stor.ListSpell(pageSize, pageNumber)
	if err != nil {
		return
	}
	for _, spell := range spells {
		err = c.prepare(spell)
		if err != nil {
			return
		}
	}
	return
}

//ListCount handles logic
func (c *SpellRepository) ListCount() (count int64, err error) {

	count, err = c.stor.ListSpellCount()
	if err != nil {
		return
	}
	return
}

func (c *SpellRepository) prepare(spell *model.Spell) (err error) {
	var ok bool
	spell.ClassesList = spellClassesList(spell)
	spell.LowestLevel = spellLowestLevel(spell)

	/*spell.TypeName, ok = spellTypeNames[spell.Spellgroup]
	if !ok {
		spell.TypeName = fmt.Sprintf("Uknown (%d)", spell.Spellgroup)
	}*/
	skillRepo := &SkillRepository{}
	if err = skillRepo.Initialize(c.stor); err != nil {
		err = errors.Wrap(err, "Failed to initialize spell skill")
		return
	}
	if spell.Skill, err = skillRepo.Get(spell.SkillID); err != nil {
		err = errors.Wrap(err, "Failed to get spell skill")
		return
	}
	if spell.CastTime > 0 {
		spell.CastTimeName = fmt.Sprintf("%.2f seconds", float64(spell.CastTime)/1000)
	} else {
		spell.CastTimeName = ""
	}

	if spell.RecastTime > 0 {
		spell.RecastTimeName = fmt.Sprintf("%.2f seconds", float64(spell.RecastTime)/1000)
	} else {
		spell.RecastTimeName = ""
	}

	if spell.RecoveryTime > 0 {
		spell.RecoveryTimeName = fmt.Sprintf("%.2f seconds", float64(spell.RecoveryTime)/1000)
	} else {
		spell.RecoveryTimeName = ""
	}

	spell.TargetTypeName, ok = spellTargetTypes[spell.Targettype]
	if !ok {
		spell.TargetTypeName = fmt.Sprintf("Unknown (%d)", spell.Targettype)
	}

	spell.ResistTypeName, ok = spellResistTypes[spell.ResistType]
	if !ok {
		spell.ResistTypeName = fmt.Sprintf("Unknown (%d)", spell.ResistType)
	}

	spell.BuffDurationName, ok = spellDurationFormulas[spell.BuffDurationFormula]
	if !ok {
		spell.BuffDurationName = fmt.Sprintf("Unknown (%d)", spell.BuffDurationFormula)
	}

	itemRepo := &ItemRepository{}
	if err = itemRepo.Initialize(c.stor); err != nil {
		err = errors.Wrap(err, "Failed to initialize reagent item")
		return
	}

	var item *model.Item
	var reagents = []int64{
		spell.Noexpendreagent1,
		spell.Noexpendreagent2,
		spell.Noexpendreagent3,
		spell.Noexpendreagent4,
	}
	for _, reagent := range reagents {
		if reagent > 0 {
			if item, err = itemRepo.Get(reagent); err != nil {
				err = errors.Wrap(err, "Failed to get reagent item")
				return
			}
			spell.NoExpendReagents = append(spell.NoExpendReagents, item)
		}
	}
	reagents = []int64{
		spell.Components1,
		spell.Components2,
		spell.Components3,
		spell.Components4,
	}
	for _, reagent := range reagents {
		if reagent > 0 {
			if item, err = itemRepo.Get(reagent); err != nil {
				err = errors.Wrap(err, "Failed to get reagent item")
				return
			}
			spell.Reagents = append(spell.Reagents, item)
		}
	}

	type spellEffectFields struct {
		ID        int64
		BaseValue int64
		Formula   int64
		Effect    *model.SpellEffect
	}
	var effects = []*spellEffectFields{
		&spellEffectFields{
			ID:        spell.Effectid1,
			BaseValue: spell.EffectBaseValue1,
			Formula:   spell.Formula1,
		},
		&spellEffectFields{
			ID:        spell.Effectid2,
			BaseValue: spell.EffectBaseValue2,
			Formula:   spell.Formula2,
		},
		&spellEffectFields{
			ID:        spell.Effectid3,
			BaseValue: spell.EffectBaseValue3,
			Formula:   spell.Formula3,
		},
		&spellEffectFields{
			ID:        spell.Effectid4,
			BaseValue: spell.EffectBaseValue4,
			Formula:   spell.Formula4,
		},
		&spellEffectFields{
			ID:        spell.Effectid5,
			BaseValue: spell.EffectBaseValue5,
			Formula:   spell.Formula5,
		},
		&spellEffectFields{
			ID:        spell.Effectid6,
			BaseValue: spell.EffectBaseValue6,
			Formula:   spell.Formula6,
		},
		&spellEffectFields{
			ID:        spell.Effectid7,
			BaseValue: spell.EffectBaseValue7,
			Formula:   spell.Formula7,
		},
		&spellEffectFields{
			ID:        spell.Effectid8,
			BaseValue: spell.EffectBaseValue8,
			Formula:   spell.Formula8,
		},
		&spellEffectFields{
			ID:        spell.Effectid9,
			BaseValue: spell.EffectBaseValue9,
			Formula:   spell.Formula9,
		},
		&spellEffectFields{
			ID:        spell.Effectid10,
			BaseValue: spell.EffectBaseValue10,
			Formula:   spell.Formula10,
		},
		&spellEffectFields{
			ID:        spell.Effectid11,
			BaseValue: spell.EffectBaseValue11,
			Formula:   spell.Formula11,
		},
		&spellEffectFields{
			ID:        spell.Effectid12,
			BaseValue: spell.EffectBaseValue12,
			Formula:   spell.Formula12,
		},
	}

	for _, effect := range effects {
		if effect.ID == 254 && effect.BaseValue == 0 {
			continue
		}
		spellEffect := &model.SpellEffect{}

		spellEffect, ok = spellEffects[effect.ID]
		if !ok {
			spellEffect.Name = fmt.Sprintf("Unknown (%d)", effect.ID)
		}
		spellEffect.BaseValue = effect.BaseValue

		spellEffect.FormulaName, ok = spellEffectFormulas[effect.Formula]
		if !ok {
			spellEffect.FormulaName = fmt.Sprintf("Unknown Formula (%d)", effect.Formula)
		}
		//spellEffect.Description += fmt.Sprintf(" %d", effect.BaseValue)
		spell.Effects = append(spell.Effects, spellEffect)
	}

	return
}

func (c *SpellRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
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

func (c *SpellRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "zoneID":
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

var spellTargetTypes = map[int64]string{
	0:  "Special",
	1:  "Line of Sight",
	2:  "Old PBAE (Unused)",
	3:  "Own Group",
	4:  "PBAE",
	5:  "Single",
	6:  "Self Only",
	8:  "Targeted AE",
	9:  "Animals Only",
	10: "Undead Only",
	11: "Summoned",
	13: "Life Tap",
	14: "Own Pet",
	15: "Corpse",
	16: "Plants",
	17: "Special Velious Giants",
	18: "Special Velious Dragons",
	20: "Targeted AE Life Tap",
	24: "AE Undead Only",
	25: "AE Summoned Only",
	32: "AE HateList/Casters Only",
	33: "NPC's Hate List",
	34: "Lost Dungeon Object",
	35: "Muramite",
	36: "AE PCs Only",
	37: "AE NPCs Only",
	38: "Any Summoned Pet",
	39: "Group (No Pets)",
	40: "Raid AE",
	41: "Group (Targetable)",
	42: "Directional Cone",
	43: "Group With Pets",
	44: "Beam",
	45: "Ring",
	46: "Target's Target",
	47: "Targeted Pet's Master",
	50: "Targeted AE (No Pets)",
}

var spellTypeNames = map[int64]string{
	0: "Detrimental",
	1: "Beneficial",
	2: "Beneficial (Grouped Only)",
	3: "For Testing Only",
}

var spellResistTypes = map[int64]string{
	0: "Unresistable",
	1: "Magic",
	2: "Fire",
	3: "Cold",
	4: "Poison",
	5: "Disease",
	6: "Chromatic (Lowest)",
	7: "Prismatic (Average)",
	8: "Physical",
	9: "Corruption",
}
var spellDurationFormulas = map[int64]string{
	0:    "Instant",
	1:    "Level / 2 Ticks, Up to Max",
	2:    "Level * 3/5 Ticks, Up to Max",
	3:    "Level * 3 Minutes, Up to Max Ticks",
	4:    "Max Ticks, or 5min if Max <= 0",
	5:    "Max Ticks (enforced to 1-3)",
	6:    "Level / 2 Ticks, Up to Max (Unused in Live)",
	7:    "If Max == 0, Level Ticks, otherwise Max Ticks",
	8:    "Level + 10 Ticks, Up to Max",
	9:    "Level * 2 + 10 Ticks, Up to Max",
	10:   "Level * 3 + 10 Ticks, Up to Max",
	11:   "(Level + 3) * 3 Minutes, Up to Max Ticks",
	12:   "Max Ticks",
	13:   "Max Ticks",
	14:   "Max Ticks",
	15:   "Max Ticks",
	50:   "Permanent until Effect Is Otherwise Canceled",
	51:   "Permanent as Long as Target within Range of Aura",
	3600: "If Max == 0, 6 hours, otherwise Max Ticks",
}

var spellEffectFormulas = map[int64]string{
	0:    "Base",
	1:    "(1-99) Base + Level * Formula",
	60:   "Base / 100",
	70:   "Base / 100",
	100:  "Base",
	101:  "(Level + Base) / 2",
	102:  "(Level + Base)",
	103:  "(Level + Base) * 2",
	104:  "(Level + Base) * 3",
	105:  "(Level + Base) * 4",
	106:  "(Level + Base) * 106",
	107:  "(Level / 2) + Base",
	108:  "(Level / 3) + Base",
	109:  "(Level / 4) + Base",
	110:  "(Level / 6) + Base",
	111:  "[(Level - 16) * 6] + Base",
	112:  "[(Level - 24) * 8] + Base",
	113:  "[(Level - 34) * 10] + Base",
	114:  "[(Level - 44) * 15] + Base",
	115:  "[(Level - 15) * 7] + Base",
	116:  "[(Level - 24) * 10] + Base",
	117:  "[(Level - 34) * 13] + Base",
	118:  "[(Level - 44) * 20] + Base",
	119:  "(Level / 8) + Base",
	121:  "(Level / 3) + Base",
	122:  "Base, counting down 12 per tick",
	123:  "Random from Base - Max",
	124:  "[(Level - 50) * 1] + Base",
	125:  "[(Level - 50) * 2] + Base",
	126:  "[(Level - 50) * 3] + Base",
	127:  "[(Level - 50) * 4] + Base",
	128:  "[(Level - 50) * 5] + Base",
	129:  "[(Level - 50) * 10] + Base",
	130:  "[(Level - 50) * 15] + Base",
	131:  "[(Level - 50) * 20] + Base",
	132:  "[(Level - 50) * 25] + Base",
	137:  "Base",
	138:  "Random from 0 - Base",
	139:  "[(Level - 30) / 2) + Base",
	140:  "(Level - 30) + Base",
	141:  "[(Level - 30) * 3/2) + Base",
	142:  "(Level - 30) + Base",
	143:  "(Level * 3/4) + Base",
	201:  "Max",
	203:  "Max",
	1001: "(1001-1999) Base - 12 * (Formula - 1000) Per Tick",
	2001: "(2001-2650) Base * Level * (Formula - 2000) Per Tick",
}

var spellCategoryNames = map[int64]string{
	-99: "NPC",
	-1:  "AA Procs",
	0:   "Unspecified",
	1:   "Direct Damage [Magic]",
	2:   "Direct Damage [Undead]",
	3:   "Direct Damage [Summoned]",
	4:   "Direct Damage [Life Taps]",
	5:   "Direct Damage [Plant]",
	6:   "Direct Damage [Velious Races]",
	7:   "Damage over Time [Magic]",
	8:   "Damage over Time [Undead]",
	9:   "Damage over Time [Life Taps]",
	10:  "Targeted Area of Effect Damage",
	11:  "Point Blank Area of Effect Damage",
	12:  "Area of Effect Rain",
	13:  "Direct Damage [Bolt]",
	14:  "Stun [Targeted Area of Effect]",
	15:  "Stun [Targeted]",
	16:  "Stun [Point Blank Area of Effect]",
	17:  "Drains [Health/Mana]",
	18:  "Drains [Stats]",
	19:  "Contact Innates",
	20:  "Heal [Instant]",
	21:  "Heal [Duration]",
	22:  "Group Heal [Instant]",
	23:  "Group Heal [Duration]",
	24:  "Regeneration [Single]",
	25:  "Regeneration [Group]",
	26:  "Heal [Own Pet]",
	27:  "Resurrect",
	28:  "Necromancer Life Transfer",
	29:  "Cure [Poison]",
	30:  "Health Buffs [Single]",
	32:  "AC Buff [Single]",
	34:  "Hate Mod Buffs",
	35:  "Haste [Single]",
	36:  "Haste [Pet]",
	37:  "Haste [Group]",
	38:  "Slow [Single]",
	39:  "Slow [Targeted Area]",
	40:  "Cannabalize",
	41:  "Move Speed [Single]",
	42:  "Move Speed [Group]",
	43:  "Wolf Form",
	44:  "Move Speed [Pet]",
	45:  "Illusion [Self]",
	46:  "Lich",
	47:  "Bear Form",
	48:  "Tree Form",
	49:  "Dead Man Floating",
	50:  "Root",
	51:  "Summon Pet",
	52:  "Summon Corpse",
	53:  "Sense Undead",
	54:  "Invulnerability",
	55:  "Gate [Combat Portal]",
	56:  "Gate [Self Gates]",
	58:  "Translocate",
	59:  "Shadow Step",
	60:  "Enchant Item",
	61:  "Summon [Misc Item]",
	62:  "Fear",
	63:  "Fear [Animal]",
	64:  "Fear [Undead]",
	65:  "Damage Shield [Single]",
	66:  "Damage Shield [Group]",
	67:  "Mark Of Karn",
	68:  "Damage Shield [Self]",
	69:  "Resist Debuffs",
	70:  "Resist Buffs",
	71:  "BST Pet Buffs",
	72:  "Summon Familiar",
	73:  "STR Buff",
	74:  "DEX Buff",
	75:  "AGI Buff",
	76:  "STA Buff",
	77:  "INT Buff",
	78:  "CHA Buff",
	79:  "Stat Debuffs",
	80:  "Invisible Undead",
	81:  "Invisible Animals",
	82:  "Invisibility",
	83:  "Absorb Damage",
	84:  "Casting Level Buffs",
	85:  "Clarity Line",
	86:  "Max Mana Buffs",
	87:  "Drain Mana",
	88:  "Mana Transfer",
	89:  "Instant Gain Mana",
	90:  "Lower Hate [Jolt]",
	91:  "Increase Archery",
	92:  "Attack Buff",
	93:  "Vision",
	94:  "Water Breathing",
	95:  "Improve Faction",
	96:  "Charm",
	97:  "Dispell",
	98:  "Lull",
	99:  "Mesmerise",
	100: "Spell Focus Items",
	101: "Snare [single]",
	102: "Snare [Area of Effect]",
	105: "Feign Death",
	106: "Identify",
	107: "Reclaim Energy",
	108: "Find Corpse",
	109: "Summon Player",
	110: "Spell Shield",
	112: "Blindness",
	113: "Levitation",
	114: "Extinguish Fatigue",
	115: "Death Pact",
	116: "Memory Blur",
	118: "Height",
	119: "Add Hate",
	120: "Iron Maiden",
	121: "Focus Spells",
	122: "Melee Guard",
	125: "Direct Damage [Fire]",
	126: "Direct Damage [Ice]",
	127: "Direct Damage [Poison]",
	128: "Direct Damage [Disease]",
	129: "Damage over Time [Fire]",
	130: "Damage over Time [Ice]",
	131: "Damage over Time [Poison]",
	132: "Damage over Time [Disease]",
	133: "INT Caster Chest Opening",
	134: "INT Caster Chest Trap Appraisal",
	135: "INT Caster Chest Trap Disarm",
	136: "WIS Caster Chest Trap Disarm",
	137: "WIS Caster Chest Trap Appraisal",
	138: "WIS Caster Chest Opening",
	140: "Destroy [Undead]",
	141: "Destroy [Summoned]",
	142: "Targeted Area of Effect [Fire]",
	143: "Targeted Area of Effect [Ice]",
	146: "Point Blank Area of Effect [Fire]",
	147: "Point Blank Area of Effect [Ice]",
	150: "Rain [Fire]",
	151: "Rain [Ice]",
	152: "Rain [Poison]",
	154: "Fear Song",
	155: "Fast Heals",
	156: "Mana to Health",
	157: "Pet Siphons",
	159: "Cure [Disease]",
	160: "Cure [Curse]",
	161: "Cure [Multiple]",
	162: "Cure [Blind]",
	163: "Group Cure [Multiple]",
	164: "Misc Effects",
	165: "Shielding",
	166: "PAL/RNG/BST Health Buffs",
	167: "Symbols",
	168: "Aegolism Line",
	169: "Paladin AC Buffs",
	170: "Spell Damage Mitigate",
	171: "Spell/Melee Block",
	172: "Spell Reflect",
	173: "Hybrid AC Buffs",
	174: "Health/Mana Regeneration",
	175: "Aggro Decreasers",
	200: "Misc Spells",
	201: "Disciplines",
	202: "Melee Haste",
	203: "Area of Effect Slow",
	204: "Summon Air Pet",
	205: "Summon Water Pet",
	206: "Summon Fire Pet",
	207: "Summon Earth Pet",
	208: "Summon Monster Pet",
	209: "Transport [Antonica]",
	210: "Transport [Odus]",
	211: "Transport [Faydwer]",
	212: "Transport [Kunark]",
	213: "Transport [Velious]",
	214: "Transport [Luclin]",
	215: "Transport [Planes]",
	216: "Transport [Gates/Omens]",
	217: "Summon [Weapon]",
	218: "Summon [Focus]",
	219: "Summon [Food/Drink]",
	220: "Summon [Armor]",
	999: "AA/Abilities",
}

//LowestLevel returns the lowest level a character can mem this
func spellLowestLevel(s *model.Spell) int64 {
	lowestClass := int64(255)
	if s.Classes1 < lowestClass {
		lowestClass = s.Classes1
	}
	if s.Classes2 < lowestClass {
		lowestClass = s.Classes2
	}
	if s.Classes3 < lowestClass {
		lowestClass = s.Classes3
	}
	if s.Classes4 < lowestClass {
		lowestClass = s.Classes4
	}
	if s.Classes5 < lowestClass {
		lowestClass = s.Classes5
	}
	if s.Classes6 < lowestClass {
		lowestClass = s.Classes6
	}
	if s.Classes7 < lowestClass {
		lowestClass = s.Classes7
	}
	if s.Classes8 < lowestClass {
		lowestClass = s.Classes8
	}
	if s.Classes9 < lowestClass {
		lowestClass = s.Classes9
	}
	if s.Classes10 < lowestClass {
		lowestClass = s.Classes10
	}
	if s.Classes11 < lowestClass {
		lowestClass = s.Classes11
	}
	if s.Classes12 < lowestClass {
		lowestClass = s.Classes12
	}
	if s.Classes13 < lowestClass {
		lowestClass = s.Classes13
	}
	return lowestClass
}

//ClassesList returns a list of human readable classes
func spellClassesList(s *model.Spell) string {
	classes := ""

	if s.Classes1 > 0 && s.Classes1 < 255 {
		classes += fmt.Sprintf("%s (%d), ", className(1), s.Classes1)
	}
	if s.Classes2 > 0 && s.Classes2 < 255 {
		classes += fmt.Sprintf("%s (%d), ", className(2), s.Classes2)
	}
	if s.Classes3 > 0 && s.Classes3 < 255 {
		classes += fmt.Sprintf("%s (%d), ", className(3), s.Classes3)
	}
	if s.Classes4 > 0 && s.Classes4 < 255 {
		classes += fmt.Sprintf("%s (%d), ", className(4), s.Classes4)
	}
	if s.Classes5 > 0 && s.Classes5 < 255 {
		classes += fmt.Sprintf("%s (%d), ", className(5), s.Classes5)
	}
	if s.Classes6 > 0 && s.Classes6 < 255 {
		classes += fmt.Sprintf("%s (%d), ", className(6), s.Classes6)
	}
	if s.Classes7 > 0 && s.Classes7 < 255 {
		classes += fmt.Sprintf("%s (%d), ", className(7), s.Classes7)
	}
	if s.Classes8 > 0 && s.Classes8 < 255 {
		classes += fmt.Sprintf("%s (%d), ", className(8), s.Classes8)
	}
	if s.Classes9 > 0 && s.Classes9 < 255 {
		classes += fmt.Sprintf("%s (%d), ", className(9), s.Classes9)
	}
	if s.Classes10 > 0 && s.Classes10 < 255 {
		classes += fmt.Sprintf("%s (%d), ", className(10), s.Classes10)
	}
	if s.Classes11 > 0 && s.Classes11 < 255 {
		classes += fmt.Sprintf("%s (%d), ", className(11), s.Classes11)
	}
	if s.Classes12 > 0 && s.Classes12 < 255 {
		classes += fmt.Sprintf("%s (%d), ", className(12), s.Classes12)
	}
	if s.Classes13 > 0 && s.Classes13 < 255 {
		classes += fmt.Sprintf("%s (%d), ", className(13), s.Classes13)
	}
	if len(classes) > 3 {
		classes = classes[0 : len(classes)-2]
	}
	return classes
}

var spellEffects = map[int64]*model.SpellEffect{
	0: &model.SpellEffect{
		ID:   0,
		Type: 1,
		Name: "Current HP",
	},
	1: &model.SpellEffect{
		ID:   1,
		Type: 1,
		Name: "Armor Class",
	},
	2: &model.SpellEffect{
		ID:   2,
		Type: 1,
		Name: "Attack Rating",
	},
	3: &model.SpellEffect{
		ID:   3,
		Type: 1,
		Name: "Movement Speed",
	},
	4: &model.SpellEffect{
		ID:   4,
		Type: 1,
		Name: "Strength",
	},
	5: &model.SpellEffect{
		ID:   5,
		Type: 1,
		Name: "Dexterity",
	},
	6: &model.SpellEffect{
		ID:   6,
		Type: 1,
		Name: "Agility",
	},
	7: &model.SpellEffect{
		ID:   7,
		Type: 1,
		Name: "Stamina",
	},
	8: &model.SpellEffect{
		ID:   8,
		Type: 1,
		Name: "Intelligence",
	},
	9: &model.SpellEffect{
		ID:   9,
		Type: 1,
		Name: "Wisdom",
	},
	10: &model.SpellEffect{
		ID:   10,
		Type: 1,
		Name: "Charisma",
	},
	11: &model.SpellEffect{
		ID:   11,
		Type: 1,
		Name: "Attack Speed",
	},
	12: &model.SpellEffect{
		ID:   12,
		Type: 0,
		Name: "Invisibility for a Random Duration",
	},
	13: &model.SpellEffect{
		ID:   13,
		Type: 0,
		Name: "See Invisible",
	},
	14: &model.SpellEffect{
		ID:   14,
		Type: 0,
		Name: "Water Breathing",
	},
	15: &model.SpellEffect{
		ID:   15,
		Type: 1,
		Name: "Current Mana",
	},
	16: &model.SpellEffect{
		ID:   16,
		Type: 0,
		Name: "NPC Frenzy Radius (Not Used)",
	},
	17: &model.SpellEffect{
		ID:   17,
		Type: 0,
		Name: "NPC Awareness (Not Used)",
	},
	18: &model.SpellEffect{
		ID:   18,
		Type: 2,
		Name: "Pacify",
	},
	19: &model.SpellEffect{
		ID:   19,
		Type: 1,
		Name: "Temporary Standing with Target NPC's Faction",
	},
	20: &model.SpellEffect{
		ID:   20,
		Type: 0,
		Name: "Blindness",
	},
	21: &model.SpellEffect{
		ID:   21,
		Type: 0,
		Name: "Stun Targets up to Level",
	},
	22: &model.SpellEffect{
		ID:   22,
		Type: 0,
		Name: "Charm Targets up to Level",
	},
	23: &model.SpellEffect{
		ID:   23,
		Type: 0,
		Name: "Fear Targets up to Level",
	},
	24: &model.SpellEffect{
		ID:   24,
		Type: 1,
		Name: "Stamina (No Longer Used)",
	},
	25: &model.SpellEffect{
		ID:   25,
		Type: 0,
		Name: "Bind Respawn Point",
	},
	26: &model.SpellEffect{
		ID:   26,
		Type: 0,
		Name: "Return to Respawn Point",
	},
	27: &model.SpellEffect{
		ID:   27,
		Type: 0,
		Name: "Attempt to Remove Magical Effect",
	},
	28: &model.SpellEffect{
		ID:   28,
		Type: 2,
		Name: "Invisibility to the Undead for a Random Duration",
	},
	29: &model.SpellEffect{
		ID:   29,
		Type: 2,
		Name: "Invisibility to Animals for a Random Duration",
	},
	30: &model.SpellEffect{
		ID:   30,
		Type: 0,
		Name: "NPC Frenzy Radius",
	},
	31: &model.SpellEffect{
		ID:   31,
		Type: 0,
		Name: "Mesmerize",
	},
	32: &model.SpellEffect{
		ID:   32,
		Type: 0,
		Name: "Summon Item:",
	},
	33: &model.SpellEffect{
		ID:   33,
		Type: 0,
		Name: "Summon Pet:",
	},
	34: &model.SpellEffect{
		ID:   34,
		Type: 0,
		Name: "Confuse (Not Used)",
	},
	35: &model.SpellEffect{
		ID:   35,
		Type: 1,
		Name: "Disease Counter",
	},
	36: &model.SpellEffect{
		ID:   36,
		Type: 1,
		Name: "Poison Counter",
	},
	37: &model.SpellEffect{
		ID:   37,
		Type: 2,
		Name: "Detect Hostile (Not Used)",
	},
	38: &model.SpellEffect{
		ID:   38,
		Type: 2,
		Name: "Detect Magic (Not Used)",
	},
	39: &model.SpellEffect{
		ID:   39,
		Type: 2,
		Name: "Detect Poison (Not Used)",
	},
	40: &model.SpellEffect{
		ID:   40,
		Type: 2,
		Name: "Temporary Invulnerability",
	},
	41: &model.SpellEffect{
		ID:   41,
		Type: 2,
		Name: "Destroy Target (With No Credit)",
	},
	42: &model.SpellEffect{
		ID:   42,
		Type: 2,
		Name: "Random Teleport within Spell's Range",
	},
	43: &model.SpellEffect{
		ID:   43,
		Type: 1,
		Name: "Crippling Blow Chance",
	},
	44: &model.SpellEffect{
		ID:   44,
		Type: 2,
		Name: "Lycanthropy",
	},
	45: &model.SpellEffect{
		ID:   45,
		Type: 2,
		Name: "Vampirism (Not Used)",
	},
	46: &model.SpellEffect{
		ID:   46,
		Type: 1,
		Name: "Resistance to Fire",
	},
	47: &model.SpellEffect{
		ID:   47,
		Type: 1,
		Name: "Resistance to Cold",
	},
	48: &model.SpellEffect{
		ID:   48,
		Type: 1,
		Name: "Resistance to Poison",
	},
	49: &model.SpellEffect{
		ID:   49,
		Type: 1,
		Name: "Resistance to Disease",
	},
	50: &model.SpellEffect{
		ID:   50,
		Type: 1,
		Name: "Resistance to Magic",
	},
	51: &model.SpellEffect{
		ID:   51,
		Type: 2,
		Name: "Detect Traps (Not Used)",
	},
	52: &model.SpellEffect{
		ID:   52,
		Type: 2,
		Name: "Locate Nearest Undead",
	},
	53: &model.SpellEffect{
		ID:   53,
		Type: 2,
		Name: "Locate Nearest Summoned",
	},
	54: &model.SpellEffect{
		ID:   54,
		Type: 2,
		Name: "Locate Nearest Animal",
	},
	55: &model.SpellEffect{
		ID:   55,
		Type: 1,
		Name: "Damage Absorption",
	},
	56: &model.SpellEffect{
		ID:   56,
		Type: 2,
		Name: "Spin to Face North",
	},
	57: &model.SpellEffect{
		ID:   57,
		Type: 0,
		Name: "Levitation",
	},
	58: &model.SpellEffect{
		ID:   58,
		Type: 0,
		Name: "Illusion:",
	},
	59: &model.SpellEffect{
		ID:   59,
		Type: 1,
		Name: "Damage Shield",
	},
	60: &model.SpellEffect{
		ID:   60,
		Type: 2,
		Name: "Transfer Item (Not Used)",
	},
	61: &model.SpellEffect{
		ID:   61,
		Type: 2,
		Name: "Identify Item",
	},
	62: &model.SpellEffect{
		ID:   62,
		Type: 2,
		Name: "ItemID (Not Used)",
	},
	63: &model.SpellEffect{
		ID:   63,
		Type: 0,
		Name: "Attempt to Wipe Hate List",
	},
	64: &model.SpellEffect{
		ID:   64,
		Type: 2,
		Name: "Spinning Stun",
	},
	65: &model.SpellEffect{
		ID:   65,
		Type: 0,
		Name: "Infravision (Heat Vision)",
	},
	66: &model.SpellEffect{
		ID:   66,
		Type: 0,
		Name: "Ultravision (Night Vision)",
	},
	67: &model.SpellEffect{
		ID:   67,
		Type: 0,
		Name: "Summon and Take Control of Seeing Eye",
	},
	68: &model.SpellEffect{
		ID:   68,
		Type: 2,
		Name: "Destroy Pet and Reclaim Some Mana",
	},
	69: &model.SpellEffect{
		ID:   69,
		Type: 1,
		Name: "Maximum HP",
	},
	70: &model.SpellEffect{
		ID:   70,
		Type: 2,
		Name: "Corpse Bomb (Not Used)",
	},
	71: &model.SpellEffect{
		ID:   71,
		Type: 0,
		Name: "Summon Necromancer Pet:",
	},
	72: &model.SpellEffect{
		ID:   72,
		Type: 2,
		Name: "Preserve Corpse (Not Used)",
	},
	73: &model.SpellEffect{
		ID:   73,
		Type: 2,
		Name: "Bind Sight to Target",
	},
	74: &model.SpellEffect{
		ID:   74,
		Type: 0,
		Name: "Feign Death",
	},
	75: &model.SpellEffect{
		ID:   75,
		Type: 2,
		Name: "Voice Transfer to Target",
	},
	76: &model.SpellEffect{
		ID:   76,
		Type: 2,
		Name: "Add Hostile Proximity Alarm",
	},
	77: &model.SpellEffect{
		ID:   77,
		Type: 2,
		Name: "Locate Nearest Corpse (Optionally of Target)",
	},
	78: &model.SpellEffect{
		ID:   78,
		Type: 1,
		Name: "Chance to Absorb Magical Attack",
	},
	79: &model.SpellEffect{
		ID:   79,
		Type: 1,
		Name: "Current HP",
	},
	80: &model.SpellEffect{
		ID:   80,
		Type: 2,
		Name: "Enchant Light (Not Used)",
	},
	81: &model.SpellEffect{
		ID:   81,
		Type: 0,
		Name: "Revive with Experience Gain",
	},
	82: &model.SpellEffect{
		ID:   82,
		Type: 2,
		Name: "Summon Player to Self",
	},
	83: &model.SpellEffect{
		ID:   83,
		Type: 0,
		Name: "Teleport To:",
	},
	84: &model.SpellEffect{
		ID:   84,
		Type: 0,
		Name: "Toss Up",
	},
	85: &model.SpellEffect{
		ID:   85,
		Type: 0,
		Name: "Add Melee Proc:",
	},
	86: &model.SpellEffect{
		ID:   86,
		Type: 0,
		Name: "NPC Reaction Radius",
	},
	87: &model.SpellEffect{
		ID:   87,
		Type: 1,
		Name: "Vision Magnification (Adjust Field of View)",
	},
	88: &model.SpellEffect{
		ID:   88,
		Type: 0,
		Name: "Evacuate To:",
	},
	89: &model.SpellEffect{
		ID:   89,
		Type: 1,
		Name: "Physical Size",
	},
	90: &model.SpellEffect{
		ID:   90,
		Type: 2,
		Name: "Cloak (Not Used)",
	},
	91: &model.SpellEffect{
		ID:   91,
		Type: 0,
		Name: "Summon Corpse",
	},
	92: &model.SpellEffect{
		ID:   92,
		Type: 1,
		Name: "Hate with Target",
	},
	93: &model.SpellEffect{
		ID:   93,
		Type: 0,
		Name: "Cancel Adverse Weather",
	},
	94: &model.SpellEffect{
		ID:   94,
		Type: 2,
		Name: "Limitation: Drop Spell In Combat",
	},
	95: &model.SpellEffect{
		ID:   95,
		Type: 2,
		Name: "Sacrifice Target Character with Confirmation",
	},
	96: &model.SpellEffect{
		ID:   96,
		Type: 0,
		Name: "Silence All Spellcasting",
	},
	97: &model.SpellEffect{
		ID:   97,
		Type: 1,
		Name: "Maximum Mana",
	},
	98: &model.SpellEffect{
		ID:   98,
		Type: 1,
		Name: "Attack Speed (V2)",
	},
	99: &model.SpellEffect{
		ID:   99,
		Type: 1,
		Name: "Movement Speed",
	},
	100: &model.SpellEffect{
		ID:   100,
		Type: 1,
		Name: "Current HP Over Time",
	},
	101: &model.SpellEffect{
		ID:   101,
		Type: 0,
		Name: "Complete Heal",
	},
	102: &model.SpellEffect{
		ID:   102,
		Type: 0,
		Name: "Immunity to Fear",
	},
	103: &model.SpellEffect{
		ID:   103,
		Type: 2,
		Name: "Summon Current Pet to Self",
	},
	104: &model.SpellEffect{
		ID:   104,
		Type: 0,
		Name: "Translocate",
	},
	105: &model.SpellEffect{
		ID:   105,
		Type: 0,
		Name: "Prevent Gating to Respawn Point",
	},
	106: &model.SpellEffect{
		ID:   106,
		Type: 0,
		Name: "Summon Beastlord Pet:",
	},
	107: &model.SpellEffect{
		ID:   107,
		Type: 1,
		Name: "NPC Level",
	},
	108: &model.SpellEffect{
		ID:   108,
		Type: 0,
		Name: "Summon Familiar:",
	},
	109: &model.SpellEffect{
		ID:   109,
		Type: 0,
		Name: "Summon Item Into Bag:",
	},
	110: &model.SpellEffect{
		ID:   110,
		Type: 1,
		Name: "Archery (Not Used)",
	},
	111: &model.SpellEffect{
		ID:   111,
		Type: 1,
		Name: "All Magical Resistances",
	},
	112: &model.SpellEffect{
		ID:   112,
		Type: 1,
		Name: "Effective Spellcasting Level",
	},
	113: &model.SpellEffect{
		ID:   113,
		Type: 0,
		Name: "Summon Mount:",
	},
	114: &model.SpellEffect{
		ID:   114,
		Type: 1,
		Name: "Hate Generation Modifier",
	},
	115: &model.SpellEffect{
		ID:   115,
		Type: 0,
		Name: "Satisfy Hunger and Thirst",
	},
	116: &model.SpellEffect{
		ID:   116,
		Type: 1,
		Name: "Curse Counter",
	},
	117: &model.SpellEffect{
		ID:   117,
		Type: 2,
		Name: "Allow Weapons to Hit Magical Targets",
	},
	118: &model.SpellEffect{
		ID:   118,
		Type: 1,
		Name: "Performance Amplification",
	},
	119: &model.SpellEffect{
		ID:   119,
		Type: 1,
		Name: "Attack Speed V3",
	},
	120: &model.SpellEffect{
		ID:   120,
		Type: 1,
		Name: "Healing Modifier",
	},
	121: &model.SpellEffect{
		ID:   121,
		Type: 1,
		Name: "Reverse (Healing) Damage Shield",
	},
	122: &model.SpellEffect{
		ID:   122,
		Type: 1,
		Name: "Skill Reduction (Not Used)",
	},
	123: &model.SpellEffect{
		ID:   123,
		Type: 0,
		Name: "Screech",
	},
	124: &model.SpellEffect{
		ID:   124,
		Type: 1,
		Name: "Damage Modifier (Focus)",
	},
	125: &model.SpellEffect{
		ID:   125,
		Type: 1,
		Name: "Healing Modifier (Focus)",
	},
	126: &model.SpellEffect{
		ID:   126,
		Type: 1,
		Name: "Spell Resistance (Focus)",
	},
	127: &model.SpellEffect{
		ID:   127,
		Type: 1,
		Name: "Spellcasting Speed (Focus)",
	},
	128: &model.SpellEffect{
		ID:   128,
		Type: 1,
		Name: "Spell Durations (Focus)",
	},
	129: &model.SpellEffect{
		ID:   129,
		Type: 1,
		Name: "Spellcasting Range (Focus)",
	},
	130: &model.SpellEffect{
		ID:   130,
		Type: 1,
		Name: "Spellcasting Hate Modifier (Focus)",
	},
	131: &model.SpellEffect{
		ID:   131,
		Type: 1,
		Name: "Chance to Conserve Reagents (Focus)",
	},
	132: &model.SpellEffect{
		ID:   132,
		Type: 1,
		Name: "Spellcasting Mana Cost (Focus)",
	},
	133: &model.SpellEffect{
		ID:   133,
		Type: 1,
		Name: "Stun Time Modifier (Focus)",
	},
	134: &model.SpellEffect{
		ID:   134,
		Type: 0,
		Name: "Limit to Maximum Spell Level:",
	},
	135: &model.SpellEffect{
		ID:   135,
		Type: 0,
		Name: "Limit to Resistance Type:",
	},
	136: &model.SpellEffect{
		ID:   136,
		Type: 0,
		Name: "Limit to Target Type:",
	},
	137: &model.SpellEffect{
		ID:   137,
		Type: 0,
		Name: "Limit to Effect:",
	},
	138: &model.SpellEffect{
		ID:   138,
		Type: 0,
		Name: "Limit to Spell Type:",
	},
	139: &model.SpellEffect{
		ID:   139,
		Type: 0,
		Name: "Limit to Spell:",
	},
	140: &model.SpellEffect{
		ID:   140,
		Type: 0,
		Name: "Limit to Minimum Spell Duration:",
	},
	141: &model.SpellEffect{
		ID:   141,
		Type: 0,
		Name: "Limit to Instant Spells",
	},
	142: &model.SpellEffect{
		ID:   142,
		Type: 0,
		Name: "Limit to Minimum Spell Level:",
	},
	143: &model.SpellEffect{
		ID:   143,
		Type: 0,
		Name: "Limit to Minimum Casting Time:",
	},
	144: &model.SpellEffect{
		ID:   144,
		Type: 0,
		Name: "Limit to Maximum Casting Time:",
	},
	145: &model.SpellEffect{
		ID:   145,
		Type: 0,
		Name: "Teleport (V2):",
	},
	146: &model.SpellEffect{
		ID:   146,
		Type: 1,
		Name: "Resistance to Electricity (Not Used)",
	},
	147: &model.SpellEffect{
		ID:   147,
		Type: 1,
		Name: "Current HP by Percentage of Maximum:",
	},
	148: &model.SpellEffect{
		ID:   148,
		Type: 0,
		Name: "Stacking - Block",
	},
	149: &model.SpellEffect{
		ID:   149,
		Type: 0,
		Name: "Stacking - Overwrite",
	},
	150: &model.SpellEffect{
		ID:   150,
		Type: 1,
		Name: "Chance to Save from Death",
	},
	151: &model.SpellEffect{
		ID:   151,
		Type: 0,
		Name: "Suspend Pet",
	},
	152: &model.SpellEffect{
		ID:   152,
		Type: 0,
		Name: "Summon Temporary Pets:",
	},
	153: &model.SpellEffect{
		ID:   153,
		Type: 0,
		Name: "Balance HP Across Group",
	},
	154: &model.SpellEffect{
		ID:   154,
		Type: 0,
		Name: "Attempt to Remove Detrimental Effect",
	},
	155: &model.SpellEffect{
		ID:   155,
		Type: 1,
		Name: "Spell Critical Damage",
	},
	156: &model.SpellEffect{
		ID:   156,
		Type: 2,
		Name: "Illusion: Target",
	},
	157: &model.SpellEffect{
		ID:   157,
		Type: 1,
		Name: "Spell Damage Shield",
	},
	158: &model.SpellEffect{
		ID:   158,
		Type: 1,
		Name: "Chance to Reflect Spells",
	},
	159: &model.SpellEffect{
		ID:   159,
		Type: 1,
		Name: "All Stats",
	},
	160: &model.SpellEffect{
		ID:   160,
		Type: 1,
		Name: "Drunkenness (Not Used)",
	},
	161: &model.SpellEffect{
		ID:   161,
		Type: 1,
		Name: "Spell Damage Mitigation",
	},
	162: &model.SpellEffect{
		ID:   162,
		Type: 1,
		Name: "Melee Damage Mitigation",
	},
	163: &model.SpellEffect{
		ID:   163,
		Type: 0,
		Name: "Block All Attacks",
	},
	164: &model.SpellEffect{
		ID:   164,
		Type: 0,
		Name: "Examine LDoN Chest for Traps",
	},
	165: &model.SpellEffect{
		ID:   165,
		Type: 0,
		Name: "Disarm LDoN Trap",
	},
	166: &model.SpellEffect{
		ID:   166,
		Type: 0,
		Name: "Unlock LDoN Chest",
	},
	167: &model.SpellEffect{
		ID:   167,
		Type: 1,
		Name: "Pet Power",
	},
	168: &model.SpellEffect{
		ID:   168,
		Type: 1,
		Name: "Melee Damage Mitigation V2",
	},
	169: &model.SpellEffect{
		ID:   169,
		Type: 1,
		Name: "Critical Hit Chance",
	},
	170: &model.SpellEffect{
		ID:   170,
		Type: 1,
		Name: "Spell Critical Chance",
	},
	171: &model.SpellEffect{
		ID:   171,
		Type: 1,
		Name: "Crippling Blow Chance",
	},
	172: &model.SpellEffect{
		ID:   172,
		Type: 1,
		Name: "Melee Avoidance",
	},
	173: &model.SpellEffect{
		ID:   173,
		Type: 1,
		Name: "Chance to Riposte",
	},
	174: &model.SpellEffect{
		ID:   174,
		Type: 1,
		Name: "Chance to Dodge",
	},
	175: &model.SpellEffect{
		ID:   175,
		Type: 1,
		Name: "Chance to Parry",
	},
	176: &model.SpellEffect{
		ID:   176,
		Type: 1,
		Name: "Offhand Attack Chance",
	},
	177: &model.SpellEffect{
		ID:   177,
		Type: 1,
		Name: "Double Attack Chance",
	},
	178: &model.SpellEffect{
		ID:   178,
		Type: 0,
		Name: "Lifetap from Melee Attacks",
	},
	179: &model.SpellEffect{
		ID:   179,
		Type: 1,
		Name: "All Instrument Effectiveness",
	},
	180: &model.SpellEffect{
		ID:   180,
		Type: 1,
		Name: "Chance to Resist Spells",
	},
	181: &model.SpellEffect{
		ID:   181,
		Type: 1,
		Name: "Chance to Resist Fear",
	},
	182: &model.SpellEffect{
		ID:   182,
		Type: 1,
		Name: "Attack Speed V4",
	},
	183: &model.SpellEffect{
		ID:   183,
		Type: 1,
		Name: "Skill Checks",
	},
	184: &model.SpellEffect{
		ID:   184,
		Type: 1,
		Name: "Hit Chance",
	},
	185: &model.SpellEffect{
		ID:   185,
		Type: 1,
		Name: "Damage Modifier",
	},
	186: &model.SpellEffect{
		ID:   186,
		Type: 1,
		Name: "Minimum Damage Modifier",
	},
	187: &model.SpellEffect{
		ID:   187,
		Type: 0,
		Name: "Balance Mana Across Group",
	},
	188: &model.SpellEffect{
		ID:   188,
		Type: 1,
		Name: "Chance to Block",
	},
	189: &model.SpellEffect{
		ID:   189,
		Type: 1,
		Name: "Current Endurance",
	},
	190: &model.SpellEffect{
		ID:   190,
		Type: 1,
		Name: "Maximum Endurance",
	},
	191: &model.SpellEffect{
		ID:   191,
		Type: 0,
		Name: "Prevent Melee Attacks",
	},
	192: &model.SpellEffect{
		ID:   192,
		Type: 1,
		Name: "Hate with Target",
	},
	193: &model.SpellEffect{
		ID:   193,
		Type: 0,
		Name: "Skill Attack:",
	},
	194: &model.SpellEffect{
		ID:   194,
		Type: 0,
		Name: "Remove Self from Hate List",
	},
	195: &model.SpellEffect{
		ID:   195,
		Type: 1,
		Name: "Resistance to Stun",
	},
	196: &model.SpellEffect{
		ID:   196,
		Type: 1,
		Name: "Chance to Strike Through",
	},
	197: &model.SpellEffect{
		ID:   197,
		Type: 1,
		Name: "Skill Damage Taken:",
	},
	198: &model.SpellEffect{
		ID:   198,
		Type: 1,
		Name: "Current Endurance",
	},
	199: &model.SpellEffect{
		ID:   199,
		Type: 0,
		Name: "Attempt to Taunt the Target",
	},
	200: &model.SpellEffect{
		ID:   200,
		Type: 1,
		Name: "Chance of Melee Attack Procs",
	},
	201: &model.SpellEffect{
		ID:   201,
		Type: 1,
		Name: "Chance of Ranged Attack Procs",
	},
	202: &model.SpellEffect{
		ID:   202,
		Type: 0,
		Name: "Illusion:",
	},
	203: &model.SpellEffect{
		ID:   203,
		Type: 0,
		Name: "Mass Group Next Buff",
	},
	204: &model.SpellEffect{
		ID:   204,
		Type: 0,
		Name: "Group Immunity to Fear",
	},
	205: &model.SpellEffect{
		ID:   205,
		Type: 0,
		Name: "Rampage (Attack Nearby Targets on Hate Lists)",
	},
	206: &model.SpellEffect{
		ID:   206,
		Type: 0,
		Name: "Area Attempt to Taunt",
	},
	207: &model.SpellEffect{
		ID:   207,
		Type: 0,
		Name: "Extract Bone Chips from Meat",
	},
	208: &model.SpellEffect{
		ID:   208,
		Type: 0,
		Name: "Purge Poison (Not Used)",
	},
	209: &model.SpellEffect{
		ID:   209,
		Type: 0,
		Name: "Attempt to Remove Beneficial Effect",
	},
	210: &model.SpellEffect{
		ID:   210,
		Type: 0,
		Name: "Pet Shield (Not Used)",
	},
	211: &model.SpellEffect{
		ID:   211,
		Type: 0,
		Name: "Area Melee Attack",
	},
	212: &model.SpellEffect{
		ID:   212,
		Type: 1,
		Name: "Spell Critical Chance and Mana Cost",
	},
	213: &model.SpellEffect{
		ID:   213,
		Type: 1,
		Name: "Pet's Maximum HP",
	},
	214: &model.SpellEffect{
		ID:   214,
		Type: 1,
		Name: "Maximum HP by Percentage",
	},
	215: &model.SpellEffect{
		ID:   215,
		Type: 1,
		Name: "Pet Melee Avoidance",
	},
	216: &model.SpellEffect{
		ID:   216,
		Type: 1,
		Name: "Melee Accuracy",
	},
	217: &model.SpellEffect{
		ID:   217,
		Type: 0,
		Name: "Chance to Headshot for 32K Damage",
	},
	218: &model.SpellEffect{
		ID:   218,
		Type: 1,
		Name: "Pet's Chance to Critical Hit",
	},
	219: &model.SpellEffect{
		ID:   219,
		Type: 0,
		Name: "Critical Spell Hit vs Undead",
	},
	220: &model.SpellEffect{
		ID:   220,
		Type: 1,
		Name: "Skill Damage Modifier",
	},
	221: &model.SpellEffect{
		ID:   221,
		Type: 1,
		Name: "Weight Encumbrance",
	},
	222: &model.SpellEffect{
		ID:   222,
		Type: 1,
		Name: "Chance to Block Attacks from Behind",
	},
	223: &model.SpellEffect{
		ID:   223,
		Type: 1,
		Name: "Chance to Double Riposte",
	},
	224: &model.SpellEffect{
		ID:   224,
		Type: 0,
		Name: "Perform Double Riposte",
	},
	225: &model.SpellEffect{
		ID:   225,
		Type: 0,
		Name: "Perform Double Attack",
	},
	226: &model.SpellEffect{
		ID:   226,
		Type: 0,
		Name: "Perform Two Handed Bash",
	},
	227: &model.SpellEffect{
		ID:   227,
		Type: 1,
		Name: "Skill Timer",
	},
	228: &model.SpellEffect{
		ID:   228,
		Type: 1,
		Name: "Falling Damage (Not Used)",
	},
	229: &model.SpellEffect{
		ID:   229,
		Type: 1,
		Name: "Chance to Cast Through Interruptions",
	},
	230: &model.SpellEffect{
		ID:   230,
		Type: 0,
		Name: "Extended Shielding (Not Used)",
	},
	231: &model.SpellEffect{
		ID:   231,
		Type: 1,
		Name: "Chance for Bashes to Stun",
	},
	232: &model.SpellEffect{
		ID:   232,
		Type: 1,
		Name: "Chance to Save from Death",
	},
	233: &model.SpellEffect{
		ID:   233,
		Type: 1,
		Name: "Metabolism",
	},
	234: &model.SpellEffect{
		ID:   234,
		Type: 1,
		Name: "Apply Poison Time (Not Used)",
	},
	235: &model.SpellEffect{
		ID:   235,
		Type: 1,
		Name: "Chance to Channel Spells through Interruptions",
	},
	236: &model.SpellEffect{
		ID:   236,
		Type: 0,
		Name: "Free Pet (Not Used)",
	},
	237: &model.SpellEffect{
		ID:   237,
		Type: 0,
		Name: "Give Pets Group Buffs",
	},
	238: &model.SpellEffect{
		ID:   238,
		Type: 0,
		Name: "Make Illusions Persist Across Zones",
	},
	239: &model.SpellEffect{
		ID:   239,
		Type: 0,
		Name: "Feigned Cast On Chance (Not Used)",
	},
	240: &model.SpellEffect{
		ID:   240,
		Type: 0,
		Name: "String Unbreakable (Not Used)",
	},
	241: &model.SpellEffect{
		ID:   241,
		Type: 1,
		Name: "Pet Energy Reclaiming Efficiency",
	},
	242: &model.SpellEffect{
		ID:   242,
		Type: 1,
		Name: "Chance to Wipe Hate Lists",
	},
	243: &model.SpellEffect{
		ID:   243,
		Type: 1,
		Name: "Chance to Prevent Charm from Breaking",
	},
	244: &model.SpellEffect{
		ID:   244,
		Type: 1,
		Name: "Chance to Prevent Root from Breaking",
	},
	245: &model.SpellEffect{
		ID:   245,
		Type: 0,
		Name: "Trap Circumvention (Not Used)",
	},
	246: &model.SpellEffect{
		ID:   246,
		Type: 0,
		Name: "Set Breathing Air Supply Level",
	},
	247: &model.SpellEffect{
		ID:   247,
		Type: 1,
		Name: "Skill Cap:",
	},
	248: &model.SpellEffect{
		ID:   248,
		Type: 0,
		Name: "Secondary Forte (Not Used)",
	},
	249: &model.SpellEffect{
		ID:   249,
		Type: 1,
		Name: "Damage Modifier (V3?)",
	},
	250: &model.SpellEffect{
		ID:   250,
		Type: 1,
		Name: "Chance of Spell Attack Procs",
	},
	251: &model.SpellEffect{
		ID:   251,
		Type: 0,
		Name: "Consume Projectile:",
	},
	252: &model.SpellEffect{
		ID:   252,
		Type: 1,
		Name: "Chance to Backstab from the Front",
	},
	253: &model.SpellEffect{
		ID:   253,
		Type: 1,
		Name: "Front Backstab Minimum Damage",
	},
	254: &model.SpellEffect{
		ID:   254,
		Type: 0,
		Name: "-",
	},
	255: &model.SpellEffect{
		ID:   255,
		Type: 1,
		Name: "Shield Duration (Not Used)",
	},
	256: &model.SpellEffect{
		ID:   256,
		Type: 0,
		Name: "Shroud of Stealth (Not Used)",
	},
	257: &model.SpellEffect{
		ID:   257,
		Type: 0,
		Name: "Pet Discipline (Not Used)",
	},
	258: &model.SpellEffect{
		ID:   258,
		Type: 1,
		Name: "Chance for Triple Backstab Damage",
	},
	259: &model.SpellEffect{
		ID:   259,
		Type: 1,
		Name: "Direct Damage Mitigation",
	},
	260: &model.SpellEffect{
		ID:   260,
		Type: 1,
		Name: "Singing Effectiveness",
	},
	261: &model.SpellEffect{
		ID:   261,
		Type: 1,
		Name: "Cap on Singing Effectiveness",
	},
	262: &model.SpellEffect{
		ID:   262,
		Type: 1,
		Name: "Stat Caps",
	},
	263: &model.SpellEffect{
		ID:   263,
		Type: 1,
		Name: "Tradeskill Mastery (Not Used)",
	},
	264: &model.SpellEffect{
		ID:   264,
		Type: 0,
		Name: "Hastened AA Skill (Not Used)",
	},
	265: &model.SpellEffect{
		ID:   265,
		Type: 0,
		Name: "Immunity to Fizzling Spells",
	},
	266: &model.SpellEffect{
		ID:   266,
		Type: 1,
		Name: "Chance of Two-Handed Triple Attack",
	},
	267: &model.SpellEffect{
		ID:   267,
		Type: 0,
		Name: "Pet Discipline 2 (Not Used)",
	},
	268: &model.SpellEffect{
		ID:   268,
		Type: 0,
		Name: "Reduce Tradeskill Fail (Not Used)",
	},
	269: &model.SpellEffect{
		ID:   269,
		Type: 1,
		Name: "Maximum Bind Woundable Health",
	},
	270: &model.SpellEffect{
		ID:   270,
		Type: 1,
		Name: "Range of Bard Songs",
	},
	271: &model.SpellEffect{
		ID:   271,
		Type: 1,
		Name: "Minimum Movement Speed",
	},
	272: &model.SpellEffect{
		ID:   272,
		Type: 1,
		Name: "Effective Spellcasting Level V2",
	},
	273: &model.SpellEffect{
		ID:   273,
		Type: 1,
		Name: "Critical DoT Chance",
	},
	274: &model.SpellEffect{
		ID:   274,
		Type: 1,
		Name: "Chance to Critical Heal",
	},
	275: &model.SpellEffect{
		ID:   275,
		Type: 1,
		Name: "Chance to Critical Mend",
	},
	276: &model.SpellEffect{
		ID:   276,
		Type: 1,
		Name: "Chance of Offhand Attacks",
	},
	277: &model.SpellEffect{
		ID:   277,
		Type: 1,
		Name: "Chance of Saving from Death",
	},
	278: &model.SpellEffect{
		ID:   278,
		Type: 1,
		Name: "Chance to Inflict Finishing Blow on Fleeing Enemy",
	},
	279: &model.SpellEffect{
		ID:   279,
		Type: 1,
		Name: "Chance to Flurry Attacks",
	},
	280: &model.SpellEffect{
		ID:   280,
		Type: 1,
		Name: "Chance for Pet to Flurry Attacks",
	},
	281: &model.SpellEffect{
		ID:   281,
		Type: 0,
		Name: "Pet Feign Death (Not Used)",
	},
	282: &model.SpellEffect{
		ID:   282,
		Type: 1,
		Name: "Bind Wound Effectiveness",
	},
	283: &model.SpellEffect{
		ID:   283,
		Type: 1,
		Name: "Chance to Perform Double Special Attack",
	},
	284: &model.SpellEffect{
		ID:   284,
		Type: 0,
		Name: "Set Lay-on-Hands Heal (Not Used)",
	},
	285: &model.SpellEffect{
		ID:   285,
		Type: 0,
		Name: "Nimble Evasion (Not Used)",
	},
	286: &model.SpellEffect{
		ID:   286,
		Type: 1,
		Name: "Spell Damage (V3?)",
	},
	287: &model.SpellEffect{
		ID:   287,
		Type: 1,
		Name: "Song/Spell Duration by 1 Tick (6s)",
	},
	288: &model.SpellEffect{
		ID:   288,
		Type: 1,
		Name: "Chance to Knockback with Special Attacks",
	},
	289: &model.SpellEffect{
		ID:   289,
		Type: 0,
		Name: "Cast New Spell when Wearing Off:",
	},
	290: &model.SpellEffect{
		ID:   290,
		Type: 1,
		Name: "Maximum Movement Speed",
	},
	291: &model.SpellEffect{
		ID:   291,
		Type: 0,
		Name: "Attempt to Remove Detrimental Effects",
	},
	292: &model.SpellEffect{
		ID:   292,
		Type: 1,
		Name: "Chance to Strike Through V2",
	},
	293: &model.SpellEffect{
		ID:   293,
		Type: 1,
		Name: "Resistance to Stuns from the Front",
	},
	294: &model.SpellEffect{
		ID:   294,
		Type: 1,
		Name: "Critical Spell Chance",
	},
	295: &model.SpellEffect{
		ID:   295,
		Type: 0,
		Name: "Reduce Timer Special (Not Used)",
	},
	296: &model.SpellEffect{
		ID:   296,
		Type: 1,
		Name: "Spell Vulnerability Focus",
	},
	297: &model.SpellEffect{
		ID:   297,
		Type: 1,
		Name: "Incoming Damage",
	},
	298: &model.SpellEffect{
		ID:   298,
		Type: 1,
		Name: "Physical Size",
	},
	299: &model.SpellEffect{
		ID:   299,
		Type: 0,
		Name: "Awaken Corpse as Temporary Pet",
	},
	300: &model.SpellEffect{
		ID:   300,
		Type: 0,
		Name: "Summon Doppelganger",
	},
	301: &model.SpellEffect{
		ID:   301,
		Type: 1,
		Name: "Archery Damage",
	},
	302: &model.SpellEffect{
		ID:   302,
		Type: 1,
		Name: "Critical Hit Damage %",
	},
	303: &model.SpellEffect{
		ID:   303,
		Type: 1,
		Name: "Critical Hit Damage",
	},
	304: &model.SpellEffect{
		ID:   304,
		Type: 1,
		Name: "Chance to Avoid Offhand Attack Ripostes",
	},
	305: &model.SpellEffect{
		ID:   305,
		Type: 1,
		Name: "Damage Shield Mitigation",
	},
	306: &model.SpellEffect{
		ID:   306,
		Type: 0,
		Name: "Army of the Dead (Not Used)",
	},
	307: &model.SpellEffect{
		ID:   307,
		Type: 0,
		Name: "Appraisal (Not Used)",
	},
	308: &model.SpellEffect{
		ID:   308,
		Type: 0,
		Name: "Suspend Pet",
	},
	309: &model.SpellEffect{
		ID:   309,
		Type: 0,
		Name: "Teleport to Caster's Bind Point",
	},
	310: &model.SpellEffect{
		ID:   310,
		Type: 1,
		Name: "Reuse Timer Reduction",
	},
	311: &model.SpellEffect{
		ID:   311,
		Type: 0,
		Name: "Limit - Not Applicable to Innate Weapon Procs",
	},
	312: &model.SpellEffect{
		ID:   312,
		Type: 0,
		Name: "Temporarily Drop to Bottom of Hate List",
	},
	313: &model.SpellEffect{
		ID:   313,
		Type: 1,
		Name: "Chance to Forage Additional Items",
	},
	314: &model.SpellEffect{
		ID:   314,
		Type: 0,
		Name: "Invisibility (Fixed Duration)",
	},
	315: &model.SpellEffect{
		ID:   315,
		Type: 0,
		Name: "Invisibility to Undead (Fixed Duration)",
	},
	316: &model.SpellEffect{
		ID:   316,
		Type: 0,
		Name: "Invisibility to Animals (Fixed Duration) (Not Used)",
	},
	317: &model.SpellEffect{
		ID:   317,
		Type: 1,
		Name: "Worn Item HP Regeneration Cap",
	},
	318: &model.SpellEffect{
		ID:   318,
		Type: 1,
		Name: "Worn Item Mana Regeneration Cap",
	},
	319: &model.SpellEffect{
		ID:   319,
		Type: 1,
		Name: "Chance for Heal-over-Time to Critical Tick",
	},
	320: &model.SpellEffect{
		ID:   320,
		Type: 1,
		Name: "Shield Block Chance",
	},
	321: &model.SpellEffect{
		ID:   321,
		Type: 1,
		Name: "Hate with Target",
	},
	322: &model.SpellEffect{
		ID:   322,
		Type: 0,
		Name: "Teleport to Home City",
	},
	323: &model.SpellEffect{
		ID:   323,
		Type: 0,
		Name: "Add Defensive Proc on Hit:",
	},
	324: &model.SpellEffect{
		ID:   324,
		Type: 0,
		Name: "Convert HP to Mana",
	},
	325: &model.SpellEffect{
		ID:   325,
		Type: 0,
		Name: "Chance Invs Break to AoE (Not Used)",
	},
	326: &model.SpellEffect{
		ID:   326,
		Type: 1,
		Name: "Number of Spell Gem Slots",
	},
	327: &model.SpellEffect{
		ID:   327,
		Type: 1,
		Name: "Maximum Number of Magical Effects",
	},
	328: &model.SpellEffect{
		ID:   328,
		Type: 1,
		Name: "Damage Taken when Unconscious Before Death",
	},
	329: &model.SpellEffect{
		ID:   329,
		Type: 1,
		Name: "Damage Absorption with Mana",
	},
	330: &model.SpellEffect{
		ID:   330,
		Type: 1,
		Name: "Critical Hit Damage with All Skills",
	},
	331: &model.SpellEffect{
		ID:   331,
		Type: 1,
		Name: "Chance to Salvage Tradeskill Failures",
	},
	332: &model.SpellEffect{
		ID:   332,
		Type: 0,
		Name: "Summon to Corpse (Not Used)",
	},
	333: &model.SpellEffect{
		ID:   333,
		Type: 0,
		Name: "Cast New Spell When Rune Fades:",
	},
	334: &model.SpellEffect{
		ID:   334,
		Type: 1,
		Name: "Current HP V2",
	},
	335: &model.SpellEffect{
		ID:   335,
		Type: 0,
		Name: "Block Next Spell Focus",
	},
	336: &model.SpellEffect{
		ID:   336,
		Type: 0,
		Name: "Illusionary Target (Not Used)",
	},
	337: &model.SpellEffect{
		ID:   337,
		Type: 1,
		Name: "Experience Gain Modifier",
	},
	338: &model.SpellEffect{
		ID:   338,
		Type: 0,
		Name: "Summon And Resurrect All Corpses",
	},
	339: &model.SpellEffect{
		ID:   339,
		Type: 0,
		Name: "Trigger Second Spell:",
	},
	340: &model.SpellEffect{
		ID:   340,
		Type: 0,
		Name: "Add Chance to Trigger Spell:",
	},
	341: &model.SpellEffect{
		ID:   341,
		Type: 1,
		Name: "Worn Item Attack Rating Cap",
	},
	342: &model.SpellEffect{
		ID:   342,
		Type: 0,
		Name: "Prevent from Fleeing",
	},
	343: &model.SpellEffect{
		ID:   343,
		Type: 0,
		Name: "Interrupt Casting",
	},
	344: &model.SpellEffect{
		ID:   344,
		Type: 1,
		Name: "Chance of Channeling Item Effect through Interruption",
	},
	345: &model.SpellEffect{
		ID:   345,
		Type: 1,
		Name: "Maximum Permitted Level for Assassinate",
	},
	346: &model.SpellEffect{
		ID:   346,
		Type: 1,
		Name: "Maximum Permitted Level for HeadShot",
	},
	347: &model.SpellEffect{
		ID:   347,
		Type: 1,
		Name: "Chance to Perform Double Ranged Attack",
	},
	348: &model.SpellEffect{
		ID:   348,
		Type: 0,
		Name: "Limit - Minimum Mana Cost:",
	},
	349: &model.SpellEffect{
		ID:   349,
		Type: 1,
		Name: "Hate Modifier with Shield Equipped",
	},
	350: &model.SpellEffect{
		ID:   350,
		Type: 1,
		Name: "Current HP by Draining Mana",
	},
	351: &model.SpellEffect{
		ID:   351,
		Type: 0,
		Name: "Persistent Effect (Not Used)",
	},
	352: &model.SpellEffect{
		ID:   352,
		Type: 0,
		Name: "Increase Trap Count (Not Used)",
	},
	353: &model.SpellEffect{
		ID:   353,
		Type: 0,
		Name: "Additional Aura (Not Used)",
	},
	354: &model.SpellEffect{
		ID:   354,
		Type: 0,
		Name: "Deactivate All Traps (Not Used)",
	},
	355: &model.SpellEffect{
		ID:   355,
		Type: 0,
		Name: "Learn Trap (Not Used)",
	},
	356: &model.SpellEffect{
		ID:   356,
		Type: 0,
		Name: "Change Trigger Type (Not Used)",
	},
	357: &model.SpellEffect{
		ID:   357,
		Type: 0,
		Name: "Silence Casting of Spells Affecting:",
	},
	358: &model.SpellEffect{
		ID:   358,
		Type: 1,
		Name: "Current Mana",
	},
	359: &model.SpellEffect{
		ID:   359,
		Type: 0,
		Name: "Passive Sense Trap (Not Used)",
	},
	360: &model.SpellEffect{
		ID:   360,
		Type: 0,
		Name: "Add Chance to Proc On Kill Shot:",
	},
	361: &model.SpellEffect{
		ID:   361,
		Type: 0,
		Name: "Add Chance to Cast Spell on Death:",
	},
	362: &model.SpellEffect{
		ID:   362,
		Type: 1,
		Name: "Potion Belt Slots (Not Used)",
	},
	363: &model.SpellEffect{
		ID:   363,
		Type: 1,
		Name: "Bandolier Slots (Not Used)",
	},
	364: &model.SpellEffect{
		ID:   364,
		Type: 1,
		Name: "Chance to Triple Attack",
	},
	365: &model.SpellEffect{
		ID:   365,
		Type: 0,
		Name: "Add Chance to Proc on Spellcast Kill Shot:",
	},
	366: &model.SpellEffect{
		ID:   366,
		Type: 1,
		Name: "Damage Modifier with Shield Equipped",
	},
	367: &model.SpellEffect{
		ID:   367,
		Type: 0,
		Name: "Set Body Type:",
	},
	368: &model.SpellEffect{
		ID:   368,
		Type: 0,
		Name: "Faction Mod (Not Used)",
	},
	369: &model.SpellEffect{
		ID:   369,
		Type: 1,
		Name: "Corruption Counter",
	},
	370: &model.SpellEffect{
		ID:   370,
		Type: 1,
		Name: "Resistance to Corruption",
	},
	371: &model.SpellEffect{
		ID:   371,
		Type: 1,
		Name: "Attack Speed V4",
	},
	372: &model.SpellEffect{
		ID:   372,
		Type: 1,
		Name: "Forage Skill (Not Used)",
	},
	373: &model.SpellEffect{
		ID:   373,
		Type: 0,
		Name: "Cast New Spell on Fade:",
	},
	374: &model.SpellEffect{
		ID:   374,
		Type: 0,
		Name: "Apply Effect:",
	},
	375: &model.SpellEffect{
		ID:   375,
		Type: 1,
		Name: "Critical DoT Tick Damage",
	},
	376: &model.SpellEffect{
		ID:   376,
		Type: 0,
		Name: "Fling (Not Used)",
	},
	377: &model.SpellEffect{
		ID:   377,
		Type: 0,
		Name: "NPC - Cast New Spell on Fade:",
	},
	378: &model.SpellEffect{
		ID:   378,
		Type: 0,
		Name: "Add Chance to Resist Spell Effect:",
	},
	379: &model.SpellEffect{
		ID:   379,
		Type: 0,
		Name: "Directional Short Range Teleport",
	},
	380: &model.SpellEffect{
		ID:   380,
		Type: 0,
		Name: "Knockdown",
	},
	381: &model.SpellEffect{
		ID:   381,
		Type: 0,
		Name: "Knock Toward Caster (Not Used)",
	},
	382: &model.SpellEffect{
		ID:   382,
		Type: 0,
		Name: "Negate Spell Effect:",
	},
	383: &model.SpellEffect{
		ID:   383,
		Type: 0,
		Name: "Add Chance of Spellcasting to Proc:",
	},
	384: &model.SpellEffect{
		ID:   384,
		Type: 0,
		Name: "Leap",
	},
	385: &model.SpellEffect{
		ID:   385,
		Type: 0,
		Name: "Limit - Group Spells Only",
	},
	386: &model.SpellEffect{
		ID:   386,
		Type: 0,
		Name: "Cast New Spell on Curer when Cured:",
	},
	387: &model.SpellEffect{
		ID:   387,
		Type: 0,
		Name: "Cast New Spell when Cured:",
	},
	388: &model.SpellEffect{
		ID:   388,
		Type: 0,
		Name: "Summon Corpse Zone (Not Used)",
	},
	389: &model.SpellEffect{
		ID:   389,
		Type: 0,
		Name: "Reactivate All Spell Gems in Cooldown",
	},
	390: &model.SpellEffect{
		ID:   390,
		Type: 0,
		Name: "Timer Lockout (Not Used)",
	},
	391: &model.SpellEffect{
		ID:   391,
		Type: 1,
		Name: "Melee Vulnerability",
	},
	392: &model.SpellEffect{
		ID:   392,
		Type: 1,
		Name: "Healing Amount V3",
	},
	393: &model.SpellEffect{
		ID:   393,
		Type: 1,
		Name: "Incoming Healing",
	},
	394: &model.SpellEffect{
		ID:   394,
		Type: 1,
		Name: "Incoming Healing",
	},
	395: &model.SpellEffect{
		ID:   395,
		Type: 1,
		Name: "Chance of Critical Incoming Heals",
	},
	396: &model.SpellEffect{
		ID:   396,
		Type: 1,
		Name: "Critical Healing Amount",
	},
	397: &model.SpellEffect{
		ID:   397,
		Type: 1,
		Name: "Pet's Melee Mitigation",
	},
	398: &model.SpellEffect{
		ID:   398,
		Type: 1,
		Name: "Swarm Pet Duration",
	},
	399: &model.SpellEffect{
		ID:   399,
		Type: 0,
		Name: "Add Chance to Double-Cast Spells",
	},
	400: &model.SpellEffect{
		ID:   400,
		Type: 1,
		Name: "Current HP of Group from Own Mana Pool",
	},
	401: &model.SpellEffect{
		ID:   401,
		Type: 1,
		Name: "Current HP from Own Mana Pool",
	},
	402: &model.SpellEffect{
		ID:   402,
		Type: 1,
		Name: "Current HP from Own Endurance Pool",
	},
	403: &model.SpellEffect{
		ID:   403,
		Type: 0,
		Name: "Limit - Spell Class:",
	},
	404: &model.SpellEffect{
		ID:   404,
		Type: 0,
		Name: "Limit - Spell Subclass:",
	},
	405: &model.SpellEffect{
		ID:   405,
		Type: 0,
		Name: "Add Chance to Block with Two-Handed Blunt Weapons (Staves)",
	},
	406: &model.SpellEffect{
		ID:   406,
		Type: 0,
		Name: "Cast New Spell when Fade from NumHits Depleted:",
	},
	407: &model.SpellEffect{
		ID:   407,
		Type: 0,
		Name: "Cast New Spell when Focus Effect Is Applied:",
	},
	408: &model.SpellEffect{
		ID:   408,
		Type: 0,
		Name: "Limit - Maximum Percent of HP",
	},
	409: &model.SpellEffect{
		ID:   409,
		Type: 0,
		Name: "Limit - Maximum Percent of Mana",
	},
	410: &model.SpellEffect{
		ID:   410,
		Type: 0,
		Name: "Limit - Maximum Percent of Endurance",
	},
	411: &model.SpellEffect{
		ID:   411,
		Type: 0,
		Name: "Limit to Class:",
	},
	412: &model.SpellEffect{
		ID:   412,
		Type: 0,
		Name: "Limit to Race:",
	},
	413: &model.SpellEffect{
		ID:   413,
		Type: 1,
		Name: "Base Power of Skills and Songs",
	},
	414: &model.SpellEffect{
		ID:   414,
		Type: 0,
		Name: "Limit - Casting Skill:",
	},
	415: &model.SpellEffect{
		ID:   415,
		Type: 1,
		Name: "FF Item Class (Not Used)",
	},
	416: &model.SpellEffect{
		ID:   416,
		Type: 1,
		Name: "Armor Class V2",
	},
	417: &model.SpellEffect{
		ID:   417,
		Type: 1,
		Name: "Mana Regeneration V2",
	},
	418: &model.SpellEffect{
		ID:   418,
		Type: 1,
		Name: "Skill Damage V2",
	},
	419: &model.SpellEffect{
		ID:   419,
		Type: 0,
		Name: "Add Chance of Melee Proc:",
	},
	420: &model.SpellEffect{
		ID:   420,
		Type: 1,
		Name: "NumHits for All Spells",
	},
	421: &model.SpellEffect{
		ID:   421,
		Type: 1,
		Name: "NumHits for Spell:",
	},
	422: &model.SpellEffect{
		ID:   422,
		Type: 0,
		Name: "Limit - Minimum NumHits:",
	},
	423: &model.SpellEffect{
		ID:   423,
		Type: 0,
		Name: "Limit - NumHits Type:",
	},
	424: &model.SpellEffect{
		ID:   424,
		Type: 1,
		Name: "Gravitation to/from Mob",
	},
	425: &model.SpellEffect{
		ID:   425,
		Type: 0,
		Name: "Display (Not Used)",
	},
	426: &model.SpellEffect{
		ID:   426,
		Type: 1,
		Name: "Extended Targets (Not Used)",
	},
	427: &model.SpellEffect{
		ID:   427,
		Type: 0,
		Name: "Add Chance to Proc with Skill:",
	},
	428: &model.SpellEffect{
		ID:   428,
		Type: 0,
		Name: "Limit - Skill:",
	},
	429: &model.SpellEffect{
		ID:   429,
		Type: 0,
		Name: "Add Chance to Proc with Skull Success:",
	},
	430: &model.SpellEffect{
		ID:   430,
		Type: 0,
		Name: "Post Effect (Not Used)",
	},
	431: &model.SpellEffect{
		ID:   431,
		Type: 0,
		Name: "Post Effect Data (Not Used)",
	},
	432: &model.SpellEffect{
		ID:   432,
		Type: 0,
		Name: "Expand Max Active Trophy Ben (Not Used)",
	},
	433: &model.SpellEffect{
		ID:   433,
		Type: 1,
		Name: "Critical DoT Chance, Decay",
	},
	434: &model.SpellEffect{
		ID:   434,
		Type: 1,
		Name: "Critical Heal Chance, Decay",
	},
	435: &model.SpellEffect{
		ID:   435,
		Type: 1,
		Name: "Critical Heal over Time Chance, Decay",
	},
	436: &model.SpellEffect{
		ID:   436,
		Type: 0,
		Name: "Beneficial Countdown Hold (Not Used)",
	},
	437: &model.SpellEffect{
		ID:   437,
		Type: 0,
		Name: "Teleport to Anchor (Not Yet Used)",
	},
	438: &model.SpellEffect{
		ID:   438,
		Type: 0,
		Name: "Translocate to Anchor (Not Yet Used)",
	},
	439: &model.SpellEffect{
		ID:   439,
		Type: 0,
		Name: "Add Chance to Assassinate Backstab for 32K Damage",
	},
	440: &model.SpellEffect{
		ID:   440,
		Type: 1,
		Name: "Maximum Permitted Level for Finishing Blow",
	},
	441: &model.SpellEffect{
		ID:   441,
		Type: 0,
		Name: "Limit - While within Distance from Buff Location",
	},
	442: &model.SpellEffect{
		ID:   442,
		Type: 0,
		Name: "Trigger New Spell when Target",
	},
	443: &model.SpellEffect{
		ID:   443,
		Type: 0,
		Name: "Trigger New Spell when Caster",
	},
	444: &model.SpellEffect{
		ID:   444,
		Type: 0,
		Name: "Lock Aggro on Caster, Lowering Others' Hate by Percentage",
	},
	445: &model.SpellEffect{
		ID:   445,
		Type: 1,
		Name: "Mercenary Slots (Not Used)",
	},
	446: &model.SpellEffect{
		ID:   446,
		Type: 0,
		Name: "Limit - Stacking Blocker A",
	},
	447: &model.SpellEffect{
		ID:   447,
		Type: 0,
		Name: "Limit - Stacking Blocker B",
	},
	448: &model.SpellEffect{
		ID:   448,
		Type: 0,
		Name: "Limit - Stacking Blocker C",
	},
	449: &model.SpellEffect{
		ID:   449,
		Type: 0,
		Name: "Limit - Stacking Blocker D",
	},
	450: &model.SpellEffect{
		ID:   450,
		Type: 1,
		Name: "DoT Damage Mitigation",
	},
	451: &model.SpellEffect{
		ID:   451,
		Type: 1,
		Name: "Melee Damage Mitigation for Hits Over",
	},
	452: &model.SpellEffect{
		ID:   452,
		Type: 1,
		Name: "Spell Damage Mitigation for Hits Over",
	},
	453: &model.SpellEffect{
		ID:   453,
		Type: 0,
		Name: "Trigger on a Melee Hit Over",
	},
	454: &model.SpellEffect{
		ID:   454,
		Type: 0,
		Name: "Trigger on a Spell Hit Over",
	},
	455: &model.SpellEffect{
		ID:   455,
		Type: 1,
		Name: "Total Hate Amount",
	},
	456: &model.SpellEffect{
		ID:   456,
		Type: 1,
		Name: "Total Hate Amount",
	},
	457: &model.SpellEffect{
		ID:   457,
		Type: 1,
		Name: "HP/Mana/Endurance from Spell Damage",
	},
	458: &model.SpellEffect{
		ID:   458,
		Type: 1,
		Name: "Faction Gain/Loss Modifier",
	},
	459: &model.SpellEffect{
		ID:   459,
		Type: 1,
		Name: "Damage Modifier V2 for Skill",
	},
}
