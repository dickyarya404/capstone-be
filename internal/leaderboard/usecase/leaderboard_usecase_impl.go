package usecase

import (
	"github.com/sawalreverr/recything/internal/leaderboard/dto"
	"github.com/sawalreverr/recything/internal/leaderboard/repository"
)

type LeaderboardUsecaseImpl struct {
	LeaderboardRepository repository.LeaderboardRepository
}

func NewLeaderboardUsecase(leaderboardRepository repository.LeaderboardRepository) *LeaderboardUsecaseImpl {
	return &LeaderboardUsecaseImpl{LeaderboardRepository: leaderboardRepository}
}

func (usecase *LeaderboardUsecaseImpl) GetLeaderboardUsecase() (*dto.LeaderboardResponse, error) {
	users, err := usecase.LeaderboardRepository.GetLeaderboard()
	if err != nil {
		return nil, err
	}
	var dataLeaderboard []*dto.DataLeaderboard
	for _, user := range *users {
		dataLeaderboard = append(dataLeaderboard, &dto.DataLeaderboard{
			Id:         user.ID,
			Name:       user.Name,
			PictureURL: user.PictureURL,
			Point:      int(user.Point),
			Badge:      user.Badge,
			Address:    user.Address,
		})
	}
	return &dto.LeaderboardResponse{DataLeaderboard: dataLeaderboard}, nil
}
