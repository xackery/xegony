package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//SpellRepository handles SpellRepository cases and is a gateway to storage
type SpellRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *SpellRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *SpellRepository) Get(spellID int64) (spell *model.Spell, err error) {
	if spellID == 0 {
		err = fmt.Errorf("Invalid Spell ID")
		return
	}
	spell, err = c.stor.GetSpell(spellID)
	return
}

//Search handles logic
func (c *SpellRepository) Search(search string) (spells []*model.Spell, err error) {
	spells, err = c.stor.SearchSpell(search)
	if err != nil {
		return
	}
	return
}

//Create handles logic
func (c *SpellRepository) Create(spell *model.Spell) (err error) {
	if spell == nil {
		err = fmt.Errorf("Empty spell")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	spell.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(spell))
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
	err = c.stor.CreateSpell(spell)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *SpellRepository) Edit(spellID int64, spell *model.Spell) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spell))
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

	err = c.stor.EditSpell(spellID, spell)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *SpellRepository) Delete(spellID int64) (err error) {
	err = c.stor.DeleteSpell(spellID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *SpellRepository) List(pageSize int64, pageNumber int64) (spells []*model.Spell, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	spells, err = c.stor.ListSpell(pageSize, pageNumber)
	if err != nil {
		return
	}
	return
}

//ListCount handles logic
func (c *SpellRepository) ListCount() (count int64, err error) {

	count, err = c.stor.ListSpellCount()
	if err != nil {
		return
	}
	return
}

func (c *SpellRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *SpellRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
