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

	page.Total, err = reader.ListSpellTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list spell toal count")
		return
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

	if spell.BuffDurationFormulaID > 0 {
		spell.BuffDurationFormula = &model.SpellDurationFormula{
			ID: spell.BuffDurationFormulaID,
		}
		err = GetSpellDurationFormula(spell.BuffDurationFormula, user)
		if err != nil {
			err = errors.Wrap(err, "failed on formula 1")
			return
		}
	}

	if spell.EffectID1 > 0 {
		spell.Effect1 = &model.SpellEffectType{
			ID: spell.EffectID1,
		}
		err = GetSpellEffectType(spell.Effect1, user)
		if err != nil {
			err = errors.Wrap(err, "failed on effect 1")
			return
		}
	}
	if spell.EffectID2 > 0 {
		spell.Effect2 = &model.SpellEffectType{
			ID: spell.EffectID2,
		}
		err = GetSpellEffectType(spell.Effect2, user)
		if err != nil {
			err = errors.Wrap(err, "failed on effect 2")
			return
		}
	}
	if spell.EffectID3 > 0 {
		spell.Effect3 = &model.SpellEffectType{
			ID: spell.EffectID3,
		}
		err = GetSpellEffectType(spell.Effect3, user)
		if err != nil {
			err = errors.Wrap(err, "failed on effect 3")
			return
		}
	}
	if spell.EffectID4 > 0 {
		spell.Effect4 = &model.SpellEffectType{
			ID: spell.EffectID4,
		}
		err = GetSpellEffectType(spell.Effect4, user)
		if err != nil {
			err = errors.Wrap(err, "failed on effect 4")
			return
		}
	}
	if spell.EffectID5 > 0 {
		spell.Effect5 = &model.SpellEffectType{
			ID: spell.EffectID5,
		}
		err = GetSpellEffectType(spell.Effect5, user)
		if err != nil {
			err = errors.Wrap(err, "failed on effect 5")
			return
		}
	}
	if spell.EffectID6 > 0 {
		spell.Effect6 = &model.SpellEffectType{
			ID: spell.EffectID6,
		}
		err = GetSpellEffectType(spell.Effect6, user)
		if err != nil {
			err = errors.Wrap(err, "failed on effect 6")
			return
		}
	}
	if spell.EffectID7 > 0 {
		spell.Effect7 = &model.SpellEffectType{
			ID: spell.EffectID7,
		}
		err = GetSpellEffectType(spell.Effect7, user)
		if err != nil {
			err = errors.Wrap(err, "failed on effect 7")
			return
		}
	}
	if spell.EffectID8 > 0 {
		spell.Effect8 = &model.SpellEffectType{
			ID: spell.EffectID8,
		}
		err = GetSpellEffectType(spell.Effect8, user)
		if err != nil {
			err = errors.Wrap(err, "failed on effect 8")
			return
		}
	}
	if spell.EffectID9 > 0 {
		spell.Effect9 = &model.SpellEffectType{
			ID: spell.EffectID9,
		}
		err = GetSpellEffectType(spell.Effect9, user)
		if err != nil {
			err = errors.Wrap(err, "failed on effect 9")
			return
		}
	}
	if spell.EffectID10 > 0 {
		spell.Effect10 = &model.SpellEffectType{
			ID: spell.EffectID10,
		}
		err = GetSpellEffectType(spell.Effect10, user)
		if err != nil {
			err = errors.Wrap(err, "failed on effect 10")
			return
		}
	}
	if spell.EffectID11 > 0 {
		spell.Effect11 = &model.SpellEffectType{
			ID: spell.EffectID11,
		}
		err = GetSpellEffectType(spell.Effect11, user)
		if err != nil {
			err = errors.Wrap(err, "failed on effect 11")
			return
		}
	}
	if spell.EffectID12 > 0 {
		spell.Effect12 = &model.SpellEffectType{
			ID: spell.EffectID12,
		}
		err = GetSpellEffectType(spell.Effect12, user)
		if err != nil {
			err = errors.Wrap(err, "failed on effect 12")
			return
		}
	}

	if spell.DeityID0 > 0 {
		spell.Deity0 = &model.Deity{
			SpellID: 0,
		}
		err = GetDeityBySpell(spell, spell.Deity0, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 0")
			return
		}
	}
	if spell.DeityID1 > 0 {
		spell.Deity1 = &model.Deity{
			SpellID: 1,
		}
		err = GetDeityBySpell(spell, spell.Deity1, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 1")
			return
		}
	}
	if spell.DeityID2 > 0 {
		spell.Deity2 = &model.Deity{
			SpellID: 2,
		}
		err = GetDeityBySpell(spell, spell.Deity2, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 2")
			return
		}
	}
	if spell.DeityID3 > 0 {
		spell.Deity3 = &model.Deity{
			SpellID: 3,
		}
		err = GetDeityBySpell(spell, spell.Deity3, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 3")
			return
		}
	}
	if spell.DeityID4 > 0 {
		spell.Deity4 = &model.Deity{
			SpellID: 4,
		}
		err = GetDeityBySpell(spell, spell.Deity4, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 4")
			return
		}
	}
	if spell.DeityID5 > 0 {
		spell.Deity5 = &model.Deity{
			SpellID: 5,
		}
		err = GetDeityBySpell(spell, spell.Deity5, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 5")
			return
		}
	}
	if spell.DeityID6 > 0 {
		spell.Deity6 = &model.Deity{
			SpellID: 6,
		}
		err = GetDeityBySpell(spell, spell.Deity6, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 6")
			return
		}
	}
	if spell.DeityID7 > 0 {
		spell.Deity7 = &model.Deity{
			SpellID: 7,
		}
		err = GetDeityBySpell(spell, spell.Deity7, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 7")
			return
		}
	}
	if spell.DeityID8 > 0 {
		spell.Deity8 = &model.Deity{
			SpellID: 8,
		}
		err = GetDeityBySpell(spell, spell.Deity8, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 8")
			return
		}
	}
	if spell.DeityID9 > 0 {
		spell.Deity9 = &model.Deity{
			SpellID: 9,
		}
		err = GetDeityBySpell(spell, spell.Deity9, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 9")
			return
		}
	}
	if spell.DeityID10 > 0 {
		spell.Deity10 = &model.Deity{
			SpellID: 10,
		}
		err = GetDeityBySpell(spell, spell.Deity10, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 10")
			return
		}
	}
	if spell.DeityID11 > 0 {
		spell.Deity11 = &model.Deity{
			SpellID: 11,
		}
		err = GetDeityBySpell(spell, spell.Deity11, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 11")
			return
		}
	}
	if spell.DeityID12 > 0 {
		spell.Deity12 = &model.Deity{
			SpellID: 12,
		}
		err = GetDeityBySpell(spell, spell.Deity12, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 12")
			return
		}
	}
	if spell.DeityID13 > 0 {
		spell.Deity13 = &model.Deity{
			SpellID: 13,
		}
		err = GetDeityBySpell(spell, spell.Deity13, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 13")
			return
		}
	}
	if spell.DeityID14 > 0 {
		spell.Deity14 = &model.Deity{
			SpellID: 14,
		}
		err = GetDeityBySpell(spell, spell.Deity14, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 14")
			return
		}
	}
	if spell.DeityID15 > 0 {
		spell.Deity15 = &model.Deity{
			SpellID: 15,
		}
		err = GetDeityBySpell(spell, spell.Deity15, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 15")
			return
		}
	}
	if spell.DeityID16 > 0 {
		spell.Deity16 = &model.Deity{
			SpellID: 16,
		}
		err = GetDeityBySpell(spell, spell.Deity16, user)
		if err != nil {
			err = errors.Wrap(err, "failed on deity 16")
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
