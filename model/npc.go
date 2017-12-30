package model

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

type Npc struct {
	Id                  int64          `json:"id" db:"id"`                                     //`id` int(11) NOT NULL AUTO_INCREMENT,
	Name                string         `json:"name" db:"name"`                                 //`name` text NOT NULL,
	Lastname            sql.NullString `json:"lastname" db:"lastname"`                         //`lastname` varchar(32) DEFAULT NULL,
	Level               int64          `json:"level" db:"level"`                               //`level` tinyint(2) unsigned NOT NULL DEFAULT '0',
	Race                int64          `json:"race" db:"race"`                                 //`race` smallint(5) unsigned NOT NULL DEFAULT '0',
	Class               int64          `json:"class" db:"class"`                               //`class` tinyint(2) unsigned NOT NULL DEFAULT '0',
	Bodytype            int64          `json:"bodytype" db:"bodytype"`                         //`bodytype` int(11) NOT NULL DEFAULT '1',
	Hp                  int64          `json:"hp" db:"hp"`                                     //`hp` int(11) NOT NULL DEFAULT '0',
	Mana                int64          `json:"mana" db:"mana"`                                 //`mana` int(11) NOT NULL DEFAULT '0',
	Gender              int64          `json:"gender" db:"gender"`                             //`gender` tinyint(2) unsigned NOT NULL DEFAULT '0',
	Texture             int64          `json:"texture" db:"texture"`                           //`texture` tinyint(2) unsigned NOT NULL DEFAULT '0',
	Helmtexture         int64          `json:"helmtexture" db:"helmtexture"`                   //`helmtexture` tinyint(2) unsigned NOT NULL DEFAULT '0',
	Herosforgemodel     int64          `json:"herosforgemodel" db:"herosforgemodel"`           //`herosforgemodel` int(11) NOT NULL DEFAULT '0',
	Size                float64        `json:"size" db:"size"`                                 //`size` float NOT NULL DEFAULT '0',
	HpRegenRate         int64          `json:"hpRegenRate" db:"hp_regen_rate"`                 //`hp_regen_rate` int(11) unsigned NOT NULL DEFAULT '0',
	ManaRegenRate       int64          `json:"manaRegenRate" db:"mana_regen_rate"`             //`mana_regen_rate` int(11) unsigned NOT NULL DEFAULT '0',
	LoottableId         int64          `json:"loottableId" db:"loottable_id"`                  //`loottable_id` int(11) unsigned NOT NULL DEFAULT '0',
	MerchantId          int64          `json:"merchantId" db:"merchant_id"`                    //`merchant_id` int(11) unsigned NOT NULL DEFAULT '0',
	AltCurrencyId       int64          `json:"altCurrencyId" db:"alt_currency_id"`             //`alt_currency_id` int(11) unsigned NOT NULL DEFAULT '0',
	NpcSpellsId         int64          `json:"npcSpellsId" db:"npc_spells_id"`                 //`npc_spells_id` int(11) unsigned NOT NULL DEFAULT '0',
	NpcSpellsEffectsId  int64          `json:"npcSpellsEffectsId" db:"npc_spells_effects_id"`  //`npc_spells_effects_id` int(11) unsigned NOT NULL DEFAULT '0',
	NpcFactionId        int64          `json:"npcFactionId" db:"npc_faction_id"`               //`npc_faction_id` int(11) NOT NULL DEFAULT '0',
	AdventureTemplateId int64          `json:"adventureTemplateId" db:"adventure_template_id"` //`adventure_template_id` int(10) unsigned NOT NULL DEFAULT '0',
	TrapTemplate        sql.NullInt64  `json:"trapTemplate" db:"trap_template"`                //`trap_template` int(10) unsigned DEFAULT '0',
	Mindmg              int64          `json:"mindmg" db:"mindmg"`                             //`mindmg` int(10) unsigned NOT NULL DEFAULT '0',
	Maxdmg              int64          `json:"maxdmg" db:"maxdmg"`                             //`maxdmg` int(10) unsigned NOT NULL DEFAULT '0',
	AttackCount         int64          `json:"attackCount" db:"attack_count"`                  //`attack_count` smallint(6) NOT NULL DEFAULT '-1',
	Npcspecialattks     string         `json:"npcspecialattks" db:"npcspecialattks"`           //`npcspecialattks` varchar(36) NOT NULL DEFAULT '',
	SpecialAbilities    sql.NullString `json:"specialAbilities" db:"special_abilities"`        //`special_abilities` text,
	Aggroradius         int64          `json:"aggroradius" db:"aggroradius"`                   //`aggroradius` int(10) unsigned NOT NULL DEFAULT '0',
	Assistradius        int64          `json:"assistradius" db:"assistradius"`                 //`assistradius` int(10) unsigned NOT NULL DEFAULT '0',
	Face                int64          `json:"face" db:"face"`                                 //`face` int(10) unsigned NOT NULL DEFAULT '1',
	LuclinHairstyle     int64          `json:"luclinHairstyle" db:"luclin_hairstyle"`          //`luclin_hairstyle` int(10) unsigned NOT NULL DEFAULT '1',
	LuclinHaircolor     int64          `json:"luclinHaircolor" db:"luclin_haircolor"`          //`luclin_haircolor` int(10) unsigned NOT NULL DEFAULT '1',
	LuclinEyecolor      int64          `json:"luclinEyecolor" db:"luclin_eyecolor"`            //`luclin_eyecolor` int(10) unsigned NOT NULL DEFAULT '1',
	LuclinEyecolor2     int64          `json:"luclinEyecolor2" db:"luclin_eyecolor2"`          //`luclin_eyecolor2` int(10) unsigned NOT NULL DEFAULT '1',
	LuclinBeardcolor    int64          `json:"luclinBeardcolor" db:"luclin_beardcolor"`        //`luclin_beardcolor` int(10) unsigned NOT NULL DEFAULT '1',
	LuclinBeard         int64          `json:"luclinBeard" db:"luclin_beard"`                  //`luclin_beard` int(10) unsigned NOT NULL DEFAULT '0',
	DrakkinHeritage     int64          `json:"drakkinHeritage" db:"drakkin_heritage"`          //`drakkin_heritage` int(10) NOT NULL DEFAULT '0',
	DrakkinTattoo       int64          `json:"drakkinTattoo" db:"drakkin_tattoo"`              //`drakkin_tattoo` int(10) NOT NULL DEFAULT '0',
	DrakkinDetails      int64          `json:"drakkinDetails" db:"drakkin_details"`            //`drakkin_details` int(10) NOT NULL DEFAULT '0',
	ArmortintId         int64          `json:"armortintId" db:"armortint_id"`                  //`armortint_id` int(10) unsigned NOT NULL DEFAULT '0',
	ArmortintRed        int64          `json:"armortintRed" db:"armortint_red"`                //`armortint_red` tinyint(3) unsigned NOT NULL DEFAULT '0',
	ArmortintGreen      int64          `json:"armortintGreen" db:"armortint_green"`            //`armortint_green` tinyint(3) unsigned NOT NULL DEFAULT '0',
	ArmortintBlue       int64          `json:"armortintBlue" db:"armortint_blue"`              //`armortint_blue` tinyint(3) unsigned NOT NULL DEFAULT '0',
	DMeleeTexture1      int64          `json:"dMeleeTexture1" db:"d_melee_texture1"`           //`d_melee_texture1` int(11) NOT NULL DEFAULT '0',
	DMeleeTexture2      int64          `json:"dMeleeTexture2" db:"d_melee_texture2"`           //`d_melee_texture2` int(11) NOT NULL DEFAULT '0',
	AmmoIdfile          string         `json:"ammoIdfile" db:"ammo_idfile"`                    //`ammo_idfile` varchar(30) NOT NULL DEFAULT 'IT10',
	PrimMeleeType       int64          `json:"primMeleeType" db:"prim_melee_type"`             //`prim_melee_type` tinyint(4) unsigned NOT NULL DEFAULT '28',
	SecMeleeType        int64          `json:"secMeleeType" db:"sec_melee_type"`               //`sec_melee_type` tinyint(4) unsigned NOT NULL DEFAULT '28',
	RangedType          int64          `json:"rangedType" db:"ranged_type"`                    //`ranged_type` tinyint(4) unsigned NOT NULL DEFAULT '7',
	Runspeed            float64        `json:"runspeed" db:"runspeed"`                         //`runspeed` float NOT NULL DEFAULT '0',
	Mr                  int64          `json:"MR" db:"MR"`                                     //`MR` smallint(5) NOT NULL DEFAULT '0',
	Cr                  int64          `json:"CR" db:"CR"`                                     //`CR` smallint(5) NOT NULL DEFAULT '0',
	Dr                  int64          `json:"DR" db:"DR"`                                     //`DR` smallint(5) NOT NULL DEFAULT '0',
	Fr                  int64          `json:"FR" db:"FR"`                                     //`FR` smallint(5) NOT NULL DEFAULT '0',
	Pr                  int64          `json:"PR" db:"PR"`                                     //`PR` smallint(5) NOT NULL DEFAULT '0',
	Corrup              int64          `json:"Corrup" db:"Corrup"`                             //`Corrup` smallint(5) NOT NULL DEFAULT '0',
	Phr                 int64          `json:"PhR" db:"PhR"`                                   //`PhR` smallint(5) unsigned NOT NULL DEFAULT '0',
	SeeInvis            int64          `json:"seeInvis" db:"see_invis"`                        //`see_invis` smallint(4) NOT NULL DEFAULT '0',
	SeeInvisUndead      int64          `json:"seeInvisUndead" db:"see_invis_undead"`           //`see_invis_undead` smallint(4) NOT NULL DEFAULT '0',
	Qglobal             int64          `json:"qglobal" db:"qglobal"`                           //`qglobal` int(2) unsigned NOT NULL DEFAULT '0',
	Ac                  int64          `json:"AC" db:"AC"`                                     //`AC` smallint(5) NOT NULL DEFAULT '0',
	NpcAggro            int64          `json:"npcAggro" db:"npc_aggro"`                        //`npc_aggro` tinyint(4) NOT NULL DEFAULT '0',
	SpawnLimit          int64          `json:"spawnLimit" db:"spawn_limit"`                    //`spawn_limit` tinyint(4) NOT NULL DEFAULT '0',
	AttackSpeed         float64        `json:"attackSpeed" db:"attack_speed"`                  //`attack_speed` float NOT NULL DEFAULT '0',
	AttackDelay         int64          `json:"attackDelay" db:"attack_delay"`                  //`attack_delay` tinyint(3) unsigned NOT NULL DEFAULT '30',
	Findable            int64          `json:"findable" db:"findable"`                         //`findable` tinyint(4) NOT NULL DEFAULT '0',
	Str                 int64          `json:"STR" db:"STR"`                                   //`STR` mediumint(8) unsigned NOT NULL DEFAULT '75',
	Sta                 int64          `json:"STA" db:"STA"`                                   //`STA` mediumint(8) unsigned NOT NULL DEFAULT '75',
	Dex                 int64          `json:"DEX" db:"DEX"`                                   //`DEX` mediumint(8) unsigned NOT NULL DEFAULT '75',
	Agi                 int64          `json:"AGI" db:"AGI"`                                   //`AGI` mediumint(8) unsigned NOT NULL DEFAULT '75',
	Int                 int64          `json:"INT" db:"_INT"`                                  //`_INT` mediumint(8) unsigned NOT NULL DEFAULT '80',
	Wis                 int64          `json:"WIS" db:"WIS"`                                   //`WIS` mediumint(8) unsigned NOT NULL DEFAULT '75',
	Cha                 int64          `json:"CHA" db:"CHA"`                                   //`CHA` mediumint(8) unsigned NOT NULL DEFAULT '75',
	SeeHide             int64          `json:"seeHide" db:"see_hide"`                          //`see_hide` tinyint(4) NOT NULL DEFAULT '0',
	SeeImprovedHide     int64          `json:"seeImprovedHide" db:"see_improved_hide"`         //`see_improved_hide` tinyint(4) NOT NULL DEFAULT '0',
	Trackable           int64          `json:"trackable" db:"trackable"`                       //`trackable` tinyint(4) NOT NULL DEFAULT '1',
	Isbot               int64          `json:"isbot" db:"isbot"`                               //`isbot` tinyint(4) NOT NULL DEFAULT '0',
	Exclude             int64          `json:"exclude" db:"exclude"`                           //`exclude` tinyint(4) NOT NULL DEFAULT '1',
	Atk                 int64          `json:"ATK" db:"ATK"`                                   //`ATK` mediumint(9) NOT NULL DEFAULT '0',
	Accuracy            int64          `json:"Accuracy" db:"Accuracy"`                         //`Accuracy` mediumint(9) NOT NULL DEFAULT '0',
	Avoidance           int64          `json:"Avoidance" db:"Avoidance"`                       //`Avoidance` mediumint(9) unsigned NOT NULL DEFAULT '0',
	SlowMitigation      int64          `json:"slowMitigation" db:"slow_mitigation"`            //`slow_mitigation` smallint(4) NOT NULL DEFAULT '0',
	Version             int64          `json:"version" db:"version"`                           //`version` smallint(5) unsigned NOT NULL DEFAULT '0',
	Maxlevel            int64          `json:"maxlevel" db:"maxlevel"`                         //`maxlevel` tinyint(3) NOT NULL DEFAULT '0',
	Scalerate           int64          `json:"scalerate" db:"scalerate"`                       //`scalerate` int(11) NOT NULL DEFAULT '100',
	PrivateCorpse       int64          `json:"privateCorpse" db:"private_corpse"`              //`private_corpse` tinyint(3) unsigned NOT NULL DEFAULT '0',
	UniqueSpawnByName   int64          `json:"uniqueSpawnByName" db:"unique_spawn_by_name"`    //`unique_spawn_by_name` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Underwater          int64          `json:"underwater" db:"underwater"`                     //`underwater` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Isquest             int64          `json:"isquest" db:"isquest"`                           //`isquest` tinyint(3) NOT NULL DEFAULT '0',
	Emoteid             int64          `json:"emoteid" db:"emoteid"`                           //`emoteid` int(10) unsigned NOT NULL DEFAULT '0',
	Spellscale          float64        `json:"spellscale" db:"spellscale"`                     //`spellscale` float NOT NULL DEFAULT '100',
	Healscale           float64        `json:"healscale" db:"healscale"`                       //`healscale` float NOT NULL DEFAULT '100',
	NoTargetHotkey      int64          `json:"noTargetHotkey" db:"no_target_hotkey"`           //`no_target_hotkey` tinyint(1) unsigned NOT NULL DEFAULT '0',
	RaidTarget          int64          `json:"raidTarget" db:"raid_target"`                    //`raid_target` tinyint(1) unsigned NOT NULL DEFAULT '0',
	Armtexture          int64          `json:"armtexture" db:"armtexture"`                     //`armtexture` tinyint(2) NOT NULL DEFAULT '0',
	Bracertexture       int64          `json:"bracertexture" db:"bracertexture"`               //`bracertexture` tinyint(2) NOT NULL DEFAULT '0',
	Handtexture         int64          `json:"handtexture" db:"handtexture"`                   //`handtexture` tinyint(2) NOT NULL DEFAULT '0',
	Legtexture          int64          `json:"legtexture" db:"legtexture"`                     //`legtexture` tinyint(2) NOT NULL DEFAULT '0',
	Feettexture         int64          `json:"feettexture" db:"feettexture"`                   //`feettexture` tinyint(2) NOT NULL DEFAULT '0',
	Light               int64          `json:"light" db:"light"`                               //`light` tinyint(2) NOT NULL DEFAULT '0',
	Walkspeed           int64          `json:"walkspeed" db:"walkspeed"`                       //`walkspeed` tinyint(2) NOT NULL DEFAULT '0',
	Peqid               int64          `json:"peqid" db:"peqid"`                               //`peqid` int(11) NOT NULL DEFAULT '0',
	Unique              int64          `json:"unique" db:"unique_"`                            //`unique_` tinyint(2) NOT NULL DEFAULT '0',
	Fixed               int64          `json:"fixed" db:"fixed"`                               //`fixed` tinyint(2) NOT NULL DEFAULT '0',
	IgnoreDespawn       int64          `json:"ignoreDespawn" db:"ignore_despawn"`              //`ignore_despawn` tinyint(2) NOT NULL DEFAULT '0',
	ShowName            int64          `json:"showName" db:"show_name"`                        //`show_name` tinyint(2) NOT NULL DEFAULT '1',
	Untargetable        int64          `json:"untargetable" db:"untargetable"`                 //`untargetable` tinyint(2) NOT NULL DEFAULT '0',
}

func (c *Npc) ZoneId() int64 {
	if c.Id > 1000 {
		return c.Id / 1000
	}
	return 0
}

func (c *Npc) ZoneName() string {
	return ZoneName(c.ZoneId())
}

func (c *Npc) ClassIcon() string {
	return ClassIcon(c.Class)
}
func (c *Npc) RaceIcon() string {
	return RaceIcon(c.Race)
}

func (c *Npc) SpecialAbilitiesList() map[string]string {
	abilities := make(map[string]string)
	rawAbils := strings.Split(c.SpecialAbilities.String, ",")
	for _, abil := range rawAbils {
		breakdown := strings.Split(abil, ",")
		key := ""
		description := ""
		var val int64
		if len(breakdown) < 1 {
			continue
		}
		fmt.Println(breakdown[0], "foo")
		switch breakdown[0] { //based on http://wiki.eqemulator.org/p?NPC_Special_Attacks
		case "1": //Summon
			key = "Summons"
			if len(breakdown) == 1 {
				description += fmt.Sprintf("you to them at lvl 41, 90%% or less HP, 6s cooldown, ")
			}
			if len(breakdown) > 1 {
				//level enabled at
				description += fmt.Sprintf("you to them at lvl %s, ", breakdown[1])
			}
			if len(breakdown) > 2 {
				description += fmt.Sprintf("to you at lvl %s, ", breakdown[2])
			}
			if len(breakdown) > 3 {
				val, _ = strconv.ParseInt(breakdown[3], 10, 64)
				if val != 0 {
					description += fmt.Sprintf("%ds cooldown, ", val/1000)
				} else {
					description += fmt.Sprintf("360s cooldown, ")
				}
			}
			if len(breakdown) > 4 {
				description += fmt.Sprintf("at %s%% or less HP, ", breakdown[4])
			}
		case "2": //Enrage
			key = "Enrages"
			if len(breakdown) > 1 {
				val, _ = strconv.ParseInt(breakdown[1], 10, 64)
				if val == 0 {
					//todo: get default Rule NPC:StartEnrageValue
				} else {
					description += fmt.Sprintf("when hp is less than %d%%, ", val)
				}
			} else {
				description += fmt.Sprintf("when hp is less than [defaultenrage]%%, ")
			}

			if len(breakdown) > 2 {
				val, _ = strconv.ParseInt(breakdown[2], 10, 64)
				if val != 0 {
					description += fmt.Sprintf("enrages for %ds, ", val/1000)
				}
			} else {
				description += fmt.Sprintf("enrages for 10s, ")
			}

			if len(breakdown) > 3 {
				val, _ = strconv.ParseInt(breakdown[2], 10, 64)
				if val != 0 {
					description += fmt.Sprintf("%ds cooldown, ", val/1000)
				} else {
					description += fmt.Sprintf("360s cooldown, ")
				}
			} else {
				description += fmt.Sprintf("360s cooldown, ")
			}
		case "3": //Rampage
			key = "Rampages"
			if len(breakdown) == 1 {
				description += "20%% chance, [[maxramptargets]], 100%% normal damage, "
			}
			if len(breakdown) > 1 { //% chance
				description += fmt.Sprintf("%s%% chance, ", breakdown[1])
			}
			if len(breakdown) > 2 { //target count
				description += fmt.Sprintf("%s ramp targets, ", breakdown[2])
			}
			if len(breakdown) > 3 { //noraml damage
				description += fmt.Sprintf("%s%% normal damage, ", breakdown[3])
			}
			if len(breakdown) > 4 { //flat dmg bonus
				description += fmt.Sprintf("%s bonus damage, ", breakdown[4])
			}
			if len(breakdown) > 5 { //ignore % armor
				description += fmt.Sprintf("%s%% ignored armor, ", breakdown[5])
			}
			if len(breakdown) > 6 { //ignore armor
				description += fmt.Sprintf("%s ignored armor, ", breakdown[6])
			}
			if len(breakdown) > 7 { //crit
				description += fmt.Sprintf("%s%% natural crit, ", breakdown[7])
			}
			if len(breakdown) > 8 { //crit bonus
				description += fmt.Sprintf("%s%% crit bonus, ", breakdown[8])
			}
		case "4": //Rampage
			key = "AE Rampages"
			if len(breakdown) == 1 {
				description += "20%% chance, all within range, 100%% normal damage, "
			}
			if len(breakdown) > 1 { //% chance
				description += fmt.Sprintf("%s%% chance, ", breakdown[1])
			}
			if len(breakdown) > 2 { //target count
				description += fmt.Sprintf("%s ramp targets, ", breakdown[2])
			}
			if len(breakdown) > 3 { //noraml damage
				description += fmt.Sprintf("%s%% normal damage, ", breakdown[3])
			}
			if len(breakdown) > 4 { //flat dmg bonus
				description += fmt.Sprintf("%s bonus damage, ", breakdown[4])
			}
			if len(breakdown) > 5 { //ignore % armor
				description += fmt.Sprintf("%s%% ignored armor, ", breakdown[5])
			}
			if len(breakdown) > 6 { //ignore armor
				description += fmt.Sprintf("%s ignored armor, ", breakdown[6])
			}
			if len(breakdown) > 7 { //crit
				description += fmt.Sprintf("%s%% natural crit, ", breakdown[7])
			}
			if len(breakdown) > 8 { //crit bonus
				description += fmt.Sprintf("%s%% crit bonus, ", breakdown[8])
			}
		case "5": //Flurry
			key = "Flurries"
			description = "Attacks you multiple times, "
		//todo
		case "6": //Triple Attack
			key = "Triple Attacks"
			description = "Attacks 3 times, "
		case "7": //Quad Attack
			key = "Quad Attacks"
			description = "Attack 4 times, "
		case "8": //Dual Wield
			key = "Dual Wields"
			description = "Attacks with both hands, "
		case "9": //Bane Attack
			key = "Bane Attacks"
			description = "Uses bane attacks, "
		case "10": //Magical Attack
			key = "Magical Attacks"
			description = "Uses magical attacks, "
		case "11": //Ranged Attack
			key = "Ranged Attacks"
			description = "Uses ranged attacks, "
		//todo: add ranged details
		case "12": //Unslwoable
			key = "Unslowable"
			description = "Immune to slow, "
		case "13": //Unmezzable
			key = "Unmezzable"
			description = "Immune to mez, "
		case "14": //Uncharmable
			key = "Uncharmable"
			description = "Immune to Charm, "
		case "15": //unstunable
			key = "Unstunable"
			description = "Immune to Stun, "
		case "16": //Unsnarable
			key = "Unsnarable"
			description = "Immune to Snare, "
		case "17": //Unfearable
			key = "Unfearable"
			description = "Immune to Fear, "
		case "18": //Immune to Dispell
			key = "Undispellable"
			description = "Immune to Dispell, "
		case "19": //Immune to Melee
			key = "Melee Invulnerable"
			description = "Immune to all damage, "
		case "20": //Immune to Magic
			key = "Magic Invulnerable"
			description = "Immune to magic, "
		case "21": //Unfleeable
			key = "Does not flee"
			description = "Immune to fleeing, "
		case "22": //immune to bane
			key = "Non-bane immune"
			description = "Immune to non-bane attacks, "
		case "23": //immune to non-magical
			key = "Non-magic immune"
			description = "Immune to non-magic attacks, "
		case "24": //never aggro
			key = "Non-KOS"
			description = "Never aggros, "
		case "25": //immune target
			key = "Immune Target"
			description = "Immune to Targeting, "
		case "26": //immune from casting fro range
			key = "Immune Ranged"
			description = "Immune from casting from range, "
		case "27": //immune fd
			key = "Immune Feign Death"
			description = "Ignores Feign Death, "
		case "28": //immune taunt
			key = "Immune Taunt"
			description = "Ignores Taunt, "
		case "29": //tunnel vision
			key = "Tunnel Vision"
			description = "Vision is focused, "
		//aggromod tunnelvision todo
		case "30": //non-assist
			key = "Non-assist"
			description = "Does not assist allies, "
		case "31": //pacify-immune
			key = "Pacify-immune"
			description = "Ignores pacify, "
		case "32": //leash
			key = "Leashed"
			description = "Leash to aggro range, full heal, memwipe, "
		//todo add param
		case "33": //leash to aggro range
			key = "Leashed To Aggro"
			description = "Leashed to aggro range, "
		case "34":
			key = "Destructable"
			description = "Destructable Object"
		case "35":
			key = "Immune Client"
			description = "Immune to Client Attacks"
		case "36":
			key = "Flees"
			description = "Flees"
		case "37":
			key = "Flee Percent"
			description = "Flees at percent"
			//todo percent to flee
		case "38":
			key = "Allow Beneficial"
			description = "Allow Beneficial"
		case "39":
			key = "Disable melee"
			description = "Does not melee"
		case "40":
			key = "Chance Distance"
			description = "Chances to a distance"
			//todo: chance distance
		case "41":
			key = "Allow to Tank"
			description = "Allow NPC to tank instead of client"
		case "42":
			key = "Ignore Root"
			description = "Ignores root rules"
		case "43":
			key = "Innate Resist Diff"
			description = "Gives innate resist"
		//todo: add more
		case "44":
			key = "Counter Avoid Damage"
			description = "Avoid Damage"
			//todo: add more
		}

		if len(key) < 1 {
			continue
		}
		if len(description) > 2 {
			description = description[0 : len(description)-2]
		}
		abilities[key] = description
	}
	fmt.Println(abilities)
	return abilities
}

func (c *Npc) Experience() int64 {
	xp := c.Level * c.Level * 75 * 35 / 10 //EXP_FORMULA

	totalMod := float64(1.0)
	zemMod := float64(1.0)

	expMultiplier := RuleR("Character:ExpMultiplier")
	if expMultiplier >= 0 {
		totalMod *= expMultiplier
	}

	if false { //if(zone->IsHotzone())
		totalMod += RuleR("Zone:HotZoneBonus")
	}

	xp = int64(float64(xp) * totalMod * zemMod)

	return xp
}

func (c *Npc) CleanName() string {
	return CleanName(c.Name)
}

func (c *Npc) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]Schema)
	var field string
	var prop Schema
	for _, field = range requiredFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	jsRef := gojsonschema.NewGoLoader(s)
	schema, err = gojsonschema.NewSchema(jsRef)
	if err != nil {
		return
	}
	return
}

func (c *Npc) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "zoneId":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}

func (c *Npc) ClassName() string {
	return ClassName(c.Class)
}

func (c *Npc) RaceName() string {
	return RaceName(c.Race)
}
