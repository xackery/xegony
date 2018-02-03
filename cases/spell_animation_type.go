package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadSpellAnimationTypeFromFileToMemory is ran during initialization
func LoadSpellAnimationTypeFromFileToMemory() (err error) {

	fr, err := file.New("config", "spellAnimationType.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("spellAnimationType-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellAnimationType-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("spellAnimationType-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellAnimationType-memory")
		return
	}

	fileReader, err := getReader("spellAnimationType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimationType-file reader")
		return
	}

	memWriter, err := getWriter("spellAnimationType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimationType-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListSpellAnimationTypeTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list spellAnimationType count")
		return
	}
	page.Limit = page.Total

	spellAnimationTypes, err := fileReader.ListSpellAnimationType(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellAnimationTypes")
		return
	}

	for _, spellAnimationType := range spellAnimationTypes {
		err = memWriter.CreateSpellAnimationType(spellAnimationType)
		if err != nil {
			err = errors.Wrap(err, "failed to create spellAnimationType")
			return
		}
	}

	fmt.Printf("%d spellAnimationTypes, ", len(spellAnimationTypes))
	return
}

//ListSpellAnimationType lists all spellAnimationTypes accessible by provided user
func ListSpellAnimationType(page *model.Page, user *model.User) (spellAnimationTypes []*model.SpellAnimationType, err error) {
	err = validateOrderBySpellAnimationTypeField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("spellAnimationType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for spellAnimationType")
		return
	}

	page.Total, err = reader.ListSpellAnimationTypeTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list spellAnimationType toal count")
		return
	}

	spellAnimationTypes, err = reader.ListSpellAnimationType(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellAnimationType")
		return
	}
	for i, spellAnimationType := range spellAnimationTypes {
		err = sanitizeSpellAnimationType(spellAnimationType, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize spellAnimationType element %d", i)
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

//ListSpellAnimationTypeBySearch will request any spellAnimationType matching the pattern of name
func ListSpellAnimationTypeBySearch(page *model.Page, spellAnimationType *model.SpellAnimationType, user *model.User) (spellAnimationTypes []*model.SpellAnimationType, err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellAnimationType by search without guide+")
		return
	}

	err = validateOrderBySpellAnimationTypeField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSpellAnimationType(spellAnimationType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre spellAnimationType")
		return
	}

	err = validateSpellAnimationType(spellAnimationType, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellAnimationType")
		return
	}
	reader, err := getReader("spellAnimationType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimationType-memory reader")
		return
	}

	spellAnimationTypes, err = reader.ListSpellAnimationTypeBySearch(page, spellAnimationType)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellAnimationType by search")
		return
	}

	for _, spellAnimationType := range spellAnimationTypes {
		err = sanitizeSpellAnimationType(spellAnimationType, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize spellAnimationType")
			return
		}
	}

	err = sanitizeSpellAnimationType(spellAnimationType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search spellAnimationType")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSpellAnimationType will create an spellAnimationType using provided information
func CreateSpellAnimationType(spellAnimationType *model.SpellAnimationType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellAnimationType by search without guide+")
		return
	}
	err = prepareSpellAnimationType(spellAnimationType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellAnimationType")
		return
	}

	err = validateSpellAnimationType(spellAnimationType, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellAnimationType")
		return
	}
	spellAnimationType.ID = 0
	writer, err := getWriter("spellAnimationType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellAnimationType")
		return
	}
	err = writer.CreateSpellAnimationType(spellAnimationType)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellAnimationType")
		return
	}

	fileWriter, err := getWriter("spellAnimationType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimationType-file writer")
		return
	}
	err = fileWriter.CreateSpellAnimationType(spellAnimationType)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellAnimationType-file")
		return
	}
	err = sanitizeSpellAnimationType(spellAnimationType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellAnimationType")
		return
	}
	return
}

//GetSpellAnimationType gets an spellAnimationType by provided spellAnimationTypeID
func GetSpellAnimationType(spellAnimationType *model.SpellAnimationType, user *model.User) (err error) {
	err = prepareSpellAnimationType(spellAnimationType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellAnimationType")
		return
	}

	err = validateSpellAnimationType(spellAnimationType, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellAnimationType")
		return
	}

	reader, err := getReader("spellAnimationType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimationType-memory reader")
		return
	}

	err = reader.GetSpellAnimationType(spellAnimationType)
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimationType")
		return
	}

	err = sanitizeSpellAnimationType(spellAnimationType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellAnimationType")
		return
	}

	return
}

//EditSpellAnimationType edits an existing spellAnimationType
func EditSpellAnimationType(spellAnimationType *model.SpellAnimationType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellAnimationType by search without guide+")
		return
	}
	err = prepareSpellAnimationType(spellAnimationType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellAnimationType")
		return
	}

	err = validateSpellAnimationType(spellAnimationType,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellAnimationType")
		return
	}
	writer, err := getWriter("spellAnimationType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellAnimationType")
		return
	}
	err = writer.EditSpellAnimationType(spellAnimationType)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellAnimationType")
		return
	}

	fileWriter, err := getWriter("spellAnimationType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimationType-file writer")
		return
	}
	err = fileWriter.EditSpellAnimationType(spellAnimationType)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellAnimationType-file")
		return
	}

	err = sanitizeSpellAnimationType(spellAnimationType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellAnimationType")
		return
	}
	return
}

//DeleteSpellAnimationType deletes an spellAnimationType by provided spellAnimationTypeID
func DeleteSpellAnimationType(spellAnimationType *model.SpellAnimationType, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete spellAnimationType without admin+")
		return
	}
	err = prepareSpellAnimationType(spellAnimationType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellAnimationType")
		return
	}

	err = validateSpellAnimationType(spellAnimationType, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellAnimationType")
		return
	}
	writer, err := getWriter("spellAnimationType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimationType-memory writer")
		return
	}
	err = writer.DeleteSpellAnimationType(spellAnimationType)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellAnimationType")
		return
	}

	fileWriter, err := getWriter("spellAnimationType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimationType-file writer")
		return
	}
	err = fileWriter.DeleteSpellAnimationType(spellAnimationType)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellAnimationType-file")
		return
	}

	return
}

func prepareSpellAnimationType(spellAnimationType *model.SpellAnimationType, user *model.User) (err error) {
	if spellAnimationType == nil {
		err = fmt.Errorf("empty spellAnimationType")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSpellAnimationType(spellAnimationType *model.SpellAnimationType, required []string, optional []string) (err error) {
	schema, err := newSchemaSpellAnimationType(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spellAnimationType))
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

func validateOrderBySpellAnimationTypeField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "id"
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

func sanitizeSpellAnimationType(spellAnimationType *model.SpellAnimationType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaSpellAnimationType(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySpellAnimationType(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySpellAnimationType(field); err != nil {
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

func getSchemaPropertySpellAnimationType(field string) (prop model.Schema, err error) {
	switch field {
	case "ID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
