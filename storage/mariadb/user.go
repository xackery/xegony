package mariadb

import (
	"fmt"
	"strings"

	"github.com/xackery/xegony/model"
	"golang.org/x/crypto/bcrypt"
)

//GetUser will grab data from storage
func (s *Storage) GetUser(user *model.User) (err error) {
	user = &model.User{}
	err = s.db.Get(user, "SELECT id, name, account_id, character_id FROM user WHERE id = ?", user.ID)
	if err != nil {
		return
	}
	return
}

//LoginUser will grab data from storage
func (s *Storage) LoginUser(user *model.User, passwordConfirm string) (err error) {
	user = &model.User{}
	err = s.db.Get(user, "SELECT id, name, password, account_id, character_id email FROM user WHERE name = ?", user.Name)
	if err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordConfirm)); err != nil {
		return
	}

	//don't expose password
	user.Password = ""
	return
}

//CreateUser will grab data from storage
func (s *Storage) CreateUser(user *model.User) (err error) {
	if user == nil {
		err = fmt.Errorf("Must provide user")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user.Password = string(hash)

	result, err := s.db.NamedExec(`INSERT INTO user(name, password, email, account_id, character_id)
		VALUES (:name, :password, :email, :account_id, :character_id)`, user)
	if err != nil {
		if strings.Index(err.Error(), "Error 1062:") == 0 {
			vErr := &model.ErrValidation{
				Message: "Duplicate entry",
			}
			field := err.Error()
			startPos := strings.Index(field, "for key") + 9
			if startPos > 0 && startPos < len(field) {
				field = field[startPos:]
				endPos := strings.Index(field, "'")
				if endPos > 0 && endPos < len(field) {
					field = field[0:endPos]
				}
				vErr.Reasons = map[string]string{}
				vErr.Reasons[field] = "duplicate entry"
				vErr.Message = fmt.Sprintf("%s: duplicate entry", field)
				err = vErr
			}
			return
		}
		return
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return
	}
	user.ID = userID
	//don't expose password
	user.Password = ""
	return
}

//ListUser will grab data from storage
func (s *Storage) ListUser() (users []*model.User, err error) {
	rows, err := s.db.Queryx(`SELECT id, name, email, account_id, character_id FROM user ORDER BY id DESC`)
	if err != nil {
		return
	}

	for rows.Next() {
		user := model.User{}
		if err = rows.StructScan(&user); err != nil {
			return
		}
		users = append(users, &user)
	}
	if len(users) == 0 {
		err = &model.ErrNoContent{}
		return
	}
	return
}

//EditUser will grab data from storage
func (s *Storage) EditUser(user *model.User) (err error) {
	result, err := s.db.NamedExec(`UPDATE user SET name=:name, email=:email, account_id=:account_id character_id=:character_id WHERE id = :id`, user)
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

//DeleteUser will grab data from storage
func (s *Storage) DeleteUser(user *model.User) (err error) {
	result, err := s.db.Exec(`DELETE FROM user WHERE id = ?`, user.ID)
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

//createTableUser will grab data from storage
func (s *Storage) createTableUser() (err error) {
	_, err = s.db.Exec(`CREATE TABLE user (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(32) NOT NULL DEFAULT '',
  account_id int(11) unsigned NOT NULL,
  character_id int(11) unsigned NOT NULL,
  email varchar(32) NOT NULL DEFAULT '',
  password varchar(80) NOT NULL DEFAULT '',
  last_modified timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  create_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  icon varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`)
	if err != nil {
		return
	}
	return
}
