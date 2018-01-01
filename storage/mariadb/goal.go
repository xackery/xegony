package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

func (s *Storage) GetGoal(listID int64, entryID int64) (goal *model.Goal, err error) {
	goal = &model.Goal{}
	err = s.db.Get(goal, "SELECT listid, entry FROM goallists WHERE listid = ? AND entry = ?", listID, entryID)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateGoal(goal *model.Goal) (err error) {
	if goal == nil {
		err = fmt.Errorf("Must provide goal")
		return
	}
	//we need to find out the highest id before adding.
	newId := int64(0)
	if err = s.db.Get(&newId, "SELECT listid FROM goallists ORDER BY listid DESC LIMIT 1"); err != nil {
		return
	}

	newId++
	goal.ListID = newId
	_, err = s.db.NamedExec(`INSERT INTO goallists(listid, entry)
		VALUES (:listid, :entry)`, goal)
	if err != nil {
		return
	}

	return
}

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

func (s *Storage) EditGoal(listID int64, goal *model.Goal) (err error) {
	goal.ListID = listID
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

func (s *Storage) DeleteGoal(listID int64, entryID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM goallists WHERE listid = ? and entry = ?`, listID, entryID)
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
