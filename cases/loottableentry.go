package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type LootTableEntryRepository struct {
	stor storage.Storage
}

func (g *LootTableEntryRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *LootTableEntryRepository) Get(lootTableId int64, lootDropId int64) (lootTableEntry *model.LootTableEntry, err error) {

	lootTableEntry, err = g.stor.GetLootTableEntry(lootTableId, lootDropId)
	return
}

func (g *LootTableEntryRepository) Create(lootTableEntry *model.LootTableEntry) (err error) {
	if lootTableEntry == nil {
		err = fmt.Errorf("Empty lootTableEntry")
		return
	}
	schema, err := lootTableEntry.NewSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootTableEntry))
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
	err = g.stor.CreateLootTableEntry(lootTableEntry)
	if err != nil {
		return
	}
	return
}

func (g *LootTableEntryRepository) Edit(lootTableId int64, lootDropId int64, lootTableEntry *model.LootTableEntry) (err error) {
	schema, err := lootTableEntry.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootTableEntry))
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

	err = g.stor.EditLootTableEntry(lootTableId, lootDropId, lootTableEntry)
	if err != nil {
		return
	}
	return
}

func (g *LootTableEntryRepository) Delete(lootTableId int64, lootDropId int64) (err error) {
	err = g.stor.DeleteLootTableEntry(lootTableId, lootDropId)
	if err != nil {
		return
	}
	return
}

func (g *LootTableEntryRepository) List(lootTableId int64) (lootTableEntrys []*model.LootTableEntry, err error) {
	lootTableEntrys, err = g.stor.ListLootTableEntry(lootTableId)
	if err != nil {
		return
	}
	return
}
