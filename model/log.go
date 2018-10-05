package model

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// NewLogger returns a default logger
func NewLogger() (logger zerolog.Logger) {
	logger = log.With().Logger()
	return
}
