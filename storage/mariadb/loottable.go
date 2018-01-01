package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	lootTableSets   = `name=:name, mincash=:mincash, maxcash=:maxcash, avgcoin=:avgcoin, done=:done`
	lootTableFields = `name, mincash, maxcash, avgcoin, done`
	lootTableBinds  = `:name, :mincash, :maxcash, :avgcoin, :done`
)

//GetLootTable will grab data from storage
func (s *Storage) GetLootTable(lootTableID int64) (lootTable *model.LootTable, err error) {
	lootTable = &model.LootTable{}
	err = s.db.Get(lootTable, fmt.Sprintf(`SELECT loottable.id, %s FROM loottable WHERE id = ?`, lootTableFields), lootTableID)
	if err != nil {
		return
	}

	lootTable.Entries, err = s.ListLootTableEntry(lootTable.ID)
	if err != nil {
		return
	}
	lootTable.Npcs, err = s.ListNpcByLootTable(lootTable.ID)
	if err != nil {
		return
	}
	return
}

//CreateLootTable will grab data from storage
func (s *Storage) CreateLootTable(lootTable *model.LootTable) (err error) {
	if lootTable == nil {
		err = fmt.Errorf("Must provide lootTable")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO loottable(%s)
		VALUES (%s)`, lootTableFields, lootTableBinds), lootTable)
	if err != nil {
		return
	}
	lootTableID, err := result.LastInsertId()
	if err != nil {
		return
	}
	lootTable.ID = lootTableID
	return
}

//ListLootTable will grab data from storage
func (s *Storage) ListLootTable() (lootTables []*model.LootTable, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT loottable.id, %s FROM loottable LIMIT 50`, lootTableFields))
	if err != nil {
		return
	}

	for rows.Next() {
		lootTable := model.LootTable{}
		if err = rows.StructScan(&lootTable); err != nil {
			return
		}
		lootTables = append(lootTables, &lootTable)
	}
	return
}

//EditLootTable will grab data from storage
func (s *Storage) EditLootTable(lootTableID int64, lootTable *model.LootTable) (err error) {
	lootTable.ID = lootTableID
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE loottable SET %s WHERE id = :id`, lootTableSets), lootTable)
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

//DeleteLootTable will grab data from storage
func (s *Storage) DeleteLootTable(lootTableID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM loottable WHERE id = ?`, lootTableID)
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
