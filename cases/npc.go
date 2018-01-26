package cases

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//NpcRepository handles NpcRepository cases and is a gateway to storage
type NpcRepository struct {
	stor       storage.Storage
	globalRepo *GlobalRepository
}

//Initialize handles logic
func (c *NpcRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor

	c.globalRepo = &GlobalRepository{}
	err = c.globalRepo.Initialize(stor)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize global repository")
		return
	}
	return
}

//Get handles logic
func (c *NpcRepository) Get(npc *model.Npc, user *model.User) (err error) {
	err = c.stor.GetNpc(npc)
	if err != nil {
		err = errors.Wrap(err, "failed to get npc")
		return
	}

	err = c.prepare(npc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare npc")
		return
	}
	return
}

//Create handles logic
func (c *NpcRepository) Create(npc *model.Npc, user *model.User) (err error) {
	if npc == nil {
		err = fmt.Errorf("Empty npc")
		return
	}
	schema, err := c.newSchema([]string{"name", "accountID"}, nil)
	if err != nil {
		return
	}

	npc.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(npc))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}
	err = c.stor.CreateNpc(npc)
	if err != nil {
		return
	}
	err = c.prepare(npc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare npc")
		return
	}
	return
}

//SearchByName handles logic
func (c *NpcRepository) SearchByName(npc *model.Npc, user *model.User) (npcs []*model.Npc, err error) {
	npc.Name = strings.Replace(npc.Name, " ", "_", -1)
	npcs, err = c.stor.SearchNpcByName(npc)
	if err != nil {
		return
	}
	for _, npc := range npcs {
		err = c.prepare(npc, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare npc")
			return
		}
	}
	return
}

//Edit handles logic
func (c *NpcRepository) Edit(npc *model.Npc, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(npc))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}

	err = c.stor.EditNpc(npc)
	if err != nil {
		return
	}
	err = c.prepare(npc, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare npc")
		return
	}
	return
}

//Delete handles logic
func (c *NpcRepository) Delete(npc *model.Npc, user *model.User) (err error) {
	err = c.stor.DeleteNpc(npc)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *NpcRepository) List(pageSize int64, pageNumber int64, user *model.User) (npcs []*model.Npc, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	npcs, err = c.stor.ListNpc(pageSize, pageNumber)
	if err != nil {
		return
	}
	for _, npc := range npcs {
		err = c.prepare(npc, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare npc")
			return
		}
	}
	return
}

//ListCount handles logic
func (c *NpcRepository) ListCount(user *model.User) (count int64, err error) {

	count, err = c.stor.ListNpcCount()
	if err != nil {
		return
	}
	return
}

//ListByMerchant handles logic
func (c *NpcRepository) ListByMerchant(merchant *model.Merchant, user *model.User) (npcs []*model.Npc, err error) {
	npcs, err = c.stor.ListNpcByMerchant(merchant)
	if err != nil {
		return
	}
	for _, npc := range npcs {
		err = c.prepare(npc, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare npc")
			return
		}
	}
	return
}

//ListBySpell handles logic
func (c *NpcRepository) ListBySpell(spell *model.Spell, user *model.User) (npcs []*model.Npc, err error) {
	npcs, err = c.stor.ListNpcBySpell(spell)
	if err != nil {
		return
	}
	for _, npc := range npcs {
		err = c.prepare(npc, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare npc")
			return
		}
	}
	return
}

//ListByItem handles logic
func (c *NpcRepository) ListByItem(item *model.Item, user *model.User) (npcs []*model.Npc, err error) {
	npcs, err = c.stor.ListNpcByItem(item)
	if err != nil {
		return
	}
	for _, npc := range npcs {
		err = c.prepare(npc, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare npc")
			return
		}
	}
	return
}

//ListByZone handles logic
func (c *NpcRepository) ListByZone(zone *model.Zone, user *model.User) (npcs []*model.Npc, err error) {
	npcs, err = c.stor.ListNpcByZone(zone)
	if err != nil {
		return
	}
	for _, npc := range npcs {
		err = c.prepare(npc, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare npc")
			return
		}
	}
	return
}

//ListByFaction handles logic
func (c *NpcRepository) ListByFaction(faction *model.Faction, user *model.User) (npcs []*model.Npc, err error) {
	npcs, err = c.stor.ListNpcByFaction(faction)
	if err != nil {
		return
	}
	for _, npc := range npcs {
		err = c.prepare(npc, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare npc")
			return
		}
	}
	return
}

func (c *NpcRepository) prepare(npc *model.Npc, user *model.User) (err error) {

	npc.Race, err = c.globalRepo.GetRace(npc.RaceID, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get race")
		return
	}
	npc.Class, err = c.globalRepo.GetClass(npc.ClassID, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get class")
		return
	}

	zoneID := int64(npc.ID / 1000)
	npc.Zone, err = c.globalRepo.GetZone(zoneID, user)
	if err != nil {
		err = errors.Wrap(err, "Failed to get zone of npc")
		return
	}
	npc.CleanName = cleanName(npc.Name)
	npc.SpecialAbilitiesList = specialAbilitiesList(npc)
	if npc.Experience, err = c.experience(npc, user); err != nil {
		err = errors.Wrap(err, "failed to get experience on npc")
		return
	}

	return
}

func (c *NpcRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
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

func (c *NpcRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "zoneID":
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

func cleanName(name string) string {
	var re = regexp.MustCompile(`[^0-9A-Za-z_]+`)
	cleanName := strings.Replace(name, " ", "_", -1)
	cleanName = strings.Replace(cleanName, "#", "", -1)
	cleanName = strings.TrimSpace(re.ReplaceAllString(cleanName, ""))
	cleanName = strings.Replace(cleanName, "_", " ", -1)
	return cleanName
}

//SpecialAbilitiesList returns a list of special abilities
func specialAbilitiesList(npc *model.Npc) map[string]string {
	abilities := make(map[string]string)
	rawAbils := strings.Split(npc.SpecialAbilities.String, ",")
	for _, abil := range rawAbils {
		breakdown := strings.Split(abil, ",")
		key := ""
		description := ""
		var val int64
		if len(breakdown) < 1 {
			continue
		}

		switch breakdown[0] { //based on http://wiki.eqemulator.org/p?NPC_Special_Attacks
		case "1": //Summon
			key = "Summons"
			if len(breakdown) == 1 {
				if npc.Level < 50 { //even if flagged summon, won't summon if below summon level
					continue
				}
				description += fmt.Sprintf("you to them at lvl 50, 90%% or less HP, 6s cooldown, ")
			}
			if len(breakdown) > 1 {
				//todo: convert  breakdown to integer, see if npc leve is less
				if npc.Level < 50 { //even if flagged summon, won't summon if below summon level
					continue
				}
				//level enabled at
				description += fmt.Sprintf("you to them at lvl %s, ", breakdown[1])
			}
			if len(breakdown) > 2 {
				//todo: convert  breakdown to integer, see if npc leve is less
				if npc.Level < 50 { //even if flagged summon, won't summon if below summon level
					continue
				}
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

//Experience returns calculated experience to kill a npc
func (c *NpcRepository) experience(npc *model.Npc, user *model.User) (exp int64, err error) {

	ruleExpMultiplier, err := c.globalRepo.GetRule("Character:ExpMultiplier", user)
	if err != nil {
		err = errors.Wrap(err, "failed to get rule, exp multiplier")
		return
	}
	expMultiplier := ruleExpMultiplier.ValueFloat

	ruleHotZoneBonus, err := c.globalRepo.GetRule("Zone:HotZoneBonus", user)
	if err != nil {
		err = errors.Wrap(err, "failed to get rule, hot zone bonus")
		return
	}
	hotZoneBonus := ruleHotZoneBonus.ValueFloat

	xp := npc.Level * npc.Level * 75 * 35 / 10 //EXP_FORMULA

	totalMod := float64(1.0)
	zemMod := float64(1.0)

	if expMultiplier >= 0 {
		totalMod *= expMultiplier
	}

	if npc.Zone.Hotzone > 0 && hotZoneBonus > 0 {
		totalMod += hotZoneBonus
	}

	exp = int64(float64(xp) * totalMod * zemMod)
	return
}
