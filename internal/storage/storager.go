package storage

import (
	"context"

	"github.com/xackery/xegony/pb"
)

// Storager is an interface for all storage services
type Storager interface {
	Connect(host string, username string, password string, database string) (err error)
	Close() (err error)

	NpcSearch(ctx context.Context, req *pb.NpcSearchRequest) (resp *pb.NpcSearchResponse, err error)
	NpcCreate(ctx context.Context, req *pb.NpcCreateRequest) (resp *pb.NpcCreateResponse, err error)
	NpcRead(ctx context.Context, req *pb.NpcReadRequest) (resp *pb.NpcReadResponse, err error)
	NpcUpdate(ctx context.Context, req *pb.NpcUpdateRequest) (resp *pb.NpcUpdateResponse, err error)
	NpcDelete(ctx context.Context, req *pb.NpcDeleteRequest) (resp *pb.NpcDeleteResponse, err error)
	NpcPatch(ctx context.Context, req *pb.NpcPatchRequest) (resp *pb.NpcPatchResponse, err error)
}
