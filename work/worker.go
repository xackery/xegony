package work

import (
	"github.com/xackery/xegony/model"
)

//Worker wraps the bot implementation
type Worker interface {
	CreateBot(bot *model.Bot) (err error)
	GetBot(bot *model.Bot) (err error)
	EditBot(bot *model.Bot) (err error)
	ListBot(page *model.Page) (bots []*model.Bot, err error)
}
