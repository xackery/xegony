package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	characterGraphFields = `character_graph.id, character_graph.account_id, character_graph.name, character_graph.last_name, character_graph.title, character_graph.suffix, character_graph.zone_id, character_graph.zone_instance, character_graph.y, character_graph.x, character_graph.z, character_graph.heading, character_graph.gender, character_graph.race, character_graph.class, character_graph.level, character_graph.deity, character_graph.birthday, character_graph.last_login, character_graph.time_played, character_graph.level2, character_graph.anon, character_graph.gm, character_graph.face, character_graph.hair_color, character_graph.hair_style, character_graph.beard, character_graph.beard_color, character_graph.eye_color_1, character_graph.eye_color_2, character_graph.drakkin_heritage, character_graph.drakkin_tattoo, character_graph.drakkin_details, character_graph.ability_time_seconds, character_graph.ability_number, character_graph.ability_time_minutes, character_graph.ability_time_hours, character_graph.exp, character_graph.aa_points_spent, character_graph.aa_exp, character_graph.aa_points, character_graph.group_leadership_exp, character_graph.raid_leadership_exp, character_graph.group_leadership_points, character_graph.raid_leadership_points, character_graph.points, character_graph.cur_hp, character_graph.mana, character_graph.endurance, character_graph.intoxication, character_graph.str, character_graph.sta, character_graph.cha, character_graph.dex, character_graph.` + "`int`" + `, character_graph.agi, character_graph.wis, character_graph.zone_change_count, character_graph.toxicity, character_graph.hunger_level, character_graph.thirst_level, character_graph.ability_up, character_graph.ldon_points_guk, character_graph.ldon_points_mir, character_graph.ldon_points_mmc, character_graph.ldon_points_ruj, character_graph.ldon_points_tak, character_graph.ldon_points_available, character_graph.tribute_time_remaining, character_graph.career_tribute_points, character_graph.tribute_points, character_graph.tribute_active, character_graph.pvp_status, character_graph.pvp_kills, character_graph.pvp_deaths, character_graph.pvp_current_points, character_graph.pvp_career_points, character_graph.pvp_best_kill_streak, character_graph.pvp_worst_death_streak, character_graph.pvp_current_kill_streak, character_graph.pvp2, character_graph.pvp_type, character_graph.show_helm, character_graph.group_auto_consent, character_graph.raid_auto_consent, character_graph.guild_auto_consent, character_graph.leadership_exp_on, character_graph.RestTimer, character_graph.air_remaining, character_graph.autosplit_enabled, character_graph.lfp, character_graph.lfg, character_graph.mailkey, character_graph.xtargets, character_graph.firstlogon, character_graph.e_aa_effects, character_graph.e_percent_to_aa, character_graph.e_expended_aa_spent, character_graph.aa_points_spent_old, character_graph.aa_points_old, e_last_invsnapshot`
	characterGraphSets   = `character_graph.id=:id, character_graph.account_id=:account_id, character_graph.name=:name, character_graph.last_name=:last_name, character_graph.title=:title, character_graph.suffix=:suffix, character_graph.zone_id=:zone_id, character_graph.zone_instance=:zone_instance, character_graph.y=:y, character_graph.x=:x, character_graph.z=:z, character_graph.heading=:heading, character_graph.gender=:gender, character_graph.race=:race, character_graph.class=:class, character_graph.level=:level, character_graph.deity=:deity, character_graph.birthday=:birthday, character_graph.last_login=:last_login, character_graph.time_played=:time_played, character_graph.level2=:level2, character_graph.anon=:anon, character_graph.gm=:gm, character_graph.face=:face, character_graph.hair_color=:hair_color, character_graph.hair_style=:hair_style, character_graph.beard=:beard, character_graph.beard_color=:beard_color, character_graph.eye_color_1=:eye_color_1, character_graph.eye_color_2=:eye_color_2, character_graph.drakkin_heritage=:drakkin_heritage, character_graph.drakkin_tattoo=:drakkin_tattoo, character_graph.drakkin_details=:drakkin_details, character_graph.ability_time_seconds=:ability_time_seconds, character_graph.ability_number=:ability_number, character_graph.ability_time_minutes=:ability_time_minutes, character_graph.ability_time_hours=:ability_time_hours, character_graph.exp=:exp, character_graph.aa_points_spent=:aa_points_spent, character_graph.aa_exp=:aa_exp, character_graph.aa_points=:aa_points, character_graph.group_leadership_exp=:group_leadership_exp, character_graph.raid_leadership_exp=:raid_leadership_exp, character_graph.group_leadership_points=:group_leadership_points, character_graph.raid_leadership_points=:raid_leadership_points, character_graph.points=:points, character_graph.cur_hp=:cur_hp, character_graph.mana=:mana, character_graph.endurance=:endurance, character_graph.intoxication=:intoxication, character_graph.str=:str, character_graph.sta=:sta, character_graph.cha=:cha, character_graph.dex=:dex, character_graph.int=:int, character_graph.agi=:agi, character_graph.wis=:wis, character_graph.zone_change_count=:zone_change_count, character_graph.toxicity=:toxicity, character_graph.hunger_level=:hunger_level, character_graph.thirst_level=:thirst_level, character_graph.ability_up=:ability_up, character_graph.ldon_points_guk=:ldon_points_guk, character_graph.ldon_points_mir=:ldon_points_mir, character_graph.ldon_points_mmc=:ldon_points_mmc, character_graph.ldon_points_ruj=:ldon_points_ruj, character_graph.ldon_points_tak=:ldon_points_tak, character_graph.ldon_points_available=:ldon_points_available, character_graph.tribute_time_remaining=:tribute_time_remaining, character_graph.career_tribute_points=:career_tribute_points, character_graph.tribute_points=:tribute_points, character_graph.tribute_active=:tribute_active, character_graph.pvp_status=:pvp_status, character_graph.pvp_kills=:pvp_kills, character_graph.pvp_deaths=:pvp_deaths, character_graph.pvp_current_points=:pvp_current_points, character_graph.pvp_career_points=:pvp_career_points, character_graph.pvp_best_kill_streak=:pvp_best_kill_streak, character_graph.pvp_worst_death_streak=:pvp_worst_death_streak, character_graph.pvp_current_kill_streak=:pvp_current_kill_streak, character_graph.pvp2=:pvp2, character_graph.pvp_type=:pvp_type, character_graph.show_helm=:show_helm, character_graph.group_auto_consent=:group_auto_consent, character_graph.raid_auto_consent=:raid_auto_consent, character_graph.guild_auto_consent=:guild_auto_consent, character_graph.leadership_exp_on=:leadership_exp_on, character_graph.RestTimer=:RestTimer, character_graph.air_remaining=:air_remaining, character_graph.autosplit_enabled=:autosplit_enabled, character_graph.lfp=:lfp, character_graph.lfg=:lfg, character_graph.mailkey=:mailkey, character_graph.xtargets=:xtargets, character_graph.firstlogon=:firstlogon, character_graph.e_aa_effects=:e_aa_effects, character_graph.e_percent_to_aa=:e_percent_to_aa, character_graph.e_expended_aa_spent=:e_expended_aa_spent, character_graph.aa_points_spent_old=:aa_points_spent_old, character_graph.aa_points_old=:aa_points_old, character_graph.e_last_invsnapshot=:e_last_invsnapshot`
	characterGraphBinds  = `:id, :account_id, :name, :last_name, :title, :suffix, :zone_id, :zone_instance, :y, :x, :z, :heading, :gender, :race, :class, :level, :deity, :birthday, :last_login, :time_played, :level2, :anon, :gm, :face, :hair_color, :hair_style, :beard, :beard_color, :eye_color_1, :eye_color_2, :drakkin_heritage, :drakkin_tattoo, :drakkin_details, :ability_time_seconds, :ability_number, :ability_time_minutes, :ability_time_hours, :exp, :aa_points_spent, :aa_exp, :aa_points, :group_leadership_exp, :raid_leadership_exp, :group_leadership_points, :raid_leadership_points, :points, :cur_hp, :mana, :endurance, :intoxication, :str, :sta, :cha, :dex, :int, :agi, :wis, :zone_change_count, :toxicity, :hunger_level, :thirst_level, :ability_up, :ldon_points_guk, :ldon_points_mir, :ldon_points_mmc, :ldon_points_ruj, :ldon_points_tak, :ldon_points_available, :tribute_time_remaining, :career_tribute_points, :tribute_points, :tribute_active, :pvp_status, :pvp_kills, :pvp_deaths, :pvp_current_points, :pvp_career_points, :pvp_best_kill_streak, :pvp_worst_death_streak, :pvp_current_kill_streak, :pvp2, :pvp_type, :show_helm, :group_auto_consent, :raid_auto_consent, :guild_auto_consent, :leadership_exp_on, :RestTimer, :air_remaining, :autosplit_enabled, :lfp, :lfg, :mailkey, :xtargets, :firstlogon, :e_aa_effects, :e_percent_to_aa, :e_expended_aa_spent, :aa_points_spent_old, :aa_points_old, :e_last_invsnapshot`
)

//GetCharacterGraph will grab data from storage
func (s *Storage) GetCharacterGraph(id int64) (characterGraph *model.CharacterGraph, err error) {
	characterGraph = &model.CharacterGraph{}
	err = s.db.Get(characterGraph, fmt.Sprintf("SELECT id, %s FROM character_graph WHERE id = ?", characterGraphFields), id)
	if err != nil {
		return
	}
	return
}

//CreateCharacterGraph will grab data from storage
func (s *Storage) CreateCharacterGraph(characterGraph *model.CharacterGraph) (err error) {
	if characterGraph == nil {
		err = fmt.Errorf("Must provide characterGraph")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO character_graph(%s)
		VALUES (%s)`, characterGraphFields, characterGraphBinds), characterGraph)
	if err != nil {
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	characterGraph.ID = id
	return
}

//ListCharacterGraph will grab data from storage
func (s *Storage) ListCharacterGraph(characterID int64) (characterGraphs []*model.CharacterGraph, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM character_graph WHERE character_id = ? ORDER BY create_date DESC`, characterGraphFields), characterID)
	if err != nil {
		return
	}

	for rows.Next() {
		characterGraph := model.CharacterGraph{}
		if err = rows.StructScan(&characterGraph); err != nil {
			return
		}
		characterGraphs = append(characterGraphs, &characterGraph)
	}
	return
}

//EditCharacterGraph will grab data from storage
func (s *Storage) EditCharacterGraph(id int64, characterGraph *model.CharacterGraph) (err error) {
	characterGraph.ID = id
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE character_graph SET %s WHERE id = :id`, characterGraphSets), characterGraph)
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

//DeleteCharacterGraph will grab data from storage
func (s *Storage) DeleteCharacterGraph(id int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM character_graph WHERE id = ?`, id)
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

//createTableCharacterGraph will grab data from storage
func (s *Storage) createTableCharacterGraph() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE character_graph (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  character_id int(11) DEFAULT NULL,
  create_date datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  experience int(10) unsigned NOT NULL DEFAULT '0',
  aa_experience int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	return
}
