package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListLootDropEntry lists all lootDropEntrys accessible by provided user
func ListLootDropEntry(page *model.Page, lootDrop *model.LootDrop, user *model.User) (lootDropEntrys []*model.LootDropEntry, err error) {
	err = validateOrderByLootDropEntryField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("lootDropEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for lootDropEntry")
		return
	}

	page.Total, err = reader.ListLootDropEntryTotalCount(lootDrop)
	if err != nil {
		err = errors.Wrap(err, "failed to list lootDropEntry total count")
		return
	}

	lootDropEntrys, err = reader.ListLootDropEntry(page, lootDrop)
	if err != nil {
		err = errors.Wrap(err, "failed to list lootDropEntry")
		return
	}
	for i, lootDropEntry := range lootDropEntrys {
		err = sanitizeLootDropEntry(lootDropEntry, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize lootDropEntry element %d", i)
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

//ListLootDropEntryBySearch will request any lootDropEntry matching the pattern of name
func ListLootDropEntryBySearch(page *model.Page, lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry, user *model.User) (lootDropEntrys []*model.LootDropEntry, err error) {

	err = validateOrderByLootDropEntryField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareLootDropEntry(lootDropEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre lootDropEntry")
		return
	}

	err = validateLootDropEntry(lootDropEntry, nil, []string{ //optional
		"lootDropDropID",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootDropEntry")
		return
	}
	reader, err := getReader("lootDropEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get lootDropEntry reader")
		return
	}

	lootDropEntrys, err = reader.ListLootDropEntryBySearch(page, lootDrop, lootDropEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to list lootDropEntry by search")
		return
	}

	page.Total, err = reader.ListLootDropEntryBySearchTotalCount(lootDrop, lootDropEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to list lootDropEntry by search total count")
		return
	}
	for _, lootDropEntry := range lootDropEntrys {
		err = sanitizeLootDropEntry(lootDropEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize lootDropEntry")
			return
		}
	}

	err = sanitizeLootDropEntry(lootDropEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search lootDropEntry")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateLootDropEntry will create an lootDropEntry using provided information
func CreateLootDropEntry(lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list lootDropEntry by search without guide+")
		return
	}
	err = prepareLootDropEntry(lootDropEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare lootDropEntry")
		return
	}

	err = validateLootDropEntry(lootDropEntry, []string{"lootDropDropID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootDropEntry")
		return
	}
	lootDropEntry.LootDropID = lootDrop.ID
	//lootDropEntry.TimeCreation = time.Now().Unix()
	writer, err := getWriter("lootDropEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for lootDropEntry")
		return
	}
	err = writer.CreateLootDropEntry(lootDrop, lootDropEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to create lootDropEntry")
		return
	}
	err = sanitizeLootDropEntry(lootDropEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize lootDropEntry")
		return
	}
	return
}

//GetLootDropEntry gets an lootDropEntry by provided lootDropEntryID
func GetLootDropEntry(lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry, user *model.User) (err error) {
	err = prepareLootDropEntry(lootDropEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare lootDropEntry")
		return
	}

	err = validateLootDropEntry(lootDropEntry, []string{"lootDropDropID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootDropEntry")
		return
	}

	reader, err := getReader("lootDropEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get lootDropEntry reader")
		return
	}

	err = reader.GetLootDropEntry(lootDrop, lootDropEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to get lootDropEntry")
		return
	}

	err = sanitizeLootDropEntry(lootDropEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize lootDropEntry")
		return
	}

	return
}

//EditLootDropEntry edits an existing lootDropEntry
func EditLootDropEntry(lootDrop *model.LootDrop, lootDropEntry *model.LootDropEntry, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list lootDropEntry by search without guide+")
		return
	}
	err = prepareLootDropEntry(lootDropEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare lootDropEntry")
		return
	}

	err = validateLootDropEntry(lootDropEntry,
		[]string{"lootDropDropID"}, //required
		[]string{                   //optional
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootDropEntry")
		return
	}
	writer, err := getWriter("lootDropEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for lootDropEntry")
		return
	}
	err = writer.EditLootDropEntry(lootDrop, lootDropEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to edit lootDropEntry")
		return
	}
	err = sanitizeLootDropEntry(lootDropEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize lootDropEntry")
		return
	}
	return
}

//DeleteLootDropEntry deletes an lootDropEntry by provided lootDropEntryID
func DeleteLootDropEntry(lootDropEntry *model.LootDropEntry, lootDrop *model.LootDrop, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete lootDropEntry without admin+")
		return
	}
	err = prepareLootDropEntry(lootDropEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare lootDropEntry")
		return
	}

	err = validateLootDropEntry(lootDropEntry, []string{"lootDropDropID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootDropEntry")
		return
	}
	writer, err := getWriter("lootDropEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get lootDropEntry writer")
		return
	}
	err = writer.DeleteLootDropEntry(lootDrop, lootDropEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to delete lootDropEntry")
		return
	}
	return
}

func prepareLootDropEntry(lootDropEntry *model.LootDropEntry, user *model.User) (err error) {
	if lootDropEntry == nil {
		err = fmt.Errorf("empty lootDropEntry")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateLootDropEntry(lootDropEntry *model.LootDropEntry, required []string, optional []string) (err error) {
	schema, err := newSchemaLootDropEntry(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootDropEntry))
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

func validateOrderByLootDropEntryField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "item_id"
	}

	validNames := []string{
		"item_id",
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

func sanitizeLootDropEntry(lootDropEntry *model.LootDropEntry, user *model.User) (err error) {
	lootDropEntry.Item = &model.Item{
		ID: lootDropEntry.ItemID,
	}
	err = GetItem(lootDropEntry.Item, user)
	if err != nil {
		err = nil
	}

	return
}

func newSchemaLootDropEntry(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyLootDropEntry(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyLootDropEntry(field); err != nil {
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

func getSchemaPropertyLootDropEntry(field string) (prop model.Schema, err error) {
	switch field {
	case "lootDropDropID":
		prop.Type = "integer"
		prop.Minimum = 0
	case "lootDropID":
		prop.Type = "integer"
		prop.Minimum = 0
	case "chance":
		prop.Type = "integer"
		prop.Minimum = 0
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
