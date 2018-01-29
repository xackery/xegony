package cases

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/xackery/xegony/model"
)

func preparePage(page *model.Page, user *model.User) (err error) {
	if page == nil {
		page = &model.Page{
			Limit:  10,
			Offset: 0,
		}
	}
	if page.Limit < 1 {
		page.Limit = 1
	}
	if page.Limit > 100 {
		page.Limit = 100
	}

	if page.IsDescending > 0 {
		page.IsDescending = 1
	}

	page.OrderBy = strings.ToLower(page.OrderBy)
	var r *regexp.Regexp
	if len(page.OrderBy) > 0 {
		r, err = regexp.Compile("[a-z0-9_]+")
		if err != nil {
			err = fmt.Errorf("Invalid orderby: %s", err.Error())
			return
		}

		if !r.MatchString(page.OrderBy) {
			err = &model.ErrValidation{
				Message: "orderBy is invalid",
				Reasons: map[string]string{
					"orderBy": "invalid syntax",
				},
			}
			return
		}
	}

	return
}

func sanitizePage(page *model.Page, user *model.User) (err error) {
	if page == nil {
		page = &model.Page{
			Limit:  10,
			Offset: 0,
		}
	}
	return
}
