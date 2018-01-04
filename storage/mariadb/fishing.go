package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	fishingFields = `zoneid, Itemid, skill_level, chance, npc_id, npc_chance`
	fishingSets   = `zoneid=:zoneid, Itemid=:Itemid, skill_level=:skill_level, chance=:chance, npc_id=:npc_id, npc_chance=:npc_chance`
	fishingBinds  = `:zoneid, :Itemid, :skill_level, :chance, :npc_id, :npc_chance`
)

//GetFishing will grab data from storage
func (s *Storage) GetFishing(fishingID int64) (fishing *model.Fishing, err error) {
	fishing = &model.Fishing{}
	err = s.db.Get(fishing, fmt.Sprintf("SELECT id, %s FROM fishing WHERE id = ?", fishingFields), fishingID)
	if err != nil {
		return
	}
	return
}

//CreateFishing will grab data from storage
func (s *Storage) CreateFishing(fishing *model.Fishing) (err error) {
	if fishing == nil {
		err = fmt.Errorf("Must provide fishing")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO fishing(%s)
		VALUES (%s)`, fishingFields, fishingBinds), fishing)
	if err != nil {
		return
	}
	fishingID, err := result.LastInsertId()
	if err != nil {
		return
	}
	fishing.ID = fishingID
	return
}

//ListFishing will grab data from storage
func (s *Storage) ListFishing(pageSize int64, pageNumber int64) (fishings []*model.Fishing, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM fishing
		ORDER BY id ASC LIMIT %d OFFSET %d`, fishingFields, pageSize, pageSize*pageNumber))
	if err != nil {
		return
	}

	for rows.Next() {
		fishing := model.Fishing{}
		if err = rows.StructScan(&fishing); err != nil {
			return
		}
		fishings = append(fishings, &fishing)
	}
	return
}

//ListFishingCount will grab data from storage
func (s *Storage) ListFishingCount() (count int64, err error) {
	err = s.db.Get(&count, `SELECT count(id) FROM fishing`)
	if err != nil {
		return
	}
	return
}

//ListFishingByItem will grab data from storage
func (s *Storage) ListFishingByItem(itemID int64) (fishings []*model.Fishing, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT fishing.id, %s FROM fishing		
		WHERE fishing.itemid = ?`, fishingFields), itemID)
	if err != nil {
		return
	}

	for rows.Next() {
		fishing := model.Fishing{}
		if err = rows.StructScan(&fishing); err != nil {
			return
		}
		fishings = append(fishings, &fishing)
	}
	return
}

//ListFishingByNpc will grab data from storage
func (s *Storage) ListFishingByNpc(npcID int64) (fishings []*model.Fishing, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT fishing.id, %s FROM fishing		
		WHERE fishing.npcid = ?`, fishingFields), npcID)
	if err != nil {
		return
	}

	for rows.Next() {
		fishing := model.Fishing{}
		if err = rows.StructScan(&fishing); err != nil {
			return
		}
		fishings = append(fishings, &fishing)
	}
	return
}

//ListFishingByZone will grab data from storage
func (s *Storage) ListFishingByZone(zoneID int64) (fishings []*model.Fishing, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT fishing.id, %s FROM fishing		
		WHERE fishing.zoneid = ?`, fishingFields), zoneID)
	if err != nil {
		return
	}

	for rows.Next() {
		fishing := model.Fishing{}
		if err = rows.StructScan(&fishing); err != nil {
			return
		}
		fishings = append(fishings, &fishing)
	}
	return
}

//EditFishing will grab data from storage
func (s *Storage) EditFishing(fishingID int64, fishing *model.Fishing) (err error) {
	fishing.ID = fishingID
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE fishing SET %s WHERE id = :id`, fishingSets), fishing)
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

//DeleteFishing will grab data from storage
func (s *Storage) DeleteFishing(fishingID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM fishing WHERE id = ?`, fishingID)
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
