package cases

import (
	"fmt"

	"github.com/pkg/errors"
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
func (c *GoalRepository) Get(goal *model.Goal, user *model.User) (err error) {

	err = c.stor.GetGoal(goal)
	if err != nil {
		err = errors.Wrap(err, "failed to get goal")
		return
	}
	err = c.prepare(goal, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare goal")
		return
	}
	return
}

//Create handles logic
func (c *GoalRepository) Create(goal *model.Goal, user *model.User) (err error) {
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
	err = c.prepare(goal, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare goal")
		return
	}
	return
}

//Edit handles logic
func (c *GoalRepository) Edit(goal *model.Goal, user *model.User) (err error) {
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

	err = c.stor.EditGoal(goal)
	if err != nil {
		return
	}
	err = c.prepare(goal, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare goal")
		return
	}
	return
}

//Delete handles logic
func (c *GoalRepository) Delete(goal *model.Goal, user *model.User) (err error) {
	err = c.stor.DeleteGoal(goal)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *GoalRepository) List(user *model.User) (goals []*model.Goal, err error) {
	goals, err = c.stor.ListGoal()
	if err != nil {
		return
	}
	for _, goal := range goals {
		err = c.prepare(goal, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare goal")
			return
		}
	}
	return
}

func (c *GoalRepository) prepare(goal *model.Goal, user *model.User) (err error) {

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
