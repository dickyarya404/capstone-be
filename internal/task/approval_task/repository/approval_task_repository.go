package repository

import (
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
)

type ApprovalTaskRepository interface {
	GetAllApprovalTaskPagination(limit int, offset int) ([]*user_task.UserTaskChallenge, int, error)
	FindUserTask(userTaskId string) (*user_task.UserTaskChallenge, error)
	ApproveUserTask(userTaskId string) error
	RejectUserTask(data *user_task.UserTaskChallenge, userTaskId string) error
	GetUserTaskDetails(userTaskId string) (*user_task.UserTaskChallenge, []*user_task.UserTaskImage, error)
	FindUserTaskForApprove(userTaskId string) (*user_task.UserTaskChallenge, error)
	FindUserTaskForReject(userTaskId string) (*user_task.UserTaskChallenge, error)
}
