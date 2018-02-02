package memory

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/xackery/xegony/model"
)

var (
	ruleEntrysDatabase = map[int64][]*model.RuleEntry{}
	ruleEntryLock      = sync.RWMutex{}
)

//GetRuleEntry will grab data from storage
func (s *Storage) GetRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {
	ruleEntryLock.RLock()
	defer ruleEntryLock.RUnlock()
	entryDB, ok := ruleEntrysDatabase[rule.ID]
	if !ok {

		err = &model.ErrNoContent{}
		return
	}
	for _, tmpRuleEntry := range entryDB {
		if tmpRuleEntry.Name == ruleEntry.Name {
			*ruleEntry = *tmpRuleEntry
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateRuleEntry will grab data from storage
func (s *Storage) CreateRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {
	ruleEntryLock.Lock()
	defer ruleEntryLock.Unlock()
	entryDB, ok := ruleEntrysDatabase[rule.ID]
	if !ok {
		ruleEntrysDatabase[rule.ID] = []*model.RuleEntry{}
	}
	for _, tmpRuleEntry := range entryDB {
		if tmpRuleEntry.Name == ruleEntry.Name {
			err = fmt.Errorf("ruleEntry already exists (%d %s)", ruleEntry.RuleID, ruleEntry.Name)
			return
		}
	}

	ruleEntrysDatabase[rule.ID] = append(ruleEntrysDatabase[rule.ID], ruleEntry)
	return
}

//ListRuleEntry will grab data from storage
func (s *Storage) ListRuleEntry(page *model.Page, rule *model.Rule) (ruleEntrys []*model.RuleEntry, err error) {
	ruleEntryLock.RLock()
	defer ruleEntryLock.RUnlock()

	entryDB, ok := ruleEntrysDatabase[rule.ID]
	if !ok {
		err = &model.ErrNoContent{}
		return
	}
	ruleEntrys = make([]*model.RuleEntry, len(entryDB))

	ruleEntrys = entryDB

	switch page.OrderBy {
	case "name":
		sort.Slice(ruleEntrys, func(i, j int) bool {
			return ruleEntrys[i].Name < ruleEntrys[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(ruleEntrys))
		}
	*/
	return
}

//ListRuleEntryTotalCount will grab data from storage
func (s *Storage) ListRuleEntryTotalCount(rule *model.Rule) (count int64, err error) {
	ruleEntryLock.RLock()
	defer ruleEntryLock.RUnlock()
	entryDB, ok := ruleEntrysDatabase[rule.ID]
	if !ok {
		return
	}
	count = int64(len(entryDB))
	return
}

//ListRuleEntryBySearch will grab data from storage
func (s *Storage) ListRuleEntryBySearch(page *model.Page, rule *model.Rule, ruleEntry *model.RuleEntry) (ruleEntrys []*model.RuleEntry, err error) {
	ruleEntryLock.RLock()
	defer ruleEntryLock.RUnlock()
	entryDB, ok := ruleEntrysDatabase[rule.ID]
	if !ok {
		err = &model.ErrNoContent{}
		return
	}
	if len(ruleEntry.Name) > 0 {
		for i := range entryDB {
			if strings.Contains(entryDB[i].Name, ruleEntry.Name) {
				ruleEntrys = append(ruleEntrys, entryDB[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(ruleEntrys, func(i, j int) bool {
			return ruleEntrys[i].Name < ruleEntrys[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(ruleEntrys))
	//}
	return
}

//ListRuleEntryBySearchTotalCount will grab data from storage
func (s *Storage) ListRuleEntryBySearchTotalCount(rule *model.Rule, ruleEntry *model.RuleEntry) (count int64, err error) {
	ruleEntryLock.RLock()
	defer ruleEntryLock.RUnlock()

	entryDB, ok := ruleEntrysDatabase[rule.ID]
	if !ok {
		err = &model.ErrNoContent{}
		return
	}
	ruleEntrys := []*model.RuleEntry{}
	if len(ruleEntry.Name) > 0 {
		for i := range entryDB {
			if strings.Contains(entryDB[i].Name, ruleEntry.Name) {
				ruleEntrys = append(ruleEntrys, entryDB[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(ruleEntrys))
	return
}

//EditRuleEntry will grab data from storage
func (s *Storage) EditRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {
	ruleEntryLock.Lock()
	defer ruleEntryLock.Unlock()
	entryDB, ok := ruleEntrysDatabase[rule.ID]
	if !ok {
		err = &model.ErrNoContent{}
		return
	}
	for i := range entryDB {
		if entryDB[i].Name == ruleEntry.Name {
			*entryDB[i] = *ruleEntry
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteRuleEntry will grab data from storage
func (s *Storage) DeleteRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {
	ruleEntryLock.Lock()
	defer ruleEntryLock.Unlock()
	indexToDelete := 0
	entryDB, ok := ruleEntrysDatabase[rule.ID]
	if !ok {
		err = &model.ErrNoContent{}
		return
	}
	for i := range entryDB {
		if entryDB[i].Name == ruleEntry.Name {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	entryDB[len(entryDB)-1], entryDB[indexToDelete] = entryDB[indexToDelete], entryDB[len(entryDB)-1]
	entryDB = entryDB[:len(entryDB)-1]
	return
}
