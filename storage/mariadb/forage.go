package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	forageFields = `zoneid, Itemid, level, chance`
	forageSets   = `zoneid=:zoneid, Itemid=:Itemid, level=:level, chance=:chance`
	forageBinds  = `:zoneid, :Itemid, :level, :chance`
)

//GetForage will grab data from storage
func (s *Storage) GetForage(forage *model.Forage) (err error) {
	err = s.db.Get(forage, fmt.Sprintf("SELECT id, %s FROM forage WHERE id = ?", forageFields), forage.ID)
	if err != nil {
		return
	}
	return
}

//CreateForage will grab data from storage
func (s *Storage) CreateForage(forage *model.Forage) (err error) {
	if forage == nil {
		err = fmt.Errorf("Must provide forage")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO forage(%s)
		VALUES (%s)`, forageFields, forageBinds), forage)
	if err != nil {
		return
	}
	forageID, err := result.LastInsertId()
	if err != nil {
		return
	}
	forage.ID = forageID
	return
}

//ListForage will grab data from storage
func (s *Storage) ListForage(pageSize int64, pageNumber int64) (forages []*model.Forage, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM forage
		ORDER BY id ASC LIMIT %d OFFSET %d`, forageFields, pageSize, pageSize*pageNumber))
	if err != nil {
		return
	}

	for rows.Next() {
		forage := model.Forage{}
		if err = rows.StructScan(&forage); err != nil {
			return
		}
		forages = append(forages, &forage)
	}
	return
}

//ListForageCount will grab data from storage
func (s *Storage) ListForageCount() (count int64, err error) {
	err = s.db.Get(&count, `SELECT count(id) FROM forage`)
	if err != nil {
		return
	}
	return
}

//ListForageByItem will grab data from storage
func (s *Storage) ListForageByItem(item *model.Item) (forages []*model.Forage, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT forage.id, %s FROM forage		
		WHERE forage.itemid = ?`, forageFields), item.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		forage := model.Forage{}
		if err = rows.StructScan(&forage); err != nil {
			return
		}
		forages = append(forages, &forage)
	}
	return
}

//ListForageByZone will grab data from storage
func (s *Storage) ListForageByZone(zone *model.Zone) (forages []*model.Forage, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT forage.id, %s FROM forage		
		WHERE forage.zoneid = ?`, forageFields), zone.ZoneIDNumber)
	if err != nil {
		return
	}

	for rows.Next() {
		forage := model.Forage{}
		if err = rows.StructScan(&forage); err != nil {
			return
		}
		forages = append(forages, &forage)
	}
	return
}

//EditForage will grab data from storage
func (s *Storage) EditForage(forage *model.Forage) (err error) {
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE forage SET %s WHERE id = :id`, forageSets), forage)
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

//DeleteForage will grab data from storage
func (s *Storage) DeleteForage(forage *model.Forage) (err error) {
	result, err := s.db.Exec(`DELETE FROM forage WHERE id = ?`, forage.ID)
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
