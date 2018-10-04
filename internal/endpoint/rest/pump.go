package rest

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/xackery/xegony/model"
)

func (s *Server) pump() {
	var queryReq *model.QueryRequest
	var queryResp *model.QueryResponse

	var templateReq *templateRequest
	var templateResp *templateResponse

	for {
		logger := model.NewLogger().With().Str("type", "rest pump").Logger()
		select {
		case queryReq = <-s.queryChan:
			queryResp = &model.QueryResponse{}
			queryResp.Resp, queryResp.Error = s.onQueryRequest(queryReq)
			if queryReq.RespChan == nil {
				continue
			}
			select {
			case <-time.After(3 * time.Second):
				logger.Error().Msg("timeout on reply")
			case queryReq.RespChan <- queryResp:
			case <-queryReq.Ctx.Done():
			}

		case templateReq = <-s.templateChan:
			templateResp = &templateResponse{}
			templateResp.Resp, templateResp.Error = s.onTemplateRequest(templateReq)
			if templateReq.RespChan == nil {
				continue
			}
			select {
			case <-time.After(3 * time.Second):
				logger.Error().Msg("timeout on reply")
			case templateReq.RespChan <- templateResp:
			case <-templateReq.Ctx.Done():
			}
		}

	}
}

func (s *Server) onQueryRequest(query *model.QueryRequest) (resp proto.Message, err error) {
	if query.Ctx == nil {
		query.Ctx = context.Background()
	}
	if query.Req == nil {
		err = fmt.Errorf("empty request")
		return
	}
	switch query.Req.(type) {
	default:
		err = fmt.Errorf("unknown request type")
		return
	}
}
