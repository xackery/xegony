package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//PostRepository handles PostRepository cases and is a gateway to storage
type PostRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *PostRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *PostRepository) Get(post *model.Post, user *model.User) (err error) {
	err = c.stor.GetPost(post)
	if err != nil {
		err = errors.Wrap(err, "failed to get post")
		return
	}
	err = c.prepare(post, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare post")
		return
	}
	return
}

//Create handles logic
func (c *PostRepository) Create(post *model.Post, user *model.User) (err error) {
	if post == nil {
		err = fmt.Errorf("Empty post")
		return
	}
	schema, err := c.newSchema([]string{"body"}, nil)
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
	err = c.stor.CreatePost(post)
	if err != nil {
		return
	}
	err = c.prepare(post, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare post")
		return
	}
	return
}

//Edit handles logic
func (c *PostRepository) Edit(post *model.Post, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"body"}, nil)
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

	err = c.stor.EditPost(post)
	if err != nil {
		return
	}
	err = c.prepare(post, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare post")
		return
	}
	return
}

//Delete handles logic
func (c *PostRepository) Delete(post *model.Post, user *model.User) (err error) {
	err = c.stor.DeletePost(post)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *PostRepository) ListByTopic(topic *model.Topic, user *model.User) (posts []*model.Post, err error) {
	posts, err = c.stor.ListPostByTopic(topic)
	if err != nil {
		return
	}
	for _, post := range posts {
		err = c.prepare(post, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare post")
			return
		}
	}
	return
}

func (c *PostRepository) prepare(post *model.Post, user *model.User) (err error) {

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
