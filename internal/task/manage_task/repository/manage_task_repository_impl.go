package repository

import (
	"log"
	"time"

	"github.com/sawalreverr/recything/internal/database"
	"github.com/sawalreverr/recything/internal/task/manage_task/entity"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	"gorm.io/gorm"
)

type ManageTaskRepositoryImpl struct {
	DB database.Database
}

func NewManageTaskRepository(db database.Database) ManageTaskRepository {
	return &ManageTaskRepositoryImpl{DB: db}
}

func (r *ManageTaskRepositoryImpl) CreateTask(task *task.TaskChallenge) (*task.TaskChallenge, error) {
	if err := r.DB.GetDB().Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (repository *ManageTaskRepositoryImpl) FindLastIdTaskChallenge() (string, error) {
	var task *task.TaskChallenge
	if err := repository.DB.GetDB().Unscoped().Order("id desc").First(&task).Error; err != nil {
		return "TM0000", err
	}
	return task.ID, nil
}

func (repository *ManageTaskRepositoryImpl) GetTaskChallengePagination(page int, limit int, status string, endDate string) ([]task.TaskChallenge, int, error) {
	var tasks []task.TaskChallenge
	var total int64
	offset := (page - 1) * limit

	baseQuery := repository.DB.GetDB().Model(&task.TaskChallenge{})

	if status != "" {
		if status == "true" {
			baseQuery = baseQuery.Where("status = ?", true)
		} else if status == "false" {
			baseQuery = baseQuery.Where("status = ?", false)
		}
	}
	totalQuery := baseQuery
	if err := totalQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	dataQuery := baseQuery.
		Preload("TaskSteps").
		Preload("Admin").
		Limit(limit).
		Offset(offset)

	if endDate != "" {
		if endDate == "desc" {
			dataQuery = dataQuery.Order("end_date DESC")
		} else if endDate == "asc" {
			dataQuery = dataQuery.Order("end_date ASC")
		}
	} else {
		dataQuery = dataQuery.Order("created_at DESC")
	}

	if err := dataQuery.Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, int(total), nil
}

func (repository *ManageTaskRepositoryImpl) GetTaskById(id string) (*task.TaskChallenge, error) {
	var task *task.TaskChallenge
	if err := repository.DB.GetDB().
		Preload("TaskSteps").
		Preload("Admin").
		First(&task, "id = ?", id).
		Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (repository *ManageTaskRepositoryImpl) FindTask(id string) (*task.TaskChallenge, error) {
	var task task.TaskChallenge
	if err := repository.DB.GetDB().Preload("TaskSteps").Where("id = ?", id).First(&task).Error; err != nil {
		log.Println("Error finding task:", err)
		return nil, err
	}
	return &task, nil
}

func (repository *ManageTaskRepositoryImpl) UpdateTaskChallenge(taskChallenge *task.TaskChallenge, taskId string) (*task.TaskChallenge, error) {
	tx := repository.DB.GetDB().Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("Transaction rollback due to panic:", r)
		}
	}()

	if len(taskChallenge.TaskSteps) != 0 {
		if err := tx.Where("task_challenge_id = ?", taskId).Delete(&task.TaskStep{}).Error; err != nil {
			log.Println("Error deleting task steps:", err)
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(taskChallenge).Error; err != nil {
		log.Println("Error updating task challenge:", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Error committing transaction:", err)
		return nil, err
	}

	return taskChallenge, nil
}

func (repository *ManageTaskRepositoryImpl) DeleteTaskChallenge(taskId string) error {
	if err := repository.DB.GetDB().Where("id = ?", taskId).Delete(&task.TaskChallenge{}).Error; err != nil {
		return err
	}
	return nil
}

func (repository *ManageTaskRepositoryImpl) UpdateTaskChallengeStatus() error {
	now := time.Now()
	result := repository.DB.GetDB().Model(&entity.TaskChallenge{}).
		Where("end_date < ? AND status = ?", now, true).Update("status", false)

	if result.Error != nil {
		log.Printf("Error updating task challenge status: %v", result.Error)
		return result.Error
	} else {
		log.Printf("Updated %d task challenge(s) status to false", result.RowsAffected)
		return nil
	}
}
