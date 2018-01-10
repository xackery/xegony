package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//LootDropRepository handles LootDropRepository cases and is a gateway to storage
type LootDropRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *LootDropRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *LootDropRepository) Get(lootDropID int64) (lootDrop *model.LootDrop, err error) {
	if lootDropID == 0 {
		err = fmt.Errorf("Invalid LootDrop ID")
		return
	}
	lootDrop, err = c.stor.GetLootDrop(lootDropID)
	return
}

//Create handles logic
func (c *LootDropRepository) Create(lootDrop *model.LootDrop) (err error) {
	if lootDrop == nil {
		err = fmt.Errorf("Empty lootDrop")
		return
	}
	schema, err := c.newSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}

	lootDrop.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(lootDrop))
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
	err = c.stor.CreateLootDrop(lootDrop)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *LootDropRepository) Edit(lootDropID int64, lootDrop *model.LootDrop) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootDrop))
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

	err = c.stor.EditLootDrop(lootDropID, lootDrop)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *LootDropRepository) Delete(lootDropID int64) (err error) {
	err = c.stor.DeleteLootDrop(lootDropID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *LootDropRepository) List() (lootDrops []*model.LootDrop, err error) {
	lootDrops, err = c.stor.ListLootDrop()
	if err != nil {
		return
	}
	return
}

func (c *LootDropRepository) prepare(lootDrop *model.LootDrop) (err error) {

	return
}

func (c *LootDropRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *LootDropRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
