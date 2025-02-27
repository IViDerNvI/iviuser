package service

import (
	"context"

	"github.com/ividernvi/iviuser/internal/apiserver/cache"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
	v1 "github.com/ividernvi/iviuser/model/v1"
)

type UserSrv interface {
	Create(ctx context.Context, user *v1.User, opts *v1.CreateOptions) error
	Get(ctx context.Context, username string, opts *v1.GetOptions) (*v1.User, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.UserList, error)
	Update(ctx context.Context, user *v1.User, opts *v1.UpdateOptions) error
	Delete(ctx context.Context, username string, opts *v1.DeleteOptions) error
}

type userSrv struct {
	store store.Store
	cache cache.Cache
}

func newUserSrv(s *service) *userSrv {
	return &userSrv{
		store: s.store,
		cache: s.cache,
	}
}

func (s *userSrv) Create(ctx context.Context, user *v1.User, opts *v1.CreateOptions) error {
	return s.store.Users().Create(ctx, user, opts)
}

func (s *userSrv) Get(ctx context.Context, username string, opts *v1.GetOptions) (*v1.User, error) {
	user, err := s.cache.Users().Get(ctx, username, opts)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}
	user, err = s.store.Users().Get(ctx, username, opts)
	if err != nil {
		return nil, err
	}

	if user != nil {
		go func() {
			err = s.cache.Users().Set(ctx, user, nil)
			if err != nil {
				return
			}
		}()
	}
	return user, nil
}

func (s *userSrv) List(ctx context.Context, opts *v1.ListOptions) (*v1.UserList, error) {
	return s.store.Users().List(ctx, opts)
}

func (s *userSrv) Update(ctx context.Context, user *v1.User, opts *v1.UpdateOptions) error {
	return s.store.Users().Update(ctx, user, opts)
}

func (s *userSrv) Delete(ctx context.Context, username string, opts *v1.DeleteOptions) error {
	err := s.store.Users().Delete(ctx, username, opts)
	go func() {
		err = s.cache.Users().Del(ctx, username, nil)
		if err != nil {
			return
		}
	}()
	return err
}
