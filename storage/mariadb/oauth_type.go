package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetOauthType will grab data from storage
func (s *Storage) GetOauthType(oauthType *model.OauthType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateOauthType will grab data from storage
func (s *Storage) CreateOauthType(oauthType *model.OauthType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListOauthType will grab data from storage
func (s *Storage) ListOauthType(page *model.Page) (oauthTypes []*model.OauthType, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListOauthTypeTotalCount will grab data from storage
func (s *Storage) ListOauthTypeTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListOauthTypeBySearch will grab data from storage
func (s *Storage) ListOauthTypeBySearch(page *model.Page, oauthType *model.OauthType) (oauthTypes []*model.OauthType, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListOauthTypeBySearchTotalCount will grab data from storage
func (s *Storage) ListOauthTypeBySearchTotalCount(oauthType *model.OauthType) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditOauthType will grab data from storage
func (s *Storage) EditOauthType(oauthType *model.OauthType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteOauthType will grab data from storage
func (s *Storage) DeleteOauthType(oauthType *model.OauthType) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
