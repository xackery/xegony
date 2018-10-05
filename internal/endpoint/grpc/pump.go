package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/pb"
)

func (s *Server) runQuery(ctx context.Context, method string, req interface{}) (resp interface{}, err error) {
	queryReq := &model.QueryRequest{
		Ctx:      ctx,
		Method:   method,
		Req:      req,
		RespChan: make(chan *model.QueryResponse),
	}
	if queryReq.Ctx == nil {
		queryReq.Ctx = context.Background()
	}

	select {
	case <-ctx.Done():
		err = fmt.Errorf("context cancelled")
	case <-time.After(3 * time.Second):
		err = fmt.Errorf("timeout waiting for response")
	case s.queryChan <- queryReq:
		select {
		case <-ctx.Done():
			err = fmt.Errorf("context cancelled")
		case <-time.After(3 * time.Second):
			err = fmt.Errorf("timeout waiting for response")
		case respMsg := <-queryReq.RespChan:
			resp = respMsg.Resp
			err = respMsg.Error
		}
	}
	return
}

func (s *Server) pump() {
	var queryReq *model.QueryRequest
	var ctx context.Context
	var err error
	var resp interface{}
	for {
		logger := model.NewLogger().With().Str("type", "rest pump").Logger()
		err = nil

		select {
		case queryReq = <-s.queryChan:
			ctx = queryReq.Ctx
			switch queryReq.Method {
			case "SetEndpoint":
				req, ok := queryReq.Req.(*pb.Endpoint)
				if !ok {
					err = fmt.Errorf("invalid request type")
				} else {
					err = s.onSetEndpoint(ctx, req)
				}
			case "GetEndpoint":
				resp, err = s.onGetEndpoint(ctx)
			case "Listen":
				req, ok := queryReq.Req.(string)
				if !ok {
					err = fmt.Errorf("invalid request type")
				} else {
					err = s.onListen(ctx, req)
				}
			case "Close":
				err = s.onClose(ctx)
			}
			if queryReq.RespChan == nil {
				continue
			}
			select {
			case <-time.After(3 * time.Second):
				logger.Error().Msg("timeout on reply")
			case queryReq.RespChan <- &model.QueryResponse{Resp: resp, Error: err}:
			case <-queryReq.Ctx.Done():
			}
		}
	}
}
