package usecase

import (
	"mime/multipart"

	"github.com/sawalreverr/recything/internal/task/manage_task/dto"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
)

type ManageTaskUsecase interface {
	CreateTaskUsecase(request *dto.CreateTaskResquest, thumbnail []*multipart.FileHeader, adminId string) (*task.TaskChallenge, error)
	GetTaskChallengePagination(page int, limit int, status string, endDate string) ([]task.TaskChallenge, int, error)
	GetTaskByIdUsecase(id string) (*task.TaskChallenge, error)
	UpdateTaskChallengeUsecase(request *dto.UpdateTaskRequest, thumbnail []*multipart.FileHeader, id string) (*task.TaskChallenge, error)
	DeleteTaskChallengeUsecase(id string) error
}
