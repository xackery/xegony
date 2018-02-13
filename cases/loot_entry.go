package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListLootEntry lists all lootEntrys accessible by provided user
func ListLootEntry(page *model.Page, loot *model.Loot, user *model.User) (lootEntrys []*model.LootEntry, err error) {
	err = validateOrderByLootEntryField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("lootEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for lootEntry")
		return
	}

	page.Total, err = reader.ListLootEntryTotalCount(loot)
	if err != nil {
		err = errors.Wrap(err, "failed to list lootEntry total count")
		return
	}

	lootEntrys, err = reader.ListLootEntry(page, loot)
	if err != nil {
		err = errors.Wrap(err, "failed to list lootEntry")
		return
	}
	for i, lootEntry := range lootEntrys {
		err = sanitizeLootEntry(lootEntry, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize lootEntry element %d", i)
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

//ListLootEntryBySearch will request any lootEntry matching the pattern of name
func ListLootEntryBySearch(page *model.Page, loot *model.Loot, lootEntry *model.LootEntry, user *model.User) (lootEntrys []*model.LootEntry, err error) {

	err = validateOrderByLootEntryField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareLootEntry(lootEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre lootEntry")
		return
	}

	err = validateLootEntry(lootEntry, nil, []string{ //optional
		"lootDropID",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootEntry")
		return
	}
	reader, err := getReader("lootEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get lootEntry reader")
		return
	}

	lootEntrys, err = reader.ListLootEntryBySearch(page, loot, lootEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to list lootEntry by search")
		return
	}

	page.Total, err = reader.ListLootEntryBySearchTotalCount(loot, lootEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to list lootEntry by search total count")
		return
	}
	for _, lootEntry := range lootEntrys {
		err = sanitizeLootEntry(lootEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize lootEntry")
			return
		}
	}

	err = sanitizeLootEntry(lootEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search lootEntry")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateLootEntry will create an lootEntry using provided information
func CreateLootEntry(loot *model.Loot, lootEntry *model.LootEntry, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list lootEntry by search without guide+")
		return
	}
	err = prepareLootEntry(lootEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare lootEntry")
		return
	}

	err = validateLootEntry(lootEntry, []string{"lootDropID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootEntry")
		return
	}
	lootEntry.LootID = loot.ID
	//lootEntry.TimeCreation = time.Now().Unix()
	writer, err := getWriter("lootEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for lootEntry")
		return
	}
	err = writer.CreateLootEntry(loot, lootEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to create lootEntry")
		return
	}
	err = sanitizeLootEntry(lootEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize lootEntry")
		return
	}
	return
}

//GetLootEntry gets an lootEntry by provided lootEntryID
func GetLootEntry(loot *model.Loot, lootEntry *model.LootEntry, user *model.User) (err error) {
	err = prepareLootEntry(lootEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare lootEntry")
		return
	}

	err = validateLootEntry(lootEntry, []string{"lootDropID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootEntry")
		return
	}

	reader, err := getReader("lootEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get lootEntry reader")
		return
	}

	err = reader.GetLootEntry(loot, lootEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to get lootEntry")
		return
	}

	err = sanitizeLootEntry(lootEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize lootEntry")
		return
	}

	return
}

//EditLootEntry edits an existing lootEntry
func EditLootEntry(loot *model.Loot, lootEntry *model.LootEntry, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list lootEntry by search without guide+")
		return
	}
	err = prepareLootEntry(lootEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare lootEntry")
		return
	}

	err = validateLootEntry(lootEntry,
		[]string{"lootDropID"}, //required
		[]string{               //optional
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootEntry")
		return
	}
	writer, err := getWriter("lootEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for lootEntry")
		return
	}
	err = writer.EditLootEntry(loot, lootEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to edit lootEntry")
		return
	}
	err = sanitizeLootEntry(lootEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize lootEntry")
		return
	}
	return
}

//DeleteLootEntry deletes an lootEntry by provided lootEntryID
func DeleteLootEntry(lootEntry *model.LootEntry, loot *model.Loot, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete lootEntry without admin+")
		return
	}
	err = prepareLootEntry(lootEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare lootEntry")
		return
	}

	err = validateLootEntry(lootEntry, []string{"lootDropID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate lootEntry")
		return
	}
	writer, err := getWriter("lootEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get lootEntry writer")
		return
	}
	err = writer.DeleteLootEntry(loot, lootEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to delete lootEntry")
		return
	}
	return
}

func prepareLootEntry(lootEntry *model.LootEntry, user *model.User) (err error) {
	if lootEntry == nil {
		err = fmt.Errorf("empty lootEntry")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateLootEntry(lootEntry *model.LootEntry, required []string, optional []string) (err error) {
	schema, err := newSchemaLootEntry(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootEntry))
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

func validateOrderByLootEntryField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "lootdrop_id"
	}

	validNames := []string{
		"lootdrop_id",
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

func sanitizeLootEntry(lootEntry *model.LootEntry, user *model.User) (err error) {
	return
}

func newSchemaLootEntry(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyLootEntry(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyLootEntry(field); err != nil {
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

func getSchemaPropertyLootEntry(field string) (prop model.Schema, err error) {
	switch field {
	case "lootDropID":
		prop.Type = "integer"
		prop.Minimum = 0
	case "lootID":
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
