package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type SpawnRepository struct {
	stor storage.Storage
}

func (g *SpawnRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *SpawnRepository) Get(spawnID int64) (spawn *model.Spawn, err error) {
	if spawnID == 0 {
		err = fmt.Errorf("Invalid Spawn ID")
		return
	}
	spawn, err = g.stor.GetSpawn(spawnID)
	return
}

func (g *SpawnRepository) Create(spawn *model.Spawn) (err error) {
	if spawn == nil {
		err = fmt.Errorf("Empty spawn")
		return
	}
	schema, err := g.newSchema([]string{"shortName"}, nil)
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
	err = g.stor.CreateSpawn(spawn)
	if err != nil {
		return
	}
	return
}

func (g *SpawnRepository) Edit(spawnID int64, spawn *model.Spawn) (err error) {
	schema, err := g.newSchema([]string{"name"}, nil)
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

	err = g.stor.EditSpawn(spawnID, spawn)
	if err != nil {
		return
	}
	return
}

func (g *SpawnRepository) Delete(spawnID int64) (err error) {
	err = g.stor.DeleteSpawn(spawnID)
	if err != nil {
		return
	}
	return
}

func (g *SpawnRepository) List() (spawns []*model.Spawn, err error) {
	spawns, err = g.stor.ListSpawn()
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
