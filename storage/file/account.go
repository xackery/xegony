package file

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetAccount will grab data from storage
func (s *Storage) GetAccount(account *model.Account) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateAccount will grab data from storage
func (s *Storage) CreateAccount(account *model.Account) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListAccount will grab data from storage
func (s *Storage) ListAccount(page *model.Page) (accounts []*model.Account, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListAccountTotalCount will grab data from storage
func (s *Storage) ListAccountTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListAccountBySearch will grab data from storage
func (s *Storage) ListAccountBySearch(page *model.Page, account *model.Account) (accounts []*model.Account, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListAccountBySearchTotalCount will grab data from storage
func (s *Storage) ListAccountBySearchTotalCount(account *model.Account) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditAccount will grab data from storage
func (s *Storage) EditAccount(account *model.Account) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteAccount will grab data from storage
func (s *Storage) DeleteAccount(account *model.Account) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
