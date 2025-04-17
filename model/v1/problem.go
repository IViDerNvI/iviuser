package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ividernvi/iviuser/pkg/util/idutil"
	"gorm.io/gorm"
)

type Problem struct {
	ObjMeta    `json:",inline"`
	Title      string `json:"title" gorm:"column:title" validate:"required"`
	Descrition string `json:"descrition" gorm:"column:descrition" validate:"required"`
	Author     string `json:"author" gorm:"column:author" validate:"required"`
}

type ProblemList struct {
	ListMeta `json:",inline"`
	Items    []Problem `json:"items"`
}

func (p *Problem) TableName() string {
	return "problems"
}

func (p *Problem) Validate() error {
	validator := validator.New()
	if err := validator.Struct(p); err != nil {
		return err
	}
	return nil
}

func (p *Problem) BeforeCreate(tx *gorm.DB) error {
	p.InstanceID = uint(idutil.SnowflakeID())
	return nil
}

func (p *Problem) BeforeUpdate(tx *gorm.DB) error {
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Problem) AfterGet(tx *gorm.DB) error {
	p.ID = 0
	return nil
}
