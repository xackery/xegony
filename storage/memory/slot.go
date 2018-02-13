package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	slotsDatabase = []*model.Slot{}
	slotLock      = sync.RWMutex{}
)

//GetSlot will grab data from storage
func (s *Storage) GetSlot(slot *model.Slot) (err error) {
	slotLock.RLock()
	defer slotLock.RUnlock()
	for _, tmpSlot := range slotsDatabase {
		if tmpSlot.ID == slot.ID {
			*slot = *tmpSlot
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSlot will grab data from storage
func (s *Storage) CreateSlot(slot *model.Slot) (err error) {
	slotLock.Lock()
	defer slotLock.Unlock()
	for _, tmpSlot := range slotsDatabase {
		if tmpSlot.ID == slot.ID {
			err = fmt.Errorf("slot already exists")
			return
		}
	}
	slotsDatabase = append(slotsDatabase, slot)
	return
}

//ListSlot will grab data from storage
func (s *Storage) ListSlot(page *model.Page) (slots []*model.Slot, err error) {
	slotLock.RLock()
	defer slotLock.RUnlock()

	slots = make([]*model.Slot, len(slotsDatabase))

	slots = slotsDatabase

	switch page.OrderBy {
	case "name":
		sort.Slice(slots, func(i, j int) bool {
			return slots[i].Name < slots[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(slots))
		}
	*/
	return
}

//ListSlotTotalCount will grab data from storage
func (s *Storage) ListSlotTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(slotsDatabase))
	return
}

//ListSlotByBit will grab data from storage
func (s *Storage) ListSlotByBit(page *model.Page, slot *model.Slot) (slots []*model.Slot, err error) {
	slotLock.RLock()
	defer slotLock.RUnlock()

	for i := range slotsDatabase {
		if slot.Bit < 1 || slotsDatabase[i].Bit < 1 {
			continue
		}
		if slot.Bit&slotsDatabase[i].Bit == slotsDatabase[i].Bit {
			slots = append(slots, slotsDatabase[i])
		}
	}
	switch page.OrderBy {
	case "name":
		sort.Slice(slots, func(i, j int) bool {
			return slots[i].Name < slots[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(slots))
	//}
	return
}

//ListSlotByBitTotalCount will grab data from storage
func (s *Storage) ListSlotByBitTotalCount(slot *model.Slot) (count int64, err error) {
	slotLock.RLock()
	defer slotLock.RUnlock()

	slots := []*model.Slot{}
	for i := range slotsDatabase {
		if slot.Bit < 1 || slotsDatabase[i].Bit < 1 {
			continue
		}
		if slot.Bit&slotsDatabase[i].Bit == slotsDatabase[i].Bit {
			slots = append(slots, slotsDatabase[i])
		}
	}
	count = int64(len(slots))
	return
}

//ListSlotBySearch will grab data from storage
func (s *Storage) ListSlotBySearch(page *model.Page, slot *model.Slot) (slots []*model.Slot, err error) {
	slotLock.RLock()
	defer slotLock.RUnlock()

	if len(slot.Name) > 0 {
		for i := range slotsDatabase {
			if strings.Contains(slotsDatabase[i].Name, slot.Name) {
				slots = append(slots, slotsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(slots, func(i, j int) bool {
			return slots[i].Name < slots[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(slots))
	//}
	return
}

//ListSlotBySearchTotalCount will grab data from storage
func (s *Storage) ListSlotBySearchTotalCount(slot *model.Slot) (count int64, err error) {
	slotLock.RLock()
	defer slotLock.RUnlock()

	slots := []*model.Slot{}
	if len(slot.Name) > 0 {
		for i := range slotsDatabase {
			if strings.Contains(slotsDatabase[i].Name, slot.Name) {
				slots = append(slots, slotsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(slots))
	return
}

//EditSlot will grab data from storage
func (s *Storage) EditSlot(slot *model.Slot) (err error) {
	slotLock.Lock()
	defer slotLock.Unlock()
	for i := range slotsDatabase {
		if slotsDatabase[i].ID == slot.ID {
			*slotsDatabase[i] = *slot
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSlot will grab data from storage
func (s *Storage) DeleteSlot(slot *model.Slot) (err error) {
	slotLock.Lock()
	defer slotLock.Unlock()
	indexToDelete := 0
	for i := range slotsDatabase {
		if slotsDatabase[i].ID == slot.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	slotsDatabase[len(slotsDatabase)-1], slotsDatabase[indexToDelete] = slotsDatabase[indexToDelete], slotsDatabase[len(slotsDatabase)-1]
	slotsDatabase = slotsDatabase[:len(slotsDatabase)-1]
	return
}
