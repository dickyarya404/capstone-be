package usecase

import "github.com/sawalreverr/recything/internal/achievements/user_achievements/dto"

type UserAchievementUsecase interface {
	GetAvhievementsByUser(userId string) (*dto.GetAchievementByUserResponse, error)
}
