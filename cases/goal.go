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

func (g *GoalRepository) Get(listId int64, entryId int64) (goal *model.Goal, err error) {
	if listId == 0 {
		err = fmt.Errorf("Invalid List ID")
		return
	}
	if entryId == 0 {
		err = fmt.Errorf("Invalid Entry ID")
		return
	}
	goal, err = g.stor.GetGoal(listId, entryId)
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

func (g *GoalRepository) Edit(listId int64, goal *model.Goal) (err error) {
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

	err = g.stor.EditGoal(listId, goal)
	if err != nil {
		return
	}
	return
}

func (g *GoalRepository) Delete(listId int64, entryId int64) (err error) {
	err = g.stor.DeleteGoal(listId, entryId)
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
