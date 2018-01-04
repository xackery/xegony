package cases

import (
	"fmt"

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
func (c *MerchantRepository) Get(merchantID int64) (merchant *model.Merchant, err error) {
	if merchantID == 0 {
		err = fmt.Errorf("Invalid Merchant ID")
		return
	}
	merchant, err = c.stor.GetMerchant(merchantID)
	return
}

//Search handles logic
func (c *MerchantRepository) Search(search string) (merchants []*model.Merchant, err error) {
	merchants, err = c.stor.SearchMerchant(search)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *MerchantRepository) Delete(merchantID int64) (err error) {
	err = c.stor.DeleteMerchant(merchantID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *MerchantRepository) List(pageSize int64, pageNumber int64) (merchants []*model.Merchant, err error) {
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
	return
}

//ListCount handles logic
func (c *MerchantRepository) ListCount() (count int64, err error) {

	count, err = c.stor.ListMerchantCount()
	if err != nil {
		return
	}
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
