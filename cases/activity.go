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

func (g *ActivityRepository) Get(taskId int64, activityId int64) (activity *model.Activity, err error) {
	if activityId == 0 {
		err = fmt.Errorf("Invalid Activity ID")
		return
	}
	if taskId == 0 {
		err = fmt.Errorf("Invalid Task ID")
		return
	}
	activity, err = g.stor.GetActivity(taskId, activityId)
	return
}

func (g *ActivityRepository) Create(activity *model.Activity) (err error) {
	if activity == nil {
		err = fmt.Errorf("Empty activity")
		return
	}
	schema, err := activity.NewSchema([]string{"body"}, nil)
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

func (g *ActivityRepository) Edit(activityId int64, activity *model.Activity) (err error) {
	schema, err := activity.NewSchema([]string{"body"}, nil)
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

	err = g.stor.EditActivity(activityId, activity)
	if err != nil {
		return
	}
	return
}

func (g *ActivityRepository) Delete(activityId int64) (err error) {
	err = g.stor.DeleteActivity(activityId)
	if err != nil {
		return
	}
	return
}

func (g *ActivityRepository) List(taskId int64) (activitys []*model.Activity, err error) {
	activitys, err = g.stor.ListActivity(taskId)
	if err != nil {
		return
	}
	return
}
