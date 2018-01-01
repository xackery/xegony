package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type SpawnEntryRepository struct {
	stor storage.Storage
}

func (g *SpawnEntryRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *SpawnEntryRepository) Get(spawnGroupID int64, npcID int64) (spawnEntry *model.SpawnEntry, err error) {
	if spawnGroupID == 0 {
		err = fmt.Errorf("Invalid SpawnEntry ID")
		return
	}
	if npcID == 0 {
		err = fmt.Errorf("Invalid Npc ID")
		return
	}
	spawnEntry, err = g.stor.GetSpawnEntry(spawnGroupID, npcID)
	return
}

func (g *SpawnEntryRepository) Create(spawnEntry *model.SpawnEntry) (err error) {
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
	schema, err := g.newSchema(nil, nil)
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
	err = g.stor.CreateSpawnEntry(spawnEntry)
	if err != nil {
		return
	}
	return
}

func (g *SpawnEntryRepository) Edit(spawnGroupID int64, npcID int64, spawnEntry *model.SpawnEntry) (err error) {
	schema, err := g.newSchema([]string{"name"}, nil)
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

	err = g.stor.EditSpawnEntry(spawnGroupID, npcID, spawnEntry)
	if err != nil {
		return
	}
	return
}

func (g *SpawnEntryRepository) Delete(spawnGroupID int64, npcID int64) (err error) {
	err = g.stor.DeleteSpawnEntry(spawnGroupID, npcID)
	if err != nil {
		return
	}
	return
}

func (g *SpawnEntryRepository) List(spawnGroupID int64) (spawnEntrys []*model.SpawnEntry, err error) {
	spawnEntrys, err = g.stor.ListSpawnEntry(spawnGroupID)
	if err != nil {
		return
	}
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
