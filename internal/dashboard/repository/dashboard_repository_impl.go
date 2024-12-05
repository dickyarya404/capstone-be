package repository

import (
	art "github.com/sawalreverr/recything/internal/article"
	"github.com/sawalreverr/recything/internal/dashboard/dto"
	"github.com/sawalreverr/recything/internal/database"
	rep "github.com/sawalreverr/recything/internal/report"
	ch "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	usr "github.com/sawalreverr/recything/internal/user"
	vid "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type DashboardRepositoryImpl struct {
	DB database.Database
}

func NewDashboardRepository(db database.Database) DashboardRepository {
	return &DashboardRepositoryImpl{DB: db}
}

// GetTotalArticle implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetTotalArticle() (int, int, error) {
	var totalArticle int64
	var additionContentToday int64

	if err := d.DB.GetDB().Model(&art.Article{}).Count(&totalArticle).Error; err != nil {
		return 0, 0, err
	}
	if err := d.DB.GetDB().Model(&art.Article{}).Where("created_at >= CURRENT_DATE").Count(&additionContentToday).Error; err != nil {
		return 0, 0, err
	}
	return int(totalArticle), int(additionContentToday), nil
}

// GetTotalVideo implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetTotalVideo() (int, int, error) {
	var totalVideo int64
	var additionVideoToday int64

	// Hitung total video
	if err := d.DB.GetDB().Model(&vid.Video{}).Count(&totalVideo).Error; err != nil {
		return 0, 0, err
	}

	// Hitung video yang ditambahkan hari ini
	if err := d.DB.GetDB().Model(&vid.Video{}).Where("created_at >= CURRENT_DATE").Count(&additionVideoToday).Error; err != nil {
		return 0, 0, err
	}

	return int(totalVideo), int(additionVideoToday), nil
}

// GetReportLittering implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetReportLittering() (int, error) {
	var totalLitering int64

	if err := d.DB.GetDB().Model(&rep.Report{}).Where("report_type = ?", "littering").Count(&totalLitering).Error; err != nil {
		return 0, err
	}
	return int(totalLitering), nil
}

// GetReportRubbish implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetReportRubbish() (int, error) {
	var totalRubbish int64

	if err := d.DB.GetDB().Model(&rep.Report{}).Where("report_type = ?", "rubbish").Count(&totalRubbish).Error; err != nil {
		return 0, err
	}
	return int(totalRubbish), nil
}

// GetTotalChange implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetTotalChallenge() (int, int, error) {
	var totalChallenge int64
	var additionChallengeSinceLastWeek int64

	// Hitung total tantangan
	if err := d.DB.GetDB().Model(&ch.TaskChallenge{}).Count(&totalChallenge).Error; err != nil {
		return 0, 0, err
	}

	// Hitung tantangan yang ditambahkan dalam 1 minggu terakhir
	if err := d.DB.GetDB().Model(&ch.TaskChallenge{}).Where("created_at > NOW() - INTERVAL 1 WEEK").Count(&additionChallengeSinceLastWeek).Error; err != nil {
		return 0, 0, err
	}

	return int(totalChallenge), int(additionChallengeSinceLastWeek), nil
}

// GetTotalReport implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetTotalReport() (int, int, error) {
	var totalReport int64
	var additionReportSinceYesterday int64

	if err := d.DB.GetDB().Model(&rep.Report{}).Count(&totalReport).Error; err != nil {
		return 0, 0, err
	}

	if err := d.DB.GetDB().Model(&rep.Report{}).Where("created_at > now() - interval 1 day").Count(&additionReportSinceYesterday).Error; err != nil {
		return 0, 0, err
	}

	return int(totalReport), int(additionReportSinceYesterday), nil
}

// GetTotalUser implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetTotalUser() (int, int, error) {
	var totalUser int64
	var additionUserSinceYesterday int64

	// Hitung total user
	if err := d.DB.GetDB().Model(&usr.User{}).Count(&totalUser).Error; err != nil {
		return 0, 0, err
	}

	// Hitung user yang ditambahkan sejak kemarin
	if err := d.DB.GetDB().Model(&usr.User{}).Where("created_at >= NOW() - INTERVAL 1 DAY").Count(&additionUserSinceYesterday).Error; err != nil {
		return 0, 0, err
	}

	return int(totalUser), int(additionUserSinceYesterday), nil
}

// GetUserClassic implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetUserClassic() (int, error) {
	classic := "https://res.cloudinary.com/dymhvau8n/image/upload/v1718189121/user_badge/htaemsjtlhfof7ww01ss.png"

	var totalUser int64
	if err := d.DB.GetDB().Model(&usr.User{}).Where("badge = ?", classic).Count(&totalUser).Error; err != nil {
		return 0, err
	}
	return int(totalUser), nil
}

// GetUserGold implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetUserGold() (int, error) {
	gold := "https://res.cloudinary.com/dymhvau8n/image/upload/v1718189184/user_badge/jshs1s2fwevahgtvjkgj.png"

	var totalUser int64
	if err := d.DB.GetDB().Model(&usr.User{}).Where("badge = ?", gold).Count(&totalUser).Error; err != nil {
		return 0, err
	}
	return int(totalUser), nil
}

func (d *DashboardRepositoryImpl) GetUserPlatinum() (int, error) {
	platinum := "https://res.cloudinary.com/dymhvau8n/image/upload/v1718188250/user_badge/icureiapdvtzyu5b99zu.png"

	var totalUser int64
	if err := d.DB.GetDB().Model(&usr.User{}).Where("badge = ?", platinum).Count(&totalUser).Error; err != nil {
		return 0, err
	}
	return int(totalUser), nil
}

// GetUserSilver implements DashboardRepository.
func (d *DashboardRepositoryImpl) GetUserSilver() (int, error) {
	silver := "https://res.cloudinary.com/dymhvau8n/image/upload/v1718189221/user_badge/oespnjdgoynkairlutbk.png"

	var totalUser int64
	if err := d.DB.GetDB().Model(&usr.User{}).Where("badge = ?", silver).Count(&totalUser).Error; err != nil {
		return 0, err
	}
	return int(totalUser), nil
}

func (d *DashboardRepositoryImpl) GetMonthlyReport(year int, reportType string) ([]dto.MonthlyReportStats, error) {
	var stats []dto.MonthlyReportStats

	for month := 1; month <= 12; month++ {
		var dailyStats []dto.DailyReportStats

		// Tentukan jumlah hari dalam bulan tersebut
		var daysInMonth int
		if month == 2 {
			if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
				daysInMonth = 29 // Tahun kabisat
			} else {
				daysInMonth = 28
			}
		} else if month == 4 || month == 6 || month == 9 || month == 11 {
			daysInMonth = 30
		} else {
			daysInMonth = 31
		}

		// Inisialisasi hasil dengan hari dari 1 hingga jumlah hari dalam bulan tersebut
		for day := 1; day <= daysInMonth; day++ {
			dailyStats = append(dailyStats, dto.DailyReportStats{
				Day: int64(day),
			})
		}

		// Query untuk menghitung jumlah laporan per hari dalam bulan tertentu
		query := `
            SELECT
                DAY(created_at) AS day,
                COUNT(*) AS total_reports
            FROM
                reports
            WHERE
                YEAR(created_at) = ? AND MONTH(created_at) = ? AND report_type = ?
            GROUP BY
                DAY(created_at)
            ORDER BY
                DAY(created_at)
        `

		var queryResults []dto.DailyReportStats
		if err := d.DB.GetDB().Raw(query, year, month, reportType).Scan(&queryResults).Error; err != nil {
			return nil, err
		}

		// Isi hasil dengan data yang diambil dari query
		for _, queryResult := range queryResults {
			dailyStats[queryResult.Day-1].TotalReports = queryResult.TotalReports
		}

		// Hitung total laporan dalam bulan tersebut
		var totalReports int64
		for _, stat := range dailyStats {
			totalReports += stat.TotalReports
		}

		// Tambahkan statistik bulanan ke hasil akhir
		stats = append(stats, dto.MonthlyReportStats{
			Month:        getMonthName(month),
			DailyStats:   dailyStats,
			TotalReports: totalReports,
		})
	}

	return stats, nil
}

func getMonthName(month int) string {
	months := []string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}
	if month >= 1 && month <= 12 {
		return months[month-1]
	}
	return ""
}

func (d *DashboardRepositoryImpl) GetReportByCity() ([]dto.DataReportByCity, error) {
	var result []dto.DataReportByCity

	if err := d.DB.GetDB().Model(&rep.Report{}).
		Select("city, COUNT(*) as total_report").
		Group("city").
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil

}

func (d *DashboardRepositoryImpl) GetUserByGender() ([]dto.DataUserByGender, error) {
	var genderStats []dto.DataUserByGender
	var totalUsers int64

	// Hitung total pengguna
	if err := d.DB.GetDB().Model(&usr.User{}).Count(&totalUsers).Error; err != nil {
		return nil, err
	}

	// Hitung total pengguna per gender
	type Result struct {
		Gender    string
		TotalUser int64
	}

	var results []Result
	if err := d.DB.GetDB().Model(&usr.User{}).Select("gender, COUNT(*) as total_user").Group("gender").Scan(&results).Error; err != nil {
		return nil, err
	}

	// Hitung persentase
	for _, result := range results {
		percentage := float64(result.TotalUser) * 100 / float64(totalUsers)
		genderStats = append(genderStats, dto.DataUserByGender{
			Gender:     result.Gender,
			TotalUser:  result.TotalUser,
			Percentage: percentage,
		})
	}

	return genderStats, nil
}

func (d *DashboardRepositoryImpl) GetDataReportByWasteType(reportType string, wasteTypes []string) ([]dto.DataReportByWasteType, error) {
	var wasteStats []dto.DataReportByWasteType
	var totalReports int64

	// Hitung total laporan untuk reportType tertentu
	if err := d.DB.GetDB().Model(&rep.Report{}).Where("report_type = ?", reportType).Count(&totalReports).Error; err != nil {
		return nil, err
	}

	// Hitung jumlah laporan per waste type
	type Result struct {
		WasteType    string
		TotalReports int64
	}

	var results []Result
	if err := d.DB.GetDB().Model(&rep.Report{}).
		Select("waste_type, COUNT(*) as total_reports").
		Where("report_type = ? AND waste_type IN ?", reportType, wasteTypes).
		Group("waste_type").
		Scan(&results).Error; err != nil {
		return nil, err
	}

	// Hitung persentase
	for _, result := range results {
		percentage := float64(result.TotalReports) * 100 / float64(totalReports)
		wasteStats = append(wasteStats, dto.DataReportByWasteType{
			ReportType:   reportType,
			WasteType:    result.WasteType,
			TotalReports: result.TotalReports,
			Percentage:   percentage,
		})
	}

	return wasteStats, nil
}
