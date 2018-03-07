package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
)

func prepareBot(bot *model.Bot, user *model.User) (err error) {
	if bot == nil {
		err = fmt.Errorf("empty bot")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	if bot.Parameters == nil {
		bot.Parameters = make(map[string]string)
	}
	return
}
