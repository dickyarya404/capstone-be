package dto

import "mime/multipart"

type CreateDataVideoRequest struct {
	Title           string                `json:"title" validate:"required"`
	Description     string                `json:"description" validate:"required"`
	LinkVideo       string                `json:"link_video" validate:"required"`
	ContentCategory []DataCategoryVideo   `json:"content_categories" validate:"required"`
	WasteCategory   []DataTrashCategory   `json:"waste_categories" validate:"required"`
	Thumbnail       *multipart.FileHeader `json:"-"`
}

type DataCategoryVideo struct {
	Name string `json:"name"`
}

type DataTrashCategory struct {
	Name string `json:"name"`
}

type CreateCategoryVideoRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateDataVideoRequest struct {
	Title             string                `json:"title"`
	Description       string                `json:"description"`
	LinkVideo         string                `json:"link_video"`
	ContentCategories []DataCategoryVideo   `json:"content_categories"`
	WasteCategories   []DataTrashCategory   `json:"waste_categories"`
	Thumbnail         *multipart.FileHeader `json:"-"`
}
