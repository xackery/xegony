package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	npcsDatabase = []*model.Npc{}
	npcLock      = sync.RWMutex{}
)

//GetNpc will grab data from storage
func (s *Storage) GetNpc(npc *model.Npc) (err error) {
	npcLock.RLock()
	defer npcLock.RUnlock()
	for _, tmpNpc := range npcsDatabase {
		if tmpNpc.ID == npc.ID {
			*npc = *tmpNpc
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateNpc will grab data from storage
func (s *Storage) CreateNpc(npc *model.Npc) (err error) {
	npcLock.Lock()
	defer npcLock.Unlock()
	for _, tmpNpc := range npcsDatabase {
		if tmpNpc.ID == npc.ID {
			err = fmt.Errorf("npc already exists")
			return
		}
	}
	npcsDatabase = append(npcsDatabase, npc)
	return
}

//ListNpc will grab data from storage
func (s *Storage) ListNpc(page *model.Page) (npcs []*model.Npc, err error) {
	npcLock.RLock()
	defer npcLock.RUnlock()

	npcs = make([]*model.Npc, len(npcsDatabase))

	npcs = npcsDatabase

	switch page.OrderBy {
	case "short_name":
		sort.Slice(npcs, func(i, j int) bool {
			return npcs[i].Name < npcs[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(npcs))
		}
	*/
	return
}

//ListNpcTotalCount will grab data from storage
func (s *Storage) ListNpcTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(npcsDatabase))
	return
}

//ListNpcBySearch will grab data from storage
func (s *Storage) ListNpcBySearch(page *model.Page, npc *model.Npc) (npcs []*model.Npc, err error) {
	npcLock.RLock()
	defer npcLock.RUnlock()

	if len(npc.Name) > 0 {
		for i := range npcsDatabase {
			if strings.Contains(npcsDatabase[i].Name, npc.Name) {
				npcs = append(npcs, npcsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "short_name":
		sort.Slice(npcs, func(i, j int) bool {
			return npcs[i].Name < npcs[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(npcs))
	//}
	return
}

//ListNpcBySearchTotalCount will grab data from storage
func (s *Storage) ListNpcBySearchTotalCount(npc *model.Npc) (count int64, err error) {
	npcLock.RLock()
	defer npcLock.RUnlock()

	npcs := []*model.Npc{}
	if len(npc.Name) > 0 {
		for i := range npcsDatabase {
			if strings.Contains(npcsDatabase[i].Name, npc.Name) {
				npcs = append(npcs, npcsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(npcs))
	return
}

//EditNpc will grab data from storage
func (s *Storage) EditNpc(npc *model.Npc) (err error) {
	npcLock.Lock()
	defer npcLock.Unlock()
	for i := range npcsDatabase {
		if npcsDatabase[i].ID == npc.ID {
			*npcsDatabase[i] = *npc
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteNpc will grab data from storage
func (s *Storage) DeleteNpc(npc *model.Npc) (err error) {
	npcLock.Lock()
	defer npcLock.Unlock()
	indexToDelete := 0
	for i := range npcsDatabase {
		if npcsDatabase[i].ID == npc.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	npcsDatabase[len(npcsDatabase)-1], npcsDatabase[indexToDelete] = npcsDatabase[indexToDelete], npcsDatabase[len(npcsDatabase)-1]
	npcsDatabase = npcsDatabase[:len(npcsDatabase)-1]
	return
}
