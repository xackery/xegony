package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	aaRankSets   = `upper_hotkey_sid, lower_hotkey_sid, title_sid, desc_sid, cost, level_req, spell, spell_type, recast_time, expansion, prev_id, next_id,`
	aaRankFields = `upper_hotkey_sid, lower_hotkey_sid, title_sid, desc_sid, cost, level_req, spell, spell_type, recast_time, expansion, prev_id, next_id,`
	aaRankBinds  = `upper_hotkey_sid, lower_hotkey_sid, title_sid, desc_sid, cost, level_req, spell, spell_type, recast_time, expansion, prev_id, next_id,`
)

//GetAaRank will grab data from storage
func (s *Storage) GetAaRank(rankID int64) (query string, aaRank *model.AaRank, err error) {
	aaRank = &model.AaRank{}
	query = fmt.Sprintf(`SELECT %s FROM aa_ranks 
		WHERE id = ?`, aaRankFields)
	err = s.db.Get(aaRank, query, rankID)
	if err != nil {
		return
	}
	return
}

//CreateAaRank will grab data from storage
func (s *Storage) CreateAaRank(aaRank *model.AaRank) (query string, err error) {
	if aaRank == nil {
		err = fmt.Errorf("Must provide aaRank")
		return
	}

	query = fmt.Sprintf(`INSERT INTO aa_ranks(%s)
		VALUES (%s)`, aaRankFields, aaRankBinds)
	_, err = s.db.NamedExec(query, aaRank)
	if err != nil {
		return
	}
	return
}

//ListAaRank will grab data from storage
func (s *Storage) ListAaRank() (query string, aaRanks []*model.AaRank, err error) {
	query = fmt.Sprintf(`SELECT %s FROM aa_ranks`, aaRankFields)
	rows, err := s.db.Queryx(query)
	if err != nil {
		return
	}

	for rows.Next() {
		aaRank := model.AaRank{}
		if err = rows.StructScan(&aaRank); err != nil {
			return
		}
		aaRanks = append(aaRanks, &aaRank)
	}
	return
}

//EditAaRank will grab data from storage
func (s *Storage) EditAaRank(rankID int64, aaRank *model.AaRank) (query string, err error) {

	query = fmt.Sprintf(`UPDATE aa_ranks SET %s WHERE id = ?`, aaRankSets)
	aaRank.ID = rankID
	result, err := s.db.NamedExec(query, aaRank)
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

//DeleteAaRank will grab data from storage
func (s *Storage) DeleteAaRank(rankID int64) (query string, err error) {
	query = `DELETE FROM aa_ranks WHERE id = ?`
	result, err := s.db.Exec(query, rankID)
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
