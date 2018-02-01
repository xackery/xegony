package mariadb

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

const (
	accountTable  = "account"
	accountFields = "name, charname, sharedplat, password, status, lsaccount_id, gmspeed, revoked, karma, minilogin_ip, hideme, rulesflag, suspendeduntil, time_creation, expansion, ban_reason, suspend_reason"
	accountBinds  = ":name, :charname, :sharedplat, :password, :status, :lsaccount_id, :gmspeed, :revoked, :karma, :minilogin_ip, :hideme, :rulesflag, :suspendeduntil, :time_creation, :expansion, :ban_reason, :suspend_reason"
)

//GetAccount will grab data from storage
func (s *Storage) GetAccount(account *model.Account) (err error) {
	query := fmt.Sprintf("SELECT id, %s FROM %s WHERE id = ?", accountFields, accountTable)
	err = s.db.Get(account, query, account.ID)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//CreateAccount will grab data from storage
func (s *Storage) CreateAccount(account *model.Account) (err error) {
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", accountTable, accountFields, accountBinds)
	result, err := s.db.NamedExec(query, account)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	accountID, err := result.LastInsertId()
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	account.ID = accountID
	return
}

//ListAccount will grab data from storage
func (s *Storage) ListAccount(page *model.Page) (accounts []*model.Account, err error) {

	if len(page.OrderBy) < 1 {
		page.OrderBy = "id"
	}

	orderField := page.OrderBy
	if page.IsDescending > 0 {
		orderField += " DESC"
	} else {
		orderField += " ASC"
	}

	query := fmt.Sprintf("SELECT id, %s FROM %s ORDER BY %s LIMIT %d OFFSET %d", accountFields, accountTable, orderField, page.Limit, page.Limit*page.Offset)

	rows, err := s.db.Queryx(query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		account := model.Account{}
		if err = rows.StructScan(&account); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		accounts = append(accounts, &account)
	}
	return
}

//ListAccountTotalCount will grab data from storage
func (s *Storage) ListAccountTotalCount() (count int64, err error) {
	query := fmt.Sprintf("SELECT count(id) FROM %s", accountTable)
	err = s.db.Get(&count, query)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}
	return
}

//ListAccountBySearch will grab data from storage
func (s *Storage) ListAccountBySearch(page *model.Page, account *model.Account) (accounts []*model.Account, err error) {

	field := ""

	if len(account.Name) > 0 {
		field += `name LIKE :name OR`
		account.Name = fmt.Sprintf("%%%s%%", account.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT id, %s FROM %s WHERE %s LIMIT %d OFFSET %d", accountFields, accountTable, field, page.Limit, page.Limit*page.Offset)
	rows, err := s.db.NamedQuery(query, account)
	if err != nil {
		err = errors.Wrapf(err, "query: %s", query)
		return
	}

	for rows.Next() {
		account := model.Account{}
		if err = rows.StructScan(&account); err != nil {
			err = errors.Wrapf(err, "query: %s", query)
			return
		}
		accounts = append(accounts, &account)
	}
	return
}

//ListAccountBySearchTotalCount will grab data from storage
func (s *Storage) ListAccountBySearchTotalCount(account *model.Account) (count int64, err error) {
	field := ""
	if len(account.Name) > 0 {
		field += `name LIKE :name OR`
		account.Name = fmt.Sprintf("%%%s%%", account.Name)
	}

	if len(field) == 0 {
		err = fmt.Errorf("No parameters to search by provided")
		return
	}
	field = field[0 : len(field)-3]

	query := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", accountTable, field)

	rows, err := s.db.NamedQuery(query, account)
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

//EditAccount will grab data from storage
func (s *Storage) EditAccount(account *model.Account) (err error) {

	prevAccount := &model.Account{
		ID: account.ID,
	}
	err = s.GetAccount(prevAccount)
	if err != nil {
		err = errors.Wrap(err, "failed to get previous account")
		return
	}

	field := ""
	if len(account.Name) > 0 && prevAccount.Name != account.Name {
		field += "name = :name, "
	}

	if len(account.Charname) > 0 && prevAccount.Charname != account.Charname {
		field += "charname = :charname, "
	}

	if len(account.Name) > 0 && prevAccount.Name != account.Name {
		field += "name=:name, "
	}
	if len(account.Charname) > 0 && prevAccount.Charname != account.Charname {
		field += "charname=:charname, "
	}
	if account.Sharedplat > 0 && prevAccount.Sharedplat != account.Sharedplat {
		field += "sharedplat=:sharedplat, "
	}
	if len(account.Password) > 0 && prevAccount.Password != account.Password {
		field += "password=:password, "
	}
	if account.Status > 0 && prevAccount.Status != account.Status {
		field += "status=:status, "
	}
	if account.LsaccountID.Int64 > 0 && prevAccount.LsaccountID != account.LsaccountID {
		field += "lsaccount_id=:lsaccount_id, "
	}
	if account.Gmspeed > 0 && prevAccount.Gmspeed != account.Gmspeed {
		field += "gmspeed=:gmspeed, "
	}
	if account.Revoked > 0 && prevAccount.Revoked != account.Revoked {
		field += "revoked=:revoked, "
	}
	if account.Karma > 0 && prevAccount.Karma != account.Karma {
		field += "karma=:karma, "
	}
	if len(account.MiniloginIP) > 0 && prevAccount.MiniloginIP != account.MiniloginIP {
		field += "minilogin_ip=:minilogin_ip, "
	}
	if account.Hideme > 0 && prevAccount.Hideme != account.Hideme {
		field += "hideme=:hideme, "
	}
	if account.Rulesflag > 0 && prevAccount.Rulesflag != account.Rulesflag {
		field += "rulesflag=:rulesflag, "
	}
	if !account.Suspendeduntil.IsZero() && prevAccount.Suspendeduntil != account.Suspendeduntil {
		field += "suspendeduntil=:suspendeduntil, "
	}
	if account.TimeCreation > 0 && prevAccount.TimeCreation != account.TimeCreation {
		field += "time_creation=:time_creation, "
	}
	if account.Expansion > 0 && prevAccount.Expansion != account.Expansion {
		field += "expansion=:expansion, "
	}
	if len(account.BanReason.String) > 0 && prevAccount.BanReason != account.BanReason {
		field += "ban_reason=:ban_reason, "
	}
	if len(account.SuspendReason.String) > 0 && prevAccount.SuspendReason != account.SuspendReason {
		field += "suspend_reason=:suspend_reason, "
	}

	if len(field) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	field = field[0 : len(field)-2]

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", accountTable, field)
	result, err := s.db.NamedExec(query, account)
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

//DeleteAccount will grab data from storage
func (s *Storage) DeleteAccount(account *model.Account) (err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", accountTable)
	result, err := s.db.Exec(query, account.ID)
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

//createTableAccount will grab data from storage
func (s *Storage) createTableAccount() (err error) {
	_, err = s.db.Exec(`
CREATE TABLE account (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(30) NOT NULL DEFAULT '',
  charname varchar(64) NOT NULL DEFAULT '',
  sharedplat int(11) unsigned NOT NULL DEFAULT '0',
  password varchar(50) NOT NULL DEFAULT '',
  status int(5) NOT NULL DEFAULT '0',
  lsaccount_id int(11) unsigned DEFAULT NULL,
  gmspeed tinyint(3) unsigned NOT NULL DEFAULT '0',
  revoked tinyint(3) unsigned NOT NULL DEFAULT '0',
  karma int(5) unsigned NOT NULL DEFAULT '0',
  minilogin_ip varchar(32) NOT NULL DEFAULT '',
  hideme tinyint(4) NOT NULL DEFAULT '0',
  rulesflag tinyint(1) unsigned NOT NULL DEFAULT '0',
  suspendeduntil datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  time_creation int(10) unsigned NOT NULL DEFAULT '0',
  expansion tinyint(4) NOT NULL DEFAULT '0',
  ban_reason text,
  suspend_reason text,
  PRIMARY KEY (id),
  UNIQUE KEY name (name),
  UNIQUE KEY lsaccount_id (lsaccount_id)
) ENGINE=INNODB AUTO_INCREMENT=82152 DEFAULT CHARSET=latin1;`)
	if err != nil {
		return
	}
	return
}
