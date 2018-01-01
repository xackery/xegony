package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//ItemRepository handles ItemRepository cases and is a gateway to storage
type ItemRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *ItemRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *ItemRepository) Get(itemID int64) (item *model.Item, err error) {
	if itemID == 0 {
		err = fmt.Errorf("Invalid Item ID")
		return
	}
	item, err = c.stor.GetItem(itemID)
	return
}

//Search handles logic
func (c *ItemRepository) Search(search string) (items []*model.Item, err error) {
	items, err = c.stor.SearchItem(search)
	if err != nil {
		return
	}
	return
}

//SearchByAccount handles logic
func (c *ItemRepository) SearchByAccount(accountID int64, search string) (items []*model.Item, err error) {
	items, err = c.stor.SearchItemByAccount(accountID, search)
	if err != nil {
		return
	}
	return
}

//Create handles logic
func (c *ItemRepository) Create(item *model.Item) (err error) {
	if item == nil {
		err = fmt.Errorf("Empty item")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	item.ID = 0 //strip ID
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
	err = c.stor.CreateItem(item)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *ItemRepository) Edit(itemID int64, item *model.Item) (err error) {
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

	err = c.stor.EditItem(itemID, item)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *ItemRepository) Delete(itemID int64) (err error) {
	err = c.stor.DeleteItem(itemID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *ItemRepository) List(pageSize int64, pageNumber int64) (items []*model.Item, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	items, err = c.stor.ListItem(pageSize, pageNumber)
	if err != nil {
		return
	}
	return
}

//ListCount handles logic
func (c *ItemRepository) ListCount() (count int64, err error) {

	count, err = c.stor.ListItemCount()
	if err != nil {
		return
	}
	return
}

//ListByCharacter handles logic
func (c *ItemRepository) ListByCharacter(characterID int64) (items []*model.Item, err error) {
	items, err = c.stor.ListItemByCharacter(characterID)
	if err != nil {
		return
	}
	return
}

//ListBySlot handles logic
func (c *ItemRepository) ListBySlot() (items []*model.Item, err error) {
	items = []*model.Item{
		&model.Item{
			Itemtype: 0,
			Name:     "1 Hand Slash",
		},
		&model.Item{
			Itemtype: 1,
			Name:     "2 Hand Slash",
		},
		&model.Item{
			Itemtype: 2,
			Name:     "Piercing",
		},
		&model.Item{
			Itemtype: 3,
			Name:     "1 Hand Blunt",
		},
		&model.Item{
			Itemtype: 4,
			Name:     "2 Hand Blunt",
		},
		&model.Item{
			Itemtype: 5,
			Name:     "Archery",
		},
		/*&model.Item{
			Itemtype: 6,
			Name:     "Unused",
		},*/
		&model.Item{
			Itemtype: 7,
			Name:     "Throwing",
		},
		&model.Item{
			Itemtype: 8,
			Name:     "Shields",
		},
		/*&model.Item{
			Itemtype: 9,
			Name:     "Unused",
		},*/
		&model.Item{
			Itemtype: 10,
			Name:     "Armor",
		},
		&model.Item{
			Itemtype: 11,
			Name:     "Involves Tradeskills",
		},
		&model.Item{
			Itemtype: 12,
			Name:     "Lock Picking",
		},
		/*&model.Item{
			Itemtype: 13,
			Name:     "Unused",
		},*/
		&model.Item{
			Itemtype: 14,
			Name:     "Food",
		},
		&model.Item{
			Itemtype: 15,
			Name:     "Drink",
		},
		&model.Item{
			Itemtype: 16,
			Name:     "Light Source",
		},
		&model.Item{
			Itemtype: 17,
			Name:     "Common Inventory Item",
		},
		&model.Item{
			Itemtype: 18,
			Name:     "Bind Wound",
		},
		&model.Item{
			Itemtype: 19,
			Name:     "Thrown Casting Items",
		},
		&model.Item{
			Itemtype: 20,
			Name:     "Spells / Song Sheets",
		},
		&model.Item{
			Itemtype: 21,
			Name:     "Potions",
		},
		&model.Item{
			Itemtype: 22,
			Name:     "Fletched Arrows",
		},
		&model.Item{
			Itemtype: 23,
			Name:     "Wind Instruments",
		},
		&model.Item{
			Itemtype: 24,
			Name:     "Stringed Instruments",
		},
		&model.Item{
			Itemtype: 25,
			Name:     "Brass Instruments",
		},
		&model.Item{
			Itemtype: 26,
			Name:     "Drum Instruments",
		},
		&model.Item{
			Itemtype: 27,
			Name:     "Ammo",
		},
		/*&model.Item{
			Itemtype: 28,
			Name:     "Unused",
		},*/
		&model.Item{
			Itemtype: 29,
			Name:     "Jewlery Items",
		},
		/*&model.Item{
			Itemtype: 30,
			Name:     "Unused",
		},*/
		&model.Item{
			Itemtype: 31,
			Name:     "Rolled Up Notes and Scrolls",
		},
		&model.Item{
			Itemtype: 32,
			Name:     "Books",
		},
		&model.Item{
			Itemtype: 33,
			Name:     "Keys",
		},
		&model.Item{
			Itemtype: 34,
			Name:     "Odd Items",
		},
		&model.Item{
			Itemtype: 35,
			Name:     "2 Hand Pierce",
		},
		&model.Item{
			Itemtype: 36,
			Name:     "Fishing Poles",
		},
		&model.Item{
			Itemtype: 37,
			Name:     "Fishing Bait",
		},
		&model.Item{
			Itemtype: 38,
			Name:     "Alcoholic Beverages",
		},
		&model.Item{
			Itemtype: 39,
			Name:     "More Keys",
		},
		&model.Item{
			Itemtype: 40,
			Name:     "Compasses",
		},
		/*&model.Item{
			Itemtype: 41,
			Name:     "Unused",
		},*/
		&model.Item{
			Itemtype: 42,
			Name:     "Poisons",
		},
		/*&model.Item{
			Itemtype: 43,
			Name:     "Unused",
		},
		&model.Item{
			Itemtype: 44,
			Name:     "Unused",
		},*/
		&model.Item{
			Itemtype: 45,
			Name:     "Hand to Hand",
		},
		/*&model.Item{
			Itemtype: 46,
			Name:     "Unused",
		},
		&model.Item{
			Itemtype: 47,
			Name:     "Unused",
		},
		&model.Item{
			Itemtype: 48,
			Name:     "Unused",
		},
		&model.Item{
			Itemtype: 49,
			Name:     "Unused",
		},
		&model.Item{
			Itemtype: 50,
			Name:     "Unused",
		},
		&model.Item{
			Itemtype: 51,
			Name:     "Unused",
		},*/
		&model.Item{
			Itemtype: 52,
			Name:     "Charms",
		},
		&model.Item{
			Itemtype: 53,
			Name:     "Dyes",
		},
		&model.Item{
			Itemtype: 54,
			Name:     "Augments",
		},
		&model.Item{
			Itemtype: 55,
			Name:     "Augment Solvents",
		},
		&model.Item{
			Itemtype: 56,
			Name:     "Augment Distillers",
		},
		&model.Item{
			Itemtype: 58,
			Name:     "Fellowship Banner Materials",
		},
		&model.Item{
			Itemtype: 60,
			Name:     "Cultural Armor Manuals",
		},
		&model.Item{
			Itemtype: 63,
			Name:     "Currencies",
		},
	}

	return
}

//GetBySlot handles logic
func (c *ItemRepository) GetBySlot(slotID int64) (items []*model.Item, err error) {
	items, err = c.stor.ListItemBySlot(slotID)
	if err != nil {
		return
	}
	return
}

//GetByZone handles logic
func (c *ItemRepository) GetByZone(zoneID int64) (items []*model.Item, err error) {
	items, err = c.stor.ListItemByZone(zoneID)
	if err != nil {
		return
	}
	return
}

func (c *ItemRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *ItemRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
