package cases

import (
	"fmt"

	"github.com/pkg/errors"
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
func (c *MerchantEntryRepository) Get(merchantEntry *model.MerchantEntry, user *model.User) (err error) {

	err = c.stor.GetMerchantEntry(merchantEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to get merchant entry")
		return
	}
	err = c.prepare(merchantEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare merchant entry")
		return
	}
	return
}

//Create handles logic
func (c *MerchantEntryRepository) Create(merchantEntry *model.MerchantEntry, user *model.User) (err error) {
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
	err = c.stor.CreateMerchantEntry(merchantEntry)
	if err != nil {
		return
	}
	err = c.prepare(merchantEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare merchant entry")
		return
	}
	return
}

//Edit handles logic
func (c *MerchantEntryRepository) Edit(merchantEntry *model.MerchantEntry, user *model.User) (err error) {
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

	err = c.stor.EditMerchantEntry(merchantEntry)
	if err != nil {
		return
	}
	err = c.prepare(merchantEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare merchant entry")
		return
	}
	return
}

//Delete handles logic
func (c *MerchantEntryRepository) Delete(merchantEntry *model.MerchantEntry, user *model.User) (err error) {
	err = c.stor.DeleteMerchantEntry(merchantEntry)
	if err != nil {
		return
	}
	return
}

//ListByMerchant handles logic
func (c *MerchantEntryRepository) ListByMerchant(merchant *model.Merchant, user *model.User) (merchantEntrys []*model.MerchantEntry, err error) {
	merchantEntrys, err = c.stor.ListMerchantEntryByMerchant(merchant)
	if err != nil {
		return
	}
	for _, merchantEntry := range merchantEntrys {
		err = c.prepare(merchantEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare merchant entry")
			return
		}
	}
	return
}

//ListByItem handles logic
func (c *MerchantEntryRepository) ListByItem(item *model.Item, user *model.User) (merchantEntrys []*model.MerchantEntry, err error) {
	merchantEntrys, err = c.stor.ListMerchantEntryByItem(item)
	if err != nil {
		return
	}
	for _, merchantEntry := range merchantEntrys {
		err = c.prepare(merchantEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare merchant entry")
			return
		}
	}
	return
}

func (c *MerchantEntryRepository) prepare(merchantEntry *model.MerchantEntry, user *model.User) (err error) {

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
