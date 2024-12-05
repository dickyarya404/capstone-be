package usecase

import (
	"mime/multipart"

	"github.com/sawalreverr/recything/internal/achievements/manage_achievements/dto"
	archievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
)

type ManageAchievementUsecase interface {
	GetAllArchievementUsecase() ([]*archievement.Achievement, error)
	GetAchievementByIdUsecase(id int) (*archievement.Achievement, error)
	UpdateAchievementUsecase(request *dto.UpdateAchievementRequest, badge *multipart.FileHeader, id int) error
	DeleteAchievementUsecase(id int) error
}
