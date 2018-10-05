// Copyright 2015 Daniel Theophanes.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// simple does nothing except block while running the service.
package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/kardianos/service"
	"github.com/rs/zerolog/log"
	"github.com/xackery/xegony/client"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	ctx, cancel := context.WithCancel(context.Background())
	c, err := client.New(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start new client")
	}
	closeChan := make(chan os.Signal, 1)
	signal.Notify(closeChan, os.Interrupt)
	go func() {
		for sig := range closeChan {
			log.Info().Msgf("got close signal %s", sig.String())
			errors := c.Close(ctx)
			if len(errors) > 0 {
				for _, err = range errors {
					log.Error().Err(err).Msg("error closing client")
				}
				cancel()
				os.Exit(1)
			}
			cancel()
			os.Exit(0)
		}
	}()

}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "Xegony",
		DisplayName: "Xegony",
		Description: "Xegony is a EQEMU editor",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start service")
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start logger")
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
