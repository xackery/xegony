package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListSpell lists all spells accessible by provided user
func ListSpell(page *model.Page, user *model.User) (spells []*model.Spell, err error) {
	err = validateOrderBySpellField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("spell")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for spell")
		return
	}

	page.Total, err = reader.ListSpellTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list spell toal count")
		return
	}

	spells, err = reader.ListSpell(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spell")
		return
	}
	for i, spell := range spells {
		err = sanitizeSpell(spell, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize spell element %d", i)
			return
		}
	}
	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}

	return
}

//ListSpellBySearch will request any spell matching the pattern of name
func ListSpellBySearch(page *model.Page, spell *model.Spell, user *model.User) (spells []*model.Spell, err error) {
	/*err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spell by search without guide+")
		return
	}
	*/
	err = validateOrderBySpellField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre spell")
		return
	}

	err = validateSpell(spell, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate spell")
		return
	}
	reader, err := getReader("spell")
	if err != nil {
		err = errors.Wrap(err, "failed to get spell reader")
		return
	}

	spells, err = reader.ListSpellBySearch(page, spell)
	if err != nil {
		err = errors.Wrap(err, "failed to list spell by search")
		return
	}

	page.Total, err = reader.ListSpellBySearchTotalCount(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to get page total")
		return
	}

	for _, spell := range spells {
		err = sanitizeSpell(spell, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize spell")
			return
		}
	}

	err = sanitizeSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search spell")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSpell will create an spell using provided information
func CreateSpell(spell *model.Spell, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spell by search without guide+")
		return
	}
	err = prepareSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spell")
		return
	}

	err = validateSpell(spell, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spell")
		return
	}
	spell.ID = 0
	//spell.TimeCreation = time.Now().Unix()
	writer, err := getWriter("spell")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spell")
		return
	}
	err = writer.CreateSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to create spell")
		return
	}

	memWriter, err := getWriter("spell")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spell")
		return
	}
	err = memWriter.CreateSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spell")
		return
	}

	err = sanitizeSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spell")
		return
	}
	return
}

//GetSpell gets an spell by provided spellID
func GetSpell(spell *model.Spell, user *model.User) (err error) {
	err = prepareSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spell")
		return
	}

	err = validateSpell(spell, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spell")
		return
	}

	reader, err := getReader("spell")
	if err != nil {
		err = errors.Wrap(err, "failed to get spell reader")
		return
	}

	err = reader.GetSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to get spell")
		return
	}

	err = sanitizeSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spell")
		return
	}

	return
}

//EditSpell edits an existing spell
func EditSpell(spell *model.Spell, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spell by search without guide+")
		return
	}
	err = prepareSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spell")
		return
	}

	err = validateSpell(spell,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spell")
		return
	}
	writer, err := getWriter("spell")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spell")
		return
	}
	err = writer.EditSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spell")
		return
	}

	memWriter, err := getWriter("spell")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spell")
		return
	}
	err = memWriter.EditSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spell")
		return
	}

	err = sanitizeSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spell")
		return
	}
	return
}

//DeleteSpell deletes an spell by provided spellID
func DeleteSpell(spell *model.Spell, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete spell without admin+")
		return
	}
	err = prepareSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spell")
		return
	}

	err = validateSpell(spell, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spell")
		return
	}
	writer, err := getWriter("spell")
	if err != nil {
		err = errors.Wrap(err, "failed to get spell writer")
		return
	}
	err = writer.DeleteSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spell")
		return
	}

	memWriter, err := getWriter("spell")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spell")
		return
	}
	err = memWriter.DeleteSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spell")
		return
	}
	return
}

func prepareSpell(spell *model.Spell, user *model.User) (err error) {
	if spell == nil {
		err = fmt.Errorf("empty spell")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSpell(spell *model.Spell, required []string, optional []string) (err error) {
	schema, err := newSchemaSpell(required, optional)
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
	return
}

func validateOrderBySpellField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "name"
	}

	validNames := []string{
		"id",
		"name",
	}

	possibleNames := ""
	for _, name := range validNames {
		if page.OrderBy == name {
			return
		}
		possibleNames += name + ", "
	}
	if len(possibleNames) > 0 {
		possibleNames = possibleNames[0 : len(possibleNames)-2]
	}
	err = &model.ErrValidation{
		Message: "orderBy is invalid. Possible fields: " + possibleNames,
		Reasons: map[string]string{
			"orderBy": "field is not valid",
		},
	}
	return
}

func sanitizeSpell(spell *model.Spell, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}

	if spell.AnimationID > 0 {
		spell.Animation = &model.SpellAnimation{
			ID: spell.AnimationID,
		}
		err = GetSpellAnimation(spell.Animation, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get spell animation during sanitize spell")
			return
		}
	}

	return
}

func newSchemaSpell(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySpell(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySpell(field); err != nil {
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

func getSchemaPropertySpell(field string) (prop model.Schema, err error) {
	switch field {

	case "ID":
		prop.Type = "integer"
		prop.Minimum = 0
	case "name":
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
