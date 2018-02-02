package api

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// UserGoogleStartRequest is a list of parameters used for initiating an oauth process
// swagger:parameters getUserGoogleStart
type UserGoogleStartRequest struct {
	// ReturnURL is the URL to return once Oauth completes
	// in: query
	// example: http://everzek.com
	ReturnURL string `json:"returnURL"`
}

// UserGoogleStartResponse is what endpoints respond with
// swagger:response
type UserGoogleStartResponse struct {
	RedirectURL string `json:"redirectURL"`
}

// UserGoogleCallbackResponse is what endpoints respond with. Typically this is just a 302 redirect.
// swagger:response
type UserGoogleCallbackResponse struct {
	RedirectURL string `json:"redirectURL"`
}

func userGoogleRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /user/google/start user getUserGoogleStart
		//
		// Start a google single sign on process
		//
		// Creates a single sign on for google chain
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//     - application/xml
		//     - application/yaml
		//
		//
		//     Responses:
		//       default: ErrInternal
		//       302: ErrRedirect
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetUserGoogleStart",
			"GET",
			"/user/google/start",
			getUserGoogleStart,
		},
		// swagger:route GET /user/google/callback user getUserGoogleCallback
		//
		// Works with a callback from google oauth
		//
		// Creates a single sign on for google chain
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//     - application/xml
		//     - application/yaml
		//
		//
		//     Responses:
		//       default: ErrInternal
		//       302: ErrRedirect
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetUserGoogleCallback",
			"GET",
			"/user/google/callback",
			getUserGoogleCallback,
		},
	}
	return
}

func getUserGoogleStart(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserGoogleStartRequest{
		ReturnURL: getQuery(r, "returnURL"),
	}
	if len(request.ReturnURL) == 0 {
		request.ReturnURL = fmt.Sprintf("%s%s/", cases.GetConfigForHTTP(), cases.GetConfigValue("apiSuffix"))
	}

	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)

	session, _ := cookieStore.Get(r, "sess")
	session.Values["state"] = state
	session.Values["returnURL"] = request.ReturnURL
	session.Save(r, w)

	redirectURL, err := cases.GetUserGoogleStart(state)
	if err != nil {
		err = errors.Wrap(err, "failed to get user google start")
		return
	}
	response := &UserGoogleStartResponse{
		RedirectURL: redirectURL,
	}
	content = response
	return
}

func getUserGoogleCallback(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	return
}
