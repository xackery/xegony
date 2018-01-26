package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//TopicRepository handles TopicRepository cases and is a gateway to storage
type TopicRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *TopicRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *TopicRepository) Get(topic *model.Topic, user *model.User) (err error) {

	err = c.stor.GetTopic(topic)
	if err != nil {
		err = errors.Wrap(err, "failed to get topic")
		return
	}
	err = c.prepare(topic, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare topic")
		return
	}
	return
}

//Create handles logic
func (c *TopicRepository) Create(topic *model.Topic, user *model.User) (err error) {
	if topic == nil {
		err = fmt.Errorf("Empty topic")
		return
	}
	schema, err := c.newSchema([]string{"body"}, nil)
	if err != nil {
		return
	}
	topic.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(topic))
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
	err = c.stor.CreateTopic(topic)
	if err != nil {
		return
	}
	err = c.prepare(topic, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare topic")
		return
	}
	return
}

//Edit handles logic
func (c *TopicRepository) Edit(topic *model.Topic, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"body"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(topic))
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

	err = c.stor.EditTopic(topic)
	if err != nil {
		return
	}
	err = c.prepare(topic, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare topic")
		return
	}
	return
}

//Delete handles logic
func (c *TopicRepository) Delete(topic *model.Topic, user *model.User) (err error) {
	err = c.stor.DeleteTopic(topic)
	if err != nil {
		return
	}
	return
}

//ListByForum handles logic
func (c *TopicRepository) ListByForum(forum *model.Forum, user *model.User) (topics []*model.Topic, err error) {
	topics, err = c.stor.ListTopicByForum(forum)
	if err != nil {
		return
	}
	for _, topic := range topics {
		err = c.prepare(topic, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare topic")
			return
		}
	}
	return
}

func (c *TopicRepository) prepare(topic *model.Topic, user *model.User) (err error) {

	return
}

func (c *TopicRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *TopicRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
