package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//CharacterGraphRepository handles CharacterGraphRepository cases and is a gateway to storage
type CharacterGraphRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *CharacterGraphRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *CharacterGraphRepository) Get(id int64) (characterGraph *model.CharacterGraph, err error) {
	if id == 0 {
		err = fmt.Errorf("Invalid CharacterGraph ID")
		return
	}
	characterGraph, err = c.stor.GetCharacterGraph(id)
	return
}

//Create handles logic
func (c *CharacterGraphRepository) Create(characterGraph *model.CharacterGraph) (err error) {
	if characterGraph == nil {
		err = fmt.Errorf("Empty characterGraph")
		return
	}
	schema, err := c.newSchema([]string{"name", "accountID"}, nil)
	if err != nil {
		return
	}

	characterGraph.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(characterGraph))
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
	err = c.stor.CreateCharacterGraph(characterGraph)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *CharacterGraphRepository) Edit(id int64, characterGraph *model.CharacterGraph) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(characterGraph))
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

	err = c.stor.EditCharacterGraph(id, characterGraph)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *CharacterGraphRepository) Delete(characterID int64) (err error) {
	err = c.stor.DeleteCharacterGraph(characterID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *CharacterGraphRepository) List(characterID int64) (characterGraphs []*model.CharacterGraph, err error) {
	characterGraphs, err = c.stor.ListCharacterGraph(characterID)
	if err != nil {
		return
	}
	return
}

func (c *CharacterGraphRepository) prepare(characterGraph *model.CharacterGraph) (err error) {

	return
}

func (c *CharacterGraphRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *CharacterGraphRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
