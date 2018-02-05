package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadSpellDurationFormulaFromFileToMemory is ran during initialization
func LoadSpellDurationFormulaFromFileToMemory() (err error) {

	fr, err := file.New("config", "spellDurationFormula.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("spellDurationFormula-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellDurationFormula-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("spellDurationFormula-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellDurationFormula-memory")
		return
	}

	fileReader, err := getReader("spellDurationFormula-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellDurationFormula-file reader")
		return
	}

	memWriter, err := getWriter("spellDurationFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellDurationFormula-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListSpellDurationFormulaTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list spellDurationFormula count")
		return
	}
	page.Limit = page.Total

	spellDurationFormulas, err := fileReader.ListSpellDurationFormula(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellDurationFormulas")
		return
	}

	for _, spellDurationFormula := range spellDurationFormulas {
		err = memWriter.CreateSpellDurationFormula(spellDurationFormula)
		if err != nil {
			err = errors.Wrap(err, "failed to create spellDurationFormula")
			return
		}
	}

	fmt.Printf("%d spellDurationFormulas, ", len(spellDurationFormulas))
	return
}

//ListSpellDurationFormula lists all spellDurationFormulas accessible by provided user
func ListSpellDurationFormula(page *model.Page, user *model.User) (spellDurationFormulas []*model.SpellDurationFormula, err error) {
	err = validateOrderBySpellDurationFormulaField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("spellDurationFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for spellDurationFormula")
		return
	}

	page.Total, err = reader.ListSpellDurationFormulaTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list spellDurationFormula toal count")
		return
	}

	spellDurationFormulas, err = reader.ListSpellDurationFormula(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellDurationFormula")
		return
	}
	for i, spellDurationFormula := range spellDurationFormulas {
		err = sanitizeSpellDurationFormula(spellDurationFormula, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize spellDurationFormula element %d", i)
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

//ListSpellDurationFormulaBySearch will request any spellDurationFormula matching the pattern of name
func ListSpellDurationFormulaBySearch(page *model.Page, spellDurationFormula *model.SpellDurationFormula, user *model.User) (spellDurationFormulas []*model.SpellDurationFormula, err error) {

	err = validateOrderBySpellDurationFormulaField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSpellDurationFormula(spellDurationFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre spellDurationFormula")
		return
	}

	err = validateSpellDurationFormula(spellDurationFormula, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellDurationFormula")
		return
	}
	reader, err := getReader("spellDurationFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellDurationFormula-memory reader")
		return
	}

	spellDurationFormulas, err = reader.ListSpellDurationFormulaBySearch(page, spellDurationFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellDurationFormula by search")
		return
	}

	for _, spellDurationFormula := range spellDurationFormulas {
		err = sanitizeSpellDurationFormula(spellDurationFormula, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize spellDurationFormula")
			return
		}
	}

	err = sanitizeSpellDurationFormula(spellDurationFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search spellDurationFormula")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSpellDurationFormula will create an spellDurationFormula using provided information
func CreateSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellDurationFormula by search without guide+")
		return
	}
	err = prepareSpellDurationFormula(spellDurationFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellDurationFormula")
		return
	}

	err = validateSpellDurationFormula(spellDurationFormula, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellDurationFormula")
		return
	}
	spellDurationFormula.ID = 0
	writer, err := getWriter("spellDurationFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellDurationFormula")
		return
	}
	err = writer.CreateSpellDurationFormula(spellDurationFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellDurationFormula")
		return
	}

	fileWriter, err := getWriter("spellDurationFormula-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellDurationFormula-file writer")
		return
	}
	err = fileWriter.CreateSpellDurationFormula(spellDurationFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellDurationFormula-file")
		return
	}
	err = sanitizeSpellDurationFormula(spellDurationFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellDurationFormula")
		return
	}
	return
}

//GetSpellDurationFormula gets an spellDurationFormula by provided spellDurationFormulaID
func GetSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula, user *model.User) (err error) {
	err = prepareSpellDurationFormula(spellDurationFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellDurationFormula")
		return
	}

	err = validateSpellDurationFormula(spellDurationFormula, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellDurationFormula")
		return
	}

	reader, err := getReader("spellDurationFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellDurationFormula-memory reader")
		return
	}

	err = reader.GetSpellDurationFormula(spellDurationFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to get spellDurationFormula")
		return
	}

	err = sanitizeSpellDurationFormula(spellDurationFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellDurationFormula")
		return
	}

	return
}

//EditSpellDurationFormula edits an existing spellDurationFormula
func EditSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellDurationFormula by search without guide+")
		return
	}
	err = prepareSpellDurationFormula(spellDurationFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellDurationFormula")
		return
	}

	err = validateSpellDurationFormula(spellDurationFormula,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"male",
			"female",
			"neutral",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellDurationFormula")
		return
	}
	writer, err := getWriter("spellDurationFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellDurationFormula")
		return
	}
	err = writer.EditSpellDurationFormula(spellDurationFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellDurationFormula")
		return
	}

	fileWriter, err := getWriter("spellDurationFormula-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellDurationFormula-file writer")
		return
	}
	err = fileWriter.EditSpellDurationFormula(spellDurationFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellDurationFormula-file")
		return
	}

	err = sanitizeSpellDurationFormula(spellDurationFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellDurationFormula")
		return
	}
	return
}

//DeleteSpellDurationFormula deletes an spellDurationFormula by provided spellDurationFormulaID
func DeleteSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete spellDurationFormula without admin+")
		return
	}
	err = prepareSpellDurationFormula(spellDurationFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellDurationFormula")
		return
	}

	err = validateSpellDurationFormula(spellDurationFormula, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellDurationFormula")
		return
	}
	writer, err := getWriter("spellDurationFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellDurationFormula-memory writer")
		return
	}
	err = writer.DeleteSpellDurationFormula(spellDurationFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellDurationFormula")
		return
	}

	fileWriter, err := getWriter("spellDurationFormula-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellDurationFormula-file writer")
		return
	}
	err = fileWriter.DeleteSpellDurationFormula(spellDurationFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellDurationFormula-file")
		return
	}

	return
}

func prepareSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula, user *model.User) (err error) {
	if spellDurationFormula == nil {
		err = fmt.Errorf("empty spellDurationFormula")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula, required []string, optional []string) (err error) {
	schema, err := newSchemaSpellDurationFormula(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spellDurationFormula))
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

func validateOrderBySpellDurationFormulaField(page *model.Page) (err error) {
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

func sanitizeSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaSpellDurationFormula(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySpellDurationFormula(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySpellDurationFormula(field); err != nil {
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

func getSchemaPropertySpellDurationFormula(field string) (prop model.Schema, err error) {
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
