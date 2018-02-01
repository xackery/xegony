package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	userTable  = "user"
	userFields = "display_name, primary_account_id, primary_character_id, email, password, google_token"
	userBinds  = ":display_name, :primary_account_id, :primary_character_id, :email, :password, :google_token,"
)

//GetUser will grab data from storage
func (s *Storage) GetUser(user *model.User) (err error) {
	query := fmt.Sprintf("SELECT id, %s FROM %s WHERE id = ?", userFields, userTable)
	err = s.db.Get(user, query, user.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateUser will grab data from storage
func (s *Storage) CreateUser(user *model.User) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", userTable, userFields, userBinds)
	result, err := s.db.NamedExec(query, user)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	userID, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	user.ID = userID
	return
}

//ListUser will grab data from storage
func (s *Storage) ListUser(page *model.Page) (users []*model.User, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT id, %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", userFields, userTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		user := model.User{}
		if err = rows.StructScan(&user); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		users = append(users, &user)
	}
	return
}

//ListUserTotalCount will grab data from storage
func (s *Storage) ListUserTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(id) FROM %s", userTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListUserBySearch will grab data from storage
func (s *Storage) ListUserBySearch(page *model.Page, user *model.User) (users []*model.User, err error) {

	field := ""

	if len(user.DisplayName) > 0 {
		field += `display_name LIKE :display_name OR`
		user.DisplayName = fmt.Sprintf("%%%s%%", user.DisplayName)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT id, %s FROM %s WHERE %s LIMIT %d OFFSET %d", userFields, userTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, user)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		user := model.User{}
		if err = rows.StructScan(&user); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		users = append(users, &user)
	}
	return
}

//ListUserBySearchTotalCount will grab data from storage
func (s *Storage) ListUserBySearchTotalCount(user *model.User) (count int64, err error) {
	field := ""
	if len(user.DisplayName) > 0 {
		field += `display_name LIKE :display_name OR`
		user.DisplayName = fmt.Sprintf("%%%s%%", user.DisplayName)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", userTable, field)

	rows, err := s.db.NamedQuery(query, user)
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

//EditUser will grab data from storage
func (s *Storage) EditUser(user *model.User) (err error) {

	prevUser := &model.User{
		ID: user.ID,
	}
	err = s.GetUser(prevUser)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous user")
		return
	}

	field := ""

	if len(user.DisplayName) > 0 && prevUser.DisplayName != user.DisplayName {
		field += "display_name = :display_name, "
	}

	if user.PrimaryAccountID > 0 && prevUser.PrimaryAccountID != user.PrimaryAccountID {
		field += "primary_account_id = :primary_account_id, "
	}

	if user.PrimaryCharacterID > 0 && prevUser.PrimaryCharacterID != user.PrimaryCharacterID {
		field += "primary_character_id = :primary_character_id, "
	}

	if len(user.Email) > 0 && prevUser.Email != user.Email {
		field += "email = :email, "
	}

	//if len(user.Password) > 0 {
	//	field += "password = :password, "
	//}

	if len(user.GoogleToken) > 0 && prevUser.GoogleToken != user.GoogleToken {
		field += "google_token = :google_token, "
	}

	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", userTable, field)
	result, err := s.db.NamedExec(query, user)
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

//DeleteUser will grab data from storage
func (s *Storage) DeleteUser(user *model.User) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", userTable)
	result, err := s.db.Exec(query, user.ID)
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

//createTableUser will grab data from storage
func (s *Storage) createTableUser() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE user (
  id int(11) NOT NULL AUTO_INCREMENT,
  display_name varchar(64) NOT NULL DEFAULT '',
  primary_account_id int(11) unsigned NOT NULL DEFAULT '0',
  primary_character_id int(11) unsigned NOT NULL DEFAULT '0',
  email varchar(128) NOT NULL DEFAULT '',
  password varchar(128) NOT NULL DEFAULT '',
  google_token varchar(128) NOT NULL DEFAULT '',
  PRIMARY KEY (id),
  UNIQUE KEY email (email)
) ENGINE=INNODB DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
