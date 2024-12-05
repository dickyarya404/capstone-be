package database

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	rpt "github.com/sawalreverr/recything/internal/report"
)

func (m *mysqlDatabase) InitReport() {
	if err := m.GetDB().Migrator().DropTable(&rpt.ReportImage{}); err != nil {
		return
	}
	if err := m.GetDB().Migrator().DropTable(&rpt.ReportWasteMaterial{}); err != nil {
		return
	}
	if err := m.GetDB().Migrator().DropTable(&rpt.Report{}); err != nil {
		return
	}

	if err := m.GetDB().AutoMigrate(&rpt.ReportImage{}); err != nil {
		return
	}
	if err := m.GetDB().AutoMigrate(&rpt.ReportWasteMaterial{}); err != nil {
		return
	}
	if err := m.GetDB().AutoMigrate(&rpt.Report{}); err != nil {
		return
	}

	reports, reportWasteMaterials, reportImages := generateReports()

	for _, report := range reports {
		m.GetDB().FirstOrCreate(&report, report)
	}

	for _, reportWasteMaterial := range reportWasteMaterials {
		m.GetDB().FirstOrCreate(&reportWasteMaterial, reportWasteMaterial)
	}

	for _, reportImage := range reportImages {
		m.GetDB().FirstOrCreate(&reportImage, reportImage)
	}
}

var wasteMaterials = []struct {
	ID   string
	Type string
}{
	{ID: "MTR01", Type: "plastik"},
	{ID: "MTR02", Type: "kaca"},
	{ID: "MTR03", Type: "kayu"},
	{ID: "MTR04", Type: "kertas"},
	{ID: "MTR05", Type: "baterai"},
	{ID: "MTR06", Type: "besi"},
	{ID: "MTR07", Type: "limbah berbahaya"},
	{ID: "MTR08", Type: "limbah beracun"},
	{ID: "MTR09", Type: "sisa makanan"},
	{ID: "MTR10", Type: "tak terdeteksi"},
}

var addressesReport = []struct {
	Address  string
	City     string
	Province string
}{
	{"Jalan Jendral Sudirman", "Jakarta", "DKI Jakarta"},
	{"Jalan MH Thamrin", "Jakarta", "DKI Jakarta"},
	{"Jalan Gatot Subroto", "Jakarta", "DKI Jakarta"},
	{"Jalan Merdeka Selatan", "Jakarta", "DKI Jakarta"},
	{"Jalan Merdeka Utara", "Jakarta", "DKI Jakarta"},
	{"Jalan Kuningan", "Jakarta", "DKI Jakarta"},
	{"Jalan Braga", "Bandung", "Jawa Barat"},
	{"Jalan Asia Afrika", "Bandung", "Jawa Barat"},
	{"Jalan Dago", "Bandung", "Jawa Barat"},
	{"Jalan Riau", "Bandung", "Jawa Barat"},
	{"Jalan Malioboro", "Yogyakarta", "DI Yogyakarta"},
	{"Jalan Solo", "Yogyakarta", "DI Yogyakarta"},
	{"Jalan Diponegoro", "Surabaya", "Jawa Timur"},
	{"Jalan Darmo", "Surabaya", "Jawa Timur"},
	{"Jalan Imam Bonjol", "Semarang", "Jawa Tengah"},
	{"Jalan Pandanaran", "Semarang", "Jawa Tengah"},
	{"Jalan Adi Sucipto", "Solo", "Jawa Tengah"},
	{"Jalan Slamet Riyadi", "Solo", "Jawa Tengah"},
	{"Jalan Gajah Mada", "Medan", "Sumatera Utara"},
	{"Jalan Sisingamangaraja", "Medan", "Sumatera Utara"},
	{"Jalan Pahlawan", "Denpasar", "Bali"},
	{"Jalan Hayam Wuruk", "Denpasar", "Bali"},
	{"Jalan Pattimura", "Pontianak", "Kalimantan Barat"},
	{"Jalan Ahmad Yani", "Pontianak", "Kalimantan Barat"},
	{"Jalan Jendral Sudirman", "Makassar", "Sulawesi Selatan"},
	{"Jalan Dr. Sam Ratulangi", "Makassar", "Sulawesi Selatan"},
	{"Jalan Diponegoro", "Balikpapan", "Kalimantan Timur"},
	{"Jalan Jenderal Sudirman", "Balikpapan", "Kalimantan Timur"},
	{"Jalan Pemuda", "Manado", "Sulawesi Utara"},
	{"Jalan Sam Ratulangi", "Manado", "Sulawesi Utara"},
}

func generateReports() ([]rpt.Report, []rpt.ReportWasteMaterial, []rpt.ReportImage) {
	gofakeit.Seed(0)

	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 6, 30, 23, 59, 59, 999, time.UTC)

	reports := make([]rpt.Report, 100)
	reportWasteMaterials := make([]rpt.ReportWasteMaterial, 0)
	reportImages := make([]rpt.ReportImage, 0)

	for i := 0; i < 100; i++ {
		reportID := fmt.Sprintf("RPT%04d", i+1)
		address := addressesReport[rand.Intn(len(addressesReport))]
		reportType := gofakeit.RandomString([]string{"littering", "rubbish"})
		var wasteType string
		if reportType == "rubbish" {
			wasteType = gofakeit.RandomString([]string{"sampah basah", "sampah kering", "sampah basah,sampah kering"})
		} else {
			wasteType = gofakeit.RandomString([]string{"organik", "anorganik", "berbahaya"})
		}
		status := gofakeit.RandomString([]string{"need review", "approve", "reject"})
		reason := ""
		if status == "reject" {
			reason = gofakeit.Sentence(5)
		}

		report := rpt.Report{
			ID:          reportID,
			AuthorID:    fmt.Sprintf("USR%04d", gofakeit.Number(1, 50)),
			ReportType:  reportType,
			Title:       gofakeit.Sentence(6),
			Description: gofakeit.Paragraph(1, 2, 3, ""),
			WasteType:   wasteType,
			Latitude:    gofakeit.Latitude(),
			Longitude:   gofakeit.Longitude(),
			Address:     address.Address,
			City:        address.City,
			Province:    address.Province,
			Status:      status,
			Reason:      reason,
			CreatedAt:   randomDate(startDate, endDate),
		}

		if reportType == "rubbish" {
			wasteMaterialCount := rand.Intn(3) + 1
			for j := 0; j < wasteMaterialCount; j++ {
				wasteMaterial := rpt.ReportWasteMaterial{
					ID:              uuid.New(),
					ReportID:        reportID,
					WasteMaterialID: wasteMaterials[rand.Intn(len(wasteMaterials))].ID,
				}
				reportWasteMaterials = append(reportWasteMaterials, wasteMaterial)
			}
		}

		imageCount := rand.Intn(3) + 3
		for j := 0; j < imageCount; j++ {
			reportImage := rpt.ReportImage{
				ID:       uuid.New(),
				ReportID: reportID,
				ImageURL: gofakeit.ImageURL(640, 480),
			}
			reportImages = append(reportImages, reportImage)
		}

		reports[i] = report
	}

	return reports, reportWasteMaterials, reportImages
}
