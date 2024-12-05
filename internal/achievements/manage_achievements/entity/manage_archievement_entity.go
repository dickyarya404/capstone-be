package entity

import (
	"time"

	"gorm.io/gorm"
)

type Achievement struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Level        string `gorm:"unique"`
	TargetPoint  int
	BadgeUrl     string
	BadgeUrlUser string
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
