package model

// OauthTypes is an array of OauthType
// swagger:model
type OauthTypes []*OauthType

// OauthType identifies the type of oauth
// swagger:model
type OauthType struct {
	ID        int64  `json:"ID,omitempty"`
	ShortName string `json:"shortName,omitempty"`
	Name      string `json:"name,omitempty"`
}
