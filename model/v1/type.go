package model

import "time"

type ObjMeta struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type ListMeta struct {
	TotalItems int `json:"totalItems"`
}
