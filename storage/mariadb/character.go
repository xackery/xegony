package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	characterFields = `character_data.id, character_data.account_id, character_data.name, character_data.last_name, character_data.title, character_data.suffix, character_data.zone_id, character_data.zone_instance, character_data.y, character_data.x, character_data.z, character_data.heading, character_data.gender, character_data.race, character_data.class, character_data.level, character_data.deity, character_data.birthday, character_data.last_login, character_data.time_played, character_data.level2, character_data.anon, character_data.gm, character_data.face, character_data.hair_color, character_data.hair_style, character_data.beard, character_data.beard_color, character_data.eye_color_1, character_data.eye_color_2, character_data.drakkin_heritage, character_data.drakkin_tattoo, character_data.drakkin_details, character_data.ability_time_seconds, character_data.ability_number, character_data.ability_time_minutes, character_data.ability_time_hours, character_data.exp, character_data.aa_points_spent, character_data.aa_exp, character_data.aa_points, character_data.group_leadership_exp, character_data.raid_leadership_exp, character_data.group_leadership_points, character_data.raid_leadership_points, character_data.points, character_data.cur_hp, character_data.mana, character_data.endurance, character_data.intoxication, character_data.str, character_data.sta, character_data.cha, character_data.dex, character_data.` + "`int`" + `, character_data.agi, character_data.wis, character_data.zone_change_count, character_data.toxicity, character_data.hunger_level, character_data.thirst_level, character_data.ability_up, character_data.ldon_points_guk, character_data.ldon_points_mir, character_data.ldon_points_mmc, character_data.ldon_points_ruj, character_data.ldon_points_tak, character_data.ldon_points_available, character_data.tribute_time_remaining, character_data.career_tribute_points, character_data.tribute_points, character_data.tribute_active, character_data.pvp_status, character_data.pvp_kills, character_data.pvp_deaths, character_data.pvp_current_points, character_data.pvp_career_points, character_data.pvp_best_kill_streak, character_data.pvp_worst_death_streak, character_data.pvp_current_kill_streak, character_data.pvp2, character_data.pvp_type, character_data.show_helm, character_data.group_auto_consent, character_data.raid_auto_consent, character_data.guild_auto_consent, character_data.leadership_exp_on, character_data.RestTimer, character_data.air_remaining, character_data.autosplit_enabled, character_data.lfp, character_data.lfg, character_data.mailkey, character_data.xtargets, character_data.firstlogon, character_data.e_aa_effects, character_data.e_percent_to_aa, character_data.e_expended_aa_spent, character_data.aa_points_spent_old, character_data.aa_points_old, e_last_invsnapshot`
	characterSets   = `character_data.id=:id, character_data.account_id=:account_id, character_data.name=:name, character_data.last_name=:last_name, character_data.title=:title, character_data.suffix=:suffix, character_data.zone_id=:zone_id, character_data.zone_instance=:zone_instance, character_data.y=:y, character_data.x=:x, character_data.z=:z, character_data.heading=:heading, character_data.gender=:gender, character_data.race=:race, character_data.class=:class, character_data.level=:level, character_data.deity=:deity, character_data.birthday=:birthday, character_data.last_login=:last_login, character_data.time_played=:time_played, character_data.level2=:level2, character_data.anon=:anon, character_data.gm=:gm, character_data.face=:face, character_data.hair_color=:hair_color, character_data.hair_style=:hair_style, character_data.beard=:beard, character_data.beard_color=:beard_color, character_data.eye_color_1=:eye_color_1, character_data.eye_color_2=:eye_color_2, character_data.drakkin_heritage=:drakkin_heritage, character_data.drakkin_tattoo=:drakkin_tattoo, character_data.drakkin_details=:drakkin_details, character_data.ability_time_seconds=:ability_time_seconds, character_data.ability_number=:ability_number, character_data.ability_time_minutes=:ability_time_minutes, character_data.ability_time_hours=:ability_time_hours, character_data.exp=:exp, character_data.aa_points_spent=:aa_points_spent, character_data.aa_exp=:aa_exp, character_data.aa_points=:aa_points, character_data.group_leadership_exp=:group_leadership_exp, character_data.raid_leadership_exp=:raid_leadership_exp, character_data.group_leadership_points=:group_leadership_points, character_data.raid_leadership_points=:raid_leadership_points, character_data.points=:points, character_data.cur_hp=:cur_hp, character_data.mana=:mana, character_data.endurance=:endurance, character_data.intoxication=:intoxication, character_data.str=:str, character_data.sta=:sta, character_data.cha=:cha, character_data.dex=:dex, character_data.int=:int, character_data.agi=:agi, character_data.wis=:wis, character_data.zone_change_count=:zone_change_count, character_data.toxicity=:toxicity, character_data.hunger_level=:hunger_level, character_data.thirst_level=:thirst_level, character_data.ability_up=:ability_up, character_data.ldon_points_guk=:ldon_points_guk, character_data.ldon_points_mir=:ldon_points_mir, character_data.ldon_points_mmc=:ldon_points_mmc, character_data.ldon_points_ruj=:ldon_points_ruj, character_data.ldon_points_tak=:ldon_points_tak, character_data.ldon_points_available=:ldon_points_available, character_data.tribute_time_remaining=:tribute_time_remaining, character_data.career_tribute_points=:career_tribute_points, character_data.tribute_points=:tribute_points, character_data.tribute_active=:tribute_active, character_data.pvp_status=:pvp_status, character_data.pvp_kills=:pvp_kills, character_data.pvp_deaths=:pvp_deaths, character_data.pvp_current_points=:pvp_current_points, character_data.pvp_career_points=:pvp_career_points, character_data.pvp_best_kill_streak=:pvp_best_kill_streak, character_data.pvp_worst_death_streak=:pvp_worst_death_streak, character_data.pvp_current_kill_streak=:pvp_current_kill_streak, character_data.pvp2=:pvp2, character_data.pvp_type=:pvp_type, character_data.show_helm=:show_helm, character_data.group_auto_consent=:group_auto_consent, character_data.raid_auto_consent=:raid_auto_consent, character_data.guild_auto_consent=:guild_auto_consent, character_data.leadership_exp_on=:leadership_exp_on, character_data.RestTimer=:RestTimer, character_data.air_remaining=:air_remaining, character_data.autosplit_enabled=:autosplit_enabled, character_data.lfp=:lfp, character_data.lfg=:lfg, character_data.mailkey=:mailkey, character_data.xtargets=:xtargets, character_data.firstlogon=:firstlogon, character_data.e_aa_effects=:e_aa_effects, character_data.e_percent_to_aa=:e_percent_to_aa, character_data.e_expended_aa_spent=:e_expended_aa_spent, character_data.aa_points_spent_old=:aa_points_spent_old, character_data.aa_points_old=:aa_points_old, character_data.e_last_invsnapshot=:e_last_invsnapshot`
	characterBinds  = `:id, :account_id, :name, :last_name, :title, :suffix, :zone_id, :zone_instance, :y, :x, :z, :heading, :gender, :race, :class, :level, :deity, :birthday, :last_login, :time_played, :level2, :anon, :gm, :face, :hair_color, :hair_style, :beard, :beard_color, :eye_color_1, :eye_color_2, :drakkin_heritage, :drakkin_tattoo, :drakkin_details, :ability_time_seconds, :ability_number, :ability_time_minutes, :ability_time_hours, :exp, :aa_points_spent, :aa_exp, :aa_points, :group_leadership_exp, :raid_leadership_exp, :group_leadership_points, :raid_leadership_points, :points, :cur_hp, :mana, :endurance, :intoxication, :str, :sta, :cha, :dex, :int, :agi, :wis, :zone_change_count, :toxicity, :hunger_level, :thirst_level, :ability_up, :ldon_points_guk, :ldon_points_mir, :ldon_points_mmc, :ldon_points_ruj, :ldon_points_tak, :ldon_points_available, :tribute_time_remaining, :career_tribute_points, :tribute_points, :tribute_active, :pvp_status, :pvp_kills, :pvp_deaths, :pvp_current_points, :pvp_career_points, :pvp_best_kill_streak, :pvp_worst_death_streak, :pvp_current_kill_streak, :pvp2, :pvp_type, :show_helm, :group_auto_consent, :raid_auto_consent, :guild_auto_consent, :leadership_exp_on, :RestTimer, :air_remaining, :autosplit_enabled, :lfp, :lfg, :mailkey, :xtargets, :firstlogon, :e_aa_effects, :e_percent_to_aa, :e_expended_aa_spent, :aa_points_spent_old, :aa_points_old, :e_last_invsnapshot`
)

func (s *Storage) GetCharacter(characterID int64) (character *model.Character, err error) {
	character = &model.Character{}
	err = s.db.Get(character, fmt.Sprintf("SELECT id, %s FROM character_data WHERE id = ?", characterFields), characterID)
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
	characterID, err := result.LastInsertId()
	if err != nil {
		return
	}
	character.ID = characterID
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
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT character_data.id, %s FROM character_data 
		INNER JOIN account ON account.id = character_data.account_id 
		WHERE account.status < 100 ORDER BY cur_hp DESC LIMIT 10`, characterFields))
	if err != nil {
		return
	}

	for rows.Next() {
		character := model.Character{}
		if err = rows.StructScan(&character); err != nil {
			return
		}

		character.Base, err = s.GetBase(character.Level, character.Class)
		if err != nil {
			return
		}

		character.Inventory, err = s.ListItemByCharacter(character.ID)
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

func (s *Storage) ListCharacterByAccount(accountID int64) (characters []*model.Character, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM character_data WHERE account_id = ?`, characterFields), accountID)
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

func (s *Storage) EditCharacter(characterID int64, character *model.Character) (err error) {
	character.ID = characterID
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

func (s *Storage) DeleteCharacter(characterID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM character_data WHERE id = ?`, characterID)
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

func (s *Storage) createTableCharacter() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE character_data (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  account_id int(11) NOT NULL DEFAULT '0',
  name varchar(64) NOT NULL DEFAULT '',
  last_name varchar(64) NOT NULL DEFAULT '',
  title varchar(32) NOT NULL DEFAULT '',
  suffix varchar(32) NOT NULL DEFAULT '',
  zone_id int(11) unsigned NOT NULL DEFAULT '0',
  zone_instance int(11) unsigned NOT NULL DEFAULT '0',
  y float NOT NULL DEFAULT '0',
  x float NOT NULL DEFAULT '0',
  z float NOT NULL DEFAULT '0',
  heading float NOT NULL DEFAULT '0',
  gender tinyint(11) unsigned NOT NULL DEFAULT '0',
  race smallint(11) unsigned NOT NULL DEFAULT '0',
  class tinyint(11) unsigned NOT NULL DEFAULT '0',
  level int(11) unsigned NOT NULL DEFAULT '0',
  deity int(11) unsigned NOT NULL DEFAULT '0',
  birthday int(11) unsigned NOT NULL DEFAULT '0',
  last_login int(11) unsigned NOT NULL DEFAULT '0',
  time_played int(11) unsigned NOT NULL DEFAULT '0',
  level2 tinyint(11) unsigned NOT NULL DEFAULT '0',
  anon tinyint(11) unsigned NOT NULL DEFAULT '0',
  gm tinyint(11) unsigned NOT NULL DEFAULT '0',
  face int(11) unsigned NOT NULL DEFAULT '0',
  hair_color tinyint(11) unsigned NOT NULL DEFAULT '0',
  hair_style tinyint(11) unsigned NOT NULL DEFAULT '0',
  beard tinyint(11) unsigned NOT NULL DEFAULT '0',
  beard_color tinyint(11) unsigned NOT NULL DEFAULT '0',
  eye_color_1 tinyint(11) unsigned NOT NULL DEFAULT '0',
  eye_color_2 tinyint(11) unsigned NOT NULL DEFAULT '0',
  drakkin_heritage int(11) unsigned NOT NULL DEFAULT '0',
  drakkin_tattoo int(11) unsigned NOT NULL DEFAULT '0',
  drakkin_details int(11) unsigned NOT NULL DEFAULT '0',
  ability_time_seconds tinyint(11) unsigned NOT NULL DEFAULT '0',
  ability_number tinyint(11) unsigned NOT NULL DEFAULT '0',
  ability_time_minutes tinyint(11) unsigned NOT NULL DEFAULT '0',
  ability_time_hours tinyint(11) unsigned NOT NULL DEFAULT '0',
  exp int(11) unsigned NOT NULL DEFAULT '0',
  aa_points_spent int(11) unsigned NOT NULL DEFAULT '0',
  aa_exp int(11) unsigned NOT NULL DEFAULT '0',
  aa_points int(11) unsigned NOT NULL DEFAULT '0',
  group_leadership_exp int(11) unsigned NOT NULL DEFAULT '0',
  raid_leadership_exp int(11) unsigned NOT NULL DEFAULT '0',
  group_leadership_points int(11) unsigned NOT NULL DEFAULT '0',
  raid_leadership_points int(11) unsigned NOT NULL DEFAULT '0',
  points int(11) unsigned NOT NULL DEFAULT '0',
  cur_hp int(11) unsigned NOT NULL DEFAULT '0',
  mana int(11) unsigned NOT NULL DEFAULT '0',
  endurance int(11) unsigned NOT NULL DEFAULT '0',
  intoxication int(11) unsigned NOT NULL DEFAULT '0',
  str int(11) unsigned NOT NULL DEFAULT '0',
  sta int(11) unsigned NOT NULL DEFAULT '0',
  cha int(11) unsigned NOT NULL DEFAULT '0',
  dex int(11) unsigned NOT NULL DEFAULT '0',
  ` + "`int`" + ` int(11) unsigned NOT NULL DEFAULT '0',
  agi int(11) unsigned NOT NULL DEFAULT '0',
  wis int(11) unsigned NOT NULL DEFAULT '0',
  zone_change_count int(11) unsigned NOT NULL DEFAULT '0',
  toxicity int(11) unsigned NOT NULL DEFAULT '0',
  hunger_level int(11) unsigned NOT NULL DEFAULT '0',
  thirst_level int(11) unsigned NOT NULL DEFAULT '0',
  ability_up int(11) unsigned NOT NULL DEFAULT '0',
  ldon_points_guk int(11) unsigned NOT NULL DEFAULT '0',
  ldon_points_mir int(11) unsigned NOT NULL DEFAULT '0',
  ldon_points_mmc int(11) unsigned NOT NULL DEFAULT '0',
  ldon_points_ruj int(11) unsigned NOT NULL DEFAULT '0',
  ldon_points_tak int(11) unsigned NOT NULL DEFAULT '0',
  ldon_points_available int(11) unsigned NOT NULL DEFAULT '0',
  tribute_time_remaining int(11) unsigned NOT NULL DEFAULT '0',
  career_tribute_points int(11) unsigned NOT NULL DEFAULT '0',
  tribute_points int(11) unsigned NOT NULL DEFAULT '0',
  tribute_active int(11) unsigned NOT NULL DEFAULT '0',
  pvp_status tinyint(11) unsigned NOT NULL DEFAULT '0',
  pvp_kills int(11) unsigned NOT NULL DEFAULT '0',
  pvp_deaths int(11) unsigned NOT NULL DEFAULT '0',
  pvp_current_points int(11) unsigned NOT NULL DEFAULT '0',
  pvp_career_points int(11) unsigned NOT NULL DEFAULT '0',
  pvp_best_kill_streak int(11) unsigned NOT NULL DEFAULT '0',
  pvp_worst_death_streak int(11) unsigned NOT NULL DEFAULT '0',
  pvp_current_kill_streak int(11) unsigned NOT NULL DEFAULT '0',
  pvp2 int(11) unsigned NOT NULL DEFAULT '0',
  pvp_type int(11) unsigned NOT NULL DEFAULT '0',
  show_helm int(11) unsigned NOT NULL DEFAULT '0',
  group_auto_consent tinyint(11) unsigned NOT NULL DEFAULT '0',
  raid_auto_consent tinyint(11) unsigned NOT NULL DEFAULT '0',
  guild_auto_consent tinyint(11) unsigned NOT NULL DEFAULT '0',
  leadership_exp_on tinyint(11) unsigned NOT NULL DEFAULT '0',
  RestTimer int(11) unsigned NOT NULL DEFAULT '0',
  air_remaining int(11) unsigned NOT NULL DEFAULT '0',
  autosplit_enabled int(11) unsigned NOT NULL DEFAULT '0',
  lfp tinyint(1) unsigned NOT NULL DEFAULT '0',
  lfg tinyint(1) unsigned NOT NULL DEFAULT '0',
  mailkey char(16) NOT NULL DEFAULT '',
  xtargets tinyint(3) unsigned NOT NULL DEFAULT '5',
  firstlogon tinyint(3) NOT NULL DEFAULT '0',
  e_aa_effects int(11) unsigned NOT NULL DEFAULT '0',
  e_percent_to_aa int(11) unsigned NOT NULL DEFAULT '0',
  e_expended_aa_spent int(11) unsigned NOT NULL DEFAULT '0',
  aa_points_spent_old int(11) unsigned NOT NULL DEFAULT '0',
  aa_points_old int(11) unsigned NOT NULL DEFAULT '0',
  e_last_invsnapshot int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (id),
  UNIQUE KEY name (name),
  KEY account_id (account_id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
