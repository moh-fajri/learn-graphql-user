package repo

import (
	"context"

	pb "github.com/moh-fajri/learn-garphql-user/entity/user"

	"github.com/moh-fajri/learn-garphql-user/repo/userdb"
)

type IUser interface {
	List(context.Context) ([]userdb.User, error)
	Get(context.Context, string) (userdb.User, error)
	Create(context.Context, *pb.CreateRequest) ([]userdb.User, error)
}
