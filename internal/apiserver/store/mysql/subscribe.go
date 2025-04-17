package mysql

import (
	"context"

	v1 "github.com/ividernvi/iviuser/model/v1"
	"gorm.io/gorm"
)

type subscribeStore struct {
	db *gorm.DB
}

func newSubscribeStore(db *gorm.DB) *subscribeStore {
	return &subscribeStore{db: db}
}

func (s *subscribeStore) Create(ctx context.Context, subscribe *v1.Subscribe, opts *v1.CreateOptions) error {
	return s.db.Create(subscribe).Error
}

func (s *subscribeStore) Get(ctx context.Context, name string, opts *v1.GetOptions) (*v1.Subscribe, error) {
	subscribe := new(v1.Subscribe)
	err := s.db.Where("name = ?", name).First(subscribe).Error
	return subscribe, err
}

func (s *subscribeStore) List(ctx context.Context, opts *v1.ListOptions) (*v1.SubscribeList, error) {
	var subscribes []v1.Subscribe
	var total int64

	query := opts.ApplyListOptions(s.db)
	err := query.Model(&v1.Subscribe{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = query.Offset(opts.Offset).Limit(opts.Limit).Find(&subscribes).Error
	if err != nil {
		return nil, err
	}

	return &v1.SubscribeList{
		ListMeta: v1.ListMeta{
			TotalItems: total,
		},
		Items: subscribes,
	}, nil
}

func (s *subscribeStore) Update(ctx context.Context, subscribe *v1.Subscribe, opts *v1.UpdateOptions) error {
	return s.db.Save(subscribe).Error
}

func (s *subscribeStore) Delete(ctx context.Context, name string, opts *v1.DeleteOptions) error {
	return s.db.Where("name = ?", name).Delete(&v1.Subscribe{}).Error
}
