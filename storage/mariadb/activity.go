package mariadb

import (
	"database/sql"
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

//GetActivityNextStep will grab data from storage
func (s *Storage) GetActivityNextStep(taskID int64, activityID int64) (step int64, err error) {
	err = s.db.Get(&step, "SELECT step FROM activities WHERE taskid = ? AND activityid = ? ORDER BY step DESC LIMIT 1", taskID, activityID)
	if err != nil {
		if err == sql.ErrNoRows {
			step = 0
			err = nil
			return
		}
		return
	}
	step++
	return
}

//CreateActivity will grab data from storage
func (s *Storage) CreateActivity(activity *model.Activity) (err error) {
	if activity == nil {
		err = fmt.Errorf("Must provide activity")
		return
	}

	query := fmt.Sprintf(`INSERT INTO activities(%s)
		VALUES (%s)`, activityFields, activityBinds)
	result, err := s.db.NamedExec(query, activity)
	if err != nil {
		return
	}
	activityID, err := result.LastInsertId()
	if err != nil {
		return
	}
	activity.ActivityID = activityID
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
	activity.ActivityID = activityID
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

//createTableUser will grab data from storage
func (s *Storage) createTableActivity() (err error) {
	_, err = s.db.Exec(`
  CREATE TABLE activities (
  taskid int(11) unsigned NOT NULL DEFAULT '0',
  activityid int(11) unsigned NOT NULL DEFAULT '0',
  step int(11) unsigned NOT NULL DEFAULT '0',
  activitytype tinyint(3) unsigned NOT NULL DEFAULT '0',
  text1 varchar(64) NOT NULL DEFAULT '',
  text2 varchar(64) NOT NULL DEFAULT '',
  text3 varchar(128) NOT NULL DEFAULT '',
  goalid int(11) unsigned NOT NULL DEFAULT '0',
  goalmethod int(10) unsigned NOT NULL DEFAULT '0',
  goalcount int(11) DEFAULT '1',
  delivertonpc int(11) unsigned NOT NULL DEFAULT '0',
  zoneid int(11) NOT NULL DEFAULT '0',
  optional tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (taskid,activityid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	return
}
