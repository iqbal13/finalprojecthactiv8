package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string         `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt *time.Time     `gorm:"not null;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt *time.Time     `gorm:"not null;autoCreateTime" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
