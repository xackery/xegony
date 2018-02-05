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
	deityLock = sync.RWMutex{}
)

//GetDeity will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetDeity(deity *model.Deity) (err error) {
	deityLock.Lock()
	defer deityLock.Unlock()
	deitysDatabase, err := s.readDeityFile()
	if err != nil {
		return
	}
	for _, tmpDeity := range deitysDatabase {
		if tmpDeity.ID == deity.ID {
			*deity = *tmpDeity
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//GetDeityBySpell will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetDeityBySpell(spell *model.Spell, deity *model.Deity) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateDeity will grab data from storage
func (s *Storage) CreateDeity(deity *model.Deity) (err error) {
	deityLock.Lock()
	defer deityLock.Unlock()
	deitysDatabase, err := s.readDeityFile()
	if err != nil {
		return
	}
	for _, tmpDeity := range deitysDatabase {
		if tmpDeity.ID == deity.ID {
			err = fmt.Errorf("deity already exists")
			return
		}
	}
	deitysDatabase = append(deitysDatabase, deity)
	err = s.writeDeityFile(deitysDatabase)
	if err != nil {
		return
	}
	return
}

//ListDeity will grab data from storage
func (s *Storage) ListDeity(page *model.Page) (deitys []*model.Deity, err error) {
	deityLock.Lock()
	defer deityLock.Unlock()
	deitysDatabase, err := s.readDeityFile()
	if err != nil {
		return
	}

	deitys = make([]*model.Deity, len(deitysDatabase))

	deitys = deitysDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
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

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(deitys))
		}
	*/
	return
}

//ListDeityTotalCount will grab data from storage
func (s *Storage) ListDeityTotalCount() (count int64, err error) {
	deityLock.Lock()
	defer deityLock.Unlock()
	deitysDatabase, err := s.readDeityFile()
	if err != nil {
		return
	}
	count = int64(len(deitysDatabase))
	return
}

//ListDeityByBit will grab data from storage
func (s *Storage) ListDeityByBit(page *model.Page, deity *model.Deity) (deitys []*model.Deity, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListDeityByBitTotalCount will grab data from storage
func (s *Storage) ListDeityByBitTotalCount(deity *model.Deity) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListDeityBySearch will grab data from storage
func (s *Storage) ListDeityBySearch(page *model.Page, deity *model.Deity) (deitys []*model.Deity, err error) {
	deityLock.Lock()
	defer deityLock.Unlock()
	deitysDatabase, err := s.readDeityFile()
	if err != nil {
		return
	}
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
	deityLock.Lock()
	defer deityLock.Unlock()
	deitysDatabase, err := s.readDeityFile()
	if err != nil {
		return
	}

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
	deitysDatabase, err := s.readDeityFile()
	if err != nil {
		return
	}

	for i := range deitysDatabase {
		if deitysDatabase[i].ID == deity.ID {
			*deitysDatabase[i] = *deity
			err = s.writeDeityFile(deitysDatabase)
			if err != nil {
				return
			}
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
	deitysDatabase, err := s.readDeityFile()
	if err != nil {
		return
	}
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
	err = s.writeDeityFile(deitysDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readDeityFile() (deitys []*model.Deity, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			deitys = loadDeityDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeDeityFile(deitys)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default deity data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &deitys)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeDeityFile(deitys []*model.Deity) (err error) {

	bData, err := yaml.Marshal(deitys)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal deitys")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
