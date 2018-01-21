package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//TaskRepository handles TaskRepository cases and is a gateway to storage
type TaskRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *TaskRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *TaskRepository) Get(task *model.Task, user *model.User) (err error) {
	err = c.stor.GetTask(task)
	if err != nil {
		err = errors.Wrap(err, "failed to get task")
		return
	}
	err = c.prepare(task, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare task")
		return
	}
	return
}

//Create handles logic
func (c *TaskRepository) Create(task *model.Task, user *model.User) (err error) {
	if task == nil {
		err = fmt.Errorf("Empty task")
		return
	}
	schema, err := c.newSchema([]string{"title"}, nil)
	if err != nil {
		return
	}

	task.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(task))
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
	err = c.stor.CreateTask(task)
	if err != nil {
		return
	}
	err = c.prepare(task, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare task")
		return
	}
	return
}

//Edit handles logic
func (c *TaskRepository) Edit(task *model.Task, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"title"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(task))
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

	err = c.stor.EditTask(task)
	if err != nil {
		return
	}
	err = c.prepare(task, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare task")
		return
	}
	return
}

//Delete handles logic
func (c *TaskRepository) Delete(task *model.Task, user *model.User) (err error) {
	err = c.stor.DeleteTask(task)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *TaskRepository) List(user *model.User) (tasks []*model.Task, err error) {
	tasks, err = c.stor.ListTask()
	if err != nil {
		return
	}
	for _, task := range tasks {
		err = c.prepare(task, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare task")
			return
		}
	}
	return
}

func (c *TaskRepository) prepare(task *model.Task, user *model.User) (err error) {

	return
}

func (c *TaskRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *TaskRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "title":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
