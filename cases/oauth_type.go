package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadOauthTypeFromFileToMemory is ran during initialization
func LoadOauthTypeFromFileToMemory() (err error) {

	fr, err := file.New("config", "oauthType.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("oauthType-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize oauthType-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("oauthType-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize oauthType-memory")
		return
	}

	fileReader, err := getReader("oauthType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get oauthType-file reader")
		return
	}

	memWriter, err := getWriter("oauthType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get oauthType-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListOauthTypeTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list oauthType count")
		return
	}
	page.Limit = page.Total

	oauthTypes, err := fileReader.ListOauthType(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list oauthTypes")
		return
	}

	for _, oauthType := range oauthTypes {
		err = memWriter.CreateOauthType(oauthType)
		if err != nil {
			err = errors.Wrap(err, "failed to create oauthType")
			return
		}
	}

	fmt.Printf("%d oauthTypes, ", len(oauthTypes))
	return
}

//ListOauthType lists all oauthTypes accessible by provided user
func ListOauthType(page *model.Page, user *model.User) (oauthTypes []*model.OauthType, err error) {
	err = validateOrderByOauthTypeField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("oauthType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for oauthType")
		return
	}

	page.Total, err = reader.ListOauthTypeTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list oauthType toal count")
		return
	}

	oauthTypes, err = reader.ListOauthType(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list oauthType")
		return
	}
	for i, oauthType := range oauthTypes {
		err = sanitizeOauthType(oauthType, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize oauthType element %d", i)
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

//ListOauthTypeBySearch will request any oauthType matching the pattern of name
func ListOauthTypeBySearch(page *model.Page, oauthType *model.OauthType, user *model.User) (oauthTypes []*model.OauthType, err error) {

	err = validateOrderByOauthTypeField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareOauthType(oauthType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre oauthType")
		return
	}

	err = validateOauthType(oauthType, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate oauthType")
		return
	}
	reader, err := getReader("oauthType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get oauthType-memory reader")
		return
	}

	oauthTypes, err = reader.ListOauthTypeBySearch(page, oauthType)
	if err != nil {
		err = errors.Wrap(err, "failed to list oauthType by search")
		return
	}

	for _, oauthType := range oauthTypes {
		err = sanitizeOauthType(oauthType, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize oauthType")
			return
		}
	}

	err = sanitizeOauthType(oauthType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search oauthType")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateOauthType will create an oauthType using provided information
func CreateOauthType(oauthType *model.OauthType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list oauthType by search without guide+")
		return
	}
	err = prepareOauthType(oauthType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare oauthType")
		return
	}

	err = validateOauthType(oauthType, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate oauthType")
		return
	}
	oauthType.ID = 0
	writer, err := getWriter("oauthType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for oauthType")
		return
	}
	err = writer.CreateOauthType(oauthType)
	if err != nil {
		err = errors.Wrap(err, "failed to create oauthType")
		return
	}

	fileWriter, err := getWriter("oauthType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get oauthType-file writer")
		return
	}
	err = fileWriter.CreateOauthType(oauthType)
	if err != nil {
		err = errors.Wrap(err, "failed to create oauthType-file")
		return
	}
	err = sanitizeOauthType(oauthType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize oauthType")
		return
	}
	return
}

//GetOauthType gets an oauthType by provided oauthTypeID
func GetOauthType(oauthType *model.OauthType, user *model.User) (err error) {
	err = prepareOauthType(oauthType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare oauthType")
		return
	}

	err = validateOauthType(oauthType, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate oauthType")
		return
	}

	reader, err := getReader("oauthType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get oauthType-memory reader")
		return
	}

	err = reader.GetOauthType(oauthType)
	if err != nil {
		err = errors.Wrap(err, "failed to get oauthType")
		return
	}

	err = sanitizeOauthType(oauthType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize oauthType")
		return
	}

	return
}

//EditOauthType edits an existing oauthType
func EditOauthType(oauthType *model.OauthType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list oauthType by search without guide+")
		return
	}
	err = prepareOauthType(oauthType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare oauthType")
		return
	}

	err = validateOauthType(oauthType,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate oauthType")
		return
	}
	writer, err := getWriter("oauthType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for oauthType")
		return
	}
	err = writer.EditOauthType(oauthType)
	if err != nil {
		err = errors.Wrap(err, "failed to edit oauthType")
		return
	}

	fileWriter, err := getWriter("oauthType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get oauthType-file writer")
		return
	}
	err = fileWriter.EditOauthType(oauthType)
	if err != nil {
		err = errors.Wrap(err, "failed to edit oauthType-file")
		return
	}

	err = sanitizeOauthType(oauthType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize oauthType")
		return
	}
	return
}

//DeleteOauthType deletes an oauthType by provided oauthTypeID
func DeleteOauthType(oauthType *model.OauthType, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete oauthType without admin+")
		return
	}
	err = prepareOauthType(oauthType, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare oauthType")
		return
	}

	err = validateOauthType(oauthType, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate oauthType")
		return
	}
	writer, err := getWriter("oauthType-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get oauthType-memory writer")
		return
	}
	err = writer.DeleteOauthType(oauthType)
	if err != nil {
		err = errors.Wrap(err, "failed to delete oauthType")
		return
	}

	fileWriter, err := getWriter("oauthType-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get oauthType-file writer")
		return
	}
	err = fileWriter.DeleteOauthType(oauthType)
	if err != nil {
		err = errors.Wrap(err, "failed to delete oauthType-file")
		return
	}

	return
}

func prepareOauthType(oauthType *model.OauthType, user *model.User) (err error) {
	if oauthType == nil {
		err = fmt.Errorf("empty oauthType")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateOauthType(oauthType *model.OauthType, required []string, optional []string) (err error) {
	schema, err := newSchemaOauthType(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(oauthType))
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

func validateOrderByOauthTypeField(page *model.Page) (err error) {
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

func sanitizeOauthType(oauthType *model.OauthType, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaOauthType(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyOauthType(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyOauthType(field); err != nil {
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

func getSchemaPropertyOauthType(field string) (prop model.Schema, err error) {
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
