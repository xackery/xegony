package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type AccountRepository struct {
	stor storage.Storage
}

func (g *AccountRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *AccountRepository) Get(accountId int64) (account *model.Account, err error) {
	if accountId == 0 {
		err = fmt.Errorf("Invalid Account ID")
		return
	}
	account, err = g.stor.GetAccount(accountId)
	return
}

func (g *AccountRepository) Create(account *model.Account) (err error) {
	if account == nil {
		err = fmt.Errorf("Empty account")
		return
	}
	schema, err := account.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	account.Id = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(account))
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
	err = g.stor.CreateAccount(account)
	if err != nil {
		return
	}
	return
}

func (g *AccountRepository) Edit(accountId int64, account *model.Account) (err error) {
	schema, err := account.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(account))
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

	err = g.stor.EditAccount(accountId, account)
	if err != nil {
		return
	}
	return
}

func (g *AccountRepository) Delete(accountId int64) (err error) {
	err = g.stor.DeleteAccount(accountId)
	if err != nil {
		return
	}
	return
}

func (g *AccountRepository) List() (accounts []*model.Account, err error) {
	accounts, err = g.stor.ListAccount()
	if err != nil {
		return
	}
	return
}
