package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	spawnNpcFields = `spawngroupID, npcID, chance`
	spawnNpcBinds  = `:spawngroupID, :npcID, :chance`
	spawnNpcSets   = `spawngroupID=:spawngroupID, npcID=:npcID, chance=:chance`
	spawnNpcTable  = `spawnentry`
)

//GetSpawnNpc will grab data from storage
func (s *Storage) GetSpawnNpc(spawnNpc *model.SpawnNpc) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE spawngroupid = ? AND npcid = ?", spawnNpcFields, spawnNpcTable)
	err = s.db.Get(spawnNpc, query, spawnNpc.SpawnID, spawnNpc.NpcID)
	if err != nil {
		return
	}
	return
}

//CreateSpawnNpc will grab data from storage
func (s *Storage) CreateSpawnNpc(spawnNpc *model.SpawnNpc) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", spawnNpcTable, spawnNpcFields, spawnNpcBinds)
	_, err = s.db.NamedExec(query, spawnNpc)
	if err != nil {
		return
	}
	return
}

//ListSpawnNpc will grab data from storage
func (s *Storage) ListSpawnNpc() (spawnNpcs []*model.SpawnNpc, err error) {
	query := fmt.Sprintf(`SELECT %s FROM %s ORDER BY id DESC`, spawnNpcFields, spawnNpcTable)
	rows, err := s.db.Queryx(query)
	if err != nil {
		return
	}

	for rows.Next() {
		spawnNpc := model.SpawnNpc{}
		if err = rows.StructScan(&spawnNpc); err != nil {
			return
		}
		spawnNpcs = append(spawnNpcs, &spawnNpc)
	}
	return
}

//ListSpawnNpcBySpawn will grab data from storage
func (s *Storage) ListSpawnNpcBySpawn(spawn *model.Spawn) (spawnNpcs []*model.SpawnNpc, err error) {
	query := fmt.Sprintf(`SELECT %s FROM %s WHERE spawngroupid = ? ORDER BY spawngroupid DESC`, spawnNpcFields, spawnNpcTable)
	rows, err := s.db.Queryx(query, spawn.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		spawnNpc := model.SpawnNpc{}
		if err = rows.StructScan(&spawnNpc); err != nil {
			return
		}
		spawnNpcs = append(spawnNpcs, &spawnNpc)
	}
	return
}

//ListSpawnNpcByNpc will grab data from storage
func (s *Storage) ListSpawnNpcByNpc(npc *model.Npc) (spawnNpcs []*model.SpawnNpc, err error) {
	query := fmt.Sprintf(`SELECT %s FROM %s WHERE npcid = ? ORDER BY spawngroupid DESC`, spawnNpcFields, spawnNpcTable)
	rows, err := s.db.Queryx(query, npc.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		spawnNpc := model.SpawnNpc{}
		if err = rows.StructScan(&spawnNpc); err != nil {
			return
		}
		spawnNpcs = append(spawnNpcs, &spawnNpc)
	}
	return
}

//EditSpawnNpc will grab data from storage
func (s *Storage) EditSpawnNpc(spawnNpc *model.SpawnNpc) (err error) {
	query := fmt.Sprintf("UPDATE %s SET %s WHERE spawngroupid = :spawngroupid AND npcid= :npcid", spawnNpcTable, spawnNpcSets)
	result, err := s.db.NamedExec(query, spawnNpc)
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

//DeleteSpawnNpc will grab data from storage
func (s *Storage) DeleteSpawnNpc(spawnNpc *model.SpawnNpc) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE spawngroupid = ? AND npcid = ?", spawnNpcTable)
	result, err := s.db.Exec(query, spawnNpc.SpawnID, spawnNpc.NpcID)
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

//createTableSpawnNpc will grab data from storage
func (s *Storage) createTableSpawnNpc() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE spawnentry (
  spawngroupID int(11) NOT NULL DEFAULT '0',
  npcID int(11) NOT NULL DEFAULT '0',
  chance smallint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (spawngroupID,npcID)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
