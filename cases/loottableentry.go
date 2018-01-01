package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//LootTableEntryRepository handles LootTableEntryRepository cases and is a gateway to storage
type LootTableEntryRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *LootTableEntryRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *LootTableEntryRepository) Get(lootTableID int64, lootDropID int64) (lootTableEntry *model.LootTableEntry, err error) {

	lootTableEntry, err = c.stor.GetLootTableEntry(lootTableID, lootDropID)
	return
}

//Create handles logic
func (c *LootTableEntryRepository) Create(lootTableEntry *model.LootTableEntry) (err error) {
	if lootTableEntry == nil {
		err = fmt.Errorf("Empty lootTableEntry")
		return
	}
	schema, err := c.newSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootTableEntry))
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
	err = c.stor.CreateLootTableEntry(lootTableEntry)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *LootTableEntryRepository) Edit(lootTableID int64, lootDropID int64, lootTableEntry *model.LootTableEntry) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootTableEntry))
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

	err = c.stor.EditLootTableEntry(lootTableID, lootDropID, lootTableEntry)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *LootTableEntryRepository) Delete(lootTableID int64, lootDropID int64) (err error) {
	err = c.stor.DeleteLootTableEntry(lootTableID, lootDropID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *LootTableEntryRepository) List(lootTableID int64) (lootTableEntrys []*model.LootTableEntry, err error) {
	lootTableEntrys, err = c.stor.ListLootTableEntry(lootTableID)
	if err != nil {
		return
	}
	return
}

func (c *LootTableEntryRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *LootTableEntryRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "lootTableID":
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
