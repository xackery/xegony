package endpoint

import (
	"context"

	"github.com/xackery/xegony/pb"
)

// Endpointer wraps endpoints
type Endpointer interface {
	Listen(ctx context.Context, host string) (err error)
	SetEndpoint(ctx context.Context, endpoint *pb.Endpoint) (err error)
	GetEndpoint(ctx context.Context) (endpoint *pb.Endpoint, err error)
	Close(ctx context.Context) (err error)
}
