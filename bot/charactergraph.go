package bot

import (
	"net/http"
	//	"github.com/pkg/errors"
	//"github.com/xackery/xegony/model"
)

func (a *Bot) characterGraphStatus(w http.ResponseWriter, r *http.Request) {
	type Content struct {
		Message string
	}
	content := &Content{
		Message: "Idle",
	}
	writeData(w, r, content, http.StatusOK)
	return
}

//ProcessCharacterGraphCache processes map files and saves them to www/images/maps/ then update zonelevel
func (a *Bot) ProcessCharacterGraphCache() (err error) {

	return
}
