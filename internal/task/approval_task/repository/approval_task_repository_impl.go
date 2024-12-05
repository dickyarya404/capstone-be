package repository

import (
	"time"

	achievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
	"github.com/sawalreverr/recything/internal/database"
	"github.com/sawalreverr/recything/internal/helper"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
	user_entity "github.com/sawalreverr/recything/internal/user"
)

type ApprovalTaskRepositoryImpl struct {
	DB database.Database
}

func NewApprovalTaskRepositoryImpl(db database.Database) *ApprovalTaskRepositoryImpl {
	return &ApprovalTaskRepositoryImpl{
		DB: db,
	}
}

func (repository *ApprovalTaskRepositoryImpl) GetAllApprovalTaskPagination(limit int, offset int) ([]*user_task.UserTaskChallenge, int, error) {
	var tasks []*user_task.UserTaskChallenge
	var total int64
	offset = (offset - 1) * limit

	if err := repository.DB.GetDB().Model(&user_task.UserTaskChallenge{}).
		Where("status_progress = ?", "done").
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}
	if err := repository.DB.GetDB().
		Preload("TaskChallenge.TaskSteps").
		Preload("User").
		Limit(limit).
		Offset(offset).Order("id desc").
		Where("status_progress = ?", "done").
		Find(&tasks).Error; err != nil {
		return nil, 0, err
	}
	return tasks, int(total), nil
}

func (repository *ApprovalTaskRepositoryImpl) FindUserTask(userTaskId string) (*user_task.UserTaskChallenge, error) {
	var userTask user_task.UserTaskChallenge
	if err := repository.DB.GetDB().
		Where("id = ?", userTaskId).
		First(&userTask).Error; err != nil {
		return nil, err
	}
	return &userTask, nil
}

func (repository *ApprovalTaskRepositoryImpl) ApproveUserTask(userTaskId string) error {
	var userTask user_task.UserTaskChallenge
	tx := repository.DB.GetDB().Begin()

	if err := tx.Where("id = ?", userTaskId).First(&userTask).Error; err != nil {
		tx.Rollback()
		return err
	}

	acceptedAt := time.Now()
	if err := tx.Model(&userTask).Where("id = ?", userTaskId).Updates(map[string]interface{}{
		"status_accept": "accept",
		"accepted_at":   acceptedAt,
		"reason":        "",
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	point := userTask.Point

	var user user_entity.User
	if err := tx.Where("id = ?", userTask.UserId).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	pointBonus := helper.BonusTask(user.Badge, point)

	pointUpdate := int(user.Point) + pointBonus

	var achievements []achievement.Achievement

	if err := tx.Model(&achievement.Achievement{}).Order("target_point desc").Find(&achievements).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(achievements) == 0 {
		tx.Rollback()
		return nil
	}
	var badge string
	for _, ach := range achievements {
		if pointUpdate >= ach.TargetPoint {
			badge = ach.BadgeUrlUser
			break
		}
	}

	if badge != "" {
		if err := tx.Model(&user_entity.User{}).Where("id = ?", userTask.UserId).
			Updates(map[string]interface{}{"badge": badge, "point": pointUpdate}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repository *ApprovalTaskRepositoryImpl) RejectUserTask(data *user_task.UserTaskChallenge, userTaskId string) error {
	if err := repository.DB.GetDB().Where("id = ?", userTaskId).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (repository *ApprovalTaskRepositoryImpl) GetUserTaskDetails(userTaskId string) (*user_task.UserTaskChallenge, []*user_task.UserTaskImage, error) {
	var userTask user_task.UserTaskChallenge
	if err := repository.DB.GetDB().
		Preload("User").
		Preload("TaskChallenge").
		Where("id = ?", userTaskId).
		Where("status_progress = ?", "done").
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

func (repository *ApprovalTaskRepositoryImpl) FindUserTaskForApprove(userTaskId string) (*user_task.UserTaskChallenge, error) {
	statusAccept := "need_rivew"
	var userTask user_task.UserTaskChallenge
	if err := repository.DB.GetDB().
		Where("id = ?", userTaskId).
		Where("status_accept = ?", statusAccept).
		First(&userTask).Error; err != nil {
		return nil, err
	}
	return &userTask, nil
}

func (repository *ApprovalTaskRepositoryImpl) FindUserTaskForReject(userTaskId string) (*user_task.UserTaskChallenge, error) {
	statusAccept := "need_rivew"
	var userTask user_task.UserTaskChallenge
	if err := repository.DB.GetDB().
		Where("id = ?", userTaskId).
		Where("status_accept = ?", statusAccept).
		First(&userTask).Error; err != nil {
		return nil, err
	}
	return &userTask, nil
}
