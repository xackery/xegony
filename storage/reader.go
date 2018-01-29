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
