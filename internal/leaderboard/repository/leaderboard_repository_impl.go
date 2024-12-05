package repository

import (
	"github.com/sawalreverr/recything/internal/database"
	user "github.com/sawalreverr/recything/internal/user"
)

type LeaderboardRepositoryImpl struct {
	DB database.Database
}

func NewLeaderboardRepository(db database.Database) *LeaderboardRepositoryImpl {
	return &LeaderboardRepositoryImpl{DB: db}
}

func (repository *LeaderboardRepositoryImpl) GetLeaderboard() (*[]user.User, error) {
	var users []user.User
	if err := repository.DB.GetDB().
		Order("point desc").
		Limit(10).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}
