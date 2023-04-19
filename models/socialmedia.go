package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Represents a social media account for a user
type SocialMedia struct {
	GormModel
	Name             string `json:"name" form:"name" valid:"required~Name is required"`
	Social_media_url string `json:"social_media_url" form:"social_media_url" valid:"required~URL is required"`
	UserID           uint
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
