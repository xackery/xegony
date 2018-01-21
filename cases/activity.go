package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//ActivityRepository handles ActivityRepository cases and is a gateway to storage
type ActivityRepository struct {
	stor       storage.Storage
	globalRepo *GlobalRepository
}

//Initialize handles logic
func (c *ActivityRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	c.globalRepo = &GlobalRepository{}
	err = c.globalRepo.Initialize(stor)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize global repository")
		return
	}

	return
}

//Get handles logic
func (c *ActivityRepository) Get(activity *model.Activity, user *model.User) (err error) {

	err = c.stor.GetActivity(activity)
	if err != nil {
		err = errors.Wrap(err, "failed to get activity")
	}
	err = c.prepare(activity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare activity")
	}
	return
}

//Create handles logic
func (c *ActivityRepository) Create(activity *model.Activity, user *model.User) (err error) {
	if activity == nil {
		err = fmt.Errorf("Empty activity")
		return
	}

	schema, err := c.newSchema([]string{"zoneid", "activityType"}, nil)
	if err != nil {
		return
	}

	//Check if step is valid
	activity.Step, err = c.stor.GetActivityNextStep(activity)
	if err != nil {
		err = errors.Wrap(err, "Failed to get next activity step")
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
	err = c.stor.CreateActivity(activity)
	if err != nil {
		return
	}
	err = c.prepare(activity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare activity")
	}
	return
}

//Edit handles logic
func (c *ActivityRepository) Edit(activity *model.Activity, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"body"}, nil)
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

	err = c.stor.EditActivity(activity)
	if err != nil {
		return
	}
	err = c.prepare(activity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare activity")
	}
	return
}

//Delete handles logic
func (c *ActivityRepository) Delete(activity *model.Activity, user *model.User) (err error) {

	err = c.stor.DeleteActivity(activity)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *ActivityRepository) ListByTask(task *model.Task, user *model.User) (activitys []*model.Activity, err error) {
	activitys, err = c.stor.ListActivityByTask(task)
	if err != nil {
		return
	}
	for _, activity := range activitys {
		err = c.prepare(activity, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare activity")
		}
	}
	return
}

func (c *ActivityRepository) prepare(activity *model.Activity, user *model.User) (err error) {
	activity.Zone, err = c.globalRepo.GetZone(activity.ZoneID, user)
	if err != nil {
		err = errors.Wrap(err, "Failed to get activity zone ID")
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
	case "zoneid":
		prop.Type = "integer"
		prop.Minimum = 1
	case "activityType":
		prop.Type = "integer"
		prop.EnumInt = []int64{1, 2, 3, 4, 5, 6, 7, 8, 11, 100}
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
