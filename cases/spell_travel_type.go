package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadSpellTravelTypeFromFileToMemory is ran during initialization
func LoadSpellTravelTypeFromFileToMemory() (err error) {

	fr, err := file.New("config", "spellTravelType.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("spellTravelType-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellTravelType-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("spellTravelType-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellTravelType-memory")
		return
	}

	fileReader, err := getReader("spellTravelType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTravelType-file reader")
		return
	}

	memWriter, err := getWriter("spellTravelType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTravelType-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListSpellTravelTypeTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list spellTravelType count")
		return
	}
	page.Limit = page.Total

	spellTravelTypes, err := fileReader.ListSpellTravelType(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellTravelTypes")
		return
	}

	for _, spellTravelType := range spellTravelTypes {
		err = memWriter.CreateSpellTravelType(spellTravelType)
		if err != nil {
			err = errors.Wrap(err, "failed to create spellTravelType")
			return
		}
	}

	fmt.Printf("%d spellTravelTypes, ", len(spellTravelTypes))
	return
}

//ListSpellTravelType lists all spellTravelTypes accessible by provided user
func ListSpellTravelType(page *model.Page, user *model.User) (spellTravelTypes []*model.SpellTravelType, err error) {
	err = validateOrderBySpellTravelTypeField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("spellTravelType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for spellTravelType")
		return
	}

	page.Total, err = reader.ListSpellTravelTypeTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list spellTravelType toal count")
		return
	}

	spellTravelTypes, err = reader.ListSpellTravelType(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellTravelType")
		return
	}
	for i, spellTravelType := range spellTravelTypes {
		err = sanitizeSpellTravelType(spellTravelType, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize spellTravelType element %d", i)
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

//ListSpellTravelTypeBySearch will request any spellTravelType matching the pattern of name
func ListSpellTravelTypeBySearch(page *model.Page, spellTravelType *model.SpellTravelType, user *model.User) (spellTravelTypes []*model.SpellTravelType, err error) {

	err = validateOrderBySpellTravelTypeField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSpellTravelType(spellTravelType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre spellTravelType")
		return
	}

	err = validateSpellTravelType(spellTravelType, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellTravelType")
		return
	}
	reader, err := getReader("spellTravelType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTravelType-memory reader")
		return
	}

	spellTravelTypes, err = reader.ListSpellTravelTypeBySearch(page, spellTravelType)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellTravelType by search")
		return
	}

	for _, spellTravelType := range spellTravelTypes {
		err = sanitizeSpellTravelType(spellTravelType, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize spellTravelType")
			return
		}
	}

	err = sanitizeSpellTravelType(spellTravelType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search spellTravelType")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSpellTravelType will create an spellTravelType using provided information
func CreateSpellTravelType(spellTravelType *model.SpellTravelType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellTravelType by search without guide+")
		return
	}
	err = prepareSpellTravelType(spellTravelType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellTravelType")
		return
	}

	err = validateSpellTravelType(spellTravelType, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellTravelType")
		return
	}
	spellTravelType.ID = 0
	writer, err := getWriter("spellTravelType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellTravelType")
		return
	}
	err = writer.CreateSpellTravelType(spellTravelType)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellTravelType")
		return
	}

	fileWriter, err := getWriter("spellTravelType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTravelType-file writer")
		return
	}
	err = fileWriter.CreateSpellTravelType(spellTravelType)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellTravelType-file")
		return
	}
	err = sanitizeSpellTravelType(spellTravelType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellTravelType")
		return
	}
	return
}

//GetSpellTravelType gets an spellTravelType by provided spellTravelTypeID
func GetSpellTravelType(spellTravelType *model.SpellTravelType, user *model.User) (err error) {
	err = prepareSpellTravelType(spellTravelType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellTravelType")
		return
	}

	err = validateSpellTravelType(spellTravelType, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellTravelType")
		return
	}

	reader, err := getReader("spellTravelType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTravelType-memory reader")
		return
	}

	err = reader.GetSpellTravelType(spellTravelType)
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTravelType")
		return
	}

	err = sanitizeSpellTravelType(spellTravelType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellTravelType")
		return
	}

	return
}

//EditSpellTravelType edits an existing spellTravelType
func EditSpellTravelType(spellTravelType *model.SpellTravelType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellTravelType by search without guide+")
		return
	}
	err = prepareSpellTravelType(spellTravelType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellTravelType")
		return
	}

	err = validateSpellTravelType(spellTravelType,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellTravelType")
		return
	}
	writer, err := getWriter("spellTravelType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellTravelType")
		return
	}
	err = writer.EditSpellTravelType(spellTravelType)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellTravelType")
		return
	}

	fileWriter, err := getWriter("spellTravelType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTravelType-file writer")
		return
	}
	err = fileWriter.EditSpellTravelType(spellTravelType)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellTravelType-file")
		return
	}

	err = sanitizeSpellTravelType(spellTravelType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellTravelType")
		return
	}
	return
}

//DeleteSpellTravelType deletes an spellTravelType by provided spellTravelTypeID
func DeleteSpellTravelType(spellTravelType *model.SpellTravelType, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete spellTravelType without admin+")
		return
	}
	err = prepareSpellTravelType(spellTravelType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellTravelType")
		return
	}

	err = validateSpellTravelType(spellTravelType, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellTravelType")
		return
	}
	writer, err := getWriter("spellTravelType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTravelType-memory writer")
		return
	}
	err = writer.DeleteSpellTravelType(spellTravelType)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellTravelType")
		return
	}

	fileWriter, err := getWriter("spellTravelType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellTravelType-file writer")
		return
	}
	err = fileWriter.DeleteSpellTravelType(spellTravelType)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellTravelType-file")
		return
	}

	return
}

func prepareSpellTravelType(spellTravelType *model.SpellTravelType, user *model.User) (err error) {
	if spellTravelType == nil {
		err = fmt.Errorf("empty spellTravelType")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSpellTravelType(spellTravelType *model.SpellTravelType, required []string, optional []string) (err error) {
	schema, err := newSchemaSpellTravelType(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spellTravelType))
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

func validateOrderBySpellTravelTypeField(page *model.Page) (err error) {
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

func sanitizeSpellTravelType(spellTravelType *model.SpellTravelType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaSpellTravelType(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySpellTravelType(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySpellTravelType(field); err != nil {
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

func getSchemaPropertySpellTravelType(field string) (prop model.Schema, err error) {
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
