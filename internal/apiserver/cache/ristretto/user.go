package ristretto

import (
	"context"
	"encoding/json"

	"github.com/dgraph-io/ristretto/v2"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/sirupsen/logrus"
)

type userCache struct {
	cache *ristretto.Cache[string, string]
	keys  map[string]bool
}

func newUserCache(d *datacache) *userCache {
	return &userCache{
		cache: d.cache,
		keys:  make(map[string]bool),
	}
}

func (u *userCache) Set(ctx context.Context, user *v1.User, opts *v1.CreateOptions) error {
	userInfo := user.String()
	u.keys[user.UserName] = true
	u.cache.Set(user.UserName, userInfo, int64(len(userInfo)))
	logrus.Debugf("User cache set: %s:\n\t %s", user.UserName, userInfo)
	return nil
}

func (u *userCache) Get(ctx context.Context, key string, opts *v1.GetOptions) (*v1.User, error) {
	value, ok := u.cache.Get(key)
	if !ok {
		return nil, nil
	}
	var user v1.User
	err := json.Unmarshal([]byte(value), &user)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("User cache get: %s:\n\t %s", key, user.String())
	return &user, nil
}

func (u *userCache) Del(ctx context.Context, key string, opts *v1.DeleteOptions) error {
	u.cache.Del(key)
	logrus.Debugf("User cache del: %s", key)
	return nil
}
