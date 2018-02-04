package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	classsDatabase = []*model.Class{}
	classLock      = sync.RWMutex{}
)

//GetClass will grab data from storage
func (s *Storage) GetClass(class *model.Class) (err error) {
	classLock.RLock()
	defer classLock.RUnlock()
	for _, tmpClass := range classsDatabase {
		if tmpClass.ID == class.ID {
			*class = *tmpClass
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateClass will grab data from storage
func (s *Storage) CreateClass(class *model.Class) (err error) {
	classLock.Lock()
	defer classLock.Unlock()
	for _, tmpClass := range classsDatabase {
		if tmpClass.ID == class.ID {
			err = fmt.Errorf("class already exists")
			return
		}
	}
	classsDatabase = append(classsDatabase, class)
	return
}

//ListClass will grab data from storage
func (s *Storage) ListClass(page *model.Page) (classs []*model.Class, err error) {
	classLock.RLock()
	defer classLock.RUnlock()

	classs = make([]*model.Class, len(classsDatabase))

	classs = classsDatabase

	switch page.OrderBy {
	case "name":
		sort.Slice(classs, func(i, j int) bool {
			return classs[i].Name < classs[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(classs))
		}
	*/
	return
}

//ListClassTotalCount will grab data from storage
func (s *Storage) ListClassTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(classsDatabase))
	return
}

//ListClassByBit will grab data from storage
func (s *Storage) ListClassByBit(page *model.Page, class *model.Class) (classs []*model.Class, err error) {
	classLock.RLock()
	defer classLock.RUnlock()

	for i := range classsDatabase {
		if class.Bit < 1 || classsDatabase[i].Bit < 1 {
			continue
		}
		if class.Bit&classsDatabase[i].Bit == classsDatabase[i].Bit {
			classs = append(classs, classsDatabase[i])
		}
	}
	switch page.OrderBy {
	case "name":
		sort.Slice(classs, func(i, j int) bool {
			return classs[i].Name < classs[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(classs))
	//}
	return
}

//ListClassByBitTotalCount will grab data from storage
func (s *Storage) ListClassByBitTotalCount(class *model.Class) (count int64, err error) {
	classLock.RLock()
	defer classLock.RUnlock()

	classs := []*model.Class{}
	for i := range classsDatabase {
		if class.Bit < 1 || classsDatabase[i].Bit < 1 {
			continue
		}
		if class.Bit&classsDatabase[i].Bit == classsDatabase[i].Bit {
			classs = append(classs, classsDatabase[i])
		}
	}
	count = int64(len(classs))
	return
}

//ListClassBySearch will grab data from storage
func (s *Storage) ListClassBySearch(page *model.Page, class *model.Class) (classs []*model.Class, err error) {
	classLock.RLock()
	defer classLock.RUnlock()

	if len(class.Name) > 0 {
		for i := range classsDatabase {
			if strings.Contains(classsDatabase[i].Name, class.Name) {
				classs = append(classs, classsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(classs, func(i, j int) bool {
			return classs[i].Name < classs[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(classs))
	//}
	return
}

//ListClassBySearchTotalCount will grab data from storage
func (s *Storage) ListClassBySearchTotalCount(class *model.Class) (count int64, err error) {
	classLock.RLock()
	defer classLock.RUnlock()

	classs := []*model.Class{}
	if len(class.Name) > 0 {
		for i := range classsDatabase {
			if strings.Contains(classsDatabase[i].Name, class.Name) {
				classs = append(classs, classsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(classs))
	return
}

//EditClass will grab data from storage
func (s *Storage) EditClass(class *model.Class) (err error) {
	classLock.Lock()
	defer classLock.Unlock()
	for i := range classsDatabase {
		if classsDatabase[i].ID == class.ID {
			*classsDatabase[i] = *class
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteClass will grab data from storage
func (s *Storage) DeleteClass(class *model.Class) (err error) {
	classLock.Lock()
	defer classLock.Unlock()
	indexToDelete := 0
	for i := range classsDatabase {
		if classsDatabase[i].ID == class.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	classsDatabase[len(classsDatabase)-1], classsDatabase[indexToDelete] = classsDatabase[indexToDelete], classsDatabase[len(classsDatabase)-1]
	classsDatabase = classsDatabase[:len(classsDatabase)-1]
	return
}
