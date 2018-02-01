package storage

import (
	"github.com/xackery/xegony/model"
)

//Writer is a generic interface of all storage types
type Writer interface {
	/*//Aa
	CreateAa(aa *model.Aa) (err error)
	EditAa(aa *model.Aa) (err error)
	DeleteAa(aa *model.Aa) (err error)

	//AaRank
	CreateAaRank(aaRank *model.AaRank) (err error)
	EditAaRank(aaRank *model.AaRank) (err error)
	DeleteAaRank(aaRank *model.AaRank) (err error)
	*/
	//Account
	CreateAccount(account *model.Account) (err error)
	EditAccount(account *model.Account) (err error)
	DeleteAccount(account *model.Account) (err error)

	//Character
	CreateCharacter(character *model.Character) (err error)
	EditCharacter(character *model.Character) (err error)
	DeleteCharacter(character *model.Character) (err error)

	//Race
	CreateRace(race *model.Race) (err error)
	EditRace(race *model.Race) (err error)
	DeleteRace(race *model.Race) (err error)

	//User
	CreateUser(user *model.User) (err error)
	EditUser(user *model.User) (err error)
	DeleteUser(user *model.User) (err error)

	//Zone
	CreateZone(zone *model.Zone) (err error)
	EditZone(zone *model.Zone) (err error)
	DeleteZone(zone *model.Zone) (err error)

	//ZoneExpansion
	CreateZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error)
	EditZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error)
	DeleteZoneExpansion(zoneExpansion *model.ZoneExpansion) (err error)

	/*
		//Activity
		CreateActivity(activity *model.Activity) (err error)
		EditActivity(activity *model.Activity) (err error)
		DeleteActivity(activity *model.Activity) (err error)

		//Base
		CreateBase(base *model.Base) (err error)
		EditBase(base *model.Base) (err error)
		DeleteBase(base *model.Base) (err error)

		//Bazaar
		CreateBazaar(bazaar *model.Bazaar) (err error)
		EditBazaar(bazaar *model.Bazaar) (err error)
		DeleteBazaar(bazaar *model.Bazaar) (err error)

		//Character
		CreateCharacter(character *model.Character) (err error)
		EditCharacter(character *model.Character) (err error)
		DeleteCharacter(character *model.Character) (err error)
		SearchCharacterByName(character *model.Character) (characters []*model.Character, err error)

		//CharacterGraph
		CreateCharacterGraph(characterGraph *model.CharacterGraph) (err error)
		EditCharacterGraph(characterGraph *model.CharacterGraph) (err error)
		DeleteCharacterGraph(characterGraph *model.CharacterGraph) (err error)

		//Error
		CreateError(errStruct *model.Error) (err error)
		SearchErrorByMessage(errStruct *model.Error) (errors []*model.Error, err error)
		EditError(errorStruct *model.Error) (err error)
		DeleteError(errStruct *model.Error) (err error)

		//Faction
		CreateFaction(faction *model.Faction) (err error)
		EditFaction(faction *model.Faction) (err error)
		DeleteFaction(faction *model.Faction) (err error)

		//Fishing
		CreateFishing(fishing *model.Fishing) (err error)
		EditFishing(fishing *model.Fishing) (err error)
		DeleteFishing(fishing *model.Fishing) (err error)

		//Forage
		CreateForage(forage *model.Forage) (err error)
		EditForage(forage *model.Forage) (err error)
		DeleteForage(forage *model.Forage) (err error)

		//Forum
		CreateForum(forum *model.Forum) (err error)
		EditForum(forum *model.Forum) (err error)
		DeleteForum(forum *model.Forum) (err error)

		//Goal
		CreateGoal(goal *model.Goal) (err error)
		EditGoal(goal *model.Goal) (err error)
		DeleteGoal(goal *model.Goal) (err error)

		//Hacker
		CreateHacker(hacker *model.Hacker) (err error)
		SearchHackerByMessage(hacker *model.Hacker) (hackers []*model.Hacker, err error)
		EditHacker(hacker *model.Hacker) (err error)
		DeleteHacker(hacker *model.Hacker) (err error)

		//Item
		CreateItem(item *model.Item) (err error)
		SearchItemByName(item *model.Item) (items []*model.Item, err error)
		SearchItemByAccount(item *model.Item, account *model.Account) (items []*model.Item, err error)
		EditItem(item *model.Item) (err error)
		DeleteItem(item *model.Item) (err error)

		//LootDrop
		CreateLootDrop(lootDrop *model.LootDrop) (err error)
		EditLootDrop(lootDrop *model.LootDrop) (err error)
		DeleteLootDrop(lootDrop *model.LootDrop) (err error)

		//LootDropEntry
		CreateLootDropEntry(lootDropEntry *model.LootDropEntry) (err error)
		EditLootDropEntry(lootDropEntry *model.LootDropEntry) (err error)
		DeleteLootDropEntry(lootDropEntry *model.LootDropEntry) (err error)

		//LootTable
		CreateLootTable(lootTable *model.LootTable) (err error)
		EditLootTable(lootTable *model.LootTable) (err error)
		DeleteLootTable(lootTable *model.LootTable) (err error)

		//LootTableEntry
		CreateLootTableEntry(lootTableEntry *model.LootTableEntry) (err error)
		EditLootTableEntry(lootTableEntry *model.LootTableEntry) (err error)
		DeleteLootTableEntry(lootTableEntry *model.LootTableEntry) (err error)

		//Mail
		CreateMail(mail *model.Mail) (err error)
		SearchMailByBody(mail *model.Mail) (mails []*model.Mail, err error)
		SearchMailByCharacter(character *model.Character, mail *model.Mail) (mails []*model.Mail, err error)
		EditMail(mail *model.Mail) (err error)
		DeleteMail(mail *model.Mail) (err error)

		//Merchant
		DeleteMerchant(merchant *model.Merchant) (err error)

		//MerchantEntry
		CreateMerchantEntry(merchantEntry *model.MerchantEntry) (err error)
		EditMerchantEntry(merchantEntry *model.MerchantEntry) (err error)
		DeleteMerchantEntry(merchantEntry *model.MerchantEntry) (err error)

		//Npc
		CreateNpc(npc *model.Npc) (err error)
		EditNpc(npc *model.Npc) (err error)
		DeleteNpc(npc *model.Npc) (err error)
		SearchNpcByName(npc *model.Npc) (npcs []*model.Npc, err error)

		//NpcLoot
		CreateNpcLoot(npcLoot *model.NpcLoot) (err error)
		EditNpcLoot(npcLoot *model.NpcLoot) (err error)
		TruncateNpcLoot() (err error)
		DeleteNpcLoot(npcLoot *model.NpcLoot) (err error)

		//Post
		CreatePost(post *model.Post) (err error)
		EditPost(post *model.Post) (err error)
		DeletePost(post *model.Post) (err error)

		//Recipe
		CreateRecipe(recipe *model.Recipe) (err error)
		SearchRecipeByName(recipe *model.Recipe) (recipes []*model.Recipe, err error)
		EditRecipe(recipe *model.Recipe) (err error)
		DeleteRecipe(recipe *model.Recipe) (err error)

		//RecipeEntry
		CreateRecipeEntry(recipeEntry *model.RecipeEntry) (err error)
		EditRecipeEntry(recipeEntry *model.RecipeEntry) (err error)
		DeleteRecipeEntry(recipeEntry *model.RecipeEntry) (err error)

		//Rule
		CreateRule(rule *model.Rule) (err error)
		EditRule(rule *model.Rule) (err error)
		DeleteRule(rule *model.Rule) (err error)

		//SharedBank
		CreateSharedBank(sharedBank *model.SharedBank) (err error)
		EditSharedBank(sharedBank *model.SharedBank) (err error)
		DeleteSharedBank(sharedBank *model.SharedBank) (err error)

		//Spawn
		CreateSpawn(spawn *model.Spawn) (err error)
		EditSpawn(spawn *model.Spawn) (err error)
		DeleteSpawn(spawn *model.Spawn) (err error)

		//SpawnEntry
		CreateSpawnEntry(spawnEntry *model.SpawnEntry) (err error)
		EditSpawnEntry(spawnEntry *model.SpawnEntry) (err error)
		DeleteSpawnEntry(spawnEntry *model.SpawnEntry) (err error)

		//SpawnNpc
		CreateSpawnNpc(spawnNpc *model.SpawnNpc) (err error)
		EditSpawnNpc(spawnNpc *model.SpawnNpc) (err error)
		DeleteSpawnNpc(spawnNpc *model.SpawnNpc) (err error)

		//Spell
		CreateSpell(spell *model.Spell) (err error)
		SearchSpellByName(spell *model.Spell) (spells []*model.Spell, err error)
		EditSpell(spell *model.Spell) (err error)
		DeleteSpell(spell *model.Spell) (err error)

		//Task
		CreateTask(task *model.Task) (err error)
		EditTask(task *model.Task) (err error)
		DeleteTask(task *model.Task) (err error)

		//Topic
		CreateTopic(topic *model.Topic) (err error)
		EditTopic(topic *model.Topic) (err error)
		DeleteTopic(topic *model.Topic) (err error)

		//User
		LoginUser(user *model.User, passwordConfirm string) (err error)
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

		//ZoneLevel
		CreateZoneLevel(zoneLevel *model.ZoneLevel) (err error)
		EditZoneLevel(zoneLevel *model.ZoneLevel) (err error)
		TruncateZoneLevel() (err error)
		DeleteZoneLevel(zoneLevel *model.ZoneLevel) (err error)*/
}
