package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	npcSets   = `:id, name=:name, lastname=:lastname, level=:level, race=:race, class=:class, bodytype=:bodytype, hp=:hp, mana=:mana, gender=:gender, texture=:texture, helmtexture=:helmtexture, herosforgemodel=:herosforgemodel, size=:size, hp_regen_rate=:hp_regen_rate, mana_regen_rate=:mana_regen_rate, loottable_id=:loottable_id, merchant_id=:merchant_id, alt_currency_id=:alt_currency_id, npc_spells_id=:npc_spells_id, npc_spells_effects_id=:npc_spells_effects_id, npc_faction_id=:npc_faction_id, adventure_template_id=:adventure_template_id, trap_template=:trap_template, mindmg=:mindmg, maxdmg=:maxdmg, attack_count=:attack_count, npcspecialattks=:npcspecialattks, special_abilities=:special_abilities, aggroradius=:aggroradius, assistradius=:assistradius, face=:face, luclin_hairstyle=:luclin_hairstyle, luclin_haircolor=:luclin_haircolor, luclin_eyecolor=:luclin_eyecolor, luclin_eyecolor2=:luclin_eyecolor2, luclin_beardcolor=:luclin_beardcolor, luclin_beard=:luclin_beard, drakkin_heritage=:drakkin_heritage, drakkin_tattoo=:drakkin_tattoo, drakkin_details=:drakkin_details, armortint_id=:armortint_id, armortint_red=:armortint_red, armortint_green=:armortint_green, armortint_blue=:armortint_blue, d_melee_texture1=:d_melee_texture1, d_melee_texture2=:d_melee_texture2, ammo_idfile=:ammo_idfile, prim_melee_type=:prim_melee_type, sec_melee_type=:sec_melee_type, ranged_type=:ranged_type, runspeed=:runspeed, MR=:MR, CR=:CR, DR=:DR, FR=:FR, PR=:PR, Corrup=:Corrup, PhR=:PhR, see_invis=:see_invis, see_invis_undead=:see_invis_undead, qglobal=:qglobal, AC=:AC, npc_aggro=:npc_aggro, spawn_limit=:spawn_limit, attack_speed=:attack_speed, attack_delay=:attack_delay, findable=:findable, STR=:STR, STA=:STA, DEX=:DEX, AGI=:AGI, _INT=:_INT, WIS=:WIS, CHA=:CHA, see_hide=:see_hide, see_improved_hide=:see_improved_hide, trackable=:trackable, isbot=:isbot, exclude=:exclude, ATK=:ATK, Accuracy=:Accuracy, Avoidance=:Avoidance, slow_mitigation=:slow_mitigation, version=:version, maxlevel=:maxlevel, scalerate=:scalerate, private_corpse=:private_corpse, unique_spawn_by_name=:unique_spawn_by_name, underwater=:underwater, isquest=:isquest, emoteid=:emoteid, spellscale=:spellscale, healscale=:healscale, no_target_hotkey=:no_target_hotkey, raid_target=:raid_target, armtexture=:armtexture, bracertexture=:bracertexture, handtexture=:handtexture, legtexture=:legtexture, feettexture=:feettexture, light=:light, walkspeed=:walkspeed, peqid=:peqid, unique_=:unique_, fixed=:fixed, ignore_despawn=:ignore_despawn, show_name=:show_name, untargetable=:untargetable`
	npcFields = ` npc_types.name, npc_types.lastname, npc_types.level, npc_types.race, npc_types.class, npc_types.bodytype, npc_types.hp, npc_types.mana, npc_types.gender, npc_types.texture, npc_types.helmtexture, npc_types.herosforgemodel, npc_types.size, npc_types.hp_regen_rate, npc_types.mana_regen_rate, npc_types.loottable_id, npc_types.merchant_id, npc_types.alt_currency_id, npc_types.npc_spells_id, npc_types.npc_spells_effects_id, npc_types.npc_faction_id, npc_types.adventure_template_id, npc_types.trap_template, npc_types.mindmg, npc_types.maxdmg, npc_types.attack_count, npc_types.npcspecialattks, npc_types.special_abilities, npc_types.aggroradius, npc_types.assistradius, npc_types.face, npc_types.luclin_hairstyle, npc_types.luclin_haircolor, npc_types.luclin_eyecolor, npc_types.luclin_eyecolor2, npc_types.luclin_beardcolor, npc_types.luclin_beard, npc_types.drakkin_heritage, npc_types.drakkin_tattoo, npc_types.drakkin_details, npc_types.armortint_id, npc_types.armortint_red, npc_types.armortint_green, npc_types.armortint_blue, npc_types.d_melee_texture1, npc_types.d_melee_texture2, npc_types.ammo_idfile, npc_types.prim_melee_type, npc_types.sec_melee_type, npc_types.ranged_type, npc_types.runspeed, npc_types.MR, npc_types.CR, npc_types.DR, npc_types.FR, npc_types.PR, npc_types.Corrup, npc_types.PhR, npc_types.see_invis, npc_types.see_invis_undead, npc_types.qglobal, npc_types.AC, npc_types.npc_aggro, npc_types.spawn_limit, npc_types.attack_speed, npc_types.attack_delay, npc_types.findable, npc_types.STR, npc_types.STA, npc_types.DEX, npc_types.AGI, npc_types._INT, npc_types.WIS, npc_types.CHA, npc_types.see_hide, npc_types.see_improved_hide, npc_types.trackable, npc_types.isbot, npc_types.exclude, npc_types.ATK, npc_types.Accuracy, npc_types.Avoidance, npc_types.slow_mitigation, npc_types.version, npc_types.maxlevel, npc_types.scalerate, npc_types.private_corpse, npc_types.unique_spawn_by_name, npc_types.underwater, npc_types.isquest, npc_types.emoteid, npc_types.spellscale, npc_types.healscale, npc_types.no_target_hotkey, npc_types.raid_target, npc_types.armtexture, npc_types.bracertexture, npc_types.handtexture, npc_types.legtexture, npc_types.feettexture, npc_types.light, npc_types.walkspeed, npc_types.peqid, npc_types.unique_, npc_types.fixed, npc_types.ignore_despawn, npc_types.show_name, untargetable`
	npcBinds  = ` :name, :lastname, :level, :race, :class, :bodytype, :hp, :mana, :gender, :texture, :helmtexture, :herosforgemodel, :size, :hp_regen_rate, :mana_regen_rate, :loottable_id, :merchant_id, :alt_currency_id, :npc_spells_id, :npc_spells_effects_id, :npc_faction_id, :adventure_template_id, :trap_template, :mindmg, :maxdmg, :attack_count, :npcspecialattks, :special_abilities, :aggroradius, :assistradius, :face, :luclin_hairstyle, :luclin_haircolor, :luclin_eyecolor, :luclin_eyecolor2, :luclin_beardcolor, :luclin_beard, :drakkin_heritage, :drakkin_tattoo, :drakkin_details, :armortint_id, :armortint_red, :armortint_green, :armortint_blue, :d_melee_texture1, :d_melee_texture2, :ammo_idfile, :prim_melee_type, :sec_melee_type, :ranged_type, :runspeed, :MR, :CR, :DR, :FR, :PR, :Corrup, :PhR, :see_invis, :see_invis_undead, :qglobal, :AC, :npc_aggro, :spawn_limit, :attack_speed, :attack_delay, :findable, :STR, :STA, :DEX, :AGI, :_INT, :WIS, :CHA, :see_hide, :see_improved_hide, :trackable, :isbot, :exclude, :ATK, :Accuracy, :Avoidance, :slow_mitigation, :version, :maxlevel, :scalerate, :private_corpse, :unique_spawn_by_name, :underwater, :isquest, :emoteid, :spellscale, :healscale, :no_target_hotkey, :raid_target, :armtexture, :bracertexture, :handtexture, :legtexture, :feettexture, :light, :walkspeed, :peqid, :unique_, :fixed, :ignore_despawn, :show_name, :untargetable`
)

//GetNpc will grab data from storage
func (s *Storage) GetNpc(npc *model.Npc) (err error) {
	err = s.db.Get(npc, fmt.Sprintf("SELECT npc_types.id, %s FROM npc_types WHERE id = ?", npcFields), npc.ID)
	if err != nil {
		return
	}
	return
}

//CreateNpc will grab data from storage
func (s *Storage) CreateNpc(npc *model.Npc) (err error) {
	if npc == nil {
		err = fmt.Errorf("Must provide npc")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO npc_types(%s)
		VALUES (%s)`, npcFields, npcBinds), npc)
	if err != nil {
		return
	}
	npcID, err := result.LastInsertId()
	if err != nil {
		return
	}
	npc.ID = npcID
	return
}

//ListNpc will grab data from storage
func (s *Storage) ListNpc(pageSize int64, pageNumber int64) (npcs []*model.Npc, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM npc_types ORDER BY id ASC LIMIT %d OFFSET %d`, npcFields, pageSize, pageSize*pageNumber))
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

//ListNpcCount will grab data from storage
func (s *Storage) ListNpcCount() (count int64, err error) {
	err = s.db.Get(&count, `SELECT count(id) FROM npc_types`)
	if err != nil {
		return
	}
	return
}

//ListNpcByZone will grab data from storage
func (s *Storage) ListNpcByZone(zone *model.Zone) (npcs []*model.Npc, err error) {

	upperID := (zone.ID * 1000) + 1000 - 1
	lowerID := (zone.ID * 1000) - 1
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT npc_types.id, %s FROM npc_types
	WHERE npc_types.id < ? AND npc_types.id > ?`, npcFields), upperID, lowerID)
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

//ListNpcByFaction will grab data from storage
func (s *Storage) ListNpcByFaction(faction *model.Faction) (npcs []*model.Npc, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT npc_types.id, %s FROM npc_types	
	INNER JOIN npc_faction ON npc_faction.id = npc_types.npc_faction_id
	INNER JOIN faction_list on faction_list.id = npc_faction.primaryfaction
	WHERE faction_list.id = ?`, npcFields), faction.ID)
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

//ListNpcByLootTable will grab data from storage
func (s *Storage) ListNpcByLootTable(lootTable *model.LootTable) (npcs []*model.Npc, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT npc_types.id, %s FROM npc_types
		WHERE loottable_id = ?`, npcFields), lootTable.ID)
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

//ListNpcByMerchant will grab data from storage
func (s *Storage) ListNpcByMerchant(merchant *model.Merchant) (npcs []*model.Npc, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT npc_types.id, %s FROM npc_types
		WHERE merchant_id = ?`, npcFields), merchant.ID)
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

//ListNpcByItem will grab data from storage
func (s *Storage) ListNpcByItem(item *model.Item) (npcs []*model.Npc, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT npc_types.id, %s FROM npc_types
		INNER JOIN npc_loot_cache ON npc_loot_cache.npc_id = npc_types.id
		WHERE npc_loot_cache.item_id = ?`, npcFields), item.ID)
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

//ListNpcBySpell will grab data from storage
func (s *Storage) ListNpcBySpell(spell *model.Spell) (npcs []*model.Npc, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT npc_types.id, %s FROM npc_types
		INNER JOIN npc_spells_entries ON npc_spells_entries.id = npc_types.npc_spells_id
		WHERE npc_spells_entries.spellid = ?`, npcFields), spell.ID)
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

//EditNpc will grab data from storage
func (s *Storage) EditNpc(npc *model.Npc) (err error) {
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE npc_types SET %s WHERE id = :id`, npcSets), npc)
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

//DeleteNpc will grab data from storage
func (s *Storage) DeleteNpc(npc *model.Npc) (err error) {
	result, err := s.db.Exec(`DELETE FROM npc_types WHERE id = ?`, npc.ID)
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

//SearchNpcByName will grab data from storage
func (s *Storage) SearchNpcByName(npc *model.Npc) (npcs []*model.Npc, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM npc_types WHERE name like ? ORDER BY id DESC`, npcFields), "%"+npc.Name+"%")
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
