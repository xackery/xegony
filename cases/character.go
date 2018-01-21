package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//CharacterRepository handles CharacterRepository cases and is a gateway to storage
type CharacterRepository struct {
	stor       storage.Storage
	globalRepo *GlobalRepository
}

func (c *CharacterRepository) isStorageInitialized() (err error) {
	if c.stor == nil {
		err = fmt.Errorf("Storage not initialized")
		return
	}
	return
}

//Initialize handles logic
func (c *CharacterRepository) Initialize(stor storage.Storage) (err error) {
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
func (c *CharacterRepository) Get(character *model.Character, user *model.User) (err error) {

	if err = c.isStorageInitialized(); err != nil {
		return
	}
	err = c.stor.GetCharacter(character)
	if err != nil {
		err = errors.Wrap(err, "failed to get character")
		return
	}

	err = c.prepare(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare character")
		return
	}
	return
}

//GetByName handles logic
func (c *CharacterRepository) GetByName(character *model.Character, user *model.User) (err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	err = c.stor.GetCharacterByName(character)
	if err != nil {
		err = errors.Wrap(err, "failed to get character")
		return
	}
	err = c.prepare(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare character")
		return
	}

	return
}

//Create handles logic
func (c *CharacterRepository) Create(character *model.Character, user *model.User) (err error) {
	if character == nil {
		err = fmt.Errorf("Empty character")
		return
	}
	schema, err := c.newSchema([]string{"name", "accountID"}, nil)
	if err != nil {
		return
	}
	if character.AccountID < 1 {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		vErr.Reasons["accountID"] = "Account ID must be greater than 0"
		err = vErr
		return
	}
	character.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(character))
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
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	err = c.stor.CreateCharacter(character)
	if err != nil {
		err = errors.Wrap(err, "failed to create character")
		return
	}
	err = c.prepare(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare character")
		return
	}
	return
}

//Search handles logic
func (c *CharacterRepository) SearchByName(character *model.Character, user *model.User) (characters []*model.Character, err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	characters, err = c.stor.SearchCharacterByName(character)
	if err != nil {
		err = errors.Wrap(err, "failed to search character")
		return
	}
	for _, character := range characters {
		err = c.prepare(character, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare character")
			return
		}
	}
	return
}

//Edit handles logic
func (c *CharacterRepository) Edit(character *model.Character, user *model.User) (err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(character))
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

	err = c.stor.EditCharacter(character)
	if err != nil {
		return
	}
	err = c.prepare(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare character")
		return
	}
	return
}

//Delete handles logic
func (c *CharacterRepository) Delete(character *model.Character, user *model.User) (err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	err = c.stor.DeleteCharacter(character)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *CharacterRepository) List(user *model.User) (characters []*model.Character, err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	characters, err = c.stor.ListCharacter()
	if err != nil {
		err = errors.Wrap(err, "failed to list characters")
		return
	}
	for _, character := range characters {
		err = c.prepare(character, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare character")
			return
		}
	}
	return
}

//ListByRanking handles logic
func (c *CharacterRepository) ListByRanking(user *model.User) (characters []*model.Character, err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	characters, err = c.stor.ListCharacterByRanking()
	if err != nil {
		if err != nil {
			err = errors.Wrap(err, "failed to list characters")
			return
		}
		return
	}
	for _, character := range characters {
		err = c.prepare(character, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare character")
			return
		}
		/*
			//needs to be moved to inside detail call
				character.Base = &model.Base{
					Level: character.Level,
					Class: character.ClassID,
				}
				err = s.GetBase(character.Base)
				if err != nil {
					err = errors.Wrap(err, "failed to get base data")
					return
				}*/

		character.Inventory, err = c.stor.ListItemByCharacter(character)
		if err != nil {
			err = errors.Wrap(err, "failed to get inventory")
			return
		}
	}

	return
}

//ListByOnline handles logic
func (c *CharacterRepository) ListByOnline(user *model.User) (characters []*model.Character, err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	characters, err = c.stor.ListCharacterByOnline()
	if err != nil {
		err = errors.Wrap(err, "failed to list characters")
		return
	}
	for _, character := range characters {
		err = c.prepare(character, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare character")
			return
		}
	}
	return
}

//ListByAccount handles logic
func (c *CharacterRepository) ListByAccount(account *model.Account, user *model.User) (characters []*model.Character, err error) {
	if err = c.isStorageInitialized(); err != nil {
		return
	}
	characters, err = c.stor.ListCharacterByAccount(account)
	if err != nil {
		err = errors.Wrap(err, "failed to list characters")
		return
	}
	for _, character := range characters {
		err = c.prepare(character, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare character")
			return
		}
	}
	return

}

func (c *CharacterRepository) prepare(character *model.Character, user *model.User) (err error) {

	character.Class, err = c.globalRepo.GetClass(character.ClassID, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get class of character")
		return
	}

	character.Race, err = c.globalRepo.GetRace(character.RaceID, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get race of character")
		return
	}

	zoneRepo := &ZoneRepository{}
	if err = zoneRepo.Initialize(c.stor); err != nil {
		err = errors.Wrap(err, "Failed to initialize zone of character")
		return
	}

	character.Zone, err = c.globalRepo.GetZone(character.ZoneID, user)
	if err != nil {
		err = errors.Wrap(err, "Failed to get zone of character")
		return
	}

	character.AASpent = c.aaSpent(character)
	if len(character.Inventory) > 0 && character.Base != nil { //don't bother calculating hp if no inventory known
		character.TotalHP = c.totalHP(character)
		character.TotalMana = c.totalMana(character)
	}

	return
}

func (c *CharacterRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *CharacterRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "accountID":
		prop.Type = "integer"
		prop.Minimum = 1
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

func (c *CharacterRepository) aaSpent(character *model.Character) int64 {
	return 0
}

//TotalHP returns total HP
func (c *CharacterRepository) totalHP(character *model.Character) int64 {
	var nd float64
	nd = 10000

	maxHp := c.baseHP(character) + c.itemBonusHP(character)

	//The AA desc clearly says it only applies to base hp..
	//but the actual effect sent on live causes the client
	//to apply it to (basehp + itemhp).. I will oblige to the client's whims over
	//the aa description
	nd += float64(c.aABonusMaxHP(character))   //Natural Durability, Physical Enhancement, Planar Durability
	maxHp = int64(float64(maxHp) * nd / 10000) //this is to fix the HP-above-495k issue
	//not needed for unbuffed?
	//maxHp += c.SpellBonusHP + AABonusHP
	maxHp += c.groupLeadershipBonusHP(character) //GroupLeadershipAAHealthEnhancement();
	//maxHp += maxHp * ((spellbonuses.MaxHPChange + itembonuses.MaxHPChange) / 10000.0f);

	return maxHp
}

//ItemBonusHP returns the total HP bonus from items
func (c *CharacterRepository) itemBonusHP(character *model.Character) int64 {
	var hp int64
	for _, item := range character.Inventory {
		if item.SlotID >= 0 && item.SlotID < 21 { //charm to one less than ammo
			hp += item.Hp
		}
		if item.SlotID == 22 { //powersource
			hp += item.Hp
		}
		//todo: tribute
	}
	return hp
}

//AABonusMaxHP returns bonus of HP from AAs
func (c *CharacterRepository) aABonusMaxHP(character *model.Character) int64 {
	return 0
}

//GroupLeadershipBonusHP returns how much hp bonus is being received from hp
func (c *CharacterRepository) groupLeadershipBonusHP(character *model.Character) int64 {
	return 0
}

//BaseHP on source
func (c *CharacterRepository) baseHP(character *model.Character) int64 {
	var baseHP int64
	stats := character.Sta

	if stats > 255 {
		stats = (stats - 255) / 2
		stats += 255
	}
	baseHP = 5

	if character.Base != nil {
		baseHP += int64(character.Base.Hp) + (int64(character.Base.HpFac) * stats)
		baseHP += (c.heroicSTA(character) * 10)
	}

	return baseHP
}

//HeroicSTA is based on GetHeroicSTA on source
func (c *CharacterRepository) heroicSTA(character *model.Character) int64 {
	return 0
}

//TotalMana returns mana
func (c *CharacterRepository) totalMana(character *model.Character) int64 {
	mana := character.Mana
	return mana
}

//ATK returns player attack
func (c *CharacterRepository) aTK(character *model.Character) int64 {
	atk := character.Dex
	return atk
}

//AC returns total AC
func (c *CharacterRepository) aC(character *model.Character) int64 {
	ac := character.Agi
	return ac
}

//HPRegen returns total hp regeneration
func (c *CharacterRepository) hPRegen(character *model.Character) int64 {
	return 0
}

//ManaRegen returns total mana regeneration
func (c *CharacterRepository) manaRegen(character *model.Character) int64 {
	return 0
}
