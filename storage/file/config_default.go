package file

import (
	"github.com/xackery/xegony/model"
)

func loadConfigDefault() model.Configs {
	return model.Configs{
		{
			Category:    "Google",
			Key:         "googleSecret",
			Value:       "",
			Description: "Google Secret Token, used for SSO/Google endpoints (by default empty)",
		},
		{
			Category:    "Google",
			Key:         "googleToken",
			Value:       "",
			Description: "Google Auth Token, used for SSO/Google endpoints (by default empty)",
		},
		{
			Category:    "MySQL",
			Key:         "mysqlUsername",
			Value:       "eqemu",
			Description: "username used for database (default: eqemu)",
		},
		{
			Category:    "MySQL",
			Key:         "mysqlPassword",
			Value:       "eqemu",
			Description: "password used for database (default: eqemu)",
		},
		{
			Category:    "MySQL",
			Key:         "mysqlHostname",
			Value:       "127.0.0.1",
			Description: "hostname used for database (default: 127.0.0.1)",
		},
		{
			Category:    "MySQL",
			Key:         "mysqlPort",
			Value:       "3306",
			Description: "port used for database (default: 3306)",
		},
		{
			Category:    "MySQL",
			Key:         "mysqlDatabase",
			Value:       "eqemu",
			Description: "port used for database (default: eqemu)",
		},
	}
}
