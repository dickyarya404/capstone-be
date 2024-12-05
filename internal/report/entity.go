package report

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// struct
type Report struct {
	ID          string  `json:"id" gorm:"primaryKey"`
	AuthorID    string  `json:"author_id"`
	ReportType  string  `json:"report_type" gorm:"type:enum('littering', 'rubbish');"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	WasteType   string  `json:"waste_type" gorm:"type:enum('sampah basah', 'sampah kering', 'sampah basah,sampah kering', 'organik', 'anorganik', 'berbahaya');" `
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Address     string  `json:"address"`
	City        string  `json:"city"`
	Province    string  `json:"province"`
	Status      string  `json:"status" gorm:"type:enum('need review', 'approve', 'reject');default:'need review'"`
	Reason      string  `json:"reason"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type WasteMaterial struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Type string `json:"type" gorm:"type:enum('plastik', 'kaca', 'kayu', 'kertas', 'baterai', 'besi', 'limbah berbahaya', 'limbah beracun', 'sisa makanan', 'tak terdeteksi');" `

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type ReportWasteMaterial struct {
	ID              uuid.UUID `json:"id" gorm:"primaryKey"`
	ReportID        string    `json:"report_id"`
	WasteMaterialID string    `json:"waste_material_id"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type ReportImage struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey"`
	ReportID string    `json:"report_id"`
	ImageURL string    `json:"image_url"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// interface
type ReportRepository interface {
	Create(report Report) (*Report, error)
	FindByID(reportID string) (*Report, error)
	FindAll(page, limit int, reportType, status string, date time.Time) (*[]Report, int64, error)
	FindAllReportsByUser(userID string, limit int) (*[]Report, error)
	FindLastID() (string, error)
	Update(report Report) error
	Delete(reportID string) error

	AddImage(image ReportImage) (*ReportImage, error)
	DeleteImage(imageID string, reportID string) error
	DeleteAllImage(reportID string) error
	FindAllImage(reportID string) (*[]string, error)

	AddReportMaterial(material ReportWasteMaterial) (*ReportWasteMaterial, error)
	DeleteAllReportMaterial(reportID string) error
	FindAllReportMaterial(reportID string) (*[]WasteMaterial, error)

	FindWasteMaterialByID(materialID string) (*WasteMaterial, error)
	FindWasteMaterialByType(materialType string) (*WasteMaterial, error)
}

type ReportUsecase interface {
	CreateReport(report ReportInput, authorID string, imageURLs []string) (*ReportDetail, error)
	FindHistoryUserReports(authorID string) (*[]ReportDetail, error)

	UpdateStatusReport(report UpdateStatus, reportID string) error
	FindAllReports(page, limit int, reportType, status string, date time.Time) (*[]ReportDetail, int64, error)
}

type ReportHandler interface {
	NewReport(c echo.Context) error
	GetHistoryUserReports(c echo.Context) error

	UpdateStatus(c echo.Context) error
	GetAllReports(c echo.Context) error
}
