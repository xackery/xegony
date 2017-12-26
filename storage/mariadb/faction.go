package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

func (s *Storage) GetFaction(factionId int64) (faction *model.Faction, err error) {
	faction = &model.Faction{}
	err = s.db.Get(faction, "SELECT id, name, base FROM faction_list WHERE id = ?", factionId)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateFaction(faction *model.Faction) (err error) {
	if faction == nil {
		err = fmt.Errorf("Must provide faction")
		return
	}

	result, err := s.db.NamedExec(`INSERT INTO faction_list(name, base)
		VALUES (:name, :base)`, faction)
	if err != nil {
		return
	}
	factionId, err := result.LastInsertId()
	if err != nil {
		return
	}
	faction.Id = factionId
	return
}

func (s *Storage) ListFaction() (factions []*model.Faction, err error) {
	rows, err := s.db.Queryx(`SELECT id, name, base FROM faction_list ORDER BY id DESC`)
	if err != nil {
		return
	}

	for rows.Next() {
		faction := model.Faction{}
		if err = rows.StructScan(&faction); err != nil {
			return
		}
		factions = append(factions, &faction)
	}
	return
}

func (s *Storage) EditFaction(factionId int64, faction *model.Faction) (err error) {
	faction.Id = factionId
	result, err := s.db.NamedExec(`UPDATE faction_list SET name=:name, base=:base WHERE id = :id`, faction)
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

func (s *Storage) DeleteFaction(factionId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM faction_list WHERE id = ?`, factionId)
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
