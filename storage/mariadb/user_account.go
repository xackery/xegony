package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	userAccountTable  = "user_account"
	userAccountFields = "id,user_id,account_id,character_id,create_date"
	userAccountBinds  = ":id,:user_id,:account_id,:character_id,:create_date"
)

//GetUserAccount will grab data from storage
func (s *Storage) GetUserAccount(user *model.User, userAccount *model.UserAccount) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = ? AND account_id = ?", userAccountFields, userAccountTable)
	err = s.db.Get(userAccount, query, user.ID, userAccount.AccountID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateUserAccount will grab data from storage
func (s *Storage) CreateUserAccount(user *model.User, userAccount *model.UserAccount) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", userAccountTable, userAccountFields, userAccountBinds)
	_, err = s.db.NamedExec(query, userAccount)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListUserAccount will grab data from storage
func (s *Storage) ListUserAccount(page *model.Page, user *model.User) (userAccounts []*model.UserAccount, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "user_id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = ? ORDER BY %s LIMIT %d OFFSET %d", userAccountFields, userAccountTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query, user.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		userAccount := model.UserAccount{}
		if err = rows.StructScan(&userAccount); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		userAccounts = append(userAccounts, &userAccount)
	}
	return
}

//ListUserAccountTotalCount will grab data from storage
func (s *Storage) ListUserAccountTotalCount(user *model.User) (count int64, err error) {
	query := fmt.Sprintf("SELECT count(user_id) FROM %s WHERE user_id = ?", userAccountTable)
	err = s.db.Get(&count, query, user.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListUserAccountBySearch will grab data from storage
func (s *Storage) ListUserAccountBySearch(page *model.Page, user *model.User, userAccount *model.UserAccount) (userAccounts []*model.UserAccount, err error) {

	field := ""

	if userAccount.AccountID > 0 {
		field += `user_id = :user_id OR`
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]
	userAccount.UserID = user.ID

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", userAccountFields, userAccountTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, userAccount)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		userAccount := model.UserAccount{}
		if err = rows.StructScan(&userAccount); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		userAccounts = append(userAccounts, &userAccount)
	}
	return
}

//ListUserAccountBySearchTotalCount will grab data from storage
func (s *Storage) ListUserAccountBySearchTotalCount(user *model.User, userAccount *model.UserAccount) (count int64, err error) {
	field := ""
	if userAccount.AccountID > 0 {
		field += `user_id = :user_id OR`
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	userAccount.UserID = user.ID
	query := fmt.Sprintf("SELECT count(usergroupID) FROM %s WHERE %s", userAccountTable, field)

	rows, err := s.db.NamedQuery(query, userAccount)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
	}
	return
}

//EditUserAccount will grab data from storage
func (s *Storage) EditUserAccount(user *model.User, userAccount *model.UserAccount) (err error) {

	prevUserAccount := &model.UserAccount{
		UserID:    userAccount.UserID,
		AccountID: userAccount.AccountID,
	}
	err = s.GetUserAccount(user, prevUserAccount)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous userAccount")
		return
	}

	field := ""
	if prevUserAccount.CharacterID != userAccount.CharacterID {
		field += "character_id = :character_id, "
	}

	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id = :user_id AND account_id = :account_id", userAccountTable, field)
	result, err := s.db.NamedExec(query, userAccount)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//DeleteUserAccount will grab data from storage
func (s *Storage) DeleteUserAccount(user *model.User, userAccount *model.UserAccount) (err error) {

	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = ? AND account_id = ?", userAccountTable)
	result, err := s.db.Exec(query, user.ID, userAccount.AccountID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//createTableUserAccount will grab data from storage
func (s *Storage) createTableUserAccount() (err error) {
	_, err = s.db.Exec(`
		CREATE TABLE user_account (
			id int(11) NOT NULL AUTO_INCREMENT,
			user_id int(11) unsigned NOT NULL DEFAULT '0',
			account_id int(11) unsigned NOT NULL DEFAULT '0',
			character_id int(11) unsigned NOT NULL DEFAULT '0',
			create_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		  ) ENGINE=INNODB DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
