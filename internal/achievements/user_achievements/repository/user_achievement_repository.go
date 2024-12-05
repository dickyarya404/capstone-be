package repository

import (
	archievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
	user "github.com/sawalreverr/recything/internal/user"
)

type UserAchievementRepository interface {
	GetAvhievementsByUser() (*[]archievement.Achievement, error)
	GetHistoryUserPoint(userId string) (*[]user_task.UserTaskChallenge, error)
	GetPoinUser(userId string) (*user.User, error)
}
