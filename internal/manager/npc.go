package manager

import (
	"context"
	"fmt"

	"github.com/xackery/xegony/pb"
)

// NpcSearch implements a SCRUD endpoint
func (m *Manager) NpcSearch(ctx context.Context, req *pb.NpcSearchRequest) (resp *pb.NpcSearchResponse, err error) {
	if req.Limit < 1 {
		req.Limit = 1
	}
	if req.Limit > 1000 {
		req.Limit = 1000
	}
	resp, err = m.db.NpcSearch(ctx, req)
	return
}

// NpcCreate implements a SCRUD endpoint
func (m *Manager) NpcCreate(ctx context.Context, req *pb.NpcCreateRequest) (resp *pb.NpcCreateResponse, err error) {
	if req.Npc == nil {
		err = fmt.Errorf("no npc passed")
		return
	}
	resp, err = m.db.NpcCreate(ctx, req)
	return
}

// NpcRead implements a SCRUD endpoint
func (m *Manager) NpcRead(ctx context.Context, req *pb.NpcReadRequest) (resp *pb.NpcReadResponse, err error) {
	if req.Id < 1000 {
		err = fmt.Errorf("id must be > 1000")
		return
	}
	if req.Id > 1000000 {
		err = fmt.Errorf("id must be < 1000000")
		return
	}
	resp, err = m.db.NpcRead(ctx, req)
	return
}

// NpcUpdate implements a SCRUD endpoint
func (m *Manager) NpcUpdate(ctx context.Context, req *pb.NpcUpdateRequest) (resp *pb.NpcUpdateResponse, err error) {
	if req.Id < 1000 {
		err = fmt.Errorf("id must be > 1000")
		return
	}
	if req.Id > 1000000 {
		err = fmt.Errorf("id must be < 1000000")
		return
	}
	if req.Npc == nil {
		err = fmt.Errorf("no npc passed")
		return
	}
	resp, err = m.db.NpcUpdate(ctx, req)
	return
}

// NpcDelete implements a SCRUD endpoint
func (m *Manager) NpcDelete(ctx context.Context, req *pb.NpcDeleteRequest) (resp *pb.NpcDeleteResponse, err error) {
	if req.Id < 1000 {
		err = fmt.Errorf("id must be > 1000")
		return
	}
	if req.Id > 1000000 {
		err = fmt.Errorf("id must be < 1000000")
		return
	}
	resp, err = m.db.NpcDelete(ctx, req)
	return
}

// NpcPatch implements a SCRUD endpoint
func (m *Manager) NpcPatch(ctx context.Context, req *pb.NpcPatchRequest) (resp *pb.NpcPatchResponse, err error) {
	if req.Id < 1000 {
		err = fmt.Errorf("id must be > 1000")
		return
	}
	if req.Id > 1000000 {
		err = fmt.Errorf("id must be < 1000000")
		return
	}
	resp, err = m.db.NpcPatch(ctx, req)
	return
}
