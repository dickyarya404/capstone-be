package usecase

import (
	"github.com/sawalreverr/recything/internal/achievements/user_achievements/dto"
	"github.com/sawalreverr/recything/internal/achievements/user_achievements/repository"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type UserAchievementUsecaseImpl struct {
	userAchievementRepository repository.UserAchievementRepository
}

func NewUserAchievementUsecase(userAchievementRepository repository.UserAchievementRepository) *UserAchievementUsecaseImpl {
	return &UserAchievementUsecaseImpl{userAchievementRepository: userAchievementRepository}
}

func (usecase *UserAchievementUsecaseImpl) GetAvhievementsByUser(userId string) (*dto.GetAchievementByUserResponse, error) {

	achievements, err := usecase.userAchievementRepository.GetAvhievementsByUser()
	if err != nil {
		return nil, err
	}

	historyUserPoint, err := usecase.userAchievementRepository.GetHistoryUserPoint(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, pkg.ErrUserNotHasHistoryPoint
		}
		return nil, err
	}
	dataUser, err := usecase.userAchievementRepository.GetPoinUser(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, pkg.ErrUserNotFound
		}
		return nil, err
	}
	var dataachievement []*dto.DataAchievement
	var dataHistoryUserPoint []*dto.HistoryUserPoint
	data := dto.GetAchievementByUserResponse{
		DataAchievement: dataachievement,
		DataUser: &dto.DataUser{
			Id:    userId,
			Name:  dataUser.Name,
			Point: int(dataUser.Point),
			Badge: dataUser.Badge,
		},
		HistoryUserPoint: dataHistoryUserPoint,
	}

	for _, v := range *achievements {
		dataachievement = append(dataachievement, &dto.DataAchievement{
			Id:          v.ID,
			Level:       v.Level,
			TargetPoint: v.TargetPoint,
			BadgeUrl:    v.BadgeUrl,
		})
	}
	for _, v := range *historyUserPoint {
		dataHistoryUserPoint = append(dataHistoryUserPoint, &dto.HistoryUserPoint{
			Point:      v.Point,
			AcceptedAt: v.AcceptedAt,
		})
	}
	data.DataAchievement = dataachievement
	data.HistoryUserPoint = dataHistoryUserPoint
	return &data, nil

}
