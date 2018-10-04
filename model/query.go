package model

import (
	"context"

	"github.com/golang/protobuf/proto"
)

// QueryRequest wraps a query request
type QueryRequest struct {
	Ctx      context.Context
	Req      proto.Message
	RespChan chan *QueryResponse
}

// QueryResponse wraps the response to a query
type QueryResponse struct {
	Resp  proto.Message
	Error error
}
