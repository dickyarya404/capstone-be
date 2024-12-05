package database

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	art "github.com/sawalreverr/recything/internal/article"
	vid "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

func (m *mysqlDatabase) InitComment() {
	if err := m.GetDB().Migrator().DropTable(&art.ArticleComment{}); err != nil {
		return
	}
	if err := m.GetDB().Migrator().DropTable(&vid.Comment{}); err != nil {
		return
	}

	if err := m.GetDB().AutoMigrate(&art.ArticleComment{}); err != nil {
		return
	}
	if err := m.GetDB().AutoMigrate(&vid.Comment{}); err != nil {
		return
	}

	articleComments, videoComments := generateComment()

	for _, articleComment := range articleComments {
		m.GetDB().FirstOrCreate(&articleComment, articleComment)
	}

	for _, videoComment := range videoComments {
		m.GetDB().FirstOrCreate(&videoComment, videoComment)
	}

	log.Println("Comments data added!")
}

func randomUserID() string {
	return fmt.Sprintf("USR%04d", rand.Intn(50)+1)
}

func generateComment() ([]art.ArticleComment, []vid.Comment) {
	gofakeit.Seed(0)

	articleComments := make([]art.ArticleComment, 0)
	commentID := 1
	for i := 1; i <= 50; i++ {
		articleID := fmt.Sprintf("ART%04d", i)
		for j := 0; j < 30; j++ {
			comment := art.ArticleComment{
				ID:        uint(commentID),
				UserID:    randomUserID(),
				ArticleID: articleID,
				Comment:   gofakeit.Sentence(5),
				CreatedAt: gofakeit.DateRange(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()),
			}
			articleComments = append(articleComments, comment)
			commentID++
		}
	}

	videoComments := make([]vid.Comment, 0)
	for i := 1; i <= 50; i++ {
		videoID := i
		for j := 0; j < 30; j++ {
			comment := vid.Comment{
				ID:        commentID,
				UserID:    randomUserID(),
				VideoID:   videoID,
				Comment:   gofakeit.Sentence(5),
				CreatedAt: gofakeit.DateRange(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()),
			}
			videoComments = append(videoComments, comment)
			commentID++
		}
	}

	return articleComments, videoComments
}
