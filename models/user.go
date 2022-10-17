package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `json:"username" gorm:"unique_index, not null" form:"username" valid:"required~Username is required"`
	Email    string `json:"email" gorm:"unique_index, not null" form:"email" valid:"required~Email is required, email~Email is invalid"`
	Password string `json:"password" gorm:"not null" form:"password" valid:"required~Password is required, length(6)~Password must be at least 6 characters"`
	Age      int    `json:"age" gorm:"not null" form:"age" valid:"required~Age is required, range(8)~Age must be at least 8"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
