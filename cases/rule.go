package cases

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

var ()

//RuleRepository handles RuleRepository cases and is a gateway to storage
type RuleRepository struct {
	stor              storage.Storage
	ruleCache         map[string]*model.Rule
	isRuleCacheLoaded bool
}

//Initialize handler
func (c *RuleRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}

	c.stor = stor
	c.isRuleCacheLoaded = false
	if err = c.rebuildCache(); err != nil {
		return
	}
	return
}

func (c *RuleRepository) rebuildCache() (err error) {
	if c.isRuleCacheLoaded {
		return
	}
	c.isRuleCacheLoaded = true
	c.ruleCache = make(map[string]*model.Rule)
	rules, err := c.list()
	if err != nil {
		return
	}

	for _, rule := range rules {
		c.ruleCache[rule.Name] = rule
	}
	fmt.Println("Rebuilt Rule Cache")
	return
}

//Get handler
func (c *RuleRepository) Get(rule *model.Rule, user *model.User) (err error) {
	rule = c.ruleCache[rule.Name]
	err = c.prepare(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare rule")
		return
	}
	return
}

//Create handler
func (c *RuleRepository) Create(rule *model.Rule, user *model.User) (err error) {
	if rule == nil {
		err = fmt.Errorf("Empty rule")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
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
	err = c.stor.CreateRule(rule)
	if err != nil {
		return
	}
	//c.isRuleCacheLoaded = false
	//c.rebuildCache()
	err = c.prepare(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare rule")
		return
	}
	return
}

//Edit handler
func (c *RuleRepository) Edit(rule *model.Rule, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
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

	if err = c.stor.EditRule(rule); err != nil {
		return
	}
	//if err = c.rebuildCache(); err != nil {
	//	return
	//}
	err = c.prepare(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare rule")
		return
	}
	return
}

//Delete handler
func (c *RuleRepository) Delete(rule *model.Rule, user *model.User) (err error) {
	err = c.stor.DeleteRule(rule)
	if err != nil {
		return
	}
	//if err = c.rebuildCache(); err != nil {
	//	return
	//}
	return
}

func (c *RuleRepository) list() (rules []*model.Rule, err error) {
	if rules, err = c.stor.ListRule(); err != nil {
		return
	}
	return
}

//List handler
func (c *RuleRepository) List(user *model.User) (rules []*model.Rule, err error) {
	for _, rule := range c.ruleCache {
		err = c.prepare(rule, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare rule")
			return
		}
		rules = append(rules, rule)
	}
	return
}

func (c *RuleRepository) prepare(rule *model.Rule, user *model.User) (err error) {
	rule.Scope = c.scope(rule)
	if rule.ValueFloat, err = strconv.ParseFloat(rule.Value, 10); err != nil {
		//error handling for float64 for now is blank
	}
	if rule.ValueInt, err = strconv.ParseInt(rule.Value, 64, 10); err != nil {
		//error handling for int64 for now is blank
	}

	return
}

//Scope is what scope a rule is for
func (c *RuleRepository) scope(rule *model.Rule) string {
	scope := ""
	if strings.Contains(rule.Name, ":") {
		scope = rule.Name[0:strings.Index(rule.Name, ":")]
	}
	return scope
}

//newSchema handler
func (c *RuleRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
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

//getSchemaProperty handler
func (c *RuleRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "shortName":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
