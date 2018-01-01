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
func (s *Storage) GetNpcLoot(npcID int64, itemID int64) (npcLoot *model.NpcLoot, err error) {
	npcLoot = &model.NpcLoot{}
	err = s.db.Get(npcLoot, fmt.Sprintf(`SELECT %s, %s FROM npc_loot_cache
	INNER JOIN items ON items.id = npc_loot_cache.item_id 
	WHERE npc_loot_cache.npc_id = ? AND npc_loot_cache.item_id = ?`, npcLootFields, itemFields), npcID, itemID)
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
func (s *Storage) ListNpcLoot(npcID int64) (npcLoots []*model.NpcLoot, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s, %s FROM npc_loot_cache
	INNER JOIN items ON items.id = npc_loot_cache.item_id 
	WHERE npc_loot_cache.npc_id = ?`, npcLootFields, itemFields), npcID)
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
func (s *Storage) ListNpcLootByZone(zoneID int64) (npcLoots []*model.NpcLoot, err error) {
	upperID := (zoneID * 1000) + 1000 - 1
	lowerID := (zoneID * 1000) - 1

	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT npc_types.name npc_name, %s, %s, %s FROM npc_loot_cache
	INNER JOIN items ON items.id = npc_loot_cache.item_id 
	INNER JOIN npc_types ON npc_types.id = npc_loot_cache.npc_id
	WHERE npc_loot_cache.npc_id < ? AND npc_loot_cache.npc_id > ? GROUP BY npc_loot_cache.item_id ORDER BY npc_loot_cache.npc_id ASC`, npcFields, npcLootFields, itemFields), upperID, lowerID)
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
func (s *Storage) EditNpcLoot(npcID int64, itemID int64, npcLoot *model.NpcLoot) (err error) {
	npcLoot.NpcID = npcID
	npcLoot.ItemID = itemID
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
func (s *Storage) DeleteNpcLoot(npcID int64, itemID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM npc_loot_cache WHERE npc_id = ? AND item_id = ?`, npcID, itemID)
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
