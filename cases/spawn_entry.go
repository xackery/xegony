package cases

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListSpawnEntry lists all spawnEntrys accessible by provided user
func ListSpawnEntry(page *model.Page, spawn *model.Spawn, user *model.User) (spawnEntrys []*model.SpawnEntry, err error) {
	err = validateOrderBySpawnEntryField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("spawnEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for spawnEntry")
		return
	}

	page.Total, err = reader.ListSpawnEntryTotalCount(spawn)
	if err != nil {
		err = errors.Wrap(err, "failed to list spawnEntry total count")
		return
	}

	spawnEntrys, err = reader.ListSpawnEntry(page, spawn)
	if err != nil {
		err = errors.Wrap(err, "failed to list spawnEntry")
		return
	}
	for i, spawnEntry := range spawnEntrys {
		err = sanitizeSpawnEntry(spawnEntry, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize spawnEntry element %d", i)
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

//ListSpawnEntryBySearch will request any spawnEntry matching the pattern of name
func ListSpawnEntryBySearch(page *model.Page, spawn *model.Spawn, spawnEntry *model.SpawnEntry, user *model.User) (spawnEntrys []*model.SpawnEntry, err error) {

	err = validateOrderBySpawnEntryField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSpawnEntry(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre spawnEntry")
		return
	}

	err = validateSpawnEntry(spawnEntry, nil, []string{ //optional
		"entryID",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawnEntry")
		return
	}
	reader, err := getReader("spawnEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get spawnEntry reader")
		return
	}

	spawnEntrys, err = reader.ListSpawnEntryBySearch(page, spawn, spawnEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to list spawnEntry by search")
		return
	}

	page.Total, err = reader.ListSpawnEntryBySearchTotalCount(spawn, spawnEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to list spawnEntry by search total count")
		return
	}
	for _, spawnEntry := range spawnEntrys {
		err = sanitizeSpawnEntry(spawnEntry, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize spawnEntry")
			return
		}
	}

	err = sanitizeSpawnEntry(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search spawnEntry")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSpawnEntry will create an spawnEntry using provided information
func CreateSpawnEntry(spawn *model.Spawn, spawnEntry *model.SpawnEntry, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spawnEntry by search without guide+")
		return
	}
	err = prepareSpawnEntry(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnEntry")
		return
	}

	err = validateSpawnEntry(spawnEntry, []string{"entryID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawnEntry")
		return
	}
	spawnEntry.SpawnID = spawn.ID
	//spawnEntry.TimeCreation = time.Now().Unix()
	writer, err := getWriter("spawnEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spawnEntry")
		return
	}
	err = writer.CreateSpawnEntry(spawn, spawnEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to create spawnEntry")
		return
	}
	err = sanitizeSpawnEntry(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spawnEntry")
		return
	}
	return
}

//GetSpawnEntry gets an spawnEntry by provided spawnEntryID
func GetSpawnEntry(spawn *model.Spawn, spawnEntry *model.SpawnEntry, user *model.User) (err error) {
	err = prepareSpawnEntry(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnEntry")
		return
	}

	err = validateSpawnEntry(spawnEntry, []string{"entryID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawnEntry")
		return
	}

	reader, err := getReader("spawnEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get spawnEntry reader")
		return
	}

	err = reader.GetSpawnEntry(spawn, spawnEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawnEntry")
		return
	}

	err = sanitizeSpawnEntry(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spawnEntry")
		return
	}

	return
}

//EditSpawnEntry edits an existing spawnEntry
func EditSpawnEntry(spawn *model.Spawn, spawnEntry *model.SpawnEntry, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spawnEntry by search without guide+")
		return
	}
	err = prepareSpawnEntry(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnEntry")
		return
	}

	err = validateSpawnEntry(spawnEntry,
		[]string{"entryID"}, //required
		[]string{            //optional
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawnEntry")
		return
	}
	writer, err := getWriter("spawnEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spawnEntry")
		return
	}
	err = writer.EditSpawnEntry(spawn, spawnEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spawnEntry")
		return
	}
	err = sanitizeSpawnEntry(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spawnEntry")
		return
	}
	return
}

//DeleteSpawnEntry deletes an spawnEntry by provided spawnEntryID
func DeleteSpawnEntry(spawnEntry *model.SpawnEntry, spawn *model.Spawn, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete spawnEntry without admin+")
		return
	}
	err = prepareSpawnEntry(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnEntry")
		return
	}

	err = validateSpawnEntry(spawnEntry, []string{"entryID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawnEntry")
		return
	}
	writer, err := getWriter("spawnEntry")
	if err != nil {
		err = errors.Wrap(err, "failed to get spawnEntry writer")
		return
	}
	err = writer.DeleteSpawnEntry(spawn, spawnEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spawnEntry")
		return
	}
	return
}

func prepareSpawnEntry(spawnEntry *model.SpawnEntry, user *model.User) (err error) {
	if spawnEntry == nil {
		err = fmt.Errorf("empty spawnEntry")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSpawnEntry(spawnEntry *model.SpawnEntry, required []string, optional []string) (err error) {
	schema, err := newSchemaSpawnEntry(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spawnEntry))
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

func validateOrderBySpawnEntryField(page *model.Page) (err error) {
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

func sanitizeSpawnEntry(spawnEntry *model.SpawnEntry, user *model.User) (err error) {
	if len(spawnEntry.ZoneShortName.String) > 0 {
		spawnEntry.Zone = &model.Zone{}

		spawnEntry.Zone.ShortName.String = strings.ToLower(spawnEntry.ZoneShortName.String)
		spawnEntry.Zone.ShortName.Valid = true

		err = GetZoneByShortName(spawnEntry.Zone, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get zone by shortname")
			return
		}
	}
	return
}

func newSchemaSpawnEntry(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySpawnEntry(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySpawnEntry(field); err != nil {
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

func getSchemaPropertySpawnEntry(field string) (prop model.Schema, err error) {
	switch field {
	case "entryID":
		prop.Type = "integer"
		prop.Minimum = 0
	case "spawnID":
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
