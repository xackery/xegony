package model

import (
	"net/http"
	"strings"
)

// MakePath returns an append-ready path name
func MakePath(r *http.Request) (path string) {

	path = r.RequestURI
	if !strings.Contains(path, "?") {
		path += "?"
	}
	return
}
