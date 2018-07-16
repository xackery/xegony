package memory

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetCharacter will grab data from storage
func (s *Storage) GetCharacter(character *model.Character) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateCharacter will grab data from storage
func (s *Storage) CreateCharacter(character *model.Character) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListCharacter will grab data from storage
func (s *Storage) ListCharacter(page *model.Page) (characters []*model.Character, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListCharacterTotalCount will grab data from storage
func (s *Storage) ListCharacterTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListCharacterByOnline will grab data from storage
func (s *Storage) ListCharacterByOnline(page *model.Page) (characters []*model.Character, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListCharacterByOnlineTotalCount will grab data from storage
func (s *Storage) ListCharacterByOnlineTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListCharacterBySearch will grab data from storage
func (s *Storage) ListCharacterBySearch(page *model.Page, character *model.Character) (characters []*model.Character, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListCharacterBySearchTotalCount will grab data from storage
func (s *Storage) ListCharacterBySearchTotalCount(character *model.Character) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditCharacter will grab data from storage
func (s *Storage) EditCharacter(character *model.Character) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteCharacter will grab data from storage
func (s *Storage) DeleteCharacter(character *model.Character) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
