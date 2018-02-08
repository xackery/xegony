package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadSkillFromFileToMemory is ran during initialization
func LoadSkillFromFileToMemory() (err error) {

	fr, err := file.New("config", "skill.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("skill-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize skill-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("skill-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize skill-memory")
		return
	}

	fileReader, err := getReader("skill-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get skill-file reader")
		return
	}

	memWriter, err := getWriter("skill-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get skill-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListSkillTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list skill count")
		return
	}
	page.Limit = page.Total

	skills, err := fileReader.ListSkill(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list skills")
		return
	}

	for _, skill := range skills {
		err = memWriter.CreateSkill(skill)
		if err != nil {
			err = errors.Wrap(err, "failed to create skill")
			return
		}
	}

	fmt.Printf("%d skills, ", len(skills))
	return
}

//ListSkill lists all skills accessible by provided user
func ListSkill(page *model.Page, user *model.User) (skills []*model.Skill, err error) {
	err = validateOrderBySkillField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("skill-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for skill")
		return
	}

	page.Total, err = reader.ListSkillTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list skill toal count")
		return
	}

	skills, err = reader.ListSkill(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list skill")
		return
	}
	for i, skill := range skills {
		err = sanitizeSkill(skill, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize skill element %d", i)
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

//ListSkillBySearch will request any skill matching the pattern of name
func ListSkillBySearch(page *model.Page, skill *model.Skill, user *model.User) (skills []*model.Skill, err error) {

	err = validateOrderBySkillField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSkill(skill, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre skill")
		return
	}

	err = validateSkill(skill, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate skill")
		return
	}
	reader, err := getReader("skill-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get skill-memory reader")
		return
	}

	skills, err = reader.ListSkillBySearch(page, skill)
	if err != nil {
		err = errors.Wrap(err, "failed to list skill by search")
		return
	}

	for _, skill := range skills {
		err = sanitizeSkill(skill, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize skill")
			return
		}
	}

	err = sanitizeSkill(skill, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search skill")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSkill will create an skill using provided information
func CreateSkill(skill *model.Skill, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list skill by search without guide+")
		return
	}
	err = prepareSkill(skill, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare skill")
		return
	}

	err = validateSkill(skill, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate skill")
		return
	}
	skill.ID = 0
	writer, err := getWriter("skill-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for skill")
		return
	}
	err = writer.CreateSkill(skill)
	if err != nil {
		err = errors.Wrap(err, "failed to create skill")
		return
	}

	fileWriter, err := getWriter("skill-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get skill-file writer")
		return
	}
	err = fileWriter.CreateSkill(skill)
	if err != nil {
		err = errors.Wrap(err, "failed to create skill-file")
		return
	}
	err = sanitizeSkill(skill, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize skill")
		return
	}
	return
}

//GetSkill gets an skill by provided skillID
func GetSkill(skill *model.Skill, user *model.User) (err error) {
	err = prepareSkill(skill, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare skill")
		return
	}

	err = validateSkill(skill, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate skill")
		return
	}

	reader, err := getReader("skill-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get skill-memory reader")
		return
	}

	err = reader.GetSkill(skill)
	if err != nil {
		err = errors.Wrap(err, "failed to get skill")
		return
	}

	err = sanitizeSkill(skill, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize skill")
		return
	}

	return
}

//EditSkill edits an existing skill
func EditSkill(skill *model.Skill, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list skill by search without guide+")
		return
	}
	err = prepareSkill(skill, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare skill")
		return
	}

	err = validateSkill(skill,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"male",
			"female",
			"neutral",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate skill")
		return
	}
	writer, err := getWriter("skill-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for skill")
		return
	}
	err = writer.EditSkill(skill)
	if err != nil {
		err = errors.Wrap(err, "failed to edit skill")
		return
	}

	fileWriter, err := getWriter("skill-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get skill-file writer")
		return
	}
	err = fileWriter.EditSkill(skill)
	if err != nil {
		err = errors.Wrap(err, "failed to edit skill-file")
		return
	}

	err = sanitizeSkill(skill, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize skill")
		return
	}
	return
}

//DeleteSkill deletes an skill by provided skillID
func DeleteSkill(skill *model.Skill, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete skill without admin+")
		return
	}
	err = prepareSkill(skill, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare skill")
		return
	}

	err = validateSkill(skill, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate skill")
		return
	}
	writer, err := getWriter("skill-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get skill-memory writer")
		return
	}
	err = writer.DeleteSkill(skill)
	if err != nil {
		err = errors.Wrap(err, "failed to delete skill")
		return
	}

	fileWriter, err := getWriter("skill-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get skill-file writer")
		return
	}
	err = fileWriter.DeleteSkill(skill)
	if err != nil {
		err = errors.Wrap(err, "failed to delete skill-file")
		return
	}

	return
}

func prepareSkill(skill *model.Skill, user *model.User) (err error) {
	if skill == nil {
		err = fmt.Errorf("empty skill")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSkill(skill *model.Skill, required []string, optional []string) (err error) {
	schema, err := newSchemaSkill(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(skill))
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

func validateOrderBySkillField(page *model.Page) (err error) {
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

func sanitizeSkill(skill *model.Skill, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	if skill.Icon == "" {
		skill.Icon = "xa-octopus"
	}
	if len(skill.Icon) > 2 && skill.Icon[2:3] == "-" {
		skill.Icon = fmt.Sprintf("%s %s", skill.Icon[0:2], skill.Icon)
	}
	return
}

func newSchemaSkill(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySkill(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySkill(field); err != nil {
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

func getSchemaPropertySkill(field string) (prop model.Schema, err error) {
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
