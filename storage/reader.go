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

	//SpellTravelType
	GetSpellTravelType(spellTravelType *model.SpellTravelType) (err error)
	ListSpellTravelType(page *model.Page) (spellTravelTypes []*model.SpellTravelType, err error)
	ListSpellTravelTypeTotalCount() (count int64, err error)
	ListSpellTravelTypeBySearch(page *model.Page, spellTravelType *model.SpellTravelType) (spellTravelTypes []*model.SpellTravelType, err error)
	ListSpellTravelTypeBySearchTotalCount(spellTravelType *model.SpellTravelType) (count int64, err error)

	//User
	GetUser(user *model.User) (err error)
	ListUser(page *model.Page) (users []*model.User, err error)
	ListUserTotalCount() (count int64, err error)
	ListUserBySearch(page *model.Page, user *model.User) (users []*model.User, err error)
	ListUserBySearchTotalCount(user *model.User) (count int64, err error)

	//Variable
	GetVariable(variable *model.Variable) (err error)
	ListVariable(page *model.Page) (variables []*model.Variable, err error)
	ListVariableTotalCount() (count int64, err error)
	ListVariableBySearch(page *model.Page, variable *model.Variable) (variables []*model.Variable, err error)
	ListVariableBySearchTotalCount(variable *model.Variable) (count int64, err error)

	//Zone
	GetZone(zone *model.Zone) (err error)
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

	/*
		//Activity
		GetActivity(activity *model.Activity) (err error)
		GetActivityNextStep(activity *model.Activity) (step int64, err error)
		ListActivityByTask(task *model.Task) (activitys []*model.Activity, err error)

		//Base
		GetBase(base *model.Base) (err error)
		ListBase() (bases []*model.Base, err error)

		//Bazaar
		GetBazaar(bazaar *model.Bazaar) (err error)
		ListBazaar() (bazaars []*model.Bazaar, err error)

		//Character
		GetCharacter(character *model.Character) (err error)
		GetCharacterByName(character *model.Character) (err error)
		ListCharacter() (characters []*model.Character, err error)
		ListCharacterByRanking() (characters []*model.Character, err error)
		ListCharacterByOnline() (characters []*model.Character, err error)
		ListCharacterByAccount(account *model.Account) (characters []*model.Character, err error)
		SearchCharacterByName(character *model.Character) (characters []*model.Character, err error)

		//CharacterGraph
		GetCharacterGraph(characterGraph *model.CharacterGraph) (err error)
		ListCharacterGraphByCharacter(character *model.Character) (characterGraphs []*model.CharacterGraph, err error)

		//Error
		GetError(errStruct *model.Error) (err error)
		ListErrorCount() (count int64, err error)
		ListError(pageSize int64, pageNumber int64) (errors []*model.Error, err error)
		ListErrorByScope(errStruct *model.Error) (errors []*model.Error, err error)
		SearchErrorByMessage(errStruct *model.Error) (errors []*model.Error, err error)

		//Faction
		GetFaction(faction *model.Faction) (err error)
		ListFaction() (factions []*model.Faction, err error)

		//Fishing
		GetFishing(fishing *model.Fishing) (err error)
		ListFishing(pageSize int64, pageNumber int64) (fishings []*model.Fishing, err error)
		ListFishingCount() (count int64, err error)
		ListFishingByItem(item *model.Item) (fishings []*model.Fishing, err error)
		ListFishingByNpc(npc *model.Npc) (fishings []*model.Fishing, err error)
		ListFishingByZone(zone *model.Zone) (fishings []*model.Fishing, err error)

		//Forage
		GetForage(forage *model.Forage) (err error)
		ListForage(pageSize int64, pageNumber int64) (forages []*model.Forage, err error)
		ListForageCount() (count int64, err error)
		ListForageByItem(item *model.Item) (forages []*model.Forage, err error)
		ListForageByZone(zone *model.Zone) (forages []*model.Forage, err error)

		//Forum
		GetForum(forum *model.Forum) (err error)
		ListForum() (forums []*model.Forum, err error)

		//Goal
		GetGoal(goal *model.Goal) (err error)
		ListGoal() (goals []*model.Goal, err error)

		//Hacker
		GetHacker(hacker *model.Hacker) (err error)
		ListHacker(pageSize int64, pageNumber int64) (hackers []*model.Hacker, err error)
		ListHackerCount() (count int64, err error)
		SearchHackerByMessage(hacker *model.Hacker) (hackers []*model.Hacker, err error)

		//Item
		GetItem(item *model.Item) (err error)
		ListItem(pageSize int64, pageNumber int64) (items []*model.Item, err error)
		ListItemCount() (count int64, err error)
		SearchItemByName(item *model.Item) (items []*model.Item, err error)
		SearchItemByAccount(item *model.Item, account *model.Account) (items []*model.Item, err error)
		ListItemByCharacter(character *model.Character) (items []*model.Item, err error)
		ListItemByItemCategory(itemCategory *model.ItemCategory) (items []*model.Item, err error)
		ListItemBySpell(spell *model.Spell) (items []*model.Item, err error)
		ListItemByZone(zone *model.Zone) (items []*model.Item, err error)

		//LootDrop
		GetLootDrop(lootDrop *model.LootDrop) (err error)
		ListLootDrop() (lootDrops []*model.LootDrop, err error)

		//LootDropEntry
		GetLootDropEntry(lootDropEntry *model.LootDropEntry) (err error)
		ListLootDropEntryByLootDrop(lootDrop *model.LootDrop) (lootDropEntrys []*model.LootDropEntry, err error)

		//LootTable
		GetLootTable(lootTable *model.LootTable) (err error)
		ListLootTable() (lootTables []*model.LootTable, err error)

		//LootTableEntry
		GetLootTableEntry(lootTableEntry *model.LootTableEntry) (err error)
		ListLootTableEntryByLootTable(lootTable *model.LootTable) (lootTableEntrys []*model.LootTableEntry, err error)

		//Mail
		GetMail(mail *model.Mail) (err error)
		ListMail(pageSize int64, pageNumber int64) (mails []*model.Mail, err error)
		ListMailCount() (count int64, err error)
		SearchMailByBody(mail *model.Mail) (mails []*model.Mail, err error)
		SearchMailByCharacter(character *model.Character, mail *model.Mail) (mails []*model.Mail, err error)
		ListMailByCharacter(character *model.Character) (mails []*model.Mail, err error)

		//Merchant
		GetMerchant(merchant *model.Merchant) (err error)
		ListMerchant(pageSize int64, pageNumber int64) (merchants []*model.Merchant, err error)
		ListMerchantCount() (count int64, err error)

		//MerchantEntry
		GetMerchantEntry(merchantEntry *model.MerchantEntry) (err error)
		ListMerchantEntryByMerchant(merchant *model.Merchant) (merchantEntrys []*model.MerchantEntry, err error)
		ListMerchantEntryByItem(item *model.Item) (merchantEntrys []*model.MerchantEntry, err error)

		//Npc
		GetNpc(npc *model.Npc) (err error)
		ListNpc(pageSize int64, pageNumber int64) (npcs []*model.Npc, err error)
		ListNpcCount() (count int64, err error)
		ListNpcByZone(zone *model.Zone) (npcs []*model.Npc, err error)
		ListNpcByFaction(faction *model.Faction) (npcs []*model.Npc, err error)
		ListNpcByLootTable(lootTable *model.LootTable) (npcs []*model.Npc, err error)
		ListNpcByMerchant(merchant *model.Merchant) (npcs []*model.Npc, err error)
		ListNpcByItem(item *model.Item) (npcs []*model.Npc, err error)
		ListNpcBySpell(spell *model.Spell) (npcs []*model.Npc, err error)
		SearchNpcByName(npc *model.Npc) (npcs []*model.Npc, err error)

		//NpcLoot
		GetNpcLoot(npcLoot *model.NpcLoot) (err error)
		ListNpcLootByNpc(npc *model.Npc) (npcLoots []*model.NpcLoot, err error)
		ListNpcLootByZone(zone *model.Zone) (npcLoots []*model.NpcLoot, err error)
		TruncateNpcLoot() (err error)

		//Post
		GetPost(post *model.Post) (err error)
		ListPostByTopic(topic *model.Topic) (posts []*model.Post, err error)

		//Recipe
		GetRecipe(recipe *model.Recipe) (err error)
		ListRecipeBySkill(skill *model.Skill, pageSize int64, pageNumber int64) (recipes []*model.Recipe, err error)
		ListRecipeBySkillCount(skill *model.Skill) (count int64, err error)
		ListRecipe(pageSize int64, pageNumber int64) (recipes []*model.Recipe, err error)
		ListRecipeCount() (count int64, err error)
		SearchRecipeByName(recipe *model.Recipe) (recipes []*model.Recipe, err error)

		//RecipeEntry
		GetRecipeEntry(recipeEntry *model.RecipeEntry) (err error)
		ListRecipeEntryByRecipe(recipe *model.Recipe) (recipeEntrys []*model.RecipeEntry, err error)
		ListRecipeEntryByItem(item *model.Item) (recipeEntrys []*model.RecipeEntry, err error)

		//Rule
		GetRule(rule *model.Rule) (err error)
		ListRule() (rules []*model.Rule, err error)

		//SharedBank
		GetSharedBank(sharedBank *model.SharedBank) (err error)
		ListSharedBankByAccount(account *model.Account, pageSize int64, pageNumber int64) (sharedBanks []*model.SharedBank, err error)
		ListSharedBankByAccountCount(account *model.Account) (count int64, err error)
		ListSharedBankByAccountAndItem(account *model.Account, item *model.Item) (sharedBanks []*model.SharedBank, err error)

		//Spawn
		GetSpawn(spawn *model.Spawn) (err error)
		ListSpawn() (spawns []*model.Spawn, err error)
		ListSpawnEntryBySpawn(spawn *model.Spawn) (spawnEntrys []*model.SpawnEntry, err error)

		//SpawnEntry
		GetSpawnEntry(spawnEntry *model.SpawnEntry) (err error)
		ListSpawnEntry() (spawnEntrys []*model.SpawnEntry, err error)

		//SpawnNpc
		GetSpawnNpc(spawnNpc *model.SpawnNpc) (err error)
		ListSpawnNpc() (spawnNpcs []*model.SpawnNpc, err error)
		ListSpawnNpcBySpawn(spawn *model.Spawn) (spawnNpcs []*model.SpawnNpc, err error)
		ListSpawnNpcByNpc(npc *model.Npc) (spawnNpcs []*model.SpawnNpc, err error)

		//Spell
		GetSpell(spell *model.Spell) (err error)
		ListSpell(pageSize int64, pageNumber int64) (spells []*model.Spell, err error)
		ListSpellCount() (count int64, err error)
		SearchSpellByName(spell *model.Spell) (spells []*model.Spell, err error)

		//Task
		GetTask(task *model.Task) (err error)
		GetTaskNextID() (taskID int64, err error)
		ListTask() (tasks []*model.Task, err error)

		//Topic
		GetTopic(topic *model.Topic) (err error)
		ListTopicByForum(forum *model.Forum) (topics []*model.Topic, err error)

		//User
		GetUser(user *model.User) (err error)
		LoginUser(user *model.User, passwordConfirm string) (err error)
		ListUser() (users []*model.User, err error)

		//Variable
		GetVariable(variable *model.Variable) (err error)
		ListVariable() (variables []*model.Variable, err error)

		//Zone
		GetZone(zone *model.Zone) (err error)
		ListZone() (zones []*model.Zone, err error)
		ListZoneByHotzone() (zones []*model.Zone, err error)

		//ZoneLevel
		GetZoneLevel(zoneLevel *model.ZoneLevel) (err error)
		ListZoneLevel() (zoneLevels []*model.ZoneLevel, err error)
		TruncateZoneLevel() (err error)*/
}
