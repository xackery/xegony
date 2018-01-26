package cases

import (
	"fmt"

	"github.com/pkg/errors"
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
func (c *RecipeEntryRepository) Get(recipeEntry *model.RecipeEntry, user *model.User) (err error) {

	err = c.stor.GetRecipeEntry(recipeEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to get recipe")
		return
	}
	err = c.prepare(recipeEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get recipe")
	}
	return
}

//Create handles logic
func (c *RecipeEntryRepository) Create(recipeEntry *model.RecipeEntry, user *model.User) (err error) {
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
	err = c.stor.CreateRecipeEntry(recipeEntry)
	if err != nil {
		return
	}
	err = c.prepare(recipeEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get recipe")
	}
	return
}

//Edit handles logic
func (c *RecipeEntryRepository) Edit(recipeEntry *model.RecipeEntry, user *model.User) (err error) {
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

	err = c.stor.EditRecipeEntry(recipeEntry)
	if err != nil {
		return
	}
	err = c.prepare(recipeEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get recipe")
	}
	return
}

//Delete handles logic
func (c *RecipeEntryRepository) Delete(recipeEntry *model.RecipeEntry, user *model.User) (err error) {
	err = c.stor.DeleteRecipeEntry(recipeEntry)
	if err != nil {
		return
	}
	return
}

//ListByRecipe handles logic
func (c *RecipeEntryRepository) ListByRecipe(recipe *model.Recipe, user *model.User) (recipeEntrys []*model.RecipeEntry, err error) {
	recipeEntrys, err = c.stor.ListRecipeEntryByRecipe(recipe)
	if err != nil {
		return
	}
	for _, recipeEntry := range recipeEntrys {
		err = c.prepare(recipeEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get recipe")
		}
	}
	return
}

//ListByItem handles logic
func (c *RecipeEntryRepository) ListByItem(item *model.Item, user *model.User) (recipeEntrys []*model.RecipeEntry, err error) {
	recipeEntrys, err = c.stor.ListRecipeEntryByItem(item)
	if err != nil {
		return
	}
	for _, recipeEntry := range recipeEntrys {
		err = c.prepare(recipeEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get recipe")
		}
	}
	return
}

func (c *RecipeEntryRepository) prepare(recipeEntry *model.RecipeEntry, user *model.User) (err error) {

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
