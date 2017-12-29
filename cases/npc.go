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

func (g *NpcRepository) Get(npcId int64) (npc *model.Npc, err error) {
	if npcId == 0 {
		err = fmt.Errorf("Invalid Npc ID")
		return
	}
	npc, err = g.stor.GetNpc(npcId)
	return
}

func (g *NpcRepository) Create(npc *model.Npc) (err error) {
	if npc == nil {
		err = fmt.Errorf("Empty npc")
		return
	}
	schema, err := npc.NewSchema([]string{"name", "accountId"}, nil)
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

func (g *NpcRepository) Edit(npcId int64, npc *model.Npc) (err error) {
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

	err = g.stor.EditNpc(npcId, npc)
	if err != nil {
		return
	}
	return
}

func (g *NpcRepository) Delete(npcId int64) (err error) {
	err = g.stor.DeleteNpc(npcId)
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

func (g *NpcRepository) ListByZone(zoneId int64) (npcs []*model.Npc, err error) {
	npcs, err = g.stor.ListNpcByZone(zoneId)
	if err != nil {
		return
	}
	return
}

func (g *NpcRepository) ListByFaction(factionId int64) (npcs []*model.Npc, err error) {
	npcs, err = g.stor.ListNpcByFaction(factionId)
	if err != nil {
		return
	}
	return
}
