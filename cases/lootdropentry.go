package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//LootDropEntryRepository handles LootDropEntryRepository cases and is a gateway to storage
type LootDropEntryRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *LootDropEntryRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *LootDropEntryRepository) Get(lootDropID int64, itemID int64) (lootDropEntry *model.LootDropEntry, err error) {

	lootDropEntry, err = c.stor.GetLootDropEntry(lootDropID, itemID)
	return
}

//Create handles logic
func (c *LootDropEntryRepository) Create(lootDropEntry *model.LootDropEntry) (err error) {
	if lootDropEntry == nil {
		err = fmt.Errorf("Empty lootDropEntry")
		return
	}
	schema, err := c.newSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootDropEntry))
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
	err = c.stor.CreateLootDropEntry(lootDropEntry)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *LootDropEntryRepository) Edit(lootDropID int64, itemID int64, lootDropEntry *model.LootDropEntry) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootDropEntry))
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

	err = c.stor.EditLootDropEntry(lootDropID, itemID, lootDropEntry)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *LootDropEntryRepository) Delete(lootDropID int64, itemID int64) (err error) {
	err = c.stor.DeleteLootDropEntry(lootDropID, itemID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *LootDropEntryRepository) List(lootDropID int64) (lootDropEntrys []*model.LootDropEntry, err error) {
	lootDropEntrys, err = c.stor.ListLootDropEntry(lootDropID)
	if err != nil {
		return
	}
	return
}

func (c *LootDropEntryRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *LootDropEntryRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "lootDropID":
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
