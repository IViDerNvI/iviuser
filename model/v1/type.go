package model

import "time"

type ObjMeta struct {
	ID         uint       `gorm:"primary_key"`
	InstanceID uint       `json:"instanceID" gorm:"column:instance_id;unique"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `sql:"index"`
}

type ListMeta struct {
	TotalItems int64 `json:"totalItems"`
}
