package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	spawnEntryFields = `spawngroupID, zone, version, x, y, z, heading, respawntime, variance, pathgrid, _condition, cond_value, enabled, animation`
	spawnEntryBinds  = `:spawngroupID, :zone, :version, :x, :y, :z, :heading, :respawntime, :variance, :pathgrid, :_condition, :cond_value, :enabled, :animation`
	spawnEntrySets   = `spawngroupID=:spawngroupID, zone=:zone, version=:version, x, y=:y, z=:z, heading=:heading, respawntime=:respawntime, variance=:variance, pathgrid=:pathgrid, _condition=:_condition, cond_value=:cond_value, enabled, animation=:animation`
	spawnEntryTable  = `spawn2`
)

//GetSpawnEntry will grab data from storage
func (s *Storage) GetSpawnEntry(spawnEntry *model.SpawnEntry) (err error) {
	query := fmt.Sprintf("SELECT id, %s FROM %s WHERE id = ?", spawnEntryFields, spawnEntryTable)
	err = s.db.Get(spawnEntry, query, spawnEntry.ID)
	if err != nil {
		return
	}
	return
}

//CreateSpawnEntry will grab data from storage
func (s *Storage) CreateSpawnEntry(spawnEntry *model.SpawnEntry) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", spawnEntryTable, spawnEntryFields, spawnEntryBinds)
	result, err := s.db.NamedExec(query, spawnEntry)
	if err != nil {
		return
	}
	spawnEntryID, err := result.LastInsertId()
	if err != nil {
		return
	}
	spawnEntry.ID = spawnEntryID
	return
}

//ListSpawnEntry will grab data from storage
func (s *Storage) ListSpawnEntry() (spawnEntrys []*model.SpawnEntry, err error) {
	query := fmt.Sprintf(`SELECT id, %s FROM %s ORDER BY id DESC`, spawnEntryFields, spawnEntryTable)
	rows, err := s.db.Queryx(query)
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

//ListSpawnEntryBySpawn will grab data from storage
func (s *Storage) ListSpawnEntryBySpawn(spawn *model.Spawn) (spawnEntrys []*model.SpawnEntry, err error) {
	query := fmt.Sprintf(`SELECT id, %s FROM %s WHERE spawngroupid = ? ORDER BY id DESC`, spawnEntryFields, spawnEntryTable)
	rows, err := s.db.Queryx(query, spawn.ID)
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
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", spawnEntryTable, spawnEntrySets)
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
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", spawnEntryTable)
	result, err := s.db.Exec(query, spawnEntry.ID)
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

//createTableSpawnEntry will grab data from storage
func (s *Storage) createTableSpawnEntry() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE spawn2 (
  id int(11) NOT NULL AUTO_INCREMENT,
  spawngroupID int(11) NOT NULL DEFAULT '0',
  zone varchar(32) DEFAULT NULL,
  version smallint(5) unsigned NOT NULL DEFAULT '0',
  x float(14,6) NOT NULL DEFAULT '0.000000',
  y float(14,6) NOT NULL DEFAULT '0.000000',
  z float(14,6) NOT NULL DEFAULT '0.000000',
  heading float(14,6) NOT NULL DEFAULT '0.000000',
  respawntime int(11) NOT NULL DEFAULT '0',
  variance int(11) NOT NULL DEFAULT '0',
  pathgrid int(10) NOT NULL DEFAULT '0',
  _condition mediumint(8) unsigned NOT NULL DEFAULT '0',
  cond_value mediumint(9) NOT NULL DEFAULT '1',
  enabled tinyint(3) unsigned NOT NULL DEFAULT '1',
  animation tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (id),
  KEY ZoneGroup (zone),
  KEY spawn2_spawngroupid_idx (spawngroupID) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
