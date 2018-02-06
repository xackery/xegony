package oauth

import (
	"github.com/xackery/xegony/model"
)

//Wrapper wraps different Oauth endpoints
type Wrapper interface {
	//Oauth
	GetLoginURL(state string) (redirectURL string)
	ExchangeToken(code string, userOauth *model.UserOauth) (err error)

	//People
	GetUser(user *model.User) (err error)
}
