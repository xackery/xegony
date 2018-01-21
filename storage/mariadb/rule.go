package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	ruleSets   = `ruleset_id=:ruleset_id, rule_name=:rule_name, rule_value=:rule_value, notes=:notes`
	ruleFields = `ruleset_id, rule_name, rule_value, notes`
	ruleBinds  = `:ruleset_id, :rule_name, :rule_value, :notes`
)

//GetRule will grab data from storage
func (s *Storage) GetRule(rule *model.Rule) (err error) {
	rule = &model.Rule{}
	err = s.db.Get(rule, fmt.Sprintf(`SELECT %s 
		FROM rule_values
		WHERE name = ?`, ruleFields), rule.Name)
	if err != nil {
		return
	}
	return
}

//CreateRule will grab data from storage
func (s *Storage) CreateRule(rule *model.Rule) (err error) {
	if rule == nil {
		err = fmt.Errorf("Must provide rule")
		return
	}

	_, err = s.db.NamedExec(fmt.Sprintf(`INSERT INTO rule_values(%s)
		VALUES (%s)`, ruleFields, ruleBinds), rule)
	if err != nil {
		return
	}
	return
}

//ListRule will grab data from storage
func (s *Storage) ListRule() (rules []*model.Rule, err error) {
	query := fmt.Sprintf(`SELECT %s 
		FROM rule_values`, ruleFields)
	rows, err := s.db.Queryx(query)
	if err != nil {
		return
	}

	for rows.Next() {
		rule := model.Rule{}
		if err = rows.StructScan(&rule); err != nil {
			return
		}
		rules = append(rules, &rule)
	}
	return
}

//EditRule will grab data from storage
func (s *Storage) EditRule(rule *model.Rule) (err error) {
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE rule_values SET %s WHERE varname = :varname`, ruleSets), rule)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}

//DeleteRule will grab data from storage
func (s *Storage) DeleteRule(rule *model.Rule) (err error) {
	result, err := s.db.Exec(`DELETE FROM rule_values WHERE name = ?`, rule.Name)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affected < 1 {
		err = &model.ErrNoContent{}
		return
	}
	return
}

//createTableRule will grab data from storage
func (s *Storage) createTableRule() (err error) {
	_, err = s.db.Exec(`CREATE TABLE rule_values (
  ruleset_id tinyint(3) unsigned NOT NULL DEFAULT '0',
  rule_name varchar(64) NOT NULL DEFAULT '',
  rule_value varchar(30) NOT NULL DEFAULT '',
  notes text,
  PRIMARY KEY (ruleset_id,rule_name),
  KEY ruleset_id (ruleset_id)
) ENGINE=INNODB DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
