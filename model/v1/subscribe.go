package model

import "github.com/go-playground/validator/v10"

type Subscribe struct {
	ObjMeta  `json:",inline"`
	UserName string `json:"username" gorm:"column:username" validate:"required"`
	ItemType string `json:"item_type" gorm:"column:item_type" validate:"required"`
	ItemID   string `json:"item_id" gorm:"column:item_id" validate:"required"`
}

type SubscribeList struct {
	ListMeta `json:",inline"`
	Items    []Subscribe `json:"items"`
}

func (s *Subscribe) TableName() string {
	return "subscribes"
}

func (s *Subscribe) Validate() error {
	validator := validator.New()
	if err := validator.Struct(s); err != nil {
		return err
	}
	return nil
}
