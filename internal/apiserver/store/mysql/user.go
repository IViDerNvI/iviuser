package mysql

import (
	"context"

	v1 "github.com/ividernvi/iviuser/model/v1"
	"gorm.io/gorm"
)

type userStore struct {
	db *gorm.DB
}

func newUserStore(db *gorm.DB) *userStore {
	return &userStore{db: db}
}

func (s *userStore) Create(ctx context.Context, user *v1.User, opts *v1.CreateOptions) error {
	return s.db.Create(user).Error
}

func (s *userStore) Get(ctx context.Context, username string, opts *v1.GetOptions) (*v1.User, error) {
	var user v1.User
	err := s.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (s *userStore) List(ctx context.Context, opts *v1.ListOptions) (*v1.UserList, error) {
	var users []v1.User
	err := s.db.Find(&users).Error
	return &v1.UserList{Items: users}, err
}

func (s *userStore) Update(ctx context.Context, user *v1.User, opts *v1.UpdateOptions) error {
	return s.db.Save(user).Error
}

func (s *userStore) Delete(ctx context.Context, username string, opts *v1.DeleteOptions) error {
	return s.db.Delete(&v1.User{}, username).Error
}
