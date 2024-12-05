package dto

import "time"

type CreateTaskResponse struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Thumbnail   string      `json:"thumbnail"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     time.Time   `json:"end_date"`
	Point       int         `json:"point"`
	Status      bool        `json:"status"`
	Steps       []TaskSteps `json:"steps"`
}

type DataTasks struct {
	Id          string           `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Thumbnail   string           `json:"thumbnail"`
	StartDate   time.Time        `json:"start_date"`
	EndDate     time.Time        `json:"end_date"`
	Point       int              `json:"point"`
	Status      bool             `json:"status"`
	Steps       []TaskSteps      `json:"steps"`
	TaskCreator TaskCreatorAdmin `json:"task_creator"`
}

type TaskCreatorAdmin struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetTaskPagination struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      []DataTasks `json:"data"`
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	TotalData int         `json:"total_data"`
	TotalPage int         `json:"total_page"`
}

type TaskGetByIdResponse struct {
	Id          string           `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Thumbnail   string           `json:"thumbnail"`
	StartDate   time.Time        `json:"start_date"`
	EndDate     time.Time        `json:"end_date"`
	Point       int              `json:"point"`
	Status      bool             `json:"status"`
	Steps       []TaskSteps      `json:"steps"`
	TaskCreator TaskCreatorAdmin `json:"task_creator"`
}

type UpdateTaskResponse struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Thumbnail   string      `json:"thumbnail"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     time.Time   `json:"end_date"`
	Point       int         `json:"point"`
	Steps       []TaskSteps `json:"steps"`
}
