package usecase

import (
	"mime/multipart"

	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	"github.com/sawalreverr/recything/internal/task/user_task/dto"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
)

type UserTaskUsecase interface {
	GetAllTasksUsecase() ([]task.TaskChallenge, error)
	GetTaskByIdUsecase(id string) (*task.TaskChallenge, error)
	CreateUserTaskUsecase(taskChallengeId string, userId string) (*user_task.UserTaskChallenge, error)
	UploadImageTaskUsecase(request *dto.UploadImageTask, fileImage []*multipart.FileHeader, userId string, userTaskId string) (*user_task.UserTaskChallenge, error)
	GetUserTaskByUserIdUsecase(userId string) ([]user_task.UserTaskChallenge, error)
	GetUserTaskDoneByUserIdUsecase(userId string) ([]user_task.UserTaskChallenge, error)
	UpdateUserTaskUsecase(request *dto.UpdateUserTaskRequest, fileImage []*multipart.FileHeader, userId string, userTaskId string) (*user_task.UserTaskChallenge, error)
	GetUserTaskDetailsUsecase(userTaskId string, userId string) (*user_task.UserTaskChallenge, []*user_task.UserTaskImage, error)
	GetHistoryPointByUserIdUsecase(userId string) ([]user_task.UserTaskChallenge, int, error)
	UpdateTaskStepUsecase(request *dto.UpdateTaskStepRequest, userId string) (*user_task.UserTaskChallenge, error)
	GetUserTaskByUserTaskId(userId string, userTaskId string) (*user_task.UserTaskChallenge, error)
	GetUserTaskRejectedByUserId(userId string, userTaskId string) (*user_task.UserTaskChallenge, error)
}
