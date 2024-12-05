package usecase

import "github.com/sawalreverr/recything/internal/dashboard/dto"

type DashboardUsecase interface {
	GetDashboardUsecase() (*dto.DashboardResponse, error)
}
