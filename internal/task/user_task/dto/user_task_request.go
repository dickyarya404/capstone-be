package dto

import "mime/multipart"

type UploadImageTask struct {
	Description string                  `json:"description" validate:"required"`
	Images      []*multipart.FileHeader `json:"-"`
}

type UpdateUserTaskRequest struct {
	Description string                  `json:"description" validate:"required"`
	Images      []*multipart.FileHeader `json:"-"`
}

type UpdateTaskStepRequest struct {
	UserTaskId string `json:"user_task_id"`
	TaskStepId int    `json:"task_step_id"`
}
