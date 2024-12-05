package repository

import user "github.com/sawalreverr/recything/internal/user"

type LeaderboardRepository interface {
	GetLeaderboard() (*[]user.User, error)
}
