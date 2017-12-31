package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type CharacterRepository struct {
	stor storage.Storage
}

func (g *CharacterRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *CharacterRepository) Get(characterID int64) (character *model.Character, err error) {
	if characterID == 0 {
		err = fmt.Errorf("Invalid Character ID")
		return
	}
	character, err = g.stor.GetCharacter(characterID)
	return
}

func (g *CharacterRepository) Create(character *model.Character) (err error) {
	if character == nil {
		err = fmt.Errorf("Empty character")
		return
	}
	schema, err := character.NewSchema([]string{"name", "accountID"}, nil)
	if err != nil {
		return
	}
	if character.AccountId < 1 {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		vErr.Reasons["accountID"] = "Account ID must be greater than 0"
		err = vErr
		return
	}
	character.Id = 0 //strip ID
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
	err = g.stor.CreateCharacter(character)
	if err != nil {
		return
	}
	return
}

func (g *CharacterRepository) Search(search string) (characters []*model.Character, err error) {
	characters, err = g.stor.SearchCharacter(search)
	if err != nil {
		return
	}
	return
}

func (g *CharacterRepository) Edit(characterID int64, character *model.Character) (err error) {
	schema, err := character.NewSchema([]string{"name"}, nil)
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

	err = g.stor.EditCharacter(characterID, character)
	if err != nil {
		return
	}
	return
}

func (g *CharacterRepository) Delete(characterID int64) (err error) {
	err = g.stor.DeleteCharacter(characterID)
	if err != nil {
		return
	}
	return
}

func (g *CharacterRepository) List() (characters []*model.Character, err error) {
	characters, err = g.stor.ListCharacter()
	if err != nil {
		return
	}
	return
}

func (g *CharacterRepository) ListByRanking() (characters []*model.Character, err error) {
	characters, err = g.stor.ListCharacterByRanking()
	if err != nil {
		return
	}
	return
}

func (g *CharacterRepository) ListByOnline() (characters []*model.Character, err error) {
	characters, err = g.stor.ListCharacterByOnline()
	if err != nil {
		return
	}
	return
}

func (g *CharacterRepository) ListByAccount(accountID int64) (characters []*model.Character, err error) {
	characters, err = g.stor.ListCharacterByAccount(accountID)
	if err != nil {
		return
	}
	return
}
