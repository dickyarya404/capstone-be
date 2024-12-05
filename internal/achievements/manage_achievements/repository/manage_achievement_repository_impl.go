package repository

import (
	archievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
	"github.com/sawalreverr/recything/internal/database"
)

type ManageAchievementRepositoryImpl struct {
	DB database.Database
}

func NewManageAchievementRepository(db database.Database) *ManageAchievementRepositoryImpl {
	return &ManageAchievementRepositoryImpl{DB: db}
}

func (repository ManageAchievementRepositoryImpl) FindArchievementByLevel(level string) (*archievement.Achievement, error) {
	var achievement archievement.Achievement
	if err := repository.DB.GetDB().Where("level = ?", level).First(&achievement).Error; err != nil {
		return nil, err
	}
	return &achievement, nil
}

func (repository ManageAchievementRepositoryImpl) GetAllArchievement() ([]*archievement.Achievement, error) {
	var achievements []*archievement.Achievement
	if err := repository.DB.GetDB().Order("target_point desc").Find(&achievements).Error; err != nil {
		return nil, err
	}
	return achievements, nil
}

func (repository ManageAchievementRepositoryImpl) GetAchievementById(id int) (*archievement.Achievement, error) {
	var achievement archievement.Achievement
	if err := repository.DB.GetDB().Where("id = ?", id).First(&achievement).Error; err != nil {
		return nil, err
	}
	return &achievement, nil
}

func (repository ManageAchievementRepositoryImpl) UpdateAchievement(achievement *archievement.Achievement, id int) error {
	if err := repository.DB.GetDB().Where("id = ?", id).Updates(achievement).Error; err != nil {
		return err
	}
	return nil
}

func (repository ManageAchievementRepositoryImpl) DeleteAchievement(id int) error {
	if err := repository.DB.GetDB().Delete(&archievement.Achievement{}, id).Error; err != nil {
		return err
	}
	return nil
}
