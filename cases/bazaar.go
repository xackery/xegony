package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type BazaarRepository struct {
	stor storage.Storage
}

func (g *BazaarRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *BazaarRepository) Get(bazaarId int64) (bazaar *model.Bazaar, err error) {
	if bazaarId == 0 {
		err = fmt.Errorf("Invalid Bazaar ID")
		return
	}
	bazaar, err = g.stor.GetBazaar(bazaarId)
	return
}

func (g *BazaarRepository) Create(bazaar *model.Bazaar) (err error) {
	if bazaar == nil {
		err = fmt.Errorf("Empty bazaar")
		return
	}
	schema, err := bazaar.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	bazaar.Id = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(bazaar))
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
	err = g.stor.CreateBazaar(bazaar)
	if err != nil {
		return
	}
	return
}

func (g *BazaarRepository) Edit(bazaarId int64, bazaar *model.Bazaar) (err error) {
	schema, err := bazaar.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(bazaar))
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

	err = g.stor.EditBazaar(bazaarId, bazaar)
	if err != nil {
		return
	}
	return
}

func (g *BazaarRepository) Delete(bazaarId int64) (err error) {
	err = g.stor.DeleteBazaar(bazaarId)
	if err != nil {
		return
	}
	return
}

func (g *BazaarRepository) List() (bazaars []*model.Bazaar, err error) {
	bazaars, err = g.stor.ListBazaar()
	if err != nil {
		return
	}
	return
}
