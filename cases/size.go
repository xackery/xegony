package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadSizeFromFileToMemory is ran during initialization
func LoadSizeFromFileToMemory() (err error) {

	fr, err := file.New("config", "size.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("size-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize size-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("size-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize size-memory")
		return
	}

	fileReader, err := getReader("size-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get size-file reader")
		return
	}

	memWriter, err := getWriter("size-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get size-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListSizeTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list size count")
		return
	}
	page.Limit = page.Total

	sizes, err := fileReader.ListSize(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list sizes")
		return
	}

	for _, size := range sizes {
		err = memWriter.CreateSize(size)
		if err != nil {
			err = errors.Wrap(err, "failed to create size")
			return
		}
	}

	fmt.Printf("%d sizees, ", len(sizes))
	return
}

//ListSize lists all sizes accessible by provided user
func ListSize(page *model.Page, user *model.User) (sizes []*model.Size, err error) {
	err = validateOrderBySizeField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("size-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for size")
		return
	}

	page.Total, err = reader.ListSizeTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list size toal count")
		return
	}

	sizes, err = reader.ListSize(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list size")
		return
	}
	for i, size := range sizes {
		err = sanitizeSize(size, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize size element %d", i)
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

//ListSizeBySearch will request any size matching the pattern of name
func ListSizeBySearch(page *model.Page, size *model.Size, user *model.User) (sizes []*model.Size, err error) {

	err = validateOrderBySizeField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSize(size, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre size")
		return
	}

	err = validateSize(size, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate size")
		return
	}
	reader, err := getReader("size-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get size-memory reader")
		return
	}

	sizes, err = reader.ListSizeBySearch(page, size)
	if err != nil {
		err = errors.Wrap(err, "failed to list size by search")
		return
	}

	for _, size := range sizes {
		err = sanitizeSize(size, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize size")
			return
		}
	}

	err = sanitizeSize(size, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search size")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSize will create an size using provided information
func CreateSize(size *model.Size, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list size by search without guide+")
		return
	}
	err = prepareSize(size, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare size")
		return
	}

	err = validateSize(size, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate size")
		return
	}
	size.ID = 0
	writer, err := getWriter("size-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for size")
		return
	}
	err = writer.CreateSize(size)
	if err != nil {
		err = errors.Wrap(err, "failed to create size")
		return
	}

	fileWriter, err := getWriter("size-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get size-file writer")
		return
	}
	err = fileWriter.CreateSize(size)
	if err != nil {
		err = errors.Wrap(err, "failed to create size-file")
		return
	}
	err = sanitizeSize(size, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize size")
		return
	}
	return
}

//GetSize gets an size by provided sizeID
func GetSize(size *model.Size, user *model.User) (err error) {
	err = prepareSize(size, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare size")
		return
	}

	err = validateSize(size, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate size")
		return
	}

	reader, err := getReader("size-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get size-memory reader")
		return
	}

	err = reader.GetSize(size)
	if err != nil {
		err = errors.Wrap(err, "failed to get size")
		return
	}

	err = sanitizeSize(size, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize size")
		return
	}

	return
}

//EditSize edits an existing size
func EditSize(size *model.Size, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list size by search without guide+")
		return
	}
	err = prepareSize(size, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare size")
		return
	}

	err = validateSize(size,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"male",
			"female",
			"neutral",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate size")
		return
	}
	writer, err := getWriter("size-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for size")
		return
	}
	err = writer.EditSize(size)
	if err != nil {
		err = errors.Wrap(err, "failed to edit size")
		return
	}

	fileWriter, err := getWriter("size-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get size-file writer")
		return
	}
	err = fileWriter.EditSize(size)
	if err != nil {
		err = errors.Wrap(err, "failed to edit size-file")
		return
	}

	err = sanitizeSize(size, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize size")
		return
	}
	return
}

//DeleteSize deletes an size by provided sizeID
func DeleteSize(size *model.Size, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete size without admin+")
		return
	}
	err = prepareSize(size, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare size")
		return
	}

	err = validateSize(size, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate size")
		return
	}
	writer, err := getWriter("size-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get size-memory writer")
		return
	}
	err = writer.DeleteSize(size)
	if err != nil {
		err = errors.Wrap(err, "failed to delete size")
		return
	}

	fileWriter, err := getWriter("size-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get size-file writer")
		return
	}
	err = fileWriter.DeleteSize(size)
	if err != nil {
		err = errors.Wrap(err, "failed to delete size-file")
		return
	}

	return
}

func prepareSize(size *model.Size, user *model.User) (err error) {
	if size == nil {
		err = fmt.Errorf("empty size")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSize(size *model.Size, required []string, optional []string) (err error) {
	schema, err := newSchemaSize(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(size))
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

func validateOrderBySizeField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "name"
	}

	validNames := []string{
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

func sanitizeSize(size *model.Size, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaSize(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySize(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySize(field); err != nil {
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

func getSchemaPropertySize(field string) (prop model.Schema, err error) {
	switch field {
	case "ID":
		prop.Type = "integer"
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "shortName":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
