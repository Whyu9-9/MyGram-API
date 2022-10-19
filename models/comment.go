package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `json:"message" gorm:"not null" form:"message" valid:"required~Message is required"`
	UserId  uint   `json:"user_id" form:"user_id"`
	User    *User  `json:"user"`
	PhotoId uint   `json:"photo_id" form:"photo_id"`
	Photo   *Photo `json:"photo"`
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

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	if c.Message == "" {
		err = errors.New("Message is required")
		return
	}

	err = nil
	return
}
