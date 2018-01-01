package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type TopicRepository struct {
	stor storage.Storage
}

func (g *TopicRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *TopicRepository) Get(topicID int64) (topic *model.Topic, err error) {
	if topicID == 0 {
		err = fmt.Errorf("Invalid Topic ID")
		return
	}
	topic, err = g.stor.GetTopic(topicID)
	return
}

func (g *TopicRepository) Create(topic *model.Topic) (err error) {
	if topic == nil {
		err = fmt.Errorf("Empty topic")
		return
	}
	schema, err := g.newSchema([]string{"body"}, nil)
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
	err = g.stor.CreateTopic(topic)
	if err != nil {
		return
	}
	return
}

func (g *TopicRepository) Edit(topicID int64, topic *model.Topic) (err error) {
	schema, err := g.newSchema([]string{"body"}, nil)
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

	err = g.stor.EditTopic(topicID, topic)
	if err != nil {
		return
	}
	return
}

func (g *TopicRepository) Delete(topicID int64) (err error) {
	err = g.stor.DeleteTopic(topicID)
	if err != nil {
		return
	}
	return
}

func (g *TopicRepository) List(forumID int64) (topics []*model.Topic, err error) {
	topics, err = g.stor.ListTopic(forumID)
	if err != nil {
		return
	}
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
