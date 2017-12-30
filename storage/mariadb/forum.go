package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

func (s *Storage) GetForum(forumId int64) (forum *model.Forum, err error) {
	forum = &model.Forum{}
	err = s.db.Get(forum, "SELECT id, owner_id, description, icon, name FROM forum WHERE id = ?", forumId)
	if err != nil {
		return
	}
	fmt.Println("Forum", forum)
	return
}

func (s *Storage) CreateForum(forum *model.Forum) (err error) {
	if forum == nil {
		err = fmt.Errorf("Must provide forum")
		return
	}

	result, err := s.db.NamedExec(`INSERT INTO forum(name, icon,  description, owner_id)
		VALUES (:name, :icon,  :description, :owner_id)`, forum)
	if err != nil {
		return
	}
	forumId, err := result.LastInsertId()
	if err != nil {
		return
	}
	forum.Id = forumId
	return
}

func (s *Storage) ListForum() (forums []*model.Forum, err error) {
	rows, err := s.db.Queryx(`SELECT id, icon,  description, name, owner_id FROM forum ORDER BY id DESC`)
	if err != nil {
		return
	}

	for rows.Next() {
		forum := model.Forum{}
		if err = rows.StructScan(&forum); err != nil {
			return
		}
		forums = append(forums, &forum)
	}
	return
}

func (s *Storage) EditForum(forumId int64, forum *model.Forum) (err error) {
	forum.Id = forumId
	result, err := s.db.NamedExec(`UPDATE forum SET icon=:icon, name=:name, description=:description WHERE id = :id`, forum)
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

func (s *Storage) DeleteForum(forumId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM forum WHERE id = ?`, forumId)
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

func (s *Storage) createTableForum() (err error) {
	_, err = s.db.Exec(`CREATE TABLE forum (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(32) NOT NULL DEFAULT '',
  owner_id int(11) unsigned NOT NULL,
  description varchar(128) NOT NULL DEFAULT '',
  last_modified timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  create_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  icon varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	return
}
