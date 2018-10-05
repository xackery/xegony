package mariadb

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/pb"
)

// NpcSearch implements a SCRUD endpoint
func (s *Storage) NpcSearch(ctx context.Context, req *pb.NpcSearchRequest) (resp *pb.NpcSearchResponse, err error) {
	resp = &pb.NpcSearchResponse{}
	where := ""
	whereMap := make(map[string]interface{})
	if len(req.Name) > 0 {
		where += "Name LIKE :Name AND"
		whereMap["Name"] = "%" + req.Name + "%"
	}

	if len(whereMap) < 1 {
		where += "id > 0 AND"
		return
	}

	where = where[0 : len(where)-4] //remove ' AND'

	query := fmt.Sprintf("SELECT count(id) FROM npc_types WHERE %s", where)
	rows, err := s.db.NamedQueryContext(ctx, query, whereMap)
	if err != nil {
		err = errors.Wrap(err, "failed to get total")
		return
	}
	for rows.Next() {
		err = rows.Scan(&resp.Total)
		if err != nil {
			err = errors.Wrap(err, "failed to scan total")
			return
		}
	}
	npc := &model.Npc{}
	query = fmt.Sprintf("SELECT %s FROM npc_types WHERE %s ORDER BY %s LIMIT %d OFFSET %d", npc.Fields(), where, req.OrderBy, req.Limit, req.Offset)
	rows, err = s.db.NamedQueryContext(ctx, query, whereMap)
	if err != nil {
		err = errors.Wrap(err, "failed to select")
		return
	}
	for rows.Next() {
		err = rows.StructScan(npc)
		if err != nil {
			err = errors.Wrap(err, "failed to scan npc")
			return
		}
		resp.Npcs = append(resp.Npcs, npc.ToProto())
	}
	return
}

// NpcCreate implements a SCRUD endpoint
func (s *Storage) NpcCreate(ctx context.Context, req *pb.NpcCreateRequest) (resp *pb.NpcCreateResponse, err error) {
	err = fmt.Errorf("not yet supported")
	return
}

// NpcRead implements a SCRUD endpoint
func (s *Storage) NpcRead(ctx context.Context, req *pb.NpcReadRequest) (resp *pb.NpcReadResponse, err error) {
	resp = &pb.NpcReadResponse{}
	npc := &model.Npc{}

	query := fmt.Sprintf("SELECT %s FROM npc_types WHERE id = ? LIMIT 1", npc.Fields())
	err = s.db.GetContext(ctx, npc, query, req.Id)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare statement")
		return
	}

	resp.Npc = npc.ToProto()
	return
}

// NpcUpdate implements a SCRUD endpoint
func (s *Storage) NpcUpdate(ctx context.Context, req *pb.NpcUpdateRequest) (resp *pb.NpcUpdateResponse, err error) {
	err = fmt.Errorf("not yet supported")
	return
}

// NpcDelete implements a SCRUD endpoint
func (s *Storage) NpcDelete(ctx context.Context, req *pb.NpcDeleteRequest) (resp *pb.NpcDeleteResponse, err error) {
	resp = &pb.NpcDeleteResponse{}
	where := "id = :id"
	whereMap := make(map[string]interface{})
	whereMap["id"] = req.Id

	query := fmt.Sprintf("DELETE FROM npc_types WHERE %s LIMIT 1", where)
	nstmt, err := s.db.PrepareNamedContext(ctx, query)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare statement")
		return
	}

	_, err = nstmt.ExecContext(ctx, whereMap)

	if err != nil {
		err = errors.Wrap(err, "failed to delete")
		return
	}

	return
}

// NpcPatch implements a SCRUD endpoint
func (s *Storage) NpcPatch(ctx context.Context, req *pb.NpcPatchRequest) (resp *pb.NpcPatchResponse, err error) {
	err = fmt.Errorf("not yet supported")
	return
}
