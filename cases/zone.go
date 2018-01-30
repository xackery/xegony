package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListZone lists all zones accessible by provided user
func ListZone(page *model.Page, user *model.User) (zones []*model.Zone, err error) {
	err = validateOrderByZoneField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("zone")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for zone")
		return
	}

	page.Total, err = reader.ListZoneTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list zone toal count")
		return
	}

	zones, err = reader.ListZone(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list zone")
		return
	}
	for i, zone := range zones {
		err = sanitizeZone(zone, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize zone element %d", i)
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

//ListZoneBySearch will request any zone matching the pattern of name
func ListZoneBySearch(page *model.Page, zone *model.Zone, user *model.User) (zones []*model.Zone, err error) {
	/*err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list zone by search without guide+")
		return
	}
	*/
	err = validateOrderByZoneField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareZone(zone, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre zone")
		return
	}

	err = validateZone(zone, nil, []string{ //optional
		"shortName",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate zone")
		return
	}
	reader, err := getReader("zone")
	if err != nil {
		err = errors.Wrap(err, "failed to get zone reader")
		return
	}

	zones, err = reader.ListZoneBySearch(page, zone)
	if err != nil {
		err = errors.Wrap(err, "failed to list zone by search")
		return
	}

	for _, zone := range zones {
		err = sanitizeZone(zone, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize zone")
			return
		}
	}

	err = sanitizeZone(zone, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search zone")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateZone will create an zone using provided information
func CreateZone(zone *model.Zone, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list zone by search without guide+")
		return
	}
	err = prepareZone(zone, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare zone")
		return
	}

	err = validateZone(zone, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate zone")
		return
	}
	zone.ID = 0
	//zone.TimeCreation = time.Now().Unix()
	writer, err := getWriter("zone")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for zone")
		return
	}
	err = writer.CreateZone(zone)
	if err != nil {
		err = errors.Wrap(err, "failed to create zone")
		return
	}
	err = sanitizeZone(zone, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize zone")
		return
	}
	return
}

//GetZone gets an zone by provided zoneID
func GetZone(zone *model.Zone, user *model.User) (err error) {
	err = prepareZone(zone, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare zone")
		return
	}

	err = validateZone(zone, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate zone")
		return
	}

	reader, err := getReader("zone")
	if err != nil {
		err = errors.Wrap(err, "failed to get zone reader")
		return
	}

	err = reader.GetZone(zone)
	if err != nil {
		err = errors.Wrap(err, "failed to get zone")
		return
	}

	err = sanitizeZone(zone, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize zone")
		return
	}

	return
}

//EditZone edits an existing zone
func EditZone(zone *model.Zone, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list zone by search without guide+")
		return
	}
	err = prepareZone(zone, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare zone")
		return
	}

	err = validateZone(zone,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"charname",
			"sharedplat",
			"password",
			"status",
			"lszoneID",
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
		err = errors.Wrap(err, "failed to validate zone")
		return
	}
	writer, err := getWriter("zone")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for zone")
		return
	}
	err = writer.EditZone(zone)
	if err != nil {
		err = errors.Wrap(err, "failed to edit zone")
		return
	}
	err = sanitizeZone(zone, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize zone")
		return
	}
	return
}

//DeleteZone deletes an zone by provided zoneID
func DeleteZone(zone *model.Zone, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete zone without admin+")
		return
	}
	err = prepareZone(zone, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare zone")
		return
	}

	err = validateZone(zone, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate zone")
		return
	}
	writer, err := getWriter("zone")
	if err != nil {
		err = errors.Wrap(err, "failed to get zone writer")
		return
	}
	err = writer.DeleteZone(zone)
	if err != nil {
		err = errors.Wrap(err, "failed to delete zone")
		return
	}
	return
}

func prepareZone(zone *model.Zone, user *model.User) (err error) {
	if zone == nil {
		err = fmt.Errorf("empty zone")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateZone(zone *model.Zone, required []string, optional []string) (err error) {
	schema, err := newSchemaZone(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(zone))
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

func validateOrderByZoneField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "shortName"
	}

	validNames := []string{
		"id",
		"short_name",
		"zoneidnumber",
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

func sanitizeZone(zone *model.Zone, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaZone(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyZone(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyZone(field); err != nil {
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

func getSchemaPropertyZone(field string) (prop model.Schema, err error) {
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
	case "zoneIDNumber":
		prop.Type = "integer"
		prop.Minimum = 0
	case "version":
		prop.Type = "integer"
		prop.Minimum = 0
	case "timezone":
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
	case "zoneExpMultiplier":
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
	case "hotZone":
		prop.Type = "integer"
		prop.Minimum = 0
	case "instType":
		prop.Type = "integer"
		prop.Minimum = 0
	case "shutdownDelay":
		prop.Type = "integer"
		prop.Minimum = 0
	case "peqZone":
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
