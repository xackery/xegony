package grpc

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xackery/xegony/internal/manager"
	"github.com/xackery/xegony/pb"
	"google.golang.org/grpc"
)

// Server wraps the grpc server
type Server struct {
	svr        *grpc.Server
	lis        net.Listener
	manager    *manager.Manager
	ctx        context.Context
	cancel     context.CancelFunc
	serverLock sync.RWMutex
	endpoint   *pb.Endpoint
}

// New creates a new grpc server
func New(m *manager.Manager) (s *Server, err error) {
	s = &Server{
		manager: m,
	}
	return
}

// Listen will listen on a specified host
func (s *Server) Listen(host string) (err error) {
	if s.lis != nil {
		s.Close()
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
func (s *Server) SetEndpoint(endpoint *pb.Endpoint) (err error) {
	errChan := make(chan error)
	go func() {
		s.serverLock.Lock()
		defer s.serverLock.Unlock()
		_, ok := (<-errChan)
		if !ok {
			return
		}
		s.endpoint = endpoint
		errChan <- nil
	}()
	select {
	case err = <-errChan:
	case <-time.After(1 * time.Second):
		close(errChan)
		err = fmt.Errorf("failed to set endpoint")
	}
	return
}

// Close exits the endpoint
func (s *Server) Close() (err error) {
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
func (s *Server) GetEndpoint() (endpoint *pb.Endpoint, err error) {
	endpoint = s.endpoint
	return
}
