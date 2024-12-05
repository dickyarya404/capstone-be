package database

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	taskEntity "github.com/sawalreverr/recything/internal/task/manage_task/entity"
)

func (m *mysqlDatabase) InitTasks() {
	if err := m.GetDB().Migrator().DropTable(&taskEntity.TaskStep{}); err != nil {
		return
	}
	if err := m.GetDB().Migrator().DropTable(&taskEntity.TaskChallenge{}); err != nil {
		return
	}

	if err := m.GetDB().AutoMigrate(&taskEntity.TaskStep{}); err != nil {
		return
	}
	if err := m.GetDB().AutoMigrate(&taskEntity.TaskChallenge{}); err != nil {
		return
	}

	taskChallenges, taskSteps := generateTask()

	for _, taskChallenge := range taskChallenges {
		m.GetDB().FirstOrCreate(&taskChallenge, taskChallenge)
	}

	for _, taskStep := range taskSteps {
		m.GetDB().FirstOrCreate(&taskStep, taskStep)
	}

	log.Println("Task Challenge data added!")
}

func generateTask() ([]taskEntity.TaskChallenge, []taskEntity.TaskStep) {
	gofakeit.Seed(0)

	startDateRange := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDateRange := time.Date(2024, 6, 30, 23, 59, 59, 999, time.UTC)

	taskChallenges := make([]taskEntity.TaskChallenge, 20)
	taskSteps := make([]taskEntity.TaskStep, 0)
	taskStepID := 1

	for i := 0; i < 20; i++ {
		challengeID := fmt.Sprintf("TM%04d", i+1)
		startDate := randomDate(startDateRange, endDateRange)
		endDate := startDate.Add(time.Duration(rand.Intn(7)+1) * 24 * time.Hour)

		taskChallenge := taskEntity.TaskChallenge{
			ID:          challengeID,
			Title:       gofakeit.Sentence(6),
			Description: gofakeit.Paragraph(1, 2, 3, ""),
			Thumbnail:   gofakeit.ImageURL(640, 480),
			StartDate:   startDate,
			EndDate:     endDate,
			Point:       rand.Intn(2301) + 200,
			Status:      true,
			AdminId:     "AD0001",
			CreatedAt:   randomDate(startDateRange, startDate),
		}

		stepCount := rand.Intn(4) + 2
		for j := 0; j < stepCount; j++ {
			taskStep := taskEntity.TaskStep{
				ID:              taskStepID,
				TaskChallengeId: challengeID,
				Title:           fmt.Sprintf("Step %d", j+1),
				Description:     gofakeit.Paragraph(1, 2, 3, ""),
				CreatedAt:       randomDate(startDateRange, startDate),
			}
			taskSteps = append(taskSteps, taskStep)
			taskStepID++
		}

		taskChallenges[i] = taskChallenge
	}

	return taskChallenges, taskSteps
}
