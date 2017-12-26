package mariadb

import (
	"fmt"
	"strings"

	"github.com/xackery/xegony/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *Storage) GetUser(userId int64) (user *model.User, err error) {
	user = &model.User{}
	err = s.db.Get(user, "SELECT id, name, account_id, FROM user WHERE id = ?", userId)
	if err != nil {
		return
	}
	return
}

func (s *Storage) LoginUser(username string, password string) (user *model.User, err error) {
	user = &model.User{}
	err = s.db.Get(user, "SELECT id, name, password, account_id, email FROM user WHERE name = ?", username)
	if err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return
	}

	//don't expose password
	user.Password = ""
	return
}

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

	result, err := s.db.NamedExec(`INSERT INTO user(name, password, email, account_id)
		VALUES (:name, :password, :email, :account_id)`, user)
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
	userId, err := result.LastInsertId()
	if err != nil {
		return
	}
	user.Id = userId
	//don't expose password
	user.Password = ""
	return
}

func (s *Storage) ListUser() (users []*model.User, err error) {
	rows, err := s.db.Queryx(`SELECT id, name, account_id FROM user ORDER BY id DESC`)
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

func (s *Storage) EditUser(userId int64, user *model.User) (err error) {
	user.Id = userId
	result, err := s.db.NamedExec(`UPDATE user SET name=:name, email=:email, account_id=:account_id WHERE id = :id`, user)
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

func (s *Storage) DeleteUser(userId int64) (err error) {
	result, err := s.db.Exec(`DELETE FROM user WHERE id = ?`, userId)
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
