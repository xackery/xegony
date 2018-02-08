package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

var (
	skillLock = sync.RWMutex{}
)

//GetSkill will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetSkill(skill *model.Skill) (err error) {
	skillLock.Lock()
	defer skillLock.Unlock()
	skillsDatabase, err := s.readSkillFile()
	if err != nil {
		return
	}
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
	skillsDatabase, err := s.readSkillFile()
	if err != nil {
		return
	}
	for _, tmpSkill := range skillsDatabase {
		if tmpSkill.ID == skill.ID {
			err = fmt.Errorf("skill already exists")
			return
		}
	}
	skillsDatabase = append(skillsDatabase, skill)
	err = s.writeSkillFile(skillsDatabase)
	if err != nil {
		return
	}
	return
}

//ListSkill will grab data from storage
func (s *Storage) ListSkill(page *model.Page) (skills []*model.Skill, err error) {
	skillLock.Lock()
	defer skillLock.Unlock()
	skillsDatabase, err := s.readSkillFile()
	if err != nil {
		return
	}

	skills = make([]*model.Skill, len(skillsDatabase))

	skills = skillsDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
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

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(skills))
		}
	*/
	return
}

//ListSkillTotalCount will grab data from storage
func (s *Storage) ListSkillTotalCount() (count int64, err error) {
	skillLock.Lock()
	defer skillLock.Unlock()
	skillsDatabase, err := s.readSkillFile()
	if err != nil {
		return
	}
	count = int64(len(skillsDatabase))
	return
}

//ListSkillBySearch will grab data from storage
func (s *Storage) ListSkillBySearch(page *model.Page, skill *model.Skill) (skills []*model.Skill, err error) {
	skillLock.Lock()
	defer skillLock.Unlock()
	skillsDatabase, err := s.readSkillFile()
	if err != nil {
		return
	}
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
	skillLock.Lock()
	defer skillLock.Unlock()
	skillsDatabase, err := s.readSkillFile()
	if err != nil {
		return
	}

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
	skillsDatabase, err := s.readSkillFile()
	if err != nil {
		return
	}

	for i := range skillsDatabase {
		if skillsDatabase[i].ID == skill.ID {
			*skillsDatabase[i] = *skill
			err = s.writeSkillFile(skillsDatabase)
			if err != nil {
				return
			}
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
	skillsDatabase, err := s.readSkillFile()
	if err != nil {
		return
	}
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
	err = s.writeSkillFile(skillsDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readSkillFile() (skills []*model.Skill, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			skills = loadSkillDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeSkillFile(skills)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default skill data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &skills)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeSkillFile(skills []*model.Skill) (err error) {

	bData, err := yaml.Marshal(skills)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal skills")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
