package storage

import (
	"github.com/xackery/xegony/model"
)

//Reader is a generic interface of all storage types
type Reader interface {
	//Aa
	/*GetAa(aa *model.Aa) (err error)
	ListAa() (aas []*model.Aa, err error)

	//AaRank
	GetAaRank(aaRank *model.AaRank) (err error)
	ListAaRank() (aaRanks []*model.AaRank, err error)
	*/
	//Account
	GetAccount(account *model.Account) (err error)
	ListAccount(page *model.Page) (accounts []*model.Account, err error)
	ListAccountTotalCount() (count int64, err error)
	ListAccountBySearch(page *model.Page, account *model.Account) (accounts []*model.Account, err error)
	ListAccountBySearchTotalCount(account *model.Account) (count int64, err error)

	//Character
	GetCharacter(character *model.Character) (err error)
	ListCharacter(page *model.Page) (characters []*model.Character, err error)
	ListCharacterTotalCount() (count int64, err error)
	ListCharacterBySearch(page *model.Page, character *model.Character) (characters []*model.Character, err error)
	ListCharacterBySearchTotalCount(character *model.Character) (count int64, err error)

	//Class
	GetClass(class *model.Class) (err error)
	ListClass(page *model.Page) (classs []*model.Class, err error)
	ListClassTotalCount() (count int64, err error)
	ListClassByBit(page *model.Page, class *model.Class) (races []*model.Class, err error)
	ListClassByBitTotalCount(class *model.Class) (count int64, err error)
	ListClassBySearch(page *model.Page, class *model.Class) (classs []*model.Class, err error)
	ListClassBySearchTotalCount(class *model.Class) (count int64, err error)

	//Config
	GetConfig(config *model.Config) (err error)
	ListConfig(page *model.Page) (configs []*model.Config, err error)
	ListConfigTotalCount() (count int64, err error)
	ListConfigBySearch(page *model.Page, config *model.Config) (configs []*model.Config, err error)
	ListConfigBySearchTotalCount(config *model.Config) (count int64, err error)

	//Deity
	GetDeity(deity *model.Deity) (err error)
	GetDeityBySpell(spell *model.Spell, deity *model.Deity) (err error)
	ListDeity(page *model.Page) (deitys []*model.Deity, err error)
	ListDeityTotalCount() (count int64, err error)
	ListDeityByBit(page *model.Page, deity *model.Deity) (races []*model.Deity, err error)
	ListDeityByBitTotalCount(deity *model.Deity) (count int64, err error)
	ListDeityBySearch(page *model.Page, deity *model.Deity) (deitys []*model.Deity, err error)
	ListDeityBySearchTotalCount(deity *model.Deity) (count int64, err error)

	//Item
	GetItem(item *model.Item) (err error)
	ListItem(page *model.Page) (items []*model.Item, err error)
	ListItemTotalCount() (count int64, err error)
	ListItemBySearch(page *model.Page, item *model.Item) (items []*model.Item, err error)
	ListItemBySearchTotalCount(item *model.Item) (count int64, err error)

	//Npc
	GetNpc(npc *model.Npc) (err error)
	ListNpc(page *model.Page) (npcs []*model.Npc, err error)
	ListNpcTotalCount() (count int64, err error)
	ListNpcBySearch(page *model.Page, npc *model.Npc) (npcs []*model.Npc, err error)
	ListNpcBySearchTotalCount(npc *model.Npc) (count int64, err error)

	//OauthType
	GetOauthType(oauthType *model.OauthType) (err error)
	ListOauthType(page *model.Page) (oauthTypes []*model.OauthType, err error)
	ListOauthTypeTotalCount() (count int64, err error)
	ListOauthTypeBySearch(page *model.Page, oauthType *model.OauthType) (oauthTypes []*model.OauthType, err error)
	ListOauthTypeBySearchTotalCount(oauthType *model.OauthType) (count int64, err error)

	//Race
	GetRace(race *model.Race) (err error)
	ListRace(page *model.Page) (races []*model.Race, err error)
	ListRaceTotalCount() (count int64, err error)
	ListRaceByBit(page *model.Page, race *model.Race) (races []*model.Race, err error)
	ListRaceByBitTotalCount(race *model.Race) (count int64, err error)
	ListRaceBySearch(page *model.Page, race *model.Race) (races []*model.Race, err error)
	ListRaceBySearchTotalCount(race *model.Race) (count int64, err error)

	//Rule
	GetRule(rule *model.Rule) (err error)
	ListRule(page *model.Page) (rules []*model.Rule, err error)
	ListRuleTotalCount() (count int64, err error)
	ListRuleBySearch(page *model.Page, rule *model.Rule) (rules []*model.Rule, err error)
	ListRuleBySearchTotalCount(rule *model.Rule) (count int64, err error)

	//RuleEntry
	GetRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error)
	ListRuleEntry(page *model.Page, rule *model.Rule) (ruleEntrys []*model.RuleEntry, err error)
	ListRuleEntryTotalCount(rule *model.Rule) (count int64, err error)
	ListRuleEntryBySearch(page *model.Page, rule *model.Rule, ruleEntry *model.RuleEntry) (ruleEntrys []*model.RuleEntry, err error)
	ListRuleEntryBySearchTotalCount(rule *model.Rule, ruleEntry *model.RuleEntry) (count int64, err error)

	//Spawn
	GetSpawn(spawn *model.Spawn) (err error)
	ListSpawn(page *model.Page) (spawns []*model.Spawn, err error)
	ListSpawnTotalCount() (count int64, err error)
	ListSpawnBySearch(page *model.Page, spawn *model.Spawn) (spawns []*model.Spawn, err error)
	ListSpawnBySearchTotalCount(spawn *model.Spawn) (count int64, err error)

	//SpawnNpc
	GetSpawnNpc(spawn *model.Spawn, spawnNpc *model.SpawnNpc) (err error)
	ListSpawnNpc(page *model.Page, spawn *model.Spawn) (spawnNpcs []*model.SpawnNpc, err error)
	ListSpawnNpcTotalCount(spawn *model.Spawn) (count int64, err error)
	ListSpawnNpcBySearch(page *model.Page, spawn *model.Spawn, spawnNpc *model.SpawnNpc) (spawnNpcs []*model.SpawnNpc, err error)
	ListSpawnNpcBySearchTotalCount(spawn *model.Spawn, spawnNpc *model.SpawnNpc) (count int64, err error)

	//SpawnEntry
	GetSpawnEntry(spawn *model.Spawn, spawnEntry *model.SpawnEntry) (err error)
	ListSpawnEntry(page *model.Page, spawn *model.Spawn) (spawnEntrys []*model.SpawnEntry, err error)
	ListSpawnEntryTotalCount(spawn *model.Spawn) (count int64, err error)
	ListSpawnEntryBySearch(page *model.Page, spawn *model.Spawn, spawnEntry *model.SpawnEntry) (spawnEntrys []*model.SpawnEntry, err error)
	ListSpawnEntryBySearchTotalCount(spawn *model.Spawn, spawnEntry *model.SpawnEntry) (count int64, err error)

	//Spell
	GetSpell(spell *model.Spell) (err error)
	ListSpell(page *model.Page) (spells []*model.Spell, err error)
	ListSpellTotalCount() (count int64, err error)
	ListSpellBySearch(page *model.Page, spell *model.Spell) (spells []*model.Spell, err error)
	ListSpellBySearchTotalCount(spell *model.Spell) (count int64, err error)

	//SpellAnimation
	GetSpellAnimation(spellAnimation *model.SpellAnimation) (err error)
	ListSpellAnimation(page *model.Page) (spellAnimations []*model.SpellAnimation, err error)
	ListSpellAnimationTotalCount() (count int64, err error)
	ListSpellAnimationBySearch(page *model.Page, spellAnimation *model.SpellAnimation) (spellAnimations []*model.SpellAnimation, err error)
	ListSpellAnimationBySearchTotalCount(spellAnimation *model.SpellAnimation) (count int64, err error)

	//SpellAnimationType
	GetSpellAnimationType(spellAnimationType *model.SpellAnimationType) (err error)
	ListSpellAnimationType(page *model.Page) (spellAnimationTypes []*model.SpellAnimationType, err error)
	ListSpellAnimationTypeTotalCount() (count int64, err error)
	ListSpellAnimationTypeBySearch(page *model.Page, spellAnimationType *model.SpellAnimationType) (spellAnimationTypes []*model.SpellAnimationType, err error)
	ListSpellAnimationTypeBySearchTotalCount(spellAnimationType *model.SpellAnimationType) (count int64, err error)

	//SpellDurationFormula
	GetSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula) (err error)
	ListSpellDurationFormula(page *model.Page) (spellDurationFormulas []*model.SpellDurationFormula, err error)
	ListSpellDurationFormulaTotalCount() (count int64, err error)
	ListSpellDurationFormulaBySearch(page *model.Page, spellDurationFormula *model.SpellDurationFormula) (spellDurationFormulas []*model.SpellDurationFormula, err error)
	ListSpellDurationFormulaBySearchTotalCount(spellDurationFormula *model.SpellDurationFormula) (count int64, err error)

	//SpellEffectFormula
	GetSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula) (err error)
	ListSpellEffectFormula(page *model.Page) (spellEffectFormulas []*model.SpellEffectFormula, err error)
	ListSpellEffectFormulaTotalCount() (count int64, err error)
	ListSpellEffectFormulaBySearch(page *model.Page, spellEffectFormula *model.SpellEffectFormula) (spellEffectFormulas []*model.SpellEffectFormula, err error)
	ListSpellEffectFormulaBySearchTotalCount(spellEffectFormula *model.SpellEffectFormula) (count int64, err error)

	//SpellEffectType
	GetSpellEffectType(spellEffectType *model.SpellEffectType) (err error)
	ListSpellEffectType(page *model.Page) (spellEffectTypes []*model.SpellEffectType, err error)
	ListSpellEffectTypeTotalCount() (count int64, err error)
	ListSpellEffectTypeBySearch(page *model.Page, spellEffectType *model.SpellEffectType) (spellEffectTypes []*model.SpellEffectType, err error)
	ListSpellEffectTypeBySearchTotalCount(spellEffectType *model.SpellEffectType) (count int64, err error)

	//SpellTargetType
	GetSpellTargetType(spellTargetType *model.SpellTargetType) (err error)
	ListSpellTargetType(page *model.Page) (spellTargetTypes []*model.SpellTargetType, err error)
	ListSpellTargetTypeTotalCount() (count int64, err error)
	ListSpellTargetTypeBySearch(page *model.Page, spellTargetType *model.SpellTargetType) (spellTargetTypes []*model.SpellTargetType, err error)
	ListSpellTargetTypeBySearchTotalCount(spellTargetType *model.SpellTargetType) (count int64, err error)

	//SpellTravelType
	GetSpellTravelType(spellTravelType *model.SpellTravelType) (err error)
	ListSpellTravelType(page *model.Page) (spellTravelTypes []*model.SpellTravelType, err error)
	ListSpellTravelTypeTotalCount() (count int64, err error)
	ListSpellTravelTypeBySearch(page *model.Page, spellTravelType *model.SpellTravelType) (spellTravelTypes []*model.SpellTravelType, err error)
	ListSpellTravelTypeBySearchTotalCount(spellTravelType *model.SpellTravelType) (count int64, err error)

	//User
	LoginUser(user *model.User) (err error)
	GetUser(user *model.User) (err error)
	ListUser(page *model.Page) (users []*model.User, err error)
	ListUserTotalCount() (count int64, err error)
	ListUserBySearch(page *model.Page, user *model.User) (users []*model.User, err error)
	ListUserBySearchTotalCount(user *model.User) (count int64, err error)

	//UserAccount
	GetUserAccount(user *model.User, userAccount *model.UserAccount) (err error)
	ListUserAccount(page *model.Page, user *model.User) (userAccounts []*model.UserAccount, err error)
	ListUserAccountTotalCount(user *model.User) (count int64, err error)
	ListUserAccountBySearch(page *model.Page, user *model.User, userAccount *model.UserAccount) (userAccounts []*model.UserAccount, err error)
	ListUserAccountBySearchTotalCount(user *model.User, userAccount *model.UserAccount) (count int64, err error)

	//UserLink
	GetUserLink(userLink *model.UserLink) (err error)
	ListUserLink(page *model.Page) (userLinks []*model.UserLink, err error)
	ListUserLinkTotalCount() (count int64, err error)
	ListUserLinkBySearch(page *model.Page, userLink *model.UserLink) (userLinks []*model.UserLink, err error)
	ListUserLinkBySearchTotalCount(userLink *model.UserLink) (count int64, err error)

	//Variable
	GetVariable(variable *model.Variable) (err error)
	ListVariable(page *model.Page) (variables []*model.Variable, err error)
	ListVariableTotalCount() (count int64, err error)
	ListVariableBySearch(page *model.Page, variable *model.Variable) (variables []*model.Variable, err error)
	ListVariableBySearchTotalCount(variable *model.Variable) (count int64, err error)

	//Zone
	GetZone(zone *model.Zone) (err error)
	GetZoneByShortName(zone *model.Zone) (err error)
	ListZone(page *model.Page) (zones []*model.Zone, err error)
	ListZoneTotalCount() (count int64, err error)
	ListZoneBySearch(page *model.Page, zone *model.Zone) (zones []*model.Zone, err error)
	ListZoneBySearchTotalCount(zone *model.Zone) (count int64, err error)

	//ZoneExpansion
	GetZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error)
	ListZoneExpansion(page *model.Page) (zoneExpansions []*model.ZoneExpansion, err error)
	ListZoneExpansionTotalCount() (count int64, err error)
	ListZoneExpansionBySearch(page *model.Page, zoneExpansion *model.ZoneExpansion) (zoneExpansions []*model.ZoneExpansion, err error)
	ListZoneExpansionBySearchTotalCount(zoneExpansion *model.ZoneExpansion) (count int64, err error)
}
