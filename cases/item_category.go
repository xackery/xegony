package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//ItemCategoryRepository handles ItemCategoryRepository cases and is a gateway to storage
type ItemCategoryRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *ItemCategoryRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	//Item Category is a locally stored system
	return
}

//Get handles logic
func (c *ItemCategoryRepository) Get(itemCategoryID int64, item *model.Item, user *model.User) (itemCategory *model.ItemCategory, err error) {

	var ok bool
	itemCategory, ok = itemCategories[itemCategoryID]
	if !ok {
		return
	}
	return
}

//GetByItem handles logic
func (c *ItemCategoryRepository) GetByItem(item *model.Item, user *model.User) (itemCategory *model.ItemCategory, err error) {

	var ok bool
	itemCategory, ok = itemCategories[item.Itemtype]
	if !ok {
		return
	}
	return
}

//Create handles logic
func (c *ItemCategoryRepository) Create(itemCategory *model.ItemCategory, item *model.Item, user *model.User) (err error) {

	schema, err := c.newSchema(nil, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(item))
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
	itemCategories[itemCategory.ID] = itemCategory
	return
}

//Edit handles logic
func (c *ItemCategoryRepository) Edit(itemCategoryID int64, itemCategory *model.ItemCategory, item *model.Item, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(item))
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

	itemCategories[itemCategoryID] = itemCategory
	return
}

//Delete handles logic
func (c *ItemCategoryRepository) Delete(itemCategory *model.ItemCategory, item *model.Item, user *model.User) (err error) {
	delete(itemCategories, itemCategory.ID)
	return
}

//List handles logic
func (c *ItemCategoryRepository) List(user *model.User) (itemCategories []*model.ItemCategory, err error) {
	for _, itemCategory := range itemCategories {
		itemCategories = append(itemCategories, itemCategory)
	}
	return
}

func (c *ItemCategoryRepository) prepare(item *model.Item) (err error) {

	return
}

func (c *ItemCategoryRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *ItemCategoryRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "zoneID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}

var itemCategories = map[int64]*model.ItemCategory{
	0: {
		ID:   0,
		Icon: "xa-crossed-swords", //0:
		Name: "1 Hand Slash",
	},
	1: {
		ID:   1,
		Icon: "xa-croc-sword", //1:
		Name: "2 Hand Slash",
	},
	2: {
		ID:   2,
		Icon: "xa-plain-dagger", //2:
		Name: "Piercing",
	},
	3: {
		ID:   3,
		Icon: "xa-flat-hammer", //3:
		Name: "1 Hand Blunt",
	},
	4: {
		ID:   4,
		Icon: "xa-gavel", //4:
		Name: "2 Hand Blunt",
	},
	5: {
		ID:   5,
		Icon: "xa-crossbow", //5:
		Name: "Archery",
	},
	6: {
		ID:   6,
		Icon: "xa-help", //6:
		Name: "Unused",
	},
	7: {
		ID:   7,
		Icon: "xa-hammer-drop", //7:
		Name: "Throwing",
	},
	8: {
		ID:   8,
		Icon: "xa-fire-shield", //8:
		Name: "Shields",
	},
	9: {
		ID:   9,
		Icon: "xa-help", //9:
		Name: "Unused",
	},
	10: {
		ID:   10,
		Icon: "xa-vest", //10:
		Name: "Armor",
	},
	11: {
		ID:   11,
		Icon: "xa-archery-target", //Involves Tradeskills (Not sure how), //11:
		Name: "Involves Tradeskills",
	},
	12: {
		ID:   12,
		Icon: "xa-key", //12:
		Name: "Lock Picking",
	},
	13: {
		ID:   13,
		Icon: "xa-help", //13:
		Name: "Unused",
	},
	14: {
		ID:   14,
		Icon: "xa-apple", //14:
		Name: "Food",
	},
	15: {
		ID:   15,
		Icon: "xa-brandy-bottle", //15:
		Name: "Drink",
	},
	16: {
		ID:   16,
		Icon: "xa-light-bulb", //16:
		Name: "Light Source",
	},
	17: {
		ID:   17,
		Icon: "xa-shovel", //17:
		Name: "Common Inventory Item",
	},
	18: {
		ID:   18,
		Icon: "xa-health", //18:
		Name: "Bind Wound",
	},
	19: {
		ID:   19,
		Icon: "xa-bottled-bolt", //19:
		Name: "Thrown Casting Items",
	},
	20: {
		ID:   20,
		Icon: "xa-scroll-unfurled", //20:
		Name: "Spells / Song Sheets",
	},
	21: {
		ID:   21,
		Icon: "xa-flask", //21:
		Name: "Potions",
	},
	22: {
		ID:   22,
		Icon: "xa-arrow-flights", //22:
		Name: "Fletched Arrows",
	},
	23: {
		ID:   23,
		Icon: "xa-ocarina", //23:
		Name: "Wind Instruments",
	},
	24: {
		ID:   24,
		Icon: "xa-ocarina", //24:
		Name: "Stringed Instruments",
	},
	25: {
		ID:   25,
		Icon: "xa-ocarina", //25:
		Name: "Brass Instruments",
	},
	26: {
		ID:   26,
		Icon: "xa-ocarina", //26:
		Name: "Drum Instruments",
	},
	27: {
		ID:   27,
		Icon: "xa-broadhead-arrow", //27:
		Name: "Ammo",
	},
	28: {
		ID:   28,
		Icon: "xa-help", //28:
		Name: "Unused",
	},
	29: {
		ID:   29,
		Icon: "xa-explosion", //29:
		Name: "Jewlery Items",
	},
	30: {
		ID:   30,
		Icon: "xa-help", //30:
		Name: "Unused",
	},
	31: {
		ID:   31,
		Icon: "xa-book", //Usually Readable Notes and Scrolls", //31:
		Name: "Rolled Up Notes and Scrolls",
	},
	32: {
		ID:   32,
		Icon: "xa-book", //Usually Readable Books", //32:
		Name: "Books",
	},
	33: {
		ID:   33,
		Icon: "xa-key", //33:
		Name: "Keys",
	},
	34: {
		ID:   34,
		Icon: "xa-vail", //Odd Items (Not sure what they are for)", //34:
		Name: "Odd Items",
	},
	35: {
		ID:   35,
		Icon: "xa-relic-blade", //2hp, //35:
		Name: "2 Hand Pierce",
	},
	36: {
		ID:   36,
		Icon: "xa-fish", //36:
		Name: "Fishing Poles",
	},
	37: {
		ID:   37,
		Icon: "xa-venomous-snake", //37:
		Name: "Fishing Bait",
	},
	38: {
		ID:   38,
		Icon: "xa-beer", //38:
		Name: "Alcoholic Beverages",
	},
	39: {
		ID:   39,
		Icon: "xa-key", //39:
		Name: "More Keys",
	},
	40: {
		ID:   40,
		Icon: "xa-compass", //40:
		Name: "Compasses",
	},
	41: {
		ID:   41,
		Icon: "xa-help", //41:
		Name: "Unused",
	},
	42: {
		ID:   42,
		Icon: "xa-bottle-vapors", //42:
		Name: "Poisons",
	},
	43: {
		ID:   43,
		Icon: "xa-help", //43:
		Name: "Unused",
	},
	44: {
		ID:   44,
		Icon: "xa-help", //44:
		Name: "Unused",
	},
	45: {
		ID:   45,
		Icon: "xa-hand", //45:
		Name: "Hand to Hand",
	},
	46: {
		ID:   46,
		Icon: "xa-help", //46:
		Name: "Unused",
	},
	47: {
		ID:   47,
		Icon: "xa-help", //47:
		Name: "Unused",
	},
	48: {
		ID:   48,
		Icon: "xa-help", //48:
		Name: "Unused",
	},
	49: {
		ID:   49,
		Icon: "xa-help", //49:
		Name: "Unused",
	},
	50: {
		ID:   50,
		Icon: "xa-help", //50:
		Name: "Unused",
	},
	51: {
		ID:   51,
		Icon: "xa-help", //51:
		Name: "Unused",
	},
	52: {
		ID:   52,
		Icon: "xa-sapphire", //52:
		Name: "Charms",
	},
	53: {
		ID:   53,
		Icon: "xa-round-bottome-flask", //53:
		Name: "Dyes",
	},
	54: {
		ID:   54,
		Icon: "xa-bubbling-potion", //54:
		Name: "Augments",
	},
	55: {
		ID:   55,
		Icon: "xa-corked-tube", //55:
		Name: "Augment Solvents",
	},
	56: {
		ID:   56,
		Icon: "xa-corked-tube", //56:
		Name: "Augment Distillers",
	},
	58: {
		ID:   58,
		Icon: "xa-castle-flag", //58:
		Name: "Fellowship Banner Materials",
	},
	60: {
		ID:   60,
		Icon: "xa-book", //60:
		Name: "Cultural Armor Manuals",
	},
	63: {
		ID:   63,
		Icon: "xa-sapphire", //63:
		Name: "Currencies",
	},
}
