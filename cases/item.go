package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type ItemRepository struct {
	stor storage.Storage
}

func (g *ItemRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *ItemRepository) Get(itemId int64) (item *model.Item, err error) {
	if itemId == 0 {
		err = fmt.Errorf("Invalid Item ID")
		return
	}
	item, err = g.stor.GetItem(itemId)
	return
}

func (g *ItemRepository) Search(search string) (items []*model.Item, err error) {
	items, err = g.stor.SearchItem(search)
	if err != nil {
		return
	}
	return
}

func (g *ItemRepository) SearchByAccount(accountId int64, search string) (items []*model.Item, err error) {
	items, err = g.stor.SearchItemByAccount(accountId, search)
	if err != nil {
		return
	}
	return
}

func (g *ItemRepository) Create(item *model.Item) (err error) {
	if item == nil {
		err = fmt.Errorf("Empty item")
		return
	}
	schema, err := item.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	item.Id = 0 //strip ID
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
	err = g.stor.CreateItem(item)
	if err != nil {
		return
	}
	return
}

func (g *ItemRepository) Edit(itemId int64, item *model.Item) (err error) {
	schema, err := item.NewSchema([]string{"name"}, nil)
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

	err = g.stor.EditItem(itemId, item)
	if err != nil {
		return
	}
	return
}

func (g *ItemRepository) Delete(itemId int64) (err error) {
	err = g.stor.DeleteItem(itemId)
	if err != nil {
		return
	}
	return
}

func (g *ItemRepository) List() (items []*model.Item, err error) {
	items, err = g.stor.ListItem()
	if err != nil {
		return
	}
	return
}

func (g *ItemRepository) ListByCharacter(characterId int64) (items []*model.Item, err error) {
	items, err = g.stor.ListItemByCharacter(characterId)
	if err != nil {
		return
	}
	return
}

func (g *ItemRepository) ListBySlot() (items []*model.Item, err error) {
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

func (g *ItemRepository) GetBySlot(slotId int64) (items []*model.Item, err error) {
	items, err = g.stor.ListItemBySlot(slotId)
	if err != nil {
		return
	}
	return
}

func (g *ItemRepository) GetByZone(zoneId int64) (items []*model.Item, err error) {
	items, err = g.stor.ListItemByZone(zoneId)
	if err != nil {
		return
	}
	return
}
