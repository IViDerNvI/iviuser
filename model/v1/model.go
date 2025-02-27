package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

type User struct {
	ObjMeta    `json:",inline"`
	UserName   string `json:"username" gorm:"primaryKey;column:username" validate:"required"`
	NickName   string `json:"nickname" gorm:"column:nickname;not null" validate:"required"`
	Email      string `json:"email" gorm:"column:email" validate:"required,email"`
	Phone      string `json:"phone" gorm:"column:phone"`
	Bio        string `json:"bio" gorm:"column:bio"`
	Company    string `json:"company" gorm:"column:company"`
	Location   string `json:"location" gorm:"column:location"`
	ProfileURL string `json:"profile_url" gorm:"profile_url"`
}

type UserList struct {
	ListMeta `json:",inline"`
	Items    []User
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) String() string {
	user_info, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(user_info)
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ObjMeta = ObjMeta{}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.ObjMeta = ObjMeta{}
	return nil
}

func (u *User) AfterFind(tx *gorm.DB) error {
	u.ObjMeta = ObjMeta{}
	return nil
}
