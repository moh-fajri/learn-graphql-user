package userdb

import "context"

type IUserDb interface {
	List(context.Context, interface{}, string) error
	Get(context.Context, interface{}, string, string) error
}
