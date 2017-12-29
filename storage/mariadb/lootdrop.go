package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	lootDropSets   = `name=:name`
	lootDropFields = `name`
	lootDropBinds  = `:name`
)

func (s *Storage) GetLootDrop(lootDropId int64) (lootDrop *model.LootDrop, err error) {
	lootDrop = &model.LootDrop{}
	err = s.db.Get(lootDrop, fmt.Sprintf("SELECT loottable.id, %s FROM lootdrop WHERE id = ?", lootDropFields), lootDropId)
	if err != nil {
		return
	}

	lootDrop.Entries, err = s.ListLootDropEntry(lootDrop.Id)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateLootDrop(lootDrop *model.LootDrop) (err error) {
	if lootDrop == nil {
		err = fmt.Errorf("Must provide lootDrop")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO loottable(%s)
		VALUES (%s)`, lootDropFields, lootDropBinds), lootDrop)
	if err != nil {
		return
	}
	lootDropId, err := result.LastInsertId()
	if err != nil {
		return
	}
	lootDrop.Id = lootDropId
	return
}

func (s *Storage) ListLootDrop() (lootDrops []*model.LootDrop, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT loottable.id, %s FROM loottable LIMIT 50`, lootDropFields))
	if err != nil {
		return
	}

	for rows.Next() {
		lootDrop := model.LootDrop{}
		if err = rows.StructScan(&lootDrop); err != nil {
			return
		}
		lootDrops = append(lootDrops, &lootDrop)
	}
	return
}

func (s *Storage) EditLootDrop(lootDropId int64, lootDrop *model.LootDrop) (err error) {
	lootDrop.Id = lootDropId
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE loottable SET %s WHERE id = :id`, lootDropSets), lootDrop)
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

func (s *Storage) DeleteLootDrop(lootDropId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM loottable WHERE id = ?`, lootDropId)
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
