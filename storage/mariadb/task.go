package mariadb

import (
	"database/sql"
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	taskFields = `duration, title, description, reward, rewardid, cashreward, xpreward, rewardmethod, startzone, minlevel, maxlevel, repeatable`
	taskBinds  = `:duration, :title, :description, :reward, :rewardid, :cashreward, :xpreward, :rewardmethod, :startzone, :minlevel, :maxlevel, :repeatable`
	taskSets   = `duration=:duration, title=:title, description=:description, reward=:reward, rewardid=:rewardid, cashreward=:cashreward, xpreward=:xpreward, rewardmethod=:rewardmethod, startzone=:startzone, minlevel=:minlevel, maxlevel=:maxlevel, repeatable=:repeatable`
)

//GetTask will grab data from storage
func (s *Storage) GetTask(taskID int64) (task *model.Task, err error) {
	task = &model.Task{}
	err = s.db.Get(task, fmt.Sprintf("SELECT id, %s FROM tasks WHERE id = ?", taskFields), taskID)
	if err != nil {
		return
	}
	return
}

func (s *Storage) GetTaskNextID() (taskID int64, err error) {
	err = s.db.Get(&taskID, "SELECT id FROM tasks ORDER BY ID DESC LIMIT 1")
	if err != nil {
		if err == sql.ErrNoRows {
			taskID = 1
			err = nil
			return
		}
		return
	}

	taskID++
	return
}

//CreateTask will grab data from storage
func (s *Storage) CreateTask(task *model.Task) (err error) {
	if task == nil {
		err = fmt.Errorf("Must provide task")
		return
	}

	taskID, err := s.GetTaskNextID()
	if err != nil {
		return
	}

	task.ID = taskID

	fmt.Println(task)
	_, err = s.db.NamedExec(fmt.Sprintf(`INSERT INTO tasks(id, %s)
		VALUES (:id, %s)`, taskFields, taskBinds), task)
	if err != nil {
		return
	}
	return
}

//ListTask will grab data from storage
func (s *Storage) ListTask() (tasks []*model.Task, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM tasks ORDER BY ID ASC`, taskFields))
	if err != nil {
		return
	}

	for rows.Next() {
		task := model.Task{}
		if err = rows.StructScan(&task); err != nil {
			return
		}
		tasks = append(tasks, &task)
	}
	return
}

//EditTask will grab data from storage
func (s *Storage) EditTask(taskID int64, task *model.Task) (err error) {
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE tasks SET %s WHERE id = :id`, taskSets), task)
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

//DeleteTask will grab data from storage
func (s *Storage) DeleteTask(taskID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM tasks WHERE id = ?`, taskID)
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
func (s *Storage) createTableTask() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE tasks (
  id int(11) unsigned NOT NULL DEFAULT '0',
  duration int(11) unsigned NOT NULL DEFAULT '0',
  title varchar(100) NOT NULL DEFAULT '',
  description text NOT NULL,
  reward varchar(64) NOT NULL DEFAULT '',
  rewardid int(11) unsigned NOT NULL DEFAULT '0',
  cashreward int(11) unsigned NOT NULL DEFAULT '0',
  xpreward int(10) NOT NULL DEFAULT '0',
  rewardmethod tinyint(3) unsigned NOT NULL DEFAULT '2',
  startzone int(11) NOT NULL DEFAULT '0',
  minlevel tinyint(3) unsigned NOT NULL DEFAULT '0',
  maxlevel tinyint(3) unsigned NOT NULL DEFAULT '0',
  repeatable tinyint(1) unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (id)
) ENGINE=INNODB DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
