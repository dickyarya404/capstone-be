package faq

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type FAQ struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Category string `json:"category" gorm:"index"`
	Question string `json:"question"`
	Answer   string `json:"answer"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type FaqRepository interface {
	FindAll() (*[]FAQ, error)
	FindByCategory(category string) (*[]FAQ, error)
	FindByKeyword(keyword string) (*[]FAQ, error)
}

type FaqUsecase interface {
	GetAllFaqs() (*[]FaqResponse, error)
	GetFaqsByCategory(category string) (*[]FaqResponse, error)
	GetFaqsByKeyword(keyword string) (*[]FaqResponse, error)
}

type FaqHandler interface {
	GetAllFaqs(c echo.Context) error
	GetFaqsByCategory(c echo.Context) error
	GetFaqsByKeyword(c echo.Context) error
}
