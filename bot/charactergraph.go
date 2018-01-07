package bot

import (
	"fmt"
	"net/http"
	"time"
	//	"github.com/pkg/errors"
	//"github.com/xackery/xegony/model"
)

func (a *Bot) characterGraphStatus(w http.ResponseWriter, r *http.Request) {
	var err error
	type Content struct {
		Message     string
		Status      string
		Runtime     string
		LastStarted time.Time
	}

	var bot *Status
	if bot, err = a.getStatus("charactergraph"); err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	content := &Content{
		Message:     fmt.Sprintf("Bot is %s, last started %s", bot.State, bot.StartTime),
		Status:      bot.State,
		Runtime:     fmt.Sprintf("%.2f minutes", bot.getRuntime().Minutes()),
		LastStarted: bot.StartTime,
	}

	a.writeData(w, r, content, http.StatusOK)
	return
}

func (a *Bot) characterGraphCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	type Content struct {
		Message string
	}

	content := &Content{
		Message: "Starting bot to process charactergraph",
	}

	if err = a.startBot("charactergraph"); err != nil {
		a.writeError(w, r, err, http.StatusForbidden)
		return
	}

	go a.ProcessCharacterGraphCache()

	a.writeData(w, r, content, http.StatusOK)
	return
}

//ProcessCharacterGraphCache processes map files and saves them to www/images/maps/ then update zonelevel
func (a *Bot) ProcessCharacterGraphCache() (err error) {

	return
}
