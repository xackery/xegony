package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

func (s *Storage) GetTopic(topicId int64) (topic *model.Topic, err error) {
	topic = &model.Topic{}
	err = s.db.Get(topic, "SELECT id, icon, title, forum_id FROM topic WHERE id = ?", topicId)
	if err != nil {
		return
	}
	return
}

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
	topicId, err := result.LastInsertId()
	if err != nil {
		return
	}
	topic.Id = topicId
	return
}

func (s *Storage) ListTopic(forumId int64) (topics []*model.Topic, err error) {
	rows, err := s.db.Queryx(`SELECT id, title, icon, forum_id FROM topic WHERE forum_id = ? ORDER BY id DESC`, forumId)
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

func (s *Storage) EditTopic(topicId int64, topic *model.Topic) (err error) {
	topic.Id = topicId
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

func (s *Storage) DeleteTopic(topicId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM topic WHERE id = ?`, topicId)
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
