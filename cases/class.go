package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadClassFromFileToMemory is ran during initialization
func LoadClassFromFileToMemory() (err error) {

	fr, err := file.New("config", "class.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("class-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize class-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("class-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize class-memory")
		return
	}

	fileReader, err := getReader("class-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get class-file reader")
		return
	}

	memWriter, err := getWriter("class-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get class-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListClassTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list class count")
		return
	}
	page.Limit = page.Total

	classs, err := fileReader.ListClass(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list classs")
		return
	}

	for _, class := range classs {
		err = memWriter.CreateClass(class)
		if err != nil {
			err = errors.Wrap(err, "failed to create class")
			return
		}
	}

	fmt.Printf("%d classes, ", len(classs))
	return
}

//ListClass lists all classs accessible by provided user
func ListClass(page *model.Page, user *model.User) (classs []*model.Class, err error) {
	err = validateOrderByClassField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("class-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for class")
		return
	}

	page.Total, err = reader.ListClassTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list class toal count")
		return
	}

	classs, err = reader.ListClass(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list class")
		return
	}
	for i, class := range classs {
		err = sanitizeClass(class, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize class element %d", i)
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

//ListClassBySearch will request any class matching the pattern of name
func ListClassBySearch(page *model.Page, class *model.Class, user *model.User) (classs []*model.Class, err error) {

	err = validateOrderByClassField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareClass(class, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre class")
		return
	}

	err = validateClass(class, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate class")
		return
	}
	reader, err := getReader("class-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get class-memory reader")
		return
	}

	classs, err = reader.ListClassBySearch(page, class)
	if err != nil {
		err = errors.Wrap(err, "failed to list class by search")
		return
	}

	for _, class := range classs {
		err = sanitizeClass(class, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize class")
			return
		}
	}

	err = sanitizeClass(class, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search class")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//ListClassByBit lists all classs that match a bitmask
func ListClassByBit(page *model.Page, class *model.Class, user *model.User) (classs []*model.Class, err error) {
	err = validateOrderByClassField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("class-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for class")
		return
	}

	page.Total, err = reader.ListClassByBitTotalCount(class)
	if err != nil {
		err = errors.Wrap(err, "failed to list class toal count")
		return
	}

	classs, err = reader.ListClassByBit(page, class)
	if err != nil {
		err = errors.Wrap(err, "failed to list class")
		return
	}

	for i, class := range classs {
		err = sanitizeClass(class, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize class element %d", i)
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

//CreateClass will create an class using provided information
func CreateClass(class *model.Class, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list class by search without guide+")
		return
	}
	err = prepareClass(class, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare class")
		return
	}

	err = validateClass(class, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate class")
		return
	}
	class.ID = 0
	writer, err := getWriter("class-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for class")
		return
	}
	err = writer.CreateClass(class)
	if err != nil {
		err = errors.Wrap(err, "failed to create class")
		return
	}

	fileWriter, err := getWriter("class-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get class-file writer")
		return
	}
	err = fileWriter.CreateClass(class)
	if err != nil {
		err = errors.Wrap(err, "failed to create class-file")
		return
	}
	err = sanitizeClass(class, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize class")
		return
	}
	return
}

//GetClass gets an class by provided classID
func GetClass(class *model.Class, user *model.User) (err error) {
	err = prepareClass(class, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare class")
		return
	}

	err = validateClass(class, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate class")
		return
	}

	reader, err := getReader("class-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get class-memory reader")
		return
	}

	err = reader.GetClass(class)
	if err != nil {
		err = errors.Wrap(err, "failed to get class")
		return
	}

	err = sanitizeClass(class, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize class")
		return
	}

	return
}

//EditClass edits an existing class
func EditClass(class *model.Class, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list class by search without guide+")
		return
	}
	err = prepareClass(class, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare class")
		return
	}

	err = validateClass(class,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"male",
			"female",
			"neutral",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate class")
		return
	}
	writer, err := getWriter("class-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for class")
		return
	}
	err = writer.EditClass(class)
	if err != nil {
		err = errors.Wrap(err, "failed to edit class")
		return
	}

	fileWriter, err := getWriter("class-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get class-file writer")
		return
	}
	err = fileWriter.EditClass(class)
	if err != nil {
		err = errors.Wrap(err, "failed to edit class-file")
		return
	}

	err = sanitizeClass(class, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize class")
		return
	}
	return
}

//DeleteClass deletes an class by provided classID
func DeleteClass(class *model.Class, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete class without admin+")
		return
	}
	err = prepareClass(class, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare class")
		return
	}

	err = validateClass(class, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate class")
		return
	}
	writer, err := getWriter("class-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get class-memory writer")
		return
	}
	err = writer.DeleteClass(class)
	if err != nil {
		err = errors.Wrap(err, "failed to delete class")
		return
	}

	fileWriter, err := getWriter("class-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get class-file writer")
		return
	}
	err = fileWriter.DeleteClass(class)
	if err != nil {
		err = errors.Wrap(err, "failed to delete class-file")
		return
	}

	return
}

func prepareClass(class *model.Class, user *model.User) (err error) {
	if class == nil {
		err = fmt.Errorf("empty class")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateClass(class *model.Class, required []string, optional []string) (err error) {
	schema, err := newSchemaClass(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(class))
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

func validateOrderByClassField(page *model.Page) (err error) {
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

func sanitizeClass(class *model.Class, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	if class.Icon == "" {
		class.Icon = "xa-shield"
	}
	if len(class.Icon) > 2 && class.Icon[2:3] == "-" {
		class.Icon = fmt.Sprintf("%s %s", class.Icon[0:2], class.Icon)
	}
	return
}

func newSchemaClass(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyClass(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyClass(field); err != nil {
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

func getSchemaPropertyClass(field string) (prop model.Schema, err error) {
	switch field {
	case "ID":
		prop.Type = "integer"
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

func getClassListByBit(bit int64) string {
	classes := ""
	if bit == 65535 {
		return "ALL"
	}
	if bit&1 == 1 {
		classes += "WAR "
	}
	if bit&2 == 2 {
		classes += "CLR "
	}
	if bit&4 == 4 {
		classes += "PAL "
	}
	if bit&8 == 8 {
		classes += "RNG "
	}
	if bit&16 == 16 {
		classes += "SHM "
	}
	if bit&32 == 32 {
		classes += "DRU "
	}
	if bit&64 == 64 {
		classes += "MNK "
	}
	if bit&128 == 128 {
		classes += "BRD "
	}
	if bit&256 == 256 {
		classes += "ROG "
	}
	if bit&512 == 512 {
		classes += "SHD "
	}
	if bit&1024 == 1024 {
		classes += "NEC "
	}
	if bit&2048 == 2048 {
		classes += "WIZ "
	}
	if bit&4096 == 4096 {
		classes += "MAG "
	}
	if bit&8192 == 8192 {
		classes += "ENC "
	}
	if bit&16384 == 16384 {
		classes += "BST "
	}
	if bit&32768 == 32768 {
		classes += "BER "
	}
	if len(classes) > 0 {
		classes = classes[0 : len(classes)-1]
	}
	if len(classes) == 0 {
		classes = "NONE"
	}
	return classes
}
