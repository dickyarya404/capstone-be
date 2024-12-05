package usecase

import (
	"github.com/sawalreverr/recything/internal/homepage/dto"
)

type HomepageUsecase interface {
	GetHomepageUsecase(userId string) (*dto.HomepageResponse, error)
}
