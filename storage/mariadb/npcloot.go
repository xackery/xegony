package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	npcLootSets   = `npc_id=:npc_id, item_id=:item_id`
	npcLootFields = `npc_loot_cache.npc_id, npc_loot_cache.item_id`
	npcLootBinds  = `:npc_id, :item_id`
)

//GetNpcLoot will grab data from storage
func (s *Storage) GetNpcLoot(npcLoot *model.NpcLoot) (err error) {
	err = s.db.Get(npcLoot, fmt.Sprintf(`SELECT %s, %s FROM npc_loot_cache
	WHERE npc_loot_cache.npc_id = ? AND npc_loot_cache.item_id = ?`, npcLootFields, itemFields), npcLoot.NpcID, npcLoot.ItemID)
	if err != nil {
		return
	}
	return
}

//CreateNpcLoot will grab data from storage
func (s *Storage) CreateNpcLoot(npcLoot *model.NpcLoot) (err error) {
	if npcLoot == nil {
		err = fmt.Errorf("Must provide npcLoot")
		return
	}

	_, err = s.db.NamedExec(fmt.Sprintf(`INSERT INTO npc_loot_cache(%s)
		VALUES (%s)`, npcLootFields, npcLootBinds), npcLoot)
	if err != nil {
		return
	}
	return
}

//ListNpcLoot will grab data from storage
func (s *Storage) ListNpcLootByNpc(npc *model.Npc) (npcLoots []*model.NpcLoot, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s, %s FROM npc_loot_cache
	WHERE npc_loot_cache.npc_id = ?`, npcLootFields, itemFields), npc.ID)
	if err != nil {
		return
	}

	for rows.Next() {
		npcLoot := model.NpcLoot{}
		if err = rows.StructScan(&npcLoot); err != nil {
			return
		}
		npcLoots = append(npcLoots, &npcLoot)
	}
	return
}

//ListNpcLootByZone will grab data from storage
func (s *Storage) ListNpcLootByZone(zone *model.Zone) (npcLoots []*model.NpcLoot, err error) {
	upperID := (zone.ID * 1000) + 1000 - 1
	lowerID := (zone.ID * 1000) - 1

	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM npc_loot_cache
	WHERE npc_loot_cache.npc_id < ? AND npc_loot_cache.npc_id > ? GROUP BY npc_loot_cache.item_id ORDER BY npc_loot_cache.npc_id ASC`, npcLootFields), upperID, lowerID)
	if err != nil {
		return
	}

	for rows.Next() {
		npcLoot := model.NpcLoot{}
		if err = rows.StructScan(&npcLoot); err != nil {
			return
		}
		npcLoots = append(npcLoots, &npcLoot)
	}
	return
}

//EditNpcLoot will grab data from storage
func (s *Storage) EditNpcLoot(npcLoot *model.NpcLoot) (err error) {
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE npc_loot_cache SET %s WHERE npc_id = :npc_id, item_id = :item_id`, npcLootSets), npcLoot)
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

//TruncateNpcLoot will grab data from storage
func (s *Storage) TruncateNpcLoot() (err error) {
	_, err = s.db.Exec(`TRUNCATE npc_loot_cache`)
	if err != nil {
		return
	}
	return
}

//DeleteNpcLoot will grab data from storage
func (s *Storage) DeleteNpcLoot(npcLoot *model.NpcLoot) (err error) {
	result, err := s.db.Exec(`DELETE FROM npc_loot_cache WHERE npc_id = ? AND item_id = ?`, npcLoot.NpcID, npcLoot.ItemID)
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

//createTableNpcLoot will grab data from storage
func (s *Storage) createTableNpcLoot() (err error) {
	_, err = s.db.Exec(`CREATE TABLE npc_loot_cache (
  npc_id int(11) unsigned NOT NULL,
  item_id int(10) unsigned NOT NULL,
  UNIQUE KEY item_id (item_id,npc_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	return
}
