package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

var ()

//VariableRepository handles VariableRepository cases and is a gateway to storage
type VariableRepository struct {
	stor                  storage.Storage
	variableCache         map[string]*model.Variable
	isVariableCacheLoaded bool
}

//Initialize handler
func (c *VariableRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}

	c.stor = stor
	c.isVariableCacheLoaded = false
	if err = c.rebuildCache(); err != nil {
		return
	}
	return
}

func (c *VariableRepository) rebuildCache() (err error) {
	if c.isVariableCacheLoaded {
		return
	}
	c.isVariableCacheLoaded = true
	c.variableCache = make(map[string]*model.Variable)
	variables, err := c.list()
	if err != nil {
		return
	}

	for _, variable := range variables {
		c.variableCache[variable.Name] = variable
	}
	fmt.Println("Rebuilt Variable Cache")
	return
}

//Get handler
func (c *VariableRepository) Get(variableName string) (variable *model.Variable, err error) {
	variable = c.variableCache[variableName]
	return
}

//Create handler
func (c *VariableRepository) Create(variable *model.Variable) (err error) {
	if variable == nil {
		err = fmt.Errorf("Empty variable")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(variable))
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
	err = c.stor.CreateVariable(variable)
	if err != nil {
		return
	}
	c.isVariableCacheLoaded = false
	c.rebuildCache()
	return
}

//Edit handler
func (c *VariableRepository) Edit(variableName string, variable *model.Variable) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(variable))
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

	if err = c.stor.EditVariable(variableName, variable); err != nil {
		return
	}
	if err = c.rebuildCache(); err != nil {
		return
	}
	return
}

//Delete handler
func (c *VariableRepository) Delete(variableName string) (err error) {
	err = c.stor.DeleteVariable(variableName)
	if err != nil {
		return
	}
	if err = c.rebuildCache(); err != nil {
		return
	}
	return
}

func (c *VariableRepository) list() (variables []*model.Variable, err error) {
	if variables, err = c.stor.ListVariable(); err != nil {
		return
	}
	return
}

//List handler
func (c *VariableRepository) List() (variables []*model.Variable, err error) {
	for _, variable := range c.variableCache {
		variables = append(variables, variable)
	}
	return
}

func (c *VariableRepository) prepare(variable *model.Variable) (err error) {

	return
}

//newSchema handler
func (c *VariableRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

//getSchemaProperty handler
func (c *VariableRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
