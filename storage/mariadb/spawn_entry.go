package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	spawnEntryTable = "spawn2"

	spawnEntryFields = "id, spawngroupID, zone, version, x, y, z, heading, respawntime, variance, pathgrid, _condition, cond_value, enabled, animation"
	spawnEntryBinds  = ":id, :spawngroupID, :zone, :version, :x, :y, :z, :heading, :respawntime, :variance, :pathgrid, :_condition, :cond_value, :enabled, :animation"
)

//GetSpawnEntry will grab data from storage
func (s *Storage) GetSpawnEntry(spawn *model.Spawn, spawnEntry *model.SpawnEntry) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE spawngroupID = ? AND id = ?", spawnEntryFields, spawnEntryTable)
	err = s.db.Get(spawnEntry, query, spawn.ID, spawnEntry.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateSpawnEntry will grab data from storage
func (s *Storage) CreateSpawnEntry(spawn *model.Spawn, spawnEntry *model.SpawnEntry) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", spawnEntryTable, spawnEntryFields, spawnEntryBinds)
	_, err = s.db.NamedExec(query, spawnEntry)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListSpawnEntry will grab data from storage
func (s *Storage) ListSpawnEntry(page *model.Page, spawn *model.Spawn) (spawnEntrys []*model.SpawnEntry, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE spawngroupID = ? ORDER BY %s LIMIT %d OFFSET %d", spawnEntryFields, spawnEntryTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query, spawn.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		spawnEntry := model.SpawnEntry{}
		if err = rows.StructScan(&spawnEntry); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		spawnEntrys = append(spawnEntrys, &spawnEntry)
	}
	return
}

//ListSpawnEntryTotalCount will grab data from storage
func (s *Storage) ListSpawnEntryTotalCount(spawn *model.Spawn) (count int64, err error) {
	query := fmt.Sprintf("SELECT count(spawngroupID) FROM %s WHERE spawngroupID = ?", spawnEntryTable)
	err = s.db.Get(&count, query, spawn.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListSpawnEntryBySearch will grab data from storage
func (s *Storage) ListSpawnEntryBySearch(page *model.Page, spawn *model.Spawn, spawnEntry *model.SpawnEntry) (spawnEntrys []*model.SpawnEntry, err error) {

	field := ""

	if spawnEntry.ID > 0 {
		field += `id = :id OR`
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]
	spawnEntry.SpawnID = spawn.ID

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", spawnEntryFields, spawnEntryTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, spawnEntry)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		spawnEntry := model.SpawnEntry{}
		if err = rows.StructScan(&spawnEntry); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		spawnEntrys = append(spawnEntrys, &spawnEntry)
	}
	return
}

//ListSpawnEntryBySearchTotalCount will grab data from storage
func (s *Storage) ListSpawnEntryBySearchTotalCount(spawn *model.Spawn, spawnEntry *model.SpawnEntry) (count int64, err error) {
	field := ""
	if spawnEntry.ID > 0 {
		field += `id = :id OR`
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	spawnEntry.SpawnID = spawn.ID
	query := fmt.Sprintf("SELECT count(spawngroupID) FROM %s WHERE %s", spawnEntryTable, field)

	rows, err := s.db.NamedQuery(query, spawnEntry)
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

//EditSpawnEntry will grab data from storage
func (s *Storage) EditSpawnEntry(spawn *model.Spawn, spawnEntry *model.SpawnEntry) (err error) {

	prevSpawnEntry := &model.SpawnEntry{
		SpawnID: spawnEntry.SpawnID,
		ID:      spawnEntry.ID,
	}
	err = s.GetSpawnEntry(spawn, prevSpawnEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous spawnEntry")
		return
	}

	field := ""
	if prevSpawnEntry.ZoneShortName.String != spawnEntry.ZoneShortName.String {
		field += "zone = :zone, "
	}

	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE spawngroupID = :spawngroupID AND id = :id", spawnEntryTable, field)
	result, err := s.db.NamedExec(query, spawnEntry)
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

//DeleteSpawnEntry will grab data from storage
func (s *Storage) DeleteSpawnEntry(spawn *model.Spawn, spawnEntry *model.SpawnEntry) (err error) {

	query := fmt.Sprintf("DELETE FROM %s WHERE spawngroupID = ? AND id = ?", spawnEntryTable)
	result, err := s.db.Exec(query, spawn.ID, spawnEntry.ID)
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
		  ) ENGINE=InnoDB AUTO_INCREMENT=241836 DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
