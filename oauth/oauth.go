package oauth

import (
	"github.com/xackery/xegony/model"
)

type Wrapper interface {
	//Oauth
	GetLoginURL(state string) (redirectURL string)
	ExchangeToken(code string) (user *model.User, err error)

	//People
	GetUser(user *model.User) (err error)
}
