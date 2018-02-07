package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListSpawn lists all spawns accessible by provided user
func ListSpawn(page *model.Page, user *model.User) (spawns []*model.Spawn, err error) {
	err = validateOrderBySpawnField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("spawn")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for spawn")
		return
	}

	spawns, err = reader.ListSpawn(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spawn")
		return
	}
	for i, spawn := range spawns {
		err = sanitizeSpawn(spawn, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize spawn element %d", i)
			return
		}
	}

	page.Total, err = reader.ListSpawnTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list spawn toal count")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}

	return
}

//ListSpawnBySearch will request any spawn matching the pattern of name
func ListSpawnBySearch(page *model.Page, spawn *model.Spawn, user *model.User) (spawns []*model.Spawn, err error) {

	err = validateOrderBySpawnField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre spawn")
		return
	}

	err = validateSpawn(spawn, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawn")
		return
	}
	reader, err := getReader("spawn")
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn reader")
		return
	}

	spawns, err = reader.ListSpawnBySearch(page, spawn)
	if err != nil {
		err = errors.Wrap(err, "failed to list spawn by search")
		return
	}

	page.Total, err = reader.ListSpawnBySearchTotalCount(spawn)
	if err != nil {
		err = errors.Wrap(err, "failed to get page total")
		return
	}

	for _, spawn := range spawns {
		err = sanitizeSpawn(spawn, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize spawn")
			return
		}
	}

	err = sanitizeSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search spawn")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSpawn will create an spawn using provided information
func CreateSpawn(spawn *model.Spawn, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spawn by search without guide+")
		return
	}
	err = prepareSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawn")
		return
	}

	err = validateSpawn(spawn, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawn")
		return
	}
	spawn.ID = 0
	//spawn.TimeCreation = time.Now().Unix()
	writer, err := getWriter("spawn")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spawn")
		return
	}
	err = writer.CreateSpawn(spawn)
	if err != nil {
		err = errors.Wrap(err, "failed to create spawn")
		return
	}

	memWriter, err := getWriter("spawn")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spawn")
		return
	}
	err = memWriter.CreateSpawn(spawn)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spawn")
		return
	}

	err = sanitizeSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spawn")
		return
	}
	return
}

//GetSpawn gets an spawn by provided spawnID
func GetSpawn(spawn *model.Spawn, user *model.User) (err error) {
	err = prepareSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawn")
		return
	}

	err = validateSpawn(spawn, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawn")
		return
	}

	reader, err := getReader("spawn")
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn reader")
		return
	}

	err = reader.GetSpawn(spawn)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn")
		return
	}

	err = sanitizeSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spawn")
		return
	}

	return
}

//EditSpawn edits an existing spawn
func EditSpawn(spawn *model.Spawn, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spawn by search without guide+")
		return
	}
	err = prepareSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawn")
		return
	}

	err = validateSpawn(spawn,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawn")
		return
	}
	writer, err := getWriter("spawn")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spawn")
		return
	}
	err = writer.EditSpawn(spawn)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spawn")
		return
	}

	memWriter, err := getWriter("spawn")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spawn")
		return
	}
	err = memWriter.EditSpawn(spawn)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spawn")
		return
	}

	err = sanitizeSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spawn")
		return
	}
	return
}

//DeleteSpawn deletes an spawn by provided spawnID
func DeleteSpawn(spawn *model.Spawn, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete spawn without admin+")
		return
	}
	err = prepareSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawn")
		return
	}

	err = validateSpawn(spawn, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawn")
		return
	}
	writer, err := getWriter("spawn")
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn writer")
		return
	}
	err = writer.DeleteSpawn(spawn)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spawn")
		return
	}

	memWriter, err := getWriter("spawn")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spawn")
		return
	}
	err = memWriter.DeleteSpawn(spawn)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spawn")
		return
	}
	return
}

func prepareSpawn(spawn *model.Spawn, user *model.User) (err error) {
	if spawn == nil {
		err = fmt.Errorf("empty spawn")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSpawn(spawn *model.Spawn, required []string, optional []string) (err error) {
	schema, err := newSchemaSpawn(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spawn))
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

func validateOrderBySpawnField(page *model.Page) (err error) {
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

func sanitizeSpawn(spawn *model.Spawn, user *model.User) (err error) {
	if len(spawn.Name) == 0 {
		spawn.Name = fmt.Sprintf("(%d)", spawn.ID)
	}

	return
}

func newSchemaSpawn(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySpawn(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySpawn(field); err != nil {
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

func getSchemaPropertySpawn(field string) (prop model.Schema, err error) {
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
