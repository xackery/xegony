package mariadb

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

const (
	mailFields = `charid, timestamp, from, subject, body, to, status,`
	mailSets   = `charid=:charid, timestamp=:timestamp, from=:from, subject=:subject, body=:body, to=:to, status=:status,`
	mailBinds  = `:charid, :timestamp, :from, :subject, :body, :to, :status,`
)

//GetMail will grab data from storage
func (s *Storage) GetMail(mailID int64) (mail *model.Mail, err error) {
	mail = &model.Mail{}
	err = s.db.Get(mail, fmt.Sprintf("SELECT msgid, %s FROM mail WHERE msgid = ?", mailFields), mailID)
	if err != nil {
		return
	}
	return
}

//CreateMail will grab data from storage
func (s *Storage) CreateMail(mail *model.Mail) (err error) {
	if mail == nil {
		err = fmt.Errorf("Must provide mail")
		return
	}

	result, err := s.db.NamedExec(fmt.Sprintf(`INSERT INTO mail(%s)
		VALUES (%s)`, mailFields, mailBinds), mail)
	if err != nil {
		return
	}
	mailID, err := result.LastInsertId()
	if err != nil {
		return
	}
	mail.ID = mailID
	return
}

//ListMail will grab data from storage
func (s *Storage) ListMail(pageSize int64, pageNumber int64) (mails []*model.Mail, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT msgid, %s FROM mail 
		ORDER BY time DESC LIMIT %d OFFSET %d`, mailFields, pageSize, pageSize*pageNumber))
	if err != nil {
		return
	}

	for rows.Next() {
		mail := model.Mail{}
		if err = rows.StructScan(&mail); err != nil {
			return
		}
		mails = append(mails, &mail)
	}
	return
}

//ListMailCount will grab data from storage
func (s *Storage) ListMailCount() (count int64, err error) {
	err = s.db.Get(&count, `SELECT count(msgid) FROM mail`)
	if err != nil {
		return
	}
	return
}

//SearchMail will grab data from storage
func (s *Storage) SearchMail(search string) (mails []*model.Mail, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT msgid, %s FROM mail
		WHERE body like ? ORDER BY msgid DESC`, mailFields), "%"+search+"%")
	if err != nil {
		return
	}

	for rows.Next() {
		mail := model.Mail{}
		if err = rows.StructScan(&mail); err != nil {
			return
		}
		mails = append(mails, &mail)
	}
	return
}

//SearchMailByCharacter will grab data from storage
func (s *Storage) SearchMailByCharacter(characterID int64, search string) (mails []*model.Mail, err error) {

	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT msgid, %s FROM mail 
		WHERE mails.name like ? AND mail.charid = ? ORDER BY time DESC`, mailFields), "%"+search+"%", characterID)
	if err != nil {
		return
	}

	for rows.Next() {
		mail := &model.Mail{}

		if err = rows.StructScan(&mail); err != nil {
			return
		}
		mails = append(mails, mail)
	}
	return
}

//ListMailByCharacter will grab data from storage
func (s *Storage) ListMailByCharacter(characterID int64) (mails []*model.Mail, err error) {
	rows, err := s.db.Queryx(fmt.Sprintf(`SELECT msgid, %s FROM mail
		WHERE mail.charid = ? ORDER BY time DESC`, mailFields), characterID)
	if err != nil {
		return
	}

	for rows.Next() {
		mail := model.Mail{}
		if err = rows.StructScan(&mail); err != nil {
			return
		}
		mails = append(mails, &mail)
	}
	return
}

//EditMail will grab data from storage
func (s *Storage) EditMail(mailID int64, mail *model.Mail) (err error) {
	mail.ID = mailID
	result, err := s.db.NamedExec(fmt.Sprintf(`UPDATE mail SET %s WHERE msgid = :msgid`, mailSets), mail)
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

//DeleteMail will grab data from storage
func (s *Storage) DeleteMail(mailID int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM mail WHERE msgid = ?`, mailID)
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
