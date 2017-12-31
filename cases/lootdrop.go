package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type LootDropRepository struct {
	stor storage.Storage
}

func (g *LootDropRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *LootDropRepository) Get(lootDropID int64) (lootDrop *model.LootDrop, err error) {
	if lootDropID == 0 {
		err = fmt.Errorf("Invalid LootDrop ID")
		return
	}
	lootDrop, err = g.stor.GetLootDrop(lootDropID)
	return
}

func (g *LootDropRepository) Create(lootDrop *model.LootDrop) (err error) {
	if lootDrop == nil {
		err = fmt.Errorf("Empty lootDrop")
		return
	}
	schema, err := lootDrop.NewSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}

	lootDrop.Id = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(lootDrop))
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
	err = g.stor.CreateLootDrop(lootDrop)
	if err != nil {
		return
	}
	return
}

func (g *LootDropRepository) Edit(lootDropID int64, lootDrop *model.LootDrop) (err error) {
	schema, err := lootDrop.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootDrop))
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

	err = g.stor.EditLootDrop(lootDropID, lootDrop)
	if err != nil {
		return
	}
	return
}

func (g *LootDropRepository) Delete(lootDropID int64) (err error) {
	err = g.stor.DeleteLootDrop(lootDropID)
	if err != nil {
		return
	}
	return
}

func (g *LootDropRepository) List() (lootDrops []*model.LootDrop, err error) {
	lootDrops, err = g.stor.ListLootDrop()
	if err != nil {
		return
	}
	return
}
