package model

import (
	"github.com/xackery/xegony/pb"
)

// NewSite returns a default site
func NewSite() (s *pb.Site) {
	s = &pb.Site{
		User: NewUser(),
	}
	return
}
