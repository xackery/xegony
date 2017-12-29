package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type LootDropEntryRepository struct {
	stor storage.Storage
}

func (g *LootDropEntryRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *LootDropEntryRepository) Get(lootDropId int64, itemId int64) (lootDropEntry *model.LootDropEntry, err error) {

	lootDropEntry, err = g.stor.GetLootDropEntry(lootDropId, itemId)
	return
}

func (g *LootDropEntryRepository) Create(lootDropEntry *model.LootDropEntry) (err error) {
	if lootDropEntry == nil {
		err = fmt.Errorf("Empty lootDropEntry")
		return
	}
	schema, err := lootDropEntry.NewSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootDropEntry))
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
	err = g.stor.CreateLootDropEntry(lootDropEntry)
	if err != nil {
		return
	}
	return
}

func (g *LootDropEntryRepository) Edit(lootDropId int64, itemId int64, lootDropEntry *model.LootDropEntry) (err error) {
	schema, err := lootDropEntry.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootDropEntry))
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

	err = g.stor.EditLootDropEntry(lootDropId, itemId, lootDropEntry)
	if err != nil {
		return
	}
	return
}

func (g *LootDropEntryRepository) Delete(lootDropId int64, itemId int64) (err error) {
	err = g.stor.DeleteLootDropEntry(lootDropId, itemId)
	if err != nil {
		return
	}
	return
}

func (g *LootDropEntryRepository) List(lootDropId int64) (lootDropEntrys []*model.LootDropEntry, err error) {
	lootDropEntrys, err = g.stor.ListLootDropEntry(lootDropId)
	if err != nil {
		return
	}
	return
}
