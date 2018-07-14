package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	ruleTable  = "rule_sets"
	ruleFields = "name"
	ruleBinds  = ":name"
)

//GetRule will grab data from storage
func (s *Storage) GetRule(rule *model.Rule) (err error) {
	query := fmt.Sprintf("SELECT id, %s FROM %s WHERE ruleset_id = ?", ruleFields, ruleTable)
	err = s.db.Get(rule, query, rule.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateRule will grab data from storage
func (s *Storage) CreateRule(rule *model.Rule) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", ruleTable, ruleFields, ruleBinds)
	result, err := s.db.NamedExec(query, rule)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	ruleID, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	rule.ID = ruleID
	return
}

//ListRule will grab data from storage
func (s *Storage) ListRule(page *model.Page) (rules []*model.Rule, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "ruleset_id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT ruleset_id, %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", ruleFields, ruleTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		rule := model.Rule{}
		if err = rows.StructScan(&rule); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		rules = append(rules, &rule)
	}
	return
}

//ListRuleTotalCount will grab data from storage
func (s *Storage) ListRuleTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(ruleset_id) FROM %s", ruleTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListRuleBySearch will grab data from storage
func (s *Storage) ListRuleBySearch(page *model.Page, rule *model.Rule) (rules []*model.Rule, err error) {

	field := ""

	if len(rule.Name) > 0 {
		field += `name LIKE :name OR`
		rule.Name = fmt.Sprintf("%%%s%%", rule.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT ruleset_id, %s FROM %s WHERE %s LIMIT %d OFFSET %d", ruleFields, ruleTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, rule)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		rule := model.Rule{}
		if err = rows.StructScan(&rule); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		rules = append(rules, &rule)
	}
	return
}

//ListRuleBySearchTotalCount will grab data from storage
func (s *Storage) ListRuleBySearchTotalCount(rule *model.Rule) (count int64, err error) {
	field := ""
	if len(rule.Name) > 0 {
		field += `name LIKE :name OR`
		rule.Name = fmt.Sprintf("%%%s%%", rule.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(ruleset_id) FROM %s WHERE %s", ruleTable, field)

	rows, err := s.db.NamedQuery(query, rule)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
	}
	return
}

//EditRule will grab data from storage
func (s *Storage) EditRule(rule *model.Rule) (err error) {

	prevRule := &model.Rule{
		ID: rule.ID,
	}
	err = s.GetRule(prevRule)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous rule")
		return
	}

	field := ""
	if len(rule.Name) > 0 && prevRule.Name != rule.Name {
		field += "name = :name, "
	}
	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE ruleset_id = :ruleset_id", ruleTable, field)
	result, err := s.db.NamedExec(query, rule)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

func (s *Storage) insertTestRule() (err error) {

	_, err = s.db.Exec(`INSERT INTO rule_sets (ruleset_id, name)
VALUES
	(3, 'merc_test'),
	(1, 'default'),
	(2, 'pop+'),
	(10, 'EQEmu_Default'),
	(4, 'GoD'),
	(5, 'raidzone'),
	(6, 'OOW');`)
	if err != nil {
		err = errors.Wrap(err, "failed to insert npc data")
		return
	}
	return
}

//DeleteRule will grab data from storage
func (s *Storage) DeleteRule(rule *model.Rule) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE ruleset_id = ?", ruleTable)
	result, err := s.db.Exec(query, rule.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//createTableRule will grab data from storage
func (s *Storage) createTableRule() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE rule_sets (
  ruleset_id tinyint(3) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (ruleset_id)
) ENGINE=MyISAM AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
