package repository

import (
	"github.com/sawalreverr/recything/internal/database"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
	"gorm.io/gorm"
)

type UserTaskRepositoryImpl struct {
	DB database.Database
}

func NewUserTaskRepository(db database.Database) UserTaskRepository {
	return &UserTaskRepositoryImpl{DB: db}
}

func (repository *UserTaskRepositoryImpl) GetAllTasks() ([]task.TaskChallenge, error) {
	var tasks []task.TaskChallenge
	if err := repository.DB.GetDB().
		Preload("TaskSteps").
		Order("id desc").
		Find(&tasks, "status = ?", true).
		Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (repository *UserTaskRepositoryImpl) GetTaskById(id string) (*task.TaskChallenge, error) {
	var task task.TaskChallenge
	if err := repository.DB.GetDB().
		Preload("TaskSteps").
		Where("id = ?", id).
		First(&task).
		Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (repository *UserTaskRepositoryImpl) FindLastIdTaskChallenge() (string, error) {
	var task user_task.UserTaskChallenge
	if err := repository.DB.GetDB().Unscoped().Order("id desc").First(&task).Error; err != nil {
		return "UT0000", err
	}
	return task.ID, nil
}

func (repository *UserTaskRepositoryImpl) CreateUserTask(userTask *user_task.UserTaskChallenge) (*user_task.UserTaskChallenge, error) {
	var result user_task.UserTaskChallenge
	err := repository.DB.GetDB().Transaction(func(tx *gorm.DB) error {
		// Create the UserTaskChallenge
		if err := tx.Create(userTask).Error; err != nil {
			return err
		}

		// Create UserTaskSteps for each TaskStep in the TaskChallenge
		for _, taskStep := range userTask.TaskChallenge.TaskSteps {
			userTaskStep := user_task.UserTaskStep{
				UserTaskChallengeID: userTask.ID,
				TaskStepID:          taskStep.ID,
				Completed:           false,
			}
			if err := tx.Create(&userTaskStep).Error; err != nil {
				return err
			}
		}

		if err := tx.Preload("TaskChallenge.TaskSteps").
			Preload("UserTaskSteps").
			Where("user_task_challenges.id = ?", userTask.ID).
			First(&result).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repository *UserTaskRepositoryImpl) FindUserTask(userId string, userTaskId string) (*user_task.UserTaskChallenge, error) {
	var userTask user_task.UserTaskChallenge
	if err := repository.DB.GetDB().
		Preload("TaskChallenge.TaskSteps").
		Preload("UserTaskSteps").
		Where("user_id = ? and id = ?", userId, userTaskId).
		First(&userTask).Error; err != nil {
		return nil, err
	}
	return &userTask, nil
}

func (repository *UserTaskRepositoryImpl) FindTask(taskId string) (*task.TaskChallenge, error) {
	var task task.TaskChallenge
	if err := repository.DB.GetDB().
		Preload("TaskSteps").
		Where("id = ?", taskId).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (repository *UserTaskRepositoryImpl) UploadImageTask(userTask *user_task.UserTaskChallenge, userTaskId string) (*user_task.UserTaskChallenge, error) {
	tx := repository.DB.GetDB().Begin()
	if err := tx.Model(&user_task.UserTaskChallenge{}).Where("id = ?", userTaskId).Updates(map[string]interface{}{
		"description_image": userTask.DescriptionImage,
		"status_progress":   userTask.StatusProgress,
		"point":             userTask.Point,
	}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var count int64
	if err := tx.Model(&user_task.UserTaskImage{}).Where("user_task_challenge_id = ?", userTaskId).Count(&count).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if count == 0 {
		if err := tx.Where("user_task_challenge_id = ?", userTaskId).Delete(&user_task.UserTaskImage{}).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		for _, img := range userTask.ImageTask {
			img.UserTaskChallengeID = userTaskId
			if err := tx.Create(&img).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	tx.Preload("UserTaskImage").
		Preload("TaskChallenge.TaskSteps").
		Preload("UserTaskSteps").
		Where("id = ?", userTaskId).
		First(&userTask)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return userTask, nil

}

func (repository *UserTaskRepositoryImpl) GetUserTaskByUserId(userId string) ([]user_task.UserTaskChallenge, error) {
	var userTask []user_task.UserTaskChallenge
	if err := repository.DB.GetDB().
		Preload("TaskChallenge.TaskSteps").
		Preload("UserTaskSteps").
		Where("user_id = ? and status_progress = ?", userId, "in_progress").
		Order("id desc").
		Find(&userTask).Error; err != nil {
		return nil, err
	}
	return userTask, nil
}

func (repository *UserTaskRepositoryImpl) GetUserTaskDoneByUserId(userId string) ([]user_task.UserTaskChallenge, error) {
	var userTask []user_task.UserTaskChallenge
	if err := repository.DB.GetDB().
		Preload("TaskChallenge.TaskSteps").
		Preload("UserTaskSteps").
		Where("user_id = ? and status_progress = ?", userId, "done").
		Order("id desc").
		Find(&userTask).Error; err != nil {
		return nil, err
	}
	return userTask, nil
}

func (repository *UserTaskRepositoryImpl) FindUserHasSameTask(userId string, taskId string) (*user_task.UserTaskChallenge, error) {

	var userTask user_task.UserTaskChallenge
	if err := repository.DB.GetDB().Where("user_id = ? and task_challenge_id = ?", userId, taskId).First(&userTask).Error; err != nil {
		return nil, err
	}
	return &userTask, nil
}

// update user task if reject
func (repository *UserTaskRepositoryImpl) UpdateUserTask(userTask *user_task.UserTaskChallenge, userTaskId string) (*user_task.UserTaskChallenge, error) {
	tx := repository.DB.GetDB().Begin()
	if err := tx.Model(&user_task.UserTaskChallenge{}).Where("id = ?", userTaskId).Updates(map[string]interface{}{
		"description_image": userTask.DescriptionImage,
		"status_accept":     userTask.StatusAccept,
	}).
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var count int64
	if err := tx.Model(&user_task.UserTaskImage{}).Where("user_task_challenge_id = ?", userTaskId).Count(&count).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if count != 0 {
		if err := tx.Where("user_task_challenge_id = ?", userTaskId).Delete(&user_task.UserTaskImage{}).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		for _, img := range userTask.ImageTask {
			img.UserTaskChallengeID = userTaskId
			if err := tx.Create(&img).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	tx.Preload("UserTaskImage").
		Preload("TaskChallenge.TaskSteps").
		Preload("UserTaskSteps").
		Where("id = ?", userTaskId).
		First(&userTask)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return userTask, nil

}

func (repository *UserTaskRepositoryImpl) GetUserTaskDetails(userTaskId string, userId string) (*user_task.UserTaskChallenge, []*user_task.UserTaskImage, error) {
	var userTask user_task.UserTaskChallenge
	if err := repository.DB.GetDB().
		Preload("User").
		Preload("TaskChallenge").
		Preload("UserTaskSteps").
		Where("id = ? ", userTaskId).
		Where("user_id = ?", userId).
		First(&userTask).Error; err != nil {
		return nil, nil, err
	}

	var images []*user_task.UserTaskImage
	if err := repository.DB.GetDB().
		Where("user_task_challenge_id = ?", userTaskId).
		Find(&images).Error; err != nil {
		return nil, nil, err
	}
	return &userTask, images, nil
}

func (repository *UserTaskRepositoryImpl) GetHistoryPointByUserId(userId string) ([]user_task.UserTaskChallenge, error) {
	var userTask []user_task.UserTaskChallenge
	if err := repository.DB.GetDB().
		Preload("TaskChallenge").
		Where("user_id = ?", userId).
		Where("status_accept = ?", "accept").
		Order("accepted_at desc").
		Find(&userTask).Error; err != nil {
		return nil, err
	}
	return userTask, nil
}

func (repository *UserTaskRepositoryImpl) FindTaskStep(stepId int, taskId string) (*task.TaskStep, error) {
	var taskStep task.TaskStep
	if err := repository.DB.GetDB().
		Where("id = ?", stepId).
		Where("task_challenge_id = ?", taskId).
		First(&taskStep).Error; err != nil {
		return nil, err
	}
	return &taskStep, nil
}

func (repository *UserTaskRepositoryImpl) UpdateUserTaskStep(userTaskStep *user_task.UserTaskStep) error {
	return repository.DB.GetDB().Save(userTaskStep).Error
}

func (repository *UserTaskRepositoryImpl) FindUserTaskStep(userTaskChallengeID string, taskStepID int) (*user_task.UserTaskStep, error) {
	var userTaskStep user_task.UserTaskStep
	err := repository.DB.GetDB().Where("user_task_challenge_id = ? AND task_step_id = ?", userTaskChallengeID, taskStepID).First(&userTaskStep).Error
	return &userTaskStep, err
}

func (repository *UserTaskRepositoryImpl) FindUserSteps(userTaskChallengeID string) ([]user_task.UserTaskStep, error) {
	var userTaskStep []user_task.UserTaskStep
	err := repository.DB.GetDB().Where("user_task_challenge_id = ?", userTaskChallengeID).Find(&userTaskStep).Error
	return userTaskStep, err
}

func (repository *UserTaskRepositoryImpl) FindCompletedUserSteps(userTaskChallengeID string) ([]user_task.UserTaskStep, error) {
	var userTaskSteps []user_task.UserTaskStep
	err := repository.DB.GetDB().Where("user_task_challenge_id = ? AND completed = ?", userTaskChallengeID, true).Order("task_step_id asc").Find(&userTaskSteps).Error
	return userTaskSteps, err
}

func (repository *UserTaskRepositoryImpl) GetUserTaskRejectedByUserId(userId string, userTaskId string) (*user_task.UserTaskChallenge, error) {
	var userTask user_task.UserTaskChallenge
	if err := repository.DB.GetDB().
		Preload("User").
		Preload("TaskChallenge.TaskSteps").
		Preload("UserTaskSteps").
		Where("id = ? ", userTaskId).
		Where("user_id = ?", userId).
		Where("status_accept = ?", "reject").
		First(&userTask).Error; err != nil {
		return nil, err
	}
	return &userTask, nil
}
