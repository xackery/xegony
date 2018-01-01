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

//GetSpawnEntry will grab data from storage
func (s *Storage) GetSpawnEntry(spawnGroupID int64, npcID int64) (query string, spawnEntry *model.SpawnEntry, err error) {
	spawnEntry = &model.SpawnEntry{}
	query = fmt.Sprintf(`SELECT %s FROM spawnentry 
		WHERE spawnentry.spawngroupid = ? AND spawnentry.npcid = ?`, spawnEntryFields)
	err = s.db.Get(spawnEntry, query, spawnGroupID, npcID)
	if err != nil {
		return
	}
	return
}

//CreateSpawnEntry will grab data from storage
func (s *Storage) CreateSpawnEntry(spawnEntry *model.SpawnEntry) (query string, err error) {
	if spawnEntry == nil {
		err = fmt.Errorf("Must provide spawnEntry")
		return
	}

	query = fmt.Sprintf(`INSERT INTO spawnentry(%s)
		VALUES (%s)`, spawnEntryFields, spawnEntryBinds)
	_, err = s.db.NamedExec(query, spawnEntry)
	if err != nil {
		return
	}
	return
}

//ListSpawnEntry will grab data from storage
func (s *Storage) ListSpawnEntry(spawnGroupID int64) (query string, spawnEntrys []*model.SpawnEntry, err error) {
	query = fmt.Sprintf(`SELECT %s FROM spawnentry WHERE spawngroupid = ?`, spawnEntryFields)
	rows, err := s.db.Queryx(query, spawnGroupID)
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

//ListSpawnEntryByZone will grab data from storage
func (s *Storage) ListSpawnEntryByZone(zoneID int64) (query string, spawnEntrys []*model.SpawnEntry, err error) {

	query = fmt.Sprintf(`SELECT %s FROM spawnentry
	WHERE npcID < ? and npcID > ?`, spawnEntryFields)
	upperID := (zoneID * 1000) + 1000 - 1
	lowerID := (zoneID * 1000) - 1

	rows, err := s.db.Queryx(query, upperID, lowerID)
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

//EditSpawnEntry will grab data from storage
func (s *Storage) EditSpawnEntry(spawnGroupID int64, npcID int64, spawnEntry *model.SpawnEntry) (query string, err error) {

	query = fmt.Sprintf(`UPDATE spawnentry SET %s WHERE spawnentry.spawngroupid = ? AND spawnentry.npcid = ?`, spawnEntrySets)
	spawnEntry.SpawngroupID = spawnGroupID
	spawnEntry.NpcID = npcID
	result, err := s.db.NamedExec(query, spawnEntry)
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

//DeleteSpawnEntry will grab data from storage
func (s *Storage) DeleteSpawnEntry(spawnGroupID int64, npcID int64) (query string, err error) {
	query = `DELETE FROM spawnentry WHERE spawngroupid = ? AND npcid = ?`
	result, err := s.db.Exec(query, spawnGroupID, npcID)
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
