package models

type SocialMedia struct {
	GormModel
	Name             string `json:"name" form:"name" valid:"required~Name is required"`
	Social_media_url string `json:"social_media_url" form:"social_media_url" valid:"required~URL is required"`
	UserID           uint
}
