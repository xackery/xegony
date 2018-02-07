package web

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

type site struct {
	Title       string //Title of site
	Name        string
	Page        string
	Section     string
	Description string //Description for oprop
	Image       string
	//Author is who was the original creator of this page
	Author string
	User   *model.User
}

func newSite(r *http.Request, user *model.User) (siteData site) {
	re := regexp.MustCompile("[^a-z]+")

	rootTitle := strings.ToLower(r.URL.Path)
	if rootTitle[0] == '/' {
		rootTitle = rootTitle[1:]
	}
	if len(rootTitle) == 0 {
		rootTitle = "index"
	}
	title := rootTitle
	titleSplit := strings.Split(title, "/")
	title = ""
	for _, titleChunk := range titleSplit {
		title += " " + titleChunk
	}
	title = strings.Title(title)

	rootTitle = re.ReplaceAllString(rootTitle, "")
	siteData = site{
		Name:        cases.GetConfigValue("webName"),
		Title:       strings.ToTitle(rootTitle),
		Page:        rootTitle,
		Section:     rootTitle,
		Author:      cases.GetConfigValue("webAuthor"),
		Image:       cases.GetConfigValue("webImage"),
		Description: cases.GetConfigValue("webDescription"),
		User:        user,
	}

	return
}
