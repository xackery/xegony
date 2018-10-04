package rest

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/pb"
)

func (s *Server) npcMux(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	logger := model.NewLogger()
	path := r.RequestURI

	path = path[4:]
	pathURL, err := url.Parse(path)
	if err != nil {
		err = errors.Wrapf(err, "failed to parse url %s", path)
		return
	}
	path = pathURL.Path
	if len(path) > 0 && string(path[len(path)-1]) != "/" {
		path += "/"
	}
	if len(path) == 0 {
		path = "/"
	}

	logger = logger.With().Str("path", path).Logger()
	switch path {
	case "/search/":
		logger = logger.With().Str("method", "npcSearch").Logger()
		err = s.npcSearch(ctx, w, r)
	case "/":
		logger = logger.With().Str("method", "npcSearch").Logger()
		err = s.npcSearch(ctx, w, r)
	default:
		switch r.Method {
		case "GET":
			logger = logger.With().Str("method", "npcRead").Logger()
			err = s.npcRead(ctx, w, r)
		default:
			logger.Error().Msg("invalid route")
			http.Error(w, "invalid route", 404)
		}
	}

	if err != nil {
		logger.Error().Err(err).Msg("failed")
		http.Error(w, err.Error(), 404)
		return
	}
	logger.Debug().Msgf("%s", path)
}

func (s *Server) npcSearch(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
	type content struct {
		Site *pb.Site
		Page *pb.Page
		Resp *pb.NpcSearchResponse
	}
	c := &content{
		Site: model.NewSite(),
		Page: model.NewPage(),
	}

	req := &pb.NpcSearchRequest{
		Name: getQuery(r, "name", ""),
	}
	c.Resp, err = s.manager.NpcSearch(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "failed to call method")
		return
	}

	t, err := s.getTemplate(&templateRequest{Ctx: ctx, Path: "/npc/search.tpl"})
	if err != nil {
		err = errors.Wrap(err, "failed to get template")
		return
	}
	if t == nil {
		err = fmt.Errorf("no template returned")
		return
	}
	err = t.Execute(w, c)
	if err != nil {
		err = errors.Wrap(err, "failed to execute template")
		return
	}

	return
}

func (s *Server) npcRead(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
	type content struct {
		Site *pb.Site
		Page *pb.Page
		Resp *pb.NpcReadResponse
	}
	c := &content{
		Site: model.NewSite(),
		Page: model.NewPage(),
	}

	req := &pb.NpcReadRequest{
		Id: getIntQuery(r, "id", 0),
	}
	c.Resp, err = s.manager.NpcRead(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "failed to call method")
		return
	}

	t, err := s.getTemplate(&templateRequest{Ctx: ctx, Path: "/npc/read.tpl"})
	if err != nil {
		err = errors.Wrap(err, "failed to get template")
		return
	}
	if t == nil {
		err = fmt.Errorf("no template returned")
		return
	}
	err = t.Execute(w, c)
	if err != nil {
		err = errors.Wrap(err, "failed to execute template")
		return
	}

	return
}
