package google

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetRule will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetRule(rule *model.Rule) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//CreateRule will grab data from storage
func (s *Storage) CreateRule(rule *model.Rule) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRule will grab data from storage
func (s *Storage) ListRule(page *model.Page) (rules []*model.Rule, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRuleTotalCount will grab data from storage
func (s *Storage) ListRuleTotalCount() (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRuleBySearch will grab data from storage
func (s *Storage) ListRuleBySearch(page *model.Page, rule *model.Rule) (rules []*model.Rule, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//ListRuleBySearchTotalCount will grab data from storage
func (s *Storage) ListRuleBySearchTotalCount(rule *model.Rule) (count int64, err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//EditRule will grab data from storage
func (s *Storage) EditRule(rule *model.Rule) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}

//DeleteRule will grab data from storage
func (s *Storage) DeleteRule(rule *model.Rule) (err error) {
	err = fmt.Errorf("Not implemented")
	return
}
