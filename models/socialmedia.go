package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SocialMedia struct {
	Base
	Name           string `gorm:"type:VARCHAR(50);not null" valid:"required" form:"name" json:"name"`
	SocialMediaUrl string `gorm:"not null" valid:"required" form:"social_media_url" json:"social_media_url"`
	UserID         string `gorm:"type:VARCHAR(50);not null" json:"user_id"`
	User           *User  `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}

func (s *SocialMedia) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(s); err != nil {
		return err
	}

	s.ID = uuid.New().String()

	return
}

func (s *SocialMedia) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(s); err != nil {
		return err
	}
	return
}
