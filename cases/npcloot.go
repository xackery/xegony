package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//NpcLootRepository handles NpcLootRepository cases and is a gateway to storage
type NpcLootRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *NpcLootRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *NpcLootRepository) Get(npcID int64, itemID int64) (npcLoot *model.NpcLoot, err error) {
	npcLoot, err = c.stor.GetNpcLoot(npcID, itemID)
	return
}

//Truncate handles logic
func (c *NpcLootRepository) Truncate() (err error) {
	err = c.stor.TruncateNpcLoot()
	return
}

//Create handles logic
func (c *NpcLootRepository) Create(npcLoot *model.NpcLoot) (err error) {
	if npcLoot == nil {
		err = fmt.Errorf("Empty npcLoot")
		return
	}
	schema, err := c.newSchema([]string{}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(npcLoot))
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
	err = c.stor.CreateNpcLoot(npcLoot)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *NpcLootRepository) Edit(npcID int64, itemID int64, npcLoot *model.NpcLoot) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(npcLoot))
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

	err = c.stor.EditNpcLoot(npcID, itemID, npcLoot)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *NpcLootRepository) Delete(npcID int64, itemID int64) (err error) {
	err = c.stor.DeleteNpcLoot(npcID, itemID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *NpcLootRepository) List(npcID int64) (npcLoots []*model.NpcLoot, err error) {
	npcLoots, err = c.stor.ListNpcLoot(npcID)
	if err != nil {
		return
	}
	return
}

//ListByZone handles logic
func (c *NpcLootRepository) ListByZone(zoneID int64) (npcLoots []*model.NpcLoot, err error) {
	npcLoots, err = c.stor.ListNpcLootByZone(zoneID)
	if err != nil {
		return
	}
	return
}

func (c *NpcLootRepository) prepare(npcLoot *model.NpcLoot) (err error) {

	return
}

func (c *NpcLootRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *NpcLootRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "accountID":
		prop.Type = "integer"
		prop.Minimum = 1
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
