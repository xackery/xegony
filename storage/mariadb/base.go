package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	baseFields = `level, class, hp, mana, end, unk1, unk2, hp_fac, mana_fac, end_fac`
	baseSets   = `:level, :class, :hp, :mana, :end, :unk1, :unk2, :hp_fac, :mana_fac, :end_fac`
	baseBinds  = `level=:level, class=:class, hp=:hp, mana=:mana, end=:end, unk1=:unk1, unk2=:unk2, hp_fac=:hp_fac, mana_fac=:mana_fac, end_fac=:end_fac`
)

//GetBase will grab data from storage
func (s *Storage) GetBase(base *model.Base) (err error) {
	base = &model.Base{}
	err = s.db.Get(base, fmt.Sprintf("SELECT %s FROM base_data WHERE level = ? AND class = ?", baseFields), base.Level, base.Class)
	if err != nil {
		return
	}
	return
}

//CreateBase will grab data from storage
func (s *Storage) CreateBase(base *model.Base) (err error) {
	if base == nil {
		err = fmt.Errorf("Must provide base")
		return
	}

	_, err = s.db.NamedExec(fmt.Sprintf(`INSERT INTO base_data(%s)
		VALUES (%s)`, baseFields, baseBinds), base)
	if err != nil {
		return
	}

	return
}

//ListBase will grab data from storage
func (s *Storage) ListBase() (bases []*model.Base, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM base_data ORDER BY class, level ASC`, baseFields))
	if err != nil {
		return
	}

	for rows.Next() {
		base := model.Base{}
		if err = rows.StructScan(&base); err != nil {
			return
		}
		bases = append(bases, &base)
	}
	return
}

//EditBase will grab data from storage
func (s *Storage) EditBase(base *model.Base) (err error) {
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE base_data SET %s WHERE level = :level AND class = :class`, baseSets), base)
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

//DeleteBase will grab data from storage
func (s *Storage) DeleteBase(base *model.Base) (err error) {
	result, err := s.db.Exec(`DELETE FROM base_data WHERE level = ? AND class = ?`, base.Level, base.Class)
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
