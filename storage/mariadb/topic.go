package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetTopic will grab data from storage
func (s *Storage) GetTopic(topic *model.Topic) (err error) {
	err = s.db.Get(topic, "SELECT id, icon, title, forum_id FROM topic WHERE id = ?", topic.ID)
	if err != nil {
		return
	}
	return
}

//CreateTopic will grab data from storage
func (s *Storage) CreateTopic(topic *model.Topic) (err error) {
	if topic == nil {
		err = fmt.Errorf("Must provide topic")
		return
	}

	result, err := s.db.NamedExec(`INSERT INTO topic(title, icon)
		VALUES (:title, :icon)`, topic)
	if err != nil {
		return
	}
	topicID, err := result.LastInsertId()
	if err != nil {
		return
	}
	topic.ID = topicID
	return
}

//ListTopicByForum will grab data from storage
func (s *Storage) ListTopicByForum(forum *model.Forum) (topics []*model.Topic, err error) {
	rows, err := s.db.Queryx(`SELECT id, title, icon, forum_id FROM topic WHERE forum_id = ? ORDER BY id DESC`, forum.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		topic := model.Topic{}
		if err = rows.StructScan(&topic); err != nil {
			return
		}
		topics = append(topics, &topic)
	}
	return
}

//EditTopic will grab data from storage
func (s *Storage) EditTopic(topic *model.Topic) (err error) {
	result, err := s.db.NamedExec(`UPDATE topic SET icon=:icon, title=:title, forum_id=:forum_id WHERE id = :id`, topic)
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

//DeleteTopic will grab data from storage
func (s *Storage) DeleteTopic(topic *model.Topic) (err error) {
	result, err := s.db.Exec(`DELETE FROM topic WHERE id = ?`, topic.ID)
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

//createTableTopic will grab data from storage
func (s *Storage) createTableTopic() (err error) {
	_, err = s.db.Exec(`CREATE TABLE topic (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  title varchar(32) NOT NULL DEFAULT '',
  owner_id int(11) unsigned NOT NULL,
  forum_id int(11) unsigned NOT NULL,
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
