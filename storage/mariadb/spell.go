package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	spellTable  = "spells_new"
	spellFields = "id, name, player_1, teleport_zone, you_cast, other_casts, cast_on_you, cast_on_other, spell_fades, `range`, aoerange, pushback, pushup, cast_time, recovery_time, recast_time, buffdurationformula, buffduration, AEDuration, mana, icon, memicon, LightType, goodEffect, Activated, resisttype, targettype, basediff, skill, zonetype, EnvironmentType, TimeOfDay, CastingAnim, TargetAnim, TravelType, SpellAffectIndex, disallow_sit, field142, field143, new_icon, spellanim, uninterruptable, ResistDiff, dot_stacking_exempt, deleteable, RecourseLink, no_partial_resist, field152, field153, short_buff_box, descnum, typedescnum, effectdescnum, effectdescnum2, npc_no_los, field160, reflectable, bonushate, field163, field164, ldon_trap, EndurCost, EndurTimerIndex, IsDiscipline, field169, field170, field171, field172, HateAdded, EndurUpkeep, numhitstype, numhits, pvpresistbase, pvpresistcalc, pvpresistcap, spell_category, field181, field182, pcnpc_only_flag, cast_not_standing, can_mgb, nodispell, npc_category, npc_usefulness, MinResist, MaxResist, viral_targets, viral_timer, nimbuseffect, ConeStartAngle, ConeStopAngle, sneaking, not_extendable, field198, field199, suspendable, viral_range, songcap, field203, field204, no_block, field206, spellgroup, rank, field209, field210, CastRestriction, allowrest, InCombat, OutofCombat, field215, field216, field217, aemaxtargets, maxtargets, field220, field221, field222, field223, persistdeath, field225, field226, min_dist, min_dist_mod, max_dist, max_dist_mod, min_range, field232, field233, field234, field235, field236"
	spellBinds  = ":id, :name, :player_1, :teleport_zone, :you_cast, :other_casts, :cast_on_you, :cast_on_other, :spell_fades, :range, :aoerange, :pushback, :pushup, :cast_time, :recovery_time, :recast_time, :buffdurationformula, :buffduration, :AEDuration, :mana, :icon, :memicon, :LightType, :goodEffect, :Activated, :resisttype, :targettype, :basediff, :skill, :zonetype, :EnvironmentType, :TimeOfDay, :CastingAnim, :TargetAnim, :TravelType, :SpellAffectIndex, :disallow_sit, :field142, :field143, :new_icon, :spellanim, :uninterruptable, :ResistDiff, :dot_stacking_exempt, :deleteable, :RecourseLink, :no_partial_resist, :field152, :field153, :short_buff_box, :descnum, :typedescnum, :effectdescnum, :effectdescnum2, :npc_no_los, :field160, :reflectable, :bonushate, :field163, :field164, :ldon_trap, :EndurCost, :EndurTimerIndex, :IsDiscipline, :field169, :field170, :field171, :field172, :HateAdded, :EndurUpkeep, :numhitstype, :numhits, :pvpresistbase, :pvpresistcalc, :pvpresistcap, :spell_category, :field181, :field182, :pcnpc_only_flag, :cast_not_standing, :can_mgb, :nodispell, :npc_category, :npc_usefulness, :MinResist, :MaxResist, :viral_targets, :viral_timer, :nimbuseffect, :ConeStartAngle, :ConeStopAngle, :sneaking, :not_extendable, :field198, :field199, :suspendable, :viral_range, :songcap, :field203, :field204, :no_block, :field206, :spellgroup, :rank, :field209, :field210, :CastRestriction, :allowrest, :InCombat, :OutofCombat, :field215, :field216, :field217, :aemaxtargets, :maxtargets, :field220, :field221, :field222, :field223, :persistdeath, :field225, :field226, :min_dist, :min_dist_mod, :max_dist, :max_dist_mod, :min_range, :field232, :field233, :field234, :field235, field236"
)

//GetSpell will grab data from storage
func (s *Storage) GetSpell(spell *model.Spell) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", spellFields, spellTable)
	err = s.db.Get(spell, query, spell.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateSpell will grab data from storage
func (s *Storage) CreateSpell(spell *model.Spell) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", spellTable, spellFields, spellBinds)
	result, err := s.db.NamedExec(query, spell)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	spellID, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	spell.ID = spellID
	return
}

//ListSpell will grab data from storage
func (s *Storage) ListSpell(page *model.Page) (spells []*model.Spell, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", spellFields, spellTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		spell := model.Spell{}
		if err = rows.StructScan(&spell); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		spells = append(spells, &spell)
	}
	return
}

//ListSpellTotalCount will grab data from storage
func (s *Storage) ListSpellTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(id) FROM %s", spellTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListSpellBySearch will grab data from storage
func (s *Storage) ListSpellBySearch(page *model.Page, spell *model.Spell) (spells []*model.Spell, err error) {

	field := ""

	if len(spell.Name.String) > 0 {
		field += `name LIKE :name OR`
		spell.Name.String = fmt.Sprintf("%%%s%%", spell.Name.String)
		spell.Name.Valid = true
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", spellFields, spellTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, spell)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	fmt.Println(query)

	for rows.Next() {
		spell := model.Spell{}
		if err = rows.StructScan(&spell); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		spells = append(spells, &spell)
	}
	return
}

//ListSpellBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellBySearchTotalCount(spell *model.Spell) (count int64, err error) {
	field := ""
	if len(spell.Name.String) > 0 {
		field += `name LIKE :name OR`
		spell.Name.String = fmt.Sprintf("%%%s%%", spell.Name.String)
		spell.Name.Valid = true
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", spellTable, field)

	rows, err := s.db.NamedQuery(query, spell)
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

//EditSpell will grab data from storage
func (s *Storage) EditSpell(spell *model.Spell) (err error) {

	prevSpell := &model.Spell{
		ID: spell.ID,
	}
	err = s.GetSpell(prevSpell)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous spell")
		return
	}

	field := ""
	if len(spell.Name.String) > 0 && prevSpell.Name.String != spell.Name.String {
		field += "name = :name, "
	}
	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", spellTable, field)
	result, err := s.db.NamedExec(query, spell)
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

//DeleteSpell will grab data from storage
func (s *Storage) DeleteSpell(spell *model.Spell) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", spellTable)
	result, err := s.db.Exec(query, spell.ID)
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

//createTableSpell will grab data from storage
func (s *Storage) createTableSpell() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE spells_new (
  id int(11) NOT NULL DEFAULT '0',
  name varchar(64) DEFAULT NULL,
  player_1 varchar(64) DEFAULT 'BLUE_TRAIL',
  teleport_zone varchar(64) DEFAULT NULL,
  you_cast varchar(120) DEFAULT NULL,
  other_casts varchar(120) DEFAULT NULL,
  cast_on_you varchar(120) DEFAULT NULL,
  cast_on_other varchar(120) DEFAULT NULL,
  spell_fades varchar(120) DEFAULT NULL,
  range int(11) NOT NULL DEFAULT '100',
  aoerange int(11) NOT NULL DEFAULT '0',
  pushback int(11) NOT NULL DEFAULT '0',
  pushup int(11) NOT NULL DEFAULT '0',
  cast_time int(11) NOT NULL DEFAULT '0',
  recovery_time int(11) NOT NULL DEFAULT '0',
  recast_time int(11) NOT NULL DEFAULT '0',
  buffdurationformula int(11) NOT NULL DEFAULT '7',
  buffduration int(11) NOT NULL DEFAULT '65',
  AEDuration int(11) NOT NULL DEFAULT '0',
  mana int(11) NOT NULL DEFAULT '0',
  effect_base_value1 int(11) NOT NULL DEFAULT '100',
  effect_base_value2 int(11) NOT NULL DEFAULT '0',
  effect_base_value3 int(11) NOT NULL DEFAULT '0',
  effect_base_value4 int(11) NOT NULL DEFAULT '0',
  effect_base_value5 int(11) NOT NULL DEFAULT '0',
  effect_base_value6 int(11) NOT NULL DEFAULT '0',
  effect_base_value7 int(11) NOT NULL DEFAULT '0',
  effect_base_value8 int(11) NOT NULL DEFAULT '0',
  effect_base_value9 int(11) NOT NULL DEFAULT '0',
  effect_base_value10 int(11) NOT NULL DEFAULT '0',
  effect_base_value11 int(11) NOT NULL DEFAULT '0',
  effect_base_value12 int(11) NOT NULL DEFAULT '0',
  effect_limit_value1 int(11) NOT NULL DEFAULT '0',
  effect_limit_value2 int(11) NOT NULL DEFAULT '0',
  effect_limit_value3 int(11) NOT NULL DEFAULT '0',
  effect_limit_value4 int(11) NOT NULL DEFAULT '0',
  effect_limit_value5 int(11) NOT NULL DEFAULT '0',
  effect_limit_value6 int(11) NOT NULL DEFAULT '0',
  effect_limit_value7 int(11) NOT NULL DEFAULT '0',
  effect_limit_value8 int(11) NOT NULL DEFAULT '0',
  effect_limit_value9 int(11) NOT NULL DEFAULT '0',
  effect_limit_value10 int(11) NOT NULL DEFAULT '0',
  effect_limit_value11 int(11) NOT NULL DEFAULT '0',
  effect_limit_value12 int(11) NOT NULL DEFAULT '0',
  max1 int(11) NOT NULL DEFAULT '0',
  max2 int(11) NOT NULL DEFAULT '0',
  max3 int(11) NOT NULL DEFAULT '0',
  max4 int(11) NOT NULL DEFAULT '0',
  max5 int(11) NOT NULL DEFAULT '0',
  max6 int(11) NOT NULL DEFAULT '0',
  max7 int(11) NOT NULL DEFAULT '0',
  max8 int(11) NOT NULL DEFAULT '0',
  max9 int(11) NOT NULL DEFAULT '0',
  max10 int(11) NOT NULL DEFAULT '0',
  max11 int(11) NOT NULL DEFAULT '0',
  max12 int(11) NOT NULL DEFAULT '0',
  icon int(11) NOT NULL DEFAULT '0',
  memicon int(11) NOT NULL DEFAULT '0',
  components1 int(11) NOT NULL DEFAULT '-1',
  components2 int(11) NOT NULL DEFAULT '-1',
  components3 int(11) NOT NULL DEFAULT '-1',
  components4 int(11) NOT NULL DEFAULT '-1',
  component_counts1 int(11) NOT NULL DEFAULT '1',
  component_counts2 int(11) NOT NULL DEFAULT '1',
  component_counts3 int(11) NOT NULL DEFAULT '1',
  component_counts4 int(11) NOT NULL DEFAULT '1',
  NoexpendReagent1 int(11) NOT NULL DEFAULT '-1',
  NoexpendReagent2 int(11) NOT NULL DEFAULT '-1',
  NoexpendReagent3 int(11) NOT NULL DEFAULT '-1',
  NoexpendReagent4 int(11) NOT NULL DEFAULT '-1',
  formula1 int(11) NOT NULL DEFAULT '100',
  formula2 int(11) NOT NULL DEFAULT '100',
  formula3 int(11) NOT NULL DEFAULT '100',
  formula4 int(11) NOT NULL DEFAULT '100',
  formula5 int(11) NOT NULL DEFAULT '100',
  formula6 int(11) NOT NULL DEFAULT '100',
  formula7 int(11) NOT NULL DEFAULT '100',
  formula8 int(11) NOT NULL DEFAULT '100',
  formula9 int(11) NOT NULL DEFAULT '100',
  formula10 int(11) NOT NULL DEFAULT '100',
  formula11 int(11) NOT NULL DEFAULT '100',
  formula12 int(11) NOT NULL DEFAULT '100',
  LightType int(11) NOT NULL DEFAULT '0',
  goodEffect int(11) NOT NULL DEFAULT '0',
  Activated int(11) NOT NULL DEFAULT '0',
  resisttype int(11) NOT NULL DEFAULT '0',
  effectid1 int(11) NOT NULL DEFAULT '254',
  effectid2 int(11) NOT NULL DEFAULT '254',
  effectid3 int(11) NOT NULL DEFAULT '254',
  effectid4 int(11) NOT NULL DEFAULT '254',
  effectid5 int(11) NOT NULL DEFAULT '254',
  effectid6 int(11) NOT NULL DEFAULT '254',
  effectid7 int(11) NOT NULL DEFAULT '254',
  effectid8 int(11) NOT NULL DEFAULT '254',
  effectid9 int(11) NOT NULL DEFAULT '254',
  effectid10 int(11) NOT NULL DEFAULT '254',
  effectid11 int(11) NOT NULL DEFAULT '254',
  effectid12 int(11) NOT NULL DEFAULT '254',
  targettype int(11) NOT NULL DEFAULT '2',
  basediff int(11) NOT NULL DEFAULT '0',
  skill int(11) NOT NULL DEFAULT '98',
  zonetype int(11) NOT NULL DEFAULT '-1',
  EnvironmentType int(11) NOT NULL DEFAULT '0',
  TimeOfDay int(11) NOT NULL DEFAULT '0',
  classes1 int(11) NOT NULL DEFAULT '255',
  classes2 int(11) NOT NULL DEFAULT '255',
  classes3 int(11) NOT NULL DEFAULT '255',
  classes4 int(11) NOT NULL DEFAULT '255',
  classes5 int(11) NOT NULL DEFAULT '255',
  classes6 int(11) NOT NULL DEFAULT '255',
  classes7 int(11) NOT NULL DEFAULT '255',
  classes8 int(11) NOT NULL DEFAULT '255',
  classes9 int(11) NOT NULL DEFAULT '255',
  classes10 int(11) NOT NULL DEFAULT '255',
  classes11 int(11) NOT NULL DEFAULT '255',
  classes12 int(11) NOT NULL DEFAULT '255',
  classes13 int(11) NOT NULL DEFAULT '255',
  classes14 int(11) NOT NULL DEFAULT '255',
  classes15 int(11) NOT NULL DEFAULT '255',
  classes16 int(11) NOT NULL DEFAULT '255',
  CastingAnim int(11) NOT NULL DEFAULT '44',
  TargetAnim int(11) NOT NULL DEFAULT '13',
  TravelType int(11) NOT NULL DEFAULT '0',
  SpellAffectIndex int(11) NOT NULL DEFAULT '-1',
  disallow_sit int(11) NOT NULL DEFAULT '0',
  deities0 int(11) NOT NULL DEFAULT '0',
  deities1 int(11) NOT NULL DEFAULT '0',
  deities2 int(11) NOT NULL DEFAULT '0',
  deities3 int(11) NOT NULL DEFAULT '0',
  deities4 int(11) NOT NULL DEFAULT '0',
  deities5 int(11) NOT NULL DEFAULT '0',
  deities6 int(11) NOT NULL DEFAULT '0',
  deities7 int(11) NOT NULL DEFAULT '0',
  deities8 int(11) NOT NULL DEFAULT '0',
  deities9 int(11) NOT NULL DEFAULT '0',
  deities10 int(11) NOT NULL DEFAULT '0',
  deities11 int(11) NOT NULL DEFAULT '0',
  deities12 int(12) NOT NULL DEFAULT '0',
  deities13 int(11) NOT NULL DEFAULT '0',
  deities14 int(11) NOT NULL DEFAULT '0',
  deities15 int(11) NOT NULL DEFAULT '0',
  deities16 int(11) NOT NULL DEFAULT '0',
  field142 int(11) NOT NULL DEFAULT '100',
  field143 int(11) NOT NULL DEFAULT '0',
  new_icon int(11) NOT NULL DEFAULT '161',
  spellanim int(11) NOT NULL DEFAULT '0',
  uninterruptable int(11) NOT NULL DEFAULT '0',
  ResistDiff int(11) NOT NULL DEFAULT '-150',
  dot_stacking_exempt int(11) NOT NULL DEFAULT '0',
  deleteable int(11) NOT NULL DEFAULT '0',
  RecourseLink int(11) NOT NULL DEFAULT '0',
  no_partial_resist int(11) NOT NULL DEFAULT '0',
  field152 int(11) NOT NULL DEFAULT '0',
  field153 int(11) NOT NULL DEFAULT '0',
  short_buff_box int(11) NOT NULL DEFAULT '-1',
  descnum int(11) NOT NULL DEFAULT '0',
  typedescnum int(11) DEFAULT NULL,
  effectdescnum int(11) DEFAULT NULL,
  effectdescnum2 int(11) NOT NULL DEFAULT '0',
  npc_no_los int(11) NOT NULL DEFAULT '0',
  field160 int(11) NOT NULL DEFAULT '0',
  reflectable int(11) NOT NULL DEFAULT '0',
  bonushate int(11) NOT NULL DEFAULT '0',
  field163 int(11) NOT NULL DEFAULT '100',
  field164 int(11) NOT NULL DEFAULT '-150',
  ldon_trap int(11) NOT NULL DEFAULT '0',
  EndurCost int(11) NOT NULL DEFAULT '0',
  EndurTimerIndex int(11) NOT NULL DEFAULT '0',
  IsDiscipline int(11) NOT NULL DEFAULT '0',
  field169 int(11) NOT NULL DEFAULT '0',
  field170 int(11) NOT NULL DEFAULT '0',
  field171 int(11) NOT NULL DEFAULT '0',
  field172 int(11) NOT NULL DEFAULT '0',
  HateAdded int(11) NOT NULL DEFAULT '0',
  EndurUpkeep int(11) NOT NULL DEFAULT '0',
  numhitstype int(11) NOT NULL DEFAULT '0',
  numhits int(11) NOT NULL DEFAULT '0',
  pvpresistbase int(11) NOT NULL DEFAULT '-150',
  pvpresistcalc int(11) NOT NULL DEFAULT '100',
  pvpresistcap int(11) NOT NULL DEFAULT '-150',
  spell_category int(11) NOT NULL DEFAULT '-99',
  field181 int(11) NOT NULL DEFAULT '7',
  field182 int(11) NOT NULL DEFAULT '65',
  pcnpc_only_flag int(11) DEFAULT '0',
  cast_not_standing int(11) DEFAULT '0',
  can_mgb int(11) NOT NULL DEFAULT '0',
  nodispell int(11) NOT NULL DEFAULT '-1',
  npc_category int(11) NOT NULL DEFAULT '0',
  npc_usefulness int(11) NOT NULL DEFAULT '0',
  MinResist int(11) NOT NULL DEFAULT '0',
  MaxResist int(11) NOT NULL DEFAULT '0',
  viral_targets int(11) NOT NULL DEFAULT '0',
  viral_timer int(11) NOT NULL DEFAULT '0',
  nimbuseffect int(11) DEFAULT '0',
  ConeStartAngle int(11) NOT NULL DEFAULT '0',
  ConeStopAngle int(11) NOT NULL DEFAULT '0',
  sneaking int(11) NOT NULL DEFAULT '0',
  not_extendable int(11) NOT NULL DEFAULT '0',
  field198 int(11) NOT NULL DEFAULT '0',
  field199 int(11) NOT NULL DEFAULT '1',
  suspendable int(11) DEFAULT '0',
  viral_range int(11) NOT NULL DEFAULT '0',
  songcap int(11) DEFAULT '0',
  field203 int(11) DEFAULT '0',
  field204 int(11) DEFAULT '0',
  no_block int(11) NOT NULL DEFAULT '0',
  field206 int(11) DEFAULT '-1',
  spellgroup int(11) DEFAULT '0',
  rank int(11) NOT NULL DEFAULT '0',
  field209 int(11) DEFAULT '0',
  field210 int(11) DEFAULT '1',
  CastRestriction int(11) NOT NULL DEFAULT '0',
  allowrest int(11) DEFAULT '0',
  InCombat int(11) NOT NULL DEFAULT '0',
  OutofCombat int(11) NOT NULL DEFAULT '0',
  field215 int(11) DEFAULT '0',
  field216 int(11) DEFAULT '0',
  field217 int(11) DEFAULT '0',
  aemaxtargets int(11) NOT NULL DEFAULT '0',
  maxtargets int(11) DEFAULT '0',
  field220 int(11) DEFAULT '0',
  field221 int(11) DEFAULT '0',
  field222 int(11) DEFAULT '0',
  field223 int(11) DEFAULT '0',
  persistdeath int(11) DEFAULT '0',
  field225 int(11) NOT NULL DEFAULT '0',
  field226 int(11) NOT NULL DEFAULT '0',
  min_dist float NOT NULL DEFAULT '0',
  min_dist_mod float NOT NULL DEFAULT '0',
  max_dist float NOT NULL DEFAULT '0',
  max_dist_mod float NOT NULL DEFAULT '0',
  min_range int(11) NOT NULL DEFAULT '0',
  field232 int(11) NOT NULL DEFAULT '0',
  field233 int(11) NOT NULL DEFAULT '0',
  field234 int(11) NOT NULL DEFAULT '0',
  field235 int(11) NOT NULL DEFAULT '0',
  field236 int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (id)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
