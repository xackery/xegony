package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

func (s *Storage) GetCharacter(characterId int64) (character *model.Character, err error) {
	character = &model.Character{}
	err = s.db.Get(character, "SELECT id, name, level, last_name, title, class, zone_id FROM character_data WHERE id = ?", characterId)
	if err != nil {
		return
	}
	return
}

func (s *Storage) CreateCharacter(character *model.Character) (err error) {
	if character == nil {
		err = fmt.Errorf("Must provide character")
		return
	}

	result, err := s.db.NamedExec(`INSERT INTO character_data(name, level, title, last_name, class, account_id, zone_id)
		VALUES (:name, :level, :title, :last_name, :class, :account_id, :zone_id)`, character)
	if err != nil {
		return
	}
	characterId, err := result.LastInsertId()
	if err != nil {
		return
	}
	character.Id = characterId
	return
}

func (s *Storage) ListCharacter() (characters []*model.Character, err error) {
	rows, err := s.db.Queryx(`SELECT id, name, level, last_name,  title, class, zone_id FROM character_data ORDER BY id DESC`)
	if err != nil {
		return
	}

	for rows.Next() {
		character := model.Character{}
		if err = rows.StructScan(&character); err != nil {
			return
		}
		characters = append(characters, &character)
	}
	return
}

func (s *Storage) ListCharacterByRanking() (characters []*model.Character, err error) {
	rows, err := s.db.Queryx(`SELECT id, name, level, last_name,  title, class, zone_id FROM character_data ORDER BY cur_hp DESC`)
	if err != nil {
		return
	}

	for rows.Next() {
		character := model.Character{}
		if err = rows.StructScan(&character); err != nil {
			return
		}
		characters = append(characters, &character)
	}
	return
}

func (s *Storage) ListCharacterByAccount(accountId int64) (characters []*model.Character, err error) {
	rows, err := s.db.Queryx(`SELECT id, name, level, last_name, title, class, zone_id FROM character_data WHERE account_id = ?`, accountId)
	if err != nil {
		return
	}

	for rows.Next() {
		character := model.Character{}
		if err = rows.StructScan(&character); err != nil {
			return
		}
		characters = append(characters, &character)
	}
	return
}

func (s *Storage) EditCharacter(characterId int64, character *model.Character) (err error) {
	character.Id = characterId
	result, err := s.db.NamedExec(`UPDATE character_data SET level=:level, last_name=:last_name, title=:title, class=:class, name=:name, zone_id=:zone_id WHERE id = :id`, character)
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

func (s *Storage) DeleteCharacter(characterId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM character_data WHERE id = ?`, characterId)
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

func (s *Storage) SearchCharacter(search string) (characters []*model.Character, err error) {
	rows, err := s.db.Queryx(`SELECT id, name, level, last_name,  title, class, zone_id FROM character_data WHERE name like ? ORDER BY id DESC`, "%"+search+"%")
	if err != nil {
		return
	}

	for rows.Next() {
		character := model.Character{}
		if err = rows.StructScan(&character); err != nil {
			return
		}
		characters = append(characters, &character)
	}
	return
}
