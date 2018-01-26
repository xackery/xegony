package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) loginRoutes() (routes []*route) {
	routes = []*route{
		{
			"PostLogin",
			"POST",
			"/login",
			a.postLogin,
		},
	}
	return
}

func (a *API) postLogin(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	type loginRequest struct {
		*model.User
		passwordConfirm string
	}
	request := &loginRequest{}

	err = decodeBody(r, request)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	user = request.User
	err = a.userRepo.Login(user, request.passwordConfirm)
	if err != nil {
		err = errors.Wrap(err, "login failed")
		return
	}

	account := &model.Account{
		ID: user.AccountID,
	}
	err = a.accountRepo.Get(account, user)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("account not found for %s: %d", user.Name, user.AccountID))
		return
	}

	expiresAt := time.Now().Add(time.Hour * 24).Unix()

	claims := model.AuthClaim{
		OwnedLobbies: map[int64][]string{1: []string{"asdb"}},
		User:         user,
	}

	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expiresAt,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return
	}

	loginResponse := loginResponse{
		APIKey: tokenString,
		User:   user,
	}

	fmt.Println("Gave token", tokenString)
	content = loginResponse
	return
}
