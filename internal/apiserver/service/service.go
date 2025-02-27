package service

import (
	"github.com/ividernvi/iviuser/internal/apiserver/cache"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
)

type Service interface {
	Users() UserSrv
}

type service struct {
	store store.Store
	cache cache.Cache
}

func NewService(store store.Store) Service {
	return &service{
		store: store,
		cache: cache.CacheFactory(),
	}
}

func (s *service) Users() UserSrv {
	return newUserSrv(s)
}
