package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xackery/xegony/internal/manager"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/pb"
	"google.golang.org/grpc"
)

// Server wraps the grpc server
type Server struct {
	svr       *grpc.Server
	lis       net.Listener
	manager   *manager.Manager
	ctx       context.Context
	cancel    context.CancelFunc
	endpoint  *pb.Endpoint
	queryChan chan *model.QueryRequest
}

// New creates a new grpc server
func New(m *manager.Manager) (s *Server, err error) {
	s = &Server{
		manager:   m,
		queryChan: make(chan *model.QueryRequest),
	}
	return
}

// Listen will listen on a specified host
func (s *Server) Listen(ctx context.Context, host string) (err error) {
	_, err = s.runQuery(ctx, "Listen", host)
	return
}

func (s *Server) onListen(ctx context.Context, host string) (err error) {
	if s.lis != nil {
		err = s.onClose(ctx)
		if err != nil {
			return
		}
	}
	s.lis, err = net.Listen("tcp", host)
	if err != nil {
		err = errors.Wrapf(err, "failed to listen to %s", host)
		return
	}

	s.svr = grpc.NewServer()
	pb.RegisterXegonyServer(s.svr, s)

	go func() {
		err = s.svr.Serve(s.lis)
		log.Error().Err(err).Msg("grpc server stopped")
	}()
	return
}

// SetEndpoint sets an endpoint with provided details
func (s *Server) SetEndpoint(ctx context.Context, endpoint *pb.Endpoint) (err error) {
	_, err = s.runQuery(ctx, "SetEndpoint", endpoint)
	return
}
func (s *Server) onSetEndpoint(ctx context.Context, endpoint *pb.Endpoint) (err error) {
	s.endpoint = endpoint
	return
}

// Close exits the endpoint
func (s *Server) Close(ctx context.Context) (err error) {
	_, err = s.runQuery(ctx, "Close", nil)
	return
}

func (s *Server) onClose(ctx context.Context) (err error) {
	if s.cancel != nil {
		s.cancel()
	}
	if s.lis != nil {
		err = s.lis.Close()
	}
	if s.svr != nil {
		s.svr.Stop()
	}
	return
}

// GetEndpoint returns an endpoint
func (s *Server) GetEndpoint(ctx context.Context) (endpoint *pb.Endpoint, err error) {
	resp, err := s.runQuery(ctx, "GetEndpoint", nil)
	endpoint, ok := resp.(*pb.Endpoint)
	if !ok {
		err = fmt.Errorf("failed to parse response")
	}
	return
}

func (s *Server) onGetEndpoint(ctx context.Context) (endpoint *pb.Endpoint, err error) {
	endpoint = s.endpoint
	return
}
