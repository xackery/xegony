package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type ForumRepository struct {
	stor storage.Storage
}

func (g *ForumRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *ForumRepository) Get(forumID int64) (forum *model.Forum, err error) {
	if forumID == 0 {
		err = fmt.Errorf("Invalid Forum ID")
		return
	}
	forum, err = g.stor.GetForum(forumID)
	return
}

func (g *ForumRepository) Create(forum *model.Forum) (err error) {
	if forum == nil {
		err = fmt.Errorf("Empty forum")
		return
	}
	schema, err := forum.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	forum.Id = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(forum))
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
	err = g.stor.CreateForum(forum)
	if err != nil {
		return
	}
	return
}

func (g *ForumRepository) Edit(forumID int64, forum *model.Forum) (err error) {
	schema, err := forum.NewSchema([]string{"name"}, []string{"description"})
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(forum))
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

	err = g.stor.EditForum(forumID, forum)
	if err != nil {
		return
	}
	return
}

func (g *ForumRepository) Delete(forumID int64) (err error) {
	err = g.stor.DeleteForum(forumID)
	if err != nil {
		return
	}
	return
}

func (g *ForumRepository) List() (forums []*model.Forum, err error) {
	forums, err = g.stor.ListForum()
	if err != nil {
		return
	}
	return
}
