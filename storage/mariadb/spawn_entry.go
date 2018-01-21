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
func (s *Storage) GetSpawnEntry(spawnEntry *model.SpawnEntry) (err error) {
	query := fmt.Sprintf(`SELECT %s FROM spawnentry 
		WHERE spawnentry.spawngroupid = ? AND spawnentry.npcid = ?`, spawnEntryFields)
	err = s.db.Get(spawnEntry, query, spawnEntry.SpawngroupID, spawnEntry.NpcID)
	if err != nil {
		return
	}
	return
}

//CreateSpawnEntry will grab data from storage
func (s *Storage) CreateSpawnEntry(spawnEntry *model.SpawnEntry) (err error) {
	if spawnEntry == nil {
		err = fmt.Errorf("Must provide spawnEntry")
		return
	}

	query := fmt.Sprintf(`INSERT INTO spawnentry(%s)
		VALUES (%s)`, spawnEntryFields, spawnEntryBinds)
	_, err = s.db.NamedExec(query, spawnEntry)
	if err != nil {
		return
	}
	return
}

//ListSpawnEntry will grab data from storage
func (s *Storage) ListSpawnEntryBySpawnGroup(spawnGroup *model.SpawnGroup) (spawnEntrys []*model.SpawnEntry, err error) {
	query := fmt.Sprintf(`SELECT %s FROM spawnentry WHERE spawngroupid = ?`, spawnEntryFields)
	rows, err := s.db.Queryx(query, spawnGroup.ID)
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
func (s *Storage) ListSpawnEntryByZone(zone *model.Zone) (spawnEntrys []*model.SpawnEntry, err error) {

	query := fmt.Sprintf(`SELECT %s FROM spawnentry
	WHERE npcID < ? and npcID > ?`, spawnEntryFields)
	upperID := (zone.ID * 1000) + 1000 - 1
	lowerID := (zone.ID * 1000) - 1

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

//ListSpawnEntryByNpc will grab data from storage
func (s *Storage) ListSpawnEntryByNpc(npc *model.Npc) (spawnEntrys []*model.SpawnEntry, err error) {

	query := fmt.Sprintf(`SELECT %s FROM spawnentry
	WHERE npcID = ?`, spawnEntryFields)

	rows, err := s.db.Queryx(query, npc.ID)
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
func (s *Storage) EditSpawnEntry(spawnEntry *model.SpawnEntry) (err error) {
	query := fmt.Sprintf(`UPDATE spawnentry SET %s WHERE spawnentry.spawngroupid = ? AND spawnentry.npcid = ?`, spawnEntrySets)
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
func (s *Storage) DeleteSpawnEntry(spawnEntry *model.SpawnEntry) (err error) {
	query := `DELETE FROM spawnentry WHERE spawngroupid = ? AND npcid = ?`
	result, err := s.db.Exec(query, spawnEntry.SpawngroupID, spawnEntry.NpcID)
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
