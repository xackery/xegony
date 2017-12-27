package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

type AuthClaims struct {
	IsAdmin      bool               `json:"isAdmin,omitempty"`
	IsModerator  bool               `json:"isModerator,omitempty"`
	OwnedLobbies map[int64][]string `json:"ownedLobbies,omitempty"`
	User         *model.User        `json:"user"`
	AccountId    int64              `json:"accountId"`
	jwt.StandardClaims
}

func (a *Api) PostLogin(w http.ResponseWriter, r *http.Request) {
	var err error

	user := &model.User{}
	err = decodeBody(r, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	user, err = a.userRepo.Login(user.Name, user.Password)
	if err != nil {
		err = errors.Wrap(err, "login failed")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	account, err := a.accountRepo.Get(user.AccountId)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("account not found for %s: %d", user.Name, user.AccountId))
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	expiresAt := time.Now().Add(time.Hour * 24).Unix()

	claims := AuthClaims{
		OwnedLobbies: map[int64][]string{1: []string{"asdb"}},
		User:         user,
		AccountId:    account.Id,
	}

	if account.Status >= 200 {
		claims.IsAdmin = true
		user.IsAdmin = true
	}

	if account.Status >= 100 {
		claims.IsModerator = true
		user.IsAdmin = true
	}

	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expiresAt,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	loginResponse := LoginResponse{
		ApiKey: tokenString,
		User:   user,
	}

	fmt.Println("Gave token", tokenString)

	writeData(w, r, loginResponse, http.StatusOK)
}

func IsLoggedIn(r *http.Request) (err error) {
	claims, err := GetAuthClaims(r)
	if err != nil {
		err = &model.ErrPermission{
			Message: err.Error(),
		}
		return
	}
	if claims.User.Id < 1 {
		err = &model.ErrPermission{
			Message: "Must be registered",
		}
		return
	}
	return
}

func IsAdmin(r *http.Request) (err error) {
	claims, err := GetAuthClaims(r)
	if err != nil {
		err = &model.ErrPermission{
			Message: err.Error(),
		}
		return
	}
	if !claims.IsAdmin {
		err = &model.ErrPermission{
			Message: "Administrator access required",
		}
		return
	}
	return
}

func IsModerator(r *http.Request) (err error) {
	claims, err := GetAuthClaims(r)
	if err != nil {
		err = &model.ErrPermission{
			Message: err.Error(),
		}
		return
	}
	if !claims.IsAdmin && !claims.IsModerator {
		err = &model.ErrPermission{
			Message: "Moderator access required",
		}
		return
	}
	return
}

func IsUserOwner(userId int64, r *http.Request) (err error) {
	claims, err := GetAuthClaims(r)
	if err != nil {
		err = &model.ErrPermission{
			Message: err.Error(),
		}
		return
	}

	if claims.IsAdmin || claims.IsModerator {
		return
	}

	if userId == claims.User.Id {
		return
	}

	err = &model.ErrPermission{
		Message: "Owner access required",
	}
	return
}

func GetAuthClaims(r *http.Request) (*AuthClaims, error) {
	tokens, ok := r.Header["Authorization"]
	token := ""
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
	}

	if token == "" {
		cookie, err := r.Cookie("apiKey")
		if err != nil {
			if err == http.ErrNoCookie {
				return nil, nil
			}
			return nil, err
		}
		token = cookie.String()
		if len(token) > 7 { //strip out apiKey=
			token = token[7:]
		}
	}

	if token == "" {
		return nil, fmt.Errorf("No Token Provided")
	}

	parsedToken, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Invalid Token: %s", err.Error())
	}

	if claims, ok := parsedToken.Claims.(*AuthClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid Token")
}
