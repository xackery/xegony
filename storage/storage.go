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
	//Character
	GetCharacter(characterId int64) (character *model.Character, err error)
	CreateCharacter(character *model.Character) (err error)
	EditCharacter(characterId int64, character *model.Character) (err error)
	ListCharacter() (characters []*model.Character, err error)
	DeleteCharacter(characterId int64) (err error)
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
	//User
	GetUser(userId int64) (user *model.User, err error)
	CreateUser(user *model.User) (err error)
	EditUser(userId int64, user *model.User) (err error)
	ListUser() (users []*model.User, err error)
	DeleteUser(userId int64) (err error)
	LoginUser(username string, password string) (user *model.User, err error)
}
