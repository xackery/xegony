package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	spawnTable  = "spawngroup"
	spawnFields = "id, name, spawn_limit, dist, max_x, min_x, max_y, min_y, delay, mindelay, despawn, despawn_timer"
	spawnBinds  = ":id, :name, :spawn_limit, :dist, :max_x, :min_x, :max_y, :min_y, :delay, :mindelay, :despawn, :despawn_timer"
)

//GetSpawn will grab data from storage
func (s *Storage) GetSpawn(spawn *model.Spawn) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", spawnFields, spawnTable)
	err = s.db.Get(spawn, query, spawn.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateSpawn will grab data from storage
func (s *Storage) CreateSpawn(spawn *model.Spawn) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", spawnTable, spawnFields, spawnBinds)
	result, err := s.db.NamedExec(query, spawn)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	spawnID, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	spawn.ID = spawnID
	return
}

//ListSpawn will grab data from storage
func (s *Storage) ListSpawn(page *model.Page) (spawns []*model.Spawn, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", spawnFields, spawnTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		spawn := model.Spawn{}
		if err = rows.StructScan(&spawn); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		spawns = append(spawns, &spawn)
	}
	return
}

//ListSpawnTotalCount will grab data from storage
func (s *Storage) ListSpawnTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(id) FROM %s", spawnTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListSpawnBySearch will grab data from storage
func (s *Storage) ListSpawnBySearch(page *model.Page, spawn *model.Spawn) (spawns []*model.Spawn, err error) {

	field := ""

	if len(spawn.Name) > 0 {
		field += `name LIKE :name OR`
		spawn.Name = fmt.Sprintf("%%%s%%", spawn.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", spawnFields, spawnTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, spawn)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		spawn := model.Spawn{}
		if err = rows.StructScan(&spawn); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		spawns = append(spawns, &spawn)
	}
	return
}

//ListSpawnBySearchTotalCount will grab data from storage
func (s *Storage) ListSpawnBySearchTotalCount(spawn *model.Spawn) (count int64, err error) {
	field := ""
	if len(spawn.Name) > 0 {
		field += `name LIKE :name OR`
		spawn.Name = fmt.Sprintf("%%%s%%", spawn.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", spawnTable, field)

	rows, err := s.db.NamedQuery(query, spawn)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
	}
	return
}

//EditSpawn will grab data from storage
func (s *Storage) EditSpawn(spawn *model.Spawn) (err error) {

	prevSpawn := &model.Spawn{
		ID: spawn.ID,
	}
	err = s.GetSpawn(prevSpawn)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous spawn")
		return
	}

	field := ""
	if len(spawn.Name) > 0 && prevSpawn.Name != spawn.Name {
		field += "name = :name, "
	}
	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", spawnTable, field)
	result, err := s.db.NamedExec(query, spawn)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//DeleteSpawn will grab data from storage
func (s *Storage) DeleteSpawn(spawn *model.Spawn) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", spawnTable)
	result, err := s.db.Exec(query, spawn.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		err = errors.Wrapf(err, "query: %s", query)
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
    ) ENGINE=InnoDB AUTO_INCREMENT=259666 DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
