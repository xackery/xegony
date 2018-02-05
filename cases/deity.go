package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadDeityFromFileToMemory is ran during initialization
func LoadDeityFromFileToMemory() (err error) {

	fr, err := file.New("config", "deity.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("deity-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize deity-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("deity-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize deity-memory")
		return
	}

	fileReader, err := getReader("deity-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get deity-file reader")
		return
	}

	memWriter, err := getWriter("deity-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get deity-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListDeityTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list deity count")
		return
	}
	page.Limit = page.Total

	deitys, err := fileReader.ListDeity(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list deitys")
		return
	}

	for _, deity := range deitys {
		err = memWriter.CreateDeity(deity)
		if err != nil {
			err = errors.Wrap(err, "failed to create deity")
			return
		}
	}

	fmt.Printf("%d deityes, ", len(deitys))
	return
}

//ListDeity lists all deitys accessible by provided user
func ListDeity(page *model.Page, user *model.User) (deitys []*model.Deity, err error) {
	err = validateOrderByDeityField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("deity-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for deity")
		return
	}

	page.Total, err = reader.ListDeityTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list deity toal count")
		return
	}

	deitys, err = reader.ListDeity(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list deity")
		return
	}
	for i, deity := range deitys {
		err = sanitizeDeity(deity, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize deity element %d", i)
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

//ListDeityBySearch will request any deity matching the pattern of name
func ListDeityBySearch(page *model.Page, deity *model.Deity, user *model.User) (deitys []*model.Deity, err error) {

	err = validateOrderByDeityField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareDeity(deity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre deity")
		return
	}

	err = validateDeity(deity, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate deity")
		return
	}
	reader, err := getReader("deity-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get deity-memory reader")
		return
	}

	deitys, err = reader.ListDeityBySearch(page, deity)
	if err != nil {
		err = errors.Wrap(err, "failed to list deity by search")
		return
	}

	for _, deity := range deitys {
		err = sanitizeDeity(deity, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize deity")
			return
		}
	}

	err = sanitizeDeity(deity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search deity")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//ListDeityByBit lists all deitys that match a bitmask
func ListDeityByBit(page *model.Page, deity *model.Deity, user *model.User) (deitys []*model.Deity, err error) {
	err = validateOrderByDeityField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("deity-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for deity")
		return
	}

	page.Total, err = reader.ListDeityByBitTotalCount(deity)
	if err != nil {
		err = errors.Wrap(err, "failed to list deity toal count")
		return
	}

	deitys, err = reader.ListDeityByBit(page, deity)
	if err != nil {
		err = errors.Wrap(err, "failed to list deity")
		return
	}

	for i, deity := range deitys {
		err = sanitizeDeity(deity, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize deity element %d", i)
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

//CreateDeity will create an deity using provided information
func CreateDeity(deity *model.Deity, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list deity by search without guide+")
		return
	}
	err = prepareDeity(deity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare deity")
		return
	}

	err = validateDeity(deity, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate deity")
		return
	}
	deity.ID = 0
	writer, err := getWriter("deity-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for deity")
		return
	}
	err = writer.CreateDeity(deity)
	if err != nil {
		err = errors.Wrap(err, "failed to create deity")
		return
	}

	fileWriter, err := getWriter("deity-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get deity-file writer")
		return
	}
	err = fileWriter.CreateDeity(deity)
	if err != nil {
		err = errors.Wrap(err, "failed to create deity-file")
		return
	}
	err = sanitizeDeity(deity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize deity")
		return
	}
	return
}

//GetDeityBySpell gets an deity by provided deityID
func GetDeityBySpell(spell *model.Spell, deity *model.Deity, user *model.User) (err error) {
	err = prepareDeity(deity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare deity")
		return
	}

	reader, err := getReader("deity-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get deity-memory reader")
		return
	}

	err = reader.GetDeityBySpell(spell, deity)
	if err != nil {
		err = errors.Wrap(err, "failed to get deity")
		return
	}

	err = sanitizeDeity(deity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize deity")
		return
	}

	return
}

//GetDeity gets an deity by provided deityID
func GetDeity(deity *model.Deity, user *model.User) (err error) {
	err = prepareDeity(deity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare deity")
		return
	}

	err = validateDeity(deity, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate deity")
		return
	}

	reader, err := getReader("deity-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get deity-memory reader")
		return
	}

	err = reader.GetDeity(deity)
	if err != nil {
		err = errors.Wrap(err, "failed to get deity")
		return
	}

	err = sanitizeDeity(deity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize deity")
		return
	}

	return
}

//EditDeity edits an existing deity
func EditDeity(deity *model.Deity, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list deity by search without guide+")
		return
	}
	err = prepareDeity(deity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare deity")
		return
	}

	err = validateDeity(deity,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"male",
			"female",
			"neutral",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate deity")
		return
	}
	writer, err := getWriter("deity-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for deity")
		return
	}
	err = writer.EditDeity(deity)
	if err != nil {
		err = errors.Wrap(err, "failed to edit deity")
		return
	}

	fileWriter, err := getWriter("deity-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get deity-file writer")
		return
	}
	err = fileWriter.EditDeity(deity)
	if err != nil {
		err = errors.Wrap(err, "failed to edit deity-file")
		return
	}

	err = sanitizeDeity(deity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize deity")
		return
	}
	return
}

//DeleteDeity deletes an deity by provided deityID
func DeleteDeity(deity *model.Deity, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete deity without admin+")
		return
	}
	err = prepareDeity(deity, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare deity")
		return
	}

	err = validateDeity(deity, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate deity")
		return
	}
	writer, err := getWriter("deity-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get deity-memory writer")
		return
	}
	err = writer.DeleteDeity(deity)
	if err != nil {
		err = errors.Wrap(err, "failed to delete deity")
		return
	}

	fileWriter, err := getWriter("deity-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get deity-file writer")
		return
	}
	err = fileWriter.DeleteDeity(deity)
	if err != nil {
		err = errors.Wrap(err, "failed to delete deity-file")
		return
	}

	return
}

func prepareDeity(deity *model.Deity, user *model.User) (err error) {
	if deity == nil {
		err = fmt.Errorf("empty deity")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateDeity(deity *model.Deity, required []string, optional []string) (err error) {
	schema, err := newSchemaDeity(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(deity))
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

func validateOrderByDeityField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "id"
	}

	validNames := []string{
		"id",
		"name",
		"bit",
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

func sanitizeDeity(deity *model.Deity, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaDeity(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyDeity(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyDeity(field); err != nil {
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

func getSchemaPropertyDeity(field string) (prop model.Schema, err error) {
	switch field {
	case "ID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "male":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "female":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "neutral":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "icon":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
