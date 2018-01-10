package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//AaRankRepository handles AaRankRepository cases and is a gateway to storage
type AaRankRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *AaRankRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *AaRankRepository) Get(rankID int64) (aaEntry *model.AaRank, query string, err error) {

	if rankID == 0 {
		err = fmt.Errorf("Invalid Rank ID")
		return
	}
	query, aaEntry, err = c.stor.GetAaRank(rankID)
	return
}

//Create handles logic
func (c *AaRankRepository) Create(aaEntry *model.AaRank) (query string, err error) {
	if aaEntry == nil {
		err = fmt.Errorf("Empty AaRank")
		return
	}

	schema, err := c.newSchema(nil, nil)
	if err != nil {
		return
	}
	result, err := schema.Validate(gojsonschema.NewGoLoader(aaEntry))
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
	query, err = c.stor.CreateAaRank(aaEntry)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *AaRankRepository) Edit(rankID int64, aaEntry *model.AaRank) (query string, err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(aaEntry))
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

	query, err = c.stor.EditAaRank(rankID, aaEntry)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *AaRankRepository) Delete(rankID int64) (query string, err error) {
	query, err = c.stor.DeleteAaRank(rankID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *AaRankRepository) List() (aaEntrys []*model.AaRank, query string, err error) {
	query, aaEntrys, err = c.stor.ListAaRank()
	if err != nil {
		return
	}
	return
}

func (c *AaRankRepository) prepare(aarank *model.AaRank) (err error) {

	return
}

func (c *AaRankRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *AaRankRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
