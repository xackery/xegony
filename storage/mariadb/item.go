package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const itemFields = `name, aagi, ac, accuracy, acha, adex, aint, artifactflag, asta, astr, attack, augrestrict, augslot1type, augslot1visible, augslot2type, augslot2visible, augslot3type, augslot3visible, augslot4type, augslot4visible, augslot5type, augslot5visible, augslot6type, augslot6visible, augtype, avoidance, awis, bagsize, bagslots, bagtype, bagwr, banedmgamt, banedmgraceamt, banedmgbody, banedmgrace, bardtype, bardvalue, book, casttime, casttime_, charmfile, charmfileid, classes, color, combateffects, extradmgskill, extradmgamt, price, cr, damage, damageshield, deity, delay, augdistiller, dotshielding, dr, clicktype, clicklevel2, elemdmgtype, elemdmgamt, endur, factionamt1, factionamt2, factionamt3, factionamt4, factionmod1, factionmod2, factionmod3, factionmod4, filename, focuseffect, fr, fvnodrop, haste, clicklevel, hp, regen, icon, idfile, itemclass, itemtype, ldonprice, ldontheme, ldonsold, light, lore, loregroup, magic, mana, manaregen, enduranceregen, material, herosforgemodel, maxcharges, mr, nodrop, norent, pendingloreflag, pr, procrate, races, ` + "`range`" + `, reclevel, recskill, reqlevel, sellrate, shielding, size, skillmodtype, skillmodvalue, slots, clickeffect, spellshield, strikethrough, stunresist, summonedflag, tradeskills, favor, weight, UNK012, UNK013, benefitflag, UNK054, UNK059, booktype, recastdelay, recasttype, guildfavor, UNK123, UNK124, attuneable, nopet, updated, comment, UNK127, pointtype, potionbelt, potionbeltslots, stacksize, notransfer, stackable, UNK134, UNK137, proceffect, proctype, proclevel2, proclevel, UNK142, worneffect, worntype, wornlevel2, wornlevel, UNK147, focustype, focuslevel2, focuslevel, UNK152, scrolleffect, scrolltype, scrolllevel2, scrolllevel, UNK157, serialized, verified, serialization, source, UNK033, lorefile, UNK014, svcorruption, skillmodmax, UNK060, augslot1unk2, augslot2unk2, augslot3unk2, augslot4unk2, augslot5unk2, augslot6unk2, UNK120, UNK121, questitemflag, UNK132, clickunk5, clickunk6, clickunk7, procunk1, procunk2, procunk3, procunk4, procunk6, procunk7, wornunk1, wornunk2, wornunk3, wornunk4, wornunk5, wornunk6, wornunk7, focusunk1, focusunk2, focusunk3, focusunk4, focusunk5, focusunk6, focusunk7, scrollunk1, scrollunk2, scrollunk3, scrollunk4, scrollunk5, scrollunk6, scrollunk7, UNK193, purity, evoitem, evoid, evolvinglevel, evomax, clickname, procname, wornname, focusname, scrollname, dsmitigation, heroic_str, heroic_int, heroic_wis, heroic_agi, heroic_dex, heroic_sta, heroic_cha, heroic_pr, heroic_dr, heroic_fr, heroic_cr, heroic_mr, heroic_svcorrup, healamt, spelldmg, clairvoyance, backstabdmg, created, elitematerial, ldonsellbackrate, scriptfileid, expendablearrow, powersourcecapacity, bardeffect, bardeffecttype, bardlevel2, bardlevel, bardunk1, bardunk2, bardunk3, bardunk4, bardunk5, bardname, bardunk7, UNK214, UNK219, UNK220, UNK221, heirloom, UNK223, UNK224, UNK225, UNK226, UNK227, UNK228, UNK229, UNK230, UNK231, UNK232, UNK233, UNK234, placeable, UNK236, UNK237, UNK238, UNK239, UNK240, UNK241, epicitem`
const itemBinds = `:name, :aagi, :ac, :accuracy, :acha, :adex, :aint, :artifactflag, :asta, :astr, :attack, :augrestrict, :augslot1type, :augslot1visible, :augslot2type, :augslot2visible, :augslot3type, :augslot3visible, :augslot4type, :augslot4visible, :augslot5type, :augslot5visible, :augslot6type, :augslot6visible, :augtype, :avoidance, :awis, :bagsize, :bagslots, :bagtype, :bagwr, :banedmgamt, :banedmgraceamt, :banedmgbody, :banedmgrace, :bardtype, :bardvalue, :book, :casttime, :casttime_, :charmfile, :charmfileid, :classes, :color, :combateffects, :extradmgskill, :extradmgamt, :price, :cr, :damage, :damageshield, :deity, :delay, :augdistiller, :dotshielding, :dr, :clicktype, :clicklevel2, :elemdmgtype, :elemdmgamt, :endur, :factionamt1, :factionamt2, :factionamt3, :factionamt4, :factionmod1, :factionmod2, :factionmod3, :factionmod4, :filename, :focuseffect, :fr, :fvnodrop, :haste, :clicklevel, :hp, :regen, :icon, :idfile, :itemclass, :itemtype, :ldonprice, :ldontheme, :ldonsold, :light, :lore, :loregroup, :magic, :mana, :manaregen, :enduranceregen, :material, :herosforgemodel, :maxcharges, :mr, :nodrop, :norent, :pendingloreflag, :pr, :procrate, :races, :range, :reclevel, :recskill, :reqlevel, :sellrate, :shielding, :size, :skillmodtype, :skillmodvalue, :slots, :clickeffect, :spellshield, :strikethrough, :stunresist, :summonedflag, :tradeskills, :favor, :weight, :UNK012, :UNK013, :benefitflag, :UNK054, :UNK059, :booktype, :recastdelay, :recasttype, :guildfavor, :UNK123, :UNK124, :attuneable, :nopet, :updated, :comment, :UNK127, :pointtype, :potionbelt, :potionbeltslots, :stacksize, :notransfer, :stackable, :UNK134, :UNK137, :proceffect, :proctype, :proclevel2, :proclevel, :UNK142, :worneffect, :worntype, :wornlevel2, :wornlevel, :UNK147, :focustype, :focuslevel2, :focuslevel, :UNK152, :scrolleffect, :scrolltype, :scrolllevel2, :scrolllevel, :UNK157, :serialized, :verified, :serialization, :source, :UNK033, :lorefile, :UNK014, :svcorruption, :skillmodmax, :UNK060, :augslot1unk2, :augslot2unk2, :augslot3unk2, :augslot4unk2, :augslot5unk2, :augslot6unk2, :UNK120, :UNK121, :questitemflag, :UNK132, :clickunk5, :clickunk6, :clickunk7, :procunk1, :procunk2, :procunk3, :procunk4, :procunk6, :procunk7, :wornunk1, :wornunk2, :wornunk3, :wornunk4, :wornunk5, :wornunk6, :wornunk7, :focusunk1, :focusunk2, :focusunk3, :focusunk4, :focusunk5, :focusunk6, :focusunk7, :scrollunk1, :scrollunk2, :scrollunk3, :scrollunk4, :scrollunk5, :scrollunk6, :scrollunk7, :UNK193, :purity, :evoitem, :evoid, :evolvinglevel, :evomax, :clickname, :procname, :wornname, :focusname, :scrollname, :dsmitigation, :heroic_str, :heroic_int, :heroic_wis, :heroic_agi, :heroic_dex, :heroic_sta, :heroic_cha, :heroic_pr, :heroic_dr, :heroic_fr, :heroic_cr, :heroic_mr, :heroic_svcorrup, :healamt, :spelldmg, :clairvoyance, :backstabdmg, :created, :elitematerial, :ldonsellbackrate, :scriptfileid, :expendablearrow, :powersourcecapacity, :bardeffect, :bardeffecttype, :bardlevel2, :bardlevel, :bardunk1, :bardunk2, :bardunk3, :bardunk4, :bardunk5, :bardname, :bardunk7, :UNK214, :UNK219, :UNK220, :UNK221, :heirloom, :UNK223, :UNK224, :UNK225, :UNK226, :UNK227, :UNK228, :UNK229, :UNK230, :UNK231, :UNK232, :UNK233, :UNK234, :placeable, :UNK236, :UNK237, :UNK238, :UNK239, :UNK240, :UNK241, :epicitem`
const itemSet = `name=:name, aagi=:aagi, ac=:ac, accuracy=:accuracy, acha=:acha, adex=:adex, aint=:aint, artifactflag=:artifactflag, asta=:asta, astr=:astr, attack=:attack, augrestrict=:augrestrict, augslot1type=:augslot1type, augslot1visible=:augslot1visible, augslot2type=:augslot2type, augslot2visible=:augslot2visible, augslot3type=:augslot3type, augslot3visible=:augslot3visible, augslot4type=:augslot4type, augslot4visible=:augslot4visible, augslot5type=:augslot5type, augslot5visible=:augslot5visible, augslot6type=:augslot6type, augslot6visible=:augslot6visible, augtype=:augtype, avoidance=:avoidance, awis=:awis, bagsize=:bagsize, bagslots=:bagslots, bagtype=:bagtype, bagwr=:bagwr, banedmgamt=:banedmgamt, banedmgraceamt=:banedmgraceamt, banedmgbody=:banedmgbody, banedmgrace=:banedmgrace, bardtype=:bardtype, bardvalue=:bardvalue, book=:book, casttime=:casttime, casttime_=:casttime_, charmfile=:charmfile, charmfileid=:charmfileid, classes=:classes, color=:color, combateffects=:combateffects, extradmgskill=:extradmgskill, extradmgamt=:extradmgamt, price=:price, cr=:cr, damage=:damage, damageshield=:damageshield, deity=:deity, delay=:delay, augdistiller=:augdistiller, dotshielding=:dotshielding, dr=:dr, clicktype=:clicktype, clicklevel2=:clicklevel2, elemdmgtype=:elemdmgtype, elemdmgamt=:elemdmgamt, endur=:endur, factionamt1=:factionamt1, factionamt2=:factionamt2, factionamt3=:factionamt3, factionamt4=:factionamt4, factionmod1=:factionmod1, factionmod2=:factionmod2, factionmod3=:factionmod3, factionmod4=:factionmod4, filename=:filename, focuseffect=:focuseffect, fr=:fr, fvnodrop=:fvnodrop, haste=:haste, clicklevel=:clicklevel, hp=:hp, regen=:regen, icon=:icon, idfile=:idfile, itemclass=:itemclass, itemtype=:itemtype, ldonprice=:ldonprice, ldontheme=:ldontheme, ldonsold=:ldonsold, light=:light, lore=:lore, loregroup=:loregroup, magic=:magic, mana=:mana, manaregen=:manaregen, enduranceregen=:enduranceregen, material=:material, herosforgemodel=:herosforgemodel, maxcharges=:maxcharges, mr=:mr, nodrop=:nodrop, norent=:norent, pendingloreflag=:pendingloreflag, pr=:pr, procrate=:procrate, races=:races, ` + "`range`" + `=:range, reclevel=:reclevel, recskill=:recskill, reqlevel=:reqlevel, sellrate=:sellrate, shielding=:shielding, size=:size, skillmodtype=:skillmodtype, skillmodvalue=:skillmodvalue, slots=:slots, clickeffect=:clickeffect, spellshield=:spellshield, strikethrough=:strikethrough, stunresist=:stunresist, summonedflag=:summonedflag, tradeskills=:tradeskills, favor=:favor, weight=:weight, UNK012=:UNK012, UNK013=:UNK013, benefitflag=:benefitflag, UNK054=:UNK054, UNK059=:UNK059, booktype=:booktype, recastdelay=:recastdelay, recasttype=:recasttype, guildfavor=:guildfavor, UNK123=:UNK123, UNK124=:UNK124, attuneable=:attuneable, nopet=:nopet, updated=:updated, comment=:comment, UNK127=:UNK127, pointtype=:pointtype, potionbelt=:potionbelt, potionbeltslots=:potionbeltslots, stacksize=:stacksize, notransfer=:notransfer, stackable=:stackable, UNK134=:UNK134, UNK137=:UNK137, proceffect=:proceffect, proctype=:proctype, proclevel2=:proclevel2, proclevel=:proclevel, UNK142=:UNK142, worneffect=:worneffect, worntype=:worntype, wornlevel2=:wornlevel2, wornlevel=:wornlevel, UNK147=:UNK147, focustype=:focustype, focuslevel2=:focuslevel2, focuslevel=:focuslevel, UNK152=:UNK152, scrolleffect=:scrolleffect, scrolltype=:scrolltype, scrolllevel2=:scrolllevel2, scrolllevel=:scrolllevel, UNK157=:UNK157, serialized=:serialized, verified=:verified, serialization=:serialization, source=:source, UNK033=:UNK033, lorefile=:lorefile, UNK014=:UNK014, svcorruption=:svcorruption, skillmodmax=:skillmodmax, UNK060=:UNK060, augslot1unk2=:augslot1unk2, augslot2unk2=:augslot2unk2, augslot3unk2=:augslot3unk2, augslot4unk2=:augslot4unk2, augslot5unk2=:augslot5unk2, augslot6unk2=:augslot6unk2, UNK120=:UNK120, UNK121=:UNK121, questitemflag=:questitemflag, UNK132=:UNK132, clickunk5=:clickunk5, clickunk6=:clickunk6, clickunk7=:clickunk7, procunk1=:procunk1, procunk2=:procunk2, procunk3=:procunk3, procunk4=:procunk4, procunk6=:procunk6, procunk7=:procunk7, wornunk1=:wornunk1, wornunk2=:wornunk2, wornunk3=:wornunk3, wornunk4=:wornunk4, wornunk5=:wornunk5, wornunk6=:wornunk6, wornunk7=:wornunk7, focusunk1=:focusunk1, focusunk2=:focusunk2, focusunk3=:focusunk3, focusunk4=:focusunk4, focusunk5=:focusunk5, focusunk6=:focusunk6, focusunk7=:focusunk7, scrollunk1=:scrollunk1, scrollunk2=:scrollunk2, scrollunk3=:scrollunk3, scrollunk4=:scrollunk4, scrollunk5=:scrollunk5, scrollunk6=:scrollunk6, scrollunk7=:scrollunk7, UNK193=:UNK193, purity=:purity, evoitem=:evoitem, evoid=:evoid, evolvinglevel=:evolvinglevel, evomax=:evomax, clickname=:clickname, procname=:procname, wornname=:wornname, focusname=:focusname, scrollname=:scrollname, dsmitigation=:dsmitigation, heroic_str=:heroic_str, heroic_int=:heroic_int, heroic_wis=:heroic_wis, heroic_agi=:heroic_agi, heroic_dex=:heroic_dex, heroic_sta=:heroic_sta, heroic_cha=:heroic_cha, heroic_pr=:heroic_pr, heroic_dr=:heroic_dr, heroic_fr=:heroic_fr, heroic_cr=:heroic_cr, heroic_mr=:heroic_mr, heroic_svcorrup=:heroic_svcorrup, healamt=:healamt, spelldmg=:spelldmg, clairvoyance=:clairvoyance, backstabdmg=:backstabdmg, created=:created, elitematerial=:elitematerial, ldonsellbackrate=:ldonsellbackrate, scriptfileid=:scriptfileid, expendablearrow=:expendablearrow, powersourcecapacity=:powersourcecapacity, bardeffect=:bardeffect, bardeffecttype=:bardeffecttype, bardlevel2=:bardlevel2, bardlevel=:bardlevel, bardunk1=:bardunk1, bardunk2=:bardunk2, bardunk3=:bardunk3, bardunk4=:bardunk4, bardunk5=:bardunk5, bardname=:bardname, bardunk7=:bardunk7, UNK214=:UNK214, UNK219=:UNK219, UNK220=:UNK220, UNK221=:UNK221, heirloom=:heirloom, UNK223=:UNK223, UNK224=:UNK224, UNK225=:UNK225, UNK226=:UNK226, UNK227=:UNK227, UNK228=:UNK228, UNK229=:UNK229, UNK230=:UNK230, UNK231=:UNK231, UNK232=:UNK232, UNK233=:UNK233, UNK234=:UNK234, placeable=:placeable, UNK236=:UNK236, UNK237=:UNK237, UNK238=:UNK238, UNK239=:UNK239, UNK240=:UNK240, UNK241=:UNK241, epicitem=:epicitem`

func (s *Storage) GetItem(itemId int64) (item *model.Item, err error) {
	item = &model.Item{}
	err = s.db.Get(item, fmt.Sprintf("SELECT id, %s FROM items WHERE id = ?", itemFields), itemId)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateItem(item *model.Item) (err error) {
	if item == nil {
		err = fmt.Errorf("Must provide item")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO items(%s)
		VALUES (%s)`, itemFields, itemBinds), item)
	if err != nil {
		return
	}
	itemId, err := result.LastInsertId()
	if err != nil {
		return
	}
	item.Id = itemId
	return
}

func (s *Storage) ListItem() (items []*model.Item, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT %s FROM items ORDER BY id DESC`, itemFields))
	if err != nil {
		return
	}

	for rows.Next() {
		item := model.Item{}
		if err = rows.StructScan(&item); err != nil {
			return
		}
		items = append(items, &item)
	}
	return
}

func (s *Storage) EditItem(itemId int64, item *model.Item) (err error) {
	item.Id = itemId
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE items SET %s WHERE id = :id`, itemSet), item)
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

func (s *Storage) DeleteItem(itemId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM items WHERE id = ?`, itemId)
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
