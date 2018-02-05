package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	spawnNpcTable  = "spawnentry"
	spawnNpcFields = "spawngroupID, npcID, chance"
	spawnNpcBinds  = ":spawngroupID, :npcID, :chance"
)

//GetSpawnNpc will grab data from storage
func (s *Storage) GetSpawnNpc(spawn *model.Spawn, spawnNpc *model.SpawnNpc) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE spawngroupID = ? AND npcID = ?", spawnNpcFields, spawnNpcTable)
	err = s.db.Get(spawnNpc, query, spawn.ID, spawnNpc.NpcID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateSpawnNpc will grab data from storage
func (s *Storage) CreateSpawnNpc(spawn *model.Spawn, spawnNpc *model.SpawnNpc) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", spawnNpcTable, spawnNpcFields, spawnNpcBinds)
	_, err = s.db.NamedExec(query, spawnNpc)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListSpawnNpc will grab data from storage
func (s *Storage) ListSpawnNpc(page *model.Page, spawn *model.Spawn) (spawnNpcs []*model.SpawnNpc, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "npcID"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE spawngroupID = ? ORDER BY %s LIMIT %d OFFSET %d", spawnNpcFields, spawnNpcTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query, spawn.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		spawnNpc := model.SpawnNpc{}
		if err = rows.StructScan(&spawnNpc); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		spawnNpcs = append(spawnNpcs, &spawnNpc)
	}
	return
}

//ListSpawnNpcTotalCount will grab data from storage
func (s *Storage) ListSpawnNpcTotalCount(spawn *model.Spawn) (count int64, err error) {
	query := fmt.Sprintf("SELECT count(spawngroupID) FROM %s WHERE spawngroupID = ?", spawnNpcTable)
	err = s.db.Get(&count, query, spawn.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListSpawnNpcBySearch will grab data from storage
func (s *Storage) ListSpawnNpcBySearch(page *model.Page, spawn *model.Spawn, spawnNpc *model.SpawnNpc) (spawnNpcs []*model.SpawnNpc, err error) {

	field := ""

	if spawnNpc.NpcID > 0 {
		field += `npcID = :npcID OR`
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]
	spawnNpc.SpawnID = spawn.ID

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", spawnNpcFields, spawnNpcTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, spawnNpc)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		spawnNpc := model.SpawnNpc{}
		if err = rows.StructScan(&spawnNpc); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		spawnNpcs = append(spawnNpcs, &spawnNpc)
	}
	return
}

//ListSpawnNpcBySearchTotalCount will grab data from storage
func (s *Storage) ListSpawnNpcBySearchTotalCount(spawn *model.Spawn, spawnNpc *model.SpawnNpc) (count int64, err error) {
	field := ""
	if spawnNpc.NpcID > 0 {
		field += `npcID = :npcID OR`
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	spawnNpc.SpawnID = spawn.ID
	query := fmt.Sprintf("SELECT count(spawngroupID) FROM %s WHERE %s", spawnNpcTable, field)

	rows, err := s.db.NamedQuery(query, spawnNpc)
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

//EditSpawnNpc will grab data from storage
func (s *Storage) EditSpawnNpc(spawn *model.Spawn, spawnNpc *model.SpawnNpc) (err error) {

	prevSpawnNpc := &model.SpawnNpc{
		SpawnID: spawnNpc.SpawnID,
		NpcID:   spawnNpc.NpcID,
	}
	err = s.GetSpawnNpc(spawn, prevSpawnNpc)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous spawnNpc")
		return
	}

	field := ""
	if prevSpawnNpc.Chance != spawnNpc.Chance {
		field += "chance = :chance, "
	}

	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE spawngroupID = :spawngroupID AND npcID = :npcID", spawnNpcTable, field)
	result, err := s.db.NamedExec(query, spawnNpc)
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

//DeleteSpawnNpc will grab data from storage
func (s *Storage) DeleteSpawnNpc(spawn *model.Spawn, spawnNpc *model.SpawnNpc) (err error) {

	query := fmt.Sprintf("DELETE FROM %s WHERE spawngroupID = ? AND npcID = ?", spawnNpcTable)
	result, err := s.db.Exec(query, spawn.ID, spawnNpc.NpcID)
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
