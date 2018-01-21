package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetPost will grab data from storage
func (s *Storage) GetPost(post *model.Post) (err error) {
	err = s.db.Get(post, "SELECT id, body, topic_id FROM post WHERE id = ?", post.ID)
	if err != nil {
		return
	}
	return
}

//CreatePost will grab data from storage
func (s *Storage) CreatePost(post *model.Post) (err error) {
	if post == nil {
		err = fmt.Errorf("Must provide post")
		return
	}

	result, err := s.db.NamedExec(`INSERT INTO post(body, topic_id)
		VALUES (:body, :topic_id)`, post)
	if err != nil {
		return
	}
	postID, err := result.LastInsertId()
	if err != nil {
		return
	}
	post.ID = postID
	return
}

//ListPost will grab data from storage
func (s *Storage) ListPostByTopic(topic *model.Topic) (posts []*model.Post, err error) {
	rows, err := s.db.Queryx(`SELECT id, body, topic_id FROM post WHERE topic_id = ? ORDER BY id DESC`, topic.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		post := model.Post{}
		if err = rows.StructScan(&post); err != nil {
			return
		}
		posts = append(posts, &post)
	}
	return
}

//EditPost will grab data from storage
func (s *Storage) EditPost(post *model.Post) (err error) {
	result, err := s.db.NamedExec(`UPDATE post SET body=:body, topic_id=:topic_id WHERE id = :id`, post)
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

//DeletePost will grab data from storage
func (s *Storage) DeletePost(post *model.Post) (err error) {
	result, err := s.db.Exec(`DELETE FROM post WHERE id = ?`, post.ID)
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

//createTablePost will grab data from storage
func (s *Storage) createTablePost() (err error) {
	_, err = s.db.Exec(`CREATE TABLE post (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  body text NOT NULL,
  owner_id int(11) unsigned NOT NULL,
  topic_id int(11) unsigned NOT NULL,
  last_modified timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  create_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  icon varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	return
}
