package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	spawnFields = `name,spawn_limit,dist,max_x,min_x,max_y,min_y,delay,mindelay,despawn,despawn_timer`
	spawnBinds  = `:name,:spawn_limit,:dist,:max_x,:min_x,:max_y,:min_y,:delay,:mindelay,:despawn,:despawn_timer`
	spawnSets   = `name=:name,spawn_limit=:spawn_limit,dist=:dist,max_x=:max_x,min_x=:min_x,max_y=:max_y,min_y,delay=:delay,mindelay=:mindelay,despawn,despawn_timer=:despawn_timer`
	spawnTable  = `spawngroup`
)

//GetSpawn will grab data from storage
func (s *Storage) GetSpawn(spawn *model.Spawn) (err error) {
	query := fmt.Sprintf("SELECT id, %s FROM %s WHERE id = ?", spawnFields, spawnTable)
	err = s.db.Get(spawn, query, spawn.ID)
	if err != nil {
		return
	}
	return
}

//CreateSpawn will grab data from storage
func (s *Storage) CreateSpawn(spawn *model.Spawn) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", spawnTable, spawnFields, spawnBinds)
	result, err := s.db.NamedExec(query, spawn)
	if err != nil {
		return
	}
	spawnID, err := result.LastInsertId()
	if err != nil {
		return
	}
	spawn.ID = spawnID
	return
}

//ListSpawn will grab data from storage
func (s *Storage) ListSpawn() (spawns []*model.Spawn, err error) {
	query := fmt.Sprintf(`SELECT id, %s FROM %s ORDER BY id DESC LIMIT 50`, spawnFields, spawnTable)
	rows, err := s.db.Queryx(query)
	if err != nil {
		return
	}

	for rows.Next() {
		spawn := model.Spawn{}
		if err = rows.StructScan(&spawn); err != nil {
			return
		}
		spawns = append(spawns, &spawn)
	}
	return
}

//EditSpawn will grab data from storage
func (s *Storage) EditSpawn(spawn *model.Spawn) (err error) {
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", spawnTable, spawnSets)
	result, err := s.db.NamedExec(query, spawn)
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

//DeleteSpawn will grab data from storage
func (s *Storage) DeleteSpawn(spawn *model.Spawn) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", spawnTable)
	result, err := s.db.Exec(query, spawn.ID)
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

//createTableSpawn will grab data from storage
func (s *Storage) createTableSpawn() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE spawngroup (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(50) NOT NULL DEFAULT '',
  spawn_limit tinyint(4) NOT NULL DEFAULT '0',
  dist float NOT NULL DEFAULT '0',
  max_x float NOT NULL DEFAULT '0',
  min_x float NOT NULL DEFAULT '0',
  max_y float NOT NULL DEFAULT '0',
  min_y float NOT NULL DEFAULT '0',
  delay int(11) NOT NULL DEFAULT '45000',
  mindelay int(11) NOT NULL DEFAULT '15000',
  despawn tinyint(3) NOT NULL DEFAULT '0',
  despawn_timer int(11) NOT NULL DEFAULT '100',
  PRIMARY KEY (id),
  UNIQUE KEY name (name)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
