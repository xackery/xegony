package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetConfig will grab data from storage
func (s *Storage) GetConfig(config *model.Config) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateConfig will grab data from storage
func (s *Storage) CreateConfig(config *model.Config) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListConfig will grab data from storage
func (s *Storage) ListConfig(page *model.Page) (configs []*model.Config, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListConfigTotalCount will grab data from storage
func (s *Storage) ListConfigTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListConfigBySearch will grab data from storage
func (s *Storage) ListConfigBySearch(page *model.Page, config *model.Config) (configs []*model.Config, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListConfigBySearchTotalCount will grab data from storage
func (s *Storage) ListConfigBySearchTotalCount(config *model.Config) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditConfig will grab data from storage
func (s *Storage) EditConfig(config *model.Config) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteConfig will grab data from storage
func (s *Storage) DeleteConfig(config *model.Config) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
