package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type LootTableRepository struct {
	stor storage.Storage
}

func (g *LootTableRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *LootTableRepository) Get(lootTableID int64) (lootTable *model.LootTable, err error) {
	if lootTableID == 0 {
		err = fmt.Errorf("Invalid LootTable ID")
		return
	}
	lootTable, err = g.stor.GetLootTable(lootTableID)
	return
}

func (g *LootTableRepository) Create(lootTable *model.LootTable) (err error) {
	if lootTable == nil {
		err = fmt.Errorf("Empty lootTable")
		return
	}
	schema, err := lootTable.NewSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}

	lootTable.Id = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(lootTable))
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
	err = g.stor.CreateLootTable(lootTable)
	if err != nil {
		return
	}
	return
}

func (g *LootTableRepository) Edit(lootTableID int64, lootTable *model.LootTable) (err error) {
	schema, err := lootTable.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(lootTable))
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

	err = g.stor.EditLootTable(lootTableID, lootTable)
	if err != nil {
		return
	}
	return
}

func (g *LootTableRepository) Delete(lootTableID int64) (err error) {
	err = g.stor.DeleteLootTable(lootTableID)
	if err != nil {
		return
	}
	return
}

func (g *LootTableRepository) List() (lootTables []*model.LootTable, err error) {
	lootTables, err = g.stor.ListLootTable()
	if err != nil {
		return
	}
	return
}
