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
	spellAnimationLock = sync.RWMutex{}
)

//GetSpellAnimation will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetSpellAnimation(spellAnimation *model.SpellAnimation) (err error) {
	spellAnimationLock.Lock()
	defer spellAnimationLock.Unlock()
	spellAnimationsDatabase, err := s.readSpellAnimationFile()
	if err != nil {
		return
	}
	for _, tmpSpellAnimation := range spellAnimationsDatabase {
		if tmpSpellAnimation.ID == spellAnimation.ID {
			*spellAnimation = *tmpSpellAnimation
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSpellAnimation will grab data from storage
func (s *Storage) CreateSpellAnimation(spellAnimation *model.SpellAnimation) (err error) {
	spellAnimationLock.Lock()
	defer spellAnimationLock.Unlock()
	spellAnimationsDatabase, err := s.readSpellAnimationFile()
	if err != nil {
		return
	}
	for _, tmpSpellAnimation := range spellAnimationsDatabase {
		if tmpSpellAnimation.ID == spellAnimation.ID {
			err = fmt.Errorf("spellAnimation already exists")
			return
		}
	}
	spellAnimationsDatabase = append(spellAnimationsDatabase, spellAnimation)
	err = s.writeSpellAnimationFile(spellAnimationsDatabase)
	if err != nil {
		return
	}
	return
}

//ListSpellAnimation will grab data from storage
func (s *Storage) ListSpellAnimation(page *model.Page) (spellAnimations []*model.SpellAnimation, err error) {
	spellAnimationLock.Lock()
	defer spellAnimationLock.Unlock()
	spellAnimationsDatabase, err := s.readSpellAnimationFile()
	if err != nil {
		return
	}

	spellAnimations = make([]*model.SpellAnimation, len(spellAnimationsDatabase))

	spellAnimations = spellAnimationsDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellAnimations, func(i, j int) bool {
			return spellAnimations[i].Name < spellAnimations[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(spellAnimations))
		}
	*/
	return
}

//ListSpellAnimationTotalCount will grab data from storage
func (s *Storage) ListSpellAnimationTotalCount() (count int64, err error) {
	spellAnimationLock.Lock()
	defer spellAnimationLock.Unlock()
	spellAnimationsDatabase, err := s.readSpellAnimationFile()
	if err != nil {
		return
	}
	count = int64(len(spellAnimationsDatabase))
	return
}

//ListSpellAnimationBySearch will grab data from storage
func (s *Storage) ListSpellAnimationBySearch(page *model.Page, spellAnimation *model.SpellAnimation) (spellAnimations []*model.SpellAnimation, err error) {
	spellAnimationLock.Lock()
	defer spellAnimationLock.Unlock()
	spellAnimationsDatabase, err := s.readSpellAnimationFile()
	if err != nil {
		return
	}
	if len(spellAnimation.Name) > 0 {
		for i := range spellAnimationsDatabase {
			if strings.Contains(spellAnimationsDatabase[i].Name, spellAnimation.Name) {
				spellAnimations = append(spellAnimations, spellAnimationsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellAnimations, func(i, j int) bool {
			return spellAnimations[i].Name < spellAnimations[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(spellAnimations))
	//}
	return
}

//ListSpellAnimationBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellAnimationBySearchTotalCount(spellAnimation *model.SpellAnimation) (count int64, err error) {
	spellAnimationLock.Lock()
	defer spellAnimationLock.Unlock()
	spellAnimationsDatabase, err := s.readSpellAnimationFile()
	if err != nil {
		return
	}

	spellAnimations := []*model.SpellAnimation{}
	if len(spellAnimation.Name) > 0 {
		for i := range spellAnimationsDatabase {
			if strings.Contains(spellAnimationsDatabase[i].Name, spellAnimation.Name) {
				spellAnimations = append(spellAnimations, spellAnimationsDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(spellAnimations))
	return
}

//EditSpellAnimation will grab data from storage
func (s *Storage) EditSpellAnimation(spellAnimation *model.SpellAnimation) (err error) {
	spellAnimationLock.Lock()
	defer spellAnimationLock.Unlock()
	spellAnimationsDatabase, err := s.readSpellAnimationFile()
	if err != nil {
		return
	}

	for i := range spellAnimationsDatabase {
		if spellAnimationsDatabase[i].ID == spellAnimation.ID {
			*spellAnimationsDatabase[i] = *spellAnimation
			err = s.writeSpellAnimationFile(spellAnimationsDatabase)
			if err != nil {
				return
			}
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSpellAnimation will grab data from storage
func (s *Storage) DeleteSpellAnimation(spellAnimation *model.SpellAnimation) (err error) {
	spellAnimationLock.Lock()
	defer spellAnimationLock.Unlock()
	spellAnimationsDatabase, err := s.readSpellAnimationFile()
	if err != nil {
		return
	}
	indexToDelete := 0
	for i := range spellAnimationsDatabase {
		if spellAnimationsDatabase[i].ID == spellAnimation.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	spellAnimationsDatabase[len(spellAnimationsDatabase)-1], spellAnimationsDatabase[indexToDelete] = spellAnimationsDatabase[indexToDelete], spellAnimationsDatabase[len(spellAnimationsDatabase)-1]
	spellAnimationsDatabase = spellAnimationsDatabase[:len(spellAnimationsDatabase)-1]
	err = s.writeSpellAnimationFile(spellAnimationsDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readSpellAnimationFile() (spellAnimations []*model.SpellAnimation, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			spellAnimations = loadSpellAnimationDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeSpellAnimationFile(spellAnimations)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default spellAnimation data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &spellAnimations)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeSpellAnimationFile(spellAnimations []*model.SpellAnimation) (err error) {

	bData, err := yaml.Marshal(spellAnimations)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal spellAnimations")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
