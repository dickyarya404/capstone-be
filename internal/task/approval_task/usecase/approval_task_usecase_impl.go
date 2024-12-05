package usecase

import (
	"github.com/sawalreverr/recything/internal/task/approval_task/dto"
	"github.com/sawalreverr/recything/internal/task/approval_task/repository"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type ApprovalTaskUsecaseImpl struct {
	ApprovalTaskRepository repository.ApprovalTaskRepository
}

func NewApprovalTaskUsecase(approvalTaskRepository repository.ApprovalTaskRepository) *ApprovalTaskUsecaseImpl {
	return &ApprovalTaskUsecaseImpl{ApprovalTaskRepository: approvalTaskRepository}
}

func (usecase *ApprovalTaskUsecaseImpl) GetAllApprovalTaskPaginationUseCase(limit int, offset int) ([]*user_task.UserTaskChallenge, int, error) {
	task, total, err := usecase.ApprovalTaskRepository.GetAllApprovalTaskPagination(limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return task, total, nil
}

func (usecase *ApprovalTaskUsecaseImpl) ApproveUserTaskUseCase(userTaskId string) error {
	userTask, err := usecase.ApprovalTaskRepository.FindUserTask(userTaskId)
	if err != nil {
		return pkg.ErrUserTaskNotFound
	}

	if userTask.StatusAccept == "accept" {
		return pkg.ErrUserTaskAlreadyApprove
	}

	if err := usecase.ApprovalTaskRepository.ApproveUserTask(userTaskId); err != nil {
		return err
	}
	return nil

}

func (usecase *ApprovalTaskUsecaseImpl) RejectUserTaskUseCase(request *dto.RejectUserTaskRequest, userTaskId string) error {
	userTask, err := usecase.ApprovalTaskRepository.FindUserTask(userTaskId)
	if err != nil {
		return pkg.ErrUserTaskNotFound
	}
	if userTask.StatusAccept == "reject" {
		return pkg.ErrUserTaskAlreadyReject
	}

	if userTask.StatusAccept == "accept" {
		return pkg.ErrUserTaskAlreadyAccepted
	}

	status := "reject"
	if err := usecase.ApprovalTaskRepository.RejectUserTask(&user_task.UserTaskChallenge{
		StatusAccept: status,
		Reason:       request.Reason,
	}, userTaskId); err != nil {
		return err
	}
	return nil
}

func (usecase *ApprovalTaskUsecaseImpl) GetUserTaskDetailsUseCase(userTaskId string) (*user_task.UserTaskChallenge, []*user_task.UserTaskImage, error) {
	task, images, err := usecase.ApprovalTaskRepository.GetUserTaskDetails(userTaskId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, pkg.ErrUserTaskNotFound
		}
		return nil, nil, err
	}
	return task, images, nil
}
