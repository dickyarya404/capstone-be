package database

import (
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	videoEntity "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

func (m *mysqlDatabase) InitVideos() {
	if err := m.GetDB().Migrator().DropTable(&videoEntity.VideoCategory{}); err != nil {
		return
	}
	if err := m.GetDB().Migrator().DropTable(&videoEntity.Video{}); err != nil {
		return
	}

	if err := m.GetDB().AutoMigrate(&videoEntity.VideoCategory{}); err != nil {
		return
	}
	if err := m.GetDB().AutoMigrate(&videoEntity.Video{}); err != nil {
		return
	}

	videos, videoCategories := generateVideo()

	for _, video := range videos {
		m.GetDB().FirstOrCreate(&video, video)
	}

	for _, videoCategory := range videoCategories {
		m.GetDB().FirstOrCreate(&videoCategory, videoCategory)
	}

	log.Println("Video data added!")
}

func generateVideo() ([]videoEntity.Video, []videoEntity.VideoCategory) {
	gofakeit.Seed(0)

	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 6, 30, 23, 59, 59, 999, time.UTC)

	videos := make([]videoEntity.Video, 50)
	videoCategories := make([]videoEntity.VideoCategory, 0)

	for i := 0; i < 50; i++ {
		videoID := i + 1

		video := videoEntity.Video{
			ID:          videoID,
			Title:       gofakeit.Sentence(6),
			Description: gofakeit.Paragraph(1, 2, 3, ""),
			Thumbnail:   gofakeit.ImageURL(640, 480),
			Link:        "https://www.youtube.com/watch?v=CGd3lgxReFE",
			CreatedAt:   randomDate(startDate, endDate),
		}

		categoryCount := rand.Intn(3) + 1
		for j := 0; j < categoryCount; j++ {
			videoCategory := videoEntity.VideoCategory{
				VideoID:           videoID,
				ContentCategoryID: contentCategories[rand.Intn(len(contentCategories))].ID,
				WasteCategoryID:   wasteCategories[rand.Intn(len(wasteCategories))].ID,
				CreatedAt:         randomDate(startDate, endDate),
			}
			videoCategories = append(videoCategories, videoCategory)
		}

		videos[i] = video
	}

	return videos, videoCategories
}
