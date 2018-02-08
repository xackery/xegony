package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	skillsDatabase = []*model.Skill{}
	skillLock      = sync.RWMutex{}
)

//GetSkill will grab data from storage
func (s *Storage) GetSkill(skill *model.Skill) (err error) {
	skillLock.RLock()
	defer skillLock.RUnlock()
	for _, tmpSkill := range skillsDatabase {
		if tmpSkill.ID == skill.ID {
			*skill = *tmpSkill
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSkill will grab data from storage
func (s *Storage) CreateSkill(skill *model.Skill) (err error) {
	skillLock.Lock()
	defer skillLock.Unlock()
	for _, tmpSkill := range skillsDatabase {
		if tmpSkill.ID == skill.ID {
			err = fmt.Errorf("skill already exists")
			return
		}
	}
	skillsDatabase = append(skillsDatabase, skill)
	return
}

//ListSkill will grab data from storage
func (s *Storage) ListSkill(page *model.Page) (skills []*model.Skill, err error) {
	skillLock.RLock()
	defer skillLock.RUnlock()

	skills = make([]*model.Skill, len(skillsDatabase))

	skills = skillsDatabase

	switch page.OrderBy {
	case "name":
		sort.Slice(skills, func(i, j int) bool {
			return skills[i].Name < skills[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(skills))
		}
	*/
	return
}

//ListSkillTotalCount will grab data from storage
func (s *Storage) ListSkillTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(skillsDatabase))
	return
}

//ListSkillBySearch will grab data from storage
func (s *Storage) ListSkillBySearch(page *model.Page, skill *model.Skill) (skills []*model.Skill, err error) {
	skillLock.RLock()
	defer skillLock.RUnlock()

	if len(skill.Name) > 0 {
		for i := range skillsDatabase {
			if strings.Contains(skillsDatabase[i].Name, skill.Name) {
				skills = append(skills, skillsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(skills, func(i, j int) bool {
			return skills[i].Name < skills[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(skills))
	//}
	return
}

//ListSkillBySearchTotalCount will grab data from storage
func (s *Storage) ListSkillBySearchTotalCount(skill *model.Skill) (count int64, err error) {
	skillLock.RLock()
	defer skillLock.RUnlock()

	skills := []*model.Skill{}
	if len(skill.Name) > 0 {
		for i := range skillsDatabase {
			if strings.Contains(skillsDatabase[i].Name, skill.Name) {
				skills = append(skills, skillsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(skills))
	return
}

//EditSkill will grab data from storage
func (s *Storage) EditSkill(skill *model.Skill) (err error) {
	skillLock.Lock()
	defer skillLock.Unlock()
	for i := range skillsDatabase {
		if skillsDatabase[i].ID == skill.ID {
			*skillsDatabase[i] = *skill
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSkill will grab data from storage
func (s *Storage) DeleteSkill(skill *model.Skill) (err error) {
	skillLock.Lock()
	defer skillLock.Unlock()
	indexToDelete := 0
	for i := range skillsDatabase {
		if skillsDatabase[i].ID == skill.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	skillsDatabase[len(skillsDatabase)-1], skillsDatabase[indexToDelete] = skillsDatabase[indexToDelete], skillsDatabase[len(skillsDatabase)-1]
	skillsDatabase = skillsDatabase[:len(skillsDatabase)-1]
	return
}
