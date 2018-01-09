package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//BazaarRepository handles BazaarRepository cases and is a gateway to storage
type BazaarRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *BazaarRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *BazaarRepository) Get(bazaarID int64) (bazaar *model.Bazaar, err error) {
	if bazaarID == 0 {
		err = fmt.Errorf("Invalid Bazaar ID")
		return
	}
	bazaar, err = c.stor.GetBazaar(bazaarID)
	return
}

//Create handles logic
func (c *BazaarRepository) Create(bazaar *model.Bazaar) (err error) {
	if bazaar == nil {
		err = fmt.Errorf("Empty bazaar")
		return
	}
	schema, err := c.newSchema([]string{"price", "accountID", "itemID"}, nil)
	if err != nil {
		return
	}
	bazaar.ID = 0 //strip ID
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
	err = c.stor.CreateBazaar(bazaar)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *BazaarRepository) Edit(bazaarID int64, bazaar *model.Bazaar) (err error) {
	schema, err := c.newSchema([]string{"price", "accountID", "itemID"}, nil)
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

	err = c.stor.EditBazaar(bazaarID, bazaar)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *BazaarRepository) Delete(bazaarID int64) (err error) {
	err = c.stor.DeleteBazaar(bazaarID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *BazaarRepository) List() (bazaars []*model.Bazaar, err error) {
	bazaars, err = c.stor.ListBazaar()
	if err != nil {
		return
	}
	return
}

func (c *BazaarRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	jsRef := gojsonschema.NewGoLoader(s)
	schema, err = gojsonschema.NewSchema(jsRef)
	if err != nil {
		return
	}
	return
}

func (c *BazaarRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "itemID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "accountID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "price":
		prop.Type = "integer"
		prop.Minimum = 1
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
