package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetSkill will grab data from storage
func (s *Storage) GetSkill(skill *model.Skill) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateSkill will grab data from storage
func (s *Storage) CreateSkill(skill *model.Skill) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSkill will grab data from storage
func (s *Storage) ListSkill(page *model.Page) (skills []*model.Skill, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSkillTotalCount will grab data from storage
func (s *Storage) ListSkillTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSkillBySearch will grab data from storage
func (s *Storage) ListSkillBySearch(page *model.Page, skill *model.Skill) (skills []*model.Skill, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListSkillBySearchTotalCount will grab data from storage
func (s *Storage) ListSkillBySearchTotalCount(skill *model.Skill) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditSkill will grab data from storage
func (s *Storage) EditSkill(skill *model.Skill) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteSkill will grab data from storage
func (s *Storage) DeleteSkill(skill *model.Skill) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
