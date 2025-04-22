package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ividernvi/iviuser/pkg/util/idutil"
	"gorm.io/gorm"
)

type Subscribe struct {
	ObjMeta  `json:",inline"`
	UserName string `json:"username" gorm:"column:username;uniqueIndex:idx_username_itemtype_itemid;type:varchar(255)" validate:"required"`
	ItemType string `json:"item_type" gorm:"column:item_type;uniqueIndex:idx_username_itemtype_itemid;type:varchar(255)" validate:"required;oneof=user"`
	ItemID   uint   `json:"item_id" gorm:"column:item_id;uniqueIndex:idx_username_itemtype_itemid" validate:"required"`
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

func (s *Subscribe) BeforeCreate(tx *gorm.DB) error {
	s.InstanceID = uint(idutil.SnowflakeID())
	return nil
}

func (s *Subscribe) BeforeUpdate(tx *gorm.DB) error {
	s.UpdatedAt = time.Now()
	return nil
}

func (s *Subscribe) Override(new *Subscribe) *Subscribe {
	s.ItemType = new.ItemType
	s.ItemID = new.ItemID
	return s
}
