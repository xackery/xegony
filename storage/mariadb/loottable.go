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

func (s *Storage) GetLootTable(lootTableId int64) (lootTable *model.LootTable, err error) {
	lootTable = &model.LootTable{}
	err = s.db.Get(lootTable, fmt.Sprintf(`SELECT loottable.id, %s FROM loottable WHERE id = ?`, lootTableFields), lootTableId)
	if err != nil {
		return
	}

	lootTable.Entries, err = s.ListLootTableEntry(lootTable.Id)
	if err != nil {
		return
	}
	lootTable.Npcs, err = s.ListNpcByLootTable(lootTable.Id)
	if err != nil {
		return
	}
	return
}

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
	lootTableId, err := result.LastInsertId()
	if err != nil {
		return
	}
	lootTable.Id = lootTableId
	return
}

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

func (s *Storage) EditLootTable(lootTableId int64, lootTable *model.LootTable) (err error) {
	lootTable.Id = lootTableId
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

func (s *Storage) DeleteLootTable(lootTableId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM loottable WHERE id = ?`, lootTableId)
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
