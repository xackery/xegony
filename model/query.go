package model

import (
	"context"
)

// QueryRequest wraps a query request
type QueryRequest struct {
	Ctx      context.Context
	Method   string
	Req      interface{}
	RespChan chan *QueryResponse
}

// QueryResponse wraps the response to a query
type QueryResponse struct {
	Type  string
	Resp  interface{}
	Error error
}
