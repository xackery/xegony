package rest

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xackery/xegony/box"
	"github.com/xackery/xegony/internal/manager"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/pb"
	"google.golang.org/grpc"
)

type ctxKey struct{}

//https://github.com/grpc-ecosystem/grpc-gateway/blob/master/examples/gateway/main.go

// Server wraps the rest endpoints using endpointer
type Server struct {
	grpcHost     string
	host         string
	cancelGrpc   context.CancelFunc
	svr          *http.Server
	manager      *manager.Manager
	queryChan    chan *model.QueryRequest
	templateChan chan *templateRequest
	templates    map[string]*template.Template
}

// New creates a new Rest instance
func New(grpcHost string, manager *manager.Manager) (s *Server, err error) {
	s = &Server{
		grpcHost:     grpcHost,
		manager:      manager,
		queryChan:    make(chan *model.QueryRequest),
		templateChan: make(chan *templateRequest),
		templates:    make(map[string]*template.Template),
	}
	go s.pump()
	return
}

// Listen lisetns for new connection details
func (s *Server) Listen(host string) (err error) {
	if s.svr != nil {
		s.Close()
		s.svr = nil
	}
	s.host = host
	ctx := context.Background()
	ctx, s.cancelGrpc = context.WithCancel(ctx)

	runtime.OtherErrorHandler = otherErrorHandler

	r := mux.NewRouter()

	//handler := http.NewServeMux()

	muxGRPC := runtime.NewServeMux(runtime.WithMarshalerOption("application/json", &runtime.JSONPb{
		//OrigName:    true,
		EnumsAsInts: true,
	}))
	r.PathPrefix("/v1/").Handler(contextWrap(muxGRPC))
	r.PathPrefix("/npc").HandlerFunc(s.npcMux).Name("npc")
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = pb.RegisterXegonyHandlerFromEndpoint(ctx, muxGRPC, s.grpcHost, opts)
	if err != nil {
		err = errors.Wrap(err, "failed to register endpoint")
		return err
	}
	s.svr = &http.Server{
		Addr:    s.host,
		Handler: r,
	}
	go func() {

		err = s.svr.ListenAndServe()
		log.Error().Err(err).Msg("rest listen crash")
	}()
	return
}

// Close closes the rest server
func (s *Server) Close() (err error) {
	if s.cancelGrpc != nil {
		s.cancelGrpc()
	}
	if s.svr != nil {
		err = s.svr.Close()
		if err != nil {
			return
		}
	}
	return
}

func contextWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger := model.NewLogger()
		logger.Debug().Msg("test")
		//req = req.WithContext(context.WithValue(req.Context(), ctxKey{}, "test"))
		//log.Debug().Msgf("in wrapper, context value is", req.Context().Value(ctxKey{}))
		h.ServeHTTP(w, req)
	})
}

func otherErrorHandler(w http.ResponseWriter, r *http.Request, _ string, _ int) {
	var err error
	path := "web/www/" + r.URL.Path[1:]

	if _, err = os.Stat(path); err == nil {
		http.ServeFile(w, r, path)
		return
	}

	var bData []byte
	if bData, err = box.ReadFile(path); err == nil {
		reader := bytes.NewReader(bData)
		http.ServeContent(w, r, path, time.Now(), reader)
		return
	}

	http.Error(w, fmt.Sprintf("404 - Not Found: %s", r.URL), http.StatusNotFound)
	return
}

// SetEndpoint sets an endpoint
func (s *Server) SetEndpoint(endpoint *pb.Endpoint) (err error) {
	return
}

// GetEndpoint returns an endpoint
func (s *Server) GetEndpoint() (endpoint *pb.Endpoint, err error) {
	return
}

// getQuery parses query parameters based on key and returns as a string
func getQuery(r *http.Request, key string, fallback string) string {
	vals := r.URL.Query()
	keyTypes, ok := vals[key]
	if ok {
		if len(keyTypes) > 0 {
			return keyTypes[0]
		}
	}
	return fallback
}

// getIntQuery parses query parameters based on key and returns as an int64
func getIntQuery(r *http.Request, key string, fallback int64) int64 {
	var val int64
	vals := r.URL.Query()
	keyTypes, ok := vals[key]
	if ok {
		if len(keyTypes) > 0 {
			val, _ = strconv.ParseInt(keyTypes[0], 10, 64)
			return val
		}
	}
	return fallback
}
