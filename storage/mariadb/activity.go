package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	activityFields = `taskid, activityid, step, activitytype, text1, text2, text3, goalid, goalmethod, goalcount, delivertonpc, zoneid, optional`
	activityBinds  = `:taskid, :activityid, :step, :activitytype, :text1, :text2, :text3, :goalid, :goalmethod, :goalcount, :delivertonpc, :zoneid, :optional`
	activitySets   = `taskid=:taskid, activityid=:activityid, step=:step, activitytype=:activitytype, text1=:text1, text2=:text2, text3=:text3, goalid=:goalid, goalmethod=:goalmethod, goalcount=:goalcount, delivertonpc=:delivertonpc, zoneid=:zoneid, optional=:optional`
)

func (s *Storage) GetActivity(taskId int64, activityId int64) (activity *model.Activity, err error) {
	activity = &model.Activity{}
	err = s.db.Get(activity, fmt.Sprintf("SELECT %s FROM activities WHERE activityid = ? AND taskid = ?", activityFields), activityId, taskId)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateActivity(activity *model.Activity) (err error) {
	if activity == nil {
		err = fmt.Errorf("Must provide activity")
		return
	}

	result, err := s.db.NamedExec(`INSERT INTO activity(title)
		VALUES (:title)`, activity)
	if err != nil {
		return
	}
	activityId, err := result.LastInsertId()
	if err != nil {
		return
	}
	activity.Activityid = activityId
	return
}

func (s *Storage) ListActivity(taskId int64) (activitys []*model.Activity, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM activities WHERE taskid = ? ORDER BY activityid DESC`, activityFields), taskId)
	if err != nil {
		return
	}

	for rows.Next() {
		activity := model.Activity{}
		if err = rows.StructScan(&activity); err != nil {
			return
		}
		activitys = append(activitys, &activity)
	}
	return
}

func (s *Storage) EditActivity(activityId int64, activity *model.Activity) (err error) {
	activity.Activityid = activityId
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE activities SET %s WHERE id = :id`, activitySets), activity)
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

func (s *Storage) DeleteActivity(activityId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM activities WHERE activityid = ?`, activityId)
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
