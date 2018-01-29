package cases

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListCharacter lists all characters accessible by provided user
func ListCharacter(page *model.Page, user *model.User) (characters []*model.Character, err error) {
	err = validateOrderByCharacterField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("character")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for character")
		return
	}

	page.Total, err = reader.ListCharacterTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list character toal count")
		return
	}

	characters, err = reader.ListCharacter(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list character")
		return
	}
	for i, character := range characters {
		err = sanitizeCharacter(character, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize character element %d", i)
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

//ListCharacterBySearch will request any character matching the pattern of name
func ListCharacterBySearch(page *model.Page, character *model.Character, user *model.User) (characters []*model.Character, err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list character by search without guide+")
		return
	}

	err = validateOrderByCharacterField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareCharacter(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre character")
		return
	}

	err = validateCharacter(character, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate character")
		return
	}
	reader, err := getReader("character")
	if err != nil {
		err = errors.Wrap(err, "failed to get character reader")
		return
	}

	characters, err = reader.ListCharacterBySearch(page, character)
	if err != nil {
		err = errors.Wrap(err, "failed to list character by search")
		return
	}

	for _, character := range characters {
		err = sanitizeCharacter(character, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize character")
			return
		}
	}

	err = sanitizeCharacter(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search character")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateCharacter will create an character using provided information
func CreateCharacter(character *model.Character, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list character by search without guide+")
		return
	}
	err = prepareCharacter(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare character")
		return
	}

	err = validateCharacter(character, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate character")
		return
	}
	character.ID = 0
	character.Birthday = time.Now().Unix()
	writer, err := getWriter("character")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for character")
		return
	}
	err = writer.CreateCharacter(character)
	if err != nil {
		err = errors.Wrap(err, "failed to create character")
		return
	}
	err = sanitizeCharacter(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize character")
		return
	}
	return
}

//GetCharacter gets an character by provided characterID
func GetCharacter(character *model.Character, user *model.User) (err error) {
	err = prepareCharacter(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare character")
		return
	}

	err = validateCharacter(character, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate character")
		return
	}

	reader, err := getReader("character")
	if err != nil {
		err = errors.Wrap(err, "failed to get character reader")
		return
	}

	err = reader.GetCharacter(character)
	if err != nil {
		err = errors.Wrap(err, "failed to get character")
		return
	}

	err = sanitizeCharacter(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize character")
		return
	}

	return
}

//EditCharacter edits an existing character
func EditCharacter(character *model.Character, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list character by search without guide+")
		return
	}
	err = prepareCharacter(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare character")
		return
	}

	err = validateCharacter(character,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"charname",
			"sharedplat",
			"password",
			"status",
			"lscharacterID",
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
		err = errors.Wrap(err, "failed to validate character")
		return
	}
	writer, err := getWriter("character")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for character")
		return
	}
	err = writer.EditCharacter(character)
	if err != nil {
		err = errors.Wrap(err, "failed to edit character")
		return
	}
	err = sanitizeCharacter(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize character")
		return
	}
	return
}

//DeleteCharacter deletes an character by provided characterID
func DeleteCharacter(character *model.Character, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete character without admin+")
		return
	}
	err = prepareCharacter(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare character")
		return
	}

	err = validateCharacter(character, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate character")
		return
	}
	writer, err := getWriter("character")
	if err != nil {
		err = errors.Wrap(err, "failed to get character writer")
		return
	}
	err = writer.DeleteCharacter(character)
	if err != nil {
		err = errors.Wrap(err, "failed to delete character")
		return
	}
	return
}

func prepareCharacter(character *model.Character, user *model.User) (err error) {
	if character == nil {
		err = fmt.Errorf("empty character")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateCharacter(character *model.Character, required []string, optional []string) (err error) {
	schema, err := newSchemaCharacter(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(character))
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

func validateOrderByCharacterField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "id"
	}

	validNames := []string{
		"id",
		"account_id",
		"name",
		"last_name",
		"title",
		"suffix",
		"zone_id",
		"zone_instance",
		"y",
		"x",
		"z",
		"heading",
		"gender",
		"race",
		"class",
		"level",
		"deity",
		"birthday",
		"last_login",
		"time_played",
		"level2",
		"anon",
		"gm",
		"face",
		"hair_color",
		"hair_style",
		"beard",
		"beard_color",
		"eye_color_1",
		"eye_color_2",
		"drakkin_heritage",
		"drakkin_tattoo",
		"drakkin_details",
		"ability_time_seconds",
		"ability_number",
		"ability_time_minutes",
		"ability_time_hours",
		"exp",
		"aa_points_spent",
		"aa_exp",
		"aa_points",
		"group_leadership_exp",
		"raid_leadership_exp",
		"group_leadership_points",
		"raid_leadership_points",
		"points",
		"cur_hp",
		"mana",
		"endurance",
		"intoxication",
		"str",
		"sta",
		"cha",
		"dex",
		"int",
		"agi",
		"wis",
		"zone_change_count",
		"toxicity",
		"hunger_level",
		"thirst_level",
		"ability_up",
		"ldon_points_guk",
		"ldon_points_mir",
		"ldon_points_mmc",
		"ldon_points_ruj",
		"ldon_points_tak",
		"ldon_points_available",
		"tribute_time_remaining",
		"career_tribute_points",
		"tribute_points",
		"tribute_active",
		"pvp_status",
		"pvp_kills",
		"pvp_deaths",
		"pvp_current_points",
		"pvp_career_points",
		"pvp_best_kill_streak",
		"pvp_worst_death_streak",
		"pvp_current_kill_streak",
		"pvp2",
		"pvp_type",
		"show_helm",
		"group_auto_consent",
		"raid_auto_consent",
		"guild_auto_consent",
		"leadership_exp_on",
		"RestTimer",
		"air_remaining",
		"autosplit_enabled",
		"lfp",
		"lfg",
		"mailkey",
		"xtargets",
		"firstlogon",
		"e_aa_effects",
		"e_percent_to_aa",
		"e_expended_aa_spent",
		"aa_points_spent_old",
		"aa_points_old",
		"e_last_invsnapshot",
	}

	for _, name := range validNames {
		if page.OrderBy == name {
			return
		}
	}

	err = &model.ErrValidation{
		Message: "orderBy is invalid",
		Reasons: map[string]string{
			"orderBy": "field is not valid",
		},
	}
	return
}

func sanitizeCharacter(character *model.Character, user *model.User) (err error) {

	return
}

func newSchemaCharacter(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyCharacter(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyCharacter(field); err != nil {
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

func getSchemaPropertyCharacter(field string) (prop model.Schema, err error) {
	switch field {
	case "ID": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "accountID": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "name": //string
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "lastName": //string
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "title": //string
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "suffix": //string
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "zoneID": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "zoneInstance": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "y": //float64
		prop.Type = "float"
	case "x": //float64
		prop.Type = "float"
	case "z": //float64
		prop.Type = "float"
	case "heading": //float64
		prop.Type = "float"
	case "gender": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "raceID": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "classID": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "level": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "deity": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "birthday": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "lastLogin": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "timePlayed": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "level2": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "anon": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "gm": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "face": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "hairColor": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "hairStyle": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "beard": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "beardColor": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "eyeColor1": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "eyeColor2": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "drakkinHeritage": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "drakkinTattoo": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "drakkinDetails": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "abilityTimeSeconds": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "abilityNumber": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "abilityTimeMinutes": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "abilityTimeHours": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "exp": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "aaPointsSpent": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "aaExp": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "aaPoints": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "groupLeadershipExp": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "raidLeadershipExp": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "groupLeadershipPoints": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "raidLeadershipPoints": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "points": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "curHp": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "mana": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "endurance": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "intoxication": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "str": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "sta": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "cha": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "dex": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "int": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "agi": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "wis": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "zoneChangeCount": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "toxicity": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "hungerLevel": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "thirstLevel": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "abilityUp": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "ldonPointsGuk": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "ldonPointsMir": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "ldonPointsMmc": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "ldonPointsRuj": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "ldonPointsTak": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "ldonPointsAvailable": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "tributeTimeRemaining": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "careerTributePoints": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "tributePoints": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "tributeActive": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "pvpStatus": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "pvpKills": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "pvpDeaths": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "pvpCurrentPoints": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "pvpCareerPoints": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "pvpBestKillStreak": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "pvpWorstDeathStreak": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "pvpCurrentKillStreak": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "pvp2": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "pvpType": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "showHelm": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "groupAutoConsent": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "raidAutoConsent": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "guildAutoConsent": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "leadershipExpOn": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "RestTimer": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "airRemaining": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "autosplitEnabled": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "lfp": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "lfg": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "mailkey": //string
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "xtargets": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "firstlogon": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "eAaEffects": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "ePercentToAa": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "eExpendedAaSpent": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "aaPointsSpentOld": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "aaPointsOld": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	case "eLastInvsnapshot": //int64
		prop.Type = "integer"
		prop.Minimum = 1
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
