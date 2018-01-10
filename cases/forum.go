package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//ForumRepository handles ForumRepository cases and is a gateway to storage
type ForumRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *ForumRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *ForumRepository) Get(forumID int64) (forum *model.Forum, err error) {
	if forumID == 0 {
		err = fmt.Errorf("Invalid Forum ID")
		return
	}
	forum, err = c.stor.GetForum(forumID)
	return
}

//Create handles logic
func (c *ForumRepository) Create(forum *model.Forum) (err error) {
	if forum == nil {
		err = fmt.Errorf("Empty forum")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	forum.ID = 0 //strip ID
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
	err = c.stor.CreateForum(forum)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *ForumRepository) Edit(forumID int64, forum *model.Forum) (err error) {
	schema, err := c.newSchema([]string{"name"}, []string{"description"})
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

	err = c.stor.EditForum(forumID, forum)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *ForumRepository) Delete(forumID int64) (err error) {
	err = c.stor.DeleteForum(forumID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *ForumRepository) List() (forums []*model.Forum, err error) {
	forums, err = c.stor.ListForum()
	if err != nil {
		return
	}
	return
}

func (c *ForumRepository) prepare(aa *model.Forum) (err error) {

	return
}

func (c *ForumRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *ForumRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "ownerId":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z' ]*$"
	case "description":
		prop.Type = "string"
		prop.MaxLength = 128
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
