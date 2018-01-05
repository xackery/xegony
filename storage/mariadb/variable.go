package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	variableSets   = `varname=:varname, value=:value, information=:information, ts=:ts`
	variableFields = `varname, value, information, ts`
	variableBinds  = `:varname, :value, :information, :ts`
)

//GetVariable will grab data from storage
func (s *Storage) GetVariable(variableName int64) (variable *model.Variable, err error) {
	variable = &model.Variable{}
	err = s.db.Get(variable, fmt.Sprintf(`SELECT %s 
		FROM variables
		WHERE name = ?`, variableFields), variableName)
	if err != nil {
		return
	}
	return
}

//CreateVariable will grab data from storage
func (s *Storage) CreateVariable(variable *model.Variable) (err error) {
	if variable == nil {
		err = fmt.Errorf("Must provide variable")
		return
	}

	_, err = s.db.NamedExec(fmt.Sprintf(`INSERT INTO variables(%s)
		VALUES (%s)`, variableFields, variableBinds), variable)
	if err != nil {
		return
	}
	return
}

//ListVariable will grab data from storage
func (s *Storage) ListVariable() (variables []*model.Variable, err error) {
	query := fmt.Sprintf(`SELECT %s 
		FROM variables`, variableFields)
	rows, err := s.db.Queryx(query)
	if err != nil {
		return
	}

	for rows.Next() {
		variable := model.Variable{}
		if err = rows.StructScan(&variable); err != nil {
			return
		}
		variables = append(variables, &variable)
	}
	return
}

//EditVariable will grab data from storage
func (s *Storage) EditVariable(variableName string, variable *model.Variable) (err error) {
	variable.Name = variableName
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE variables SET %s WHERE varname = :varname`, variableSets), variable)
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

//DeleteVariable will grab data from storage
func (s *Storage) DeleteVariable(variableName string) (err error) {
	result, err := s.db.Exec(`DELETE FROM variables WHERE name = ?`, variableName)
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

//createTableVariable will grab data from storage
func (s *Storage) createTableVariable() (err error) {
	_, err = s.db.Exec(`CREATE TABLE variables (
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
