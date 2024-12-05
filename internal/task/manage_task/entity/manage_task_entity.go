package entity

import (
	"time"

	admin "github.com/sawalreverr/recything/internal/admin/entity"
	"gorm.io/gorm"
)

type TaskChallenge struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	Thumbnail   string
	StartDate   time.Time
	EndDate     time.Time
	Point       int
	Status      bool
	TaskSteps   []TaskStep     `gorm:"foreignKey:TaskChallengeId"`
	AdminId     string         `gorm:"index"`
	Admin       admin.Admin    `gorm:"foreignKey:AdminId"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type TaskStep struct {
	ID              int    `gorm:"primaryKey"`
	TaskChallengeId string `gorm:"index"`
	Title           string
	Description     string
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
