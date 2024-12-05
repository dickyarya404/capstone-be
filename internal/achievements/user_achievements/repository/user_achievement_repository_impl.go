package repository

import (
	archievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
	"github.com/sawalreverr/recything/internal/database"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
	user "github.com/sawalreverr/recything/internal/user"
)

type UserAchievementRepositoryImpl struct {
	DB database.Database
}

func NewUserAchievementRepository(db database.Database) *UserAchievementRepositoryImpl {
	return &UserAchievementRepositoryImpl{DB: db}
}

func (repository UserAchievementRepositoryImpl) GetAvhievementsByUser() (*[]archievement.Achievement, error) {
	var achievements []archievement.Achievement
	if err := repository.DB.GetDB().Find(&achievements).Error; err != nil {
		return nil, err
	}
	return &achievements, nil
}

func (repository UserAchievementRepositoryImpl) GetHistoryUserPoint(userId string) (*[]user_task.UserTaskChallenge, error) {
	var userTasks []user_task.UserTaskChallenge
	if err := repository.DB.GetDB().
		Where("user_id = ?", userId).
		Where("status_accept = ?", "accept").
		Order("accepted_at desc").
		Find(&userTasks).Error; err != nil {
		return nil, err
	}
	return &userTasks, nil
}

func (repository UserAchievementRepositoryImpl) GetPoinUser(userId string) (*user.User, error) {
	var user user.User
	if err := repository.DB.GetDB().Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
