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

func (s *Storage) insertTestVariable() (err error) {
	_, err = s.db.Exec(`
	INSERT INTO variables (varname, value, information, ts)
	VALUES
		('AAXPMod', '0.75', 'AA Experience multipler. Increase to increase exp rate', '2010-09-06 08:03:51'),
		('ACfail', '15', 'the percentage of time AC fails to protect. 0 would mean there was always some level of protection, 100 would mean AC has no affect. When AC fails, it will be possible to get a max dmg hit.', '2010-09-06 08:03:51'),
		('ACrandom', '20', '', '2010-09-06 08:03:51'),
		('ACreduction', '3', '', '2010-09-06 08:03:51'),
		('ailevel', '6', '', '2010-09-06 08:03:51'),
		('curInstFlagNum', '2002', 'Determines what instance flag will be handed out next', '2010-09-06 08:03:51'),
		('DBVersion', '070_pop', 'DB version info', '2010-09-06 08:03:51'),
		('decaytime 1 54', '480', 'Corpse decay time for Level\'s 1 to 54', '2010-09-06 08:03:51'),
		('decaytime 55 100', '1800', 'Corpse decay time for Level\'s 55 to 100', '2010-09-06 08:03:51'),
		('dfltInstZflag', '1000', 'Used to determine if a zone is instanced, must be 1000 or greater', '2010-09-06 08:03:51'),
		('disablecommandline', '0', 'Allow command lines to be run from world.exe | 0 - off | 1 - on |', '2010-09-06 08:03:51'),
		('Expansions', '4', 'Accessible expansions for each player', '2016-02-04 15:14:07'),
		('EXPMod', '0.75', 'Experience multipler. Increase to increase exp rate', '2010-09-06 08:03:51'),
		('GroupEXPBonus', '0.60', 'Experience multipler. Increase to increase group exp rate', '2010-09-06 08:03:51'),
		('GuildWars', '0', 'Enable Guild Wars Type Server | 0 - off | 1 - on |', '2010-09-06 08:03:51'),
		('holdzones', '0', 'Restart Crashed Zone Servers | 0 - off | 1 - on |', '2010-09-06 08:03:51'),
		('leavecorpses', '0', 'Players leave corpses | 0 - off | 1 - on |', '2010-09-06 08:03:51'),
		('loglevel', '1111', 'Commands,Merchants,Trades,Loot', '2016-01-17 22:14:18'),
		('Max_AAXP', '21626880', 'Max AA Experience', '2010-09-06 08:03:51'),
		('MerchantsKeepItems', '1', 'Merchants keep items sold to them | 0 - off | 1 - on |', '2010-09-06 08:03:51'),
		('MOTD', 'Last Patch: 2017-05-22', '', '2017-05-22 21:44:47'),
		('hotfix_name', '', '', '2017-06-01 14:09:54');
	`)
	if err != nil {
		err = errors.Wrap(err, "failed to insert user data")
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
