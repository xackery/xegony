package endpoint

import (
	"github.com/xackery/xegony/pb"
)

// Endpointer wraps endpoints
type Endpointer interface {
	Listen(host string) (err error)
	SetEndpoint(endpoint *pb.Endpoint) (err error)
	GetEndpoint() (endpoint *pb.Endpoint, err error)
	Close() (err error)
}
