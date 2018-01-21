package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//MailRepository handles MailRepository cases and is a gateway to storage
type MailRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *MailRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *MailRepository) Get(mail *model.Mail, user *model.User) (err error) {

	err = c.stor.GetMail(mail)
	if err != nil {
		err = errors.Wrap(err, "failed to get mail")
		return
	}
	err = c.prepare(mail, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare mail")
		return
	}
	return
}

//Search handles logic
func (c *MailRepository) Search(mail *model.Mail, user *model.User) (mails []*model.Mail, err error) {
	mails, err = c.stor.SearchMailByBody(mail)
	if err != nil {
		return
	}
	for _, mail := range mails {
		err = c.prepare(mail, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare mail")
			return
		}
	}
	return
}

//Create handles logic
func (c *MailRepository) Create(mail *model.Mail, user *model.User) (err error) {
	if mail == nil {
		err = fmt.Errorf("Empty mail")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	mail.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(mail))
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
	err = c.stor.CreateMail(mail)
	if err != nil {
		return
	}
	err = c.prepare(mail, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare mail")
		return
	}
	return
}

//Edit handles logic
func (c *MailRepository) Edit(mail *model.Mail, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(mail))
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

	err = c.stor.EditMail(mail)
	if err != nil {
		return
	}
	err = c.prepare(mail, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare mail")
		return
	}
	return
}

//Delete handles logic
func (c *MailRepository) Delete(mail *model.Mail, user *model.User) (err error) {
	err = c.stor.DeleteMail(mail)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *MailRepository) List(pageSize int64, pageNumber int64, user *model.User) (mails []*model.Mail, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	mails, err = c.stor.ListMail(pageSize, pageNumber)
	if err != nil {
		return
	}
	for _, mail := range mails {
		err = c.prepare(mail, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare mail")
			return
		}
	}
	return
}

//ListCount handles logic
func (c *MailRepository) ListCount(user *model.User) (count int64, err error) {

	count, err = c.stor.ListMailCount()
	if err != nil {
		return
	}
	return
}

//ListByCharacter handles logic
func (c *MailRepository) ListByCharacter(character *model.Character, user *model.User) (mails []*model.Mail, err error) {
	mails, err = c.stor.ListMailByCharacter(character)
	if err != nil {
		return
	}
	for _, mail := range mails {
		err = c.prepare(mail, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare mail")
			return
		}
	}
	return
}

func (c *MailRepository) prepare(mail *model.Mail, user *model.User) (err error) {

	return
}

func (c *MailRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *MailRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "zoneID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
