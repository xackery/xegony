package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//SpawnEntryRepository handles SpawnEntryRepository cases and is a gateway to storage
type SpawnEntryRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *SpawnEntryRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *SpawnEntryRepository) Get(spawnEntry *model.SpawnEntry, user *model.User) (err error) {
	err = c.stor.GetSpawnEntry(spawnEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawnEntry")
		return
	}

	err = c.prepare(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnEntry")
		return
	}

	return
}

//Create handles logic
func (c *SpawnEntryRepository) Create(spawnEntry *model.SpawnEntry, user *model.User) (err error) {
	if spawnEntry == nil {
		err = fmt.Errorf("Empty spawnEntry")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	spawnEntry.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(spawnEntry))
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
	err = c.stor.CreateSpawnEntry(spawnEntry)
	if err != nil {
		return
	}
	err = c.prepare(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnEntry")
		return
	}
	return
}

//Edit handles logic
func (c *SpawnEntryRepository) Edit(spawnEntry *model.SpawnEntry, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spawnEntry))
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

	err = c.stor.EditSpawnEntry(spawnEntry)
	if err != nil {
		return
	}
	err = c.prepare(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnEntry")
		return
	}
	return
}

//Delete handles logic
func (c *SpawnEntryRepository) Delete(spawnEntry *model.SpawnEntry, user *model.User) (err error) {
	err = c.stor.DeleteSpawnEntry(spawnEntry)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *SpawnEntryRepository) List(user *model.User) (spawnEntrys []*model.SpawnEntry, err error) {
	spawnEntrys, err = c.stor.ListSpawnEntry()
	if err != nil {
		return
	}
	for _, spawnEntry := range spawnEntrys {
		err = c.prepare(spawnEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare spawnEntry")
			return
		}
	}
	return
}

//List handles logic
func (c *SpawnEntryRepository) ListBySpawn(spawn *model.Spawn, user *model.User) (spawnEntrys []*model.SpawnEntry, err error) {
	spawnEntrys, err = c.stor.ListSpawnEntryBySpawn(spawn)
	if err != nil {
		return
	}
	for _, spawnEntry := range spawnEntrys {
		err = c.prepare(spawnEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare spawnEntry")
			return
		}
	}
	return
}

func (c *SpawnEntryRepository) prepare(spawnEntry *model.SpawnEntry, user *model.User) (err error) {

	return
}

func (c *SpawnEntryRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *SpawnEntryRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
