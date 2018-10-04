package model

import (
	"github.com/xackery/xegony/pb"
)

// NewUser returns a default user
func NewUser() (u *pb.User) {
	u = &pb.User{}
	return
}
