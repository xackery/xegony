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

func (g *TopicRepository) Get(topicId int64) (topic *model.Topic, err error) {
	if topicId == 0 {
		err = fmt.Errorf("Invalid Topic ID")
		return
	}
	topic, err = g.stor.GetTopic(topicId)
	return
}

func (g *TopicRepository) Create(topic *model.Topic) (err error) {
	if topic == nil {
		err = fmt.Errorf("Empty topic")
		return
	}
	schema, err := topic.NewSchema([]string{"body"}, nil)
	if err != nil {
		return
	}
	topic.Id = 0 //strip ID
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

func (g *TopicRepository) Edit(topicId int64, topic *model.Topic) (err error) {
	schema, err := topic.NewSchema([]string{"body"}, nil)
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

	err = g.stor.EditTopic(topicId, topic)
	if err != nil {
		return
	}
	return
}

func (g *TopicRepository) Delete(topicId int64) (err error) {
	err = g.stor.DeleteTopic(topicId)
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
