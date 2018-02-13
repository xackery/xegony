package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

var (
	slotLock = sync.RWMutex{}
)

//GetSlot will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetSlot(slot *model.Slot) (err error) {
	slotLock.Lock()
	defer slotLock.Unlock()
	slotsDatabase, err := s.readSlotFile()
	if err != nil {
		return
	}
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
	slotsDatabase, err := s.readSlotFile()
	if err != nil {
		return
	}
	for _, tmpSlot := range slotsDatabase {
		if tmpSlot.ID == slot.ID {
			err = fmt.Errorf("slot already exists")
			return
		}
	}
	slotsDatabase = append(slotsDatabase, slot)
	err = s.writeSlotFile(slotsDatabase)
	if err != nil {
		return
	}
	return
}

//ListSlot will grab data from storage
func (s *Storage) ListSlot(page *model.Page) (slots []*model.Slot, err error) {
	slotLock.Lock()
	defer slotLock.Unlock()
	slotsDatabase, err := s.readSlotFile()
	if err != nil {
		return
	}

	slots = make([]*model.Slot, len(slotsDatabase))

	slots = slotsDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
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

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(slots))
		}
	*/
	return
}

//ListSlotTotalCount will grab data from storage
func (s *Storage) ListSlotTotalCount() (count int64, err error) {
	slotLock.Lock()
	defer slotLock.Unlock()
	slotsDatabase, err := s.readSlotFile()
	if err != nil {
		return
	}
	count = int64(len(slotsDatabase))
	return
}

//ListSlotByBit will grab data from storage
func (s *Storage) ListSlotByBit(page *model.Page, slot *model.Slot) (slots []*model.Slot, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSlotByBitTotalCount will grab data from storage
func (s *Storage) ListSlotByBitTotalCount(slot *model.Slot) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSlotBySearch will grab data from storage
func (s *Storage) ListSlotBySearch(page *model.Page, slot *model.Slot) (slots []*model.Slot, err error) {
	slotLock.Lock()
	defer slotLock.Unlock()
	slotsDatabase, err := s.readSlotFile()
	if err != nil {
		return
	}
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
	slotLock.Lock()
	defer slotLock.Unlock()
	slotsDatabase, err := s.readSlotFile()
	if err != nil {
		return
	}

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
	slotsDatabase, err := s.readSlotFile()
	if err != nil {
		return
	}

	for i := range slotsDatabase {
		if slotsDatabase[i].ID == slot.ID {
			*slotsDatabase[i] = *slot
			err = s.writeSlotFile(slotsDatabase)
			if err != nil {
				return
			}
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
	slotsDatabase, err := s.readSlotFile()
	if err != nil {
		return
	}
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
	err = s.writeSlotFile(slotsDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readSlotFile() (slots []*model.Slot, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			slots = loadSlotDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeSlotFile(slots)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default slot data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &slots)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeSlotFile(slots []*model.Slot) (err error) {

	bData, err := yaml.Marshal(slots)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal slots")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
