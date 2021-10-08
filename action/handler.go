package action

import "github.com/moh-fajri/learn-garphql-user/repo"

type Handler struct {
	userRepo repo.IUser
}

func NewHandler() *Handler {
	return &Handler{
		userRepo: repo.NewUser(),
	}
}
