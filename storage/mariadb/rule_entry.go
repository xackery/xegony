package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	ruleEntryTable  = "rule_values"
	ruleEntryFields = "ruleset_id, rule_value, notes"
	ruleEntryBinds  = ":ruleset_id, :rule_value, :notes"
)

//GetRuleEntry will grab data from storage
func (s *Storage) GetRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {
	query := fmt.Sprintf("SELECT rule_name, %s FROM %s WHERE ruleset_id = ? AND rule_name = ?", ruleEntryFields, ruleEntryTable)
	err = s.db.Get(ruleEntry, query, rule.ID, ruleEntry.Name)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateRuleEntry will grab data from storage
func (s *Storage) CreateRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(rule_name, %s) VALUES (:rule_name, %s)", ruleEntryTable, ruleEntryFields, ruleEntryBinds)
	_, err = s.db.NamedExec(query, ruleEntry)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListRuleEntry will grab data from storage
func (s *Storage) ListRuleEntry(page *model.Page, rule *model.Rule) (ruleEntrys []*model.RuleEntry, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "rule_name"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT rule_name, %s FROM %s WHERE ruleset_id = ? ORDER BY %s LIMIT %d OFFSET %d", ruleEntryFields, ruleEntryTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query, rule.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		ruleEntry := model.RuleEntry{}
		if err = rows.StructScan(&ruleEntry); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		ruleEntrys = append(ruleEntrys, &ruleEntry)
	}
	return
}

//ListRuleEntryTotalCount will grab data from storage
func (s *Storage) ListRuleEntryTotalCount(rule *model.Rule) (count int64, err error) {
	query := fmt.Sprintf("SELECT count(ruleset_id) FROM %s WHERE ruleset_id = ?", ruleEntryTable)
	err = s.db.Get(&count, query, rule.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListRuleEntryBySearch will grab data from storage
func (s *Storage) ListRuleEntryBySearch(page *model.Page, rule *model.Rule, ruleEntry *model.RuleEntry) (ruleEntrys []*model.RuleEntry, err error) {

	field := ""

	if len(ruleEntry.Name) > 0 {
		field += `rule_name LIKE :rule_name OR`
		ruleEntry.Name = fmt.Sprintf("%%%s%%", ruleEntry.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]
	ruleEntry.RuleID = rule.ID

	query := fmt.Sprintf("SELECT rule_name, %s FROM %s WHERE %s AND ruleset_id = :ruleset_id LIMIT %d OFFSET %d", ruleEntryFields, ruleEntryTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, ruleEntry)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		ruleEntry := model.RuleEntry{}
		if err = rows.StructScan(&ruleEntry); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		ruleEntrys = append(ruleEntrys, &ruleEntry)
	}
	return
}

//ListRuleEntryBySearchTotalCount will grab data from storage
func (s *Storage) ListRuleEntryBySearchTotalCount(rule *model.Rule, ruleEntry *model.RuleEntry) (count int64, err error) {
	field := ""
	if len(ruleEntry.Name) > 0 {
		field += `rule_name LIKE :rule_name OR`
		ruleEntry.Name = fmt.Sprintf("%%%s%%", ruleEntry.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	ruleEntry.RuleID = rule.ID
	query := fmt.Sprintf("SELECT count(ruleset_id) FROM %s WHERE %s AND ruleset_id = :ruleset_id", ruleEntryTable, field)

	rows, err := s.db.NamedQuery(query, ruleEntry)
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

//EditRuleEntry will grab data from storage
func (s *Storage) EditRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {

	prevRuleEntry := &model.RuleEntry{
		Name: ruleEntry.Name,
	}
	err = s.GetRuleEntry(rule, prevRuleEntry)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous ruleEntry")
		return
	}

	field := ""
	if len(ruleEntry.Value) > 0 && prevRuleEntry.Value != ruleEntry.Value {
		field += "value = :value, "
	}
	if len(ruleEntry.Description.String) > 0 && prevRuleEntry.Description.String != ruleEntry.Description.String {
		field += "notes = :notes, "
	}
	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE ruleset_id = :ruleset_id AND rule_name = :rule_name", ruleEntryTable, field)
	result, err := s.db.NamedExec(query, ruleEntry)
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

//DeleteRuleEntry will grab data from storage
func (s *Storage) DeleteRuleEntry(rule *model.Rule, ruleEntry *model.RuleEntry) (err error) {

	query := fmt.Sprintf("DELETE FROM %s WHERE ruleset_id = ? AND rule_name = ?", ruleEntryTable)
	result, err := s.db.Exec(query, rule.ID, ruleEntry.Name)
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

//createTableRuleEntry will grab data from storage
func (s *Storage) createTableRuleEntry() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE rule_values (
  ruleset_id tinyint(3) unsigned NOT NULL DEFAULT '0',
  rule_name varchar(64) NOT NULL DEFAULT '',
  rule_value varchar(30) NOT NULL DEFAULT '',
  notes text,
  PRIMARY KEY (ruleset_id,rule_name),
  KEY ruleset_id (ruleset_id)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
