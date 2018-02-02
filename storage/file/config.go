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
	configLock = sync.RWMutex{}
)

//GetConfig will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetConfig(config *model.Config) (err error) {
	configLock.Lock()
	defer configLock.Unlock()
	configsDatabase, err := s.readConfigFile()
	if err != nil {
		return
	}
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
	configsDatabase, err := s.readConfigFile()
	if err != nil {
		return
	}
	for _, tmpConfig := range configsDatabase {
		if tmpConfig.Key == config.Key {
			err = fmt.Errorf("config already exists")
			return
		}
	}
	configsDatabase = append(configsDatabase, config)
	err = s.writeConfigFile(configsDatabase)
	if err != nil {
		return
	}
	return
}

//ListConfig will grab data from storage
func (s *Storage) ListConfig(page *model.Page) (configs []*model.Config, err error) {
	configLock.Lock()
	defer configLock.Unlock()
	configsDatabase, err := s.readConfigFile()
	if err != nil {
		return
	}

	configs = make([]*model.Config, len(configsDatabase))

	configs = configsDatabase

	if page.OrderBy == "" {
		page.OrderBy = "key"
	}

	switch page.OrderBy {
	case "key":
		sort.Slice(configs, func(i, j int) bool {
			return configs[i].Key < configs[j].Key
		})
	default:
		err = fmt.Errorf("Unsupported sort key")
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
	configLock.Lock()
	defer configLock.Unlock()
	configsDatabase, err := s.readConfigFile()
	if err != nil {
		return
	}
	count = int64(len(configsDatabase))
	return
}

//ListConfigBySearch will grab data from storage
func (s *Storage) ListConfigBySearch(page *model.Page, config *model.Config) (configs []*model.Config, err error) {
	configLock.Lock()
	defer configLock.Unlock()
	configsDatabase, err := s.readConfigFile()
	if err != nil {
		return
	}
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
		err = fmt.Errorf("Unsupported sort key")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(configs))
	//}
	return
}

//ListConfigBySearchTotalCount will grab data from storage
func (s *Storage) ListConfigBySearchTotalCount(config *model.Config) (count int64, err error) {
	configLock.Lock()
	defer configLock.Unlock()
	configsDatabase, err := s.readConfigFile()
	if err != nil {
		return
	}

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
	configsDatabase, err := s.readConfigFile()
	if err != nil {
		return
	}

	for i := range configsDatabase {
		if configsDatabase[i].Key == config.Key {
			*configsDatabase[i] = *config
			err = s.writeConfigFile(configsDatabase)
			if err != nil {
				return
			}
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
	configsDatabase, err := s.readConfigFile()
	if err != nil {
		return
	}
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
	err = s.writeConfigFile(configsDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readConfigFile() (configs []*model.Config, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			configs = loadConfigDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeConfigFile(configs)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default config data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &configs)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeConfigFile(configs []*model.Config) (err error) {

	bData, err := yaml.Marshal(configs)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal configs")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
