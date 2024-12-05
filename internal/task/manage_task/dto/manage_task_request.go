package dto

import "mime/multipart"

type CreateTaskResquest struct {
	Title       string                `json:"title" validate:"required"`
	Description string                `json:"description" validate:"required"`
	StartDate   string                `json:"start_date" validate:"required"`
	EndDate     string                `json:"end_date" validate:"required"`
	Point       int                   `json:"point" validate:"required"`
	Thumbnail   *multipart.FileHeader `json:"-"`
	TaskSteps   []TaskSteps           `json:"task_steps" validate:"required"`
}

type TaskSteps struct {
	Id          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateTaskRequest struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	StartDate   string                `json:"start_date"`
	EndDate     string                `json:"end_date"`
	Point       int                   `json:"point"`
	Thumbnail   *multipart.FileHeader `json:"-"`
	TaskSteps   []TaskSteps           `json:"task_steps" validate:"required"`
}
