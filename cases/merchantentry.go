package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//MerchantEntryRepository handles MerchantEntryRepository cases and is a gateway to storage
type MerchantEntryRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *MerchantEntryRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *MerchantEntryRepository) Get(merchantID int64, itemID int64) (merchantEntry *model.MerchantEntry, query string, err error) {
	if merchantID == 0 {
		err = fmt.Errorf("Invalid MerchantEntry ID")
		return
	}
	if itemID == 0 {
		err = fmt.Errorf("Invalid Item ID")
		return
	}
	query, merchantEntry, err = c.stor.GetMerchantEntry(merchantID, itemID)
	return
}

//Create handles logic
func (c *MerchantEntryRepository) Create(merchantEntry *model.MerchantEntry) (query string, err error) {
	if merchantEntry == nil {
		err = fmt.Errorf("Empty MerchantEntry")
		return
	}
	if merchantEntry.MerchantID == 0 {
		err = fmt.Errorf("Invalid MerchantGroup ID")
		return
	}
	if merchantEntry.ItemID == 0 {
		err = fmt.Errorf("Invalid Item ID")
		return
	}
	schema, err := c.newSchema(nil, nil)
	if err != nil {
		return
	}
	result, err := schema.Validate(gojsonschema.NewGoLoader(merchantEntry))
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
	query, err = c.stor.CreateMerchantEntry(merchantEntry)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *MerchantEntryRepository) Edit(merchantID int64, itemID int64, merchantEntry *model.MerchantEntry) (query string, err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(merchantEntry))
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

	query, err = c.stor.EditMerchantEntry(merchantID, itemID, merchantEntry)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *MerchantEntryRepository) Delete(merchantID int64, itemID int64) (query string, err error) {
	query, err = c.stor.DeleteMerchantEntry(merchantID, itemID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *MerchantEntryRepository) List(merchantID int64) (merchantEntrys []*model.MerchantEntry, query string, err error) {
	query, merchantEntrys, err = c.stor.ListMerchantEntry(merchantID)
	if err != nil {
		return
	}
	return
}

//ListByItem handles logic
func (c *MerchantEntryRepository) ListByItem(itemID int64) (merchantEntrys []*model.MerchantEntry, query string, err error) {
	query, merchantEntrys, err = c.stor.ListMerchantEntryByItem(itemID)
	if err != nil {
		return
	}
	return
}

func (c *MerchantEntryRepository) prepare(merchantEntry *model.MerchantEntry) (err error) {

	return
}

func (c *MerchantEntryRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *MerchantEntryRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
