package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	spawn2Sets       = `zone=:zone, version=:version, x=:x, y=:y, z=:z, heading=:heading, respawntime=:respawntime, variance=:variance, pathgrid=:pathgrid, _condition=:_condition, cond_value=:cond_value, enabled=:enabled, animation=:animation`
	spawn2Fields     = `spawngroupID, zone, version, x, y, z, heading, respawntime, variance, pathgrid, _condition, cond_value, enabled, animation`
	spawn2Binds      = `:spawngroupID, :zone, :version, :x, :y, :z, :heading, :respawntime, :variance, :pathgrid, :_condition, :cond_value, :enabled, :animation`
	spawnGroupSets   = `name=:name, spawn_limit=:spawn_limit, dist=:dist, max_x=:max_x, min_x=:min_x, max_y=:max_y, min_y=:min_y, delay=:delay, mindelay=:mindelay, despawn=:despawn, despawn_timer=:despawn_timer`
	spawnGroupFields = `name, spawn_limit, dist, max_x, min_x, max_y, min_y, delay, mindelay, despawn, despawn_timer`
	spawnGroupBinds  = `:name, :spawn_limit, :dist, :max_x, :min_x, :max_y, :min_y, :delay, :mindelay, :despawn, :despawn_timer`
)

func (s *Storage) GetSpawn(spawnID int64) (spawn *model.Spawn, err error) {
	spawn = &model.Spawn{}
	err = s.db.Get(spawn, fmt.Sprintf(`SELECT spawngroup.id spawngroupID, %s, %s FROM spawn2 
		INNER JOIN spawngroup ON spawngroup.id = spawn2.spawngroupid
		WHERE spawngroup.id = ?`, spawn2Fields, spawnGroupFields), spawnID)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateSpawn(spawn *model.Spawn) (err error) {
	if spawn == nil {
		err = fmt.Errorf("Must provide spawn")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO spawn2(%s)
		VALUES (%s)`, spawn2Fields, spawn2Binds), spawn)
	if err != nil {
		return
	}
	spawnID, err := result.LastInsertId()
	if err != nil {
		return
	}
	spawn.SpawngroupID = spawnID

	result, err = s.db.NamedExec(fmt.Sprintf(`INSERT INTO spawngroup(%s)
		VALUES (%s)`, spawnGroupFields, spawnGroupBinds), spawn)
	if err != nil {
		return
	}
	_, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

func (s *Storage) ListSpawn() (spawns []*model.Spawn, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT spawngroup.id spawngroupID, %s, %s FROM spawn2 
		INNER JOIN spawngroup ON spawngroup.id = spawn2.spawngroupID
		GROUP BY spawn2.spawngroupID`, spawn2Fields, spawnGroupFields))
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

func (s *Storage) EditSpawn(spawnID int64, spawn *model.Spawn) (err error) {
	spawn.SpawngroupID = spawnID
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE spawn2 SET %s WHERE spawn2.spawngroupID = :spawn2.spawngroupID`, spawn2Sets), spawn)
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

	result, err = s.db.NamedExec(fmt.Sprintf(`UPDATE spawngroup SET %s WHERE spawn2.spawngroupID = :spawn2.spawngroupID`, spawnGroupSets), spawn)
	if err != nil {
		return
	}
	affected, err = result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}

func (s *Storage) DeleteSpawn(spawnID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM spawn2 WHERE spawngroupid = ?`, spawnID)
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

	result, err = s.db.Exec(`DELETE FROM spawngroup WHERE id = ?`, spawnID)
	if err != nil {
		return
	}

	affected, err = result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}
