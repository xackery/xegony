package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type GoalRepository struct {
	stor storage.Storage
}

func (g *GoalRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *GoalRepository) Get(listID int64, entryID int64) (goal *model.Goal, err error) {
	if listID == 0 {
		err = fmt.Errorf("Invalid List ID")
		return
	}
	if entryID == 0 {
		err = fmt.Errorf("Invalid Entry ID")
		return
	}
	goal, err = g.stor.GetGoal(listID, entryID)
	return
}

func (g *GoalRepository) Create(goal *model.Goal) (err error) {
	if goal == nil {
		err = fmt.Errorf("Empty goal")
		return
	}
	schema, err := goal.NewSchema([]string{"body"}, nil)
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
	err = g.stor.CreateGoal(goal)
	if err != nil {
		return
	}
	return
}

func (g *GoalRepository) Edit(listID int64, goal *model.Goal) (err error) {
	schema, err := goal.NewSchema([]string{"body"}, nil)
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

	err = g.stor.EditGoal(listID, goal)
	if err != nil {
		return
	}
	return
}

func (g *GoalRepository) Delete(listID int64, entryID int64) (err error) {
	err = g.stor.DeleteGoal(listID, entryID)
	if err != nil {
		return
	}
	return
}

func (g *GoalRepository) List() (goals []*model.Goal, err error) {
	goals, err = g.stor.ListGoal()
	if err != nil {
		return
	}
	return
}
