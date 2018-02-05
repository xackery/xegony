package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadSpellEffectFormulaFromFileToMemory is ran during initialization
func LoadSpellEffectFormulaFromFileToMemory() (err error) {

	fr, err := file.New("config", "spellEffectFormula.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("spellEffectFormula-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellEffectFormula-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("spellEffectFormula-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellEffectFormula-memory")
		return
	}

	fileReader, err := getReader("spellEffectFormula-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectFormula-file reader")
		return
	}

	memWriter, err := getWriter("spellEffectFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectFormula-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListSpellEffectFormulaTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list spellEffectFormula count")
		return
	}
	page.Limit = page.Total

	spellEffectFormulas, err := fileReader.ListSpellEffectFormula(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellEffectFormulas")
		return
	}

	for _, spellEffectFormula := range spellEffectFormulas {
		err = memWriter.CreateSpellEffectFormula(spellEffectFormula)
		if err != nil {
			err = errors.Wrap(err, "failed to create spellEffectFormula")
			return
		}
	}

	fmt.Printf("%d spellEffectFormulas, ", len(spellEffectFormulas))
	return
}

//ListSpellEffectFormula lists all spellEffectFormulas accessible by provided user
func ListSpellEffectFormula(page *model.Page, user *model.User) (spellEffectFormulas []*model.SpellEffectFormula, err error) {
	err = validateOrderBySpellEffectFormulaField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("spellEffectFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for spellEffectFormula")
		return
	}

	page.Total, err = reader.ListSpellEffectFormulaTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list spellEffectFormula toal count")
		return
	}

	spellEffectFormulas, err = reader.ListSpellEffectFormula(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellEffectFormula")
		return
	}
	for i, spellEffectFormula := range spellEffectFormulas {
		err = sanitizeSpellEffectFormula(spellEffectFormula, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize spellEffectFormula element %d", i)
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

//ListSpellEffectFormulaBySearch will request any spellEffectFormula matching the pattern of name
func ListSpellEffectFormulaBySearch(page *model.Page, spellEffectFormula *model.SpellEffectFormula, user *model.User) (spellEffectFormulas []*model.SpellEffectFormula, err error) {

	err = validateOrderBySpellEffectFormulaField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSpellEffectFormula(spellEffectFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre spellEffectFormula")
		return
	}

	err = validateSpellEffectFormula(spellEffectFormula, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellEffectFormula")
		return
	}
	reader, err := getReader("spellEffectFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectFormula-memory reader")
		return
	}

	spellEffectFormulas, err = reader.ListSpellEffectFormulaBySearch(page, spellEffectFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellEffectFormula by search")
		return
	}

	for _, spellEffectFormula := range spellEffectFormulas {
		err = sanitizeSpellEffectFormula(spellEffectFormula, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize spellEffectFormula")
			return
		}
	}

	err = sanitizeSpellEffectFormula(spellEffectFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search spellEffectFormula")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSpellEffectFormula will create an spellEffectFormula using provided information
func CreateSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellEffectFormula by search without guide+")
		return
	}
	err = prepareSpellEffectFormula(spellEffectFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellEffectFormula")
		return
	}

	err = validateSpellEffectFormula(spellEffectFormula, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellEffectFormula")
		return
	}
	spellEffectFormula.ID = 0
	writer, err := getWriter("spellEffectFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellEffectFormula")
		return
	}
	err = writer.CreateSpellEffectFormula(spellEffectFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellEffectFormula")
		return
	}

	fileWriter, err := getWriter("spellEffectFormula-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectFormula-file writer")
		return
	}
	err = fileWriter.CreateSpellEffectFormula(spellEffectFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellEffectFormula-file")
		return
	}
	err = sanitizeSpellEffectFormula(spellEffectFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellEffectFormula")
		return
	}
	return
}

//GetSpellEffectFormula gets an spellEffectFormula by provided spellEffectFormulaID
func GetSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula, user *model.User) (err error) {
	err = prepareSpellEffectFormula(spellEffectFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellEffectFormula")
		return
	}

	err = validateSpellEffectFormula(spellEffectFormula, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellEffectFormula")
		return
	}

	reader, err := getReader("spellEffectFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectFormula-memory reader")
		return
	}

	err = reader.GetSpellEffectFormula(spellEffectFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectFormula")
		return
	}

	err = sanitizeSpellEffectFormula(spellEffectFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellEffectFormula")
		return
	}

	return
}

//EditSpellEffectFormula edits an existing spellEffectFormula
func EditSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellEffectFormula by search without guide+")
		return
	}
	err = prepareSpellEffectFormula(spellEffectFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellEffectFormula")
		return
	}

	err = validateSpellEffectFormula(spellEffectFormula,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"male",
			"female",
			"neutral",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellEffectFormula")
		return
	}
	writer, err := getWriter("spellEffectFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellEffectFormula")
		return
	}
	err = writer.EditSpellEffectFormula(spellEffectFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellEffectFormula")
		return
	}

	fileWriter, err := getWriter("spellEffectFormula-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectFormula-file writer")
		return
	}
	err = fileWriter.EditSpellEffectFormula(spellEffectFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellEffectFormula-file")
		return
	}

	err = sanitizeSpellEffectFormula(spellEffectFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellEffectFormula")
		return
	}
	return
}

//DeleteSpellEffectFormula deletes an spellEffectFormula by provided spellEffectFormulaID
func DeleteSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete spellEffectFormula without admin+")
		return
	}
	err = prepareSpellEffectFormula(spellEffectFormula, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellEffectFormula")
		return
	}

	err = validateSpellEffectFormula(spellEffectFormula, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellEffectFormula")
		return
	}
	writer, err := getWriter("spellEffectFormula-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectFormula-memory writer")
		return
	}
	err = writer.DeleteSpellEffectFormula(spellEffectFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellEffectFormula")
		return
	}

	fileWriter, err := getWriter("spellEffectFormula-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellEffectFormula-file writer")
		return
	}
	err = fileWriter.DeleteSpellEffectFormula(spellEffectFormula)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellEffectFormula-file")
		return
	}

	return
}

func prepareSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula, user *model.User) (err error) {
	if spellEffectFormula == nil {
		err = fmt.Errorf("empty spellEffectFormula")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula, required []string, optional []string) (err error) {
	schema, err := newSchemaSpellEffectFormula(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spellEffectFormula))
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

func validateOrderBySpellEffectFormulaField(page *model.Page) (err error) {
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

func sanitizeSpellEffectFormula(spellEffectFormula *model.SpellEffectFormula, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaSpellEffectFormula(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySpellEffectFormula(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySpellEffectFormula(field); err != nil {
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

func getSchemaPropertySpellEffectFormula(field string) (prop model.Schema, err error) {
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
