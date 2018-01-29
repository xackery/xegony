package model

import (
//"fmt"
//"html/template"
)

// Page represents pagination
// swagger:model
type Page struct {
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// Total number of results found
	// example: 100
	// in: query
	Total int64 `json:"total"`
	// OrderBy is which field to order a page by
	// example: id
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

/*
//PageList allows a site to create pagination on bottom
func (c *Page) PageList() template.HTML {
	page := `<div class="btn-group pull-right">`
	curPage := c.PageNumber

	var curElement int64

	if c.PageNumber > 0 {
		page += fmt.Sprintf("\n"+`<button type="button" class="btn btn-default"><a href="/%s?pageNumber=%d"><i class="fa fa-chevron-left"></i></a></button>`, c.Scope, c.PageNumber-1)
	}

	curElement = (c.PageNumber - 6) * c.PageSize
	curPage -= 6
	numCount := 0

	for curElement <= c.Total {
		if curPage < 0 {
			curPage++
			curElement += c.PageSize
			continue
		}
		curPage++
		if curPage == c.PageNumber {
			page += fmt.Sprintf("\n"+` <button class="btn btn-default active"><a href="/%s/?pageNumber=%d">%d</a></button>`, c.Scope, curPage, curPage)
		} else {
			page += fmt.Sprintf("\n"+` <button class="btn btn-default"><a href="/%s/?pageNumber=%d">%d</a></button>`, c.Scope, curPage, curPage)
		}
		curElement += c.PageSize
		numCount++
		if numCount >= 10 {
			break
		}
	}
	if c.PageNumber*c.PageSize < c.Total {
		page += fmt.Sprintf("\n"+`<button type="button" class="btn btn-default"><a href="/%s?pageNumber=%d"><i class="fa fa-chevron-right"></a></i></button>`, c.Scope, c.PageNumber+1)
	}
	page += "\n</div>"
	return template.HTML(page)
}
*/
