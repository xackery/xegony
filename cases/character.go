package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//CharacterRepository handles CharacterRepository cases and is a gateway to storage
type CharacterRepository struct {
	stor storage.Storage
}

func (c *CharacterRepository) isStorageInitialized() (err error) {
	if c.stor == nil {
		err = fmt.Errorf("Storage not initialized")
		return
	}
	return
}

//Initialize handles logic
func (c *CharacterRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *CharacterRepository) Get(characterID int64) (character *model.Character, err error) {
	if characterID == 0 {
		err = fmt.Errorf("Invalid Character ID")
		return
	}
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	character, err = c.stor.GetCharacter(characterID)
	return
}

//GetByName handles logic
func (c *CharacterRepository) GetByName(name string) (character *model.Character, err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	character, err = c.stor.GetCharacterByName(name)
	return
}

//Create handles logic
func (c *CharacterRepository) Create(character *model.Character) (err error) {
	if character == nil {
		err = fmt.Errorf("Empty character")
		return
	}
	schema, err := c.newSchema([]string{"name", "accountID"}, nil)
	if err != nil {
		return
	}
	if character.AccountID < 1 {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		vErr.Reasons["accountID"] = "Account ID must be greater than 0"
		err = vErr
		return
	}
	character.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(character))
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
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	err = c.stor.CreateCharacter(character)
	if err != nil {
		return
	}
	return
}

//Search handles logic
func (c *CharacterRepository) Search(search string) (characters []*model.Character, err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	characters, err = c.stor.SearchCharacter(search)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *CharacterRepository) Edit(characterID int64, character *model.Character) (err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(character))
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

	err = c.stor.EditCharacter(characterID, character)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *CharacterRepository) Delete(characterID int64) (err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	err = c.stor.DeleteCharacter(characterID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *CharacterRepository) List() (characters []*model.Character, err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	characters, err = c.stor.ListCharacter()
	if err != nil {
		return
	}
	return
}

//ListByRanking handles logic
func (c *CharacterRepository) ListByRanking() (characters []*model.Character, err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	characters, err = c.stor.ListCharacterByRanking()
	if err != nil {
		return
	}
	return
}

//ListByOnline handles logic
func (c *CharacterRepository) ListByOnline() (characters []*model.Character, err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	characters, err = c.stor.ListCharacterByOnline()
	if err != nil {
		return
	}
	return
}

//ListByAccount handles logic
func (c *CharacterRepository) ListByAccount(accountID int64) (characters []*model.Character, err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	characters, err = c.stor.ListCharacterByAccount(accountID)
	if err != nil {
		return
	}
	return
}

func (c *CharacterRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *CharacterRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "accountID":
		prop.Type = "integer"
		prop.Minimum = 1
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
