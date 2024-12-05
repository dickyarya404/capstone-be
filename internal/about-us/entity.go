package aboutus

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AboutUs struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Category    string `json:"category"`
	Title       string `json:"title"`
	Description string `json:"description"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// perusahaan, tim, contact_us

type AboutUsImage struct {
	ID        string `json:"id" gorm:"primaryKey"`
	AboutUsID string `json:"about_us_id"`
	Name      string `json:"name"`
	ImageURL  string `json:"image_url"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type AboutUsRepository interface {
	FindByCategory(categoryName string) (*[]AboutUs, error)
	FindAllImageByID(aboutUsID string) (*[]AboutUsImage, error)
}

type AboutUsUsecase interface {
	GetAboutUsByCategory(categoryName string) (*[]AboutUsResponse, error)
}

type AboutUsHandler interface {
	GetAboutUsByCategory(c echo.Context) error
}
