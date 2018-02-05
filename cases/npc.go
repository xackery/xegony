package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListNpc lists all npcs accessible by provided user
func ListNpc(page *model.Page, user *model.User) (npcs []*model.Npc, err error) {
	err = validateOrderByNpcField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("npc")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for npc")
		return
	}

	npcs, err = reader.ListNpc(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list npc")
		return
	}
	for i, npc := range npcs {
		err = sanitizeNpc(npc, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize npc element %d", i)
			return
		}
	}

	page.Total, err = reader.ListNpcTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list npc toal count")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}

	return
}

//ListNpcBySearch will request any npc matching the pattern of name
func ListNpcBySearch(page *model.Page, npc *model.Npc, user *model.User) (npcs []*model.Npc, err error) {

	err = validateOrderByNpcField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareNpc(npc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre npc")
		return
	}

	err = validateNpc(npc, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate npc")
		return
	}
	reader, err := getReader("npc")
	if err != nil {
		err = errors.Wrap(err, "failed to get npc reader")
		return
	}

	npcs, err = reader.ListNpcBySearch(page, npc)
	if err != nil {
		err = errors.Wrap(err, "failed to list npc by search")
		return
	}

	page.Total, err = reader.ListNpcBySearchTotalCount(npc)
	if err != nil {
		err = errors.Wrap(err, "failed to get page total")
		return
	}

	for _, npc := range npcs {
		err = sanitizeNpc(npc, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize npc")
			return
		}
	}

	err = sanitizeNpc(npc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search npc")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateNpc will create an npc using provided information
func CreateNpc(npc *model.Npc, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list npc by search without guide+")
		return
	}
	err = prepareNpc(npc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare npc")
		return
	}

	err = validateNpc(npc, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate npc")
		return
	}
	npc.ID = 0
	//npc.TimeCreation = time.Now().Unix()
	writer, err := getWriter("npc")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for npc")
		return
	}
	err = writer.CreateNpc(npc)
	if err != nil {
		err = errors.Wrap(err, "failed to create npc")
		return
	}

	memWriter, err := getWriter("npc")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for npc")
		return
	}
	err = memWriter.CreateNpc(npc)
	if err != nil {
		err = errors.Wrap(err, "failed to edit npc")
		return
	}

	err = sanitizeNpc(npc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize npc")
		return
	}
	return
}

//GetNpc gets an npc by provided npcID
func GetNpc(npc *model.Npc, user *model.User) (err error) {
	err = prepareNpc(npc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare npc")
		return
	}

	err = validateNpc(npc, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate npc")
		return
	}

	reader, err := getReader("npc")
	if err != nil {
		err = errors.Wrap(err, "failed to get npc reader")
		return
	}

	err = reader.GetNpc(npc)
	if err != nil {
		err = errors.Wrap(err, "failed to get npc")
		return
	}

	err = sanitizeNpc(npc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize npc")
		return
	}

	return
}

//EditNpc edits an existing npc
func EditNpc(npc *model.Npc, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list npc by search without guide+")
		return
	}
	err = prepareNpc(npc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare npc")
		return
	}

	err = validateNpc(npc,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate npc")
		return
	}
	writer, err := getWriter("npc")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for npc")
		return
	}
	err = writer.EditNpc(npc)
	if err != nil {
		err = errors.Wrap(err, "failed to edit npc")
		return
	}

	memWriter, err := getWriter("npc")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for npc")
		return
	}
	err = memWriter.EditNpc(npc)
	if err != nil {
		err = errors.Wrap(err, "failed to edit npc")
		return
	}

	err = sanitizeNpc(npc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize npc")
		return
	}
	return
}

//DeleteNpc deletes an npc by provided npcID
func DeleteNpc(npc *model.Npc, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete npc without admin+")
		return
	}
	err = prepareNpc(npc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare npc")
		return
	}

	err = validateNpc(npc, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate npc")
		return
	}
	writer, err := getWriter("npc")
	if err != nil {
		err = errors.Wrap(err, "failed to get npc writer")
		return
	}
	err = writer.DeleteNpc(npc)
	if err != nil {
		err = errors.Wrap(err, "failed to delete npc")
		return
	}

	memWriter, err := getWriter("npc")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for npc")
		return
	}
	err = memWriter.DeleteNpc(npc)
	if err != nil {
		err = errors.Wrap(err, "failed to delete npc")
		return
	}
	return
}

func prepareNpc(npc *model.Npc, user *model.User) (err error) {
	if npc == nil {
		err = fmt.Errorf("empty npc")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateNpc(npc *model.Npc, required []string, optional []string) (err error) {
	schema, err := newSchemaNpc(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(npc))
	if err != nil {
		return
	}

	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			if vErr.Message == "invalid" {
				vErr.Message = fmt.Sprintf("%s: %s", res.Field(), res.Description())
			}
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}
	return
}

func validateOrderByNpcField(page *model.Page) (err error) {
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

func sanitizeNpc(npc *model.Npc, user *model.User) (err error) {
	if npc.ClassID > 0 {
		npc.Class = &model.Class{
			ID: npc.ClassID,
		}
		err = GetClass(npc.Class, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to get class for ID %d", npc.ClassID)
			return
		}
	}
	if npc.RaceID > 0 {
		npc.Race = &model.Race{
			ID: npc.RaceID,
		}
		err = GetRace(npc.Race, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to get race for ID %d", npc.RaceID)
			return
		}
	}

	return
}

func newSchemaNpc(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyNpc(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyNpc(field); err != nil {
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

func getSchemaPropertyNpc(field string) (prop model.Schema, err error) {
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
