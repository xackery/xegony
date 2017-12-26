package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type FactionRepository struct {
	stor storage.Storage
}

func (g *FactionRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *FactionRepository) Get(factionId int64) (faction *model.Faction, err error) {
	if factionId == 0 {
		err = fmt.Errorf("Invalid Faction ID")
		return
	}
	faction, err = g.stor.GetFaction(factionId)
	return
}

func (g *FactionRepository) Create(faction *model.Faction) (err error) {
	if faction == nil {
		err = fmt.Errorf("Empty faction")
		return
	}
	schema, err := faction.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	faction.Id = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(faction))
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
	err = g.stor.CreateFaction(faction)
	if err != nil {
		return
	}
	return
}

func (g *FactionRepository) Edit(factionId int64, faction *model.Faction) (err error) {
	schema, err := faction.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(faction))
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

	err = g.stor.EditFaction(factionId, faction)
	if err != nil {
		return
	}
	return
}

func (g *FactionRepository) Delete(factionId int64) (err error) {
	err = g.stor.DeleteFaction(factionId)
	if err != nil {
		return
	}
	return
}

func (g *FactionRepository) List() (factions []*model.Faction, err error) {
	factions, err = g.stor.ListFaction()
	if err != nil {
		return
	}
	return
}
