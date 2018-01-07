package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

//GetAccount will grab data from storage
func (s *Storage) GetAccount(accountID int64) (account *model.Account, err error) {
	account = &model.Account{}
	err = s.db.Get(account, "SELECT id, name, status FROM account WHERE id = ?", accountID)
	if err != nil {
		return
	}
	return
}

//GetAccountByName will grab data from storage
func (s *Storage) GetAccountByName(name string) (account *model.Account, err error) {
	account = &model.Account{}
	err = s.db.Get(account, "SELECT id, name, status FROM account WHERE name = ?", name)
	if err != nil {
		return
	}
	return
}

//CreateAccount will grab data from storage
func (s *Storage) CreateAccount(account *model.Account) (err error) {
	if account == nil {
		err = fmt.Errorf("Must provide account")
		return
	}

	result, err := s.db.NamedExec(`INSERT INTO account(name, status)
		VALUES (:name, :status)`, account)
	if err != nil {
		return
	}
	accountID, err := result.LastInsertId()
	if err != nil {
		return
	}
	account.ID = accountID
	return
}

//ListAccount will grab data from storage
func (s *Storage) ListAccount() (accounts []*model.Account, err error) {
	rows, err := s.db.Queryx(`SELECT id, name, status FROM account ORDER BY id DESC`)
	if err != nil {
		return
	}

	for rows.Next() {
		account := model.Account{}
		if err = rows.StructScan(&account); err != nil {
			return
		}
		accounts = append(accounts, &account)
	}
	return
}

//EditAccount will grab data from storage
func (s *Storage) EditAccount(accountID int64, account *model.Account) (err error) {
	account.ID = accountID
	result, err := s.db.NamedExec(`UPDATE account SET name=:name, status=:status WHERE id = :id`, account)
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

//DeleteAccount will grab data from storage
func (s *Storage) DeleteAccount(accountID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM account WHERE id = ?`, accountID)
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
