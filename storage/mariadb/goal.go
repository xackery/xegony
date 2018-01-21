package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetGoal will grab data from storage
func (s *Storage) GetGoal(goal *model.Goal) (err error) {
	goal = &model.Goal{}
	err = s.db.Get(goal, "SELECT listid, entry FROM goallists WHERE listid = ? AND entry = ?", goal.ListID, goal.EntryID)
	if err != nil {
		return
	}
	return
}

//CreateGoal will grab data from storage
func (s *Storage) CreateGoal(goal *model.Goal) (err error) {
	if goal == nil {
		err = fmt.Errorf("Must provide goal")
		return
	}
	//we need to find out the highest id before adding.
	newID := int64(0)
	if err = s.db.Get(&newID, "SELECT listid FROM goallists ORDER BY listid DESC LIMIT 1"); err != nil {
		return
	}

	newID++
	goal.ListID = newID
	_, err = s.db.NamedExec(`INSERT INTO goallists(listid, entry)
		VALUES (:listid, :entry)`, goal)
	if err != nil {
		return
	}

	return
}

//ListGoal will grab data from storage
func (s *Storage) ListGoal() (goals []*model.Goal, err error) {
	rows, err := s.db.Queryx(`SELECT listid, entry FROM goallists`)
	if err != nil {
		return
	}

	for rows.Next() {
		goal := model.Goal{}
		if err = rows.StructScan(&goal); err != nil {
			return
		}
		goals = append(goals, &goal)
	}
	return
}

//EditGoal will grab data from storage
func (s *Storage) EditGoal(goal *model.Goal) (err error) {
	result, err := s.db.NamedExec(`UPDATE goallists SET listid=:listid, entry=:entry WHERE listid = :listid`, goal)
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

//DeleteGoal will grab data from storage
func (s *Storage) DeleteGoal(goal *model.Goal) (err error) {
	result, err := s.db.Exec(`DELETE FROM goallists WHERE listid = ? and entry = ?`, goal.ListID, goal.EntryID)
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
