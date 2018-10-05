package client

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/pkg/profile"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/xackery/xegony/internal/endpoint"
	"github.com/xackery/xegony/internal/endpoint/grpc"
	"github.com/xackery/xegony/internal/endpoint/rest"
	"github.com/xackery/xegony/internal/manager"
)

// Client represents the Xegony instance
type Client struct {
	manager   *manager.Manager
	endpoints []endpoint.Endpointer
}

// New creates a new client
func New(ctx context.Context) (c *Client, err error) {
	p := profile.Start()
	c = &Client{}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	c.manager, err = manager.New()
	if err != nil {
		err = errors.Wrap(err, "failed to start manager")
		return
	}

	var end endpoint.Endpointer
	end, err = grpc.New(c.manager)
	if err != nil {
		err = errors.Wrap(err, "failed to start grpc")
		return
	}
	c.endpoints = append(c.endpoints, end)
	end.Listen(ctx, ":8081")

	end, err = rest.New(":8081", c.manager)
	if err != nil {
		err = errors.Wrap(err, "failed to start rest")
		return
	}
	err = end.Listen(ctx, ":8082")
	if err != nil {
		err = errors.Wrap(err, "could not listen on rest")
		return
	}
	c.endpoints = append(c.endpoints, end)

	log.Debug().Msg("logger started, http://localhost:8082/npc")
	p.Stop()
	return
}

// Close will close the xegony instance
func (c *Client) Close(ctx context.Context) (errors []error) {
	var err error
	for _, end := range c.endpoints {
		err = end.Close(ctx)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return
}
