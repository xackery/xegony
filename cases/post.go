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
	schema, err := g.newSchema([]string{"body"}, nil)
	if err != nil {
		return
	}
	post.ID = 0 //strip ID
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
	schema, err := g.newSchema([]string{"body"}, nil)
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

func (c *PostRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *PostRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "body":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 1024
		prop.Pattern = "^[a-zA-Z' ]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
