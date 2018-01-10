package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

var ()

//SkillRepository handles SkillRepository cases and is a gateway to storage
type SkillRepository struct {
	stor               storage.Storage
	skillCache         map[int64]*model.Skill
	isSkillCacheLoaded bool
}

//Initialize handler
func (c *SkillRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}

	c.stor = stor
	c.isSkillCacheLoaded = false
	if err = c.rebuildCache(); err != nil {
		return
	}
	return
}

func (c *SkillRepository) rebuildCache() (err error) {
	if c.isSkillCacheLoaded {
		return
	}
	c.isSkillCacheLoaded = true
	c.skillCache = make(map[int64]*model.Skill)
	skills, err := c.list()
	if err != nil {
		return
	}

	for _, skill := range skills {
		c.skillCache[skill.ID] = skill
	}
	return
}

//Get handler
func (c *SkillRepository) Get(skillID int64) (skill *model.Skill, err error) {
	if skillID == 0 {
		err = fmt.Errorf("Invalid Skill ID")
		return
	}
	skill = c.skillCache[skillID]
	//skill, err = c.stor.GetSkill(skillID)
	return
}

//GetByName gets a skill by it's name
func (c *SkillRepository) GetByName(name string) (skill *model.Skill, err error) {
	for _, skillC := range c.skillCache {
		if skillC.Name == name {
			skill = skillC
			return
		}
	}
	return
}

//ListByType gets a skill by it's name
func (c *SkillRepository) ListByType(skillType int64) (skills []*model.Skill, err error) {
	for _, skillC := range c.skillCache {
		if skillC.Type == skillType {
			skills = append(skills, skillC)
		}
	}
	return
}

//Create handler
func (c *SkillRepository) Create(skill *model.Skill) (err error) {
	if skill == nil {
		err = fmt.Errorf("Empty skill")
		return
	}
	schema, err := c.newSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}

	skill.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(skill))
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
	//err = c.stor.CreateSkill(skill)
	//if err != nil {
	//	return
	//}
	c.isSkillCacheLoaded = false
	c.rebuildCache()
	return
}

//Edit handler
func (c *SkillRepository) Edit(skillID int64, skill *model.Skill) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(skill))
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

	//if err = c.stor.EditSkill(skillID, skill); err != nil {
	//	return
	//}
	if err = c.rebuildCache(); err != nil {
		return
	}
	return
}

//Delete handler
func (c *SkillRepository) Delete(skillID int64) (err error) {
	//err = c.stor.DeleteSkill(skillID)
	//if err != nil {
	//	return
	//}
	//if err = c.rebuildCache(); err != nil {
	//	return
	//}
	return
}

func (c *SkillRepository) list() (skills []*model.Skill, err error) {
	for _, skill := range skillsList {
		skills = append(skills, skill)
	}
	return
}

//List handler
func (c *SkillRepository) List() (skills []*model.Skill, err error) {
	for _, skill := range c.skillCache {
		skills = append(skills, skill)
	}
	return
}

func (c *SkillRepository) prepare(skill *model.Skill) (err error) {

	return
}

//newSchema handler
func (c *SkillRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
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

//getSchemaProperty handler
func (c *SkillRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	case "type":
		prop.Type = "integer"
		prop.Minimum = 0
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}
	return
}

var skillsList = map[int]*model.Skill{
	0: &model.Skill{
		ID:   0,
		Name: "1H Blunt",
		Type: 0,
	},
	1: &model.Skill{
		ID:   1,
		Name: "1H Slashing",
		Type: 0,
	},
	2: &model.Skill{
		ID:   2,
		Name: "2H Blunt",
		Type: 0,
	},
	3: &model.Skill{
		ID:   3,
		Name: "2H Slashing",
		Type: 0,
	},
	4: &model.Skill{
		ID:   4,
		Name: "Abjuration",
		Type: 0,
	},
	5: &model.Skill{
		ID:   5,
		Name: "Alteration",
		Type: 0,
	},
	6: &model.Skill{
		ID:   6,
		Name: "Apply Poison",
		Type: 0,
	},
	7: &model.Skill{
		ID:   7,
		Name: "Archery",
		Type: 0,
	},
	8: &model.Skill{
		ID:   8,
		Name: "Backstab",
		Type: 0,
	},
	9: &model.Skill{
		ID:   9,
		Name: "Bind Wound",
		Type: 0,
	},
	10: &model.Skill{
		ID:   10,
		Name: "Bash",
		Type: 0,
	},
	11: &model.Skill{
		ID:   11,
		Name: "Block",
		Type: 0,
	},
	12: &model.Skill{
		ID:   12,
		Name: "Brass Instruments",
		Type: 0,
	},
	13: &model.Skill{
		ID:   13,
		Name: "Channeling",
		Type: 0,
	},
	14: &model.Skill{
		ID:   14,
		Name: "Conjuration",
		Type: 0,
	},
	15: &model.Skill{
		ID:   15,
		Name: "Defense",
		Type: 0,
	},
	16: &model.Skill{
		ID:   16,
		Name: "Disarm",
		Type: 0,
	},
	17: &model.Skill{
		ID:   17,
		Name: "Disarm Traps",
		Type: 0,
	},
	18: &model.Skill{
		ID:   18,
		Name: "Divination",
		Type: 0,
	},
	19: &model.Skill{
		ID:   19,
		Name: "Dodge",
		Type: 0,
	},
	20: &model.Skill{
		ID:   20,
		Name: "Double Attack",
		Type: 0,
	},
	21: &model.Skill{
		ID:   21,
		Name: "Dragon Punch",
		Type: 0,
	},
	22: &model.Skill{
		ID:   22,
		Name: "Duel Wield",
		Type: 0,
	},
	23: &model.Skill{
		ID:   23,
		Name: "Eagle Strike",
		Type: 0,
	},
	24: &model.Skill{
		ID:   24,
		Name: "Evocation",
		Type: 0,
	},
	25: &model.Skill{
		ID:   25,
		Name: "Feign Death",
		Type: 0,
	},
	26: &model.Skill{
		ID:   26,
		Name: "Flying Kick",
		Type: 0,
	},
	27: &model.Skill{
		ID:   27,
		Name: "Forage",
		Type: 0,
	},
	28: &model.Skill{
		ID:   28,
		Name: "Hand To Hand",
		Type: 0,
	},
	29: &model.Skill{
		ID:   29,
		Name: "Hide",
		Type: 0,
	},
	30: &model.Skill{
		ID:   30,
		Name: "Kick",
		Type: 0,
	},
	31: &model.Skill{
		ID:   31,
		Name: "Meditate",
		Type: 0,
	},
	32: &model.Skill{
		ID:   32,
		Name: "Mend",
		Type: 0,
	},
	33: &model.Skill{
		ID:   33,
		Name: "Offense",
		Type: 0,
	},
	34: &model.Skill{
		ID:   34,
		Name: "Parry",
		Type: 0,
	},
	35: &model.Skill{
		ID:   35,
		Name: "Pick Lock",
		Type: 0,
	},
	36: &model.Skill{
		ID:   36,
		Name: "Piercing",
		Type: 0,
	},
	37: &model.Skill{
		ID:   37,
		Name: "Riposte",
		Type: 0,
	},
	38: &model.Skill{
		ID:   38,
		Name: "Round Kick",
		Type: 0,
	},
	39: &model.Skill{
		ID:   39,
		Name: "Safe Fall",
		Type: 0,
	},
	40: &model.Skill{
		ID:   40,
		Name: "Sense Heading",
		Type: 0,
	},
	41: &model.Skill{
		ID:   41,
		Name: "Sing",
		Type: 0,
	},
	42: &model.Skill{
		ID:   42,
		Name: "Sneak",
		Type: 0,
	},
	43: &model.Skill{
		ID:   43,
		Name: "Specialize Abjure",
		Type: 0,
	},
	44: &model.Skill{
		ID:   44,
		Name: "Specialize Alteration",
		Type: 0,
	},
	45: &model.Skill{
		ID:   45,
		Name: "Specialize Conjuration",
		Type: 0,
	},
	46: &model.Skill{
		ID:   46,
		Name: "Specialize Divinatation",
		Type: 0,
	},
	47: &model.Skill{
		ID:   47,
		Name: "Specialize Evocation",
		Type: 0,
	},
	48: &model.Skill{
		ID:   48,
		Name: "Pick Pockets",
		Type: 0,
	},
	49: &model.Skill{
		ID:   49,
		Name: "Stringed Instruments",
		Type: 0,
	},
	50: &model.Skill{
		ID:   50,
		Name: "Swimming",
		Type: 0,
	},
	51: &model.Skill{
		ID:   51,
		Name: "Throwing",
		Type: 0,
	},
	52: &model.Skill{
		ID:   52,
		Name: "Tiger Claw",
		Type: 0,
	},
	53: &model.Skill{
		ID:   53,
		Name: "Tracking",
		Type: 0,
	},
	54: &model.Skill{
		ID:   54,
		Name: "Wind Instruments",
		Type: 0,
	},
	55: &model.Skill{
		ID:   55,
		Name: "Fishing",
		Type: 1,
	},
	56: &model.Skill{
		ID:   56,
		Name: "Make Poison",
		Type: 1,
	},
	57: &model.Skill{
		ID:   57,
		Name: "Tinkering",
		Type: 1,
	},
	58: &model.Skill{
		ID:   58,
		Name: "Research",
		Type: 1,
	},
	59: &model.Skill{
		ID:   59,
		Name: "Alchemy",
		Type: 1,
	},
	60: &model.Skill{
		ID:   60,
		Name: "Baking",
		Type: 1,
	},
	61: &model.Skill{
		ID:   61,
		Name: "Tailoring",
		Type: 1,
	},
	62: &model.Skill{
		ID:   62,
		Name: "Sense Traps",
		Type: 0,
	},
	63: &model.Skill{
		ID:   63,
		Name: "Blacksmithing",
		Type: 1,
	},
	64: &model.Skill{
		ID:   64,
		Name: "Fletching",
		Type: 1,
	},
	65: &model.Skill{
		ID:   65,
		Name: "Brewing",
		Type: 1,
	},
	66: &model.Skill{
		ID:   66,
		Name: "Alcohol Tolerance",
		Type: 0,
	},
	67: &model.Skill{
		ID:   67,
		Name: "Begging",
		Type: 0,
	},
	68: &model.Skill{
		ID:   68,
		Name: "Jewelry Making",
		Type: 1,
	},
	69: &model.Skill{
		ID:   69,
		Name: "Pottery",
		Type: 1,
	},
	70: &model.Skill{
		ID:   70,
		Name: "Percussion Instruments",
		Type: 0,
	},
	71: &model.Skill{
		ID:   71,
		Name: "Intimidation",
		Type: 0,
	},
	72: &model.Skill{
		ID:   72,
		Name: "Berserking",
		Type: 0,
	},
	73: &model.Skill{
		ID:   73,
		Name: "Taunt ",
		Type: 0,
	},
	74: &model.Skill{
		ID:   74,
		Name: "Frenzy",
		Type: 0,
	},
	75: &model.Skill{
		ID:   75,
		Name: "Remove Traps",
		Type: 0,
	},
	76: &model.Skill{
		ID:   76,
		Name: "Triple Attack",
		Type: 0,
	},
	77: &model.Skill{
		ID:   77,
		Name: "2H Piercing",
		Type: 0,
	},
}
