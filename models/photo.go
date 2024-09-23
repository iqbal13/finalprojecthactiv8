package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	Base
	Title    string   `gorm:"type:VARCHAR(50);not null" valid:"required" form:"title" json:"title"`
	Caption  string   `form:"caption" json:"caption"`
	PhotoUrl string   `gorm:"not null" valid:"required" form:"photo_url" json:"photo_url"`
	UserID   string   `gorm:"type:VARCHAR(50);not null" json:"user_id"`
	User     *User    `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Comment  *Comment `json:"-"`
}

func (photo *Photo) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(photo); err != nil {
		return err
	}
	photo.ID = uuid.New().String()

	return
}

func (photo *Photo) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(photo); err != nil {
		return err
	}
	return
}
