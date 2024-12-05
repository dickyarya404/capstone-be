package database

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	art "github.com/sawalreverr/recything/internal/article"
)

func (m *mysqlDatabase) InitArticle() {
	if err := m.GetDB().Migrator().DropTable(&art.ArticleCategories{}); err != nil {
		return
	}
	if err := m.GetDB().Migrator().DropTable(&art.ArticleSection{}); err != nil {
		return
	}
	if err := m.GetDB().Migrator().DropTable(&art.Article{}); err != nil {
		return
	}

	if err := m.GetDB().AutoMigrate(&art.ArticleCategories{}); err != nil {
		return
	}
	if err := m.GetDB().AutoMigrate(&art.ArticleSection{}); err != nil {
		return
	}
	if err := m.GetDB().AutoMigrate(&art.Article{}); err != nil {
		return
	}

	articles, articleSections, articleCategories := generateArticle()

	for _, article := range articles {
		m.GetDB().FirstOrCreate(&article, article)
	}

	for _, articleSection := range articleSections {
		m.GetDB().FirstOrCreate(&articleSection, articleSection)
	}

	for _, articleCategory := range articleCategories {
		m.GetDB().FirstOrCreate(&articleCategory, articleCategory)
	}

	log.Println("Article data added!")
}

var wasteCategories = []struct {
	ID   uint
	Name string
}{
	{ID: 1, Name: "plastik"},
	{ID: 2, Name: "besi"},
	{ID: 3, Name: "kaca"},
	{ID: 4, Name: "organik"},
	{ID: 5, Name: "kayu"},
	{ID: 6, Name: "kertas"},
	{ID: 7, Name: "baterai"},
	{ID: 8, Name: "kaleng"},
	{ID: 9, Name: "elektronik"},
	{ID: 10, Name: "tekstil"},
	{ID: 11, Name: "minyak"},
	{ID: 12, Name: "bola lampu"},
	{ID: 13, Name: "berbahaya"},
}

var contentCategories = []struct {
	ID   uint
	Name string
}{
	{ID: 1, Name: "tips"},
	{ID: 2, Name: "daur ulang"},
	{ID: 3, Name: "tutorial"},
	{ID: 4, Name: "edukasi"},
	{ID: 5, Name: "kampanye"},
}

func generateArticle() ([]art.Article, []art.ArticleSection, []art.ArticleCategories) {
	gofakeit.Seed(0)

	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 6, 30, 23, 59, 59, 999, time.UTC)

	articles := make([]art.Article, 50)
	articleSections := make([]art.ArticleSection, 0)
	articleCategories := make([]art.ArticleCategories, 0)

	for i := 0; i < 50; i++ {
		articleID := fmt.Sprintf("ART%04d", i+1)
		article := art.Article{
			ID:           articleID,
			Title:        gofakeit.Sentence(6),
			Description:  gofakeit.Paragraph(1, 2, 3, ""),
			ThumbnailURL: gofakeit.ImageURL(640, 480),
			AuthorID:     "AD0001",
			CreatedAt:    randomDate(startDate, endDate),
		}

		sectionCount := rand.Intn(4) + 2
		for j := 0; j < sectionCount; j++ {
			section := art.ArticleSection{
				ID:          uint(len(articleSections) + 1),
				ArticleID:   articleID,
				Title:       gofakeit.Sentence(6),
				Description: gofakeit.Paragraph(1, 2, 3, ""),
				ImageURL:    gofakeit.ImageURL(640, 480),
				CreatedAt:   randomDate(startDate, endDate),
			}
			articleSections = append(articleSections, section)
		}

		categoryCount := rand.Intn(3) + 1
		for j := 0; j < categoryCount; j++ {
			category := art.ArticleCategories{
				ID:                uint(len(articleCategories) + 1),
				ArticleID:         articleID,
				WasteCategoryID:   wasteCategories[rand.Intn(len(wasteCategories))].ID,
				ContentCategoryID: int(contentCategories[rand.Intn(len(contentCategories))].ID),
				CreatedAt:         randomDate(startDate, endDate),
			}
			articleCategories = append(articleCategories, category)
		}

		articles[i] = article
	}

	return articles, articleSections, articleCategories
}
