package memory

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetUserAccount will grab data from storage
func (s *Storage) GetUserAccount(user *model.User, userAccount *model.UserAccount) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateUserAccount will grab data from storage
func (s *Storage) CreateUserAccount(user *model.User, userAccount *model.UserAccount) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListUserAccount will grab data from storage
func (s *Storage) ListUserAccount(page *model.Page, user *model.User) (userAccounts []*model.UserAccount, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListUserAccountTotalCount will grab data from storage
func (s *Storage) ListUserAccountTotalCount(user *model.User) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListUserAccountBySearch will grab data from storage
func (s *Storage) ListUserAccountBySearch(page *model.Page, user *model.User, userAccount *model.UserAccount) (userAccounts []*model.UserAccount, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListUserAccountBySearchTotalCount will grab data from storage
func (s *Storage) ListUserAccountBySearchTotalCount(user *model.User, userAccount *model.UserAccount) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditUserAccount will grab data from storage
func (s *Storage) EditUserAccount(user *model.User, userAccount *model.UserAccount) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteUserAccount will grab data from storage
func (s *Storage) DeleteUserAccount(user *model.User, userAccount *model.UserAccount) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
