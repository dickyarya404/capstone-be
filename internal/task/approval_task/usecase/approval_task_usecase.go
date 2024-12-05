package usecase

import (
	"github.com/sawalreverr/recything/internal/task/approval_task/dto"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
)

type ApprovalTaskUsecase interface {
	GetAllApprovalTaskPaginationUseCase(limit int, offset int) ([]*user_task.UserTaskChallenge, int, error)
	ApproveUserTaskUseCase(userTaskId string) error
	RejectUserTaskUseCase(request *dto.RejectUserTaskRequest, userTaskId string) error
	GetUserTaskDetailsUseCase(userTaskId string) (*user_task.UserTaskChallenge, []*user_task.UserTaskImage, error)
}
