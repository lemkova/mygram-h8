package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment model that belongs to a photo
type Comment struct {
	GormModel
	UserID  uint
	PhotoID uint
	Message string `json:"message" form:"message" valid:"required~Message cannot be empty"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
