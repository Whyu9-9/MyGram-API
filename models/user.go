package models

import (
	"mygram-api/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string        `json:"username" gorm:"unique;not null" form:"username" valid:"required~Username is required"`
	Email       string        `json:"email" gorm:"unique;not null" form:"email" valid:"required~Email is required, email~Email is invalid"`
	Password    string        `json:"password" gorm:"not null" form:"password" valid:"required~Password is required, minstringlength(6)~Password must be at least 6 characters"`
	Age         int           `json:"age" gorm:"not null" form:"age" valid:"required~Age is required, range(8|100)~Age must be at least 8"`
	Photo       []Photo       `json:"photo" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comment     []Comment     `json:"comment" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedia []SocialMedia `json:"social_media" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPassword(u.Password)

	err = nil
	return
}
