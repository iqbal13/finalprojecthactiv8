package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	Base
	UserID  string `gorm:"type:VARCHAR(50);not null" json:"user_id"`
	PhotoID string `gorm:"type:VARCHAR(50);not null" form:"photo_id" json:"photo_id"`
	Message string `gorm:"not null" valid:"required" form:"message" json:"message"`
	User    *User  `gorm:"foreignKey:UserID;constraint:opUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Photo   *Photo `gorm:"foreignKey:PhotoID;constraint:opUpdate:CASCADE,onDelete:CASCADE" json:"photo"`
}

func (c *Comment) BeforeCreate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(c); err != nil {
		return err
	}
	c.ID = uuid.New().String()

	return
}

func (c *Comment) BeforeUpdate(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(c); err != nil {
		return err
	}
	return
}
