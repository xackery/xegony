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
	sizeLock = sync.RWMutex{}
)

//GetSize will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetSize(size *model.Size) (err error) {
	sizeLock.Lock()
	defer sizeLock.Unlock()
	sizesDatabase, err := s.readSizeFile()
	if err != nil {
		return
	}
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
	sizesDatabase, err := s.readSizeFile()
	if err != nil {
		return
	}
	for _, tmpSize := range sizesDatabase {
		if tmpSize.ID == size.ID {
			err = fmt.Errorf("size already exists")
			return
		}
	}
	sizesDatabase = append(sizesDatabase, size)
	err = s.writeSizeFile(sizesDatabase)
	if err != nil {
		return
	}
	return
}

//ListSize will grab data from storage
func (s *Storage) ListSize(page *model.Page) (sizes []*model.Size, err error) {
	sizeLock.Lock()
	defer sizeLock.Unlock()
	sizesDatabase, err := s.readSizeFile()
	if err != nil {
		return
	}

	sizes = make([]*model.Size, len(sizesDatabase))

	sizes = sizesDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
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

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(sizes))
		}
	*/
	return
}

//ListSizeTotalCount will grab data from storage
func (s *Storage) ListSizeTotalCount() (count int64, err error) {
	sizeLock.Lock()
	defer sizeLock.Unlock()
	sizesDatabase, err := s.readSizeFile()
	if err != nil {
		return
	}
	count = int64(len(sizesDatabase))
	return
}

//ListSizeByBit will grab data from storage
func (s *Storage) ListSizeByBit(page *model.Page, size *model.Size) (sizes []*model.Size, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSizeByBitTotalCount will grab data from storage
func (s *Storage) ListSizeByBitTotalCount(size *model.Size) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSizeBySearch will grab data from storage
func (s *Storage) ListSizeBySearch(page *model.Page, size *model.Size) (sizes []*model.Size, err error) {
	sizeLock.Lock()
	defer sizeLock.Unlock()
	sizesDatabase, err := s.readSizeFile()
	if err != nil {
		return
	}
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
	sizeLock.Lock()
	defer sizeLock.Unlock()
	sizesDatabase, err := s.readSizeFile()
	if err != nil {
		return
	}

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
	sizesDatabase, err := s.readSizeFile()
	if err != nil {
		return
	}

	for i := range sizesDatabase {
		if sizesDatabase[i].ID == size.ID {
			*sizesDatabase[i] = *size
			err = s.writeSizeFile(sizesDatabase)
			if err != nil {
				return
			}
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
	sizesDatabase, err := s.readSizeFile()
	if err != nil {
		return
	}
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
	err = s.writeSizeFile(sizesDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readSizeFile() (sizes []*model.Size, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			sizes = loadSizeDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeSizeFile(sizes)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default size data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &sizes)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeSizeFile(sizes []*model.Size) (err error) {

	bData, err := yaml.Marshal(sizes)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal sizes")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
