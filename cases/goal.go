package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//GoalRepository handles GoalRepository cases and is a gateway to storage
type GoalRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *GoalRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *GoalRepository) Get(listID int64, entryID int64) (goal *model.Goal, err error) {
	if listID == 0 {
		err = fmt.Errorf("Invalid List ID")
		return
	}
	if entryID == 0 {
		err = fmt.Errorf("Invalid Entry ID")
		return
	}
	goal, err = c.stor.GetGoal(listID, entryID)
	return
}

//Create handles logic
func (c *GoalRepository) Create(goal *model.Goal) (err error) {
	if goal == nil {
		err = fmt.Errorf("Empty goal")
		return
	}
	schema, err := c.newSchema([]string{"body"}, nil)
	if err != nil {
		return
	}
	result, err := schema.Validate(gojsonschema.NewGoLoader(goal))
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
	err = c.stor.CreateGoal(goal)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *GoalRepository) Edit(listID int64, goal *model.Goal) (err error) {
	schema, err := c.newSchema([]string{"body"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(goal))
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

	err = c.stor.EditGoal(listID, goal)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *GoalRepository) Delete(listID int64, entryID int64) (err error) {
	err = c.stor.DeleteGoal(listID, entryID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *GoalRepository) List() (goals []*model.Goal, err error) {
	goals, err = c.stor.ListGoal()
	if err != nil {
		return
	}
	return
}

func (c *GoalRepository) prepare(goal *model.Goal) (err error) {

	return
}

func (c *GoalRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *GoalRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
