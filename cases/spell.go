package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadSpellFromDBToMemory is ran during initialization
func LoadSpellFromDBToMemory() (err error) {
	fmt.Printf("Loading spells...")
	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("spell-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize spell-memory")
		return
	}

	dbReader, err := getReader("spell")
	if err != nil {
		err = errors.Wrap(err, "failed to get spell reader")
		return
	}

	memWriter, err := getWriter("spell-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spell-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = dbReader.ListSpellTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list spell count")
		return
	}
	var totalSpells []*model.Spell
	var spells []*model.Spell
	for {
		spells, err = dbReader.ListSpell(page)
		if err != nil {
			err = errors.Wrap(err, "failed to list spells")
			return
		}
		totalSpells = append(totalSpells, spells...)
		if int64(len(totalSpells)) >= page.Total {
			break
		}
		page.Offset++
	}

	for _, spell := range totalSpells {
		err = memWriter.CreateSpell(spell)
		if err != nil {
			err = errors.Wrap(err, "failed to create spell")
			return
		}
	}

	fmt.Printf(" (%d)\n", len(totalSpells))
	return
}

//ListSpell lists all spells accessible by provided user
func ListSpell(page *model.Page, user *model.User) (spells []*model.Spell, err error) {
	err = validateOrderBySpellField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("spell-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for spell")
		return
	}

	page.Total, err = reader.ListSpellTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list spell toal count")
		return
	}

	spells, err = reader.ListSpell(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list spell")
		return
	}
	for i, spell := range spells {
		err = sanitizeSpell(spell, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize spell element %d", i)
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

//ListSpellBySearch will request any spell matching the pattern of name
func ListSpellBySearch(page *model.Page, spell *model.Spell, user *model.User) (spells []*model.Spell, err error) {
	/*err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spell by search without guide+")
		return
	}
	*/
	err = validateOrderBySpellField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre spell")
		return
	}

	err = validateSpell(spell, nil, []string{ //optional
		"shortName",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate spell")
		return
	}
	reader, err := getReader("spell-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spell reader")
		return
	}

	spells, err = reader.ListSpellBySearch(page, spell)
	if err != nil {
		err = errors.Wrap(err, "failed to list spell by search")
		return
	}

	for _, spell := range spells {
		err = sanitizeSpell(spell, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize spell")
			return
		}
	}

	err = sanitizeSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search spell")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateSpell will create an spell using provided information
func CreateSpell(spell *model.Spell, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spell by search without guide+")
		return
	}
	err = prepareSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spell")
		return
	}

	err = validateSpell(spell, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spell")
		return
	}
	spell.ID = 0
	//spell.TimeCreation = time.Now().Unix()
	writer, err := getWriter("spell")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spell")
		return
	}
	err = writer.CreateSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to create spell")
		return
	}

	memWriter, err := getWriter("spell-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spell-memory")
		return
	}
	err = memWriter.CreateSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spell-memory")
		return
	}

	err = sanitizeSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spell")
		return
	}
	return
}

//GetSpell gets an spell by provided spellID
func GetSpell(spell *model.Spell, user *model.User) (err error) {
	err = prepareSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spell")
		return
	}

	err = validateSpell(spell, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spell")
		return
	}

	reader, err := getReader("spell-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get spell reader")
		return
	}

	err = reader.GetSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to get spell")
		return
	}

	err = sanitizeSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spell")
		return
	}

	return
}

//EditSpell edits an existing spell
func EditSpell(spell *model.Spell, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list spell by search without guide+")
		return
	}
	err = prepareSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spell")
		return
	}

	err = validateSpell(spell,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"charname",
			"sharedplat",
			"password",
			"status",
			"lsspellID",
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
		err = errors.Wrap(err, "failed to validate spell")
		return
	}
	writer, err := getWriter("spell")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spell")
		return
	}
	err = writer.EditSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spell")
		return
	}

	memWriter, err := getWriter("spell-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spell-memory")
		return
	}
	err = memWriter.EditSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to edit spell-memory")
		return
	}

	err = sanitizeSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize spell")
		return
	}
	return
}

//DeleteSpell deletes an spell by provided spellID
func DeleteSpell(spell *model.Spell, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete spell without admin+")
		return
	}
	err = prepareSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare spell")
		return
	}

	err = validateSpell(spell, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate spell")
		return
	}
	writer, err := getWriter("spell")
	if err != nil {
		err = errors.Wrap(err, "failed to get spell writer")
		return
	}
	err = writer.DeleteSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spell")
		return
	}

	memWriter, err := getWriter("spell-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for spell-memory")
		return
	}
	err = memWriter.DeleteSpell(spell)
	if err != nil {
		err = errors.Wrap(err, "failed to delete spell-memory")
		return
	}
	return
}

func prepareSpell(spell *model.Spell, user *model.User) (err error) {
	if spell == nil {
		err = fmt.Errorf("empty spell")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateSpell(spell *model.Spell, required []string, optional []string) (err error) {
	schema, err := newSchemaSpell(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(spell))
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

func validateOrderBySpellField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "shortName"
	}

	validNames := []string{
		"id",
		"short_name",
		"spellidnumber",
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

func sanitizeSpell(spell *model.Spell, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaSpell(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertySpell(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertySpell(field); err != nil {
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

func getSchemaPropertySpell(field string) (prop model.Schema, err error) {
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
	case "spellIDNumber":
		prop.Type = "integer"
		prop.Minimum = 0
	case "version":
		prop.Type = "integer"
		prop.Minimum = 0
	case "timespell":
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
	case "spellExpMultiplier":
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
	case "hotSpell":
		prop.Type = "integer"
		prop.Minimum = 0
	case "instType":
		prop.Type = "integer"
		prop.Minimum = 0
	case "shutdownDelay":
		prop.Type = "integer"
		prop.Minimum = 0
	case "peqSpell":
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
