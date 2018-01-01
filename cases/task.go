package cases

import (
	"fmt"

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
func (c *TaskRepository) Get(taskID int64) (task *model.Task, err error) {
	if taskID == 0 {
		err = fmt.Errorf("Invalid Task ID")
		return
	}
	task, err = c.stor.GetTask(taskID)
	return
}

//Create handles logic
func (c *TaskRepository) Create(task *model.Task) (err error) {
	if task == nil {
		err = fmt.Errorf("Empty task")
		return
	}
	schema, err := c.newSchema([]string{"shortName"}, nil)
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
	return
}

//Edit handles logic
func (c *TaskRepository) Edit(taskID int64, task *model.Task) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
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

	err = c.stor.EditTask(taskID, task)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *TaskRepository) Delete(taskID int64) (err error) {
	err = c.stor.DeleteTask(taskID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *TaskRepository) List() (tasks []*model.Task, err error) {
	tasks, err = c.stor.ListTask()
	if err != nil {
		return
	}
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
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
