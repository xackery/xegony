package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	errorSets   = `scope=:scope, url=:url, message=:message, severity=:severity`
	errorFields = `scope, url, message, severity`
	errorBinds  = `:scope, :url, :message, :severity`
)

//GetError will grab data from storage
func (s *Storage) GetError(errStruct *model.Error) (err error) {
	err = s.db.Get(errStruct, fmt.Sprintf(`SELECT id, %s 
		FROM xegony_error 
		WHERE id = ?`, errorFields), errStruct.ID)
	if err != nil {
		return
	}
	return
}

//CreateError will grab data from storage
func (s *Storage) CreateError(errStruct *model.Error) (err error) {

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO xegony_error(%s)
		VALUES (%s)`, errorFields, errorBinds), errStruct)
	if err != nil {
		return
	}
	errorID, err := result.LastInsertId()
	if err != nil {
		return
	}
	errStruct.ID = errorID
	return
}

//ListErrorCount will grab data from storage
func (s *Storage) ListErrorCount() (count int64, err error) {
	err = s.db.Get(&count, `SELECT count(id) FROM xegony_error`)
	if err != nil {
		return
	}
	return
}

//ListError will grab data from storage
func (s *Storage) ListError(pageSize int64, pageNumber int64) (errors []*model.Error, err error) {
	query := fmt.Sprintf(`SELECT xegony_error.id, create_date,  %s 
		FROM xegony_error ORDER BY create_date DESC LIMIT %d OFFSET %d`, errorFields, pageSize, pageSize*pageNumber)
	rows, err := s.db.Queryx(query)
	if err != nil {
		return
	}

	for rows.Next() {
		error := model.Error{}
		if err = rows.StructScan(&error); err != nil {
			return
		}
		errors = append(errors, &error)
	}
	return
}

//ListErrorByScope will grab data from storage
func (s *Storage) ListErrorByScope(errStruct *model.Error) (errors []*model.Error, err error) {

	query := fmt.Sprintf(`SELECT xegony_error.id, create_date, %s 
		FROM xegony_error WHERE xegony_error.scope = ? ORDER BY create_date DESC`, errorFields)
	rows, err := s.db.Queryx(query, errStruct.Scope)
	if err != nil {
		return
	}

	for rows.Next() {
		error := model.Error{}
		if err = rows.StructScan(&error); err != nil {
			return
		}
		errors = append(errors, &error)
	}
	return
}

//SearchError will grab data from storage
func (s *Storage) SearchErrorByMessage(errStruct *model.Error) (errors []*model.Error, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT id, create_date, %s FROM errors 
		WHERE name like ? ORDER BY id DESC`, errorFields), "%"+errStruct.Message+"%")
	if err != nil {
		return
	}

	for rows.Next() {
		errorStruct := model.Error{}
		if err = rows.StructScan(&errorStruct); err != nil {
			return
		}
		errors = append(errors, &errorStruct)
	}
	return
}

//EditError will grab data from storage
func (s *Storage) EditError(errorStruct *model.Error) (err error) {
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE xegony_error SET %s WHERE id = :id`, errorSets), errorStruct)
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

//DeleteError will grab data from storage
func (s *Storage) DeleteError(errStruct *model.Error) (err error) {
	result, err := s.db.Exec(`DELETE FROM xegony_error WHERE id = ?`, errStruct.ID)
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

//createTableError will grab data from storage
func (s *Storage) createTableError() (err error) {
	_, err = s.db.Exec(`CREATE TABLE xegony_error (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  url varchar(32) NOT NULL DEFAULT '',
  scope varchar(32) NOT NULL DEFAULT '',
  message varchar(256) NOT NULL,
  severity int(10) unsigned NOT NULL DEFAULT '0',
  create_date datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	return
}
