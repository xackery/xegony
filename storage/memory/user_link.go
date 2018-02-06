package memory

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetUserLink will grab data from storage
func (s *Storage) GetUserLink(userLink *model.UserLink) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateUserLink will grab data from storage
func (s *Storage) CreateUserLink(userLink *model.UserLink) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListUserLink will grab data from storage
func (s *Storage) ListUserLink(page *model.Page) (userLinks []*model.UserLink, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListUserLinkTotalCount will grab data from storage
func (s *Storage) ListUserLinkTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListUserLinkBySearch will grab data from storage
func (s *Storage) ListUserLinkBySearch(page *model.Page, userLink *model.UserLink) (userLinks []*model.UserLink, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListUserLinkBySearchTotalCount will grab data from storage
func (s *Storage) ListUserLinkBySearchTotalCount(userLink *model.UserLink) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditUserLink will grab data from storage
func (s *Storage) EditUserLink(userLink *model.UserLink) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteUserLink will grab data from storage
func (s *Storage) DeleteUserLink(userLink *model.UserLink) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteUserLinkByAccount will grab data from storage
func (s *Storage) DeleteUserLinkByAccount(account *model.Account) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
