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

//GetLootDrop will grab data from storage
func (s *Storage) GetLootDrop(lootDrop *model.LootDrop) (err error) {
	err = s.db.Get(lootDrop, fmt.Sprintf("SELECT lootdrop.id, %s FROM lootdrop WHERE id = ?", lootDropFields), lootDrop.ID)
	if err != nil {
		return
	}

	return
}

//CreateLootDrop will grab data from storage
func (s *Storage) CreateLootDrop(lootDrop *model.LootDrop) (err error) {
	if lootDrop == nil {
		err = fmt.Errorf("Must provide lootDrop")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO lootdrop(%s)
		VALUES (%s)`, lootDropFields, lootDropBinds), lootDrop)
	if err != nil {
		return
	}
	lootDropID, err := result.LastInsertId()
	if err != nil {
		return
	}
	lootDrop.ID = lootDropID
	return
}

//ListLootDrop will grab data from storage
func (s *Storage) ListLootDrop() (lootDrops []*model.LootDrop, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT lootdrop.id, %s FROM lootdrop LIMIT 50`, lootDropFields))
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

//EditLootDrop will grab data from storage
func (s *Storage) EditLootDrop(lootDrop *model.LootDrop) (err error) {
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE lootdrop SET %s WHERE id = :id`, lootDropSets), lootDrop)
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

//DeleteLootDrop will grab data from storage
func (s *Storage) DeleteLootDrop(lootDrop *model.LootDrop) (err error) {
	result, err := s.db.Exec(`DELETE FROM lootdrop WHERE id = ?`, lootDrop.ID)
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
