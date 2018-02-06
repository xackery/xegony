package model

// UserOauths is an array of UserOauth
// swagger:model
type UserOauths []*UserOauth

// UserOauth represents linked oauth entities
//
// swagger:model
type UserOauth struct {
	UserID            int64  `json:"userID,omitempty" db:"user_id"`
	OauthTypeID       int64  `json:"oauthTypeID,omitempty" db:"oauth_type_id"`
	OauthToken        string `json:"oauthToken,omitempty" db:"oauth_token"`
	OauthRefreshToken string `json:"oauthRefreshToken,omitempty" db:"oauth_refresh_token"`
}
