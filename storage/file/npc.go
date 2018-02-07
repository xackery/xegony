package file

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetNpc will grab data from storage
func (s *Storage) GetNpc(npc *model.Npc) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateNpc will grab data from storage
func (s *Storage) CreateNpc(npc *model.Npc) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListNpc will grab data from storage
func (s *Storage) ListNpc(page *model.Page) (npcs []*model.Npc, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListNpcTotalCount will grab data from storage
func (s *Storage) ListNpcTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListNpcBySearch will grab data from storage
func (s *Storage) ListNpcBySearch(page *model.Page, npc *model.Npc) (npcs []*model.Npc, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListNpcBySearchTotalCount will grab data from storage
func (s *Storage) ListNpcBySearchTotalCount(npc *model.Npc) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListNpcByZone will grab data from storage
func (s *Storage) ListNpcByZone(page *model.Page, zone *model.Zone) (npcs []*model.Npc, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListNpcByZoneTotalCount will grab data from storage
func (s *Storage) ListNpcByZoneTotalCount(zone *model.Zone) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditNpc will grab data from storage
func (s *Storage) EditNpc(npc *model.Npc) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteNpc will grab data from storage
func (s *Storage) DeleteNpc(npc *model.Npc) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
