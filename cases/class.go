package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

var ()

//ClassRepository handles ClassRepository cases and is a gateway to storage
type ClassRepository struct {
	stor               storage.Storage
	classCache         map[int64]*model.Class
	isClassCacheLoaded bool
}

//Initialize handler
func (c *ClassRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}

	c.stor = stor
	c.isClassCacheLoaded = false
	if err = c.rebuildCache(); err != nil {
		return
	}
	return
}

func (c *ClassRepository) rebuildCache() (err error) {
	if c.isClassCacheLoaded {
		return
	}
	c.isClassCacheLoaded = true
	c.classCache = make(map[int64]*model.Class)
	classs, err := c.list()
	if err != nil {
		return
	}

	for _, class := range classs {
		c.classCache[class.ID] = class
	}
	return
}

//Get handler
func (c *ClassRepository) Get(class *model.Class, user *model.User) (err error) {

	class = c.classCache[class.ID]
	//class, err = c.stor.GetClass(classID)
	return
}

//GetByName gets a class by it's name
func (c *ClassRepository) GetByName(class *model.Class, user *model.User) (err error) {
	for _, classC := range c.classCache {
		if classC.Name == class.Name {
			class = classC
			return
		}
	}
	return
}

//Create handler
func (c *ClassRepository) Create(class *model.Class, user *model.User) (err error) {
	if class == nil {
		err = fmt.Errorf("Empty class")
		return
	}
	schema, err := c.newSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}

	class.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(class))
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
	//err = c.stor.CreateClass(class)
	//if err != nil {
	//	return
	//}
	c.isClassCacheLoaded = false
	c.rebuildCache()
	return
}

//Edit handler
func (c *ClassRepository) Edit(class *model.Class, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(class))
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

	//if err = c.stor.EditClass(classID, class); err != nil {
	//	return
	//}
	if err = c.rebuildCache(); err != nil {
		return
	}
	return
}

//Delete handler
func (c *ClassRepository) Delete(class *model.Class, user *model.User) (err error) {
	//err = c.stor.DeleteClass(classID)
	//if err != nil {
	//	return
	//}
	//if err = c.rebuildCache(); err != nil {
	//	return
	//}
	return
}

func (c *ClassRepository) list() (classes []*model.Class, err error) {
	for _, class := range classes {
		classes = append(classes, class)
	}
	return
}

//List handler
func (c *ClassRepository) List(user *model.User) (classes []*model.Class, err error) {
	for _, class := range c.classCache {
		classes = append(classes, class)
	}
	return
}

func (c *ClassRepository) prepare(class *model.Class, user *model.User) (err error) {

	return
}

//newSchema handler
func (c *ClassRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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
func (c *ClassRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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

var classes = map[int64]*model.Class{
	1: {
		ID:        1,
		Bit:       1,
		Name:      "Warrior",
		ShortName: "WAR",
		Icon:      "xa-shield",
	},
	2: {
		ID:        2,
		Bit:       2,
		Name:      "Cleric",
		ShortName: "CLR",
		Icon:      "xa-ankh",
	},
	3: {
		ID:        3,
		Bit:       4,
		Name:      "Paladin",
		ShortName: "PAL",
		Icon:      "xa-fireball-sword",
	},
	4: {
		ID:        4,
		Bit:       8,
		Name:      "Ranger",
		ShortName: "RNG",
		Icon:      "xa-arrow-cluster",
	},
	5: {
		ID:        5,
		Bit:       512,
		Name:      "Shadow Knight",
		ShortName: "SHD",
		Icon:      "xa-fireball-sword",
	},
	6: {
		ID:        6,
		Bit:       32,
		Name:      "Druid",
		ShortName: "DRU",
		Icon:      "xa-leaf",
	},
	7: {
		ID:        7,
		Bit:       64,
		Name:      "Monk",
		ShortName: "MNK",
		Icon:      "xa-hand-emblem",
	},
	8: {
		ID:        8,
		Bit:       128,
		Name:      "Bard",
		ShortName: "BRD",
		Icon:      "xa-ocarina",
	},
	9: {
		ID:        9,
		Bit:       256,
		Name:      "Rogue",
		ShortName: "ROG",
		Icon:      "xa-hood",
	},
	10: {
		ID:        10,
		Bit:       16,
		Name:      "Shaman",
		ShortName: "SHM",
		Icon:      "xa-incense",
	},
	11: {
		ID:        11,
		Bit:       1024,
		Name:      "Necromancer",
		ShortName: "NEC",
		Icon:      "xa-skull",
	},
	12: {
		ID:        12,
		Bit:       2048,
		Name:      "Wizard",
		ShortName: "WIZ",
		Icon:      "xa-fire",
	},
	13: {
		ID: 13,

		Bit:       4096,
		Name:      "Magician",
		ShortName: "MAG",
		Icon:      "xa-burning-book",
	},
	14: {
		ID:        14,
		Bit:       8192,
		Name:      "Enchanter",
		ShortName: "ENC",
		Icon:      "xa-crystal-ball",
	},
	15: {
		ID:        15,
		Bit:       16384,
		Name:      "Beastlord",
		ShortName: "BST",
		Icon:      "xa-pawprint",
	},
	16: {
		ID:        16,
		Bit:       32769,
		Name:      "Berserker",
		ShortName: "BER",
		Icon:      "xa-axe",
	},
	20: {
		ID:   20,
		Name: "GM Warrior",
	},
	21: {
		ID:   21,
		Name: "GM Cleric",
	},
	22: {
		ID:   22,
		Name: "GM Paladin",
	},
	23: {
		ID:   23,
		Name: "GM Ranger",
	},
	24: {
		ID:   24,
		Name: "GM Shadow Knight",
	},
	25: {
		ID:   25,
		Name: "GM Druid",
	},
	26: {
		ID:   26,
		Name: "GM Monk",
	},
	27: {
		ID:   27,
		Name: "GM Bard",
	},
	28: {
		ID:   28,
		Name: "GM Rogue",
	},
	29: {
		ID:   29,
		Name: "GM Shaman",
	},
	30: {
		ID:   30,
		Name: "GM Necromancer",
	},
	31: {
		ID:   31,
		Name: "GM Wizard",
	},
	32: {
		ID:   32,
		Name: "GM Magician",
	},
	33: {
		ID:   33,
		Name: "GM Enchanter",
	},
	34: {
		ID:   34,
		Name: "GM Beastlord",
	},
	35: {
		ID:   35,
		Name: "GM Berserker",
	},
	40: {
		ID:   40,
		Name: "Banker",
	},
	41: {
		ID:   41,
		Name: "Shopkeeper",
	},
	59: {
		ID:   59,
		Name: "Discord Merchant",
	},
	60: {
		ID:   60,
		Name: "Adventure Recruiter",
	},
	61: {
		ID:   61,
		Name: "Adventure Merchant",
	},
	63: {
		ID:   63,
		Name: "Tribute Master",
	},
	64: {
		ID:   64,
		Name: "Guild Tribute Master?",
	},
	66: {
		ID:   66,
		Name: "Guild Bank",
	},
	67: {
		ID:   67,
		Name: "Radiant Crystal Merchant",
	},
	68: {
		ID:   68,
		Name: "Ebon Crystal Merchant",
	},
	69: {
		ID:   69,
		Name: "Fellowships",
	},
	70: {
		ID:   70,
		Name: "Alternate Currency Merchant",
	},
	71: {
		ID:   71,
		Name: "Mercenary Merchant",
	},
}
