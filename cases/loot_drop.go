package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListLootDrop lists all lootDrops accessible by provided user
func ListLootDrop(page *model.Page, user *model.User) (lootDrops []*model.LootDrop, err error) {
	err = validateOrderByLootDropField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("lootDrop")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for lootDrop")
		return
	}

	lootDrops, err = reader.ListLootDrop(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list lootDrop")
		return
	}
	for i, lootDrop := range lootDrops {
		err = sanitizeLootDrop(lootDrop, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize lootDrop element %d", i)
			return
		}
	}

	page.Total, err = reader.ListLootDropTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list lootDrop toal count")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}

	return
}

//ListLootDropBySearch will request any lootDrop matching the pattern of name
func ListLootDropBySearch(page *model.Page, lootDrop *model.LootDrop, user *model.User) (lootDrops []*model.LootDrop, err error) {

	err = validateOrderByLootDropField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre lootDrop")
		return
	}

	err = validateLootDrop(lootDrop, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootDrop")
		return
	}
	reader, err := getReader("lootDrop")
	if err != nil {
		err = errors.Wrap(err, "failed to get lootDrop reader")
		return
	}

	lootDrops, err = reader.ListLootDropBySearch(page, lootDrop)
	if err != nil {
		err = errors.Wrap(err, "failed to list lootDrop by search")
		return
	}

	page.Total, err = reader.ListLootDropBySearchTotalCount(lootDrop)
	if err != nil {
		err = errors.Wrap(err, "failed to get page total")
		return
	}

	for _, lootDrop := range lootDrops {
		err = sanitizeLootDrop(lootDrop, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize lootDrop")
			return
		}
	}

	err = sanitizeLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search lootDrop")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateLootDrop will create an lootDrop using provided information
func CreateLootDrop(lootDrop *model.LootDrop, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list lootDrop by search without guide+")
		return
	}
	err = prepareLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare lootDrop")
		return
	}

	err = validateLootDrop(lootDrop, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootDrop")
		return
	}
	lootDrop.ID = 0
	//lootDrop.TimeCreation = time.Now().Unix()
	writer, err := getWriter("lootDrop")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for lootDrop")
		return
	}
	err = writer.CreateLootDrop(lootDrop)
	if err != nil {
		err = errors.Wrap(err, "failed to create lootDrop")
		return
	}

	memWriter, err := getWriter("lootDrop")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for lootDrop")
		return
	}
	err = memWriter.CreateLootDrop(lootDrop)
	if err != nil {
		err = errors.Wrap(err, "failed to edit lootDrop")
		return
	}

	err = sanitizeLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize lootDrop")
		return
	}
	return
}

//GetLootDrop gets an lootDrop by provided lootDropID
func GetLootDrop(lootDrop *model.LootDrop, user *model.User) (err error) {
	err = prepareLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare lootDrop")
		return
	}

	err = validateLootDrop(lootDrop, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootDrop")
		return
	}

	reader, err := getReader("lootDrop")
	if err != nil {
		err = errors.Wrap(err, "failed to get lootDrop reader")
		return
	}

	err = reader.GetLootDrop(lootDrop)
	if err != nil {
		err = errors.Wrap(err, "failed to get lootDrop")
		return
	}

	err = sanitizeLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize lootDrop")
		return
	}

	return
}

//EditLootDrop edits an existing lootDrop
func EditLootDrop(lootDrop *model.LootDrop, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list lootDrop by search without guide+")
		return
	}
	err = prepareLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare lootDrop")
		return
	}

	err = validateLootDrop(lootDrop,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootDrop")
		return
	}
	writer, err := getWriter("lootDrop")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for lootDrop")
		return
	}
	err = writer.EditLootDrop(lootDrop)
	if err != nil {
		err = errors.Wrap(err, "failed to edit lootDrop")
		return
	}

	memWriter, err := getWriter("lootDrop")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for lootDrop")
		return
	}
	err = memWriter.EditLootDrop(lootDrop)
	if err != nil {
		err = errors.Wrap(err, "failed to edit lootDrop")
		return
	}

	err = sanitizeLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize lootDrop")
		return
	}
	return
}

//DeleteLootDrop deletes an lootDrop by provided lootDropID
func DeleteLootDrop(lootDrop *model.LootDrop, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete lootDrop without admin+")
		return
	}
	err = prepareLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare lootDrop")
		return
	}

	err = validateLootDrop(lootDrop, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootDrop")
		return
	}
	writer, err := getWriter("lootDrop")
	if err != nil {
		err = errors.Wrap(err, "failed to get lootDrop writer")
		return
	}
	err = writer.DeleteLootDrop(lootDrop)
	if err != nil {
		err = errors.Wrap(err, "failed to delete lootDrop")
		return
	}

	memWriter, err := getWriter("lootDrop")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for lootDrop")
		return
	}
	err = memWriter.DeleteLootDrop(lootDrop)
	if err != nil {
		err = errors.Wrap(err, "failed to delete lootDrop")
		return
	}
	return
}

func prepareLootDrop(lootDrop *model.LootDrop, user *model.User) (err error) {
	if lootDrop == nil {
		err = fmt.Errorf("empty lootDrop")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateLootDrop(lootDrop *model.LootDrop, required []string, optional []string) (err error) {
	schema, err := newSchemaLootDrop(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootDrop))
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

func validateOrderByLootDropField(page *model.Page) (err error) {
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

func sanitizeLootDrop(lootDrop *model.LootDrop, user *model.User) (err error) {
	if len(lootDrop.Name) == 0 {
		lootDrop.Name = fmt.Sprintf("(%d)", lootDrop.ID)
	}

	return
}

func newSchemaLootDrop(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyLootDrop(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyLootDrop(field); err != nil {
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

func getSchemaPropertyLootDrop(field string) (prop model.Schema, err error) {
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
