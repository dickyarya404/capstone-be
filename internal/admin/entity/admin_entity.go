package entity

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	Role      string `gorm:"type:enum('super admin', 'admin')"`
	ImageUrl  string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
