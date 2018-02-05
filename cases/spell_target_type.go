package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadSpellTargetTypeFromFileToMemory is ran during initialization
func LoadSpellTargetTypeFromFileToMemory() (err error) {

	fr, err := file.New("config", "spellTargetType.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("spellTargetType-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellTargetType-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("spellTargetType-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellTargetType-memory")
		return
	}

	fileReader, err := getReader("spellTargetType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTargetType-file reader")
		return
	}

	memWriter, err := getWriter("spellTargetType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTargetType-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListSpellTargetTypeTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list spellTargetType count")
		return
	}
	page.Limit = page.Total

	spellTargetTypes, err := fileReader.ListSpellTargetType(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellTargetTypes")
		return
	}

	for _, spellTargetType := range spellTargetTypes {
		err = memWriter.CreateSpellTargetType(spellTargetType)
		if err != nil {
			err = errors.Wrap(err, "failed to create spellTargetType")
			return
		}
	}

	fmt.Printf("%d spellTargetTypes, ", len(spellTargetTypes))
	return
}

//ListSpellTargetType lists all spellTargetTypes accessible by provided user
func ListSpellTargetType(page *model.Page, user *model.User) (spellTargetTypes []*model.SpellTargetType, err error) {
	err = validateOrderBySpellTargetTypeField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("spellTargetType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for spellTargetType")
		return
	}

	page.Total, err = reader.ListSpellTargetTypeTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list spellTargetType toal count")
		return
	}

	spellTargetTypes, err = reader.ListSpellTargetType(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellTargetType")
		return
	}
	for i, spellTargetType := range spellTargetTypes {
		err = sanitizeSpellTargetType(spellTargetType, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize spellTargetType element %d", i)
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

//ListSpellTargetTypeBySearch will request any spellTargetType matching the pattern of name
func ListSpellTargetTypeBySearch(page *model.Page, spellTargetType *model.SpellTargetType, user *model.User) (spellTargetTypes []*model.SpellTargetType, err error) {

	err = validateOrderBySpellTargetTypeField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSpellTargetType(spellTargetType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre spellTargetType")
		return
	}

	err = validateSpellTargetType(spellTargetType, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellTargetType")
		return
	}
	reader, err := getReader("spellTargetType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTargetType-memory reader")
		return
	}

	spellTargetTypes, err = reader.ListSpellTargetTypeBySearch(page, spellTargetType)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellTargetType by search")
		return
	}

	for _, spellTargetType := range spellTargetTypes {
		err = sanitizeSpellTargetType(spellTargetType, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize spellTargetType")
			return
		}
	}

	err = sanitizeSpellTargetType(spellTargetType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search spellTargetType")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSpellTargetType will create an spellTargetType using provided information
func CreateSpellTargetType(spellTargetType *model.SpellTargetType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellTargetType by search without guide+")
		return
	}
	err = prepareSpellTargetType(spellTargetType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellTargetType")
		return
	}

	err = validateSpellTargetType(spellTargetType, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellTargetType")
		return
	}
	spellTargetType.ID = 0
	writer, err := getWriter("spellTargetType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellTargetType")
		return
	}
	err = writer.CreateSpellTargetType(spellTargetType)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellTargetType")
		return
	}

	fileWriter, err := getWriter("spellTargetType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTargetType-file writer")
		return
	}
	err = fileWriter.CreateSpellTargetType(spellTargetType)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellTargetType-file")
		return
	}
	err = sanitizeSpellTargetType(spellTargetType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellTargetType")
		return
	}
	return
}

//GetSpellTargetType gets an spellTargetType by provided spellTargetTypeID
func GetSpellTargetType(spellTargetType *model.SpellTargetType, user *model.User) (err error) {
	err = prepareSpellTargetType(spellTargetType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellTargetType")
		return
	}

	err = validateSpellTargetType(spellTargetType, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellTargetType")
		return
	}

	reader, err := getReader("spellTargetType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTargetType-memory reader")
		return
	}

	err = reader.GetSpellTargetType(spellTargetType)
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTargetType")
		return
	}

	err = sanitizeSpellTargetType(spellTargetType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellTargetType")
		return
	}

	return
}

//EditSpellTargetType edits an existing spellTargetType
func EditSpellTargetType(spellTargetType *model.SpellTargetType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellTargetType by search without guide+")
		return
	}
	err = prepareSpellTargetType(spellTargetType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellTargetType")
		return
	}

	err = validateSpellTargetType(spellTargetType,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellTargetType")
		return
	}
	writer, err := getWriter("spellTargetType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellTargetType")
		return
	}
	err = writer.EditSpellTargetType(spellTargetType)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellTargetType")
		return
	}

	fileWriter, err := getWriter("spellTargetType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTargetType-file writer")
		return
	}
	err = fileWriter.EditSpellTargetType(spellTargetType)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellTargetType-file")
		return
	}

	err = sanitizeSpellTargetType(spellTargetType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellTargetType")
		return
	}
	return
}

//DeleteSpellTargetType deletes an spellTargetType by provided spellTargetTypeID
func DeleteSpellTargetType(spellTargetType *model.SpellTargetType, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete spellTargetType without admin+")
		return
	}
	err = prepareSpellTargetType(spellTargetType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellTargetType")
		return
	}

	err = validateSpellTargetType(spellTargetType, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellTargetType")
		return
	}
	writer, err := getWriter("spellTargetType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTargetType-memory writer")
		return
	}
	err = writer.DeleteSpellTargetType(spellTargetType)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellTargetType")
		return
	}

	fileWriter, err := getWriter("spellTargetType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTargetType-file writer")
		return
	}
	err = fileWriter.DeleteSpellTargetType(spellTargetType)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellTargetType-file")
		return
	}

	return
}

func prepareSpellTargetType(spellTargetType *model.SpellTargetType, user *model.User) (err error) {
	if spellTargetType == nil {
		err = fmt.Errorf("empty spellTargetType")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSpellTargetType(spellTargetType *model.SpellTargetType, required []string, optional []string) (err error) {
	schema, err := newSchemaSpellTargetType(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spellTargetType))
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

func validateOrderBySpellTargetTypeField(page *model.Page) (err error) {
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

func sanitizeSpellTargetType(spellTargetType *model.SpellTargetType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaSpellTargetType(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySpellTargetType(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySpellTargetType(field); err != nil {
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

func getSchemaPropertySpellTargetType(field string) (prop model.Schema, err error) {
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
