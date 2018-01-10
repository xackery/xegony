package cases

import (
	"fmt"

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
func (c *MailRepository) Get(mailID int64) (mail *model.Mail, err error) {
	if mailID == 0 {
		err = fmt.Errorf("Invalid Mail ID")
		return
	}
	mail, err = c.stor.GetMail(mailID)
	return
}

//Search handles logic
func (c *MailRepository) Search(search string) (mails []*model.Mail, err error) {
	mails, err = c.stor.SearchMail(search)
	if err != nil {
		return
	}
	return
}

//Create handles logic
func (c *MailRepository) Create(mail *model.Mail) (err error) {
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
	return
}

//Edit handles logic
func (c *MailRepository) Edit(mailID int64, mail *model.Mail) (err error) {
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

	err = c.stor.EditMail(mailID, mail)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *MailRepository) Delete(mailID int64) (err error) {
	err = c.stor.DeleteMail(mailID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *MailRepository) List(pageSize int64, pageNumber int64) (mails []*model.Mail, err error) {
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
	return
}

//ListCount handles logic
func (c *MailRepository) ListCount() (count int64, err error) {

	count, err = c.stor.ListMailCount()
	if err != nil {
		return
	}
	return
}

//ListByCharacter handles logic
func (c *MailRepository) ListByCharacter(characterID int64) (mails []*model.Mail, err error) {
	mails, err = c.stor.ListMailByCharacter(characterID)
	if err != nil {
		return
	}
	return
}

func (c *MailRepository) prepare(mail *model.Mail) (err error) {

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
