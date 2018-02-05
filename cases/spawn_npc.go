package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListSpawnNpc lists all spawnNpcs accessible by provided user
func ListSpawnNpc(page *model.Page, spawn *model.Spawn, user *model.User) (spawnNpcs []*model.SpawnNpc, err error) {
	err = validateOrderBySpawnNpcField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("spawnNpc")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for spawnNpc")
		return
	}

	page.Total, err = reader.ListSpawnNpcTotalCount(spawn)
	if err != nil {
		err = errors.Wrap(err, "failed to list spawnNpc total count")
		return
	}

	spawnNpcs, err = reader.ListSpawnNpc(page, spawn)
	if err != nil {
		err = errors.Wrap(err, "failed to list spawnNpc")
		return
	}
	for i, spawnNpc := range spawnNpcs {
		err = sanitizeSpawnNpc(spawnNpc, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize spawnNpc element %d", i)
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

//ListSpawnNpcBySearch will request any spawnNpc matching the pattern of name
func ListSpawnNpcBySearch(page *model.Page, spawn *model.Spawn, spawnNpc *model.SpawnNpc, user *model.User) (spawnNpcs []*model.SpawnNpc, err error) {

	err = validateOrderBySpawnNpcField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSpawnNpc(spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre spawnNpc")
		return
	}

	err = validateSpawnNpc(spawnNpc, nil, []string{ //optional
		"npcID",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawnNpc")
		return
	}
	reader, err := getReader("spawnNpc")
	if err != nil {
		err = errors.Wrap(err, "failed to get spawnNpc reader")
		return
	}

	spawnNpcs, err = reader.ListSpawnNpcBySearch(page, spawn, spawnNpc)
	if err != nil {
		err = errors.Wrap(err, "failed to list spawnNpc by search")
		return
	}

	page.Total, err = reader.ListSpawnNpcBySearchTotalCount(spawn, spawnNpc)
	if err != nil {
		err = errors.Wrap(err, "failed to list spawnNpc by search total count")
		return
	}
	for _, spawnNpc := range spawnNpcs {
		err = sanitizeSpawnNpc(spawnNpc, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize spawnNpc")
			return
		}
	}

	err = sanitizeSpawnNpc(spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search spawnNpc")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSpawnNpc will create an spawnNpc using provided information
func CreateSpawnNpc(spawn *model.Spawn, spawnNpc *model.SpawnNpc, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spawnNpc by search without guide+")
		return
	}
	err = prepareSpawnNpc(spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnNpc")
		return
	}

	err = validateSpawnNpc(spawnNpc, []string{"npcID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawnNpc")
		return
	}
	spawnNpc.SpawnID = spawn.ID
	//spawnNpc.TimeCreation = time.Now().Unix()
	writer, err := getWriter("spawnNpc")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spawnNpc")
		return
	}
	err = writer.CreateSpawnNpc(spawn, spawnNpc)
	if err != nil {
		err = errors.Wrap(err, "failed to create spawnNpc")
		return
	}
	err = sanitizeSpawnNpc(spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spawnNpc")
		return
	}
	return
}

//GetSpawnNpc gets an spawnNpc by provided spawnNpcID
func GetSpawnNpc(spawn *model.Spawn, spawnNpc *model.SpawnNpc, user *model.User) (err error) {
	err = prepareSpawnNpc(spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnNpc")
		return
	}

	err = validateSpawnNpc(spawnNpc, []string{"npcID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawnNpc")
		return
	}

	reader, err := getReader("spawnNpc")
	if err != nil {
		err = errors.Wrap(err, "failed to get spawnNpc reader")
		return
	}

	err = reader.GetSpawnNpc(spawn, spawnNpc)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawnNpc")
		return
	}

	err = sanitizeSpawnNpc(spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spawnNpc")
		return
	}

	return
}

//EditSpawnNpc edits an existing spawnNpc
func EditSpawnNpc(spawn *model.Spawn, spawnNpc *model.SpawnNpc, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spawnNpc by search without guide+")
		return
	}
	err = prepareSpawnNpc(spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnNpc")
		return
	}

	err = validateSpawnNpc(spawnNpc,
		[]string{"npcID"}, //required
		[]string{          //optional
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawnNpc")
		return
	}
	writer, err := getWriter("spawnNpc")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spawnNpc")
		return
	}
	err = writer.EditSpawnNpc(spawn, spawnNpc)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spawnNpc")
		return
	}
	err = sanitizeSpawnNpc(spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spawnNpc")
		return
	}
	return
}

//DeleteSpawnNpc deletes an spawnNpc by provided spawnNpcID
func DeleteSpawnNpc(spawnNpc *model.SpawnNpc, spawn *model.Spawn, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete spawnNpc without admin+")
		return
	}
	err = prepareSpawnNpc(spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spawnNpc")
		return
	}

	err = validateSpawnNpc(spawnNpc, []string{"npcID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spawnNpc")
		return
	}
	writer, err := getWriter("spawnNpc")
	if err != nil {
		err = errors.Wrap(err, "failed to get spawnNpc writer")
		return
	}
	err = writer.DeleteSpawnNpc(spawn, spawnNpc)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spawnNpc")
		return
	}
	return
}

func prepareSpawnNpc(spawnNpc *model.SpawnNpc, user *model.User) (err error) {
	if spawnNpc == nil {
		err = fmt.Errorf("empty spawnNpc")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSpawnNpc(spawnNpc *model.SpawnNpc, required []string, optional []string) (err error) {
	schema, err := newSchemaSpawnNpc(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spawnNpc))
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

func validateOrderBySpawnNpcField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "npcid"
	}

	validNames := []string{
		"npcid",
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

func sanitizeSpawnNpc(spawnNpc *model.SpawnNpc, user *model.User) (err error) {

	spawnNpc.Npc = &model.Npc{
		ID: spawnNpc.NpcID,
	}

	err = GetNpc(spawnNpc.Npc, user)
	if err != nil {
		err = errors.Wrapf(err, "failed to get npc %d", spawnNpc.NpcID)
		return
	}
	return
}

func newSchemaSpawnNpc(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySpawnNpc(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySpawnNpc(field); err != nil {
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

func getSchemaPropertySpawnNpc(field string) (prop model.Schema, err error) {
	switch field {
	case "npcID":
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
