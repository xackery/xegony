package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//RecipeEntryRepository handles RecipeEntryRepository cases and is a gateway to storage
type RecipeEntryRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *RecipeEntryRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *RecipeEntryRepository) Get(recipeID int64, itemID int64) (recipeEntry *model.RecipeEntry, query string, err error) {
	if recipeID == 0 {
		err = fmt.Errorf("Invalid RecipeEntry ID")
		return
	}
	if itemID == 0 {
		err = fmt.Errorf("Invalid Item ID")
		return
	}
	query, recipeEntry, err = c.stor.GetRecipeEntry(recipeID, itemID)
	return
}

//Create handles logic
func (c *RecipeEntryRepository) Create(recipeEntry *model.RecipeEntry) (query string, err error) {
	if recipeEntry == nil {
		err = fmt.Errorf("Empty RecipeEntry")
		return
	}
	if recipeEntry.RecipeID == 0 {
		err = fmt.Errorf("Invalid RecipeGroup ID")
		return
	}
	if recipeEntry.ItemID == 0 {
		err = fmt.Errorf("Invalid Item ID")
		return
	}
	schema, err := c.newSchema(nil, nil)
	if err != nil {
		return
	}
	result, err := schema.Validate(gojsonschema.NewGoLoader(recipeEntry))
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
	query, err = c.stor.CreateRecipeEntry(recipeEntry)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *RecipeEntryRepository) Edit(recipeID int64, itemID int64, recipeEntry *model.RecipeEntry) (query string, err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(recipeEntry))
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

	query, err = c.stor.EditRecipeEntry(recipeID, itemID, recipeEntry)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *RecipeEntryRepository) Delete(recipeID int64, itemID int64) (query string, err error) {
	query, err = c.stor.DeleteRecipeEntry(recipeID, itemID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *RecipeEntryRepository) List(recipeID int64) (recipeEntrys []*model.RecipeEntry, query string, err error) {
	query, recipeEntrys, err = c.stor.ListRecipeEntry(recipeID)
	if err != nil {
		return
	}
	return
}

//ListByItem handles logic
func (c *RecipeEntryRepository) ListByItem(itemID int64) (recipeEntrys []*model.RecipeEntry, query string, err error) {
	query, recipeEntrys, err = c.stor.ListRecipeEntryByItem(itemID)
	if err != nil {
		return
	}
	return
}

func (c *RecipeEntryRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *RecipeEntryRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
