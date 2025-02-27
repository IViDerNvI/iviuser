package cache

import (
	"context"

	v1 "github.com/ividernvi/iviuser/model/v1"
)

type UserCache interface {
	Set(ctx context.Context, user *v1.User, opts *v1.CreateOptions) error
	Get(ctx context.Context, username string, opts *v1.GetOptions) (*v1.User, error)
	Del(ctx context.Context, username string, opts *v1.DeleteOptions) error
}
