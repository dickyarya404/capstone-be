package repository

import (
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
)

type ManageTaskRepository interface {
	CreateTask(task *task.TaskChallenge) (*task.TaskChallenge, error)
	FindLastIdTaskChallenge() (string, error)
	GetTaskChallengePagination(page int, limit int, status string, endDate string) ([]task.TaskChallenge, int, error)
	GetTaskById(id string) (*task.TaskChallenge, error)
	FindTask(id string) (*task.TaskChallenge, error)
	UpdateTaskChallenge(taskChallenge *task.TaskChallenge, taskId string) (*task.TaskChallenge, error)
	DeleteTaskChallenge(taskId string) error
	UpdateTaskChallengeStatus() error
}
