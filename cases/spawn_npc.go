package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//SpawnNpcRepository handles SpawnNpcRepository cases and is a gateway to storage
type SpawnNpcRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *SpawnNpcRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *SpawnNpcRepository) Get(spawnNpc *model.SpawnNpc, user *model.User) (err error) {
	err = c.stor.GetSpawnNpc(spawnNpc)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawnNpc")
		return
	}

	err = c.prepare(spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnNpc")
		return
	}

	return
}

//Create handles logic
func (c *SpawnNpcRepository) Create(spawnNpc *model.SpawnNpc, user *model.User) (err error) {
	if spawnNpc == nil {
		err = fmt.Errorf("Empty spawnNpc")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	result, err := schema.Validate(gojsonschema.NewGoLoader(spawnNpc))
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
	err = c.stor.CreateSpawnNpc(spawnNpc)
	if err != nil {
		return
	}
	err = c.prepare(spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnNpc")
		return
	}
	return
}

//Edit handles logic
func (c *SpawnNpcRepository) Edit(spawnNpc *model.SpawnNpc, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spawnNpc))
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

	err = c.stor.EditSpawnNpc(spawnNpc)
	if err != nil {
		return
	}
	err = c.prepare(spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnNpc")
		return
	}
	return
}

//Delete handles logic
func (c *SpawnNpcRepository) Delete(spawnNpc *model.SpawnNpc, user *model.User) (err error) {
	err = c.stor.DeleteSpawnNpc(spawnNpc)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *SpawnNpcRepository) List(user *model.User) (spawnNpcs []*model.SpawnNpc, err error) {
	spawnNpcs, err = c.stor.ListSpawnNpc()
	if err != nil {
		return
	}
	for _, spawnNpc := range spawnNpcs {
		err = c.prepare(spawnNpc, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare spawnNpc")
			return
		}
	}
	return
}

//List handles logic
func (c *SpawnNpcRepository) ListBySpawn(spawn *model.Spawn, user *model.User) (spawnNpcs []*model.SpawnNpc, err error) {
	spawnNpcs, err = c.stor.ListSpawnNpcBySpawn(spawn)
	if err != nil {
		return
	}
	for _, spawnNpc := range spawnNpcs {
		err = c.prepare(spawnNpc, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare spawnNpc")
			return
		}
	}
	return
}

func (c *SpawnNpcRepository) prepare(spawnNpc *model.SpawnNpc, user *model.User) (err error) {

	return
}

func (c *SpawnNpcRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *SpawnNpcRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "status":
		prop.Type = "integer"
		prop.Minimum = 1
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
