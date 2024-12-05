package entity

import (
	"time"

	art "github.com/sawalreverr/recything/internal/article"
	"github.com/sawalreverr/recything/internal/user"
	"gorm.io/gorm"
)

type Video struct {
	ID                int `gorm:"primaryKey"`
	Title             string
	Description       string
	Thumbnail         string
	Link              string
	Viewer            int
	Categories        []VideoCategory
	ContentCategories []art.ContentCategory `gorm:"-"`
	WasteCategories   []art.WasteCategory   `gorm:"-"`
	CreatedAt         time.Time             `gorm:"autoCreateTime"`
	UpdatedAt         time.Time             `gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt        `gorm:"index"`
}

type VideoCategory struct {
	ID                int `gorm:"primaryKey"`
	VideoID           int
	Video             Video `gorm:"foreignKey:VideoID"`
	ContentCategoryID uint
	ContentCategory   art.ContentCategory `gorm:"foreignKey:ContentCategoryID"`
	WasteCategoryID   uint
	WasteCategory     art.WasteCategory `gorm:"foreignKey:WasteCategoryID"`
	CreatedAt         time.Time         `gorm:"autoCreateTime"`
	UpdatedAt         time.Time         `gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt    `gorm:"index"`
}

type Comment struct {
	ID        int       `gorm:"primaryKey"`
	VideoID   int       `gorm:"index"`
	Video     Video     `gorm:"foreignKey:VideoID"`
	UserID    string    `gorm:"index"`
	User      user.User `gorm:"foreignKey:UserID"`
	Comment   string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
