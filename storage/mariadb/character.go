package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	characterFields = `character_data.id, character_data.account_id, character_data.name, character_data.last_name, character_data.title, character_data.suffix, character_data.zone_id, character_data.zone_instance, character_data.y, character_data.x, character_data.z, character_data.heading, character_data.gender, character_data.race, character_data.class, character_data.level, character_data.deity, character_data.birthday, character_data.last_login, character_data.time_played, character_data.level2, character_data.anon, character_data.gm, character_data.face, character_data.hair_color, character_data.hair_style, character_data.beard, character_data.beard_color, character_data.eye_color_1, character_data.eye_color_2, character_data.drakkin_heritage, character_data.drakkin_tattoo, character_data.drakkin_details, character_data.ability_time_seconds, character_data.ability_number, character_data.ability_time_minutes, character_data.ability_time_hours, character_data.exp, character_data.aa_points_spent, character_data.aa_exp, character_data.aa_points, character_data.group_leadership_exp, character_data.raid_leadership_exp, character_data.group_leadership_points, character_data.raid_leadership_points, character_data.points, character_data.cur_hp, character_data.mana, character_data.endurance, character_data.intoxication, character_data.str, character_data.sta, character_data.cha, character_data.dex, character_data.` + "`int`" + `, character_data.agi, character_data.wis, character_data.zone_change_count, character_data.toxicity, character_data.hunger_level, character_data.thirst_level, character_data.ability_up, character_data.ldon_points_guk, character_data.ldon_points_mir, character_data.ldon_points_mmc, character_data.ldon_points_ruj, character_data.ldon_points_tak, character_data.ldon_points_available, character_data.tribute_time_remaining, character_data.career_tribute_points, character_data.tribute_points, character_data.tribute_active, character_data.pvp_status, character_data.pvp_kills, character_data.pvp_deaths, character_data.pvp_current_points, character_data.pvp_career_points, character_data.pvp_best_kill_streak, character_data.pvp_worst_death_streak, character_data.pvp_current_kill_streak, character_data.pvp2, character_data.pvp_type, character_data.show_helm, character_data.group_auto_consent, character_data.raid_auto_consent, character_data.guild_auto_consent, character_data.leadership_exp_on, character_data.RestTimer, character_data.air_remaining, character_data.autosplit_enabled, character_data.lfp, character_data.lfg, character_data.mailkey, character_data.xtargets, character_data.firstlogon, character_data.e_aa_effects, character_data.e_percent_to_aa, character_data.e_expended_aa_spent, character_data.aa_points_spent_old, character_data.aa_points_old, e_last_invsnapshot`
	characterSets   = `character_data.id=:id, character_data.account_id=:account_id, character_data.name=:name, character_data.last_name=:last_name, character_data.title=:title, character_data.suffix=:suffix, character_data.zone_id=:zone_id, character_data.zone_instance=:zone_instance, character_data.y=:y, character_data.x=:x, character_data.z=:z, character_data.heading=:heading, character_data.gender=:gender, character_data.race=:race, character_data.class=:class, character_data.level=:level, character_data.deity=:deity, character_data.birthday=:birthday, character_data.last_login=:last_login, character_data.time_played=:time_played, character_data.level2=:level2, character_data.anon=:anon, character_data.gm=:gm, character_data.face=:face, character_data.hair_color=:hair_color, character_data.hair_style=:hair_style, character_data.beard=:beard, character_data.beard_color=:beard_color, character_data.eye_color_1=:eye_color_1, character_data.eye_color_2=:eye_color_2, character_data.drakkin_heritage=:drakkin_heritage, character_data.drakkin_tattoo=:drakkin_tattoo, character_data.drakkin_details=:drakkin_details, character_data.ability_time_seconds=:ability_time_seconds, character_data.ability_number=:ability_number, character_data.ability_time_minutes=:ability_time_minutes, character_data.ability_time_hours=:ability_time_hours, character_data.exp=:exp, character_data.aa_points_spent=:aa_points_spent, character_data.aa_exp=:aa_exp, character_data.aa_points=:aa_points, character_data.group_leadership_exp=:group_leadership_exp, character_data.raid_leadership_exp=:raid_leadership_exp, character_data.group_leadership_points=:group_leadership_points, character_data.raid_leadership_points=:raid_leadership_points, character_data.points=:points, character_data.cur_hp=:cur_hp, character_data.mana=:mana, character_data.endurance=:endurance, character_data.intoxication=:intoxication, character_data.str=:str, character_data.sta=:sta, character_data.cha=:cha, character_data.dex=:dex, ` + "`character_data.int`" + `=:int, character_data.agi=:agi, character_data.wis=:wis, character_data.zone_change_count=:zone_change_count, character_data.toxicity=:toxicity, character_data.hunger_level=:hunger_level, character_data.thirst_level=:thirst_level, character_data.ability_up=:ability_up, character_data.ldon_points_guk=:ldon_points_guk, character_data.ldon_points_mir=:ldon_points_mir, character_data.ldon_points_mmc=:ldon_points_mmc, character_data.ldon_points_ruj=:ldon_points_ruj, character_data.ldon_points_tak=:ldon_points_tak, character_data.ldon_points_available=:ldon_points_available, character_data.tribute_time_remaining=:tribute_time_remaining, character_data.career_tribute_points=:career_tribute_points, character_data.tribute_points=:tribute_points, character_data.tribute_active=:tribute_active, character_data.pvp_status=:pvp_status, character_data.pvp_kills=:pvp_kills, character_data.pvp_deaths=:pvp_deaths, character_data.pvp_current_points=:pvp_current_points, character_data.pvp_career_points=:pvp_career_points, character_data.pvp_best_kill_streak=:pvp_best_kill_streak, character_data.pvp_worst_death_streak=:pvp_worst_death_streak, character_data.pvp_current_kill_streak=:pvp_current_kill_streak, character_data.pvp2=:pvp2, character_data.pvp_type=:pvp_type, character_data.show_helm=:show_helm, character_data.group_auto_consent=:group_auto_consent, character_data.raid_auto_consent=:raid_auto_consent, character_data.guild_auto_consent=:guild_auto_consent, character_data.leadership_exp_on=:leadership_exp_on, character_data.RestTimer=:RestTimer, character_data.air_remaining=:air_remaining, character_data.autosplit_enabled=:autosplit_enabled, character_data.lfp=:lfp, character_data.lfg=:lfg, character_data.mailkey=:mailkey, character_data.xtargets=:xtargets, character_data.firstlogon=:firstlogon, character_data.e_aa_effects=:e_aa_effects, character_data.e_percent_to_aa=:e_percent_to_aa, character_data.e_expended_aa_spent=:e_expended_aa_spent, character_data.aa_points_spent_old=:aa_points_spent_old, character_data.aa_points_old=:aa_points_old, character_data.e_last_invsnapshot=:e_last_invsnapshot`
	characterBinds  = `:id, :account_id, :name, :last_name, :title, :suffix, :zone_id, :zone_instance, :y, :x, :z, :heading, :gender, :race, :class, :level, :deity, :birthday, :last_login, :time_played, :level2, :anon, :gm, :face, :hair_color, :hair_style, :beard, :beard_color, :eye_color_1, :eye_color_2, :drakkin_heritage, :drakkin_tattoo, :drakkin_details, :ability_time_seconds, :ability_number, :ability_time_minutes, :ability_time_hours, :exp, :aa_points_spent, :aa_exp, :aa_points, :group_leadership_exp, :raid_leadership_exp, :group_leadership_points, :raid_leadership_points, :points, :cur_hp, :mana, :endurance, :intoxication, :str, :sta, :cha, :dex, :int, :agi, :wis, :zone_change_count, :toxicity, :hunger_level, :thirst_level, :ability_up, :ldon_points_guk, :ldon_points_mir, :ldon_points_mmc, :ldon_points_ruj, :ldon_points_tak, :ldon_points_available, :tribute_time_remaining, :career_tribute_points, :tribute_points, :tribute_active, :pvp_status, :pvp_kills, :pvp_deaths, :pvp_current_points, :pvp_career_points, :pvp_best_kill_streak, :pvp_worst_death_streak, :pvp_current_kill_streak, :pvp2, :pvp_type, :show_helm, :group_auto_consent, :raid_auto_consent, :guild_auto_consent, :leadership_exp_on, :RestTimer, :air_remaining, :autosplit_enabled, :lfp, :lfg, :mailkey, :xtargets, :firstlogon, :e_aa_effects, :e_percent_to_aa, :e_expended_aa_spent, :aa_points_spent_old, :aa_points_old, :e_last_invsnapshot`
)

func (s *Storage) GetCharacter(characterId int64) (character *model.Character, err error) {
	character = &model.Character{}
	err = s.db.Get(character, fmt.Sprintf("SELECT id, %s FROM character_data WHERE id = ?", characterFields), characterId)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateCharacter(character *model.Character) (err error) {
	if character == nil {
		err = fmt.Errorf("Must provide character")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO character_data(%s)
		VALUES (%s)`, characterFields, characterBinds), character)
	if err != nil {
		return
	}
	characterId, err := result.LastInsertId()
	if err != nil {
		return
	}
	character.Id = characterId
	return
}

func (s *Storage) ListCharacter() (characters []*model.Character, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM character_data ORDER BY id DESC`, characterFields))
	if err != nil {
		return
	}

	for rows.Next() {
		character := model.Character{}
		if err = rows.StructScan(&character); err != nil {
			return
		}
		characters = append(characters, &character)
	}
	return
}

func (s *Storage) ListCharacterByRanking() (characters []*model.Character, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT character_data.id, %s FROM character_data INNER JOIN
		account ON account.id = character_data.account_id WHERE account.status < 100 ORDER BY cur_hp DESC`, characterFields))
	if err != nil {
		return
	}

	for rows.Next() {
		character := model.Character{}
		if err = rows.StructScan(&character); err != nil {
			return
		}
		characters = append(characters, &character)
	}
	return
}

func (s *Storage) ListCharacterByOnline() (characters []*model.Character, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM character_data 
		WHERE last_login >= UNIX_TIMESTAMP(NOW())-600
		ORDER BY cur_hp DESC`, characterFields))
	if err != nil {
		return
	}

	for rows.Next() {
		character := model.Character{}
		if err = rows.StructScan(&character); err != nil {
			return
		}
		characters = append(characters, &character)
	}
	return
}

func (s *Storage) ListCharacterByAccount(accountId int64) (characters []*model.Character, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM character_data WHERE account_id = ?`, characterFields), accountId)
	if err != nil {
		return
	}

	for rows.Next() {
		character := model.Character{}
		if err = rows.StructScan(&character); err != nil {
			return
		}
		characters = append(characters, &character)
	}
	return
}

func (s *Storage) EditCharacter(characterId int64, character *model.Character) (err error) {
	character.Id = characterId
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE character_data SET %s WHERE id = :id`, characterSets), character)
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

func (s *Storage) DeleteCharacter(characterId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM character_data WHERE id = ?`, characterId)
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

func (s *Storage) SearchCharacter(search string) (characters []*model.Character, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM character_data WHERE name like ? ORDER BY id DESC`, characterFields), "%"+search+"%")
	if err != nil {
		return
	}

	for rows.Next() {
		character := model.Character{}
		if err = rows.StructScan(&character); err != nil {
			return
		}
		characters = append(characters, &character)
	}
	return
}
