package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	PhotoTitle string `json"photo_title" form:"photo_title" valid:"required~Title is required"`
	Caption    string `json:"caption,omitempty" form:"caption"`
	PhotoURL   string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo is required"`
	UserID     uint
	Comments   []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments,omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
