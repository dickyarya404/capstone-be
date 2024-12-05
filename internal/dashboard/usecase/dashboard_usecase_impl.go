package usecase

import (
	"github.com/sawalreverr/recything/internal/dashboard/dto"
	"github.com/sawalreverr/recything/internal/dashboard/repository"
)

type DashboardUsecaseImpl struct {
	dashboardRepository repository.DashboardRepository
}

func NewDashboardUsecase(dashboardRepository repository.DashboardRepository) DashboardUsecase {
	return &DashboardUsecaseImpl{dashboardRepository: dashboardRepository}
}

func (usecase *DashboardUsecaseImpl) GetDashboardUsecase() (*dto.DashboardResponse, error) {
	totalUser, additionUserSinceYesterday, err := usecase.dashboardRepository.GetTotalUser()
	if err != nil {
		return nil, err
	}
	totalReport, additionReportSinceYesterday, err := usecase.dashboardRepository.GetTotalReport()
	if err != nil {
		return nil, err
	}
	totalChallenge, additionChallengeSinceLastWeek, err := usecase.dashboardRepository.GetTotalChallenge()
	if err != nil {
		return nil, err
	}
	totalVideo, additionVideoToday, err := usecase.dashboardRepository.GetTotalVideo()
	if err != nil {
		return nil, err
	}
	totalArticle, additionArticleToday, err := usecase.dashboardRepository.GetTotalArticle()
	if err != nil {
		return nil, err
	}

	totalUserClassic, err := usecase.dashboardRepository.GetUserClassic()
	if err != nil {
		return nil, err
	}
	totalUserSilver, err := usecase.dashboardRepository.GetUserSilver()
	if err != nil {
		return nil, err
	}
	totalUserGold, err := usecase.dashboardRepository.GetUserGold()
	if err != nil {
		return nil, err
	}
	totalUserPlatinum, err := usecase.dashboardRepository.GetUserPlatinum()
	if err != nil {
		return nil, err
	}
	totalLittering, err := usecase.dashboardRepository.GetReportLittering()
	if err != nil {
		return nil, err
	}
	totalRubbish, err := usecase.dashboardRepository.GetReportRubbish()
	if err != nil {
		return nil, err
	}

	reportLittering, err := usecase.dashboardRepository.GetMonthlyReport(2024, "littering")
	if err != nil {
		return nil, err
	}
	reportRubbish, err := usecase.dashboardRepository.GetMonthlyReport(2024, "rubbish")
	if err != nil {
		return nil, err
	}

	reportByCity, err := usecase.dashboardRepository.GetReportByCity()
	if err != nil {
		return nil, err
	}

	userByGender, err := usecase.dashboardRepository.GetUserByGender()
	if err != nil {
		return nil, err
	}

	dataReportByWasteRubbish, err := usecase.dashboardRepository.GetDataReportByWasteType("rubbish", []string{"sampah basah", "sampah kering", "sampah basah,sampah kering"})
	if err != nil {
		return nil, err
	}

	dataReportByWasteLittering, err := usecase.dashboardRepository.GetDataReportByWasteType("littering", []string{"organik", "anorganik", "berbahaya"})
	if err != nil {
		return nil, err
	}

	user := dto.TotalUser{
		TotalUser:                  totalUser,
		AdditionUserSinceYesterday: additionUserSinceYesterday,
	}

	userAchievement := dto.UserAchievement{
		TotalUser: totalUser,
		Classic:   totalUserClassic,
		Silver:    totalUserSilver,
		Gold:      totalUserGold,
		Platinum:  totalUserPlatinum,
	}

	report := dto.TotalReport{
		TotalReport:                  totalReport,
		AdditionReportSinceYesterday: additionReportSinceYesterday,
	}

	dataReportStats := dto.DataReportStatistic{
		ReportLittering: reportLittering,
		ReportRubbish:   reportRubbish,
	}

	totalContent := totalArticle + totalVideo

	challenge := dto.TotalChallenge{
		TotalChallenge:                 totalChallenge,
		AdditionChallengeSinceLastWeek: additionChallengeSinceLastWeek,
	}

	content := dto.TotalContent{
		TotalContent:         totalContent,
		AdditionContentToday: additionVideoToday + additionArticleToday,
	}

	return &dto.DashboardResponse{
		User:                           user,
		Report:                         report,
		Challenge:                      challenge,
		Content:                        content,
		UserAchievement:                userAchievement,
		TotalLittering:                 totalLittering,
		TotalRubbish:                   totalRubbish,
		DataReportStatistic:            dataReportStats,
		DataReportByCity:               reportByCity,
		DataUserByGender:               userByGender,
		DataReportByWasteTypeRubbish:   dataReportByWasteRubbish,
		DataReportByWasteTypeLittering: dataReportByWasteLittering,
	}, nil
}
