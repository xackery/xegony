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

var (
	signingKey = []byte("23jøˆdofijso")
)

func generateToken(user *model.User) (tokenString string, err error) {
	expiresAt := time.Now().Add(time.Hour * 24).Unix()

	claims := model.AuthClaim{
		OwnedLobbies: map[int64][]string{1: []string{"asdb"}},
		User:         user,
	}

	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expiresAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(signingKey)
	if err != nil {
		err = errors.Wrap(err, "failed to generate signed token")
		return
	}
	return
}

//GetAuthClaim gets token authorization data from request
func GetAuthClaim(r *http.Request) (*model.AuthClaim, error) {
	auth := &model.AuthClaim{
		User: &model.User{},
	}
	tokens, ok := r.Header["Authorization"]
	token := ""
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
	}

	if token == "" {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				return auth, nil
			}
			return auth, err
		}
		token = cookie.String()
		if len(token) > 6 { //strip out token=
			token = token[6:]
		}
	}

	if token == "" || token == "undefined" {
		return auth, fmt.Errorf("No Token Provided")
	}

	parsedToken, err := jwt.ParseWithClaims(token, &model.AuthClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return auth, errors.Wrap(err, "failed to parse token")
	}

	if claims, ok := parsedToken.Claims.(*model.AuthClaim); ok && parsedToken.Valid {
		return claims, nil
	}
	return auth, fmt.Errorf("Invalid Token")
}
