package report

import (
	"time"

	"github.com/google/uuid"
	"github.com/sawalreverr/recything/internal/helper"
	rpt "github.com/sawalreverr/recything/internal/report"
	user "github.com/sawalreverr/recything/internal/user"
	"github.com/sawalreverr/recything/pkg"
)

type reportUsecase struct {
	reportRepository rpt.ReportRepository
	userRepository   user.UserRepository
}

func NewReportUsecase(reportRepo rpt.ReportRepository, userRepo user.UserRepository) rpt.ReportUsecase {
	return &reportUsecase{reportRepository: reportRepo, userRepository: userRepo}
}

func (uc *reportUsecase) CreateReport(report rpt.ReportInput, authorID string, imageURLs []string) (*rpt.ReportDetail, error) {
	lastID, _ := uc.reportRepository.FindLastID()
	newID := helper.GenerateCustomID(lastID, "RPT")

	newReport := rpt.Report{
		ID:          newID,
		AuthorID:    authorID,
		ReportType:  report.ReportType,
		Title:       report.Title,
		Description: report.Description,
		WasteType:   report.WasteType,
		Latitude:    report.Latitude,
		Longitude:   report.Longitude,
		Address:     report.Address,
		City:        report.City,
		Province:    report.Province,
	}

	createdReport, err := uc.reportRepository.Create(newReport)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	for _, materialType := range report.WasteMaterials {
		material, err := uc.reportRepository.FindWasteMaterialByType(materialType)
		if err != nil {
			_ = uc.reportRepository.Delete(createdReport.ID)
			return nil, err
		}

		reportMaterial := rpt.ReportWasteMaterial{
			ID:              uuid.New(),
			ReportID:        createdReport.ID,
			WasteMaterialID: material.ID,
		}

		if _, err := uc.reportRepository.AddReportMaterial(reportMaterial); err != nil {
			_ = uc.reportRepository.Delete(createdReport.ID)
			return nil, err
		}
	}

	for _, url := range imageURLs {
		reportImage := rpt.ReportImage{
			ID:       uuid.New(),
			ReportID: createdReport.ID,
			ImageURL: url,
		}

		if _, err := uc.reportRepository.AddImage(reportImage); err != nil {
			_ = uc.reportRepository.Delete(createdReport.ID)
			_ = uc.reportRepository.DeleteAllReportMaterial(createdReport.ID)
			return nil, err
		}
	}

	images, err := uc.reportRepository.FindAllImage(createdReport.ID)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	materials, err := uc.reportRepository.FindAllReportMaterial(createdReport.ID)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	userFound, _ := uc.userRepository.FindByID(authorID)
	author := rpt.UserDetail{
		ID:       userFound.ID,
		Name:     userFound.Name,
		ImageURL: userFound.PictureURL,
	}

	reportDetail := rpt.ReportDetail{
		ID:             createdReport.ID,
		Author:         author,
		ReportType:     createdReport.ReportType,
		Title:          createdReport.Title,
		Description:    createdReport.Description,
		WasteType:      createdReport.WasteType,
		Latitude:       createdReport.Latitude,
		Longitude:      createdReport.Longitude,
		Address:        createdReport.Address,
		City:           createdReport.City,
		Province:       createdReport.Province,
		Status:         createdReport.Status,
		Reason:         createdReport.Reason,
		CreatedAt:      createdReport.CreatedAt,
		WasteMaterials: *materials,
		ReportImages:   *images,
	}

	return &reportDetail, nil
}

func (uc *reportUsecase) FindHistoryUserReports(authorID string) (*[]rpt.ReportDetail, error) {
	var reportDetails []rpt.ReportDetail
	reports, err := uc.reportRepository.FindAllReportsByUser(authorID, 10)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	for _, report := range *reports {
		images, err := uc.reportRepository.FindAllImage(report.ID)
		if err != nil {
			return nil, pkg.ErrStatusInternalError
		}

		materials, err := uc.reportRepository.FindAllReportMaterial(report.ID)
		if err != nil {
			return nil, pkg.ErrStatusInternalError
		}

		userFound, _ := uc.userRepository.FindByID(authorID)
		author := rpt.UserDetail{
			ID:       userFound.ID,
			Name:     userFound.Name,
			ImageURL: userFound.PictureURL,
		}

		reportDetail := rpt.ReportDetail{
			ID:             report.ID,
			Author:         author,
			ReportType:     report.ReportType,
			Title:          report.Title,
			Description:    report.Description,
			WasteType:      report.WasteType,
			Latitude:       report.Latitude,
			Longitude:      report.Longitude,
			Address:        report.Address,
			City:           report.City,
			Province:       report.Province,
			Status:         report.Status,
			Reason:         report.Reason,
			CreatedAt:      report.CreatedAt,
			WasteMaterials: *materials,
			ReportImages:   *images,
		}

		reportDetails = append(reportDetails, reportDetail)
	}

	return &reportDetails, nil
}

func (uc *reportUsecase) UpdateStatusReport(report rpt.UpdateStatus, reportID string) error {
	reportFound, err := uc.reportRepository.FindByID(reportID)
	if err != nil {
		return pkg.ErrReportNotFound
	}

	reportFound.Status = report.Status

	if reportFound.Status == "reject" {
		reportFound.Reason = report.Reason
	}

	if err := uc.reportRepository.Update(*reportFound); err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}

func (uc *reportUsecase) FindAllReports(page, limit int, reportType, status string, date time.Time) (*[]rpt.ReportDetail, int64, error) {
	var reportDetails []rpt.ReportDetail
	reports, total, err := uc.reportRepository.FindAll(page, limit, reportType, status, date)
	if err != nil {
		return nil, 0, pkg.ErrStatusInternalError
	}

	for _, report := range *reports {
		images, err := uc.reportRepository.FindAllImage(report.ID)
		if err != nil {
			return nil, 0, pkg.ErrStatusInternalError
		}

		materials, err := uc.reportRepository.FindAllReportMaterial(report.ID)
		if err != nil {
			return nil, 0, pkg.ErrStatusInternalError
		}

		userFound, _ := uc.userRepository.FindByID(report.AuthorID)
		author := rpt.UserDetail{
			ID:       userFound.ID,
			Name:     userFound.Name,
			ImageURL: userFound.PictureURL,
		}

		reportDetail := rpt.ReportDetail{
			ID:             report.ID,
			Author:         author,
			ReportType:     report.ReportType,
			Title:          report.Title,
			Description:    report.Description,
			WasteType:      report.WasteType,
			Latitude:       report.Latitude,
			Longitude:      report.Longitude,
			Address:        report.Address,
			City:           report.City,
			Province:       report.Province,
			Status:         report.Status,
			Reason:         report.Reason,
			CreatedAt:      report.CreatedAt,
			WasteMaterials: *materials,
			ReportImages:   *images,
		}

		reportDetails = append(reportDetails, reportDetail)
	}

	return &reportDetails, total, nil
}
