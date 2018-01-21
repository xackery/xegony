package cases

/*
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
		err = errors.Wrap(err, "failed to get spawn entry")
		return
	}
	err = c.prepare(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawn entry")
		return
	}
	return
}

//Create handles logic
func (c *SpawnEntryRepository) Create(spawnEntry *model.SpawnEntry, user *model.User) (err error) {
	if spawnEntry == nil {
		err = fmt.Errorf("Empty SpawnEntry")
		return
	}
	if spawnEntry.SpawngroupID == 0 {
		err = fmt.Errorf("Invalid SpawnGroup ID")
		return
	}
	if spawnEntry.NpcID == 0 {
		err = fmt.Errorf("Invalid Npc ID")
		return
	}
	schema, err := c.newSchema(nil, nil)
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
	err = c.stor.CreateSpawnEntry(spawnEntry)
	if err != nil {
		return
	}
	err = c.prepare(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawn entry")
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
		err = errors.Wrap(err, "failed to prepare spawn entry")
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
func (c *SpawnEntryRepository) ListBySpawnGroup(spawnGroup *model.SpawnGroup, user *model.User) (spawnEntrys []*model.SpawnEntry, err error) {
	spawnEntrys, err = c.stor.ListSpawnEntryBySpawnGroup(spawnGroup)
	if err != nil {
		return
	}
	for _, spawnEntry := range spawnEntrys {
		err = c.prepare(spawnEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare spawn entry")
			return
		}
	}
	return
}

//ListByNpc handles logic
func (c *SpawnEntryRepository) ListByNpc(npc *model.Npc, user *model.User) (spawnEntrys []*model.SpawnEntry, err error) {
	spawnEntrys, err = c.stor.ListSpawnEntryByNpc(npc)
	if err != nil {
		return
	}
	for _, spawnEntry := range spawnEntrys {
		err = c.prepare(spawnEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare spawn entry")
			return
		}
	}
	return
}

//ListByZone handles logic
func (c *SpawnEntryRepository) ListByZone(zone *model.Zone, user *model.User) (spawnEntrys []*model.SpawnEntry, err error) {
	spawnEntrys, err = c.stor.ListSpawnEntryByZone(zone)
	if err != nil {
		return
	}
	for _, spawnEntry := range spawnEntrys {
		err = c.prepare(spawnEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare spawn entry")
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
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
*/
