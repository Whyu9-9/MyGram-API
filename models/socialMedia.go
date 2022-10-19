package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `json:"name" gorm:"not null" form:"name" valid:"required~Social Media Name is required"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" form:"social_media_url" valid:"required~social media url is required"`
	UserId         uint   `json:"user_id" form:"user_id"`
	User           *User  `json:"user"`
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

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	if s.Name == "" && s.SocialMediaUrl == "" {
		err = errors.New("Name and Social Media Url is required")
		return
	} else if s.Name == "" {
		err = errors.New("Name is required")
		return
	} else if s.SocialMediaUrl == "" {
		err = errors.New("Social Media Url is required")
		return
	}

	err = nil
	return
}
