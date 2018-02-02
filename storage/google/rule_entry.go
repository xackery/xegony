package google

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetRuleEntry will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateRuleEntry will grab data from storage
func (s *Storage) CreateRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRuleEntry will grab data from storage
func (s *Storage) ListRuleEntry(page *model.Page, rule *model.Rule) (ruleEntrys []*model.RuleEntry, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRuleEntryTotalCount will grab data from storage
func (s *Storage) ListRuleEntryTotalCount(rule *model.Rule) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRuleEntryBySearch will grab data from storage
func (s *Storage) ListRuleEntryBySearch(page *model.Page, rule *model.Rule, ruleEntry *model.RuleEntry) (ruleEntrys []*model.RuleEntry, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRuleEntryBySearchTotalCount will grab data from storage
func (s *Storage) ListRuleEntryBySearchTotalCount(rule *model.Rule, ruleEntry *model.RuleEntry) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditRuleEntry will grab data from storage
func (s *Storage) EditRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteRuleEntry will grab data from storage
func (s *Storage) DeleteRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
