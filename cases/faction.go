package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//FactionRepository handles FactionRepository cases and is a gateway to storage
type FactionRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *FactionRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *FactionRepository) Get(faction *model.Faction, user *model.User) (err error) {

	err = c.stor.GetFaction(faction)
	if err != nil {
		err = errors.Wrap(err, "failed to get faction")
		return
	}
	err = c.prepare(faction, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare faction")
		return
	}
	return
}

//Create handles logic
func (c *FactionRepository) Create(faction *model.Faction, user *model.User) (err error) {
	if faction == nil {
		err = fmt.Errorf("Empty faction")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	faction.ID = 0 //strip ID
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
	err = c.stor.CreateFaction(faction)
	if err != nil {
		return
	}
	err = c.prepare(faction, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare faction")
		return
	}
	return
}

//Edit handles logic
func (c *FactionRepository) Edit(faction *model.Faction, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
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

	err = c.stor.EditFaction(faction)
	if err != nil {
		return
	}
	err = c.prepare(faction, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare faction")
		return
	}
	return
}

//Delete handles logic
func (c *FactionRepository) Delete(faction *model.Faction, user *model.User) (err error) {
	err = c.stor.DeleteFaction(faction)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *FactionRepository) List(user *model.User) (factions []*model.Faction, err error) {
	factions, err = c.stor.ListFaction()
	if err != nil {
		return
	}
	for _, faction := range factions {
		err = c.prepare(faction, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare faction")
			return
		}
	}
	return
}

func (c *FactionRepository) prepare(faction *model.Faction, user *model.User) (err error) {
	faction.CleanName = cleanName(faction.Name)
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
