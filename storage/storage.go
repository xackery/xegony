package storage

import (
	"io"

	"github.com/xackery/xegony/model"
)

//Storage is a generic interface of all storage types
type Storage interface {
	Initialize(config string, w io.Writer) (err error)
	DropTables() (err error)
	VerifyTables() (err error)
	InsertTestData() (err error)
	//Aa
	GetAa(aa *model.Aa) (err error)
	CreateAa(aa *model.Aa) (err error)
	ListAa() (aas []*model.Aa, err error)
	EditAa(aa *model.Aa) (err error)
	DeleteAa(aa *model.Aa) (err error)

	//AaRank
	GetAaRank(aaRank *model.AaRank) (err error)
	CreateAaRank(aaRank *model.AaRank) (err error)
	ListAaRank() (aaRanks []*model.AaRank, err error)
	EditAaRank(aaRank *model.AaRank) (err error)
	DeleteAaRank(aaRank *model.AaRank) (err error)

	//Account
	GetAccount(account *model.Account) (err error)
	GetAccountByName(account *model.Account) (err error)
	CreateAccount(account *model.Account) (err error)
	ListAccount() (accounts []*model.Account, err error)
	EditAccount(account *model.Account) (err error)
	DeleteAccount(account *model.Account) (err error)

	//Activity
	GetActivity(activity *model.Activity) (err error)
	GetActivityNextStep(activity *model.Activity) (step int64, err error)
	CreateActivity(activity *model.Activity) (err error)
	ListActivityByTask(task *model.Task) (activitys []*model.Activity, err error)
	EditActivity(activity *model.Activity) (err error)
	DeleteActivity(activity *model.Activity) (err error)

	//Base
	GetBase(base *model.Base) (err error)
	CreateBase(base *model.Base) (err error)
	ListBase() (bases []*model.Base, err error)
	EditBase(base *model.Base) (err error)
	DeleteBase(base *model.Base) (err error)

	//Bazaar
	GetBazaar(bazaar *model.Bazaar) (err error)
	CreateBazaar(bazaar *model.Bazaar) (err error)
	ListBazaar() (bazaars []*model.Bazaar, err error)
	EditBazaar(bazaar *model.Bazaar) (err error)
	DeleteBazaar(bazaar *model.Bazaar) (err error)

	//Character
	GetCharacter(character *model.Character) (err error)
	GetCharacterByName(character *model.Character) (err error)
	CreateCharacter(character *model.Character) (err error)
	ListCharacter() (characters []*model.Character, err error)
	ListCharacterByRanking() (characters []*model.Character, err error)
	ListCharacterByOnline() (characters []*model.Character, err error)
	ListCharacterByAccount(account *model.Account) (characters []*model.Character, err error)
	EditCharacter(character *model.Character) (err error)
	DeleteCharacter(character *model.Character) (err error)
	SearchCharacterByName(character *model.Character) (characters []*model.Character, err error)

	//CharacterGraph
	GetCharacterGraph(characterGraph *model.CharacterGraph) (err error)
	CreateCharacterGraph(characterGraph *model.CharacterGraph) (err error)
	ListCharacterGraphByCharacter(character *model.Character) (characterGraphs []*model.CharacterGraph, err error)
	EditCharacterGraph(characterGraph *model.CharacterGraph) (err error)
	DeleteCharacterGraph(characterGraph *model.CharacterGraph) (err error)

	//Error
	GetError(errStruct *model.Error) (err error)
	CreateError(errStruct *model.Error) (err error)
	ListErrorCount() (count int64, err error)
	ListError(pageSize int64, pageNumber int64) (errors []*model.Error, err error)
	ListErrorByScope(errStruct *model.Error) (errors []*model.Error, err error)
	SearchErrorByMessage(errStruct *model.Error) (errors []*model.Error, err error)
	EditError(errorStruct *model.Error) (err error)
	DeleteError(errStruct *model.Error) (err error)

	//Faction
	GetFaction(faction *model.Faction) (err error)
	CreateFaction(faction *model.Faction) (err error)
	ListFaction() (factions []*model.Faction, err error)
	EditFaction(faction *model.Faction) (err error)
	DeleteFaction(faction *model.Faction) (err error)

	//Fishing
	GetFishing(fishing *model.Fishing) (err error)
	CreateFishing(fishing *model.Fishing) (err error)
	ListFishing(pageSize int64, pageNumber int64) (fishings []*model.Fishing, err error)
	ListFishingCount() (count int64, err error)
	ListFishingByItem(item *model.Item) (fishings []*model.Fishing, err error)
	ListFishingByNpc(npc *model.Npc) (fishings []*model.Fishing, err error)
	ListFishingByZone(zone *model.Zone) (fishings []*model.Fishing, err error)
	EditFishing(fishing *model.Fishing) (err error)
	DeleteFishing(fishing *model.Fishing) (err error)

	//Forage
	GetForage(forage *model.Forage) (err error)
	CreateForage(forage *model.Forage) (err error)
	ListForage(pageSize int64, pageNumber int64) (forages []*model.Forage, err error)
	ListForageCount() (count int64, err error)
	ListForageByItem(item *model.Item) (forages []*model.Forage, err error)
	ListForageByZone(zone *model.Zone) (forages []*model.Forage, err error)
	EditForage(forage *model.Forage) (err error)
	DeleteForage(forage *model.Forage) (err error)

	//Forum
	GetForum(forum *model.Forum) (err error)
	CreateForum(forum *model.Forum) (err error)
	ListForum() (forums []*model.Forum, err error)
	EditForum(forum *model.Forum) (err error)
	DeleteForum(forum *model.Forum) (err error)

	//Goal
	GetGoal(goal *model.Goal) (err error)
	CreateGoal(goal *model.Goal) (err error)
	ListGoal() (goals []*model.Goal, err error)
	EditGoal(goal *model.Goal) (err error)
	DeleteGoal(goal *model.Goal) (err error)

	//Hacker
	GetHacker(hacker *model.Hacker) (err error)
	CreateHacker(hacker *model.Hacker) (err error)
	ListHacker(pageSize int64, pageNumber int64) (hackers []*model.Hacker, err error)
	ListHackerCount() (count int64, err error)
	SearchHackerByMessage(hacker *model.Hacker) (hackers []*model.Hacker, err error)
	EditHacker(hacker *model.Hacker) (err error)
	DeleteHacker(hacker *model.Hacker) (err error)

	//Item
	GetItem(item *model.Item) (err error)
	CreateItem(item *model.Item) (err error)
	ListItem(pageSize int64, pageNumber int64) (items []*model.Item, err error)
	ListItemCount() (count int64, err error)
	SearchItemByName(item *model.Item) (items []*model.Item, err error)
	SearchItemByAccount(item *model.Item, account *model.Account) (items []*model.Item, err error)
	ListItemByCharacter(character *model.Character) (items []*model.Item, err error)
	ListItemByItemCategory(itemCategory *model.ItemCategory) (items []*model.Item, err error)
	ListItemBySpell(spell *model.Spell) (items []*model.Item, err error)
	ListItemByZone(zone *model.Zone) (items []*model.Item, err error)
	EditItem(item *model.Item) (err error)
	DeleteItem(item *model.Item) (err error)

	//LootDrop
	GetLootDrop(lootDrop *model.LootDrop) (err error)
	CreateLootDrop(lootDrop *model.LootDrop) (err error)
	ListLootDrop() (lootDrops []*model.LootDrop, err error)
	EditLootDrop(lootDrop *model.LootDrop) (err error)
	DeleteLootDrop(lootDrop *model.LootDrop) (err error)

	//LootDropEntry
	GetLootDropEntry(lootDropEntry *model.LootDropEntry) (err error)
	CreateLootDropEntry(lootDropEntry *model.LootDropEntry) (err error)
	ListLootDropEntryByLootDrop(lootDrop *model.LootDrop) (lootDropEntrys []*model.LootDropEntry, err error)
	EditLootDropEntry(lootDropEntry *model.LootDropEntry) (err error)
	DeleteLootDropEntry(lootDropEntry *model.LootDropEntry) (err error)

	//LootTable
	GetLootTable(lootTable *model.LootTable) (err error)
	CreateLootTable(lootTable *model.LootTable) (err error)
	ListLootTable() (lootTables []*model.LootTable, err error)
	EditLootTable(lootTable *model.LootTable) (err error)
	DeleteLootTable(lootTable *model.LootTable) (err error)

	//LootTableEntry
	GetLootTableEntry(lootTableEntry *model.LootTableEntry) (err error)
	CreateLootTableEntry(lootTableEntry *model.LootTableEntry) (err error)
	ListLootTableEntryByLootTable(lootTable *model.LootTable) (lootTableEntrys []*model.LootTableEntry, err error)
	EditLootTableEntry(lootTableEntry *model.LootTableEntry) (err error)
	DeleteLootTableEntry(lootTableEntry *model.LootTableEntry) (err error)

	//Mail
	GetMail(mail *model.Mail) (err error)
	CreateMail(mail *model.Mail) (err error)
	ListMail(pageSize int64, pageNumber int64) (mails []*model.Mail, err error)
	ListMailCount() (count int64, err error)
	SearchMailByBody(mail *model.Mail) (mails []*model.Mail, err error)
	SearchMailByCharacter(character *model.Character, mail *model.Mail) (mails []*model.Mail, err error)
	ListMailByCharacter(character *model.Character) (mails []*model.Mail, err error)
	EditMail(mail *model.Mail) (err error)
	DeleteMail(mail *model.Mail) (err error)

	//Merchant
	GetMerchant(merchant *model.Merchant) (err error)
	ListMerchant(pageSize int64, pageNumber int64) (merchants []*model.Merchant, err error)
	ListMerchantCount() (count int64, err error)
	DeleteMerchant(merchant *model.Merchant) (err error)

	//MerchantEntry
	GetMerchantEntry(merchantEntry *model.MerchantEntry) (err error)
	CreateMerchantEntry(merchantEntry *model.MerchantEntry) (err error)
	ListMerchantEntryByMerchant(merchant *model.Merchant) (merchantEntrys []*model.MerchantEntry, err error)
	ListMerchantEntryByItem(item *model.Item) (merchantEntrys []*model.MerchantEntry, err error)
	EditMerchantEntry(merchantEntry *model.MerchantEntry) (err error)
	DeleteMerchantEntry(merchantEntry *model.MerchantEntry) (err error)

	//Npc
	GetNpc(npc *model.Npc) (err error)
	CreateNpc(npc *model.Npc) (err error)
	ListNpc(pageSize int64, pageNumber int64) (npcs []*model.Npc, err error)
	ListNpcCount() (count int64, err error)
	ListNpcByZone(zone *model.Zone) (npcs []*model.Npc, err error)
	ListNpcByFaction(faction *model.Faction) (npcs []*model.Npc, err error)
	ListNpcByLootTable(lootTable *model.LootTable) (npcs []*model.Npc, err error)
	ListNpcByMerchant(merchant *model.Merchant) (npcs []*model.Npc, err error)
	ListNpcByItem(item *model.Item) (npcs []*model.Npc, err error)
	ListNpcBySpell(spell *model.Spell) (npcs []*model.Npc, err error)
	EditNpc(npc *model.Npc) (err error)
	DeleteNpc(npc *model.Npc) (err error)
	SearchNpcByName(npc *model.Npc) (npcs []*model.Npc, err error)

	//NpcLoot
	GetNpcLoot(npcLoot *model.NpcLoot) (err error)
	CreateNpcLoot(npcLoot *model.NpcLoot) (err error)
	ListNpcLootByNpc(npc *model.Npc) (npcLoots []*model.NpcLoot, err error)
	ListNpcLootByZone(zone *model.Zone) (npcLoots []*model.NpcLoot, err error)
	EditNpcLoot(npcLoot *model.NpcLoot) (err error)
	TruncateNpcLoot() (err error)
	DeleteNpcLoot(npcLoot *model.NpcLoot) (err error)

	//Post
	GetPost(post *model.Post) (err error)
	CreatePost(post *model.Post) (err error)
	ListPostByTopic(topic *model.Topic) (posts []*model.Post, err error)
	EditPost(post *model.Post) (err error)
	DeletePost(post *model.Post) (err error)

	//Recipe
	GetRecipe(recipe *model.Recipe) (err error)
	CreateRecipe(recipe *model.Recipe) (err error)
	ListRecipeBySkill(skill *model.Skill, pageSize int64, pageNumber int64) (recipes []*model.Recipe, err error)
	ListRecipeBySkillCount(skill *model.Skill) (count int64, err error)
	ListRecipe(pageSize int64, pageNumber int64) (recipes []*model.Recipe, err error)
	ListRecipeCount() (count int64, err error)
	SearchRecipeByName(recipe *model.Recipe) (recipes []*model.Recipe, err error)
	EditRecipe(recipe *model.Recipe) (err error)
	DeleteRecipe(recipe *model.Recipe) (err error)

	//RecipeEntry
	GetRecipeEntry(recipeEntry *model.RecipeEntry) (err error)
	CreateRecipeEntry(recipeEntry *model.RecipeEntry) (err error)
	ListRecipeEntryByRecipe(recipe *model.Recipe) (recipeEntrys []*model.RecipeEntry, err error)
	ListRecipeEntryByItem(item *model.Item) (recipeEntrys []*model.RecipeEntry, err error)
	EditRecipeEntry(recipeEntry *model.RecipeEntry) (err error)
	DeleteRecipeEntry(recipeEntry *model.RecipeEntry) (err error)

	//Rule
	GetRule(rule *model.Rule) (err error)
	CreateRule(rule *model.Rule) (err error)
	ListRule() (rules []*model.Rule, err error)
	EditRule(rule *model.Rule) (err error)
	DeleteRule(rule *model.Rule) (err error)

	//SharedBank
	GetSharedBank(sharedBank *model.SharedBank) (err error)
	CreateSharedBank(sharedBank *model.SharedBank) (err error)
	ListSharedBankByAccount(account *model.Account, pageSize int64, pageNumber int64) (sharedBanks []*model.SharedBank, err error)
	ListSharedBankByAccountCount(account *model.Account) (count int64, err error)
	ListSharedBankByAccountAndItem(account *model.Account, item *model.Item) (sharedBanks []*model.SharedBank, err error)
	EditSharedBank(sharedBank *model.SharedBank) (err error)
	DeleteSharedBank(sharedBank *model.SharedBank) (err error)

	//Spawn
	GetSpawn(spawn *model.Spawn) (err error)
	CreateSpawn(spawn *model.Spawn) (err error)
	ListSpawnBySpawnGroup(spawnGroup *model.SpawnGroup) (spawns []*model.Spawn, err error)
	ListSpawn() (spawns []*model.Spawn, err error)
	EditSpawn(spawn *model.Spawn) (err error)
	DeleteSpawn(spawn *model.Spawn) (err error)

	//SpawnEntry
	GetSpawnEntry(spawnEntry *model.SpawnEntry) (err error)
	CreateSpawnEntry(spawnEntry *model.SpawnEntry) (err error)
	ListSpawnEntryBySpawnGroup(spawnGroup *model.SpawnGroup) (spawnEntrys []*model.SpawnEntry, err error)
	ListSpawnEntryByZone(zone *model.Zone) (spawnEntrys []*model.SpawnEntry, err error)
	ListSpawnEntryByNpc(npc *model.Npc) (spawnEntrys []*model.SpawnEntry, err error)
	EditSpawnEntry(spawnEntry *model.SpawnEntry) (err error)
	DeleteSpawnEntry(spawnEntry *model.SpawnEntry) (err error)

	//Spell
	GetSpell(spell *model.Spell) (err error)
	CreateSpell(spell *model.Spell) (err error)
	ListSpell(pageSize int64, pageNumber int64) (spells []*model.Spell, err error)
	ListSpellCount() (count int64, err error)
	SearchSpellByName(spell *model.Spell) (spells []*model.Spell, err error)
	EditSpell(spell *model.Spell) (err error)
	DeleteSpell(spell *model.Spell) (err error)

	//Task
	GetTask(task *model.Task) (err error)
	GetTaskNextID() (taskID int64, err error)
	CreateTask(task *model.Task) (err error)
	ListTask() (tasks []*model.Task, err error)
	EditTask(task *model.Task) (err error)
	DeleteTask(task *model.Task) (err error)

	//Topic
	GetTopic(topic *model.Topic) (err error)
	CreateTopic(topic *model.Topic) (err error)
	ListTopicByForum(forum *model.Forum) (topics []*model.Topic, err error)
	EditTopic(topic *model.Topic) (err error)
	DeleteTopic(topic *model.Topic) (err error)

	//User
	GetUser(user *model.User) (err error)
	LoginUser(user *model.User, passwordConfirm string) (err error)
	CreateUser(user *model.User) (err error)
	ListUser() (users []*model.User, err error)
	EditUser(user *model.User) (err error)
	DeleteUser(user *model.User) (err error)

	//Variable
	GetVariable(variable *model.Variable) (err error)
	CreateVariable(variable *model.Variable) (err error)
	ListVariable() (variables []*model.Variable, err error)
	EditVariable(variable *model.Variable) (err error)
	DeleteVariable(variable *model.Variable) (err error)

	//Zone
	GetZone(zone *model.Zone) (err error)
	CreateZone(zone *model.Zone) (err error)
	ListZone() (zones []*model.Zone, err error)
	ListZoneByHotzone() (zones []*model.Zone, err error)
	EditZone(zone *model.Zone) (err error)
	DeleteZone(zone *model.Zone) (err error)

	//ZoneLevel
	GetZoneLevel(zoneLevel *model.ZoneLevel) (err error)
	CreateZoneLevel(zoneLevel *model.ZoneLevel) (err error)
	ListZoneLevel() (zoneLevels []*model.ZoneLevel, err error)
	EditZoneLevel(zoneLevel *model.ZoneLevel) (err error)
	TruncateZoneLevel() (err error)
	DeleteZoneLevel(zoneLevel *model.ZoneLevel) (err error)
}
