package google

import (
	"golang.org/x/oauth2"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/oauth"
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
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       scopes,
			Endpoint:     google.Endpoint,
		},
	}
	return
}

//GetLoginURL Retrieves a login url from oauth
func (o *Oauth) GetLoginURL(state string) (redirectURL string) {
	return
}

//ExchangeToken will upgrade to a more permanent token
func (o *Oauth) ExchangeToken(code string, userOauth *model.UserOauth) (err error) {
	//token, err = o.config.Exchange(ctx, code)
	return
}

//GetUser will fetch user details
func (o *Oauth) GetUser(user *model.User) (err error) {
	return
}
