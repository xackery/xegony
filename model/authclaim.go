package model

import (
	"github.com/dgrijalva/jwt-go"
)

//AuthClaim wraps all token data
type AuthClaim struct {
	IsAdmin      bool               `json:"isAdmin,omitempty"`
	IsModerator  bool               `json:"isModerator,omitempty"`
	OwnedLobbies map[int64][]string `json:"ownedLobbies,omitempty"`
	User         *User              `json:"user"`
	jwt.StandardClaims
}
