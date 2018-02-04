package cases

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadVariableFromDBToMemory is ran during initialization
func LoadVariableFromDBToMemory() (err error) {

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("variable-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize variable-memory")
		return
	}

	dbReader, err := getReader("variable")
	if err != nil {
		err = errors.Wrap(err, "failed to get variable reader")
		return
	}

	memWriter, err := getWriter("variable-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get variable-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = dbReader.ListVariableTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list variable count")
		return
	}
	var totalVariables []*model.Variable
	var variables []*model.Variable
	for {
		variables, err = dbReader.ListVariable(page)
		if err != nil {
			err = errors.Wrap(err, "failed to list variables")
			return
		}
		totalVariables = append(totalVariables, variables...)
		if int64(len(totalVariables)) >= page.Total {
			break
		}
		page.Offset++
	}

	for _, variable := range totalVariables {
		err = memWriter.CreateVariable(variable)
		if err != nil {
			err = errors.Wrap(err, "failed to create variable")
			return
		}
	}

	fmt.Printf("%d variables, ", len(totalVariables))
	return
}

//ListVariable lists all variables accessible by provided user
func ListVariable(page *model.Page, user *model.User) (variables []*model.Variable, err error) {
	err = validateOrderByVariableField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("variable-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for variable")
		return
	}

	page.Total, err = reader.ListVariableTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list variable toal count")
		return
	}

	variables, err = reader.ListVariable(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list variable")
		return
	}
	for i, variable := range variables {
		err = sanitizeVariable(variable, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize variable element %d", i)
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

//ListVariableBySearch will request any variable matching the pattern of name
func ListVariableBySearch(page *model.Page, variable *model.Variable, user *model.User) (variables []*model.Variable, err error) {

	err = validateOrderByVariableField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareVariable(variable, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre variable")
		return
	}

	err = validateVariable(variable, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate variable")
		return
	}
	reader, err := getReader("variable-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get variable reader")
		return
	}

	variables, err = reader.ListVariableBySearch(page, variable)
	if err != nil {
		err = errors.Wrap(err, "failed to list variable by search")
		return
	}

	for _, variable := range variables {
		err = sanitizeVariable(variable, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize variable")
			return
		}
	}

	err = sanitizeVariable(variable, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search variable")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateVariable will create an variable using provided information
func CreateVariable(variable *model.Variable, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list variable by search without guide+")
		return
	}
	err = prepareVariable(variable, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare variable")
		return
	}

	err = validateVariable(variable, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate variable")
		return
	}
	//variable.TimeCreation = time.Now().Unix()
	writer, err := getWriter("variable")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for variable")
		return
	}
	err = writer.CreateVariable(variable)
	if err != nil {
		err = errors.Wrap(err, "failed to create variable")
		return
	}

	memWriter, err := getWriter("variable-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for variable-memory")
		return
	}
	err = memWriter.CreateVariable(variable)
	if err != nil {
		err = errors.Wrap(err, "failed to edit variable-memory")
		return
	}

	err = sanitizeVariable(variable, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize variable")
		return
	}
	return
}

//GetVariableValueFloat is a quickhand fetch for a value in float64 format
func GetVariableValueFloat(name string) (value float64) {
	valueStr := GetVariableValue(name)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return
	}
	return
}

//GetVariableValue is a quickhand fetch for a value in string format
func GetVariableValue(name string) (value string) {
	variable := &model.Variable{
		Name: name,
	}
	user := &model.User{}
	err := GetVariable(variable, user)
	if err != nil {
		return
	}
	value = variable.Value
	return
}

//GetVariable gets an variable by provided variableID
func GetVariable(variable *model.Variable, user *model.User) (err error) {
	err = prepareVariable(variable, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare variable")
		return
	}

	err = validateVariable(variable, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate variable")
		return
	}

	reader, err := getReader("variable-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get variable reader")
		return
	}

	err = reader.GetVariable(variable)
	if err != nil {
		err = errors.Wrap(err, "failed to get variable")
		return
	}

	err = sanitizeVariable(variable, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize variable")
		return
	}

	return
}

//EditVariable edits an existing variable
func EditVariable(variable *model.Variable, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list variable by search without guide+")
		return
	}
	err = prepareVariable(variable, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare variable")
		return
	}

	err = validateVariable(variable,
		[]string{"name"}, //required
		[]string{ //optional
			"value",
			"description"},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate variable")
		return
	}
	writer, err := getWriter("variable")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for variable")
		return
	}
	err = writer.EditVariable(variable)
	if err != nil {
		err = errors.Wrap(err, "failed to edit variable")
		return
	}

	memWriter, err := getWriter("variable-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for variable-memory")
		return
	}
	err = memWriter.EditVariable(variable)
	if err != nil {
		err = errors.Wrap(err, "failed to edit variable-memory")
		return
	}

	err = sanitizeVariable(variable, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize variable")
		return
	}
	return
}

//DeleteVariable deletes an variable by provided variableID
func DeleteVariable(variable *model.Variable, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete variable without admin+")
		return
	}
	err = prepareVariable(variable, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare variable")
		return
	}

	err = validateVariable(variable, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate variable")
		return
	}
	writer, err := getWriter("variable")
	if err != nil {
		err = errors.Wrap(err, "failed to get variable writer")
		return
	}
	err = writer.DeleteVariable(variable)
	if err != nil {
		err = errors.Wrap(err, "failed to delete variable")
		return
	}

	memWriter, err := getWriter("variable-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for variable-memory")
		return
	}
	err = memWriter.DeleteVariable(variable)
	if err != nil {
		err = errors.Wrap(err, "failed to delete variable-memory")
		return
	}
	return
}

func prepareVariable(variable *model.Variable, user *model.User) (err error) {
	if variable == nil {
		err = fmt.Errorf("empty variable")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateVariable(variable *model.Variable, required []string, optional []string) (err error) {
	schema, err := newSchemaVariable(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(variable))
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

func validateOrderByVariableField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "name"
	}

	validNames := []string{
		"name",
		"value",
		"description",
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

func sanitizeVariable(variable *model.Variable, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaVariable(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyVariable(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyVariable(field); err != nil {
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

func getSchemaPropertyVariable(field string) (prop model.Schema, err error) {
	switch field {
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 25
		prop.Pattern = "^[a-zA-Z_]*$"
	case "value":
		prop.Type = "string"
		prop.MaxLength = 256
	case "description":
		prop.Type = "string"
		prop.MaxLength = 256
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
