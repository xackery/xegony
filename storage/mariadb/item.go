package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	itemTable  = "items"
	itemFields = `id, name, aagi, ac, accuracy, acha, adex, aint, artifactflag, asta, astr, attack, augrestrict, augslot1type, augslot1visible, augslot2type, augslot2visible, augslot3type, augslot3visible, augslot4type, augslot4visible, augslot5type, augslot5visible, augslot6type, augslot6visible, augtype, avoidance, awis, bagsize, bagslots, bagtype, bagwr, banedmgamt, banedmgraceamt, banedmgbody, banedmgrace, bardtype, bardvalue, book, casttime, casttime_, charmfile, charmfileid, classes, color, combateffects, extradmgskill, extradmgamt, price, cr, damage, damageshield, deity, delay, augdistiller, dotshielding, dr, clicktype, clicklevel2, elemdmgtype, elemdmgamt, endur, factionamt1, factionamt2, factionamt3, factionamt4, factionmod1, factionmod2, factionmod3, factionmod4, filename, focuseffect, fr, fvnodrop, haste, clicklevel, hp, regen, icon, idfile, itemclass, itemtype, ldonprice, ldontheme, ldonsold, light, lore, loregroup, magic, mana, manaregen, enduranceregen, material, herosforgemodel, maxcharges, mr, nodrop, norent, pendingloreflag, pr, procrate, races, ` + "`range`" + `, reclevel, recskill, reqlevel, sellrate, shielding, size, skillmodtype, skillmodvalue, slots, clickeffect, spellshield, strikethrough, stunresist, summonedflag, tradeskills, favor, weight, UNK012, UNK013, benefitflag, UNK054, UNK059, booktype, recastdelay, recasttype, guildfavor, UNK123, UNK124, attuneable, nopet, updated, comment, UNK127, pointtype, potionbelt, potionbeltslots, stacksize, notransfer, stackable, UNK134, UNK137, proceffect, proctype, proclevel2, proclevel, UNK142, worneffect, worntype, wornlevel2, wornlevel, UNK147, focustype, focuslevel2, focuslevel, UNK152, scrolleffect, scrolltype, scrolllevel2, scrolllevel, UNK157, serialized, verified, serialization, source, UNK033, lorefile, UNK014, svcorruption, skillmodmax, UNK060, augslot1unk2, augslot2unk2, augslot3unk2, augslot4unk2, augslot5unk2, augslot6unk2, UNK120, UNK121, questitemflag, UNK132, clickunk5, clickunk6, clickunk7, procunk1, procunk2, procunk3, procunk4, procunk6, procunk7, wornunk1, wornunk2, wornunk3, wornunk4, wornunk5, wornunk6, wornunk7, focusunk1, focusunk2, focusunk3, focusunk4, focusunk5, focusunk6, focusunk7, scrollunk1, scrollunk2, scrollunk3, scrollunk4, scrollunk5, scrollunk6, scrollunk7, UNK193, purity, evoitem, evoid, evolvinglevel, evomax, clickname, procname, wornname, focusname, scrollname, dsmitigation, heroic_str, heroic_int, heroic_wis, heroic_agi, heroic_dex, heroic_sta, heroic_cha, heroic_pr, heroic_dr, heroic_fr, heroic_cr, heroic_mr, heroic_svcorrup, healamt, spelldmg, clairvoyance, backstabdmg, created, elitematerial, ldonsellbackrate, scriptfileid, expendablearrow, powersourcecapacity, bardeffect, bardeffecttype, bardlevel2, bardlevel, bardunk1, bardunk2, bardunk3, bardunk4, bardunk5, bardname, bardunk7, UNK214, UNK219, UNK220, UNK221, heirloom, UNK223, UNK224, UNK225, UNK226, UNK227, UNK228, UNK229, UNK230, UNK231, UNK232, UNK233, UNK234, placeable, UNK236, UNK237, UNK238, UNK239, UNK240, UNK241, epicitem`
	itemBinds  = `:id, :name, :aagi, :ac, :accuracy, :acha, :adex, :aint, :artifactflag, :asta, :astr, :attack, :augrestrict, :augslot1type, :augslot1visible, :augslot2type, :augslot2visible, :augslot3type, :augslot3visible, :augslot4type, :augslot4visible, :augslot5type, :augslot5visible, :augslot6type, :augslot6visible, :augtype, :avoidance, :awis, :bagsize, :bagslots, :bagtype, :bagwr, :banedmgamt, :banedmgraceamt, :banedmgbody, :banedmgrace, :bardtype, :bardvalue, :book, :casttime, :casttime_, :charmfile, :charmfileid, :classes, :color, :combateffects, :extradmgskill, :extradmgamt, :price, :cr, :damage, :damageshield, :deity, :delay, :augdistiller, :dotshielding, :dr, :clicktype, :clicklevel2, :elemdmgtype, :elemdmgamt, :endur, :factionamt1, :factionamt2, :factionamt3, :factionamt4, :factionmod1, :factionmod2, :factionmod3, :factionmod4, :filename, :focuseffect, :fr, :fvnodrop, :haste, :clicklevel, :hp, :regen, :icon, :idfile, :itemclass, :itemtype, :ldonprice, :ldontheme, :ldonsold, :light, :lore, :loregroup, :magic, :mana, :manaregen, :enduranceregen, :material, :herosforgemodel, :maxcharges, :mr, :nodrop, :norent, :pendingloreflag, :pr, :procrate, :races, :range, :reclevel, :recskill, :reqlevel, :sellrate, :shielding, :size, :skillmodtype, :skillmodvalue, :slots, :clickeffect, :spellshield, :strikethrough, :stunresist, :summonedflag, :tradeskills, :favor, :weight, :UNK012, :UNK013, :benefitflag, :UNK054, :UNK059, :booktype, :recastdelay, :recasttype, :guildfavor, :UNK123, :UNK124, :attuneable, :nopet, :updated, :comment, :UNK127, :pointtype, :potionbelt, :potionbeltslots, :stacksize, :notransfer, :stackable, :UNK134, :UNK137, :proceffect, :proctype, :proclevel2, :proclevel, :UNK142, :worneffect, :worntype, :wornlevel2, :wornlevel, :UNK147, :focustype, :focuslevel2, :focuslevel, :UNK152, :scrolleffect, :scrolltype, :scrolllevel2, :scrolllevel, :UNK157, :serialized, :verified, :serialization, :source, :UNK033, :lorefile, :UNK014, :svcorruption, :skillmodmax, :UNK060, :augslot1unk2, :augslot2unk2, :augslot3unk2, :augslot4unk2, :augslot5unk2, :augslot6unk2, :UNK120, :UNK121, :questitemflag, :UNK132, :clickunk5, :clickunk6, :clickunk7, :procunk1, :procunk2, :procunk3, :procunk4, :procunk6, :procunk7, :wornunk1, :wornunk2, :wornunk3, :wornunk4, :wornunk5, :wornunk6, :wornunk7, :focusunk1, :focusunk2, :focusunk3, :focusunk4, :focusunk5, :focusunk6, :focusunk7, :scrollunk1, :scrollunk2, :scrollunk3, :scrollunk4, :scrollunk5, :scrollunk6, :scrollunk7, :UNK193, :purity, :evoitem, :evoid, :evolvinglevel, :evomax, :clickname, :procname, :wornname, :focusname, :scrollname, :dsmitigation, :heroic_str, :heroic_int, :heroic_wis, :heroic_agi, :heroic_dex, :heroic_sta, :heroic_cha, :heroic_pr, :heroic_dr, :heroic_fr, :heroic_cr, :heroic_mr, :heroic_svcorrup, :healamt, :spelldmg, :clairvoyance, :backstabdmg, :created, :elitematerial, :ldonsellbackrate, :scriptfileid, :expendablearrow, :powersourcecapacity, :bardeffect, :bardeffecttype, :bardlevel2, :bardlevel, :bardunk1, :bardunk2, :bardunk3, :bardunk4, :bardunk5, :bardname, :bardunk7, :UNK214, :UNK219, :UNK220, :UNK221, :heirloom, :UNK223, :UNK224, :UNK225, :UNK226, :UNK227, :UNK228, :UNK229, :UNK230, :UNK231, :UNK232, :UNK233, :UNK234, :placeable, :UNK236, :UNK237, :UNK238, :UNK239, :UNK240, :UNK241, epicitem`
)

//GetItem will grab data from storage
func (s *Storage) GetItem(item *model.Item) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", itemFields, itemTable)
	err = s.db.Get(item, query, item.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateItem will grab data from storage
func (s *Storage) CreateItem(item *model.Item) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", itemTable, itemFields, itemBinds)
	result, err := s.db.NamedExec(query, item)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	itemID, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	item.ID = itemID
	return
}

//ListItem will grab data from storage
func (s *Storage) ListItem(page *model.Page) (items []*model.Item, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", itemFields, itemTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		item := model.Item{}
		if err = rows.StructScan(&item); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		items = append(items, &item)
	}
	return
}

//ListItemTotalCount will grab data from storage
func (s *Storage) ListItemTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(id) FROM %s", itemTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListItemBySearch will grab data from storage
func (s *Storage) ListItemBySearch(page *model.Page, item *model.Item) (items []*model.Item, err error) {

	field := ""

	if len(item.Name) > 0 {
		field += `name LIKE :name OR`
		item.Name = fmt.Sprintf("%%%s%%", item.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", itemFields, itemTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, item)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		item := model.Item{}
		if err = rows.StructScan(&item); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		items = append(items, &item)
	}
	return
}

//ListItemBySearchTotalCount will grab data from storage
func (s *Storage) ListItemBySearchTotalCount(item *model.Item) (count int64, err error) {
	field := ""
	if len(item.Name) > 0 {
		field += `name LIKE :name OR`
		item.Name = fmt.Sprintf("%%%s%%", item.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", itemTable, field)

	rows, err := s.db.NamedQuery(query, item)
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

//EditItem will grab data from storage
func (s *Storage) EditItem(item *model.Item) (err error) {

	prevItem := &model.Item{
		ID: item.ID,
	}
	err = s.GetItem(prevItem)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous item")
		return
	}

	field := ""

	/*if item.AccountID > 0 && prevItem.AccountID != item.AccountID {
		field += "account_id=:account_id, "
	}
	if len(item.Name) > 0 && prevItem.Name != item.Name {
		field += "name=:name, "
	}
	if len(item.LastName) > 0 && prevItem.LastName != item.LastName {
		field += "last_name=:last_name, "
	}
	if len(item.Title) > 0 && prevItem.Title != item.Title {
		field += "title=:title, "
	}
	if len(item.Suffix) > 0 && prevItem.Suffix != item.Suffix {
		field += "suffix=:suffix, "
	}
	if item.ZoneID > 0 && prevItem.ZoneID != item.ZoneID {
		field += "zone_id=:zone_id, "
	}
	if item.ZoneInstance > 0 && prevItem.ZoneInstance != item.ZoneInstance {
		field += "zone_instance=:zone_instance, "
	}
	if item.Y > 0 && prevItem.Y != item.Y {
		field += "y=:y, "
	}
	if item.X > 0 && prevItem.X != item.X {
		field += "x=:x, "
	}
	if item.Z > 0 && prevItem.Z != item.Z {
		field += "z=:z, "
	}
	if item.Heading > 0 && prevItem.Heading != item.Heading {
		field += "heading=:heading, "
	}
	if item.Gender > 0 && prevItem.Gender != item.Gender {
		field += "gender=:gender, "
	}
	if item.RaceID > 0 && prevItem.RaceID != item.RaceID {
		field += "race=:race, "
	}
	if item.ClassID > 0 && prevItem.ClassID != item.ClassID {
		field += "class=:class, "
	}
	if item.Level > 0 && prevItem.Level != item.Level {
		field += "level=:level, "
	}
	if item.Deity > 0 && prevItem.Deity != item.Deity {
		field += "deity=:deity, "
	}
	if item.Birthday > 0 && prevItem.Birthday != item.Birthday {
		field += "birthday=:birthday, "
	}
	if item.LastLogin > 0 && prevItem.LastLogin != item.LastLogin {
		field += "last_login=:last_login, "
	}
	if item.TimePlayed > 0 && prevItem.TimePlayed != item.TimePlayed {
		field += "time_played=:time_played, "
	}
	if item.Level2 > 0 && prevItem.Level2 != item.Level2 {
		field += "level2=:level2, "
	}
	if item.Anon > 0 && prevItem.Anon != item.Anon {
		field += "anon=:anon, "
	}
	if item.Gm > 0 && prevItem.Gm != item.Gm {
		field += "gm=:gm, "
	}
	if item.Face > 0 && prevItem.Face != item.Face {
		field += "face=:face, "
	}
	if item.HairColor > 0 && prevItem.HairColor != item.HairColor {
		field += "hair_color=:hair_color, "
	}
	if item.HairStyle > 0 && prevItem.HairStyle != item.HairStyle {
		field += "hair_style=:hair_style, "
	}
	if item.Beard > 0 && prevItem.Beard != item.Beard {
		field += "beard=:beard, "
	}
	if item.BeardColor > 0 && prevItem.BeardColor != item.BeardColor {
		field += "beard_color=:beard_color, "
	}
	if item.EyeColor1 > 0 && prevItem.EyeColor1 != item.EyeColor1 {
		field += "eye_color_1=:eye_color_1, "
	}
	if item.EyeColor2 > 0 && prevItem.EyeColor2 != item.EyeColor2 {
		field += "eye_color_2=:eye_color_2, "
	}
	if item.DrakkinHeritage > 0 && prevItem.DrakkinHeritage != item.DrakkinHeritage {
		field += "drakkin_heritage=:drakkin_heritage, "
	}
	if item.DrakkinTattoo > 0 && prevItem.DrakkinTattoo != item.DrakkinTattoo {
		field += "drakkin_tattoo=:drakkin_tattoo, "
	}
	if item.DrakkinDetails > 0 && prevItem.DrakkinDetails != item.DrakkinDetails {
		field += "drakkin_details=:drakkin_details, "
	}
	if item.AbilityTimeSeconds > 0 && prevItem.AbilityTimeSeconds != item.AbilityTimeSeconds {
		field += "ability_time_seconds=:ability_time_seconds, "
	}
	if item.AbilityNumber > 0 && prevItem.AbilityNumber != item.AbilityNumber {
		field += "ability_number=:ability_number, "
	}
	if item.AbilityTimeMinutes > 0 && prevItem.AbilityTimeMinutes != item.AbilityTimeMinutes {
		field += "ability_time_minutes=:ability_time_minutes, "
	}
	if item.AbilityTimeHours > 0 && prevItem.AbilityTimeHours != item.AbilityTimeHours {
		field += "ability_time_hours=:ability_time_hours, "
	}
	if item.Exp > 0 && prevItem.Exp != item.Exp {
		field += "exp=:exp, "
	}
	if item.AaPointsSpent > 0 && prevItem.AaPointsSpent != item.AaPointsSpent {
		field += "aa_points_spent=:aa_points_spent, "
	}
	if item.AaExp > 0 && prevItem.AaExp != item.AaExp {
		field += "aa_exp=:aa_exp, "
	}
	if item.AaPoints > 0 && prevItem.AaPoints != item.AaPoints {
		field += "aa_points=:aa_points, "
	}
	if item.GroupLeadershipExp > 0 && prevItem.GroupLeadershipExp != item.GroupLeadershipExp {
		field += "group_leadership_exp=:group_leadership_exp, "
	}
	if item.RaidLeadershipExp > 0 && prevItem.RaidLeadershipExp != item.RaidLeadershipExp {
		field += "raid_leadership_exp=:raid_leadership_exp, "
	}
	if item.GroupLeadershipPoints > 0 && prevItem.GroupLeadershipPoints != item.GroupLeadershipPoints {
		field += "group_leadership_points=:group_leadership_points, "
	}
	if item.RaidLeadershipPoints > 0 && prevItem.RaidLeadershipPoints != item.RaidLeadershipPoints {
		field += "raid_leadership_points=:raid_leadership_points, "
	}
	if item.Points > 0 && prevItem.Points != item.Points {
		field += "points=:points, "
	}
	if item.CurHp > 0 && prevItem.CurHp != item.CurHp {
		field += "cur_hp=:cur_hp, "
	}
	if item.Mana > 0 && prevItem.Mana != item.Mana {
		field += "mana=:mana, "
	}
	if item.Endurance > 0 && prevItem.Endurance != item.Endurance {
		field += "endurance=:endurance, "
	}
	if item.Intoxication > 0 && prevItem.Intoxication != item.Intoxication {
		field += "intoxication=:intoxication, "
	}
	if item.Str > 0 && prevItem.Str != item.Str {
		field += "str=:str, "
	}
	if item.Sta > 0 && prevItem.Sta != item.Sta {
		field += "sta=:sta, "
	}
	if item.Cha > 0 && prevItem.Cha != item.Cha {
		field += "cha=:cha, "
	}
	if item.Dex > 0 && prevItem.Dex != item.Dex {
		field += "dex=:dex, "
	}
	if item.Int > 0 && prevItem.Int != item.Int {
		field += "'int'=:int, "
	}
	if item.Agi > 0 && prevItem.Agi != item.Agi {
		field += "agi=:agi, "
	}
	if item.Wis > 0 && prevItem.Wis != item.Wis {
		field += "wis=:wis, "
	}
	if item.ZoneChangeCount > 0 && prevItem.ZoneChangeCount != item.ZoneChangeCount {
		field += "zone_change_count=:zone_change_count, "
	}
	if item.Toxicity > 0 && prevItem.Toxicity != item.Toxicity {
		field += "toxicity=:toxicity, "
	}
	if item.HungerLevel > 0 && prevItem.HungerLevel != item.HungerLevel {
		field += "hunger_level=:hunger_level, "
	}
	if item.ThirstLevel > 0 && prevItem.ThirstLevel != item.ThirstLevel {
		field += "thirst_level=:thirst_level, "
	}
	if item.AbilityUp > 0 && prevItem.AbilityUp != item.AbilityUp {
		field += "ability_up=:ability_up, "
	}
	if item.LdonPointsGuk > 0 && prevItem.LdonPointsGuk != item.LdonPointsGuk {
		field += "ldon_points_guk=:ldon_points_guk, "
	}
	if item.LdonPointsMir > 0 && prevItem.LdonPointsMir != item.LdonPointsMir {
		field += "ldon_points_mir=:ldon_points_mir, "
	}
	if item.LdonPointsMmc > 0 && prevItem.LdonPointsMmc != item.LdonPointsMmc {
		field += "ldon_points_mmc=:ldon_points_mmc, "
	}
	if item.LdonPointsRuj > 0 && prevItem.LdonPointsRuj != item.LdonPointsRuj {
		field += "ldon_points_ruj=:ldon_points_ruj, "
	}
	if item.LdonPointsTak > 0 && prevItem.LdonPointsTak != item.LdonPointsTak {
		field += "ldon_points_tak=:ldon_points_tak, "
	}
	if item.LdonPointsAvailable > 0 && prevItem.LdonPointsAvailable != item.LdonPointsAvailable {
		field += "ldon_points_available=:ldon_points_available, "
	}
	if item.TributeTimeRemaining > 0 && prevItem.TributeTimeRemaining != item.TributeTimeRemaining {
		field += "tribute_time_remaining=:tribute_time_remaining, "
	}
	if item.CareerTributePoints > 0 && prevItem.CareerTributePoints != item.CareerTributePoints {
		field += "career_tribute_points=:career_tribute_points, "
	}
	if item.TributePoints > 0 && prevItem.TributePoints != item.TributePoints {
		field += "tribute_points=:tribute_points, "
	}
	if item.TributeActive > 0 && prevItem.TributeActive != item.TributeActive {
		field += "tribute_active=:tribute_active, "
	}
	if item.PvpStatus > 0 && prevItem.PvpStatus != item.PvpStatus {
		field += "pvp_status=:pvp_status, "
	}
	if item.PvpKills > 0 && prevItem.PvpKills != item.PvpKills {
		field += "pvp_kills=:pvp_kills, "
	}
	if item.PvpDeaths > 0 && prevItem.PvpDeaths != item.PvpDeaths {
		field += "pvp_deaths=:pvp_deaths, "
	}
	if item.PvpCurrentPoints > 0 && prevItem.PvpCurrentPoints != item.PvpCurrentPoints {
		field += "pvp_current_points=:pvp_current_points, "
	}
	if item.PvpCareerPoints > 0 && prevItem.PvpCareerPoints != item.PvpCareerPoints {
		field += "pvp_career_points=:pvp_career_points, "
	}
	if item.PvpBestKillStreak > 0 && prevItem.PvpBestKillStreak != item.PvpBestKillStreak {
		field += "pvp_best_kill_streak=:pvp_best_kill_streak, "
	}
	if item.PvpWorstDeathStreak > 0 && prevItem.PvpWorstDeathStreak != item.PvpWorstDeathStreak {
		field += "pvp_worst_death_streak=:pvp_worst_death_streak, "
	}
	if item.PvpCurrentKillStreak > 0 && prevItem.PvpCurrentKillStreak != item.PvpCurrentKillStreak {
		field += "pvp_current_kill_streak=:pvp_current_kill_streak, "
	}
	if item.Pvp2 > 0 && prevItem.Pvp2 != item.Pvp2 {
		field += "pvp2=:pvp2, "
	}
	if item.PvpType > 0 && prevItem.PvpType != item.PvpType {
		field += "pvp_type=:pvp_type, "
	}
	if item.ShowHelm > 0 && prevItem.ShowHelm != item.ShowHelm {
		field += "show_helm=:show_helm, "
	}
	if item.GroupAutoConsent > 0 && prevItem.GroupAutoConsent != item.GroupAutoConsent {
		field += "group_auto_consent=:group_auto_consent, "
	}
	if item.RaidAutoConsent > 0 && prevItem.RaidAutoConsent != item.RaidAutoConsent {
		field += "raid_auto_consent=:raid_auto_consent, "
	}
	if item.GuildAutoConsent > 0 && prevItem.GuildAutoConsent != item.GuildAutoConsent {
		field += "guild_auto_consent=:guild_auto_consent, "
	}
	if item.LeadershipExpOn > 0 && prevItem.LeadershipExpOn != item.LeadershipExpOn {
		field += "leadership_exp_on=:leadership_exp_on, "
	}
	if item.Resttimer > 0 && prevItem.Resttimer != item.Resttimer {
		field += "RestTimer=:RestTimer, "
	}
	if item.AirRemaining > 0 && prevItem.AirRemaining != item.AirRemaining {
		field += "air_remaining=:air_remaining, "
	}
	if item.AutosplitEnabled > 0 && prevItem.AutosplitEnabled != item.AutosplitEnabled {
		field += "autosplit_enabled=:autosplit_enabled, "
	}
	if item.Lfp > 0 && prevItem.Lfp != item.Lfp {
		field += "lfp=:lfp, "
	}
	if item.Lfg > 0 && prevItem.Lfg != item.Lfg {
		field += "lfg=:lfg, "
	}
	if len(item.Mailkey) > 0 && prevItem.Mailkey != item.Mailkey {
		field += "mailkey=:mailkey, "
	}
	if item.Xtargets > 0 && prevItem.Xtargets != item.Xtargets {
		field += "xtargets=:xtargets, "
	}
	if item.Firstlogon > 0 && prevItem.Firstlogon != item.Firstlogon {
		field += "firstlogon=:firstlogon, "
	}
	if item.EAaEffects > 0 && prevItem.EAaEffects != item.EAaEffects {
		field += "e_aa_effects=:e_aa_effects, "
	}
	if item.EPercentToAa > 0 && prevItem.EPercentToAa != item.EPercentToAa {
		field += "e_percent_to_aa=:e_percent_to_aa, "
	}
	if item.EExpendedAaSpent > 0 && prevItem.EExpendedAaSpent != item.EExpendedAaSpent {
		field += "e_expended_aa_spent=:e_expended_aa_spent, "
	}
	if item.AaPointsSpentOld > 0 && prevItem.AaPointsSpentOld != item.AaPointsSpentOld {
		field += "aa_points_spent_old=:aa_points_spent_old, "
	}
	if item.AaPointsOld > 0 && prevItem.AaPointsOld != item.AaPointsOld {
		field += "aa_points_old=:aa_points_old, "
	}
	if item.ELastInvsnapshot > 0 && prevItem.ELastInvsnapshot != item.ELastInvsnapshot {
		field += "e_last_invsnapshot=:e_last_invsnapshot, "
	}*/

	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", itemTable, field)
	result, err := s.db.NamedExec(query, item)
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

//DeleteItem will grab data from storage
func (s *Storage) DeleteItem(item *model.Item) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", itemTable)
	result, err := s.db.Exec(query, item.ID)
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

//createTableItem will grab data from storage
func (s *Storage) createTableItem() (err error) {
	_, err = s.db.Exec(`
		CREATE TABLE items (
			id int(11) NOT NULL DEFAULT '0',
			minstatus smallint(5) NOT NULL DEFAULT '0',
			Name varchar(64) NOT NULL DEFAULT '',
			aagi int(11) NOT NULL DEFAULT '0',
			ac int(11) NOT NULL DEFAULT '0',
			accuracy int(11) NOT NULL DEFAULT '0',
			acha int(11) NOT NULL DEFAULT '0',
			adex int(11) NOT NULL DEFAULT '0',
			aint int(11) NOT NULL DEFAULT '0',
			artifactflag tinyint(3) unsigned NOT NULL DEFAULT '0',
			asta int(11) NOT NULL DEFAULT '0',
			astr int(11) NOT NULL DEFAULT '0',
			attack int(11) NOT NULL DEFAULT '0',
			augrestrict int(11) NOT NULL DEFAULT '0',
			augslot1type tinyint(3) NOT NULL DEFAULT '0',
			augslot1visible tinyint(3) DEFAULT NULL,
			augslot2type tinyint(3) NOT NULL DEFAULT '0',
			augslot2visible tinyint(3) DEFAULT NULL,
			augslot3type tinyint(3) NOT NULL DEFAULT '0',
			augslot3visible tinyint(3) DEFAULT NULL,
			augslot4type tinyint(3) NOT NULL DEFAULT '0',
			augslot4visible tinyint(3) DEFAULT NULL,
			augslot5type tinyint(3) NOT NULL DEFAULT '0',
			augslot5visible tinyint(3) DEFAULT NULL,
			augslot6type tinyint(3) NOT NULL DEFAULT '0',
			augslot6visible tinyint(3) NOT NULL DEFAULT '0',
			augtype int(11) NOT NULL DEFAULT '0',
			avoidance int(11) NOT NULL DEFAULT '0',
			awis int(11) NOT NULL DEFAULT '0',
			bagsize int(11) NOT NULL DEFAULT '0',
			bagslots int(11) NOT NULL DEFAULT '0',
			bagtype int(11) NOT NULL DEFAULT '0',
			bagwr int(11) NOT NULL DEFAULT '0',
			banedmgamt int(11) NOT NULL DEFAULT '0',
			banedmgraceamt int(11) NOT NULL DEFAULT '0',
			banedmgbody int(11) NOT NULL DEFAULT '0',
			banedmgrace int(11) NOT NULL DEFAULT '0',
			bardtype int(11) NOT NULL DEFAULT '0',
			bardvalue int(11) NOT NULL DEFAULT '0',
			book int(11) NOT NULL DEFAULT '0',
			casttime int(11) NOT NULL DEFAULT '0',
			casttime_ int(11) NOT NULL DEFAULT '0',
			charmfile varchar(32) NOT NULL DEFAULT '',
			charmfileid varchar(32) NOT NULL DEFAULT '',
			classes int(11) NOT NULL DEFAULT '0',
			color int(10) unsigned NOT NULL DEFAULT '0',
			combateffects varchar(10) NOT NULL DEFAULT '',
			extradmgskill int(11) NOT NULL DEFAULT '0',
			extradmgamt int(11) NOT NULL DEFAULT '0',
			price int(11) NOT NULL DEFAULT '0',
			cr int(11) NOT NULL DEFAULT '0',
			damage int(11) NOT NULL DEFAULT '0',
			damageshield int(11) NOT NULL DEFAULT '0',
			deity int(11) NOT NULL DEFAULT '0',
			delay int(11) NOT NULL DEFAULT '0',
			augdistiller int(11) NOT NULL DEFAULT '0',
			dotshielding int(11) NOT NULL DEFAULT '0',
			dr int(11) NOT NULL DEFAULT '0',
			clicktype int(11) NOT NULL DEFAULT '0',
			clicklevel2 int(11) NOT NULL DEFAULT '0',
			elemdmgtype int(11) NOT NULL DEFAULT '0',
			elemdmgamt int(11) NOT NULL DEFAULT '0',
			endur int(11) NOT NULL DEFAULT '0',
			factionamt1 int(11) NOT NULL DEFAULT '0',
			factionamt2 int(11) NOT NULL DEFAULT '0',
			factionamt3 int(11) NOT NULL DEFAULT '0',
			factionamt4 int(11) NOT NULL DEFAULT '0',
			factionmod1 int(11) NOT NULL DEFAULT '0',
			factionmod2 int(11) NOT NULL DEFAULT '0',
			factionmod3 int(11) NOT NULL DEFAULT '0',
			factionmod4 int(11) NOT NULL DEFAULT '0',
			filename varchar(32) NOT NULL DEFAULT '',
			focuseffect int(11) NOT NULL DEFAULT '0',
			fr int(11) NOT NULL DEFAULT '0',
			fvnodrop int(11) NOT NULL DEFAULT '0',
			haste int(11) NOT NULL DEFAULT '0',
			clicklevel int(11) NOT NULL DEFAULT '0',
			hp int(11) NOT NULL DEFAULT '0',
			regen int(11) NOT NULL DEFAULT '0',
			icon int(11) NOT NULL DEFAULT '0',
			idfile varchar(30) NOT NULL DEFAULT '',
			itemclass int(11) NOT NULL DEFAULT '0',
			itemtype int(11) NOT NULL DEFAULT '0',
			ldonprice int(11) NOT NULL DEFAULT '0',
			ldontheme int(11) NOT NULL DEFAULT '0',
			ldonsold int(11) NOT NULL DEFAULT '0',
			light int(11) NOT NULL DEFAULT '0',
			lore varchar(80) NOT NULL DEFAULT '',
			loregroup int(11) NOT NULL DEFAULT '0',
			magic int(11) NOT NULL DEFAULT '0',
			mana int(11) NOT NULL DEFAULT '0',
			manaregen int(11) NOT NULL DEFAULT '0',
			enduranceregen int(11) NOT NULL DEFAULT '0',
			material int(11) NOT NULL DEFAULT '0',
			herosforgemodel int(11) NOT NULL DEFAULT '0',
			maxcharges int(11) NOT NULL DEFAULT '0',
			mr int(11) NOT NULL DEFAULT '0',
			nodrop int(11) NOT NULL DEFAULT '0',
			norent int(11) NOT NULL DEFAULT '0',
			pendingloreflag tinyint(3) unsigned NOT NULL DEFAULT '0',
			pr int(11) NOT NULL DEFAULT '0',
			procrate int(11) NOT NULL DEFAULT '0',
			races int(11) NOT NULL DEFAULT '0',
			range int(11) NOT NULL DEFAULT '0',
			reclevel int(11) NOT NULL DEFAULT '0',
			recskill int(11) NOT NULL DEFAULT '0',
			reqlevel int(11) NOT NULL DEFAULT '0',
			sellrate float NOT NULL DEFAULT '0',
			shielding int(11) NOT NULL DEFAULT '0',
			size int(11) NOT NULL DEFAULT '0',
			skillmodtype int(11) NOT NULL DEFAULT '0',
			skillmodvalue int(11) NOT NULL DEFAULT '0',
			slots int(11) NOT NULL DEFAULT '0',
			clickeffect int(11) NOT NULL DEFAULT '0',
			spellshield int(11) NOT NULL DEFAULT '0',
			strikethrough int(11) NOT NULL DEFAULT '0',
			stunresist int(11) NOT NULL DEFAULT '0',
			summonedflag tinyint(3) unsigned NOT NULL DEFAULT '0',
			tradeskills int(11) NOT NULL DEFAULT '0',
			favor int(11) NOT NULL DEFAULT '0',
			weight int(11) NOT NULL DEFAULT '0',
			UNK012 int(11) NOT NULL DEFAULT '0',
			UNK013 int(11) NOT NULL DEFAULT '0',
			benefitflag int(11) NOT NULL DEFAULT '0',
			UNK054 int(11) NOT NULL DEFAULT '0',
			UNK059 int(11) NOT NULL DEFAULT '0',
			booktype int(11) NOT NULL DEFAULT '0',
			recastdelay int(11) NOT NULL DEFAULT '0',
			recasttype int(11) NOT NULL DEFAULT '0',
			guildfavor int(11) NOT NULL DEFAULT '0',
			UNK123 int(11) NOT NULL DEFAULT '0',
			UNK124 int(11) NOT NULL DEFAULT '0',
			attuneable int(11) NOT NULL DEFAULT '0',
			nopet int(11) NOT NULL DEFAULT '0',
			updated datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
			comment varchar(255) NOT NULL DEFAULT '',
			UNK127 int(11) NOT NULL DEFAULT '0',
			pointtype int(11) NOT NULL DEFAULT '0',
			potionbelt int(11) NOT NULL DEFAULT '0',
			potionbeltslots int(11) NOT NULL DEFAULT '0',
			stacksize int(11) NOT NULL DEFAULT '0',
			notransfer int(11) NOT NULL DEFAULT '0',
			stackable int(11) NOT NULL DEFAULT '0',
			UNK134 varchar(255) NOT NULL DEFAULT '',
			UNK137 int(11) NOT NULL DEFAULT '0',
			proceffect int(11) NOT NULL DEFAULT '0',
			proctype int(11) NOT NULL DEFAULT '0',
			proclevel2 int(11) NOT NULL DEFAULT '0',
			proclevel int(11) NOT NULL DEFAULT '0',
			UNK142 int(11) NOT NULL DEFAULT '0',
			worneffect int(11) NOT NULL DEFAULT '0',
			worntype int(11) NOT NULL DEFAULT '0',
			wornlevel2 int(11) NOT NULL DEFAULT '0',
			wornlevel int(11) NOT NULL DEFAULT '0',
			UNK147 int(11) NOT NULL DEFAULT '0',
			focustype int(11) NOT NULL DEFAULT '0',
			focuslevel2 int(11) NOT NULL DEFAULT '0',
			focuslevel int(11) NOT NULL DEFAULT '0',
			UNK152 int(11) NOT NULL DEFAULT '0',
			scrolleffect int(11) NOT NULL DEFAULT '0',
			scrolltype int(11) NOT NULL DEFAULT '0',
			scrolllevel2 int(11) NOT NULL DEFAULT '0',
			scrolllevel int(11) NOT NULL DEFAULT '0',
			UNK157 int(11) NOT NULL DEFAULT '0',
			serialized datetime DEFAULT NULL,
			verified datetime DEFAULT NULL,
			serialization text,
			source varchar(20) NOT NULL DEFAULT '',
			UNK033 int(11) NOT NULL DEFAULT '0',
			lorefile varchar(32) NOT NULL DEFAULT '',
			UNK014 int(11) NOT NULL DEFAULT '0',
			svcorruption int(11) NOT NULL DEFAULT '0',
			skillmodmax int(11) NOT NULL DEFAULT '0',
			UNK060 int(11) NOT NULL DEFAULT '0',
			augslot1unk2 int(11) NOT NULL DEFAULT '0',
			augslot2unk2 int(11) NOT NULL DEFAULT '0',
			augslot3unk2 int(11) NOT NULL DEFAULT '0',
			augslot4unk2 int(11) NOT NULL DEFAULT '0',
			augslot5unk2 int(11) NOT NULL DEFAULT '0',
			augslot6unk2 int(11) NOT NULL DEFAULT '0',
			UNK120 int(11) NOT NULL DEFAULT '0',
			UNK121 int(11) NOT NULL DEFAULT '0',
			questitemflag int(11) NOT NULL DEFAULT '0',
			UNK132 text NOT NULL,
			clickunk5 int(11) NOT NULL DEFAULT '0',
			clickunk6 varchar(32) NOT NULL DEFAULT '',
			clickunk7 int(11) NOT NULL DEFAULT '0',
			procunk1 int(11) NOT NULL DEFAULT '0',
			procunk2 int(11) NOT NULL DEFAULT '0',
			procunk3 int(11) NOT NULL DEFAULT '0',
			procunk4 int(11) NOT NULL DEFAULT '0',
			procunk6 varchar(32) NOT NULL DEFAULT '',
			procunk7 int(11) NOT NULL DEFAULT '0',
			wornunk1 int(11) NOT NULL DEFAULT '0',
			wornunk2 int(11) NOT NULL DEFAULT '0',
			wornunk3 int(11) NOT NULL DEFAULT '0',
			wornunk4 int(11) NOT NULL DEFAULT '0',
			wornunk5 int(11) NOT NULL DEFAULT '0',
			wornunk6 varchar(32) NOT NULL DEFAULT '',
			wornunk7 int(11) NOT NULL DEFAULT '0',
			focusunk1 int(11) NOT NULL DEFAULT '0',
			focusunk2 int(11) NOT NULL DEFAULT '0',
			focusunk3 int(11) NOT NULL DEFAULT '0',
			focusunk4 int(11) NOT NULL DEFAULT '0',
			focusunk5 int(11) NOT NULL DEFAULT '0',
			focusunk6 varchar(32) NOT NULL DEFAULT '',
			focusunk7 int(11) NOT NULL DEFAULT '0',
			scrollunk1 int(11) NOT NULL DEFAULT '0',
			scrollunk2 int(11) NOT NULL DEFAULT '0',
			scrollunk3 int(11) NOT NULL DEFAULT '0',
			scrollunk4 int(11) NOT NULL DEFAULT '0',
			scrollunk5 int(11) NOT NULL DEFAULT '0',
			scrollunk6 varchar(32) NOT NULL DEFAULT '',
			scrollunk7 int(11) NOT NULL DEFAULT '0',
			UNK193 int(11) NOT NULL DEFAULT '0',
			purity int(11) NOT NULL DEFAULT '0',
			evoitem int(11) NOT NULL DEFAULT '0',
			evoid int(11) NOT NULL DEFAULT '0',
			evolvinglevel int(11) NOT NULL DEFAULT '0',
			evomax int(11) NOT NULL DEFAULT '0',
			clickname varchar(64) NOT NULL DEFAULT '',
			procname varchar(64) NOT NULL DEFAULT '',
			wornname varchar(64) NOT NULL DEFAULT '',
			focusname varchar(64) NOT NULL DEFAULT '',
			scrollname varchar(64) NOT NULL DEFAULT '',
			dsmitigation smallint(6) NOT NULL DEFAULT '0',
			heroic_str smallint(6) NOT NULL DEFAULT '0',
			heroic_int smallint(6) NOT NULL DEFAULT '0',
			heroic_wis smallint(6) NOT NULL DEFAULT '0',
			heroic_agi smallint(6) NOT NULL DEFAULT '0',
			heroic_dex smallint(6) NOT NULL DEFAULT '0',
			heroic_sta smallint(6) NOT NULL DEFAULT '0',
			heroic_cha smallint(6) NOT NULL DEFAULT '0',
			heroic_pr smallint(6) NOT NULL DEFAULT '0',
			heroic_dr smallint(6) NOT NULL DEFAULT '0',
			heroic_fr smallint(6) NOT NULL DEFAULT '0',
			heroic_cr smallint(6) NOT NULL DEFAULT '0',
			heroic_mr smallint(6) NOT NULL DEFAULT '0',
			heroic_svcorrup smallint(6) NOT NULL DEFAULT '0',
			healamt smallint(6) NOT NULL DEFAULT '0',
			spelldmg smallint(6) NOT NULL DEFAULT '0',
			clairvoyance smallint(6) NOT NULL DEFAULT '0',
			backstabdmg smallint(6) NOT NULL DEFAULT '0',
			created varchar(64) NOT NULL DEFAULT '',
			elitematerial smallint(6) NOT NULL DEFAULT '0',
			ldonsellbackrate smallint(6) NOT NULL DEFAULT '0',
			scriptfileid smallint(6) NOT NULL DEFAULT '0',
			expendablearrow smallint(6) NOT NULL DEFAULT '0',
			powersourcecapacity smallint(6) NOT NULL DEFAULT '0',
			bardeffect smallint(6) NOT NULL DEFAULT '0',
			bardeffecttype smallint(6) NOT NULL DEFAULT '0',
			bardlevel2 smallint(6) NOT NULL DEFAULT '0',
			bardlevel smallint(6) NOT NULL DEFAULT '0',
			bardunk1 smallint(6) NOT NULL DEFAULT '0',
			bardunk2 smallint(6) NOT NULL DEFAULT '0',
			bardunk3 smallint(6) NOT NULL DEFAULT '0',
			bardunk4 smallint(6) NOT NULL DEFAULT '0',
			bardunk5 smallint(6) NOT NULL DEFAULT '0',
			bardname varchar(64) NOT NULL DEFAULT '',
			bardunk7 smallint(6) NOT NULL DEFAULT '0',
			UNK214 smallint(6) NOT NULL DEFAULT '0',
			UNK219 int(11) NOT NULL DEFAULT '0',
			UNK220 int(11) NOT NULL DEFAULT '0',
			UNK221 int(11) NOT NULL DEFAULT '0',
			heirloom int(11) NOT NULL DEFAULT '0',
			UNK223 int(11) NOT NULL DEFAULT '0',
			UNK224 int(11) NOT NULL DEFAULT '0',
			UNK225 int(11) NOT NULL DEFAULT '0',
			UNK226 int(11) NOT NULL DEFAULT '0',
			UNK227 int(11) NOT NULL DEFAULT '0',
			UNK228 int(11) NOT NULL DEFAULT '0',
			UNK229 int(11) NOT NULL DEFAULT '0',
			UNK230 int(11) NOT NULL DEFAULT '0',
			UNK231 int(11) NOT NULL DEFAULT '0',
			UNK232 int(11) NOT NULL DEFAULT '0',
			UNK233 int(11) NOT NULL DEFAULT '0',
			UNK234 int(11) NOT NULL DEFAULT '0',
			placeable int(11) NOT NULL DEFAULT '0',
			UNK236 int(11) NOT NULL DEFAULT '0',
			UNK237 int(11) NOT NULL DEFAULT '0',
			UNK238 int(11) NOT NULL DEFAULT '0',
			UNK239 int(11) NOT NULL DEFAULT '0',
			UNK240 int(11) NOT NULL DEFAULT '0',
			UNK241 int(11) NOT NULL DEFAULT '0',
			epicitem int(11) NOT NULL DEFAULT '0',
			UNIQUE KEY ID (id),
			KEY name_idx (Name),
			KEY lore_idx (lore),
			KEY minstatus (minstatus)
		  ) ENGINE=MyISAM DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
