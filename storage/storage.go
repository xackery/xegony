package storage

import (
	"github.com/xackery/xegony/model"
)

//Storage is a generic interface of all storage types
type Storage interface {
	Initialize(config string) (err error)
	DropTables() (err error)
	VerifyTables() (err error)
	InsertTestData() (err error)
	//Account
	GetAccount(accountID int64) (account *model.Account, err error)
	CreateAccount(account *model.Account) (err error)
	EditAccount(accountID int64, account *model.Account) (err error)
	ListAccount() (accounts []*model.Account, err error)
	DeleteAccount(accountID int64) (err error)
	//Activity
	GetActivity(taskID int64, activityID int64) (activity *model.Activity, err error)
	CreateActivity(activity *model.Activity) (err error)
	EditActivity(activityID int64, activity *model.Activity) (err error)
	ListActivity(taskID int64) (activitys []*model.Activity, err error)
	DeleteActivity(activityID int64) (err error)
	//Base
	GetBase(level int64, class int64) (base *model.Base, err error)
	CreateBase(base *model.Base) (err error)
	ListBase() (bases []*model.Base, err error)
	EditBase(level int64, class int64, base *model.Base) (err error)
	DeleteBase(level int64, class int64) (err error)
	//Bazaar
	GetBazaar(bazaarID int64) (bazaar *model.Bazaar, err error)
	CreateBazaar(bazaar *model.Bazaar) (err error)
	EditBazaar(bazaarID int64, bazaar *model.Bazaar) (err error)
	ListBazaar() (bazaars []*model.Bazaar, err error)
	DeleteBazaar(bazaarID int64) (err error)
	//Faction
	GetFaction(accountID int64) (account *model.Faction, err error)
	CreateFaction(account *model.Faction) (err error)
	EditFaction(accountID int64, account *model.Faction) (err error)
	ListFaction() (accounts []*model.Faction, err error)
	DeleteFaction(accountID int64) (err error)
	//Goal
	GetGoal(listID int64, entryID int64) (goal *model.Goal, err error)
	CreateGoal(goal *model.Goal) (err error)
	ListGoal() (goals []*model.Goal, err error)
	EditGoal(listID int64, goal *model.Goal) (err error)
	DeleteGoal(listID int64, entryID int64) (err error)
	//Character
	GetCharacter(characterID int64) (character *model.Character, err error)
	CreateCharacter(character *model.Character) (err error)
	EditCharacter(characterID int64, character *model.Character) (err error)
	ListCharacter() (characters []*model.Character, err error)
	ListCharacterByAccount(accountID int64) (characters []*model.Character, err error)
	ListCharacterByOnline() (characters []*model.Character, err error)
	ListCharacterByRanking() (characters []*model.Character, err error)
	SearchCharacter(search string) (characters []*model.Character, err error)
	DeleteCharacter(characterID int64) (err error)
	//Item
	GetItem(itemID int64) (item *model.Item, err error)
	CreateItem(item *model.Item) (err error)
	EditItem(itemID int64, item *model.Item) (err error)
	ListItem(pageSize int64, pageNumber int64) (items []*model.Item, err error)
	ListItemCount() (count int64, err error)
	ListItemByCharacter(characterID int64) (items []*model.Item, err error)
	ListItemBySlot(slotID int64) (items []*model.Item, err error)
	ListItemByZone(zoneID int64) (items []*model.Item, err error)
	DeleteItem(itemID int64) (err error)
	SearchItem(search string) (items []*model.Item, err error)
	SearchItemByAccount(accountID int64, search string) (items []*model.Item, err error)
	//Npc
	GetNpc(npcID int64) (npc *model.Npc, err error)
	CreateNpc(npc *model.Npc) (err error)
	EditNpc(npcID int64, npc *model.Npc) (err error)
	ListNpc() (npcs []*model.Npc, err error)
	ListNpcByZone(zoneID int64) (npcs []*model.Npc, err error)
	ListNpcByItem(itemID int64) (npcs []*model.Npc, err error)
	ListNpcByFaction(factionID int64) (npcs []*model.Npc, err error)
	ListNpcByLootTable(lootTableID int64) (npcs []*model.Npc, err error)
	DeleteNpc(npcID int64) (err error)
	SearchNpc(search string) (npcs []*model.Npc, err error)
	//NpcLoot
	GetNpcLoot(npcID int64, itemID int64) (npcLoot *model.NpcLoot, err error)
	CreateNpcLoot(npcLoot *model.NpcLoot) (err error)
	ListNpcLoot(npcID int64) (npcLoots []*model.NpcLoot, err error)
	ListNpcLootByZone(zoneID int64) (npcLoots []*model.NpcLoot, err error)
	EditNpcLoot(npcID int64, itemID int64, npcLoot *model.NpcLoot) (err error)
	DeleteNpcLoot(npcID int64, itemID int64) (err error)
	TruncateNpcLoot() (err error)
	//Forum
	GetForum(forumID int64) (forum *model.Forum, err error)
	CreateForum(forum *model.Forum) (err error)
	EditForum(forumID int64, forum *model.Forum) (err error)
	ListForum() (forums []*model.Forum, err error)
	DeleteForum(forumID int64) (err error)
	//LootDrop
	GetLootDrop(lootDropID int64) (lootDrop *model.LootDrop, err error)
	CreateLootDrop(lootDrop *model.LootDrop) (err error)
	ListLootDrop() (lootDrops []*model.LootDrop, err error)
	EditLootDrop(lootDropID int64, lootDrop *model.LootDrop) (err error)
	DeleteLootDrop(lootDropID int64) (err error)
	//LootDropEntry
	GetLootDropEntry(lootDropID int64, itemID int64) (lootDropEntry *model.LootDropEntry, err error)
	CreateLootDropEntry(lootDropEntry *model.LootDropEntry) (err error)
	ListLootDropEntry(lootDropID int64) (lootDropEntrys []*model.LootDropEntry, err error)
	EditLootDropEntry(lootDropID int64, itemID int64, lootDropEntry *model.LootDropEntry) (err error)
	DeleteLootDropEntry(lootDropID int64, itemID int64) (err error)
	//LootTable
	GetLootTable(lootDropID int64) (lootDrop *model.LootTable, err error)
	CreateLootTable(lootDrop *model.LootTable) (err error)
	ListLootTable() (lootDrops []*model.LootTable, err error)
	EditLootTable(lootDropID int64, lootDrop *model.LootTable) (err error)
	DeleteLootTable(lootDropID int64) (err error)
	//LootTableEntry
	GetLootTableEntry(lootTableID int64, lootDropID int64) (lootTableEntry *model.LootTableEntry, err error)
	CreateLootTableEntry(lootTableEntry *model.LootTableEntry) (err error)
	ListLootTableEntry(lootTableID int64) (lootTableEntrys []*model.LootTableEntry, err error)
	EditLootTableEntry(lootTableID int64, lootDropID int64, lootTableEntry *model.LootTableEntry) (err error)
	DeleteLootTableEntry(lootTableID int64, lootDropID int64) (err error)
	//Post
	GetPost(postID int64) (post *model.Post, err error)
	CreatePost(post *model.Post) (err error)
	EditPost(postID int64, post *model.Post) (err error)
	ListPost(topicID int64) (posts []*model.Post, err error)
	DeletePost(postID int64) (err error)
	//Spawn
	GetSpawn(spawnID int64) (spawn *model.Spawn, err error)
	CreateSpawn(spawn *model.Spawn) (err error)
	ListSpawn() (spawns []*model.Spawn, err error)
	EditSpawn(spawnID int64, spawn *model.Spawn) (err error)
	DeleteSpawn(spawnID int64) (err error)
	//SpawnEntry
	GetSpawnEntry(spawnGroupID int64, npcID int64) (query string, spawnEntry *model.SpawnEntry, err error)
	CreateSpawnEntry(spawnEntry *model.SpawnEntry) (query string, err error)
	ListSpawnEntry(spawnGroupID int64) (query string, spawnEntrys []*model.SpawnEntry, err error)
	ListSpawnEntryByZone(zoneID int64) (query string, spawnEntrys []*model.SpawnEntry, err error)
	EditSpawnEntry(spawnGroupID int64, npcID int64, spawnEntry *model.SpawnEntry) (query string, err error)
	DeleteSpawnEntry(spawnGroupID int64, npcID int64) (query string, err error)
	//Topic
	GetTopic(topicID int64) (topic *model.Topic, err error)
	CreateTopic(topic *model.Topic) (err error)
	EditTopic(topicID int64, topic *model.Topic) (err error)
	ListTopic(forumID int64) (topics []*model.Topic, err error)
	DeleteTopic(topicID int64) (err error)
	//Zone
	GetZone(zoneID int64) (zone *model.Zone, err error)
	CreateZone(zone *model.Zone) (err error)
	EditZone(zoneID int64, zone *model.Zone) (err error)
	ListZone() (zones []*model.Zone, err error)
	ListZoneByHotzone() (zones []*model.Zone, err error)
	DeleteZone(zoneID int64) (err error)
	//ZoneLevel
	GetZoneLevel(zoneID int64) (zoneLevel *model.ZoneLevel, err error)
	CreateZoneLevel(zoneLevel *model.ZoneLevel) (err error)
	ListZoneLevel() (zoneLevels []*model.ZoneLevel, err error)
	EditZoneLevel(zoneID int64, zoneLevel *model.ZoneLevel) (err error)
	TruncateZoneLevel() (err error)
	DeleteZoneLevel(zoneID int64) (err error)
	//Task
	GetTask(taskID int64) (task *model.Task, err error)
	CreateTask(task *model.Task) (err error)
	EditTask(taskID int64, task *model.Task) (err error)
	ListTask() (tasks []*model.Task, err error)
	DeleteTask(taskID int64) (err error)
	//User
	GetUser(userID int64) (user *model.User, err error)
	CreateUser(user *model.User) (err error)
	EditUser(userID int64, user *model.User) (err error)
	ListUser() (users []*model.User, err error)
	DeleteUser(userID int64) (err error)
	LoginUser(username string, password string) (user *model.User, err error)
}
