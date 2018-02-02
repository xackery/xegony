package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	rulesDatabase = []*model.Rule{}
	ruleLock      = sync.RWMutex{}
)

//GetRule will grab data from storage
func (s *Storage) GetRule(rule *model.Rule) (err error) {
	ruleLock.RLock()
	defer ruleLock.RUnlock()
	for _, tmpRule := range rulesDatabase {
		if tmpRule.ID == rule.ID {
			*rule = *tmpRule
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateRule will grab data from storage
func (s *Storage) CreateRule(rule *model.Rule) (err error) {
	ruleLock.Lock()
	defer ruleLock.Unlock()
	for _, tmpRule := range rulesDatabase {
		if tmpRule.ID == rule.ID {
			err = fmt.Errorf("rule already exists")
			return
		}
	}
	rulesDatabase = append(rulesDatabase, rule)
	return
}

//ListRule will grab data from storage
func (s *Storage) ListRule(page *model.Page) (rules []*model.Rule, err error) {
	ruleLock.RLock()
	defer ruleLock.RUnlock()

	rules = make([]*model.Rule, len(rulesDatabase))

	rules = rulesDatabase

	switch page.OrderBy {
	case "name":
		sort.Slice(rules, func(i, j int) bool {
			return rules[i].Name < rules[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(rules))
		}
	*/
	return
}

//ListRuleTotalCount will grab data from storage
func (s *Storage) ListRuleTotalCount() (count int64, err error) {
	configLock.RLock()
	defer configLock.RUnlock()
	count = int64(len(rulesDatabase))
	return
}

//ListRuleBySearch will grab data from storage
func (s *Storage) ListRuleBySearch(page *model.Page, rule *model.Rule) (rules []*model.Rule, err error) {
	ruleLock.RLock()
	defer ruleLock.RUnlock()

	if len(rule.Name) > 0 {
		for i := range rulesDatabase {
			if strings.Contains(rulesDatabase[i].Name, rule.Name) {
				rules = append(rules, rulesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(rules, func(i, j int) bool {
			return rules[i].Name < rules[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(rules))
	//}
	return
}

//ListRuleBySearchTotalCount will grab data from storage
func (s *Storage) ListRuleBySearchTotalCount(rule *model.Rule) (count int64, err error) {
	ruleLock.RLock()
	defer ruleLock.RUnlock()

	rules := []*model.Rule{}
	if len(rule.Name) > 0 {
		for i := range rulesDatabase {
			if strings.Contains(rulesDatabase[i].Name, rule.Name) {
				rules = append(rules, rulesDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(rules))
	return
}

//EditRule will grab data from storage
func (s *Storage) EditRule(rule *model.Rule) (err error) {
	ruleLock.Lock()
	defer ruleLock.Unlock()
	for i := range rulesDatabase {
		if rulesDatabase[i].ID == rule.ID {
			*rulesDatabase[i] = *rule
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteRule will grab data from storage
func (s *Storage) DeleteRule(rule *model.Rule) (err error) {
	ruleLock.Lock()
	defer ruleLock.Unlock()
	indexToDelete := 0
	for i := range rulesDatabase {
		if rulesDatabase[i].ID == rule.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	rulesDatabase[len(rulesDatabase)-1], rulesDatabase[indexToDelete] = rulesDatabase[indexToDelete], rulesDatabase[len(rulesDatabase)-1]
	rulesDatabase = rulesDatabase[:len(rulesDatabase)-1]
	return
}
