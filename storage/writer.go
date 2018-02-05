package storage

import (
	"github.com/xackery/xegony/model"
)

//Writer is a generic interface of all storage types
type Writer interface {

	//Account
	CreateAccount(account *model.Account) (err error)
	EditAccount(account *model.Account) (err error)
	DeleteAccount(account *model.Account) (err error)

	//Character
	CreateCharacter(character *model.Character) (err error)
	EditCharacter(character *model.Character) (err error)
	DeleteCharacter(character *model.Character) (err error)

	//Class
	CreateClass(class *model.Class) (err error)
	EditClass(class *model.Class) (err error)
	DeleteClass(class *model.Class) (err error)

	//Config
	CreateConfig(config *model.Config) (err error)
	EditConfig(config *model.Config) (err error)
	DeleteConfig(config *model.Config) (err error)

	//Item
	CreateItem(item *model.Item) (err error)
	EditItem(item *model.Item) (err error)
	DeleteItem(item *model.Item) (err error)

	//Race
	CreateRace(race *model.Race) (err error)
	EditRace(race *model.Race) (err error)
	DeleteRace(race *model.Race) (err error)

	//Rule
	CreateRule(rule *model.Rule) (err error)
	EditRule(rule *model.Rule) (err error)
	DeleteRule(rule *model.Rule) (err error)

	//RuleEntry
	CreateRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error)
	EditRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error)
	DeleteRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error)

	//Spell
	CreateSpell(spell *model.Spell) (err error)
	EditSpell(spell *model.Spell) (err error)
	DeleteSpell(spell *model.Spell) (err error)

	//SpellAnimation
	CreateSpellAnimation(spellAnimation *model.SpellAnimation) (err error)
	EditSpellAnimation(spellAnimation *model.SpellAnimation) (err error)
	DeleteSpellAnimation(spellAnimation *model.SpellAnimation) (err error)

	//SpellAnimationType
	CreateSpellAnimationType(spellAnimationType *model.SpellAnimationType) (err error)
	EditSpellAnimationType(spellAnimationType *model.SpellAnimationType) (err error)
	DeleteSpellAnimationType(spellAnimationType *model.SpellAnimationType) (err error)

	//SpellEffectFormula
	CreateSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula) (err error)
	EditSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula) (err error)
	DeleteSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula) (err error)

	//SpellEffectType
	CreateSpellEffectType(spellEffectType *model.SpellEffectType) (err error)
	EditSpellEffectType(spellEffectType *model.SpellEffectType) (err error)
	DeleteSpellEffectType(spellEffectType *model.SpellEffectType) (err error)

	//User
	CreateUser(user *model.User) (err error)
	EditUser(user *model.User) (err error)
	DeleteUser(user *model.User) (err error)

	//Variable
	CreateVariable(variable *model.Variable) (err error)
	EditVariable(variable *model.Variable) (err error)
	DeleteVariable(variable *model.Variable) (err error)

	//Zone
	CreateZone(zone *model.Zone) (err error)
	EditZone(zone *model.Zone) (err error)
	DeleteZone(zone *model.Zone) (err error)

	//ZoneExpansion
	CreateZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error)
	EditZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error)
	DeleteZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error)
}
