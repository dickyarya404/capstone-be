package report

import (
	"time"

	"github.com/sawalreverr/recything/internal/database"
	rpt "github.com/sawalreverr/recything/internal/report"
)

type reportRepository struct {
	DB database.Database
}

func NewReportRepository(db database.Database) rpt.ReportRepository {
	return &reportRepository{DB: db}
}

// Report
func (r *reportRepository) Create(report rpt.Report) (*rpt.Report, error) {
	if err := r.DB.GetDB().Create(&report).Error; err != nil {
		return nil, err
	}

	return &report, nil
}

func (r *reportRepository) FindByID(reportID string) (*rpt.Report, error) {
	var report rpt.Report
	if err := r.DB.GetDB().Where("id = ?", reportID).First(&report).Error; err != nil {
		return nil, err
	}

	return &report, nil
}

func (r *reportRepository) FindLastID() (string, error) {
	var report rpt.Report
	if err := r.DB.GetDB().Unscoped().Order("id DESC").First(&report).Error; err != nil {
		return "RPT0000", err
	}

	return report.ID, nil
}

func (r *reportRepository) Update(report rpt.Report) error {
	if err := r.DB.GetDB().Save(&report).Error; err != nil {
		return err
	}

	return nil
}

func (r *reportRepository) Delete(reportID string) error {
	var report rpt.Report
	if err := r.DB.GetDB().Where("id = ?", reportID).Delete(&report).Error; err != nil {
		return err
	}

	return nil
}

func (r *reportRepository) FindAll(page, limit int, reportType, status string, date time.Time) (*[]rpt.Report, int64, error) {
	var reports []rpt.Report
	var total int64

	db := r.DB.GetDB().Model(&rpt.Report{})

	if reportType != "" {
		db = db.Where("report_type = ?", reportType)
	}
	if status != "" {
		db = db.Where("status = ?", status)
	}
	if !date.IsZero() {
		dateStr := date.Format("2006-01-02")
		db = db.Where("DATE(created_at) = ?", dateStr)
	}

	db.Count(&total)
	offset := (page - 1) * limit
	if err := db.Offset(offset).Limit(limit).Find(&reports).Error; err != nil {
		return nil, 0, err
	}

	return &reports, total, nil
}

func (r *reportRepository) FindAllReportsByUser(userID string, limit int) (*[]rpt.Report, error) {
	var reports []rpt.Report
	if err := r.DB.GetDB().Where("author_id = ?", userID).Order("created_at desc").Limit(10).Find(&reports).Error; err != nil {
		return nil, err
	}

	return &reports, nil
}

// Report Image
func (r *reportRepository) AddImage(image rpt.ReportImage) (*rpt.ReportImage, error) {
	if err := r.DB.GetDB().Create(&image).Error; err != nil {
		return nil, err
	}

	return &image, nil
}

func (r *reportRepository) DeleteImage(imageID string, reportID string) error {
	var reportImage rpt.ReportImage
	if err := r.DB.GetDB().Where("id = ? AND report_id = ?", imageID, reportID).Delete(&reportImage).Error; err != nil {
		return err
	}

	return nil
}

func (r *reportRepository) DeleteAllImage(reportID string) error {
	var reportImage rpt.ReportImage
	if err := r.DB.GetDB().Where("report_id = ?", reportID).Delete(&reportImage).Error; err != nil {
		return err
	}

	return nil
}

func (r *reportRepository) FindAllImage(reportID string) (*[]string, error) {
	var reportImages []rpt.ReportImage
	var imageURLs []string

	if err := r.DB.GetDB().Where("report_id = ?", reportID).Find(&reportImages).Error; err != nil {
		return nil, err
	}

	for _, report := range reportImages {
		imageURLs = append(imageURLs, report.ImageURL)
	}

	return &imageURLs, nil
}

// Report Waste Materials
func (r *reportRepository) AddReportMaterial(material rpt.ReportWasteMaterial) (*rpt.ReportWasteMaterial, error) {
	if err := r.DB.GetDB().Create(&material).Error; err != nil {
		return nil, err
	}

	return &material, nil
}

func (r *reportRepository) DeleteAllReportMaterial(reportID string) error {
	var reportMaterial rpt.ReportWasteMaterial
	if err := r.DB.GetDB().Where("report_id = ?", reportID).Delete(&reportMaterial).Error; err != nil {
		return err
	}

	return nil
}

func (r *reportRepository) FindAllReportMaterial(reportID string) (*[]rpt.WasteMaterial, error) {
	var reportMaterials []rpt.ReportWasteMaterial
	var wasteMaterials []rpt.WasteMaterial

	if err := r.DB.GetDB().Where("report_id = ?", reportID).Find(&reportMaterials).Error; err != nil {
		return nil, err
	}

	for _, report := range reportMaterials {
		wasteMaterial, _ := r.FindWasteMaterialByID(report.WasteMaterialID)
		wasteMaterials = append(wasteMaterials, *wasteMaterial)
	}

	return &wasteMaterials, nil
}

func (r *reportRepository) FindWasteMaterialByID(materialID string) (*rpt.WasteMaterial, error) {
	var wasteMaterial rpt.WasteMaterial
	if err := r.DB.GetDB().Where("id = ?", materialID).First(&wasteMaterial).Error; err != nil {
		return nil, err
	}

	return &wasteMaterial, nil
}

func (r *reportRepository) FindWasteMaterialByType(materialType string) (*rpt.WasteMaterial, error) {
	var wasteMaterial rpt.WasteMaterial
	if err := r.DB.GetDB().Where("type = ?", materialType).First(&wasteMaterial).Error; err != nil {
		return nil, err
	}

	return &wasteMaterial, nil
}
