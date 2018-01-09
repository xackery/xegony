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
	stor storage.Storage
}

//Initialize handles logic
func (c *ActivityRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *ActivityRepository) Get(taskID int64, activityID int64) (activity *model.Activity, err error) {
	if activityID == 0 {
		err = fmt.Errorf("Invalid Activity ID")
		return
	}
	if taskID == 0 {
		err = fmt.Errorf("Invalid Task ID")
		return
	}
	activity, err = c.stor.GetActivity(taskID, activityID)
	return
}

//Create handles logic
func (c *ActivityRepository) Create(activity *model.Activity) (err error) {
	if activity == nil {
		err = fmt.Errorf("Empty activity")
		return
	}

	schema, err := c.newSchema([]string{"zoneid", "activityType"}, nil)
	if err != nil {
		return
	}

	//Verify taskID

	_, err = c.stor.GetTask(activity.TaskID)
	if err != nil {
		err = errors.Wrap(err, "Failed to verify TaskID")
		return
	}

	//Verify zoneID
	if activity.ZoneID == 0 {
		err = fmt.Errorf("Invalid ZoneID")
	}
	zoneRepo := &ZoneRepository{}
	err = zoneRepo.Initialize(c.stor)
	if err != nil {
		err = errors.Wrap(err, "Failed to initialize zone repo")
		return
	}

	if _, err = zoneRepo.Get(activity.ZoneID); err != nil {
		err = errors.Wrap(err, "Failed to get zone ID")
		return
	}

	//Check if step is valid
	step, err := c.stor.GetActivityNextStep(activity.TaskID, activity.ActivityID)
	if err != nil {
		err = errors.Wrap(err, "Failed to get next activity step")
		return
	}
	if activity.Step == 0 {
		activity.Step = step
	}
	if activity.Step > step {
		err = errors.Wrap(err, "Step is out of steps bounds")
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
	return
}

//Edit handles logic
func (c *ActivityRepository) Edit(activityID int64, activity *model.Activity) (err error) {
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

	err = c.stor.EditActivity(activityID, activity)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *ActivityRepository) Delete(activityID int64) (err error) {
	err = c.stor.DeleteActivity(activityID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *ActivityRepository) List(taskID int64) (activitys []*model.Activity, err error) {
	activitys, err = c.stor.ListActivity(taskID)
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
