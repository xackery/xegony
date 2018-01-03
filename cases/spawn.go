package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//SpawnRepository handles SpawnRepository cases and is a gateway to storage
type SpawnRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *SpawnRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *SpawnRepository) Get(spawnID int64) (spawn *model.Spawn, err error) {
	if spawnID == 0 {
		err = fmt.Errorf("Invalid Spawn ID")
		return
	}
	spawn, err = c.stor.GetSpawn(spawnID)
	return
}

//Create handles logic
func (c *SpawnRepository) Create(spawn *model.Spawn) (err error) {
	if spawn == nil {
		err = fmt.Errorf("Empty spawn")
		return
	}
	schema, err := c.newSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}
	result, err := schema.Validate(gojsonschema.NewGoLoader(spawn))
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
	err = c.stor.CreateSpawn(spawn)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *SpawnRepository) Edit(spawnID int64, spawn *model.Spawn) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spawn))
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

	err = c.stor.EditSpawn(spawnID, spawn)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *SpawnRepository) Delete(spawnID int64) (err error) {
	err = c.stor.DeleteSpawn(spawnID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *SpawnRepository) List() (spawns []*model.Spawn, err error) {
	spawns, err = c.stor.ListSpawn()
	if err != nil {
		return
	}
	return
}

//ListBySpawnGroup handles logic
func (c *SpawnRepository) ListBySpawnGroup(spawnGroupID int64) (spawns []*model.Spawn, err error) {
	spawns, err = c.stor.ListSpawnBySpawnGroup(spawnGroupID)
	if err != nil {
		return
	}
	return
}

func (c *SpawnRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *SpawnRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "shortName":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
