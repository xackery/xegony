package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type LootTableRepository struct {
	stor storage.Storage
}

func (g *LootTableRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *LootTableRepository) Get(lootTableID int64) (lootTable *model.LootTable, err error) {
	if lootTableID == 0 {
		err = fmt.Errorf("Invalid LootTable ID")
		return
	}
	lootTable, err = g.stor.GetLootTable(lootTableID)
	return
}

func (g *LootTableRepository) Create(lootTable *model.LootTable) (err error) {
	if lootTable == nil {
		err = fmt.Errorf("Empty lootTable")
		return
	}
	schema, err := g.newSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}

	lootTable.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(lootTable))
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
	err = g.stor.CreateLootTable(lootTable)
	if err != nil {
		return
	}
	return
}

func (g *LootTableRepository) Edit(lootTableID int64, lootTable *model.LootTable) (err error) {
	schema, err := g.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootTable))
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

	err = g.stor.EditLootTable(lootTableID, lootTable)
	if err != nil {
		return
	}
	return
}

func (g *LootTableRepository) Delete(lootTableID int64) (err error) {
	err = g.stor.DeleteLootTable(lootTableID)
	if err != nil {
		return
	}
	return
}

func (g *LootTableRepository) List() (lootTables []*model.LootTable, err error) {
	lootTables, err = g.stor.ListLootTable()
	if err != nil {
		return
	}
	return
}

func (c *LootTableRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *LootTableRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
