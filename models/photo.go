package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `json:"title" gorm:"not null" form:"title" valid:"required~Title is required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" gorm:"not null" form:"photo_url" valid:"required~PhotoUrl is required"`
	UserId   uint   `json:"user_id" form:"user_id"`
	User     *User  `json:"User"`
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

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	if p.Title == "" && p.PhotoUrl == "" {
		err = errors.New("Title and PhotoUrl is required")
		return
	} else if p.Title == "" {
		err = errors.New("Title is required")
		return
	} else if p.PhotoUrl == "" {
		err = errors.New("PhotoUrl is required")
		return
	}

	err = nil
	return
}
