package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	characterTable  = "character_data"
	characterFields = `account_id, name, last_name, title, suffix, zone_id, zone_instance, y, x, z, heading, gender, race, class, level, deity, birthday, last_login, time_played, level2, anon, gm, face, hair_color, hair_style, beard, beard_color, eye_color_1, eye_color_2, drakkin_heritage, drakkin_tattoo, drakkin_details, ability_time_seconds, ability_number, ability_time_minutes, ability_time_hours, exp, aa_points_spent, aa_exp, aa_points, group_leadership_exp, raid_leadership_exp, group_leadership_points, raid_leadership_points, points, cur_hp, mana, endurance, intoxication, str, sta, cha, dex, ` + "`int`" + `, agi, wis, zone_change_count, toxicity, hunger_level, thirst_level, ability_up, ldon_points_guk, ldon_points_mir, ldon_points_mmc, ldon_points_ruj, ldon_points_tak, ldon_points_available, tribute_time_remaining, career_tribute_points, tribute_points, tribute_active, pvp_status, pvp_kills, pvp_deaths, pvp_current_points, pvp_career_points, pvp_best_kill_streak, pvp_worst_death_streak, pvp_current_kill_streak, pvp2, pvp_type, show_helm, group_auto_consent, raid_auto_consent, guild_auto_consent, leadership_exp_on, RestTimer, air_remaining, autosplit_enabled, lfp, lfg, mailkey, xtargets, firstlogon, e_aa_effects, e_percent_to_aa, e_expended_aa_spent, aa_points_spent_old, aa_points_old, e_last_invsnapshot`
	characterBinds  = `:account_id, :name, :last_name, :title, :suffix, :zone_id, :zone_instance, :y, :x, :z, :heading, :gender, :race, :class, :level, :deity, :birthday, :last_login, :time_played, :level2, :anon, :gm, :face, :hair_color, :hair_style, :beard, :beard_color, :eye_color_1, :eye_color_2, :drakkin_heritage, :drakkin_tattoo, :drakkin_details, :ability_time_seconds, :ability_number, :ability_time_minutes, :ability_time_hours, :exp, :aa_points_spent, :aa_exp, :aa_points, :group_leadership_exp, :raid_leadership_exp, :group_leadership_points, :raid_leadership_points, :points, :cur_hp, :mana, :endurance, :intoxication, :str, :sta, :cha, :dex, :int, :agi, :wis, :zone_change_count, :toxicity, :hunger_level, :thirst_level, :ability_up, :ldon_points_guk, :ldon_points_mir, :ldon_points_mmc, :ldon_points_ruj, :ldon_points_tak, :ldon_points_available, :tribute_time_remaining, :career_tribute_points, :tribute_points, :tribute_active, :pvp_status, :pvp_kills, :pvp_deaths, :pvp_current_points, :pvp_career_points, :pvp_best_kill_streak, :pvp_worst_death_streak, :pvp_current_kill_streak, :pvp2, :pvp_type, :show_helm, :group_auto_consent, :raid_auto_consent, :guild_auto_consent, :leadership_exp_on, :RestTimer, :air_remaining, :autosplit_enabled, :lfp, :lfg, :mailkey, :xtargets, :firstlogon, :e_aa_effects, :e_percent_to_aa, :e_expended_aa_spent, :aa_points_spent_old, :aa_points_old, :e_last_invsnapshot`
)

//GetCharacter will grab data from storage
func (s *Storage) GetCharacter(character *model.Character) (err error) {
	query := fmt.Sprintf("SELECT id, %s FROM %s WHERE id = ?", characterFields, characterTable)
	err = s.db.Get(character, query, character.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateCharacter will grab data from storage
func (s *Storage) CreateCharacter(character *model.Character) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", characterTable, characterFields, characterBinds)
	result, err := s.db.NamedExec(query, character)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	characterID, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	character.ID = characterID
	return
}

//ListCharacter will grab data from storage
func (s *Storage) ListCharacter(page *model.Page) (characters []*model.Character, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT id, %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", characterFields, characterTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		character := model.Character{}
		if err = rows.StructScan(&character); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		characters = append(characters, &character)
	}
	return
}

//ListCharacterTotalCount will grab data from storage
func (s *Storage) ListCharacterTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(id) FROM %s", characterTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListCharacterBySearch will grab data from storage
func (s *Storage) ListCharacterBySearch(page *model.Page, character *model.Character) (characters []*model.Character, err error) {

	field := ""

	if len(character.Name) > 0 {
		field += `name LIKE :name OR`
		character.Name = fmt.Sprintf("%%%s%%", character.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT id, %s FROM %s WHERE %s LIMIT %d OFFSET %d", characterFields, characterTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, character)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		character := model.Character{}
		if err = rows.StructScan(&character); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		characters = append(characters, &character)
	}
	return
}

//ListCharacterBySearchTotalCount will grab data from storage
func (s *Storage) ListCharacterBySearchTotalCount(character *model.Character) (count int64, err error) {
	field := ""
	if len(character.Name) > 0 {
		field += `name LIKE :name OR`
		character.Name = fmt.Sprintf("%%%s%%", character.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", characterTable, field)

	rows, err := s.db.NamedQuery(query, character)
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

//EditCharacter will grab data from storage
func (s *Storage) EditCharacter(character *model.Character) (err error) {

	prevCharacter := &model.Character{
		ID: character.ID,
	}
	err = s.GetCharacter(prevCharacter)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous character")
		return
	}

	field := ""

	if character.AccountID > 0 && prevCharacter.AccountID != character.AccountID {
		field += "account_id=:account_id, "
	}
	if len(character.Name) > 0 && prevCharacter.Name != character.Name {
		field += "name=:name, "
	}
	if len(character.LastName) > 0 && prevCharacter.LastName != character.LastName {
		field += "last_name=:last_name, "
	}
	if len(character.Title) > 0 && prevCharacter.Title != character.Title {
		field += "title=:title, "
	}
	if len(character.Suffix) > 0 && prevCharacter.Suffix != character.Suffix {
		field += "suffix=:suffix, "
	}
	if character.ZoneID > 0 && prevCharacter.ZoneID != character.ZoneID {
		field += "zone_id=:zone_id, "
	}
	if character.ZoneInstance > 0 && prevCharacter.ZoneInstance != character.ZoneInstance {
		field += "zone_instance=:zone_instance, "
	}
	if character.Y > 0 && prevCharacter.Y != character.Y {
		field += "y=:y, "
	}
	if character.X > 0 && prevCharacter.X != character.X {
		field += "x=:x, "
	}
	if character.Z > 0 && prevCharacter.Z != character.Z {
		field += "z=:z, "
	}
	if character.Heading > 0 && prevCharacter.Heading != character.Heading {
		field += "heading=:heading, "
	}
	if character.Gender > 0 && prevCharacter.Gender != character.Gender {
		field += "gender=:gender, "
	}
	if character.RaceID > 0 && prevCharacter.RaceID != character.RaceID {
		field += "race=:race, "
	}
	if character.ClassID > 0 && prevCharacter.ClassID != character.ClassID {
		field += "class=:class, "
	}
	if character.Level > 0 && prevCharacter.Level != character.Level {
		field += "level=:level, "
	}
	if character.Deity > 0 && prevCharacter.Deity != character.Deity {
		field += "deity=:deity, "
	}
	if character.Birthday > 0 && prevCharacter.Birthday != character.Birthday {
		field += "birthday=:birthday, "
	}
	if character.LastLogin > 0 && prevCharacter.LastLogin != character.LastLogin {
		field += "last_login=:last_login, "
	}
	if character.TimePlayed > 0 && prevCharacter.TimePlayed != character.TimePlayed {
		field += "time_played=:time_played, "
	}
	if character.Level2 > 0 && prevCharacter.Level2 != character.Level2 {
		field += "level2=:level2, "
	}
	if character.Anon > 0 && prevCharacter.Anon != character.Anon {
		field += "anon=:anon, "
	}
	if character.Gm > 0 && prevCharacter.Gm != character.Gm {
		field += "gm=:gm, "
	}
	if character.Face > 0 && prevCharacter.Face != character.Face {
		field += "face=:face, "
	}
	if character.HairColor > 0 && prevCharacter.HairColor != character.HairColor {
		field += "hair_color=:hair_color, "
	}
	if character.HairStyle > 0 && prevCharacter.HairStyle != character.HairStyle {
		field += "hair_style=:hair_style, "
	}
	if character.Beard > 0 && prevCharacter.Beard != character.Beard {
		field += "beard=:beard, "
	}
	if character.BeardColor > 0 && prevCharacter.BeardColor != character.BeardColor {
		field += "beard_color=:beard_color, "
	}
	if character.EyeColor1 > 0 && prevCharacter.EyeColor1 != character.EyeColor1 {
		field += "eye_color_1=:eye_color_1, "
	}
	if character.EyeColor2 > 0 && prevCharacter.EyeColor2 != character.EyeColor2 {
		field += "eye_color_2=:eye_color_2, "
	}
	if character.DrakkinHeritage > 0 && prevCharacter.DrakkinHeritage != character.DrakkinHeritage {
		field += "drakkin_heritage=:drakkin_heritage, "
	}
	if character.DrakkinTattoo > 0 && prevCharacter.DrakkinTattoo != character.DrakkinTattoo {
		field += "drakkin_tattoo=:drakkin_tattoo, "
	}
	if character.DrakkinDetails > 0 && prevCharacter.DrakkinDetails != character.DrakkinDetails {
		field += "drakkin_details=:drakkin_details, "
	}
	if character.AbilityTimeSeconds > 0 && prevCharacter.AbilityTimeSeconds != character.AbilityTimeSeconds {
		field += "ability_time_seconds=:ability_time_seconds, "
	}
	if character.AbilityNumber > 0 && prevCharacter.AbilityNumber != character.AbilityNumber {
		field += "ability_number=:ability_number, "
	}
	if character.AbilityTimeMinutes > 0 && prevCharacter.AbilityTimeMinutes != character.AbilityTimeMinutes {
		field += "ability_time_minutes=:ability_time_minutes, "
	}
	if character.AbilityTimeHours > 0 && prevCharacter.AbilityTimeHours != character.AbilityTimeHours {
		field += "ability_time_hours=:ability_time_hours, "
	}
	if character.Exp > 0 && prevCharacter.Exp != character.Exp {
		field += "exp=:exp, "
	}
	if character.AaPointsSpent > 0 && prevCharacter.AaPointsSpent != character.AaPointsSpent {
		field += "aa_points_spent=:aa_points_spent, "
	}
	if character.AaExp > 0 && prevCharacter.AaExp != character.AaExp {
		field += "aa_exp=:aa_exp, "
	}
	if character.AaPoints > 0 && prevCharacter.AaPoints != character.AaPoints {
		field += "aa_points=:aa_points, "
	}
	if character.GroupLeadershipExp > 0 && prevCharacter.GroupLeadershipExp != character.GroupLeadershipExp {
		field += "group_leadership_exp=:group_leadership_exp, "
	}
	if character.RaidLeadershipExp > 0 && prevCharacter.RaidLeadershipExp != character.RaidLeadershipExp {
		field += "raid_leadership_exp=:raid_leadership_exp, "
	}
	if character.GroupLeadershipPoints > 0 && prevCharacter.GroupLeadershipPoints != character.GroupLeadershipPoints {
		field += "group_leadership_points=:group_leadership_points, "
	}
	if character.RaidLeadershipPoints > 0 && prevCharacter.RaidLeadershipPoints != character.RaidLeadershipPoints {
		field += "raid_leadership_points=:raid_leadership_points, "
	}
	if character.Points > 0 && prevCharacter.Points != character.Points {
		field += "points=:points, "
	}
	if character.CurHp > 0 && prevCharacter.CurHp != character.CurHp {
		field += "cur_hp=:cur_hp, "
	}
	if character.Mana > 0 && prevCharacter.Mana != character.Mana {
		field += "mana=:mana, "
	}
	if character.Endurance > 0 && prevCharacter.Endurance != character.Endurance {
		field += "endurance=:endurance, "
	}
	if character.Intoxication > 0 && prevCharacter.Intoxication != character.Intoxication {
		field += "intoxication=:intoxication, "
	}
	if character.Strength > 0 && prevCharacter.Strength != character.Strength {
		field += "str=:str, "
	}
	if character.Stamina > 0 && prevCharacter.Stamina != character.Stamina {
		field += "sta=:sta, "
	}
	if character.Charisma > 0 && prevCharacter.Charisma != character.Charisma {
		field += "cha=:cha, "
	}
	if character.Dexterity > 0 && prevCharacter.Dexterity != character.Dexterity {
		field += "dex=:dex, "
	}
	if character.Intelligence > 0 && prevCharacter.Intelligence != character.Intelligence {
		field += "'int'=:int, "
	}
	if character.Agility > 0 && prevCharacter.Agility != character.Agility {
		field += "agi=:agi, "
	}
	if character.Wisdom > 0 && prevCharacter.Wisdom != character.Wisdom {
		field += "wis=:wis, "
	}
	if character.ZoneChangeCount > 0 && prevCharacter.ZoneChangeCount != character.ZoneChangeCount {
		field += "zone_change_count=:zone_change_count, "
	}
	if character.Toxicity > 0 && prevCharacter.Toxicity != character.Toxicity {
		field += "toxicity=:toxicity, "
	}
	if character.HungerLevel > 0 && prevCharacter.HungerLevel != character.HungerLevel {
		field += "hunger_level=:hunger_level, "
	}
	if character.ThirstLevel > 0 && prevCharacter.ThirstLevel != character.ThirstLevel {
		field += "thirst_level=:thirst_level, "
	}
	if character.AbilityUp > 0 && prevCharacter.AbilityUp != character.AbilityUp {
		field += "ability_up=:ability_up, "
	}
	if character.LdonPointsGuk > 0 && prevCharacter.LdonPointsGuk != character.LdonPointsGuk {
		field += "ldon_points_guk=:ldon_points_guk, "
	}
	if character.LdonPointsMir > 0 && prevCharacter.LdonPointsMir != character.LdonPointsMir {
		field += "ldon_points_mir=:ldon_points_mir, "
	}
	if character.LdonPointsMmc > 0 && prevCharacter.LdonPointsMmc != character.LdonPointsMmc {
		field += "ldon_points_mmc=:ldon_points_mmc, "
	}
	if character.LdonPointsRuj > 0 && prevCharacter.LdonPointsRuj != character.LdonPointsRuj {
		field += "ldon_points_ruj=:ldon_points_ruj, "
	}
	if character.LdonPointsTak > 0 && prevCharacter.LdonPointsTak != character.LdonPointsTak {
		field += "ldon_points_tak=:ldon_points_tak, "
	}
	if character.LdonPointsAvailable > 0 && prevCharacter.LdonPointsAvailable != character.LdonPointsAvailable {
		field += "ldon_points_available=:ldon_points_available, "
	}
	if character.TributeTimeRemaining > 0 && prevCharacter.TributeTimeRemaining != character.TributeTimeRemaining {
		field += "tribute_time_remaining=:tribute_time_remaining, "
	}
	if character.CareerTributePoints > 0 && prevCharacter.CareerTributePoints != character.CareerTributePoints {
		field += "career_tribute_points=:career_tribute_points, "
	}
	if character.TributePoints > 0 && prevCharacter.TributePoints != character.TributePoints {
		field += "tribute_points=:tribute_points, "
	}
	if character.TributeActive > 0 && prevCharacter.TributeActive != character.TributeActive {
		field += "tribute_active=:tribute_active, "
	}
	if character.PvpStatus > 0 && prevCharacter.PvpStatus != character.PvpStatus {
		field += "pvp_status=:pvp_status, "
	}
	if character.PvpKills > 0 && prevCharacter.PvpKills != character.PvpKills {
		field += "pvp_kills=:pvp_kills, "
	}
	if character.PvpDeaths > 0 && prevCharacter.PvpDeaths != character.PvpDeaths {
		field += "pvp_deaths=:pvp_deaths, "
	}
	if character.PvpCurrentPoints > 0 && prevCharacter.PvpCurrentPoints != character.PvpCurrentPoints {
		field += "pvp_current_points=:pvp_current_points, "
	}
	if character.PvpCareerPoints > 0 && prevCharacter.PvpCareerPoints != character.PvpCareerPoints {
		field += "pvp_career_points=:pvp_career_points, "
	}
	if character.PvpBestKillStreak > 0 && prevCharacter.PvpBestKillStreak != character.PvpBestKillStreak {
		field += "pvp_best_kill_streak=:pvp_best_kill_streak, "
	}
	if character.PvpWorstDeathStreak > 0 && prevCharacter.PvpWorstDeathStreak != character.PvpWorstDeathStreak {
		field += "pvp_worst_death_streak=:pvp_worst_death_streak, "
	}
	if character.PvpCurrentKillStreak > 0 && prevCharacter.PvpCurrentKillStreak != character.PvpCurrentKillStreak {
		field += "pvp_current_kill_streak=:pvp_current_kill_streak, "
	}
	if character.Pvp2 > 0 && prevCharacter.Pvp2 != character.Pvp2 {
		field += "pvp2=:pvp2, "
	}
	if character.PvpType > 0 && prevCharacter.PvpType != character.PvpType {
		field += "pvp_type=:pvp_type, "
	}
	if character.ShowHelm > 0 && prevCharacter.ShowHelm != character.ShowHelm {
		field += "show_helm=:show_helm, "
	}
	if character.GroupAutoConsent > 0 && prevCharacter.GroupAutoConsent != character.GroupAutoConsent {
		field += "group_auto_consent=:group_auto_consent, "
	}
	if character.RaidAutoConsent > 0 && prevCharacter.RaidAutoConsent != character.RaidAutoConsent {
		field += "raid_auto_consent=:raid_auto_consent, "
	}
	if character.GuildAutoConsent > 0 && prevCharacter.GuildAutoConsent != character.GuildAutoConsent {
		field += "guild_auto_consent=:guild_auto_consent, "
	}
	if character.LeadershipExpOn > 0 && prevCharacter.LeadershipExpOn != character.LeadershipExpOn {
		field += "leadership_exp_on=:leadership_exp_on, "
	}
	if character.Resttimer > 0 && prevCharacter.Resttimer != character.Resttimer {
		field += "RestTimer=:RestTimer, "
	}
	if character.AirRemaining > 0 && prevCharacter.AirRemaining != character.AirRemaining {
		field += "air_remaining=:air_remaining, "
	}
	if character.AutosplitEnabled > 0 && prevCharacter.AutosplitEnabled != character.AutosplitEnabled {
		field += "autosplit_enabled=:autosplit_enabled, "
	}
	if character.Lfp > 0 && prevCharacter.Lfp != character.Lfp {
		field += "lfp=:lfp, "
	}
	if character.Lfg > 0 && prevCharacter.Lfg != character.Lfg {
		field += "lfg=:lfg, "
	}
	if len(character.Mailkey) > 0 && prevCharacter.Mailkey != character.Mailkey {
		field += "mailkey=:mailkey, "
	}
	if character.Xtargets > 0 && prevCharacter.Xtargets != character.Xtargets {
		field += "xtargets=:xtargets, "
	}
	if character.Firstlogon > 0 && prevCharacter.Firstlogon != character.Firstlogon {
		field += "firstlogon=:firstlogon, "
	}
	if character.EAaEffects > 0 && prevCharacter.EAaEffects != character.EAaEffects {
		field += "e_aa_effects=:e_aa_effects, "
	}
	if character.EPercentToAa > 0 && prevCharacter.EPercentToAa != character.EPercentToAa {
		field += "e_percent_to_aa=:e_percent_to_aa, "
	}
	if character.EExpendedAaSpent > 0 && prevCharacter.EExpendedAaSpent != character.EExpendedAaSpent {
		field += "e_expended_aa_spent=:e_expended_aa_spent, "
	}
	if character.AaPointsSpentOld > 0 && prevCharacter.AaPointsSpentOld != character.AaPointsSpentOld {
		field += "aa_points_spent_old=:aa_points_spent_old, "
	}
	if character.AaPointsOld > 0 && prevCharacter.AaPointsOld != character.AaPointsOld {
		field += "aa_points_old=:aa_points_old, "
	}
	if character.ELastInvsnapshot > 0 && prevCharacter.ELastInvsnapshot != character.ELastInvsnapshot {
		field += "e_last_invsnapshot=:e_last_invsnapshot, "
	}

	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", characterTable, field)
	result, err := s.db.NamedExec(query, character)
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

//DeleteCharacter will grab data from storage
func (s *Storage) DeleteCharacter(character *model.Character) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", characterTable)
	result, err := s.db.Exec(query, character.ID)
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

//createTableCharacter will grab data from storage
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
) ENGINE=INNODB AUTO_INCREMENT=82152 DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
