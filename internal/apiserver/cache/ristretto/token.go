package ristretto

import (
	"context"

	"github.com/dgraph-io/ristretto/v2"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
	"github.com/ividernvi/iviuser/pkg/util/jwtutil"
)

type tokenCache struct {
	cache *ristretto.Cache[string, string]
	keys  map[string]bool
}

func newTokenCache(d *datacache) *tokenCache {
	return &tokenCache{
		cache: d.cache,
		keys:  make(map[string]bool),
	}
}

func (u *tokenCache) Set(ctx context.Context, token string, opts *v1.CreateOptions) error {
	u.cache.SetWithTTL(token, "", 1, jwtutil.ExpireDuration)
	return nil
}

func (u *tokenCache) Get(ctx context.Context, key string, opts *v1.GetOptions) error {
	val, _ := u.cache.Get(key)
	if val == key {
		return core.ErrTokenInvalid
	}
	return nil
}
