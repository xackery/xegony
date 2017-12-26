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
			Name:     "1 Hand Blunt",
			Itemtype: 3,
		},
		&model.Item{
			Name:     "2 Hand Blunt",
			Itemtype: 4,
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
