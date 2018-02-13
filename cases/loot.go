package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListLoot lists all loots accessible by provided user
func ListLoot(page *model.Page, user *model.User) (loots []*model.Loot, err error) {
	err = validateOrderByLootField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("loot")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for loot")
		return
	}

	loots, err = reader.ListLoot(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list loot")
		return
	}
	for i, loot := range loots {
		err = sanitizeLoot(loot, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize loot element %d", i)
			return
		}
	}

	page.Total, err = reader.ListLootTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list loot toal count")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}

	return
}

//ListLootBySearch will request any loot matching the pattern of name
func ListLootBySearch(page *model.Page, loot *model.Loot, user *model.User) (loots []*model.Loot, err error) {

	err = validateOrderByLootField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre loot")
		return
	}

	err = validateLoot(loot, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate loot")
		return
	}
	reader, err := getReader("loot")
	if err != nil {
		err = errors.Wrap(err, "failed to get loot reader")
		return
	}

	loots, err = reader.ListLootBySearch(page, loot)
	if err != nil {
		err = errors.Wrap(err, "failed to list loot by search")
		return
	}

	page.Total, err = reader.ListLootBySearchTotalCount(loot)
	if err != nil {
		err = errors.Wrap(err, "failed to get page total")
		return
	}

	for _, loot := range loots {
		err = sanitizeLoot(loot, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize loot")
			return
		}
	}

	err = sanitizeLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search loot")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateLoot will create an loot using provided information
func CreateLoot(loot *model.Loot, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list loot by search without guide+")
		return
	}
	err = prepareLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare loot")
		return
	}

	err = validateLoot(loot, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate loot")
		return
	}
	loot.ID = 0
	//loot.TimeCreation = time.Now().Unix()
	writer, err := getWriter("loot")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for loot")
		return
	}
	err = writer.CreateLoot(loot)
	if err != nil {
		err = errors.Wrap(err, "failed to create loot")
		return
	}

	memWriter, err := getWriter("loot")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for loot")
		return
	}
	err = memWriter.CreateLoot(loot)
	if err != nil {
		err = errors.Wrap(err, "failed to edit loot")
		return
	}

	err = sanitizeLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize loot")
		return
	}
	return
}

//GetLoot gets an loot by provided lootID
func GetLoot(loot *model.Loot, user *model.User) (err error) {
	err = prepareLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare loot")
		return
	}

	err = validateLoot(loot, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate loot")
		return
	}

	reader, err := getReader("loot")
	if err != nil {
		err = errors.Wrap(err, "failed to get loot reader")
		return
	}

	err = reader.GetLoot(loot)
	if err != nil {
		err = errors.Wrap(err, "failed to get loot")
		return
	}

	err = sanitizeLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize loot")
		return
	}

	return
}

//EditLoot edits an existing loot
func EditLoot(loot *model.Loot, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list loot by search without guide+")
		return
	}
	err = prepareLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare loot")
		return
	}

	err = validateLoot(loot,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate loot")
		return
	}
	writer, err := getWriter("loot")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for loot")
		return
	}
	err = writer.EditLoot(loot)
	if err != nil {
		err = errors.Wrap(err, "failed to edit loot")
		return
	}

	memWriter, err := getWriter("loot")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for loot")
		return
	}
	err = memWriter.EditLoot(loot)
	if err != nil {
		err = errors.Wrap(err, "failed to edit loot")
		return
	}

	err = sanitizeLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize loot")
		return
	}
	return
}

//DeleteLoot deletes an loot by provided lootID
func DeleteLoot(loot *model.Loot, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete loot without admin+")
		return
	}
	err = prepareLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare loot")
		return
	}

	err = validateLoot(loot, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate loot")
		return
	}
	writer, err := getWriter("loot")
	if err != nil {
		err = errors.Wrap(err, "failed to get loot writer")
		return
	}
	err = writer.DeleteLoot(loot)
	if err != nil {
		err = errors.Wrap(err, "failed to delete loot")
		return
	}

	memWriter, err := getWriter("loot")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for loot")
		return
	}
	err = memWriter.DeleteLoot(loot)
	if err != nil {
		err = errors.Wrap(err, "failed to delete loot")
		return
	}
	return
}

func prepareLoot(loot *model.Loot, user *model.User) (err error) {
	if loot == nil {
		err = fmt.Errorf("empty loot")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateLoot(loot *model.Loot, required []string, optional []string) (err error) {
	schema, err := newSchemaLoot(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(loot))
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

func validateOrderByLootField(page *model.Page) (err error) {
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

func sanitizeLoot(loot *model.Loot, user *model.User) (err error) {
	if len(loot.Name) == 0 {
		loot.Name = fmt.Sprintf("(%d)", loot.ID)
	}

	return
}

func newSchemaLoot(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyLoot(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyLoot(field); err != nil {
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

func getSchemaPropertyLoot(field string) (prop model.Schema, err error) {
	switch field {

	case "ID":
		prop.Type = "integer"
		prop.Minimum = 0
	case "name":
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
