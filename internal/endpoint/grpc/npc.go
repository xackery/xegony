package grpc

import (
	"context"

	"github.com/xackery/xegony/pb"
)

// NpcSearch implements SCRUD endpoints
func (s *Server) NpcSearch(ctx context.Context, req *pb.NpcSearchRequest) (resp *pb.NpcSearchResponse, err error) {
	resp, err = s.manager.NpcSearch(ctx, req)
	return
}

// NpcCreate implements SCRUD endpoints
func (s *Server) NpcCreate(ctx context.Context, req *pb.NpcCreateRequest) (resp *pb.NpcCreateResponse, err error) {
	resp, err = s.manager.NpcCreate(ctx, req)
	return
}

// NpcRead implements SCRUD endpoints
func (s *Server) NpcRead(ctx context.Context, req *pb.NpcReadRequest) (resp *pb.NpcReadResponse, err error) {
	resp, err = s.manager.NpcRead(ctx, req)
	return
}

// NpcUpdate implements SCRUD endpoints
func (s *Server) NpcUpdate(ctx context.Context, req *pb.NpcUpdateRequest) (resp *pb.NpcUpdateResponse, err error) {
	resp, err = s.manager.NpcUpdate(ctx, req)
	return
}

// NpcDelete implements SCRUD endpoints
func (s *Server) NpcDelete(ctx context.Context, req *pb.NpcDeleteRequest) (resp *pb.NpcDeleteResponse, err error) {
	resp, err = s.manager.NpcDelete(ctx, req)
	return
}

// NpcPatch implements SCRUD endpoints
func (s *Server) NpcPatch(ctx context.Context, req *pb.NpcPatchRequest) (resp *pb.NpcPatchResponse, err error) {
	resp, err = s.manager.NpcPatch(ctx, req)
	return
}
