package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type ActivityRepository struct {
	stor storage.Storage
}

func (g *ActivityRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *ActivityRepository) Get(taskID int64, activityID int64) (activity *model.Activity, err error) {
	if activityID == 0 {
		err = fmt.Errorf("Invalid Activity ID")
		return
	}
	if taskID == 0 {
		err = fmt.Errorf("Invalid Task ID")
		return
	}
	activity, err = g.stor.GetActivity(taskID, activityID)
	return
}

func (g *ActivityRepository) Create(activity *model.Activity) (err error) {
	if activity == nil {
		err = fmt.Errorf("Empty activity")
		return
	}
	schema, err := g.newSchema([]string{"body"}, nil)
	if err != nil {
		return
	}
	result, err := schema.Validate(gojsonschema.NewGoLoader(activity))
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
	err = g.stor.CreateActivity(activity)
	if err != nil {
		return
	}
	return
}

func (g *ActivityRepository) Edit(activityID int64, activity *model.Activity) (err error) {
	schema, err := g.newSchema([]string{"body"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(activity))
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

	err = g.stor.EditActivity(activityID, activity)
	if err != nil {
		return
	}
	return
}

func (g *ActivityRepository) Delete(activityID int64) (err error) {
	err = g.stor.DeleteActivity(activityID)
	if err != nil {
		return
	}
	return
}

func (g *ActivityRepository) List(taskID int64) (activitys []*model.Activity, err error) {
	activitys, err = g.stor.ListActivity(taskID)
	if err != nil {
		return
	}
	return
}

func (c *ActivityRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *ActivityRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
