package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	spawnEntrySets   = `spawngroupID=:spawngroupID, npcid=:npcID, chance=:chance`
	spawnEntryFields = `spawngroupID, npcID, chance`
	spawnEntryBinds  = `:spawngroupID, :npcID, :chance`
)

func (s *Storage) GetSpawnEntry(spawnGroupID int64, npcID int64) (spawnEntry *model.SpawnEntry, err error) {
	spawnEntry = &model.SpawnEntry{}
	err = s.db.Get(spawnEntry, fmt.Sprintf(`SELECT %s FROM spawnentry 
		WHERE spawnentry.spawngroupid = ? AND spawnentry.npcid = ?`, spawnEntryFields), spawnGroupID, npcID)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateSpawnEntry(spawnEntry *model.SpawnEntry) (err error) {
	if spawnEntry == nil {
		err = fmt.Errorf("Must provide spawnEntry")
		return
	}

	_, err = s.db.NamedExec(fmt.Sprintf(`INSERT INTO spawnentry(%s)
		VALUES (%s)`, spawnEntryFields, spawnEntryBinds), spawnEntry)
	if err != nil {
		return
	}
	return
}

func (s *Storage) ListSpawnEntry(spawnGroupID int64) (spawnEntrys []*model.SpawnEntry, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM spawnentry WHERE spawngroupid = ?`, spawnEntryFields), spawnGroupID)
	if err != nil {
		return
	}

	for rows.Next() {
		spawnEntry := model.SpawnEntry{}
		if err = rows.StructScan(&spawnEntry); err != nil {
			return
		}
		spawnEntrys = append(spawnEntrys, &spawnEntry)
	}
	return
}

func (s *Storage) EditSpawnEntry(spawnGroupID int64, npcID int64, spawnEntry *model.SpawnEntry) (err error) {

	spawnEntry.SpawngroupID = spawnGroupID
	spawnEntry.NpcID = npcID
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE spawnentry SET %s WHERE spawnentry.spawngroupid = ? AND spawnentry.npcid = ?`, spawnEntrySets), spawnEntry)
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

func (s *Storage) DeleteSpawnEntry(spawnGroupID int64, npcID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM spawnentry WHERE spawngroupid = ? AND npcid = ?`, spawnGroupID, npcID)
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
