package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadSpellAnimationFromFileToMemory is ran during initialization
func LoadSpellAnimationFromFileToMemory() (err error) {

	fr, err := file.New("config", "spellAnimation.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("spellAnimation-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellAnimation-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("spellAnimation-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spellAnimation-memory")
		return
	}

	fileReader, err := getReader("spellAnimation-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimation-file reader")
		return
	}

	memWriter, err := getWriter("spellAnimation-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimation-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListSpellAnimationTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list spellAnimation count")
		return
	}
	page.Limit = page.Total

	spellAnimations, err := fileReader.ListSpellAnimation(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellAnimations")
		return
	}

	for _, spellAnimation := range spellAnimations {
		err = memWriter.CreateSpellAnimation(spellAnimation)
		if err != nil {
			err = errors.Wrap(err, "failed to create spellAnimation")
			return
		}
	}

	fmt.Printf("%d spellAnimations, ", len(spellAnimations))
	return
}

//ListSpellAnimation lists all spellAnimations accessible by provided user
func ListSpellAnimation(page *model.Page, user *model.User) (spellAnimations []*model.SpellAnimation, err error) {
	err = validateOrderBySpellAnimationField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("spellAnimation-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for spellAnimation")
		return
	}

	page.Total, err = reader.ListSpellAnimationTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list spellAnimation toal count")
		return
	}

	spellAnimations, err = reader.ListSpellAnimation(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellAnimation")
		return
	}
	for i, spellAnimation := range spellAnimations {
		err = sanitizeSpellAnimation(spellAnimation, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize spellAnimation element %d", i)
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

//ListSpellAnimationBySearch will request any spellAnimation matching the pattern of name
func ListSpellAnimationBySearch(page *model.Page, spellAnimation *model.SpellAnimation, user *model.User) (spellAnimations []*model.SpellAnimation, err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellAnimation by search without guide+")
		return
	}

	err = validateOrderBySpellAnimationField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSpellAnimation(spellAnimation, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre spellAnimation")
		return
	}

	err = validateSpellAnimation(spellAnimation, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellAnimation")
		return
	}
	reader, err := getReader("spellAnimation-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimation-memory reader")
		return
	}

	spellAnimations, err = reader.ListSpellAnimationBySearch(page, spellAnimation)
	if err != nil {
		err = errors.Wrap(err, "failed to list spellAnimation by search")
		return
	}

	for _, spellAnimation := range spellAnimations {
		err = sanitizeSpellAnimation(spellAnimation, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize spellAnimation")
			return
		}
	}

	err = sanitizeSpellAnimation(spellAnimation, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search spellAnimation")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSpellAnimation will create an spellAnimation using provided information
func CreateSpellAnimation(spellAnimation *model.SpellAnimation, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellAnimation by search without guide+")
		return
	}
	err = prepareSpellAnimation(spellAnimation, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellAnimation")
		return
	}

	err = validateSpellAnimation(spellAnimation, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellAnimation")
		return
	}
	spellAnimation.ID = 0
	writer, err := getWriter("spellAnimation-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellAnimation")
		return
	}
	err = writer.CreateSpellAnimation(spellAnimation)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellAnimation")
		return
	}

	fileWriter, err := getWriter("spellAnimation-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimation-file writer")
		return
	}
	err = fileWriter.CreateSpellAnimation(spellAnimation)
	if err != nil {
		err = errors.Wrap(err, "failed to create spellAnimation-file")
		return
	}
	err = sanitizeSpellAnimation(spellAnimation, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellAnimation")
		return
	}
	return
}

//GetSpellAnimation gets an spellAnimation by provided spellAnimationID
func GetSpellAnimation(spellAnimation *model.SpellAnimation, user *model.User) (err error) {
	err = prepareSpellAnimation(spellAnimation, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellAnimation")
		return
	}

	err = validateSpellAnimation(spellAnimation, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellAnimation")
		return
	}

	reader, err := getReader("spellAnimation-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimation-memory reader")
		return
	}

	err = reader.GetSpellAnimation(spellAnimation)
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimation")
		return
	}

	err = sanitizeSpellAnimation(spellAnimation, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellAnimation")
		return
	}

	return
}

//EditSpellAnimation edits an existing spellAnimation
func EditSpellAnimation(spellAnimation *model.SpellAnimation, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spellAnimation by search without guide+")
		return
	}
	err = prepareSpellAnimation(spellAnimation, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellAnimation")
		return
	}

	err = validateSpellAnimation(spellAnimation,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellAnimation")
		return
	}
	writer, err := getWriter("spellAnimation-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spellAnimation")
		return
	}
	err = writer.EditSpellAnimation(spellAnimation)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellAnimation")
		return
	}

	fileWriter, err := getWriter("spellAnimation-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimation-file writer")
		return
	}
	err = fileWriter.EditSpellAnimation(spellAnimation)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spellAnimation-file")
		return
	}

	err = sanitizeSpellAnimation(spellAnimation, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spellAnimation")
		return
	}
	return
}

//DeleteSpellAnimation deletes an spellAnimation by provided spellAnimationID
func DeleteSpellAnimation(spellAnimation *model.SpellAnimation, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete spellAnimation without admin+")
		return
	}
	err = prepareSpellAnimation(spellAnimation, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spellAnimation")
		return
	}

	err = validateSpellAnimation(spellAnimation, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spellAnimation")
		return
	}
	writer, err := getWriter("spellAnimation-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimation-memory writer")
		return
	}
	err = writer.DeleteSpellAnimation(spellAnimation)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellAnimation")
		return
	}

	fileWriter, err := getWriter("spellAnimation-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get spellAnimation-file writer")
		return
	}
	err = fileWriter.DeleteSpellAnimation(spellAnimation)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spellAnimation-file")
		return
	}

	return
}

func prepareSpellAnimation(spellAnimation *model.SpellAnimation, user *model.User) (err error) {
	if spellAnimation == nil {
		err = fmt.Errorf("empty spellAnimation")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSpellAnimation(spellAnimation *model.SpellAnimation, required []string, optional []string) (err error) {
	schema, err := newSchemaSpellAnimation(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spellAnimation))
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

func validateOrderBySpellAnimationField(page *model.Page) (err error) {
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

func sanitizeSpellAnimation(spellAnimation *model.SpellAnimation, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	spellAnimation.Type = &model.SpellAnimationType{
		ID: spellAnimation.TypeID,
	}
	err = GetSpellAnimationType(spellAnimation.Type, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get spell animation type during sanitize of spell animation")
		return
	}
	return
}

func newSchemaSpellAnimation(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySpellAnimation(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySpellAnimation(field); err != nil {
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

func getSchemaPropertySpellAnimation(field string) (prop model.Schema, err error) {
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
