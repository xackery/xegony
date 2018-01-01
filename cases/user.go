package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type UserRepository struct {
	stor storage.Storage
}

func (g *UserRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *UserRepository) Get(userID int64) (user *model.User, err error) {
	if userID == 0 {
		err = fmt.Errorf("Invalid User ID")
		return
	}
	user, err = g.stor.GetUser(userID)
	return
}

func (g *UserRepository) Create(user *model.User) (err error) {
	if user == nil {
		err = fmt.Errorf("Empty user")
		return
	}
	schema, err := user.NewSchema([]string{"name", "password", "email", "accountID"}, nil)
	if err != nil {
		return
	}
	user.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(user))
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
	err = g.stor.CreateUser(user)
	if err != nil {
		return
	}
	return
}

func (g *UserRepository) Login(username string, password string) (user *model.User, err error) {
	user = &model.User{
		Name:     username,
		Password: password,
	}
	schema, err := user.NewSchema([]string{"name", "password"}, nil)
	if err != nil {
		return
	}
	result, err := schema.Validate(gojsonschema.NewGoLoader(user))
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
	user, err = g.stor.LoginUser(user.Name, user.Password)
	if err != nil {
		return
	}

	return
}

func (g *UserRepository) Edit(userID int64, user *model.User) (err error) {
	schema, err := user.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(user))
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

	err = g.stor.EditUser(userID, user)
	if err != nil {
		return
	}
	return
}

func (g *UserRepository) Delete(userID int64) (err error) {
	err = g.stor.DeleteUser(userID)
	if err != nil {
		return
	}
	return
}

func (g *UserRepository) List() (users []*model.User, err error) {
	users, err = g.stor.ListUser()
	if err != nil {
		return
	}
	return
}
