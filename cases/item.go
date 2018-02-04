package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListItem lists all items accessible by provided user
func ListItem(page *model.Page, user *model.User) (items []*model.Item, err error) {
	err = validateOrderByItemField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("item")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for item")
		return
	}

	page.Total, err = reader.ListItemTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list item toal count")
		return
	}

	items, err = reader.ListItem(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list item")
		return
	}
	for i, item := range items {
		err = sanitizeItem(item, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize item element %d", i)
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

//ListItemBySearch will request any item matching the pattern of name
func ListItemBySearch(page *model.Page, item *model.Item, user *model.User) (items []*model.Item, err error) {

	err = validateOrderByItemField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre item")
		return
	}

	err = validateItem(item, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate item")
		return
	}
	reader, err := getReader("item")
	if err != nil {
		err = errors.Wrap(err, "failed to get item reader")
		return
	}

	items, err = reader.ListItemBySearch(page, item)
	if err != nil {
		err = errors.Wrap(err, "failed to list item by search")
		return
	}

	for _, item := range items {
		err = sanitizeItem(item, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize item")
			return
		}
	}

	err = sanitizeItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search item")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateItem will create an item using provided information
func CreateItem(item *model.Item, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list item by search without guide+")
		return
	}
	err = prepareItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare item")
		return
	}

	err = validateItem(item, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate item")
		return
	}
	item.ID = 0
	writer, err := getWriter("item")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for item")
		return
	}
	err = writer.CreateItem(item)
	if err != nil {
		err = errors.Wrap(err, "failed to create item")
		return
	}
	err = sanitizeItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize item")
		return
	}
	return
}

//GetItem gets an item by provided itemID
func GetItem(item *model.Item, user *model.User) (err error) {
	err = prepareItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare item")
		return
	}

	err = validateItem(item, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate item")
		return
	}

	reader, err := getReader("item")
	if err != nil {
		err = errors.Wrap(err, "failed to get item reader")
		return
	}

	err = reader.GetItem(item)
	if err != nil {
		err = errors.Wrap(err, "failed to get item")
		return
	}

	err = sanitizeItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize item")
		return
	}

	return
}

//EditItem edits an existing item
func EditItem(item *model.Item, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list item by search without guide+")
		return
	}
	err = prepareItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare item")
		return
	}

	err = validateItem(item,
		[]string{"ID"}, //required
		[]string{       //optional
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate item")
		return
	}
	writer, err := getWriter("item")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for item")
		return
	}
	err = writer.EditItem(item)
	if err != nil {
		err = errors.Wrap(err, "failed to edit item")
		return
	}
	err = sanitizeItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize item")
		return
	}
	return
}

//DeleteItem deletes an item by provided itemID
func DeleteItem(item *model.Item, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete item without admin+")
		return
	}
	err = prepareItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare item")
		return
	}

	err = validateItem(item, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate item")
		return
	}
	writer, err := getWriter("item")
	if err != nil {
		err = errors.Wrap(err, "failed to get item writer")
		return
	}
	err = writer.DeleteItem(item)
	if err != nil {
		err = errors.Wrap(err, "failed to delete item")
		return
	}
	return
}

func prepareItem(item *model.Item, user *model.User) (err error) {
	if item == nil {
		err = fmt.Errorf("empty item")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateItem(item *model.Item, required []string, optional []string) (err error) {
	schema, err := newSchemaItem(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(item))
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

func validateOrderByItemField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "id"
	}

	validNames := []string{
		"id",
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

func sanitizeItem(item *model.Item, user *model.User) (err error) {
	if item.ProcEffectSpellID > 0 {
		item.ProcEffect = &model.Spell{
			ID: item.ProcEffectSpellID,
		}
		err = GetSpell(item.ProcEffect, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get spell")
			return
		}
	}
	/*class := &model.Class{
		Bit: item.ClassBit,
	}
	item.Classs, err = ListClassByBit(class, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get class from item.ClassID")
		return
	}

	race := &model.Race{
		item.RaceBit,
	}
	item.Races, err = ListRaceByBit(race, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get race from item.RaceID")
		return
	}*/
	return
}

func newSchemaItem(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyItem(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyItem(field); err != nil {
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

func getSchemaPropertyItem(field string) (prop model.Schema, err error) {
	switch field {
	case "ID": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
