package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const itemFields = `items.name, items.aagi, items.ac, items.accuracy, items.acha, items.adex, items.aint, items.artifactflag, items.asta, items.astr, items.attack, items.augrestrict, items.augslot1type, items.augslot1visible, items.augslot2type, items.augslot2visible, items.augslot3type, items.augslot3visible, items.augslot4type, items.augslot4visible, items.augslot5type, items.augslot5visible, items.augslot6type, items.augslot6visible, items.augtype, items.avoidance, items.awis, items.bagsize, items.bagslots, items.bagtype, items.bagwr, items.banedmgamt, items.banedmgraceamt, items.banedmgbody, items.banedmgrace, items.bardtype, items.bardvalue, items.book, items.casttime, items.casttime_, items.charmfile, items.charmfileid, items.classes, items.color, items.combateffects, items.extradmgskill, items.extradmgamt, items.price, items.cr, items.damage, items.damageshield, items.deity, items.delay, items.augdistiller, items.dotshielding, items.dr, items.clicktype, items.clicklevel2, items.elemdmgtype, items.elemdmgamt, items.endur, items.factionamt1, items.factionamt2, items.factionamt3, items.factionamt4, items.factionmod1, items.factionmod2, items.factionmod3, items.factionmod4, items.filename, items.focuseffect, items.fr, items.fvnodrop, items.haste, items.clicklevel, items.hp, items.regen, items.icon, items.idfile, items.itemclass, items.itemtype, items.ldonprice, items.ldontheme, items.ldonsold, items.light, items.lore, items.loregroup, items.magic, items.mana, items.manaregen, items.enduranceregen, items.material, items.herosforgemodel, items.maxcharges, items.mr, items.nodrop, items.norent, items.pendingloreflag, items.pr, items.procrate, items.races, items.` + "`range`" + `, items.reclevel, items.recskill, items.reqlevel, items.sellrate, items.shielding, items.size, items.skillmodtype, items.skillmodvalue, items.slots, items.clickeffect, items.spellshield, items.strikethrough, items.stunresist, items.summonedflag, items.tradeskills, items.favor, items.weight, items.UNK012, items.UNK013, items.benefitflag, items.UNK054, items.UNK059, items.booktype, items.recastdelay, items.recasttype, items.guildfavor, items.UNK123, items.UNK124, items.attuneable, items.nopet, items.updated, items.comment, items.UNK127, items.pointtype, items.potionbelt, items.potionbeltslots, items.stacksize, items.notransfer, items.stackable, items.UNK134, items.UNK137, items.proceffect, items.proctype, items.proclevel2, items.proclevel, items.UNK142, items.worneffect, items.worntype, items.wornlevel2, items.wornlevel, items.UNK147, items.focustype, items.focuslevel2, items.focuslevel, items.UNK152, items.scrolleffect, items.scrolltype, items.scrolllevel2, items.scrolllevel, items.UNK157, items.serialized, items.verified, items.serialization, items.source, items.UNK033, items.lorefile, items.UNK014, items.svcorruption, items.skillmodmax, items.UNK060, items.augslot1unk2, items.augslot2unk2, items.augslot3unk2, items.augslot4unk2, items.augslot5unk2, items.augslot6unk2, items.UNK120, items.UNK121, items.questitemflag, items.UNK132, items.clickunk5, items.clickunk6, items.clickunk7, items.procunk1, items.procunk2, items.procunk3, items.procunk4, items.procunk6, items.procunk7, items.wornunk1, items.wornunk2, items.wornunk3, items.wornunk4, items.wornunk5, items.wornunk6, items.wornunk7, items.focusunk1, items.focusunk2, items.focusunk3, items.focusunk4, items.focusunk5, items.focusunk6, items.focusunk7, items.scrollunk1, items.scrollunk2, items.scrollunk3, items.scrollunk4, items.scrollunk5, items.scrollunk6, items.scrollunk7, items.UNK193, items.purity, items.evoitem, items.evoid, items.evolvinglevel, items.evomax, items.clickname, items.procname, items.wornname, items.focusname, items.scrollname, items.dsmitigation, items.heroic_str, items.heroic_int, items.heroic_wis, items.heroic_agi, items.heroic_dex, items.heroic_sta, items.heroic_cha, items.heroic_pr, items.heroic_dr, items.heroic_fr, items.heroic_cr, items.heroic_mr, items.heroic_svcorrup, items.healamt, items.spelldmg, items.clairvoyance, items.backstabdmg, items.created, items.elitematerial, items.ldonsellbackrate, items.scriptfileid, items.expendablearrow, items.powersourcecapacity, items.bardeffect, items.bardeffecttype, items.bardlevel2, items.bardlevel, items.bardunk1, items.bardunk2, items.bardunk3, items.bardunk4, items.bardunk5, items.bardname, items.bardunk7, items.UNK214, items.UNK219, items.UNK220, items.UNK221, items.heirloom, items.UNK223, items.UNK224, items.UNK225, items.UNK226, items.UNK227, items.UNK228, items.UNK229, items.UNK230, items.UNK231, items.UNK232, items.UNK233, items.UNK234, items.placeable, items.UNK236, items.UNK237, items.UNK238, items.UNK239, items.UNK240, items.UNK241, items.epicitem`
const itemSet = `items.name=:name, items.aagi=:aagi, items.ac=:ac, items.accuracy=:accuracy, items.acha=:acha, items.adex=:adex, items.aint=:aint, items.artifactflag=:artifactflag, items.asta=:asta, items.astr=:astr, items.attack=:attack, items.augrestrict=:augrestrict, items.augslot1type=:augslot1type, items.augslot1visible=:augslot1visible, items.augslot2type=:augslot2type, items.augslot2visible=:augslot2visible, items.augslot3type=:augslot3type, items.augslot3visible=:augslot3visible, items.augslot4type=:augslot4type, items.augslot4visible=:augslot4visible, items.augslot5type=:augslot5type, items.augslot5visible=:augslot5visible, items.augslot6type=:augslot6type, items.augslot6visible=:augslot6visible, items.augtype=:augtype, items.avoidance=:avoidance, items.awis=:awis, items.bagsize=:bagsize, items.bagslots=:bagslots, items.bagtype=:bagtype, items.bagwr=:bagwr, items.banedmgamt=:banedmgamt, items.banedmgraceamt=:banedmgraceamt, items.banedmgbody=:banedmgbody, items.banedmgrace=:banedmgrace, items.bardtype=:bardtype, items.bardvalue=:bardvalue, items.book=:book, items.casttime=:casttime, items.casttime_=:casttime_, items.charmfile=:charmfile, items.charmfileid=:charmfileid, items.classes=:classes, items.color=:color, items.combateffects=:combateffects, items.extradmgskill=:extradmgskill, items.extradmgamt=:extradmgamt, items.price=:price, items.cr=:cr, items.damage=:damage, items.damageshield=:damageshield, items.deity=:deity, items.delay=:delay, items.augdistiller=:augdistiller, items.dotshielding=:dotshielding, items.dr=:dr, items.clicktype=:clicktype, items.clicklevel2=:clicklevel2, items.elemdmgtype=:elemdmgtype, items.elemdmgamt=:elemdmgamt, items.endur=:endur, items.factionamt1=:factionamt1, items.factionamt2=:factionamt2, items.factionamt3=:factionamt3, items.factionamt4=:factionamt4, items.factionmod1=:factionmod1, items.factionmod2=:factionmod2, items.factionmod3=:factionmod3, items.factionmod4=:factionmod4, items.filename=:filename, items.focuseffect=:focuseffect, items.fr=:fr, items.fvnodrop=:fvnodrop, items.haste=:haste, items.clicklevel=:clicklevel, items.hp=:hp, items.regen=:regen, items.icon=:icon, items.idfile=:idfile, items.itemclass=:itemclass, items.itemtype=:itemtype, items.ldonprice=:ldonprice, items.ldontheme=:ldontheme, items.ldonsold=:ldonsold, items.light=:light, items.lore=:lore, items.loregroup=:loregroup, items.magic=:magic, items.mana=:mana, items.manaregen=:manaregen, items.enduranceregen=:enduranceregen, items.material=:material, items.herosforgemodel=:herosforgemodel, items.maxcharges=:maxcharges, items.mr=:mr, items.nodrop=:nodrop, items.norent=:norent, items.pendingloreflag=:pendingloreflag, items.pr=:pr, items.procrate=:procrate, items.races=:races, items.` + "`range`" + `=:range, items.reclevel=:reclevel, items.recskill=:recskill, items.reqlevel=:reqlevel, items.sellrate=:sellrate, items.shielding=:shielding, items.size=:size, items.skillmodtype=:skillmodtype, items.skillmodvalue=:skillmodvalue, items.slots=:slots, items.clickeffect=:clickeffect, items.spellshield=:spellshield, items.strikethrough=:strikethrough, items.stunresist=:stunresist, items.summonedflag=:summonedflag, items.tradeskills=:tradeskills, items.favor=:favor, items.weight=:weight, items.UNK012=:UNK012, items.UNK013=:UNK013, items.benefitflag=:benefitflag, items.UNK054=:UNK054, items.UNK059=:UNK059, items.booktype=:booktype, items.recastdelay=:recastdelay, items.recasttype=:recasttype, items.guildfavor=:guildfavor, items.UNK123=:UNK123, items.UNK124=:UNK124, items.attuneable=:attuneable, items.nopet=:nopet, items.updated=:updated, items.comment=:comment, items.UNK127=:UNK127, items.pointtype=:pointtype, items.potionbelt=:potionbelt, items.potionbeltslots=:potionbeltslots, items.stacksize=:stacksize, items.notransfer=:notransfer, items.stackable=:stackable, items.UNK134=:UNK134, items.UNK137=:UNK137, items.proceffect=:proceffect, items.proctype=:proctype, items.proclevel2=:proclevel2, items.proclevel=:proclevel, items.UNK142=:UNK142, items.worneffect=:worneffect, items.worntype=:worntype, items.wornlevel2=:wornlevel2, items.wornlevel=:wornlevel, items.UNK147=:UNK147, items.focustype=:focustype, items.focuslevel2=:focuslevel2, items.focuslevel=:focuslevel, items.UNK152=:UNK152, items.scrolleffect=:scrolleffect, items.scrolltype=:scrolltype, items.scrolllevel2=:scrolllevel2, items.scrolllevel=:scrolllevel, items.UNK157=:UNK157, items.serialized=:serialized, items.verified=:verified, items.serialization=:serialization, items.source=:source, items.UNK033=:UNK033, items.lorefile=:lorefile, items.UNK014=:UNK014, items.svcorruption=:svcorruption, items.skillmodmax=:skillmodmax, items.UNK060=:UNK060, items.augslot1unk2=:augslot1unk2, items.augslot2unk2=:augslot2unk2, items.augslot3unk2=:augslot3unk2, items.augslot4unk2=:augslot4unk2, items.augslot5unk2=:augslot5unk2, items.augslot6unk2=:augslot6unk2, items.UNK120=:UNK120, items.UNK121=:UNK121, items.questitemflag=:questitemflag, items.UNK132=:UNK132, items.clickunk5=:clickunk5, items.clickunk6=:clickunk6, items.clickunk7=:clickunk7, items.procunk1=:procunk1, items.procunk2=:procunk2, items.procunk3=:procunk3, items.procunk4=:procunk4, items.procunk6=:procunk6, items.procunk7=:procunk7, items.wornunk1=:wornunk1, items.wornunk2=:wornunk2, items.wornunk3=:wornunk3, items.wornunk4=:wornunk4, items.wornunk5=:wornunk5, items.wornunk6=:wornunk6, items.wornunk7=:wornunk7, items.focusunk1=:focusunk1, items.focusunk2=:focusunk2, items.focusunk3=:focusunk3, items.focusunk4=:focusunk4, items.focusunk5=:focusunk5, items.focusunk6=:focusunk6, items.focusunk7=:focusunk7, items.scrollunk1=:scrollunk1, items.scrollunk2=:scrollunk2, items.scrollunk3=:scrollunk3, items.scrollunk4=:scrollunk4, items.scrollunk5=:scrollunk5, items.scrollunk6=:scrollunk6, items.scrollunk7=:scrollunk7, items.UNK193=:UNK193, items.purity=:purity, items.evoitem=:evoitem, items.evoid=:evoid, items.evolvinglevel=:evolvinglevel, items.evomax=:evomax, items.clickname=:clickname, items.procname=:procname, items.wornname=:wornname, items.focusname=:focusname, items.scrollname=:scrollname, items.dsmitigation=:dsmitigation, items.heroic_str=:heroic_str, items.heroic_int=:heroic_int, items.heroic_wis=:heroic_wis, items.heroic_agi=:heroic_agi, items.heroic_dex=:heroic_dex, items.heroic_sta=:heroic_sta, items.heroic_cha=:heroic_cha, items.heroic_pr=:heroic_pr, items.heroic_dr=:heroic_dr, items.heroic_fr=:heroic_fr, items.heroic_cr=:heroic_cr, items.heroic_mr=:heroic_mr, items.heroic_svcorrup=:heroic_svcorrup, items.healamt=:healamt, items.spelldmg=:spelldmg, items.clairvoyance=:clairvoyance, items.backstabdmg=:backstabdmg, items.created=:created, items.elitematerial=:elitematerial, items.ldonsellbackrate=:ldonsellbackrate, items.scriptfileid=:scriptfileid, items.expendablearrow=:expendablearrow, items.powersourcecapacity=:powersourcecapacity, items.bardeffect=:bardeffect, items.bardeffecttype=:bardeffecttype, items.bardlevel2=:bardlevel2, items.bardlevel=:bardlevel, items.bardunk1=:bardunk1, items.bardunk2=:bardunk2, items.bardunk3=:bardunk3, items.bardunk4=:bardunk4, items.bardunk5=:bardunk5, items.bardname=:bardname, items.bardunk7=:bardunk7, items.UNK214=:UNK214, items.UNK219=:UNK219, items.UNK220=:UNK220, items.UNK221=:UNK221, items.heirloom=:heirloom, items.UNK223=:UNK223, items.UNK224=:UNK224, items.UNK225=:UNK225, items.UNK226=:UNK226, items.UNK227=:UNK227, items.UNK228=:UNK228, items.UNK229=:UNK229, items.UNK230=:UNK230, items.UNK231=:UNK231, items.UNK232=:UNK232, items.UNK233=:UNK233, items.UNK234=:UNK234, items.placeable=:placeable, items.UNK236=:UNK236, items.UNK237=:UNK237, items.UNK238=:UNK238, items.UNK239=:UNK239, items.UNK240=:UNK240, items.UNK241=:UNK241, items.epicitem=:epicitem`
const itemBinds = `:name, :aagi, :ac, :accuracy, :acha, :adex, :aint, :artifactflag, :asta, :astr, :attack, :augrestrict, :augslot1type, :augslot1visible, :augslot2type, :augslot2visible, :augslot3type, :augslot3visible, :augslot4type, :augslot4visible, :augslot5type, :augslot5visible, :augslot6type, :augslot6visible, :augtype, :avoidance, :awis, :bagsize, :bagslots, :bagtype, :bagwr, :banedmgamt, :banedmgraceamt, :banedmgbody, :banedmgrace, :bardtype, :bardvalue, :book, :casttime, :casttime_, :charmfile, :charmfileid, :classes, :color, :combateffects, :extradmgskill, :extradmgamt, :price, :cr, :damage, :damageshield, :deity, :delay, :augdistiller, :dotshielding, :dr, :clicktype, :clicklevel2, :elemdmgtype, :elemdmgamt, :endur, :factionamt1, :factionamt2, :factionamt3, :factionamt4, :factionmod1, :factionmod2, :factionmod3, :factionmod4, :filename, :focuseffect, :fr, :fvnodrop, :haste, :clicklevel, :hp, :regen, :icon, :idfile, :itemclass, :itemtype, :ldonprice, :ldontheme, :ldonsold, :light, :lore, :loregroup, :magic, :mana, :manaregen, :enduranceregen, :material, :herosforgemodel, :maxcharges, :mr, :nodrop, :norent, :pendingloreflag, :pr, :procrate, :races, :range, :reclevel, :recskill, :reqlevel, :sellrate, :shielding, :size, :skillmodtype, :skillmodvalue, :slots, :clickeffect, :spellshield, :strikethrough, :stunresist, :summonedflag, :tradeskills, :favor, :weight, :UNK012, :UNK013, :benefitflag, :UNK054, :UNK059, :booktype, :recastdelay, :recasttype, :guildfavor, :UNK123, :UNK124, :attuneable, :nopet, :updated, :comment, :UNK127, :pointtype, :potionbelt, :potionbeltslots, :stacksize, :notransfer, :stackable, :UNK134, :UNK137, :proceffect, :proctype, :proclevel2, :proclevel, :UNK142, :worneffect, :worntype, :wornlevel2, :wornlevel, :UNK147, :focustype, :focuslevel2, :focuslevel, :UNK152, :scrolleffect, :scrolltype, :scrolllevel2, :scrolllevel, :UNK157, :serialized, :verified, :serialization, :source, :UNK033, :lorefile, :UNK014, :svcorruption, :skillmodmax, :UNK060, :augslot1unk2, :augslot2unk2, :augslot3unk2, :augslot4unk2, :augslot5unk2, :augslot6unk2, :UNK120, :UNK121, :questitemflag, :UNK132, :clickunk5, :clickunk6, :clickunk7, :procunk1, :procunk2, :procunk3, :procunk4, :procunk6, :procunk7, :wornunk1, :wornunk2, :wornunk3, :wornunk4, :wornunk5, :wornunk6, :wornunk7, :focusunk1, :focusunk2, :focusunk3, :focusunk4, :focusunk5, :focusunk6, :focusunk7, :scrollunk1, :scrollunk2, :scrollunk3, :scrollunk4, :scrollunk5, :scrollunk6, :scrollunk7, :UNK193, :purity, :evoitem, :evoid, :evolvinglevel, :evomax, :clickname, :procname, :wornname, :focusname, :scrollname, :dsmitigation, :heroic_str, :heroic_int, :heroic_wis, :heroic_agi, :heroic_dex, :heroic_sta, :heroic_cha, :heroic_pr, :heroic_dr, :heroic_fr, :heroic_cr, :heroic_mr, :heroic_svcorrup, :healamt, :spelldmg, :clairvoyance, :backstabdmg, :created, :elitematerial, :ldonsellbackrate, :scriptfileid, :expendablearrow, :powersourcecapacity, :bardeffect, :bardeffecttype, :bardlevel2, :bardlevel, :bardunk1, :bardunk2, :bardunk3, :bardunk4, :bardunk5, :bardname, :bardunk7, :UNK214, :UNK219, :UNK220, :UNK221, :heirloom, :UNK223, :UNK224, :UNK225, :UNK226, :UNK227, :UNK228, :UNK229, :UNK230, :UNK231, :UNK232, :UNK233, :UNK234, :placeable, :UNK236, :UNK237, :UNK238, :UNK239, :UNK240, :UNK241, :epicitem`

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
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, %s FROM items 
		ORDER BY id DESC`, itemFields))
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

func (s *Storage) ListItemByCharacter(characterId int64) (items []*model.Item, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT items.id, slotid,charges,inventory.color invcolor,augslot1,augslot2,augslot3,augslot4,augslot5,augslot6,instnodrop,custom_data,ornamenticon,ornamentidfile,ornament_hero_model, %s FROM items 
		INNER JOIN inventory ON inventory.itemid = items.id
		WHERE inventory.charid = ? ORDER BY slotid ASC`, itemFields), characterId)
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
