package store

import (
	"context"

	v1 "github.com/ividernvi/iviuser/model/v1"
)

type SubscribeStore interface {
	Create(ctx context.Context, subscribe *v1.Subscribe, opts *v1.CreateOptions) error
	Delete(ctx context.Context, name string, opts *v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts *v1.GetOptions) (*v1.Subscribe, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.SubscribeList, error)
	Update(ctx context.Context, subscribe *v1.Subscribe, opts *v1.UpdateOptions) error
}
