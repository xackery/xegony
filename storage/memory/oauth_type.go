package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	oauthTypesDatabase = []*model.OauthType{}
	oauthTypeLock      = sync.RWMutex{}
)

//GetOauthType will grab data from storage
func (s *Storage) GetOauthType(oauthType *model.OauthType) (err error) {
	oauthTypeLock.RLock()
	defer oauthTypeLock.RUnlock()
	for _, tmpOauthType := range oauthTypesDatabase {
		if tmpOauthType.ID == oauthType.ID {
			*oauthType = *tmpOauthType
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateOauthType will grab data from storage
func (s *Storage) CreateOauthType(oauthType *model.OauthType) (err error) {
	oauthTypeLock.Lock()
	defer oauthTypeLock.Unlock()
	for _, tmpOauthType := range oauthTypesDatabase {
		if tmpOauthType.ID == oauthType.ID {
			err = fmt.Errorf("oauthType already exists")
			return
		}
	}
	oauthTypesDatabase = append(oauthTypesDatabase, oauthType)
	return
}

//ListOauthType will grab data from storage
func (s *Storage) ListOauthType(page *model.Page) (oauthTypes []*model.OauthType, err error) {
	oauthTypeLock.RLock()
	defer oauthTypeLock.RUnlock()

	oauthTypes = make([]*model.OauthType, len(oauthTypesDatabase))

	oauthTypes = oauthTypesDatabase

	switch page.OrderBy {
	case "id":
		sort.Slice(oauthTypes, func(i, j int) bool {
			return oauthTypes[i].ID < oauthTypes[j].ID
		})
	case "name":
		sort.Slice(oauthTypes, func(i, j int) bool {
			return oauthTypes[i].Name < oauthTypes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(oauthTypes))
		}
	*/
	return
}

//ListOauthTypeTotalCount will grab data from storage
func (s *Storage) ListOauthTypeTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(oauthTypesDatabase))
	return
}

//ListOauthTypeBySearch will grab data from storage
func (s *Storage) ListOauthTypeBySearch(page *model.Page, oauthType *model.OauthType) (oauthTypes []*model.OauthType, err error) {
	oauthTypeLock.RLock()
	defer oauthTypeLock.RUnlock()

	if len(oauthType.Name) > 0 {
		for i := range oauthTypesDatabase {
			if strings.Contains(oauthTypesDatabase[i].Name, oauthType.Name) {
				oauthTypes = append(oauthTypes, oauthTypesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "id":
		sort.Slice(oauthTypes, func(i, j int) bool {
			return oauthTypes[i].ID < oauthTypes[j].ID
		})
	case "name":
		sort.Slice(oauthTypes, func(i, j int) bool {
			return oauthTypes[i].Name < oauthTypes[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(oauthTypes))
	//}
	return
}

//ListOauthTypeBySearchTotalCount will grab data from storage
func (s *Storage) ListOauthTypeBySearchTotalCount(oauthType *model.OauthType) (count int64, err error) {
	oauthTypeLock.RLock()
	defer oauthTypeLock.RUnlock()

	oauthTypes := []*model.OauthType{}
	if len(oauthType.Name) > 0 {
		for i := range oauthTypesDatabase {
			if strings.Contains(oauthTypesDatabase[i].Name, oauthType.Name) {
				oauthTypes = append(oauthTypes, oauthTypesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(oauthTypes))
	return
}

//EditOauthType will grab data from storage
func (s *Storage) EditOauthType(oauthType *model.OauthType) (err error) {
	oauthTypeLock.Lock()
	defer oauthTypeLock.Unlock()
	for i := range oauthTypesDatabase {
		if oauthTypesDatabase[i].ID == oauthType.ID {
			*oauthTypesDatabase[i] = *oauthType
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteOauthType will grab data from storage
func (s *Storage) DeleteOauthType(oauthType *model.OauthType) (err error) {
	oauthTypeLock.Lock()
	defer oauthTypeLock.Unlock()
	indexToDelete := 0
	for i := range oauthTypesDatabase {
		if oauthTypesDatabase[i].ID == oauthType.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	oauthTypesDatabase[len(oauthTypesDatabase)-1], oauthTypesDatabase[indexToDelete] = oauthTypesDatabase[indexToDelete], oauthTypesDatabase[len(oauthTypesDatabase)-1]
	oauthTypesDatabase = oauthTypesDatabase[:len(oauthTypesDatabase)-1]
	return
}
