package storage

import (
	"github.com/xackery/xegony/model"
)

type Storage interface {
	Initialize(config string) (err error)
	DropTables() (err error)
	VerifyTables() (err error)
	//Account
	GetAccount(accountId int64) (account *model.Account, err error)
	CreateAccount(account *model.Account) (err error)
	EditAccount(accountId int64, account *model.Account) (err error)
	ListAccount() (accounts []*model.Account, err error)
	DeleteAccount(accountId int64) (err error)
	//Faction
	GetFaction(accountId int64) (account *model.Faction, err error)
	CreateFaction(account *model.Faction) (err error)
	EditFaction(accountId int64, account *model.Faction) (err error)
	ListFaction() (accounts []*model.Faction, err error)
	DeleteFaction(accountId int64) (err error)
	//Character
	GetCharacter(characterId int64) (character *model.Character, err error)
	CreateCharacter(character *model.Character) (err error)
	EditCharacter(characterId int64, character *model.Character) (err error)
	ListCharacter() (characters []*model.Character, err error)
	DeleteCharacter(characterId int64) (err error)
	//Item
	GetItem(itemId int64) (item *model.Item, err error)
	CreateItem(item *model.Item) (err error)
	EditItem(itemId int64, item *model.Item) (err error)
	ListItem() (items []*model.Item, err error)
	DeleteItem(itemId int64) (err error)
	//Npc
	GetNpc(npcId int64) (npc *model.Npc, err error)
	CreateNpc(npc *model.Npc) (err error)
	EditNpc(npcId int64, npc *model.Npc) (err error)
	ListNpc() (npcs []*model.Npc, err error)
	ListNpcByZone(zoneId int64) (npcs []*model.Npc, err error)
	ListNpcByFaction(factionId int64) (npcs []*model.Npc, err error)
	DeleteNpc(npcId int64) (err error)
	//Forum
	GetForum(forumId int64) (forum *model.Forum, err error)
	CreateForum(forum *model.Forum) (err error)
	EditForum(forumId int64, forum *model.Forum) (err error)
	ListForum() (forums []*model.Forum, err error)
	DeleteForum(forumId int64) (err error)
	//Topic
	GetTopic(topicId int64) (topic *model.Topic, err error)
	CreateTopic(topic *model.Topic) (err error)
	EditTopic(topicId int64, topic *model.Topic) (err error)
	ListTopic() (topics []*model.Topic, err error)
	DeleteTopic(topicId int64) (err error)
	//Zone
	GetZone(zoneId int64) (zone *model.Zone, err error)
	CreateZone(zone *model.Zone) (err error)
	EditZone(zoneId int64, zone *model.Zone) (err error)
	ListZone() (zones []*model.Zone, err error)
	DeleteZone(zoneId int64) (err error)
	//User
	GetUser(userId int64) (user *model.User, err error)
	CreateUser(user *model.User) (err error)
	EditUser(userId int64, user *model.User) (err error)
	ListUser() (users []*model.User, err error)
	DeleteUser(userId int64) (err error)
	LoginUser(username string, password string) (user *model.User, err error)
}
