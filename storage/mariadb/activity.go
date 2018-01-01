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

//GetActivity will grab data from storage
func (s *Storage) GetActivity(taskID int64, activityID int64) (activity *model.Activity, err error) {
	activity = &model.Activity{}
	err = s.db.Get(activity, fmt.Sprintf("SELECT %s FROM activities WHERE activityid = ? AND taskid = ?", activityFields), activityID, taskID)
	if err != nil {
		return
	}
	return
}

//CreateActivity will grab data from storage
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
	activityID, err := result.LastInsertId()
	if err != nil {
		return
	}
	activity.Activityid = activityID
	return
}

//ListActivity will grab data from storage
func (s *Storage) ListActivity(taskID int64) (activitys []*model.Activity, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM activities WHERE taskid = ? ORDER BY activityid DESC`, activityFields), taskID)
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

//EditActivity will grab data from storage
func (s *Storage) EditActivity(activityID int64, activity *model.Activity) (err error) {
	activity.Activityid = activityID
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

//DeleteActivity will grab data from storage
func (s *Storage) DeleteActivity(activityID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM activities WHERE activityid = ?`, activityID)
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
