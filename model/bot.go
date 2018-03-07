package model

import (
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
	//Parameters are arguments for the bot's job.
	Parameters map[string]string `json:"parameters,omitempty"`
	//Status is the status of the bot. 0: idle, 1: working
	Status int64 `json:"status,omitempty"`
	//LastStart is the time a bot started
	LastStart time.Time `json:"lastStart,omitempty"`
	//LastDuration is the last runtime of this bot
	LastDuration time.Duration `json:"lastDuration,omitempty"`
}

//GetParameterValue is used to quick grab values
func (b *Bot) GetParameterValue(key string) (value string) {
	value, _ = b.Parameters[key]
	return
}
