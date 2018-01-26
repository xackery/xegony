package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
)

//GlobalRepository handles GlobalRepository cases and is a gateway to storage
type GlobalRepository struct {
	stor            storage.Storage
	classRepository *ClassRepository
	raceRepository  *RaceRepository
	ruleRepository  *RuleRepository
	skillRepository *SkillRepository
	zoneRepository  *ZoneRepository
}

//Initialize handles logic
func (c *GlobalRepository) Initialize(stor storage.Storage) (err error) {

	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
	c.classRepository = &ClassRepository{}
	err = c.classRepository.Initialize(stor)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize class repository")
		return
	}

	c.ruleRepository = &RuleRepository{}
	err = c.ruleRepository.Initialize(stor)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize rule repository")
		return
	}
	c.raceRepository = &RaceRepository{}
	err = c.raceRepository.Initialize(stor)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize race repository")
		return
	}
	c.skillRepository = &SkillRepository{}
	err = c.skillRepository.Initialize(stor)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize skill repository")
		return
	}
	c.zoneRepository = &ZoneRepository{}
	err = c.zoneRepository.Initialize(stor)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize zone repository")
		return
	}
	return
}

func (c *GlobalRepository) GetClass(classID int64, user *model.User) (class *model.Class, err error) {
	class = &model.Class{
		ID: classID,
	}
	err = c.classRepository.Get(class, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get class")
		return
	}
	return
}

func (c *GlobalRepository) GetRace(raceID int64, user *model.User) (race *model.Race, err error) {
	race = &model.Race{
		ID: raceID,
	}
	err = c.raceRepository.Get(race, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get race")
		return
	}
	return
}

//GetRule gets a rule
func (c *GlobalRepository) GetRule(name string, user *model.User) (rule *model.Rule, err error) {
	rule = &model.Rule{
		Name: name,
	}
	err = c.ruleRepository.Get(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get rule")
		return
	}
	return
}

//GetSkill gets a skill
func (c *GlobalRepository) GetSkill(skillID int64, user *model.User) (skill *model.Skill, err error) {
	skill = &model.Skill{
		ID: skillID,
	}
	err = c.skillRepository.Get(skill, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get skill")
		return
	}
	return
}

//GetZone gets a zone
func (c *GlobalRepository) GetZone(zoneID int64, user *model.User) (zone *model.Zone, err error) {
	zone = &model.Zone{
		ZoneIDNumber: zoneID,
	}
	err = c.zoneRepository.Get(zone, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get zone")
		return
	}
	return
}
