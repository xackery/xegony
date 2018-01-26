package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	spellFields = `spells_new.id, spells_new.name, spells_new.player_1, spells_new.teleport_zone, spells_new.you_cast, spells_new.other_casts, spells_new.cast_on_you, spells_new.cast_on_other, spells_new.spell_fades, spells_new.range, spells_new.aoerange, spells_new.pushback, spells_new.pushup, spells_new.cast_time, spells_new.recovery_time, spells_new.recast_time, spells_new.buffdurationformula, spells_new.buffduration, spells_new.AEDuration, spells_new.mana, spells_new.effect_base_value1, spells_new.effect_base_value2, spells_new.effect_base_value3, spells_new.effect_base_value4, spells_new.effect_base_value5, spells_new.effect_base_value6, spells_new.effect_base_value7, spells_new.effect_base_value8, spells_new.effect_base_value9, spells_new.effect_base_value10, spells_new.effect_base_value11, spells_new.effect_base_value12, spells_new.effect_limit_value1, spells_new.effect_limit_value2, spells_new.effect_limit_value3, spells_new.effect_limit_value4, spells_new.effect_limit_value5, spells_new.effect_limit_value6, spells_new.effect_limit_value7, spells_new.effect_limit_value8, spells_new.effect_limit_value9, spells_new.effect_limit_value10, spells_new.effect_limit_value11, spells_new.effect_limit_value12, spells_new.max1, spells_new.max2, spells_new.max3, spells_new.max4, spells_new.max5, spells_new.max6, spells_new.max7, spells_new.max8, spells_new.max9, spells_new.max10, spells_new.max11, spells_new.max12, spells_new.icon, spells_new.memicon, spells_new.components1, spells_new.components2, spells_new.components3, spells_new.components4, spells_new.component_counts1, spells_new.component_counts2, spells_new.component_counts3, spells_new.component_counts4, spells_new.NoexpendReagent1, spells_new.NoexpendReagent2, spells_new.NoexpendReagent3, spells_new.NoexpendReagent4, spells_new.formula1, spells_new.formula2, spells_new.formula3, spells_new.formula4, spells_new.formula5, spells_new.formula6, spells_new.formula7, spells_new.formula8, spells_new.formula9, spells_new.formula10, spells_new.formula11, spells_new.formula12, spells_new.LightType, spells_new.goodEffect, spells_new.Activated, spells_new.resisttype, spells_new.effectid1, spells_new.effectid2, spells_new.effectid3, spells_new.effectid4, spells_new.effectid5, spells_new.effectid6, spells_new.effectid7, spells_new.effectid8, spells_new.effectid9, spells_new.effectid10, spells_new.effectid11, spells_new.effectid12, spells_new.targettype, spells_new.basediff, spells_new.skill, spells_new.zonetype, spells_new.EnvironmentType, spells_new.TimeOfDay, spells_new.classes1, spells_new.classes2, spells_new.classes3, spells_new.classes4, spells_new.classes5, spells_new.classes6, spells_new.classes7, spells_new.classes8, spells_new.classes9, spells_new.classes10, spells_new.classes11, spells_new.classes12, spells_new.classes13, spells_new.classes14, spells_new.classes15, spells_new.classes16, spells_new.CastingAnim, spells_new.TargetAnim, spells_new.TravelType, spells_new.SpellAffectIndex, spells_new.disallow_sit, spells_new.deities0, spells_new.deities1, spells_new.deities2, spells_new.deities3, spells_new.deities4, spells_new.deities5, spells_new.deities6, spells_new.deities7, spells_new.deities8, spells_new.deities9, spells_new.deities10, spells_new.deities11, spells_new.deities12, spells_new.deities13, spells_new.deities14, spells_new.deities15, spells_new.deities16, spells_new.field142, spells_new.field143, spells_new.new_icon, spells_new.spellanim, spells_new.uninterruptable, spells_new.ResistDiff, spells_new.dot_stacking_exempt, spells_new.deleteable, spells_new.RecourseLink, spells_new.no_partial_resist, spells_new.field152, spells_new.field153, spells_new.short_buff_box, spells_new.descnum, spells_new.typedescnum, spells_new.effectdescnum, spells_new.effectdescnum2, spells_new.npc_no_los, spells_new.field160, spells_new.reflectable, spells_new.bonushate, spells_new.field163, spells_new.field164, spells_new.ldon_trap, spells_new.EndurCost, spells_new.EndurTimerIndex, spells_new.IsDiscipline, spells_new.field169, spells_new.field170, spells_new.field171, spells_new.field172, spells_new.HateAdded, spells_new.EndurUpkeep, spells_new.numhitstype, spells_new.numhits, spells_new.pvpresistbase, spells_new.pvpresistcalc, spells_new.pvpresistcap, spells_new.spell_category, spells_new.field181, spells_new.field182, spells_new.pcnpc_only_flag, spells_new.cast_not_standing, spells_new.can_mgb, spells_new.nodispell, spells_new.npc_category, spells_new.npc_usefulness, spells_new.MinResist, spells_new.MaxResist, spells_new.viral_targets, spells_new.viral_timer, spells_new.nimbuseffect, spells_new.ConeStartAngle, spells_new.ConeStopAngle, spells_new.sneaking, spells_new.not_extendable, spells_new.field198, spells_new.field199, spells_new.suspendable, spells_new.viral_range, spells_new.songcap, spells_new.field203, spells_new.field204, spells_new.no_block, spells_new.field206, spells_new.spellgroup, spells_new.rank, spells_new.field209, spells_new.field210, spells_new.CastRestriction, spells_new.allowrest, spells_new.InCombat, spells_new.OutofCombat, spells_new.field215, spells_new.field216, spells_new.field217, spells_new.aemaxtargets, spells_new.maxtargets, spells_new.field220, spells_new.field221, spells_new.field222, spells_new.field223, spells_new.persistdeath, spells_new.field225, spells_new.field226, spells_new.min_dist, spells_new.min_dist_mod, spells_new.max_dist, spells_new.max_dist_mod, spells_new.min_range, spells_new.field232, spells_new.field233, spells_new.field234, spells_new.field235, spells_new.field236`
	spellSets   = `id=:id, name=:name, player_1=:player_1, teleport_zone=:teleport_zone, you_cast=:you_cast, other_casts=:other_casts, cast_on_you=:cast_on_you, cast_on_other=:cast_on_other, spell_fades=:spell_fades, range=:range, aoerange=:aoerange, pushback=:pushback, pushup=:pushup, cast_time=:cast_time, recovery_time=:recovery_time, recast_time=:recast_time, buffdurationformula=:buffdurationformula, buffduration=:buffduration, AEDuration=:AEDuration, mana=:mana, effect_base_value1=:effect_base_value1, effect_base_value2=:effect_base_value2, effect_base_value3=:effect_base_value3, effect_base_value4=:effect_base_value4, effect_base_value5=:effect_base_value5, effect_base_value6=:effect_base_value6, effect_base_value7=:effect_base_value7, effect_base_value8=:effect_base_value8, effect_base_value9=:effect_base_value9, effect_base_value10=:effect_base_value10, effect_base_value11=:effect_base_value11, effect_base_value12=:effect_base_value12, effect_limit_value1=:effect_limit_value1, effect_limit_value2=:effect_limit_value2, effect_limit_value3=:effect_limit_value3, effect_limit_value4=:effect_limit_value4, effect_limit_value5=:effect_limit_value5, effect_limit_value6=:effect_limit_value6, effect_limit_value7=:effect_limit_value7, effect_limit_value8=:effect_limit_value8, effect_limit_value9=:effect_limit_value9, effect_limit_value10=:effect_limit_value10, effect_limit_value11=:effect_limit_value11, effect_limit_value12=:effect_limit_value12, max1=:max1, max2=:max2, max3=:max3, max4=:max4, max5=:max5, max6=:max6, max7=:max7, max8=:max8, max9=:max9, max10=:max10, max11=:max11, max12=:max12, icon=:icon, memicon=:memicon, components1=:components1, components2=:components2, components3=:components3, components4=:components4, component_counts1=:component_counts1, component_counts2=:component_counts2, component_counts3=:component_counts3, component_counts4=:component_counts4, NoexpendReagent1=:NoexpendReagent1, NoexpendReagent2=:NoexpendReagent2, NoexpendReagent3=:NoexpendReagent3, NoexpendReagent4=:NoexpendReagent4, formula1=:formula1, formula2=:formula2, formula3=:formula3, formula4=:formula4, formula5=:formula5, formula6=:formula6, formula7=:formula7, formula8=:formula8, formula9=:formula9, formula10=:formula10, formula11=:formula11, formula12=:formula12, LightType=:LightType, goodEffect=:goodEffect, Activated=:Activated, resisttype=:resisttype, effectid1=:effectid1, effectid2=:effectid2, effectid3=:effectid3, effectid4=:effectid4, effectid5=:effectid5, effectid6=:effectid6, effectid7=:effectid7, effectid8=:effectid8, effectid9=:effectid9, effectid10=:effectid10, effectid11=:effectid11, effectid12=:effectid12, targettype=:targettype, basediff=:basediff, skill=:skill, zonetype=:zonetype, EnvironmentType=:EnvironmentType, TimeOfDay=:TimeOfDay, classes1=:classes1, classes2=:classes2, classes3=:classes3, classes4=:classes4, classes5=:classes5, classes6=:classes6, classes7=:classes7, classes8=:classes8, classes9=:classes9, classes10=:classes10, classes11=:classes11, classes12=:classes12, classes13=:classes13, classes14=:classes14, classes15=:classes15, classes16=:classes16, CastingAnim=:CastingAnim, TargetAnim=:TargetAnim, TravelType=:TravelType, SpellAffectIndex=:SpellAffectIndex, disallow_sit=:disallow_sit, deities0=:deities0, deities1=:deities1, deities2=:deities2, deities3=:deities3, deities4=:deities4, deities5=:deities5, deities6=:deities6, deities7=:deities7, deities8=:deities8, deities9=:deities9, deities10=:deities10, deities11=:deities11, deities12=:deities12, deities13=:deities13, deities14=:deities14, deities15=:deities15, deities16=:deities16, field142=:field142, field143=:field143, new_icon=:new_icon, spellanim=:spellanim, uninterruptable=:uninterruptable, ResistDiff=:ResistDiff, dot_stacking_exempt=:dot_stacking_exempt, deleteable=:deleteable, RecourseLink=:RecourseLink, no_partial_resist=:no_partial_resist, field152=:field152, field153=:field153, short_buff_box=:short_buff_box, descnum=:descnum, typedescnum=:typedescnum, effectdescnum=:effectdescnum, effectdescnum2=:effectdescnum2, npc_no_los=:npc_no_los, field160=:field160, reflectable=:reflectable, bonushate=:bonushate, field163=:field163, field164=:field164, ldon_trap=:ldon_trap, EndurCost=:EndurCost, EndurTimerIndex=:EndurTimerIndex, IsDiscipline=:IsDiscipline, field169=:field169, field170=:field170, field171=:field171, field172=:field172, HateAdded=:HateAdded, EndurUpkeep=:EndurUpkeep, numhitstype=:numhitstype, numhits=:numhits, pvpresistbase=:pvpresistbase, pvpresistcalc=:pvpresistcalc, pvpresistcap=:pvpresistcap, spell_category=:spell_category, field181=:field181, field182=:field182, pcnpc_only_flag=:pcnpc_only_flag, cast_not_standing=:cast_not_standing, can_mgb=:can_mgb, nodispell=:nodispell, npc_category=:npc_category, npc_usefulness=:npc_usefulness, MinResist=:MinResist, MaxResist=:MaxResist, viral_targets=:viral_targets, viral_timer=:viral_timer, nimbuseffect=:nimbuseffect, ConeStartAngle=:ConeStartAngle, ConeStopAngle=:ConeStopAngle, sneaking=:sneaking, not_extendable=:not_extendable, field198=:field198, field199=:field199, suspendable=:suspendable, viral_range=:viral_range, songcap=:songcap, field203=:field203, field204=:field204, no_block=:no_block, field206=:field206, spellgroup=:spellgroup, rank=:rank, field209=:field209, field210=:field210, CastRestriction=:CastRestriction, allowrest=:allowrest, InCombat=:InCombat, OutofCombat=:OutofCombat, field215=:field215, field216=:field216, field217=:field217, aemaxtargets=:aemaxtargets, maxtargets=:maxtargets, field220=:field220, field221=:field221, field222=:field222, field223=:field223, persistdeath=:persistdeath, field225=:field225, field226=:field226, min_dist=:min_dist, min_dist_mod=:min_dist_mod, max_dist=:max_dist, max_dist_mod=:max_dist_mod, min_range=:min_range, field232=:field232, field233=:field233, field234=:field234, field235=:field235, field236=:field236`
	spellBinds  = `:id, :name, :player_1, :teleport_zone, :you_cast, :other_casts, :cast_on_you, :cast_on_other, :spell_fades, :range, :aoerange, :pushback, :pushup, :cast_time, :recovery_time, :recast_time, :buffdurationformula, :buffduration, :AEDuration, :mana, :effect_base_value1, :effect_base_value2, :effect_base_value3, :effect_base_value4, :effect_base_value5, :effect_base_value6, :effect_base_value7, :effect_base_value8, :effect_base_value9, :effect_base_value10, :effect_base_value11, :effect_base_value12, :effect_limit_value1, :effect_limit_value2, :effect_limit_value3, :effect_limit_value4, :effect_limit_value5, :effect_limit_value6, :effect_limit_value7, :effect_limit_value8, :effect_limit_value9, :effect_limit_value10, :effect_limit_value11, :effect_limit_value12, :max1, :max2, :max3, :max4, :max5, :max6, :max7, :max8, :max9, :max10, :max11, :max12, :icon, :memicon, :components1, :components2, :components3, :components4, :component_counts1, :component_counts2, :component_counts3, :component_counts4, :NoexpendReagent1, :NoexpendReagent2, :NoexpendReagent3, :NoexpendReagent4, :formula1, :formula2, :formula3, :formula4, :formula5, :formula6, :formula7, :formula8, :formula9, :formula10, :formula11, :formula12, :LightType, :goodEffect, :Activated, :resisttype, :effectid1, :effectid2, :effectid3, :effectid4, :effectid5, :effectid6, :effectid7, :effectid8, :effectid9, :effectid10, :effectid11, :effectid12, :targettype, :basediff, :skill, :zonetype, :EnvironmentType, :TimeOfDay, :classes1, :classes2, :classes3, :classes4, :classes5, :classes6, :classes7, :classes8, :classes9, :classes10, :classes11, :classes12, :classes13, :classes14, :classes15, :classes16, :CastingAnim, :TargetAnim, :TravelType, :SpellAffectIndex, :disallow_sit, :deities0, :deities1, :deities2, :deities3, :deities4, :deities5, :deities6, :deities7, :deities8, :deities9, :deities10, :deities11, :deities12, :deities13, :deities14, :deities15, :deities16, :field142, :field143, :new_icon, :spellanim, :uninterruptable, :ResistDiff, :dot_stacking_exempt, :deleteable, :RecourseLink, :no_partial_resist, :field152, :field153, :short_buff_box, :descnum, :typedescnum, :effectdescnum, :effectdescnum2, :npc_no_los, :field160, :reflectable, :bonushate, :field163, :field164, :ldon_trap, :EndurCost, :EndurTimerIndex, :IsDiscipline, :field169, :field170, :field171, :field172, :HateAdded, :EndurUpkeep, :numhitstype, :numhits, :pvpresistbase, :pvpresistcalc, :pvpresistcap, :spell_category, :field181, :field182, :pcnpc_only_flag, :cast_not_standing, :can_mgb, :nodispell, :npc_category, :npc_usefulness, :MinResist, :MaxResist, :viral_targets, :viral_timer, :nimbuseffect, :ConeStartAngle, :ConeStopAngle, :sneaking, :not_extendable, :field198, :field199, :suspendable, :viral_range, :songcap, :field203, :field204, :no_block, :field206, :spellgroup, :rank, :field209, :field210, :CastRestriction, :allowrest, :InCombat, :OutofCombat, :field215, :field216, :field217, :aemaxtargets, :maxtargets, :field220, :field221, :field222, :field223, :persistdeath, :field225, :field226, :min_dist, :min_dist_mod, :max_dist, :max_dist_mod, :min_range, :field232, :field233, :field234, :field235, spells_new.field236`
)

//GetSpell will grab data from storage
func (s *Storage) GetSpell(spell *model.Spell) (err error) {
	err = s.db.Get(spell, fmt.Sprintf("SELECT id, %s FROM spells_new WHERE id = ?", spellFields), spell.ID)
	if err != nil {
		return
	}
	return
}

//CreateSpell will grab data from storage
func (s *Storage) CreateSpell(spell *model.Spell) (err error) {
	if spell == nil {
		err = fmt.Errorf("Must provide spell")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO spells_new(%s)
		VALUES (%s)`, spellFields, spellBinds), spell)
	if err != nil {
		return
	}
	spellID, err := result.LastInsertId()
	if err != nil {
		return
	}
	spell.ID = spellID
	return
}

//ListSpell will grab data from storage
func (s *Storage) ListSpell(pageSize int64, pageNumber int64) (spells []*model.Spell, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM spells_new 
		ORDER BY id ASC LIMIT %d OFFSET %d`, spellFields, pageSize, pageSize*pageNumber))
	if err != nil {
		return
	}

	for rows.Next() {
		spell := model.Spell{}
		if err = rows.StructScan(&spell); err != nil {
			return
		}
		spells = append(spells, &spell)
	}
	return
}

//ListSpellCount will grab data from storage
func (s *Storage) ListSpellCount() (count int64, err error) {
	err = s.db.Get(&count, `SELECT count(id) FROM spells_new`)
	if err != nil {
		return
	}
	return
}

//SearchSpellByName will grab data from storage
func (s *Storage) SearchSpellByName(spell *model.Spell) (spells []*model.Spell, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM spells_new 
		WHERE name like ? ORDER BY id DESC`, spellFields), "%"+spell.Name.String+"%")

	if err != nil {
		return
	}

	for rows.Next() {
		spell := model.Spell{}
		if err = rows.StructScan(&spell); err != nil {
			return
		}
		spells = append(spells, &spell)
	}
	return
}

//EditSpell will grab data from storage
func (s *Storage) EditSpell(spell *model.Spell) (err error) {
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE spells_new SET %s WHERE id = :id`, spellSets), spell)
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

//DeleteSpell will grab data from storage
func (s *Storage) DeleteSpell(spell *model.Spell) (err error) {
	result, err := s.db.Exec(`DELETE FROM spells_new WHERE id = ?`, spell.ID)
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
