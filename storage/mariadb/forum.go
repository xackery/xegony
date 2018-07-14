package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	forumTable  = "forum"
	forumFields = "id, name, user_id, description, icon"
	forumBinds  = ":id, :name, :user_id, :description, :icon"
)

//GetForum will grab data from storage
func (s *Storage) GetForum(forum *model.Forum) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", forumFields, forumTable)
	err = s.db.Get(forum, query, forum.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

func (s *Storage) insertTestForum() (err error) {
	_, err = s.db.Exec("INSERT INTO `forum` (`id`, `name`, `user_id`, `description`, `last_modified`, `create_date`, `icon`, `sort`) VALUES (1, 'Test', 1, 'Testing Forum', '2018-07-14 20:13:06', '2018-07-14 20:13:06', '', 0);")
	if err != nil {
		err = errors.Wrap(err, "failed to insert npc data")
		return
	}
	return
}

//CreateForum will grab data from storage
func (s *Storage) CreateForum(forum *model.Forum) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", forumTable, forumFields, forumBinds)
	result, err := s.db.NamedExec(query, forum)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	forumID, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	forum.ID = forumID
	return
}

//ListForum will grab data from storage
func (s *Storage) ListForum(page *model.Page) (forums []*model.Forum, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "sort"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", forumFields, forumTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		forum := model.Forum{}
		if err = rows.StructScan(&forum); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		forums = append(forums, &forum)
	}
	return
}

//ListForumTotalCount will grab data from storage
func (s *Storage) ListForumTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(id) FROM %s", forumTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListForumBySearch will grab data from storage
func (s *Storage) ListForumBySearch(page *model.Page, forum *model.Forum) (forums []*model.Forum, err error) {

	field := ""

	if len(forum.Name) > 0 {
		field += `name LIKE :name OR`
		forum.Name = fmt.Sprintf("%%%s%%", forum.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", forumFields, forumTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, forum)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		forum := model.Forum{}
		if err = rows.StructScan(&forum); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		forums = append(forums, &forum)
	}
	return
}

//ListForumBySearchTotalCount will grab data from storage
func (s *Storage) ListForumBySearchTotalCount(forum *model.Forum) (count int64, err error) {
	field := ""
	if len(forum.Name) > 0 {
		field += `name LIKE :name OR`
		forum.Name = fmt.Sprintf("%%%s%%", forum.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", forumTable, field)

	rows, err := s.db.NamedQuery(query, forum)
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

//EditForum will grab data from storage
func (s *Storage) EditForum(forum *model.Forum) (err error) {

	prevForum := &model.Forum{
		ID: forum.ID,
	}
	err = s.GetForum(prevForum)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous forum")
		return
	}

	field := ""
	if len(forum.Name) > 0 && prevForum.Name != forum.Name {
		field += "name = :name, "
	}
	if len(forum.Description) > 0 && prevForum.Description != forum.Description {
		field += "description = :description, "
	}
	if len(forum.Icon) > 0 && prevForum.Icon != forum.Icon {
		field += "icon = :icon, "
	}
	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", forumTable, field)
	result, err := s.db.NamedExec(query, forum)
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
	fmt.Println(forum)
	return
}

//DeleteForum will grab data from storage
func (s *Storage) DeleteForum(forum *model.Forum) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", forumTable)
	result, err := s.db.Exec(query, forum.ID)
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

//createTableForum will grab data from storage
func (s *Storage) createTableForum() (err error) {
	_, err = s.db.Exec(`
		CREATE TABLE forum (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			name varchar(32) NOT NULL DEFAULT '',
			user_id int(11) unsigned NOT NULL,
			description varchar(128) NOT NULL DEFAULT '',
			last_modified timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			create_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			icon varchar(32) NOT NULL DEFAULT '',
			sort int(10) unsigned NOT NULL,
			PRIMARY KEY (id)
		  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
		`)
	if err != nil {
		return
	}
	return
}
