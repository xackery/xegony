package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	sizesDatabase = []*model.Size{}
	sizeLock      = sync.RWMutex{}
)

//GetSize will grab data from storage
func (s *Storage) GetSize(size *model.Size) (err error) {
	sizeLock.RLock()
	defer sizeLock.RUnlock()
	for _, tmpSize := range sizesDatabase {
		if tmpSize.ID == size.ID {
			*size = *tmpSize
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSize will grab data from storage
func (s *Storage) CreateSize(size *model.Size) (err error) {
	sizeLock.Lock()
	defer sizeLock.Unlock()
	for _, tmpSize := range sizesDatabase {
		if tmpSize.ID == size.ID {
			err = fmt.Errorf("size already exists")
			return
		}
	}
	sizesDatabase = append(sizesDatabase, size)
	return
}

//ListSize will grab data from storage
func (s *Storage) ListSize(page *model.Page) (sizes []*model.Size, err error) {
	sizeLock.RLock()
	defer sizeLock.RUnlock()

	sizes = make([]*model.Size, len(sizesDatabase))

	sizes = sizesDatabase

	switch page.OrderBy {
	case "name":
		sort.Slice(sizes, func(i, j int) bool {
			return sizes[i].Name < sizes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(sizes))
		}
	*/
	return
}

//ListSizeTotalCount will grab data from storage
func (s *Storage) ListSizeTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(sizesDatabase))
	return
}

//ListSizeBySearch will grab data from storage
func (s *Storage) ListSizeBySearch(page *model.Page, size *model.Size) (sizes []*model.Size, err error) {
	sizeLock.RLock()
	defer sizeLock.RUnlock()

	if len(size.Name) > 0 {
		for i := range sizesDatabase {
			if strings.Contains(sizesDatabase[i].Name, size.Name) {
				sizes = append(sizes, sizesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(sizes, func(i, j int) bool {
			return sizes[i].Name < sizes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(sizes))
	//}
	return
}

//ListSizeBySearchTotalCount will grab data from storage
func (s *Storage) ListSizeBySearchTotalCount(size *model.Size) (count int64, err error) {
	sizeLock.RLock()
	defer sizeLock.RUnlock()

	sizes := []*model.Size{}
	if len(size.Name) > 0 {
		for i := range sizesDatabase {
			if strings.Contains(sizesDatabase[i].Name, size.Name) {
				sizes = append(sizes, sizesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(sizes))
	return
}

//EditSize will grab data from storage
func (s *Storage) EditSize(size *model.Size) (err error) {
	sizeLock.Lock()
	defer sizeLock.Unlock()
	for i := range sizesDatabase {
		if sizesDatabase[i].ID == size.ID {
			*sizesDatabase[i] = *size
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSize will grab data from storage
func (s *Storage) DeleteSize(size *model.Size) (err error) {
	sizeLock.Lock()
	defer sizeLock.Unlock()
	indexToDelete := 0
	for i := range sizesDatabase {
		if sizesDatabase[i].ID == size.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	sizesDatabase[len(sizesDatabase)-1], sizesDatabase[indexToDelete] = sizesDatabase[indexToDelete], sizesDatabase[len(sizesDatabase)-1]
	sizesDatabase = sizesDatabase[:len(sizesDatabase)-1]
	return
}
