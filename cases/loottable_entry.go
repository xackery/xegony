package cases

import (
	"fmt"

	"github.com/pkg/errors"
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
func (c *LootTableEntryRepository) Get(lootTableEntry *model.LootTableEntry, user *model.User) (err error) {
	err = c.stor.GetLootTableEntry(lootTableEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to get loot table entry")
		return
	}
	err = c.prepare(lootTableEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare loot table")
		return
	}
	return
}

//Create handles logic
func (c *LootTableEntryRepository) Create(lootTableEntry *model.LootTableEntry, user *model.User) (err error) {
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
	err = c.prepare(lootTableEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare loot table")
		return
	}
	return
}

//Edit handles logic
func (c *LootTableEntryRepository) Edit(lootTableEntry *model.LootTableEntry, user *model.User) (err error) {
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

	err = c.stor.EditLootTableEntry(lootTableEntry)
	if err != nil {
		return
	}
	err = c.prepare(lootTableEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare loot table")
		return
	}
	return
}

//Delete handles logic
func (c *LootTableEntryRepository) Delete(lootTableEntry *model.LootTableEntry, user *model.User) (err error) {
	err = c.stor.DeleteLootTableEntry(lootTableEntry)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *LootTableEntryRepository) ListByLootTable(lootTable *model.LootTable, user *model.User) (lootTableEntrys []*model.LootTableEntry, err error) {
	lootTableEntrys, err = c.stor.ListLootTableEntryByLootTable(lootTable)
	if err != nil {
		return
	}
	for _, lootTableEntry := range lootTableEntrys {
		err = c.prepare(lootTableEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare loot table")
			return
		}
	}
	return
}

func (c *LootTableEntryRepository) prepare(lootTableEntry *model.LootTableEntry, user *model.User) (err error) {

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
