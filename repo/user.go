package repo

import (
	"context"
	"fmt"

	pb "github.com/moh-fajri/learn-garphql-user/entity/user"

	"github.com/moh-fajri/learn-garphql-user/repo/userdb"
)

type User struct {
	db userdb.IUserDb
}

func NewUser() *User {
	return &User{
		db: userdb.NewClient(),
	}
}

func (u *User) List(ctx context.Context) (res []userdb.User, err error) {
	err = u.db.List(ctx, &res, userdb.QueryUser)
	return
}

func (u *User) Get(ctx context.Context, id string) (res userdb.User, err error) {
	err = u.db.Get(ctx, &res, userdb.QueryGetUser, id)
	return
}

func (u *User) Create(ctx context.Context, req *pb.CreateRequest) (res []userdb.User, err error) {
	query := fmt.Sprintf(userdb.QueryCreateUser, req.Name, req.Email)
	err = u.db.List(ctx, &res, query)
	return
}
