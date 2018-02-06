package cases

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadRuleEntryFromDBToMemory is ran during initialization
func LoadRuleEntryFromDBToMemory() (err error) {

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("ruleEntry-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize ruleEntry-memory")
		return
	}

	dbReader, err := getReader("ruleEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get ruleEntry reader")
		return
	}

	memWriter, err := getWriter("ruleEntry-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get ruleEntry-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	var totalRuleEntrys []*model.RuleEntry
	var ruleEntrys []*model.RuleEntry
	var rule *model.Rule
	for i := 0; i < 10; i++ {
		rule = &model.Rule{
			ID: int64(i),
		}
		page.Total, err = dbReader.ListRuleEntryTotalCount(rule)
		if err != nil {
			err = errors.Wrap(err, "failed to get list ruleEntry count")
			return
		}
		for {
			ruleEntrys, err = dbReader.ListRuleEntry(page, rule)
			if err != nil {
				err = errors.Wrap(err, "failed to list ruleEntrys")
				return
			}
			totalRuleEntrys = append(totalRuleEntrys, ruleEntrys...)
			if int64(len(totalRuleEntrys)) >= page.Total {
				break
			}
			page.Offset++
		}

	}
	rule = &model.Rule{}
	for _, ruleEntry := range totalRuleEntrys {
		rule.ID = ruleEntry.RuleID
		err = memWriter.CreateRuleEntry(rule, ruleEntry)
		if err != nil {
			err = errors.Wrap(err, "failed to create ruleEntry")
			return
		}
	}

	fmt.Printf("%d rule entries, ", len(totalRuleEntrys))
	return
}

//ListRuleEntry lists all ruleEntrys accessible by provided user
func ListRuleEntry(page *model.Page, rule *model.Rule, user *model.User) (ruleEntrys []*model.RuleEntry, err error) {
	err = validateOrderByRuleEntryField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("ruleEntry-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for ruleEntry")
		return
	}

	page.Total, err = reader.ListRuleEntryTotalCount(rule)
	if err != nil {
		err = errors.Wrap(err, "failed to list ruleEntry toal count")
		return
	}

	ruleEntrys, err = reader.ListRuleEntry(page, rule)
	if err != nil {
		err = errors.Wrap(err, "failed to list ruleEntry")
		return
	}
	for i, ruleEntry := range ruleEntrys {
		err = sanitizeRuleEntry(ruleEntry, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize ruleEntry element %d", i)
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

//ListRuleEntryBySearch will request any ruleEntry matching the pattern of name
func ListRuleEntryBySearch(page *model.Page, rule *model.Rule, ruleEntry *model.RuleEntry, user *model.User) (ruleEntrys []*model.RuleEntry, err error) {

	err = validateOrderByRuleEntryField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareRuleEntry(ruleEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre ruleEntry")
		return
	}

	err = validateRuleEntry(ruleEntry, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate ruleEntry")
		return
	}
	reader, err := getReader("ruleEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get ruleEntry reader")
		return
	}

	ruleEntrys, err = reader.ListRuleEntryBySearch(page, rule, ruleEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to list ruleEntry by search")
		return
	}

	page.Total, err = reader.ListRuleEntryBySearchTotalCount(rule, ruleEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to list ruleEntry by search total count")
		return
	}
	for _, ruleEntry := range ruleEntrys {
		err = sanitizeRuleEntry(ruleEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize ruleEntry")
			return
		}
	}

	err = sanitizeRuleEntry(ruleEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search ruleEntry")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateRuleEntry will create an ruleEntry using provided information
func CreateRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list ruleEntry by search without guide+")
		return
	}
	err = prepareRuleEntry(ruleEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare ruleEntry")
		return
	}

	err = validateRuleEntry(ruleEntry, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate ruleEntry")
		return
	}
	ruleEntry.RuleID = rule.ID
	//ruleEntry.TimeCreation = time.Now().Unix()
	writer, err := getWriter("ruleEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for ruleEntry")
		return
	}
	err = writer.CreateRuleEntry(rule, ruleEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to create ruleEntry")
		return
	}
	err = sanitizeRuleEntry(ruleEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize ruleEntry")
		return
	}
	return
}

//GetRuleEntryValueFloat returns a float value
func GetRuleEntryValueFloat(ruleID int64, name string) (value float64) {
	valStr := GetRuleEntryValue(ruleID, name)
	if len(valStr) == 0 {
		return
	}
	value, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		fmt.Println("(warning) Failed to get float64 of rule:", ruleID, name, err.Error())
		return
	}
	return
}

//GetRuleEntryValue quickly returns a rule
func GetRuleEntryValue(ruleID int64, name string) (value string) {
	rule := &model.Rule{
		ID: ruleID,
	}
	ruleEntry := &model.RuleEntry{
		RuleID: ruleID,
		Name:   name,
	}
	user := &model.User{}
	err := GetRuleEntry(rule, ruleEntry, user)
	if err != nil {
		fmt.Println("(warning) Failed to get rule", ruleID, name, err.Error())
		return
	}
	value = ruleEntry.Value
	return
}

//GetRuleEntry gets an ruleEntry by provided ruleEntryID
func GetRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry, user *model.User) (err error) {
	err = prepareRuleEntry(ruleEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare ruleEntry")
		return
	}

	err = validateRuleEntry(ruleEntry, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate ruleEntry")
		return
	}

	reader, err := getReader("ruleEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get ruleEntry reader")
		return
	}

	err = reader.GetRuleEntry(rule, ruleEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to get ruleEntry")
		return
	}

	err = sanitizeRuleEntry(ruleEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize ruleEntry")
		return
	}

	return
}

//EditRuleEntry edits an existing ruleEntry
func EditRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list ruleEntry by search without guide+")
		return
	}
	err = prepareRuleEntry(ruleEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare ruleEntry")
		return
	}

	err = validateRuleEntry(ruleEntry,
		[]string{"name"}, //required
		[]string{ //optional
			"value",
			"description"},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate ruleEntry")
		return
	}
	writer, err := getWriter("ruleEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for ruleEntry")
		return
	}
	err = writer.EditRuleEntry(rule, ruleEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to edit ruleEntry")
		return
	}
	err = sanitizeRuleEntry(ruleEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize ruleEntry")
		return
	}
	return
}

//DeleteRuleEntry deletes an ruleEntry by provided ruleEntryID
func DeleteRuleEntry(ruleEntry *model.RuleEntry, rule *model.Rule, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete ruleEntry without admin+")
		return
	}
	err = prepareRuleEntry(ruleEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare ruleEntry")
		return
	}

	err = validateRuleEntry(ruleEntry, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate ruleEntry")
		return
	}
	writer, err := getWriter("ruleEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get ruleEntry writer")
		return
	}
	err = writer.DeleteRuleEntry(rule, ruleEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to delete ruleEntry")
		return
	}
	return
}

func prepareRuleEntry(ruleEntry *model.RuleEntry, user *model.User) (err error) {
	if ruleEntry == nil {
		err = fmt.Errorf("empty ruleEntry")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateRuleEntry(ruleEntry *model.RuleEntry, required []string, optional []string) (err error) {
	schema, err := newSchemaRuleEntry(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(ruleEntry))
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

func validateOrderByRuleEntryField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "name"
	}

	validNames := []string{
		"name",
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

func sanitizeRuleEntry(ruleEntry *model.RuleEntry, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	if strings.Contains(ruleEntry.Name, ":") {
		ruleEntry.Scope = ruleEntry.Name[0:strings.Index(ruleEntry.Name, ":")]
	}
	return
}

func newSchemaRuleEntry(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyRuleEntry(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyRuleEntry(field); err != nil {
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

func getSchemaPropertyRuleEntry(field string) (prop model.Schema, err error) {
	switch field {
	case "ruleID":
		prop.Type = "integer"
		prop.Minimum = 0
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 64
	case "value":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 64
	case "description":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 64
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
