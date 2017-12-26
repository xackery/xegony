package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

func (s *Storage) GetAccount(accountId int64) (account *model.Account, err error) {
	account = &model.Account{}
	err = s.db.Get(account, "SELECT id, name, status FROM account WHERE id = ?", accountId)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateAccount(account *model.Account) (err error) {
	if account == nil {
		err = fmt.Errorf("Must provide account")
		return
	}

	result, err := s.db.NamedExec(`INSERT INTO account(name, status)
		VALUES (:name, :status)`, account)
	if err != nil {
		return
	}
	accountId, err := result.LastInsertId()
	if err != nil {
		return
	}
	account.Id = accountId
	return
}

func (s *Storage) ListAccount() (accounts []*model.Account, err error) {
	rows, err := s.db.Queryx(`SELECT id, name, status FROM account ORDER BY id DESC`)
	if err != nil {
		return
	}

	for rows.Next() {
		account := model.Account{}
		if err = rows.StructScan(&account); err != nil {
			return
		}
		accounts = append(accounts, &account)
	}
	return
}

func (s *Storage) EditAccount(accountId int64, account *model.Account) (err error) {
	account.Id = accountId
	result, err := s.db.NamedExec(`UPDATE account SET name=:name, status=:status WHERE id = :id`, account)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}

func (s *Storage) DeleteAccount(accountId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM account WHERE id = ?`, accountId)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}
