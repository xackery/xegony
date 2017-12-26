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
	UserId       int64              `json:"userId"`
	jwt.StandardClaims
}

func (a *Api) PostLogin(w http.ResponseWriter, r *http.Request) {
	var err error
	type LoginRequest struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	loginRequest := &LoginRequest{}

	err = decodeBody(r, loginRequest)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	user, err := a.userRepo.Login(loginRequest.Name, loginRequest.Password)
	if err != nil {
		err = errors.Wrap(err, "login failed")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	expiresAt := time.Now().Add(time.Hour * 24).Unix()

	claims := AuthClaims{
		OwnedLobbies: map[int64][]string{1: []string{"asdb"}},
		UserId:       user.Id,
	}

	//if user.Isadmin > 0 {
	//	claims.IsAdmin = true
	//}
	//if user.Ismoderator > 0 {
	//	claims.IsModerator = true
	//}
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

	writeData(w, r, loginResponse, http.StatusOK)
}

func IsLoggedIn(r *http.Request) (err error) {
	claims, err := getAuthClaims(r)
	if err != nil {
		err = &model.ErrPermission{
			Message: err.Error(),
		}
		return
	}
	if claims.UserId < 1 {
		err = &model.ErrPermission{
			Message: "Must be registered",
		}
		return
	}
	return
}

func IsAdmin(r *http.Request) (err error) {
	claims, err := getAuthClaims(r)
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
	claims, err := getAuthClaims(r)
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
	claims, err := getAuthClaims(r)
	if err != nil {
		err = &model.ErrPermission{
			Message: err.Error(),
		}
		return
	}

	if claims.IsAdmin || claims.IsModerator {
		return
	}

	if userId == claims.UserId {
		return
	}

	err = &model.ErrPermission{
		Message: "Owner access required",
	}
	return
}

func getAuthClaims(r *http.Request) (*AuthClaims, error) {
	tokens, ok := r.Header["Authorization"]
	token := ""
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
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
