package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//RecipeRepository handles RecipeRepository cases and is a gateway to storage
type RecipeRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *RecipeRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *RecipeRepository) Get(recipe *model.Recipe, user *model.User) (err error) {

	err = c.stor.GetRecipe(recipe)
	if err != nil {
		err = errors.Wrap(err, "failed to get recipe")
		return
	}
	err = c.prepare(recipe, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare recipe")
		return
	}
	return
}

//Search handles logic
func (c *RecipeRepository) SearchByName(recipe *model.Recipe, user *model.User) (recipes []*model.Recipe, err error) {
	recipes, err = c.stor.SearchRecipeByName(recipe)
	if err != nil {
		return
	}
	err = c.prepare(recipe, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare recipe")
		return
	}
	return
}

//Create handles logic
func (c *RecipeRepository) Create(recipe *model.Recipe, user *model.User) (err error) {
	if recipe == nil {
		err = fmt.Errorf("Empty recipe")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	recipe.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(recipe))
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
	err = c.stor.CreateRecipe(recipe)
	if err != nil {
		return
	}
	err = c.prepare(recipe, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare recipe")
		return
	}
	return
}

//Edit handles logic
func (c *RecipeRepository) Edit(recipe *model.Recipe, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(recipe))
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

	err = c.stor.EditRecipe(recipe)
	if err != nil {
		return
	}
	err = c.prepare(recipe, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare recipe")
		return
	}
	return
}

//Delete handles logic
func (c *RecipeRepository) Delete(recipe *model.Recipe, user *model.User) (err error) {
	err = c.stor.DeleteRecipe(recipe)
	if err != nil {
		return
	}
	return
}

//ListByTradeskill handles logic
func (c *RecipeRepository) ListBySkill(skill *model.Skill, pageSize int64, pageNumber int64, user *model.User) (recipes []*model.Recipe, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	recipes, err = c.stor.ListRecipeBySkill(skill, pageSize, pageNumber)
	if err != nil {
		return
	}
	for _, recipe := range recipes {
		err = c.prepare(recipe, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare recipe")
			return
		}
	}
	return
}

//ListByTradeskillCount handles logic
func (c *RecipeRepository) ListBySkillCount(skill *model.Skill, user *model.User) (count int64, err error) {

	count, err = c.stor.ListRecipeBySkillCount(skill)
	if err != nil {
		return
	}

	return
}

//List handles logic
func (c *RecipeRepository) List(pageSize int64, pageNumber int64, user *model.User) (recipes []*model.Recipe, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	recipes, err = c.stor.ListRecipe(pageSize, pageNumber)
	if err != nil {
		return
	}
	for _, recipe := range recipes {
		err = c.prepare(recipe, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare recipe")
			return
		}
	}
	return
}

//ListCount handles logic
func (c *RecipeRepository) ListCount(user *model.User) (count int64, err error) {

	count, err = c.stor.ListRecipeCount()
	if err != nil {
		return
	}
	return
}

func (c *RecipeRepository) prepare(recipe *model.Recipe, user *model.User) (err error) {

	return
}

func (c *RecipeRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *RecipeRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
