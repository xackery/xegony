package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type FactionRepository struct {
	stor storage.Storage
}

func (g *FactionRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *FactionRepository) Get(factionID int64) (faction *model.Faction, err error) {
	if factionID == 0 {
		err = fmt.Errorf("Invalid Faction ID")
		return
	}
	faction, err = g.stor.GetFaction(factionID)
	return
}

func (g *FactionRepository) Create(faction *model.Faction) (err error) {
	if faction == nil {
		err = fmt.Errorf("Empty faction")
		return
	}
	schema, err := g.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	faction.Id = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(faction))
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
	err = g.stor.CreateFaction(faction)
	if err != nil {
		return
	}
	return
}

func (g *FactionRepository) Edit(factionID int64, faction *model.Faction) (err error) {
	schema, err := g.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(faction))
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

	err = g.stor.EditFaction(factionID, faction)
	if err != nil {
		return
	}
	return
}

func (g *FactionRepository) Delete(factionID int64) (err error) {
	err = g.stor.DeleteFaction(factionID)
	if err != nil {
		return
	}
	return
}

func (g *FactionRepository) List() (factions []*model.Faction, err error) {
	factions, err = g.stor.ListFaction()
	if err != nil {
		return
	}
	return
}

func (c *FactionRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *FactionRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "base":
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
