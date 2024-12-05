package customdata

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomData struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Topic       string `json:"topic"`
	Description string `json:"description"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type CustomDataRepository interface {
	Create(data CustomData) (*CustomData, error)
	FindByID(dataID string) (*CustomData, error)
	FindAll(page int, limit int, sortBy string, sortType string) (*[]CustomData, int64, error)
	FindLastID() (string, error)
	Update(data CustomData) error
	Delete(dataID string) error
}

type CustomDataUsecase interface {
	NewCustomData(data CustomDataInput) (*CustomDataResponse, error)
	FindByID(dataID string) (*CustomDataResponse, error)
	FindAll(page int, limit int, sortBy string, sortType string) (*CustomDataPaginationResponse, error)
	UpdateData(dataID string, data CustomDataInput) error
	DeleteData(dataID string) error
}

type CustomDataHandler interface {
	NewCustomData(c echo.Context) error
	GetDataByID(c echo.Context) error
	GetAllData(c echo.Context) error
	UpdateData(c echo.Context) error
	DeleteData(c echo.Context) error
}
