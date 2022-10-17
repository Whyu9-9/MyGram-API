package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `json:"title" gorm:"not null" form:"title" valid:"required~Title is required"`
	Caption  string `json:"caption" form:"caption" valid:"required~Caption is required"`
	PhotoUrl string `json:"photo_url" gorm:"not null" form:"photo_url" valid:"required~PhotoUrl is required"`
	UserId   uint   `json:"user_id" form:"user_id"`
	User     *User
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
