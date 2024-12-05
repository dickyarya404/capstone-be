package usecase

import (
	"mime/multipart"
	"strings"

	"github.com/sawalreverr/recything/internal/achievements/manage_achievements/dto"
	archievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
	"github.com/sawalreverr/recything/internal/achievements/manage_achievements/repository"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type ManageAchievementUsecaseImpl struct {
	repository repository.ManageAchievementRepository
}

func NewManageAchievementUsecase(repository repository.ManageAchievementRepository) *ManageAchievementUsecaseImpl {
	return &ManageAchievementUsecaseImpl{repository: repository}
}

func (repository ManageAchievementUsecaseImpl) GetAllArchievementUsecase() ([]*archievement.Achievement, error) {
	achievements, err := repository.repository.GetAllArchievement()
	if err != nil {
		return nil, err
	}
	return achievements, nil
}

func (repository ManageAchievementUsecaseImpl) GetAchievementByIdUsecase(id int) (*archievement.Achievement, error) {
	achievement, err := repository.repository.GetAchievementById(id)
	if err != nil {
		return nil, pkg.ErrAchievementNotFound
	}

	return achievement, nil
}

func (repository ManageAchievementUsecaseImpl) UpdateAchievementUsecase(request *dto.UpdateAchievementRequest, badge *multipart.FileHeader, id int) error {
	achievement, err := repository.repository.GetAchievementById(id)
	if err != nil {
		return pkg.ErrAchievementNotFound
	}
	var urlBadge string
	if badge != nil {
		if badge.Size > 2*1024*1024 {
			return pkg.ErrFileTooLarge
		}
		if !strings.HasPrefix(badge.Header.Get("Content-Type"), "image") {
			return pkg.ErrInvalidFileType
		}
		src, errOpen := badge.Open()
		if errOpen != nil {
			return pkg.ErrOpenFile
		}

		urlBadgeUpload, errUpload := helper.UploadToCloudinary(src, "achievement_badge")
		if errUpload != nil {
			return errUpload
		}
		urlBadge = urlBadgeUpload
		defer src.Close()
	}

	if request.Level != "" {
		achievement.Level = request.Level
	}
	if request.TargetPoint != 0 {
		achievement.TargetPoint = request.TargetPoint
	}
	if urlBadge != "" {
		achievement.BadgeUrl = urlBadge
	}

	if err := repository.repository.UpdateAchievement(achievement, id); err != nil {
		return err
	}
	return nil
}

func (repository ManageAchievementUsecaseImpl) DeleteAchievementUsecase(id int) error {
	if _, err := repository.repository.GetAchievementById(id); err != nil {
		return pkg.ErrAchievementNotFound
	}
	if err := repository.repository.DeleteAchievement(id); err != nil {
		return err
	}
	return nil
}
