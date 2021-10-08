package action

import (
	"context"

	"github.com/jinzhu/copier"
	pb "github.com/moh-fajri/learn-garphql-user/entity/user"
	"github.com/moh-fajri/learn-garphql-user/repo/userdb"
)

func (h *Handler) List(ctx context.Context, req *pb.NoRequest) (res *pb.UserResponse, err error) {
	var users []userdb.User
	users, err = h.userRepo.List(ctx)
	if err != nil {
		return
	}
	var response []*pb.User
	copier.Copy(&response, users)
	res = &pb.UserResponse{Users: response}
	return
}
