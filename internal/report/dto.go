package report

import (
	"mime/multipart"
	"time"
)

type ReportInput struct {
	ReportType     string                  `json:"report_type" validate:"required"`
	Title          string                  `json:"title" validate:"max=100"`
	Description    string                  `json:"description" validate:"required"`
	WasteType      string                  `json:"waste_type" validate:"required"`
	WasteMaterials []string                `json:"waste_materials"`
	Latitude       float64                 `json:"latitude" validate:"required,latitude"`
	Longitude      float64                 `json:"longitude" validate:"required,longitude"`
	Address        string                  `json:"address" validate:"required"`
	City           string                  `json:"city" validate:"required"`
	Province       string                  `json:"province" validate:"required"`
	ReportImages   []*multipart.FileHeader `json:"-"`
}

type UpdateStatus struct {
	Status string `json:"status" validate:"required,oneof='approve' 'reject'"`
	Reason string `json:"reason"`
}

type UserDetail struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type ReportDetail struct {
	ID          string     `json:"id"`
	Author      UserDetail `json:"author"`
	ReportType  string     `json:"report_type"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	WasteType   string     `json:"waste_type"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	Address     string     `json:"address"`
	City        string     `json:"city"`
	Province    string     `json:"province"`
	Status      string     `json:"status"`
	Reason      string     `json:"reason"`

	WasteMaterials []WasteMaterial `json:"waste_materials"`
	ReportImages   []string        `json:"report_images"`
	CreatedAt      time.Time       `json:"created_at"`
}

type ReportResponsePagination struct {
	Total  int64          `json:"total"`
	Page   int            `json:"page"`
	Limit  int            `json:"limit"`
	Report []ReportDetail `json:"reports"`
}
