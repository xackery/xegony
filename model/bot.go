package model

import (
	"context"
	"time"
)

// Bots is an array of Bot
// swagger:model
type Bots []*Bot

// Bot represents workers
// swagger:model
type Bot struct {
	//Index of bot. Typically is auto incremented
	ID int64 `json:"ID,omitempty"`
	//Context is used for cancalling a bot. You should use the status editing instead.
	Context context.Context `json:"context,omitempty"`
	//Parameters are arguments for the bot's job.
	Parameters []string `json:"parameters,omitempty"`
	//Status is the status of the bot. 0: idle, 1: working
	Status int64 `json:"status,omitempty"`
	//LastStart is the time a bot started
	LastStart time.Time `json:"lastStart,omitempty"`
	//LastDuration is the last runtime of this bot
	LastDuration time.Duration `json:"lastDuration,omitempty"`
}
