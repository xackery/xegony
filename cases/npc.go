package cases

import (
	"fmt"
	"strings"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type NpcRepository struct {
	stor storage.Storage
}

func (g *NpcRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *NpcRepository) Get(npcID int64) (npc *model.Npc, err error) {
	if npcID == 0 {
		err = fmt.Errorf("Invalid Npc ID")
		return
	}
	npc, err = g.stor.GetNpc(npcID)
	return
}

func (g *NpcRepository) Create(npc *model.Npc) (err error) {
	if npc == nil {
		err = fmt.Errorf("Empty npc")
		return
	}
	schema, err := npc.NewSchema([]string{"name", "accountID"}, nil)
	if err != nil {
		return
	}

	npc.Id = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(npc))
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
	err = g.stor.CreateNpc(npc)
	if err != nil {
		return
	}
	return
}

func (g *NpcRepository) Search(search string) (npcs []*model.Npc, err error) {
	sanitySearch := strings.Replace(search, " ", "_", -1)
	npcs, err = g.stor.SearchNpc(sanitySearch)
	if err != nil {
		return
	}
	return
}

func (g *NpcRepository) Edit(npcID int64, npc *model.Npc) (err error) {
	schema, err := npc.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(npc))
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

	err = g.stor.EditNpc(npcID, npc)
	if err != nil {
		return
	}
	return
}

func (g *NpcRepository) Delete(npcID int64) (err error) {
	err = g.stor.DeleteNpc(npcID)
	if err != nil {
		return
	}
	return
}

func (g *NpcRepository) List() (npcs []*model.Npc, err error) {
	npcs, err = g.stor.ListNpc()
	if err != nil {
		return
	}
	return
}

func (g *NpcRepository) ListByZone(zoneID int64) (npcs []*model.Npc, err error) {
	npcs, err = g.stor.ListNpcByZone(zoneID)
	if err != nil {
		return
	}
	return
}

func (g *NpcRepository) ListByFaction(factionID int64) (npcs []*model.Npc, err error) {
	npcs, err = g.stor.ListNpcByFaction(factionID)
	if err != nil {
		return
	}
	return
}
