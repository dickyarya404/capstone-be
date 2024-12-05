package repository

import "github.com/sawalreverr/recything/internal/dashboard/dto"

type DashboardRepository interface {
	GetTotalUser() (int, int, error)
	GetTotalReport() (int, int, error)
	GetTotalChallenge() (int, int, error)
	GetTotalVideo() (int, int, error)
	GetTotalArticle() (int, int, error)
	GetUserClassic() (int, error)
	GetUserSilver() (int, error)
	GetUserGold() (int, error)
	GetUserPlatinum() (int, error)
	GetReportLittering() (int, error)
	GetReportRubbish() (int, error)
	GetMonthlyReport(year int, reportType string) ([]dto.MonthlyReportStats, error)
	GetReportByCity() ([]dto.DataReportByCity, error)
	GetUserByGender() ([]dto.DataUserByGender, error)
	GetDataReportByWasteType(reportType string, wasteTypes []string) ([]dto.DataReportByWasteType, error)
}
