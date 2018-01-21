package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	aaSets   = `name=:name, category=:category, classes=:classes, races=:races, drakkin_heritage=:drakkin_heritage, deities=:deities, status=:status, type=:type, charges=:charges, grant_only=:grant_only, first_rank_id=:first_rank_id, enabled=:enabled`
	aaFields = `name, category, classes, races, drakkin_heritage, deities, status, type, charges, grant_only, first_rank_id, enabled`
	aaBinds  = `:name, :category, :classes, :races, :drakkin_heritage, :deities, :status, :type, :charges, :grant_only, :first_rank_id, :enabled`
)

//GetAa will grab data from storage
func (s *Storage) GetAa(aa *model.Aa) (err error) {
	aa = &model.Aa{}
	err = s.db.Get(aa, fmt.Sprintf(`SELECT id, %s FROM aa_ability 
		WHERE id = ?`, aaFields), aa.ID)
	if err != nil {
		return
	}
	return
}

//CreateAa will grab data from storage
func (s *Storage) CreateAa(aa *model.Aa) (err error) {
	if aa == nil {
		err = fmt.Errorf("Must provide aa")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO aa_ability(%s)
		VALUES (%s)`, aaFields, aaBinds), aa)
	if err != nil {
		return
	}

	aaID, err := result.LastInsertId()
	if err != nil {
		return
	}
	aa.ID = aaID
	return
}

//ListAa will grab data from storage
func (s *Storage) ListAa() (aas []*model.Aa, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM aa_ability`, aaFields))
	if err != nil {
		return
	}

	for rows.Next() {
		aa := model.Aa{}
		if err = rows.StructScan(&aa); err != nil {
			return
		}
		aas = append(aas, &aa)
	}
	return
}

//EditAa will grab data from storage
func (s *Storage) EditAa(aa *model.Aa) (err error) {
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE aa_ability SET %s WHERE id = :id`, aaSets), aa)
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

//DeleteAa will grab data from storage
func (s *Storage) DeleteAa(aa *model.Aa) (err error) {
	result, err := s.db.Exec(`DELETE FROM aa WHERE id = ?`, aa.ID)
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
