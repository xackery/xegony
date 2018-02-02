package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	configsDatabase = []*model.Config{}
	configLock      = sync.RWMutex{}
)

//GetConfig will grab data from storage
func (s *Storage) GetConfig(config *model.Config) (err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	for _, tmpConfig := range configsDatabase {
		if tmpConfig.Key == config.Key {
			*config = *tmpConfig
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateConfig will grab data from storage
func (s *Storage) CreateConfig(config *model.Config) (err error) {
	configLock.Lock()
	defer configLock.Unlock()
	for _, tmpConfig := range configsDatabase {
		if tmpConfig.Key == config.Key {
			err = fmt.Errorf("config already exists")
			return
		}
	}
	configsDatabase = append(configsDatabase, config)
	return
}

//ListConfig will grab data from storage
func (s *Storage) ListConfig(page *model.Page) (configs []*model.Config, err error) {
	configLock.RLock()
	defer configLock.RUnlock()

	configs = make([]*model.Config, len(configsDatabase))

	configs = configsDatabase

	switch page.OrderBy {
	case "key":
		sort.Slice(configs, func(i, j int) bool {
			return configs[i].Key < configs[j].Key
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(configs))
		}
	*/
	return
}

//ListConfigTotalCount will grab data from storage
func (s *Storage) ListConfigTotalCount() (count int64, err error) {
	count = int64(len(configsDatabase))
	return
}

//ListConfigBySearch will grab data from storage
func (s *Storage) ListConfigBySearch(page *model.Page, config *model.Config) (configs []*model.Config, err error) {
	configLock.RLock()
	defer configLock.RUnlock()

	if len(config.Key) > 0 {
		for i := range configsDatabase {
			if strings.Contains(configsDatabase[i].Key, config.Key) {
				configs = append(configs, configsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "key":
		sort.Slice(configs, func(i, j int) bool {
			return configs[i].Key < configs[j].Key
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(configs))
	//}
	return
}

//ListConfigBySearchTotalCount will grab data from storage
func (s *Storage) ListConfigBySearchTotalCount(config *model.Config) (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()

	configs := []*model.Config{}
	if len(config.Key) > 0 {
		for i := range configsDatabase {
			if strings.Contains(configsDatabase[i].Key, config.Key) {
				configs = append(configs, configsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(configs))
	return
}

//EditConfig will grab data from storage
func (s *Storage) EditConfig(config *model.Config) (err error) {
	configLock.Lock()
	defer configLock.Unlock()
	for i := range configsDatabase {
		if configsDatabase[i].Key == config.Key {
			*configsDatabase[i] = *config
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteConfig will grab data from storage
func (s *Storage) DeleteConfig(config *model.Config) (err error) {
	configLock.Lock()
	defer configLock.Unlock()
	indexToDelete := 0
	for i := range configsDatabase {
		if configsDatabase[i].Key == config.Key {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	configsDatabase[len(configsDatabase)-1], configsDatabase[indexToDelete] = configsDatabase[indexToDelete], configsDatabase[len(configsDatabase)-1]
	configsDatabase = configsDatabase[:len(configsDatabase)-1]
	return
}
