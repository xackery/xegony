package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadRuleFromDBToMemory is ran during initialization
func LoadRuleFromDBToMemory() (err error) {

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("rule-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize rule-memory")
		return
	}

	dbReader, err := getReader("rule")
	if err != nil {
		err = errors.Wrap(err, "failed to get rule reader")
		return
	}

	memWriter, err := getWriter("rule-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get rule-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = dbReader.ListRuleTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list rule count")
		return
	}
	var totalRules []*model.Rule
	var rules []*model.Rule
	for {
		rules, err = dbReader.ListRule(page)
		if err != nil {
			err = errors.Wrap(err, "failed to list rules")
			return
		}
		totalRules = append(totalRules, rules...)
		if int64(len(totalRules)) >= page.Total {
			break
		}
		page.Offset++
	}

	for _, rule := range totalRules {
		err = memWriter.CreateRule(rule)
		if err != nil {
			err = errors.Wrap(err, "failed to create rule")
			return
		}
	}

	fmt.Printf("%d rules, ", len(totalRules))
	return
}

//ListRule lists all rules accessible by provided user
func ListRule(page *model.Page, user *model.User) (rules []*model.Rule, err error) {
	err = validateOrderByRuleField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("rule-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for rule")
		return
	}

	page.Total, err = reader.ListRuleTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list rule toal count")
		return
	}

	rules, err = reader.ListRule(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list rule")
		return
	}
	for i, rule := range rules {
		err = sanitizeRule(rule, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize rule element %d", i)
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

//ListRuleBySearch will request any rule matching the pattern of name
func ListRuleBySearch(page *model.Page, rule *model.Rule, user *model.User) (rules []*model.Rule, err error) {
	/*err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list rule by search without guide+")
		return
	}
	*/
	err = validateOrderByRuleField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre rule")
		return
	}

	err = validateRule(rule, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate rule")
		return
	}
	reader, err := getReader("rule-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get rule reader")
		return
	}

	rules, err = reader.ListRuleBySearch(page, rule)
	if err != nil {
		err = errors.Wrap(err, "failed to list rule by search")
		return
	}

	for _, rule := range rules {
		err = sanitizeRule(rule, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize rule")
			return
		}
	}

	err = sanitizeRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search rule")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateRule will create an rule using provided information
func CreateRule(rule *model.Rule, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list rule by search without guide+")
		return
	}
	err = prepareRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare rule")
		return
	}

	err = validateRule(rule, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate rule")
		return
	}
	rule.ID = 0
	//rule.TimeCreation = time.Now().Unix()
	writer, err := getWriter("rule")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for rule")
		return
	}
	err = writer.CreateRule(rule)
	if err != nil {
		err = errors.Wrap(err, "failed to create rule")
		return
	}

	memWriter, err := getWriter("rule-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for rule-memory")
		return
	}
	err = memWriter.CreateRule(rule)
	if err != nil {
		err = errors.Wrap(err, "failed to edit rule-memory")
		return
	}

	err = sanitizeRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize rule")
		return
	}
	return
}

//GetRule gets an rule by provided ruleID
func GetRule(rule *model.Rule, user *model.User) (err error) {
	err = prepareRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare rule")
		return
	}

	err = validateRule(rule, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate rule")
		return
	}

	reader, err := getReader("rule-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get rule reader")
		return
	}

	err = reader.GetRule(rule)
	if err != nil {
		err = errors.Wrap(err, "failed to get rule")
		return
	}

	err = sanitizeRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize rule")
		return
	}

	return
}

//EditRule edits an existing rule
func EditRule(rule *model.Rule, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list rule by search without guide+")
		return
	}
	err = prepareRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare rule")
		return
	}

	err = validateRule(rule,
		[]string{"ID"}, //required
		[]string{ //optional
			"name"},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate rule")
		return
	}
	writer, err := getWriter("rule")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for rule")
		return
	}
	err = writer.EditRule(rule)
	if err != nil {
		err = errors.Wrap(err, "failed to edit rule")
		return
	}

	memWriter, err := getWriter("rule-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for rule-memory")
		return
	}
	err = memWriter.EditRule(rule)
	if err != nil {
		err = errors.Wrap(err, "failed to edit rule-memory")
		return
	}

	err = sanitizeRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize rule")
		return
	}
	return
}

//DeleteRule deletes an rule by provided ruleID
func DeleteRule(rule *model.Rule, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete rule without admin+")
		return
	}
	err = prepareRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare rule")
		return
	}

	err = validateRule(rule, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate rule")
		return
	}
	writer, err := getWriter("rule")
	if err != nil {
		err = errors.Wrap(err, "failed to get rule writer")
		return
	}
	err = writer.DeleteRule(rule)
	if err != nil {
		err = errors.Wrap(err, "failed to delete rule")
		return
	}

	memWriter, err := getWriter("rule-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for rule-memory")
		return
	}
	err = memWriter.DeleteRule(rule)
	if err != nil {
		err = errors.Wrap(err, "failed to delete rule-memory")
		return
	}
	return
}

func prepareRule(rule *model.Rule, user *model.User) (err error) {
	if rule == nil {
		err = fmt.Errorf("empty rule")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateRule(rule *model.Rule, required []string, optional []string) (err error) {
	schema, err := newSchemaRule(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(rule))
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

func validateOrderByRuleField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "name"
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

func sanitizeRule(rule *model.Rule, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaRule(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyRule(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyRule(field); err != nil {
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

func getSchemaPropertyRule(field string) (prop model.Schema, err error) {
	switch field {
	case "ID":
		prop.Type = "integer"
		prop.Minimum = 0
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 64
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
