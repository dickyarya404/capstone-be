package usecase

import "github.com/sawalreverr/recything/internal/leaderboard/dto"

type LeaderboardUsecase interface {
	GetLeaderboardUsecase() (*dto.LeaderboardResponse, error)
}
