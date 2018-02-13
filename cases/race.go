package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadRaceFromFileToMemory is ran during initialization
func LoadRaceFromFileToMemory() (err error) {

	fr, err := file.New("config", "race.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("race-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize race-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("race-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize race-memory")
		return
	}

	fileReader, err := getReader("race-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get race-file reader")
		return
	}

	memWriter, err := getWriter("race-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get race-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListRaceTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list race count")
		return
	}
	page.Limit = page.Total

	races, err := fileReader.ListRace(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list races")
		return
	}

	for _, race := range races {
		err = memWriter.CreateRace(race)
		if err != nil {
			err = errors.Wrap(err, "failed to create race")
			return
		}
	}

	fmt.Printf("%d races, ", len(races))
	return
}

//ListRace lists all races accessible by provided user
func ListRace(page *model.Page, user *model.User) (races []*model.Race, err error) {
	err = validateOrderByRaceField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("race-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for race")
		return
	}

	page.Total, err = reader.ListRaceTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list race toal count")
		return
	}

	races, err = reader.ListRace(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list race")
		return
	}
	for i, race := range races {
		err = sanitizeRace(race, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize race element %d", i)
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

//ListRaceByBit will request any race matching the pattern of name
func ListRaceByBit(page *model.Page, race *model.Race, user *model.User) (races []*model.Race, err error) {

	err = validateOrderByRaceField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareRace(race, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre race")
		return
	}

	reader, err := getReader("race-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get race-memory reader")
		return
	}

	races, err = reader.ListRaceByBit(page, race)
	if err != nil {
		err = errors.Wrap(err, "failed to list race by bit")
		return
	}

	for _, race := range races {
		err = sanitizeRace(race, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize race")
			return
		}
	}

	err = sanitizeRace(race, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search race")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//ListRaceBySearch will request any race matching the pattern of name
func ListRaceBySearch(page *model.Page, race *model.Race, user *model.User) (races []*model.Race, err error) {

	err = validateOrderByRaceField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareRace(race, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre race")
		return
	}

	err = validateRace(race, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate race")
		return
	}
	reader, err := getReader("race-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get race-memory reader")
		return
	}

	races, err = reader.ListRaceBySearch(page, race)
	if err != nil {
		err = errors.Wrap(err, "failed to list race by search")
		return
	}

	for _, race := range races {
		err = sanitizeRace(race, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize race")
			return
		}
	}

	err = sanitizeRace(race, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search race")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateRace will create an race using provided information
func CreateRace(race *model.Race, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list race by search without guide+")
		return
	}
	err = prepareRace(race, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare race")
		return
	}

	err = validateRace(race, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate race")
		return
	}
	race.ID = 0
	writer, err := getWriter("race-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for race")
		return
	}
	err = writer.CreateRace(race)
	if err != nil {
		err = errors.Wrap(err, "failed to create race")
		return
	}

	fileWriter, err := getWriter("race-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get race-file writer")
		return
	}
	err = fileWriter.CreateRace(race)
	if err != nil {
		err = errors.Wrap(err, "failed to create race-file")
		return
	}
	err = sanitizeRace(race, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize race")
		return
	}
	return
}

//GetRace gets an race by provided raceID
func GetRace(race *model.Race, user *model.User) (err error) {
	err = prepareRace(race, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare race")
		return
	}

	err = validateRace(race, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate race")
		return
	}

	reader, err := getReader("race-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get race-memory reader")
		return
	}

	err = reader.GetRace(race)
	if err != nil {
		err = errors.Wrap(err, "failed to get race")
		return
	}

	err = sanitizeRace(race, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize race")
		return
	}

	return
}

//EditRace edits an existing race
func EditRace(race *model.Race, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list race by search without guide+")
		return
	}
	err = prepareRace(race, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare race")
		return
	}

	err = validateRace(race,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"male",
			"female",
			"neutral",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate race")
		return
	}
	writer, err := getWriter("race-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for race")
		return
	}
	err = writer.EditRace(race)
	if err != nil {
		err = errors.Wrap(err, "failed to edit race")
		return
	}

	fileWriter, err := getWriter("race-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get race-file writer")
		return
	}
	err = fileWriter.EditRace(race)
	if err != nil {
		err = errors.Wrap(err, "failed to edit race-file")
		return
	}

	err = sanitizeRace(race, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize race")
		return
	}
	return
}

//DeleteRace deletes an race by provided raceID
func DeleteRace(race *model.Race, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete race without admin+")
		return
	}
	err = prepareRace(race, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare race")
		return
	}

	err = validateRace(race, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate race")
		return
	}
	writer, err := getWriter("race-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get race-memory writer")
		return
	}
	err = writer.DeleteRace(race)
	if err != nil {
		err = errors.Wrap(err, "failed to delete race")
		return
	}

	fileWriter, err := getWriter("race-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get race-file writer")
		return
	}
	err = fileWriter.DeleteRace(race)
	if err != nil {
		err = errors.Wrap(err, "failed to delete race-file")
		return
	}

	return
}

func prepareRace(race *model.Race, user *model.User) (err error) {
	if race == nil {
		err = fmt.Errorf("empty race")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateRace(race *model.Race, required []string, optional []string) (err error) {
	schema, err := newSchemaRace(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(race))
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

func validateOrderByRaceField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "id"
	}

	validNames := []string{
		"id",
		"name",
		"bit",
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

func sanitizeRace(race *model.Race, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	if race.Icon == "" {
		race.Icon = "xa-octopus"
	}
	if len(race.Icon) > 2 && race.Icon[2:3] == "-" {
		race.Icon = fmt.Sprintf("%s %s", race.Icon[0:2], race.Icon)
	}
	return
}

func newSchemaRace(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyRace(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyRace(field); err != nil {
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

func getSchemaPropertyRace(field string) (prop model.Schema, err error) {
	switch field {
	case "ID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "male":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "female":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "neutral":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "icon":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}

func getRaceListByBit(bit int64) string {
	races := ""
	if bit == 65535 {
		return "ALL"
	}
	if bit&1 == 1 {
		races += "HUM "
	}
	if bit&2 == 2 {
		races += "BAR "
	}
	if bit&4 == 4 {
		races += "ERU "
	}
	if bit&8 == 8 {
		races += "WEF "
	}
	if bit&16 == 16 {
		races += "HEF "
	}
	if bit&32 == 32 {
		races += "DEF "
	}
	if bit&64 == 64 {
		races += "HLF "
	}
	if bit&128 == 128 {
		races += "DWF "
	}
	if bit&256 == 256 {
		races += "TRL "
	}
	if bit&512 == 512 {
		races += "OGR "
	}
	if bit&1024 == 1024 {
		races += "HLF "
	}
	if bit&2048 == 2048 {
		races += "GNM "
	}
	if bit&4096 == 4096 {
		races += "IKS "
	}
	if bit&8192 == 8192 {
		races += "VAH "
	}
	if bit&16384 == 16384 {
		races += "FRO "
	}
	if bit&32768 == 32768 {
		races += "DRA "
	}
	if len(races) > 0 {
		races = races[0 : len(races)-1]
	}
	if len(races) == 0 {
		races = "NONE"
	}
	return races
}
