package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	variableTable  = "variables"
	variableFields = "varname, value, information, ts"
	variableBinds  = ":varname, :value, :information, :ts"
)

//GetVariable will grab data from storage
func (s *Storage) GetVariable(variable *model.Variable) (err error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE varname = ?", variableFields, variableTable)
	err = s.db.Get(variable, query, variable.Name)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateVariable will grab data from storage
func (s *Storage) CreateVariable(variable *model.Variable) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", variableTable, variableFields, variableBinds)
	_, err = s.db.NamedExec(query, variable)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListVariable will grab data from storage
func (s *Storage) ListVariable(page *model.Page) (variables []*model.Variable, err error) {
	if len(page.OrderBy) < 1 {
		page.OrderBy = "name"
	}

	orderField := page.OrderBy
	switch orderField {
	case "name":
		orderField = "varname"
	default:
		err = fmt.Errorf("unknown orderBy: %s", orderField)
		return
	}

	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", variableFields, variableTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		variable := model.Variable{}
		if err = rows.StructScan(&variable); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		variables = append(variables, &variable)
	}
	return
}

//ListVariableTotalCount will grab data from storage
func (s *Storage) ListVariableTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(varname) FROM %s", variableTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListVariableBySearch will grab data from storage
func (s *Storage) ListVariableBySearch(page *model.Page, variable *model.Variable) (variables []*model.Variable, err error) {

	field := ""

	if len(variable.Name) > 0 {
		field += `varname LIKE :varname OR`
		variable.Name = fmt.Sprintf("%%%s%%", variable.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d OFFSET %d", variableFields, variableTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, variable)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		variable := model.Variable{}
		if err = rows.StructScan(&variable); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		variables = append(variables, &variable)
	}
	return
}

//ListVariableBySearchTotalCount will grab data from storage
func (s *Storage) ListVariableBySearchTotalCount(variable *model.Variable) (count int64, err error) {
	field := ""
	if len(variable.Name) > 0 {
		field += `varname LIKE :varname OR`
		variable.Name = fmt.Sprintf("%%%s%%", variable.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(varname) FROM %s WHERE %s", variableTable, field)

	rows, err := s.db.NamedQuery(query, variable)
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

//EditVariable will grab data from storage
func (s *Storage) EditVariable(variable *model.Variable) (err error) {

	prevVariable := &model.Variable{
		Name: variable.Name,
	}
	err = s.GetVariable(prevVariable)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous variable")
		return
	}

	field := ""
	if len(variable.Name) > 0 && prevVariable.Name != variable.Name {
		field += "varname = :varname, "
	}
	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE varname = :varname", variableTable, field)
	result, err := s.db.NamedExec(query, variable)
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

//DeleteVariable will grab data from storage
func (s *Storage) DeleteVariable(variable *model.Variable) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE varname = ?", variableTable)
	result, err := s.db.Exec(query, variable.Name)
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

//createTableVariable will grab data from storage
func (s *Storage) createTableVariable() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE variables (
  varname varchar(25) NOT NULL DEFAULT '',
  value text NOT NULL,
  information text NOT NULL,
  ts timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (varname)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
