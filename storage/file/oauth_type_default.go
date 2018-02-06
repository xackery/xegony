package file

import (
	"github.com/xackery/xegony/model"
)

func loadOauthTypeDefault() model.OauthTypes {
	return model.OauthTypes{
		{
			ID:        0,
			ShortName: "google",
			Name:      "Google",
		},
		{
			ID:        1,
			ShortName: "facebook",
			Name:      "Facebook",
		},
		{
			ID:        2,
			ShortName: "discord",
			Name:      "Discord",
		},
	}
}
