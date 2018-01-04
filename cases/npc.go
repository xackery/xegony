package cases

import (
	"fmt"
	"strings"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//NpcRepository handles NpcRepository cases and is a gateway to storage
type NpcRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *NpcRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *NpcRepository) Get(npcID int64) (npc *model.Npc, err error) {
	if npcID == 0 {
		err = fmt.Errorf("Invalid Npc ID")
		return
	}
	npc, err = c.stor.GetNpc(npcID)
	return
}

//Create handles logic
func (c *NpcRepository) Create(npc *model.Npc) (err error) {
	if npc == nil {
		err = fmt.Errorf("Empty npc")
		return
	}
	schema, err := c.newSchema([]string{"name", "accountID"}, nil)
	if err != nil {
		return
	}

	npc.ID = 0 //strip ID
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
	err = c.stor.CreateNpc(npc)
	if err != nil {
		return
	}
	return
}

//Search handles logic
func (c *NpcRepository) Search(search string) (npcs []*model.Npc, err error) {
	sanitySearch := strings.Replace(search, " ", "_", -1)
	npcs, err = c.stor.SearchNpc(sanitySearch)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *NpcRepository) Edit(npcID int64, npc *model.Npc) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
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

	err = c.stor.EditNpc(npcID, npc)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *NpcRepository) Delete(npcID int64) (err error) {
	err = c.stor.DeleteNpc(npcID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *NpcRepository) List() (npcs []*model.Npc, err error) {
	npcs, err = c.stor.ListNpc()
	if err != nil {
		return
	}
	return
}

//ListBySpell handles logic
func (c *NpcRepository) ListBySpell(spellID int64) (npcs []*model.Npc, err error) {
	npcs, err = c.stor.ListNpcBySpell(spellID)
	if err != nil {
		return
	}
	return
}

//ListByItem handles logic
func (c *NpcRepository) ListByItem(itemID int64) (npcs []*model.Npc, err error) {
	npcs, err = c.stor.ListNpcByItem(itemID)
	if err != nil {
		return
	}
	return
}

//ListByZone handles logic
func (c *NpcRepository) ListByZone(zoneID int64) (npcs []*model.Npc, err error) {
	npcs, err = c.stor.ListNpcByZone(zoneID)
	if err != nil {
		return
	}
	return
}

//ListByFaction handles logic
func (c *NpcRepository) ListByFaction(factionID int64) (npcs []*model.Npc, err error) {
	npcs, err = c.stor.ListNpcByFaction(factionID)
	if err != nil {
		return
	}
	return
}

func (c *NpcRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *NpcRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
