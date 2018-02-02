package file

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetUser will grab data from storage
func (s *Storage) GetUser(user *model.User) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateUser will grab data from storage
func (s *Storage) CreateUser(user *model.User) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListUser will grab data from storage
func (s *Storage) ListUser(page *model.Page) (users []*model.User, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListUserTotalCount will grab data from storage
func (s *Storage) ListUserTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListUserBySearch will grab data from storage
func (s *Storage) ListUserBySearch(page *model.Page, user *model.User) (users []*model.User, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListUserBySearchTotalCount will grab data from storage
func (s *Storage) ListUserBySearchTotalCount(user *model.User) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditUser will grab data from storage
func (s *Storage) EditUser(user *model.User) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteUser will grab data from storage
func (s *Storage) DeleteUser(user *model.User) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
