package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	npcTable  = "npc_types"
	npcFields = "id,name,lastname,level,race,class,bodytype,hp,mana,gender,texture,helmtexture,herosforgemodel,size,hp_regen_rate,mana_regen_rate,loottable_id,merchant_id,alt_currency_id,npc_spells_id,npc_spells_effects_id,npc_faction_id,adventure_template_id,trap_template,mindmg,maxdmg,attack_count,npcspecialattks,special_abilities,aggroradius,assistradius,face,luclin_hairstyle,luclin_haircolor,luclin_eyecolor,luclin_eyecolor2,luclin_beardcolor,luclin_beard,drakkin_heritage,drakkin_tattoo,drakkin_details,armortint_id,armortint_red,armortint_green,armortint_blue,d_melee_texture1,d_melee_texture2,ammo_idfile,prim_melee_type,sec_melee_type,ranged_type,runspeed,MR,CR,DR,FR,PR,Corrup,PhR,see_invis,see_invis_undead,qglobal,AC,npc_aggro,spawn_limit,attack_speed,attack_delay,findable,STR,STA,DEX,AGI,_INT,WIS,CHA,see_hide,see_improved_hide,trackable,isbot,exclude,ATK,Accuracy,Avoidance,slow_mitigation,version,maxlevel,scalerate,private_corpse,unique_spawn_by_name,underwater,isquest,emoteid,spellscale,healscale,no_target_hotkey,raid_target,armtexture,bracertexture,handtexture,legtexture,feettexture,light,walkspeed,peqid,unique_,fixed,ignore_despawn,show_name,untargetable"
	npcBinds  = ":id, :name, :lastname, :level, :race, :class, :bodytype, :hp, :mana, :gender, :texture, :helmtexture, :herosforgemodel, :size, :hp_regen_rate, :mana_regen_rate, :loottable_id, :merchant_id, :alt_currency_id, :npc_spells_id, :npc_spells_effects_id, :npc_faction_id, :adventure_template_id, :trap_template, :mindmg, :maxdmg, :attack_count, :npcspecialattks, :special_abilities, :aggroradius, :assistradius, :face, :luclin_hairstyle, :luclin_haircolor, :luclin_eyecolor, :luclin_eyecolor2, :luclin_beardcolor, :luclin_beard, :drakkin_heritage, :drakkin_tattoo, :drakkin_details, :armortint_id, :armortint_red, :armortint_green, :armortint_blue, :d_melee_texture1, :d_melee_texture2, :ammo_idfile, :prim_melee_type, :sec_melee_type, :ranged_type, :runspeed, :MR, :CR, :DR, :FR, :PR, :Corrup, :PhR, :see_invis, :see_invis_undead, :qglobal, :AC, :npc_aggro, :spawn_limit, :attack_speed, :attack_delay, :findable, :STR, :STA, :DEX, :AGI, :_INT, :WIS, :CHA, :see_hide, :see_improved_hide, :trackable, :isbot, :exclude, :ATK, :Accuracy, :Avoidance, :slow_mitigation, :version, :maxlevel, :scalerate, :private_corpse, :unique_spawn_by_name, :underwater, :isquest, :emoteid, :spellscale, :healscale, :no_target_hotkey, :raid_target, :armtexture, :bracertexture, :handtexture, :legtexture, :feettexture, :light, :walkspeed, :peqid, :unique_, :fixed, :ignore_despawn, :show_name, untargetable"
)

//GetNpc will grab data from storage
func (s *Storage) GetNpc(npc *model.Npc) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", npcFields, npcTable)
	err = s.db.Get(npc, query, npc.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateNpc will grab data from storage
func (s *Storage) CreateNpc(npc *model.Npc) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", npcTable, npcFields, npcBinds)
	result, err := s.db.NamedExec(query, npc)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	npcID, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	npc.ID = npcID
	return
}

//ListNpc will grab data from storage
func (s *Storage) ListNpc(page *model.Page) (npcs []*model.Npc, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", npcFields, npcTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		npc := model.Npc{}
		if err = rows.StructScan(&npc); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		npcs = append(npcs, &npc)
	}
	return
}

//ListNpcTotalCount will grab data from storage
func (s *Storage) ListNpcTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(id) FROM %s", npcTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListNpcBySearch will grab data from storage
func (s *Storage) ListNpcBySearch(page *model.Page, npc *model.Npc) (npcs []*model.Npc, err error) {

	field := ""

	if len(npc.Name) > 0 {
		field += `name LIKE :name OR`
		npc.Name = fmt.Sprintf("%%%s%%", npc.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", npcFields, npcTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, npc)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		npc := model.Npc{}
		if err = rows.StructScan(&npc); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		npcs = append(npcs, &npc)
	}
	return
}

//ListNpcBySearchTotalCount will grab data from storage
func (s *Storage) ListNpcBySearchTotalCount(npc *model.Npc) (count int64, err error) {
	field := ""
	if len(npc.Name) > 0 {
		field += `name LIKE :name OR`
		npc.Name = fmt.Sprintf("%%%s%%", npc.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", npcTable, field)

	rows, err := s.db.NamedQuery(query, npc)
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

//ListNpcByZone will grab data from storage
func (s *Storage) ListNpcByZone(page *model.Page, zone *model.Zone) (npcs []*model.Npc, err error) {

	field := ""

	if zone.ID > 0 {
		field += fmt.Sprintf(`id < %d AND id > %d OR`, (zone.ID*1000)+1000-1, (zone.ID*1000)-1)
	}
	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", npcFields, npcTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, zone)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		npc := model.Npc{}
		if err = rows.StructScan(&npc); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		npcs = append(npcs, &npc)
	}
	return
}

//ListNpcByZoneTotalCount will grab data from storage
func (s *Storage) ListNpcByZoneTotalCount(zone *model.Zone) (count int64, err error) {
	field := ""
	if zone.ID > 0 {
		field += fmt.Sprintf(`id < %d AND id > %d OR`, (zone.ID*1000)+1000-1, (zone.ID*1000)-1)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", npcTable, field)

	rows, err := s.db.NamedQuery(query, zone)
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

//EditNpc will grab data from storage
func (s *Storage) EditNpc(npc *model.Npc) (err error) {

	prevNpc := &model.Npc{
		ID: npc.ID,
	}
	err = s.GetNpc(prevNpc)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous npc")
		return
	}

	field := ""
	if len(npc.Name) > 0 && prevNpc.Name != npc.Name {
		field += "name = :name, "
	}
	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", npcTable, field)
	result, err := s.db.NamedExec(query, npc)
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

//DeleteNpc will grab data from storage
func (s *Storage) DeleteNpc(npc *model.Npc) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", npcTable)
	result, err := s.db.Exec(query, npc.ID)
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

func (s *Storage) insertTestNpc() (err error) {
	_, err = s.db.Exec("INSERT INTO `npc_types` (`id`, `name`, `lastname`, `level`, `race`, `class`, `bodytype`, `hp`, `mana`, `gender`, `texture`, `helmtexture`, `herosforgemodel`, `size`, `hp_regen_rate`, `mana_regen_rate`, `loottable_id`, `merchant_id`, `alt_currency_id`, `npc_spells_id`, `npc_spells_effects_id`, `npc_faction_id`, `adventure_template_id`, `trap_template`, `mindmg`, `maxdmg`, `attack_count`, `npcspecialattks`, `special_abilities`, `aggroradius`, `assistradius`, `face`, `luclin_hairstyle`, `luclin_haircolor`, `luclin_eyecolor`, `luclin_eyecolor2`, `luclin_beardcolor`, `luclin_beard`, `drakkin_heritage`, `drakkin_tattoo`, `drakkin_details`, `armortint_id`, `armortint_red`, `armortint_green`, `armortint_blue`, `d_melee_texture1`, `d_melee_texture2`, `ammo_idfile`, `prim_melee_type`, `sec_melee_type`, `ranged_type`, `runspeed`, `MR`, `CR`, `DR`, `FR`, `PR`, `Corrup`, `PhR`, `see_invis`, `see_invis_undead`, `qglobal`, `AC`, `npc_aggro`, `spawn_limit`, `attack_speed`, `attack_delay`, `findable`, `STR`, `STA`, `DEX`, `AGI`, `_INT`, `WIS`, `CHA`, `see_hide`, `see_improved_hide`, `trackable`, `isbot`, `exclude`, `ATK`, `Accuracy`, `Avoidance`, `slow_mitigation`, `version`, `maxlevel`, `scalerate`, `private_corpse`, `unique_spawn_by_name`, `underwater`, `isquest`, `emoteid`, `spellscale`, `healscale`, `no_target_hotkey`, `raid_target`, `armtexture`, `bracertexture`, `handtexture`, `legtexture`, `feettexture`, `light`, `walkspeed`, `peqid`, `unique_`, `fixed`, `ignore_despawn`) VALUES (1008, 'Topala_Xenem', 'Bard Songs', 45, 71, 41, 1, 5875, 0, 1, 1, 1, 0, 6, 12, 12, 0, 1008, 0, 0, 0, 144, 0, 0, 36, 139, -1, '', '', 55, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 'IT10', 28, 28, 7, 1.325, 18, 18, 18, 18, 18, 28, 10, 0, 1, 0, 311, 0, 0, -10, 32, 1, 156, 156, 156, 156, 156, 156, 156, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 100, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0);")
	if err != nil {
		err = errors.Wrap(err, "failed to insert npc data")
		return
	}
	return
}

//createTableNpc will grab data from storage
func (s *Storage) createTableNpc() (err error) {
	_, err = s.db.Exec(`
		CREATE TABLE npc_types (
			id int(11) NOT NULL AUTO_INCREMENT,
			name text NOT NULL,
			lastname varchar(32) DEFAULT NULL,
			level tinyint(2) unsigned NOT NULL DEFAULT '0',
			race smallint(5) unsigned NOT NULL DEFAULT '0',
			class tinyint(2) unsigned NOT NULL DEFAULT '0',
			bodytype int(11) NOT NULL DEFAULT '1',
			hp int(11) NOT NULL DEFAULT '0',
			mana int(11) NOT NULL DEFAULT '0',
			gender tinyint(2) unsigned NOT NULL DEFAULT '0',
			texture tinyint(2) unsigned NOT NULL DEFAULT '0',
			helmtexture tinyint(2) unsigned NOT NULL DEFAULT '0',
			herosforgemodel int(11) NOT NULL DEFAULT '0',
			size float NOT NULL DEFAULT '0',
			hp_regen_rate int(11) unsigned NOT NULL DEFAULT '0',
			mana_regen_rate int(11) unsigned NOT NULL DEFAULT '0',
			loottable_id int(11) unsigned NOT NULL DEFAULT '0',
			merchant_id int(11) unsigned NOT NULL DEFAULT '0',
			alt_currency_id int(11) unsigned NOT NULL DEFAULT '0',
			npc_spells_id int(11) unsigned NOT NULL DEFAULT '0',
			npc_spells_effects_id int(11) unsigned NOT NULL DEFAULT '0',
			npc_faction_id int(11) NOT NULL DEFAULT '0',
			adventure_template_id int(10) unsigned NOT NULL DEFAULT '0',
			trap_template int(10) unsigned DEFAULT '0',
			mindmg int(10) unsigned NOT NULL DEFAULT '0',
			maxdmg int(10) unsigned NOT NULL DEFAULT '0',
			attack_count smallint(6) NOT NULL DEFAULT '-1',
			npcspecialattks varchar(36) NOT NULL DEFAULT '',
			special_abilities text,
			aggroradius int(10) unsigned NOT NULL DEFAULT '0',
			assistradius int(10) unsigned NOT NULL DEFAULT '0',
			face int(10) unsigned NOT NULL DEFAULT '1',
			luclin_hairstyle int(10) unsigned NOT NULL DEFAULT '1',
			luclin_haircolor int(10) unsigned NOT NULL DEFAULT '1',
			luclin_eyecolor int(10) unsigned NOT NULL DEFAULT '1',
			luclin_eyecolor2 int(10) unsigned NOT NULL DEFAULT '1',
			luclin_beardcolor int(10) unsigned NOT NULL DEFAULT '1',
			luclin_beard int(10) unsigned NOT NULL DEFAULT '0',
			drakkin_heritage int(10) NOT NULL DEFAULT '0',
			drakkin_tattoo int(10) NOT NULL DEFAULT '0',
			drakkin_details int(10) NOT NULL DEFAULT '0',
			armortint_id int(10) unsigned NOT NULL DEFAULT '0',
			armortint_red tinyint(3) unsigned NOT NULL DEFAULT '0',
			armortint_green tinyint(3) unsigned NOT NULL DEFAULT '0',
			armortint_blue tinyint(3) unsigned NOT NULL DEFAULT '0',
			d_melee_texture1 int(11) NOT NULL DEFAULT '0',
			d_melee_texture2 int(11) NOT NULL DEFAULT '0',
			ammo_idfile varchar(30) NOT NULL DEFAULT 'IT10',
			prim_melee_type tinyint(4) unsigned NOT NULL DEFAULT '28',
			sec_melee_type tinyint(4) unsigned NOT NULL DEFAULT '28',
			ranged_type tinyint(4) unsigned NOT NULL DEFAULT '7',
			runspeed float NOT NULL DEFAULT '0',
			MR smallint(5) NOT NULL DEFAULT '0',
			CR smallint(5) NOT NULL DEFAULT '0',
			DR smallint(5) NOT NULL DEFAULT '0',
			FR smallint(5) NOT NULL DEFAULT '0',
			PR smallint(5) NOT NULL DEFAULT '0',
			Corrup smallint(5) NOT NULL DEFAULT '0',
			PhR smallint(5) unsigned NOT NULL DEFAULT '0',
			see_invis smallint(4) NOT NULL DEFAULT '0',
			see_invis_undead smallint(4) NOT NULL DEFAULT '0',
			qglobal int(2) unsigned NOT NULL DEFAULT '0',
			AC smallint(5) NOT NULL DEFAULT '0',
			npc_aggro tinyint(4) NOT NULL DEFAULT '0',
			spawn_limit tinyint(4) NOT NULL DEFAULT '0',
			attack_speed float NOT NULL DEFAULT '0',
			attack_delay tinyint(3) unsigned NOT NULL DEFAULT '30',
			findable tinyint(4) NOT NULL DEFAULT '0',
			STR mediumint(8) unsigned NOT NULL DEFAULT '75',
			STA mediumint(8) unsigned NOT NULL DEFAULT '75',
			DEX mediumint(8) unsigned NOT NULL DEFAULT '75',
			AGI mediumint(8) unsigned NOT NULL DEFAULT '75',
			_INT mediumint(8) unsigned NOT NULL DEFAULT '80',
			WIS mediumint(8) unsigned NOT NULL DEFAULT '75',
			CHA mediumint(8) unsigned NOT NULL DEFAULT '75',
			see_hide tinyint(4) NOT NULL DEFAULT '0',
			see_improved_hide tinyint(4) NOT NULL DEFAULT '0',
			trackable tinyint(4) NOT NULL DEFAULT '1',
			isbot tinyint(4) NOT NULL DEFAULT '0',
			exclude tinyint(4) NOT NULL DEFAULT '1',
			ATK mediumint(9) NOT NULL DEFAULT '0',
			Accuracy mediumint(9) NOT NULL DEFAULT '0',
			Avoidance mediumint(9) unsigned NOT NULL DEFAULT '0',
			slow_mitigation smallint(4) NOT NULL DEFAULT '0',
			version smallint(5) unsigned NOT NULL DEFAULT '0',
			maxlevel tinyint(3) NOT NULL DEFAULT '0',
			scalerate int(11) NOT NULL DEFAULT '100',
			private_corpse tinyint(3) unsigned NOT NULL DEFAULT '0',
			unique_spawn_by_name tinyint(3) unsigned NOT NULL DEFAULT '0',
			underwater tinyint(3) unsigned NOT NULL DEFAULT '0',
			isquest tinyint(3) NOT NULL DEFAULT '0',
			emoteid int(10) unsigned NOT NULL DEFAULT '0',
			spellscale float NOT NULL DEFAULT '100',
			healscale float NOT NULL DEFAULT '100',
			no_target_hotkey tinyint(1) unsigned NOT NULL DEFAULT '0',
			raid_target tinyint(1) unsigned NOT NULL DEFAULT '0',
			armtexture tinyint(2) NOT NULL DEFAULT '0',
			bracertexture tinyint(2) NOT NULL DEFAULT '0',
			handtexture tinyint(2) NOT NULL DEFAULT '0',
			legtexture tinyint(2) NOT NULL DEFAULT '0',
			feettexture tinyint(2) NOT NULL DEFAULT '0',
			light tinyint(2) NOT NULL DEFAULT '0',
			walkspeed tinyint(2) NOT NULL DEFAULT '0',
			peqid int(11) NOT NULL DEFAULT '0',
			unique_ tinyint(2) NOT NULL DEFAULT '0',
			fixed tinyint(2) NOT NULL DEFAULT '0',
			ignore_despawn tinyint(2) NOT NULL DEFAULT '0',
			show_name tinyint(2) NOT NULL DEFAULT '1',
			untargetable tinyint(2) NOT NULL DEFAULT '0',
			PRIMARY KEY (id)
		  ) ENGINE=MyISAM AUTO_INCREMENT=2000059 DEFAULT CHARSET=latin1 PACK_KEYS=0;`)
	if err != nil {
		return
	}
	return
}
