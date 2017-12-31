package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type PostRepository struct {
	stor storage.Storage
}

func (g *PostRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *PostRepository) Get(postID int64) (post *model.Post, err error) {
	if postID == 0 {
		err = fmt.Errorf("Invalid Post ID")
		return
	}
	post, err = g.stor.GetPost(postID)
	return
}

func (g *PostRepository) Create(post *model.Post) (err error) {
	if post == nil {
		err = fmt.Errorf("Empty post")
		return
	}
	schema, err := post.NewSchema([]string{"body"}, nil)
	if err != nil {
		return
	}
	post.Id = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(post))
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
	err = g.stor.CreatePost(post)
	if err != nil {
		return
	}
	return
}

func (g *PostRepository) Edit(postID int64, post *model.Post) (err error) {
	schema, err := post.NewSchema([]string{"body"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(post))
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

	err = g.stor.EditPost(postID, post)
	if err != nil {
		return
	}
	return
}

func (g *PostRepository) Delete(postID int64) (err error) {
	err = g.stor.DeletePost(postID)
	if err != nil {
		return
	}
	return
}

func (g *PostRepository) List(topicID int64) (posts []*model.Post, err error) {
	posts, err = g.stor.ListPost(topicID)
	if err != nil {
		return
	}
	return
}
