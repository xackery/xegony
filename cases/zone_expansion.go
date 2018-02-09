package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadZoneExpansionFromFileToMemory is ran during initialization
func LoadZoneExpansionFromFileToMemory() (err error) {

	fr, err := file.New("config", "zoneExpansion.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("zoneExpansion-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize zoneExpansion-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("zoneExpansion-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize zoneExpansion-memory")
		return
	}

	fileReader, err := getReader("zoneExpansion-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneExpansion-file reader")
		return
	}

	memWriter, err := getWriter("zoneExpansion-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneExpansion-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListZoneExpansionTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list zoneExpansion count")
		return
	}
	page.Limit = page.Total

	zoneExpansions, err := fileReader.ListZoneExpansion(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list zoneExpansions")
		return
	}

	for _, zoneExpansion := range zoneExpansions {
		err = memWriter.CreateZoneExpansion(zoneExpansion)
		if err != nil {
			err = errors.Wrap(err, "failed to create zoneExpansion")
			return
		}
	}

	fmt.Printf("%d zone expansions, ", len(zoneExpansions))
	return
}

//ListZoneExpansion lists all zoneExpansions accessible by provided user
func ListZoneExpansion(page *model.Page, user *model.User) (zoneExpansions []*model.ZoneExpansion, err error) {
	err = validateOrderByZoneExpansionField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("zoneExpansion-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for zoneExpansion")
		return
	}

	page.Total, err = reader.ListZoneExpansionTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list zoneExpansion toal count")
		return
	}

	zoneExpansions, err = reader.ListZoneExpansion(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list zoneExpansion")
		return
	}
	for i, zoneExpansion := range zoneExpansions {
		err = sanitizeZoneExpansion(zoneExpansion, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize zoneExpansion element %d", i)
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

//ListZoneExpansionBySearch will request any zoneExpansion matching the pattern of name
func ListZoneExpansionBySearch(page *model.Page, zoneExpansion *model.ZoneExpansion, user *model.User) (zoneExpansions []*model.ZoneExpansion, err error) {

	err = validateOrderByZoneExpansionField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareZoneExpansion(zoneExpansion, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre zoneExpansion")
		return
	}

	err = validateZoneExpansion(zoneExpansion, nil, []string{ //optional
		"shortName",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate zoneExpansion")
		return
	}
	reader, err := getReader("zoneExpansion")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneExpansion reader")
		return
	}

	zoneExpansions, err = reader.ListZoneExpansionBySearch(page, zoneExpansion)
	if err != nil {
		err = errors.Wrap(err, "failed to list zoneExpansion by search")
		return
	}

	for _, zoneExpansion := range zoneExpansions {
		err = sanitizeZoneExpansion(zoneExpansion, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize zoneExpansion")
			return
		}
	}

	err = sanitizeZoneExpansion(zoneExpansion, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search zoneExpansion")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateZoneExpansion will create an zoneExpansion using provided information
func CreateZoneExpansion(zoneExpansion *model.ZoneExpansion, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list zoneExpansion by search without guide+")
		return
	}
	err = prepareZoneExpansion(zoneExpansion, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare zoneExpansion")
		return
	}

	err = validateZoneExpansion(zoneExpansion, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate zoneExpansion")
		return
	}
	zoneExpansion.ID = 0
	//zoneExpansion.TimeCreation = time.Now().Unix()
	writer, err := getWriter("zoneExpansion")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for zoneExpansion")
		return
	}
	err = writer.CreateZoneExpansion(zoneExpansion)
	if err != nil {
		err = errors.Wrap(err, "failed to create zoneExpansion")
		return
	}
	err = sanitizeZoneExpansion(zoneExpansion, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize zoneExpansion")
		return
	}
	return
}

//GetZoneExpansion gets an zoneExpansion by provided zoneExpansionID
func GetZoneExpansion(zoneExpansion *model.ZoneExpansion, user *model.User) (err error) {
	err = prepareZoneExpansion(zoneExpansion, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare zoneExpansion")
		return
	}

	err = validateZoneExpansion(zoneExpansion, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate zoneExpansion")
		return
	}

	reader, err := getReader("zoneExpansion-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneExpansion reader")
		return
	}

	err = reader.GetZoneExpansion(zoneExpansion)
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneExpansion")
		return
	}

	err = sanitizeZoneExpansion(zoneExpansion, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize zoneExpansion")
		return
	}

	return
}

//EditZoneExpansion edits an existing zoneExpansion
func EditZoneExpansion(zoneExpansion *model.ZoneExpansion, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list zoneExpansion by search without guide+")
		return
	}
	err = prepareZoneExpansion(zoneExpansion, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare zoneExpansion")
		return
	}

	err = validateZoneExpansion(zoneExpansion,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"charname",
			"sharedplat",
			"password",
			"status",
			"lszoneExpansionID",
			"gmspeed",
			"revoked",
			"karma",
			"miniloginIp",
			"hideme",
			"rulesflag",
			"suspendeduntil",
			"timeCreation",
			"expansion",
			"banReason",
			"suspendReason"},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate zoneExpansion")
		return
	}
	writer, err := getWriter("zoneExpansion")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for zoneExpansion")
		return
	}
	err = writer.EditZoneExpansion(zoneExpansion)
	if err != nil {
		err = errors.Wrap(err, "failed to edit zoneExpansion")
		return
	}
	err = sanitizeZoneExpansion(zoneExpansion, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize zoneExpansion")
		return
	}
	return
}

//DeleteZoneExpansion deletes an zoneExpansion by provided zoneExpansionID
func DeleteZoneExpansion(zoneExpansion *model.ZoneExpansion, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete zoneExpansion without admin+")
		return
	}
	err = prepareZoneExpansion(zoneExpansion, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare zoneExpansion")
		return
	}

	err = validateZoneExpansion(zoneExpansion, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate zoneExpansion")
		return
	}
	writer, err := getWriter("zoneExpansion")
	if err != nil {
		err = errors.Wrap(err, "failed to get zoneExpansion writer")
		return
	}
	err = writer.DeleteZoneExpansion(zoneExpansion)
	if err != nil {
		err = errors.Wrap(err, "failed to delete zoneExpansion")
		return
	}
	return
}

func prepareZoneExpansion(zoneExpansion *model.ZoneExpansion, user *model.User) (err error) {
	if zoneExpansion == nil {
		err = fmt.Errorf("empty zoneExpansion")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateZoneExpansion(zoneExpansion *model.ZoneExpansion, required []string, optional []string) (err error) {
	schema, err := newSchemaZoneExpansion(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(zoneExpansion))
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

func validateOrderByZoneExpansionField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "id"
	}

	validNames := []string{
		"id",
		"short_name",
		"zoneExpansionidnumber",
		"long_name",
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

func sanitizeZoneExpansion(zoneExpansion *model.ZoneExpansion, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaZoneExpansion(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyZoneExpansion(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyZoneExpansion(field); err != nil {
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

func getSchemaPropertyZoneExpansion(field string) (prop model.Schema, err error) {
	switch field {

	case "shortName":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 64
	case "ID":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fileName":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 64
	case "longName":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 64
	case "mapFileName":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 64
	case "safeX":
		prop.Type = "float"
		prop.Minimum = 0
	case "safeY":
		prop.Type = "float"
		prop.Minimum = 0
	case "safeZ":
		prop.Type = "float"
		prop.Minimum = 0
	case "graveyardID":
		prop.Type = "float"
		prop.Minimum = 0
	case "minLevel":
		prop.Type = "integer"
		prop.Minimum = 0
	case "minStatus":
		prop.Type = "integer"
		prop.Minimum = 0
	case "zoneExpansionIDNumber":
		prop.Type = "integer"
		prop.Minimum = 0
	case "version":
		prop.Type = "integer"
		prop.Minimum = 0
	case "timezoneExpansion":
		prop.Type = "integer"
		prop.Minimum = 0
	case "maxClients":
		prop.Type = "integer"
		prop.Minimum = 0
	case "ruleset":
		prop.Type = "integer"
		prop.Minimum = 0
	case "note":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 64
	case "underworld":
		prop.Type = "float"
		prop.Minimum = 0
	case "MinClip":
		prop.Type = "float"
		prop.Minimum = 0
	case "MaxClip":
		prop.Type = "float"
		prop.Minimum = 0
	case "fogMinClip":
		prop.Type = "float"
		prop.Minimum = 0
	case "fogMaxClip":
		prop.Type = "float"
		prop.Minimum = 0
	case "fogBlue":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogRed":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogGreen":
		prop.Type = "integer"
		prop.Minimum = 0
	case "sky":
		prop.Type = "integer"
		prop.Minimum = 0
	case "zType":
		prop.Type = "integer"
		prop.Minimum = 0
	case "zoneExpansionExpMultiplier":
		prop.Type = "float"
		prop.Minimum = 0
	case "walkSpeed":
		prop.Type = "float"
		prop.Minimum = 0
	case "timeType":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogRed1":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogGreen1":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogBlue1":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogMinClip1":
		prop.Type = "float"
		prop.Minimum = 0
	case "fogMaxClip1":
		prop.Type = "float"
		prop.Minimum = 0
	case "fogRed2":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogGreen2":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogBlue2":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogMinClip2":
		prop.Type = "float"
		prop.Minimum = 0
	case "fogMaxClip2":
		prop.Type = "float"
		prop.Minimum = 0
	case "fogRed3":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogGreen3":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogBlue3":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogMinClip3":
		prop.Type = "float"
		prop.Minimum = 0
	case "fogMaxClip3":
		prop.Type = "float"
		prop.Minimum = 0
	case "fogRed4":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogGreen4":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogBlue4":
		prop.Type = "integer"
		prop.Minimum = 0
	case "fogMinClip4":
		prop.Type = "float"
		prop.Minimum = 0
	case "fogMaxClip4":
		prop.Type = "float"
		prop.Minimum = 0
	case "fogDensity":
		prop.Type = "float"
		prop.Minimum = 0
	case "flagNeeded":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 64
	case "canBind":
		prop.Type = "integer"
		prop.Minimum = 0
	case "canCombat":
		prop.Type = "integer"
		prop.Minimum = 0
	case "canLevitate":
		prop.Type = "integer"
		prop.Minimum = 0
	case "castOutdoor":
		prop.Type = "integer"
		prop.Minimum = 0
	case "hotZoneExpansion":
		prop.Type = "integer"
		prop.Minimum = 0
	case "instType":
		prop.Type = "integer"
		prop.Minimum = 0
	case "shutdownDelay":
		prop.Type = "integer"
		prop.Minimum = 0
	case "peqZoneExpansion":
		prop.Type = "integer"
		prop.Minimum = 0
	case "expansion":
		prop.Type = "integer"
		prop.Minimum = 0
	case "suspendBuffs":
		prop.Type = "integer"
		prop.Minimum = 0
	case "rainChance1":
		prop.Type = "integer"
		prop.Minimum = 0
	case "rainChance2":
		prop.Type = "integer"
		prop.Minimum = 0
	case "rainChance3":
		prop.Type = "integer"
		prop.Minimum = 0
	case "rainChance4":
		prop.Type = "integer"
		prop.Minimum = 0
	case "rainDuration1":
		prop.Type = "integer"
		prop.Minimum = 0
	case "rainDuration2":
		prop.Type = "integer"
		prop.Minimum = 0
	case "rainDuration3":
		prop.Type = "integer"
		prop.Minimum = 0
	case "rainDuration4":
		prop.Type = "integer"
		prop.Minimum = 0
	case "snowChance1":
		prop.Type = "integer"
		prop.Minimum = 0
	case "snowChance2":
		prop.Type = "integer"
		prop.Minimum = 0
	case "snowChance3":
		prop.Type = "integer"
		prop.Minimum = 0
	case "snowChance4":
		prop.Type = "integer"
		prop.Minimum = 0
	case "snowDuration1":
		prop.Type = "integer"
		prop.Minimum = 0
	case "snowDuration2":
		prop.Type = "integer"
		prop.Minimum = 0
	case "snowDuration3":
		prop.Type = "integer"
		prop.Minimum = 0
	case "snowDuration4":
		prop.Type = "integer"
		prop.Minimum = 0
	case "gravity":
		prop.Type = "float"
		prop.Minimum = 0
	case "type":
		prop.Type = "integer"
		prop.Minimum = 0
	case "skylock":
		prop.Type = "integer"
		prop.Minimum = 0
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
