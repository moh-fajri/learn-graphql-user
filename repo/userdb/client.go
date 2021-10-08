package userdb

import (
	"context"
)

// Client is a struct to mongo
type Client struct{}

// NewClient for initiate function in redis
func NewClient() *Client {
	return &Client{}
}

func (c *Client) List(ctx context.Context, data interface{}, query string) (err error) {
	err = UserDB.SelectContext(ctx, data, UserDB.Rebind(query))
	return
}

func (c *Client) Get(ctx context.Context, data interface{}, query string, id string) (err error) {
	err = UserDB.GetContext(ctx, data, query, id)
	return
}
