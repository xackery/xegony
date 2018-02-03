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
	classLock = sync.RWMutex{}
)

//GetClass will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetClass(class *model.Class) (err error) {
	classLock.Lock()
	defer classLock.Unlock()
	classsDatabase, err := s.readClassFile()
	if err != nil {
		return
	}
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
	classsDatabase, err := s.readClassFile()
	if err != nil {
		return
	}
	for _, tmpClass := range classsDatabase {
		if tmpClass.ID == class.ID {
			err = fmt.Errorf("class already exists")
			return
		}
	}
	classsDatabase = append(classsDatabase, class)
	err = s.writeClassFile(classsDatabase)
	if err != nil {
		return
	}
	return
}

//ListClass will grab data from storage
func (s *Storage) ListClass(page *model.Page) (classs []*model.Class, err error) {
	classLock.Lock()
	defer classLock.Unlock()
	classsDatabase, err := s.readClassFile()
	if err != nil {
		return
	}

	classs = make([]*model.Class, len(classsDatabase))

	classs = classsDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
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

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(classs))
		}
	*/
	return
}

//ListClassTotalCount will grab data from storage
func (s *Storage) ListClassTotalCount() (count int64, err error) {
	classLock.Lock()
	defer classLock.Unlock()
	classsDatabase, err := s.readClassFile()
	if err != nil {
		return
	}
	count = int64(len(classsDatabase))
	return
}

//ListClassBySearch will grab data from storage
func (s *Storage) ListClassBySearch(page *model.Page, class *model.Class) (classs []*model.Class, err error) {
	classLock.Lock()
	defer classLock.Unlock()
	classsDatabase, err := s.readClassFile()
	if err != nil {
		return
	}
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
	classLock.Lock()
	defer classLock.Unlock()
	classsDatabase, err := s.readClassFile()
	if err != nil {
		return
	}

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
	classsDatabase, err := s.readClassFile()
	if err != nil {
		return
	}

	for i := range classsDatabase {
		if classsDatabase[i].ID == class.ID {
			*classsDatabase[i] = *class
			err = s.writeClassFile(classsDatabase)
			if err != nil {
				return
			}
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
	classsDatabase, err := s.readClassFile()
	if err != nil {
		return
	}
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
	err = s.writeClassFile(classsDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readClassFile() (classs []*model.Class, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			classs = loadClassDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeClassFile(classs)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default class data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &classs)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeClassFile(classs []*model.Class) (err error) {

	bData, err := yaml.Marshal(classs)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal classs")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
