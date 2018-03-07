package file

import (
	"github.com/xackery/xegony/model"
)

func loadConfigDefault() model.Configs {
	return model.Configs{
		{
			Category:    "API",
			Key:         "httpPort",
			Value:       "8080",
			Description: "Port to listen for HTTP routing.",
		},
		{
			Category:    "API",
			Key:         "httpHostname",
			Value:       "",
			Description: "Hostname for HTTP routing. Used for redirects. Supports DNS.",
		},
		{
			Category:    "API",
			Key:         "apiSuffix",
			Value:       "/api",
			Description: "API Suffix is appended to all API endpoints (default: /api/)",
		},
		{
			Category:    "BOT",
			Key:         "mapDir",
			Value:       "map/",
			Description: "Directory map txt files are found. (default: map/)",
		},
		{
			Category:    "WEB",
			Key:         "webSuffix",
			Value:       "/",
			Description: "Web Suffix is appended to all WEB endpoints (default: /)",
		},
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
			Category:    "Google",
			Key:         "googleRedirectURL",
			Value:       "",
			Description: "Google Redirect URL, used for SSO/Google endpoints (by default empty)",
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
		{
			Category:    "Web",
			Key:         "webCacheTemplate",
			Value:       "0",
			Description: "Should templates be cached? 0 for no, 1 for yes. Caching speeds up the load time of web pages, but if you edit templates, they won't show up on web (default: 0)",
		},
		{
			Category:    "Web",
			Key:         "webName",
			Value:       "Xegony",
			Description: "Name of the website. (default: Xegony)",
		},
		{
			Category:    "Web",
			Key:         "webAuthor",
			Value:       "Xackery",
			Description: "Default author of each web page. (default: Xackery)",
		},
		{
			Category:    "Web",
			Key:         "webImage",
			Value:       "/images/logo.png",
			Description: "Default image of each web page. (default: /images/logo.png)",
		},
		{
			Category:    "Web",
			Key:         "webDescription",
			Value:       "Xegony is a project",
			Description: "Default description of each web page. (default: Xegony is a project)",
		},
	}
}
