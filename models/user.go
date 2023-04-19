package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"

	"github.com/lemkova/mygram-h8/helpers"
)

// User model
type User struct {
	GormModel
	Username     string        `json:"username" form:"username" gorm:"not null;uniqueIndex" valid:"required~Username is required"`
	Email        string        `json:"email" form:"email" gorm:"not null;uniqueIndex" valid:"required~Email is required,email~Invalid email format"`
	Password     string        `json:"password" form:"password" gorm:"not null" valid:"required~Password is required,minstringlength(6)~Password must be at least 6 characters"`
	Age          int           `json:"age" form:"age" valid:"required~Age is required,range(8|100)~Age must be at least 8 years old"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos,omitempty"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments,omitempty"`
	Social_Media []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
