package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

func (s *Storage) GetNpc(npcId int64) (npc *model.Npc, err error) {
	npc = &model.Npc{}
	err = s.db.Get(npc, "SELECT id, name, level, lastname, hp, class, loottable_id FROM npc_types WHERE id = ?", npcId)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateNpc(npc *model.Npc) (err error) {
	if npc == nil {
		err = fmt.Errorf("Must provide npc")
		return
	}

	result, err := s.db.NamedExec(`INSERT INTO npc_types(name, level, hp, lastname, class, account_id, loottable_id)
		VALUES (:name, :level, :hp, :lastname, :class, :account_id, :loottable_id)`, npc)
	if err != nil {
		return
	}
	npcId, err := result.LastInsertId()
	if err != nil {
		return
	}
	npc.Id = npcId
	return
}

func (s *Storage) ListNpc() (npcs []*model.Npc, err error) {
	rows, err := s.db.Queryx(`SELECT id, name, level, lastname, hp, class, loottable_id FROM npc_types ORDER BY id DESC`)
	if err != nil {
		return
	}

	for rows.Next() {
		npc := model.Npc{}
		if err = rows.StructScan(&npc); err != nil {
			return
		}
		npcs = append(npcs, &npc)
	}
	return
}

func (s *Storage) ListNpcByZone(zoneId int64) (npcs []*model.Npc, err error) {
	rows, err := s.db.Queryx(`SELECT n.id, n.name, n.level, n.lastname, n.hp, n.class, n.loottable_id FROM npc_types n	
	INNER JOIN spawnentry ON spawnentry.npcid = n.id
	INNER JOIN spawn2 ON spawn2.spawngroupid = spawnentry.spawngroupid
	INNER JOIN zone ON zone.short_name = spawn2.zone
	WHERE zone.zoneidnumber = ? ORDER BY id DESC`, zoneId)
	if err != nil {
		return
	}

	for rows.Next() {
		npc := model.Npc{}
		if err = rows.StructScan(&npc); err != nil {
			return
		}
		npcs = append(npcs, &npc)
	}
	return
}

func (s *Storage) ListNpcByFaction(factionId int64) (npcs []*model.Npc, err error) {
	rows, err := s.db.Queryx(`SELECT n.id, n.name, n.level, n.lastname, n.hp, n.class, n.loottable_id FROM npc_types n	
	INNER JOIN npc_faction ON npc_faction.id = n.npc_faction_id
	INNER JOIN faction_list on faction_list.id = npc_faction.primaryfaction
	WHERE faction_list.id = ?`, factionId)
	if err != nil {
		return
	}

	for rows.Next() {
		npc := model.Npc{}
		if err = rows.StructScan(&npc); err != nil {
			return
		}
		npcs = append(npcs, &npc)
	}
	return
}

func (s *Storage) EditNpc(npcId int64, npc *model.Npc) (err error) {
	npc.Id = npcId
	result, err := s.db.NamedExec(`UPDATE npc_types SET level=:level, lastname=:lastname, hp=:hp, class=:class, name=:name, loottable_id=:loottable_id WHERE id = :id`, npc)
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

func (s *Storage) DeleteNpc(npcId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM npc_types WHERE id = ?`, npcId)
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
