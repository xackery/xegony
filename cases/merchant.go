package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//MerchantRepository handles MerchantRepository cases and is a gateway to storage
type MerchantRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *MerchantRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *MerchantRepository) Get(merchant *model.Merchant, user *model.User) (err error) {

	err = c.stor.GetMerchant(merchant)
	if err != nil {
		err = errors.Wrap(err, "failed to get merchant")
		return
	}
	err = c.prepare(merchant, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare merchant")
		return
	}
	return
}

//Delete handles logic
func (c *MerchantRepository) Delete(merchant *model.Merchant, user *model.User) (err error) {
	err = c.stor.DeleteMerchant(merchant)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *MerchantRepository) List(pageSize int64, pageNumber int64, user *model.User) (merchants []*model.Merchant, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	merchants, err = c.stor.ListMerchant(pageSize, pageNumber)
	if err != nil {
		return
	}
	for _, merchant := range merchants {
		err = c.prepare(merchant, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare merchant")
			return
		}
	}
	return
}

//ListCount handles logic
func (c *MerchantRepository) ListCount(user *model.User) (count int64, err error) {

	count, err = c.stor.ListMerchantCount()
	if err != nil {
		return
	}
	return
}

func (c *MerchantRepository) prepare(merchant *model.Merchant, user *model.User) (err error) {

	return
}

func (c *MerchantRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *MerchantRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "zoneID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
