package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	userLinkTable  = "user_link"
	userLinkFields = "id, link, account_id, character_id, create_date"
	userLinkBinds  = ":id, :link, :account_id, :character_id, :create_date"
)

//GetUserLink will grab data from storage
func (s *Storage) GetUserLink(userLink *model.UserLink) (err error) {

	query := fmt.Sprintf("SELECT %s FROM %s WHERE link = ?", userLinkFields, userLinkTable)
	err = s.db.Get(userLink, query, userLink.Link)
	if err != nil {
		err = errors.Wrapf(err, "query: %s, param: %s", query, userLink.Link)
		return
	}
	return
}

//CreateUserLink will grab data from storage
func (s *Storage) CreateUserLink(userLink *model.UserLink) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", userLinkTable, userLinkFields, userLinkBinds)
	result, err := s.db.NamedExec(query, userLink)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	userLinkID, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	userLink.ID = userLinkID
	return
}

//ListUserLink will grab data from storage
func (s *Storage) ListUserLink(page *model.Page) (userLinks []*model.UserLink, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "link"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", userLinkFields, userLinkTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		userLink := model.UserLink{}
		if err = rows.StructScan(&userLink); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		userLinks = append(userLinks, &userLink)
	}
	return
}

//ListUserLinkTotalCount will grab data from storage
func (s *Storage) ListUserLinkTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(id) FROM %s", userLinkTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListUserLinkBySearch will grab data from storage
func (s *Storage) ListUserLinkBySearch(page *model.Page, userLink *model.UserLink) (userLinks []*model.UserLink, err error) {

	field := ""

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", userLinkFields, userLinkTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, userLink)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		userLink := model.UserLink{}
		if err = rows.StructScan(&userLink); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		userLinks = append(userLinks, &userLink)
	}
	return
}

//ListUserLinkBySearchTotalCount will grab data from storage
func (s *Storage) ListUserLinkBySearchTotalCount(userLink *model.UserLink) (count int64, err error) {
	field := ""
	if len(userLink.Link) > 0 {
		field += `link LIKE :link OR`
		userLink.Link = fmt.Sprintf("%%%s%%", userLink.Link)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", userLinkTable, field)

	rows, err := s.db.NamedQuery(query, userLink)
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

//EditUserLink will grab data from storage
func (s *Storage) EditUserLink(userLink *model.UserLink) (err error) {

	prevUserLink := &model.UserLink{
		ID: userLink.ID,
	}
	err = s.GetUserLink(prevUserLink)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous userLink")
		return
	}

	field := ""

	//if len(userLink.Password) > 0 {
	//	field += "password = :password, "
	//}

	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", userLinkTable, field)
	result, err := s.db.NamedExec(query, userLink)
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

//DeleteUserLink will grab data from storage
func (s *Storage) DeleteUserLink(userLink *model.UserLink) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", userLinkTable)
	result, err := s.db.Exec(query, userLink.ID)
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

//DeleteUserLinkByAccount will grab data from storage
func (s *Storage) DeleteUserLinkByAccount(account *model.Account) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE account_id = ?", userLinkTable)
	result, err := s.db.Exec(query, account.ID)
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

//createTableUserLink will grab data from storage
func (s *Storage) createTableUserLink() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE user_link (
  id int(11) NOT NULL AUTO_INCREMENT,
  link varchar(64) NOT NULL DEFAULT '',
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
