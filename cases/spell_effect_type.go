package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadSpellEffectTypeFromFileToMemory is ran during initialization
func LoadSpellEffectTypeFromFileToMemory() (err error) {

	fr, err := file.New("config", "spellEffectType.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("spellEffectType-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellEffectType-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("spellEffectType-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellEffectType-memory")
		return
	}

	fileReader, err := getReader("spellEffectType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectType-file reader")
		return
	}

	memWriter, err := getWriter("spellEffectType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectType-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListSpellEffectTypeTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list spellEffectType count")
		return
	}
	page.Limit = page.Total

	spellEffectTypes, err := fileReader.ListSpellEffectType(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellEffectTypes")
		return
	}

	for _, spellEffectType := range spellEffectTypes {
		err = memWriter.CreateSpellEffectType(spellEffectType)
		if err != nil {
			err = errors.Wrap(err, "failed to create spellEffectType")
			return
		}
	}

	fmt.Printf("%d spellEffectTypes, ", len(spellEffectTypes))
	return
}

//ListSpellEffectType lists all spellEffectTypes accessible by provided user
func ListSpellEffectType(page *model.Page, user *model.User) (spellEffectTypes []*model.SpellEffectType, err error) {
	err = validateOrderBySpellEffectTypeField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("spellEffectType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for spellEffectType")
		return
	}

	page.Total, err = reader.ListSpellEffectTypeTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list spellEffectType toal count")
		return
	}

	spellEffectTypes, err = reader.ListSpellEffectType(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellEffectType")
		return
	}
	for i, spellEffectType := range spellEffectTypes {
		err = sanitizeSpellEffectType(spellEffectType, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize spellEffectType element %d", i)
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

//ListSpellEffectTypeBySearch will request any spellEffectType matching the pattern of name
func ListSpellEffectTypeBySearch(page *model.Page, spellEffectType *model.SpellEffectType, user *model.User) (spellEffectTypes []*model.SpellEffectType, err error) {

	err = validateOrderBySpellEffectTypeField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSpellEffectType(spellEffectType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre spellEffectType")
		return
	}

	err = validateSpellEffectType(spellEffectType, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellEffectType")
		return
	}
	reader, err := getReader("spellEffectType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectType-memory reader")
		return
	}

	spellEffectTypes, err = reader.ListSpellEffectTypeBySearch(page, spellEffectType)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellEffectType by search")
		return
	}

	for _, spellEffectType := range spellEffectTypes {
		err = sanitizeSpellEffectType(spellEffectType, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize spellEffectType")
			return
		}
	}

	err = sanitizeSpellEffectType(spellEffectType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search spellEffectType")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSpellEffectType will create an spellEffectType using provided information
func CreateSpellEffectType(spellEffectType *model.SpellEffectType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellEffectType by search without guide+")
		return
	}
	err = prepareSpellEffectType(spellEffectType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellEffectType")
		return
	}

	err = validateSpellEffectType(spellEffectType, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellEffectType")
		return
	}
	spellEffectType.ID = 0
	writer, err := getWriter("spellEffectType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellEffectType")
		return
	}
	err = writer.CreateSpellEffectType(spellEffectType)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellEffectType")
		return
	}

	fileWriter, err := getWriter("spellEffectType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectType-file writer")
		return
	}
	err = fileWriter.CreateSpellEffectType(spellEffectType)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellEffectType-file")
		return
	}
	err = sanitizeSpellEffectType(spellEffectType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellEffectType")
		return
	}
	return
}

//GetSpellEffectType gets an spellEffectType by provided spellEffectTypeID
func GetSpellEffectType(spellEffectType *model.SpellEffectType, user *model.User) (err error) {
	err = prepareSpellEffectType(spellEffectType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellEffectType")
		return
	}

	err = validateSpellEffectType(spellEffectType, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellEffectType")
		return
	}

	reader, err := getReader("spellEffectType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectType-memory reader")
		return
	}

	err = reader.GetSpellEffectType(spellEffectType)
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectType")
		return
	}

	err = sanitizeSpellEffectType(spellEffectType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellEffectType")
		return
	}

	return
}

//EditSpellEffectType edits an existing spellEffectType
func EditSpellEffectType(spellEffectType *model.SpellEffectType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellEffectType by search without guide+")
		return
	}
	err = prepareSpellEffectType(spellEffectType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellEffectType")
		return
	}

	err = validateSpellEffectType(spellEffectType,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellEffectType")
		return
	}
	writer, err := getWriter("spellEffectType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellEffectType")
		return
	}
	err = writer.EditSpellEffectType(spellEffectType)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellEffectType")
		return
	}

	fileWriter, err := getWriter("spellEffectType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectType-file writer")
		return
	}
	err = fileWriter.EditSpellEffectType(spellEffectType)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellEffectType-file")
		return
	}

	err = sanitizeSpellEffectType(spellEffectType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellEffectType")
		return
	}
	return
}

//DeleteSpellEffectType deletes an spellEffectType by provided spellEffectTypeID
func DeleteSpellEffectType(spellEffectType *model.SpellEffectType, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete spellEffectType without admin+")
		return
	}
	err = prepareSpellEffectType(spellEffectType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellEffectType")
		return
	}

	err = validateSpellEffectType(spellEffectType, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellEffectType")
		return
	}
	writer, err := getWriter("spellEffectType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectType-memory writer")
		return
	}
	err = writer.DeleteSpellEffectType(spellEffectType)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellEffectType")
		return
	}

	fileWriter, err := getWriter("spellEffectType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectType-file writer")
		return
	}
	err = fileWriter.DeleteSpellEffectType(spellEffectType)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellEffectType-file")
		return
	}

	return
}

func prepareSpellEffectType(spellEffectType *model.SpellEffectType, user *model.User) (err error) {
	if spellEffectType == nil {
		err = fmt.Errorf("empty spellEffectType")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSpellEffectType(spellEffectType *model.SpellEffectType, required []string, optional []string) (err error) {
	schema, err := newSchemaSpellEffectType(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spellEffectType))
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

func validateOrderBySpellEffectTypeField(page *model.Page) (err error) {
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

func sanitizeSpellEffectType(spellEffectType *model.SpellEffectType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaSpellEffectType(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySpellEffectType(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySpellEffectType(field); err != nil {
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

func getSchemaPropertySpellEffectType(field string) (prop model.Schema, err error) {
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
