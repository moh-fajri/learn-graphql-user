package action

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/moh-fajri/learn-garphql-user/repo/userdb"

	pb "github.com/moh-fajri/learn-garphql-user/entity/user"
)

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest) (res *pb.UserResponse, err error) {
	var users []userdb.User

	users, err = h.userRepo.Create(ctx, req)
	if err != nil {
		return
	}

	var response []*pb.User
	copier.Copy(&response, users)

	res = &pb.UserResponse{Users: response}
	return
}
