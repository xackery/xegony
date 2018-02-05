package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	deitysDatabase = []*model.Deity{}
	deityLock      = sync.RWMutex{}
)

//GetDeity will grab data from storage
func (s *Storage) GetDeity(deity *model.Deity) (err error) {
	deityLock.RLock()
	defer deityLock.RUnlock()
	for _, tmpDeity := range deitysDatabase {
		if tmpDeity.ID == deity.ID {
			*deity = *tmpDeity
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//GetDeityBySpell will grab data from storage
func (s *Storage) GetDeityBySpell(spell *model.Spell, deity *model.Deity) (err error) {
	deityLock.RLock()
	defer deityLock.RUnlock()
	for _, tmpDeity := range deitysDatabase {
		if tmpDeity.SpellID == deity.SpellID {
			*deity = *tmpDeity
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateDeity will grab data from storage
func (s *Storage) CreateDeity(deity *model.Deity) (err error) {
	deityLock.Lock()
	defer deityLock.Unlock()
	for _, tmpDeity := range deitysDatabase {
		if tmpDeity.ID == deity.ID {
			err = fmt.Errorf("deity already exists")
			return
		}
	}
	deitysDatabase = append(deitysDatabase, deity)
	return
}

//ListDeity will grab data from storage
func (s *Storage) ListDeity(page *model.Page) (deitys []*model.Deity, err error) {
	deityLock.RLock()
	defer deityLock.RUnlock()

	deitys = make([]*model.Deity, len(deitysDatabase))

	deitys = deitysDatabase

	switch page.OrderBy {
	case "name":
		sort.Slice(deitys, func(i, j int) bool {
			return deitys[i].Name < deitys[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(deitys))
		}
	*/
	return
}

//ListDeityTotalCount will grab data from storage
func (s *Storage) ListDeityTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(deitysDatabase))
	return
}

//ListDeityByBit will grab data from storage
func (s *Storage) ListDeityByBit(page *model.Page, deity *model.Deity) (deitys []*model.Deity, err error) {
	deityLock.RLock()
	defer deityLock.RUnlock()

	for i := range deitysDatabase {
		if deity.Bit < 1 || deitysDatabase[i].Bit < 1 {
			continue
		}
		if deity.Bit&deitysDatabase[i].Bit == deitysDatabase[i].Bit {
			deitys = append(deitys, deitysDatabase[i])
		}
	}
	switch page.OrderBy {
	case "name":
		sort.Slice(deitys, func(i, j int) bool {
			return deitys[i].Name < deitys[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(deitys))
	//}
	return
}

//ListDeityByBitTotalCount will grab data from storage
func (s *Storage) ListDeityByBitTotalCount(deity *model.Deity) (count int64, err error) {
	deityLock.RLock()
	defer deityLock.RUnlock()

	deitys := []*model.Deity{}
	for i := range deitysDatabase {
		if deity.Bit < 1 || deitysDatabase[i].Bit < 1 {
			continue
		}
		if deity.Bit&deitysDatabase[i].Bit == deitysDatabase[i].Bit {
			deitys = append(deitys, deitysDatabase[i])
		}
	}
	count = int64(len(deitys))
	return
}

//ListDeityBySearch will grab data from storage
func (s *Storage) ListDeityBySearch(page *model.Page, deity *model.Deity) (deitys []*model.Deity, err error) {
	deityLock.RLock()
	defer deityLock.RUnlock()

	if len(deity.Name) > 0 {
		for i := range deitysDatabase {
			if strings.Contains(deitysDatabase[i].Name, deity.Name) {
				deitys = append(deitys, deitysDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(deitys, func(i, j int) bool {
			return deitys[i].Name < deitys[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(deitys))
	//}
	return
}

//ListDeityBySearchTotalCount will grab data from storage
func (s *Storage) ListDeityBySearchTotalCount(deity *model.Deity) (count int64, err error) {
	deityLock.RLock()
	defer deityLock.RUnlock()

	deitys := []*model.Deity{}
	if len(deity.Name) > 0 {
		for i := range deitysDatabase {
			if strings.Contains(deitysDatabase[i].Name, deity.Name) {
				deitys = append(deitys, deitysDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(deitys))
	return
}

//EditDeity will grab data from storage
func (s *Storage) EditDeity(deity *model.Deity) (err error) {
	deityLock.Lock()
	defer deityLock.Unlock()
	for i := range deitysDatabase {
		if deitysDatabase[i].ID == deity.ID {
			*deitysDatabase[i] = *deity
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteDeity will grab data from storage
func (s *Storage) DeleteDeity(deity *model.Deity) (err error) {
	deityLock.Lock()
	defer deityLock.Unlock()
	indexToDelete := 0
	for i := range deitysDatabase {
		if deitysDatabase[i].ID == deity.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	deitysDatabase[len(deitysDatabase)-1], deitysDatabase[indexToDelete] = deitysDatabase[indexToDelete], deitysDatabase[len(deitysDatabase)-1]
	deitysDatabase = deitysDatabase[:len(deitysDatabase)-1]
	return
}
