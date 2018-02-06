package google

import (
	"golang.org/x/oauth2"

	"github.com/xackery/xegony/oauth"
	"github.com/xackery/xegony/cases"
	"golang.org/x/oauth2/google"
)

//Oauth implements the google oauth endpoint
type Oauth struct {
	config *oauth2.Config	
}

//New creates a new oauth instance
func New(clientID string, clientSecret string, redirectURL string, scopes []string) (db oauth.Wrapper, err error) {
	db = &Oauth{
		config: &oauth2.Config{
			ClientID: clientID,
			ClientSecret: clientSecret,
			RedirectURL: redirectURL,
			Scopes: scopes,
			Endpoint: google.Endpoint
		},
	}
	return
}

func (o *Oauth) GetLoginURL(state string) (redirectURL string) {
	return
}

func (o *Oauth) ExchangeToken(code string, userOauth *model.UserOauth) (err error) {
	token, err = o.config.Exchange(ctx, code)	
	return
}

//People
func (o *Oauth) GetUser(user *model.User) (err error) {
	return
}
