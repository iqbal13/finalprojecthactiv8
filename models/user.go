package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/iqbal13/finalprojecthactiv8/helpers"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username     string         `gorm:"type:VARCHAR(50);uniqueIndex;not null" valid:"required" form:"username" json:"username"`
	Email        string         `gorm:"type:VARCHAR(50);uniqueIndex;not null" valid:"email,required" form:"email" json:"email"`
	Password     string         `gorm:"not null" valid:"required,minstringlength(6)" form:"password" json:"password,omitempty"`
	Age          uint           `gorm:"not null" valid:"required,range(8|100)" form:"age" json:"age,omitempty"`
	Photos       *[]Photo       `json:"-"`
	SocialMedias *[]SocialMedia `json:"-"`
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(user); err != nil {
		return err
	}
	user.ID = uuid.New().String()

	user.Password = helpers.Hash(user.Password)

	return
}

func (user *User) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(user); err != nil {
		return err
	}

	return
}
