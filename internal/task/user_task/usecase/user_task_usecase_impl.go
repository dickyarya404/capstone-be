package usecase

import (
	"errors"
	"mime/multipart"
	"time"

	"github.com/sawalreverr/recything/internal/helper"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	"github.com/sawalreverr/recything/internal/task/user_task/dto"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
	"github.com/sawalreverr/recything/internal/task/user_task/repository"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type UserTaskUsecaseImpl struct {
	UserTaskRepository repository.UserTaskRepository
}

func NewUserTaskUsecase(repository repository.UserTaskRepository) UserTaskUsecase {
	return &UserTaskUsecaseImpl{UserTaskRepository: repository}
}

func (usecase *UserTaskUsecaseImpl) GetAllTasksUsecase() ([]task.TaskChallenge, error) {
	userTask, err := usecase.UserTaskRepository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return userTask, nil
}

func (usecase *UserTaskUsecaseImpl) GetTaskByIdUsecase(id string) (*task.TaskChallenge, error) {
	userTask, err := usecase.UserTaskRepository.GetTaskById(id)
	if err != nil {
		return nil, pkg.ErrTaskNotFound
	}

	return userTask, nil
}

func (usecase *UserTaskUsecaseImpl) CreateUserTaskUsecase(taskChallengeId string, userId string) (*user_task.UserTaskChallenge, error) {
	findtask, errFindTask := usecase.UserTaskRepository.FindTask(taskChallengeId)

	if errFindTask != nil {
		return nil, pkg.ErrTaskNotFound
	}

	if !findtask.Status {
		return nil, pkg.ErrTaskCannotBeFollowed
	}

	if _, err := usecase.UserTaskRepository.FindUserTask(userId, taskChallengeId); err == nil {
		return nil, pkg.ErrUserTaskExist
	}

	if _, err := usecase.UserTaskRepository.FindUserHasSameTask(userId, taskChallengeId); err == nil {
		return nil, pkg.ErrUserTaskExist
	}

	lastId, _ := usecase.UserTaskRepository.FindLastIdTaskChallenge()
	id := helper.GenerateCustomID(lastId, "UT")
	userTask := &user_task.UserTaskChallenge{
		ID:              id,
		UserId:          userId,
		TaskChallengeId: taskChallengeId,
		AcceptedAt:      time.Now(),
		ImageTask:       []user_task.UserTaskImage{},
		StatusProgress:  "in_progress",
		UserTaskSteps:   []user_task.UserTaskStep{},
	}

	userTask.TaskChallenge = *findtask

	userTaskData, err := usecase.UserTaskRepository.CreateUserTask(userTask)
	if err != nil {
		return nil, err
	}
	return userTaskData, nil
}

func (usecase *UserTaskUsecaseImpl) UploadImageTaskUsecase(request *dto.UploadImageTask, fileImage []*multipart.FileHeader, userId string, userTaskId string) (*user_task.UserTaskChallenge, error) {
	findUserTask, errFind := usecase.UserTaskRepository.FindUserTask(userId, userTaskId)

	if errFind != nil {
		return nil, pkg.ErrUserTaskNotFound
	}

	findTask, errFindTask := usecase.UserTaskRepository.FindTask(findUserTask.TaskChallengeId)

	if errFindTask != nil {
		return nil, pkg.ErrTaskNotFound
	}

	if !findTask.Status {
		return nil, pkg.ErrTaskCannotBeFollowed
	}
	countImage := len(fileImage)
	countTaskSteps := len(findTask.TaskSteps) * 3
	if countImage > countTaskSteps {
		return nil, pkg.ErrImagesExceed
	}

	if findUserTask.StatusProgress == "done" {
		return nil, pkg.ErrUserTaskDone
	}

	findUserSteps, errFindUserSteps := usecase.UserTaskRepository.FindUserSteps(userTaskId)

	if errFindUserSteps != nil {
		return nil, errFindUserSteps
	}

	for _, userStep := range findUserSteps {
		if !userStep.Completed {
			return nil, pkg.ErrUserTaskNotCompleted
		}
	}

	validImages, errImages := helper.ImagesValidation(fileImage)
	if errImages != nil {
		return nil, errImages
	}

	var imageUrls []string
	for _, image := range validImages {
		imageUrl, err := helper.UploadToCloudinary(image, "task_images+"+userTaskId)
		if err != nil {
			return nil, pkg.ErrUploadCloudinary
		}
		imageUrls = append(imageUrls, imageUrl)
	}

	data := &user_task.UserTaskChallenge{
		DescriptionImage: request.Description,
		StatusProgress:   "done",
		Point:            findTask.Point,
		ImageTask:        []user_task.UserTaskImage{},
	}

	for _, image := range imageUrls {
		data.ImageTask = append(data.ImageTask, user_task.UserTaskImage{
			ImageUrl: image,
		})
	}

	userTask, err := usecase.UserTaskRepository.UploadImageTask(data, userTaskId)
	if err != nil {
		return nil, err
	}
	return userTask, nil

}

func (usecase *UserTaskUsecaseImpl) GetUserTaskByUserIdUsecase(userId string) ([]user_task.UserTaskChallenge, error) {
	userTask, err := usecase.UserTaskRepository.GetUserTaskByUserId(userId)
	if err != nil {
		return nil, err
	}
	if len(userTask) == 0 {
		return nil, pkg.ErrUserNoHasTask
	}
	return userTask, nil
}

func (usecase *UserTaskUsecaseImpl) GetUserTaskDoneByUserIdUsecase(userId string) ([]user_task.UserTaskChallenge, error) {
	userTask, err := usecase.UserTaskRepository.GetUserTaskDoneByUserId(userId)
	if err != nil {
		return nil, err
	}
	if len(userTask) == 0 {
		return nil, pkg.ErrUserNoHasTask
	}
	return userTask, nil
}

func (usecase *UserTaskUsecaseImpl) UpdateUserTaskUsecase(request *dto.UpdateUserTaskRequest, fileImage []*multipart.FileHeader, userId string, userTaskId string) (*user_task.UserTaskChallenge, error) {
	findUserTask, errFind := usecase.UserTaskRepository.FindUserTask(userId, userTaskId)

	if errFind != nil {
		return nil, pkg.ErrUserTaskNotFound
	}

	if findUserTask.StatusAccept != "reject" {
		return nil, pkg.ErrUserTaskNotReject
	}

	findTask, errFindTask := usecase.UserTaskRepository.FindTask(findUserTask.TaskChallengeId)

	if errFindTask != nil {
		return nil, pkg.ErrTaskNotFound
	}
	countImage := len(fileImage)
	lenTaskSteps := len(findTask.TaskSteps)
	countTaskSteps := lenTaskSteps * 3

	if countImage > countTaskSteps {
		return nil, pkg.ErrImagesExceed
	}

	validImages, errImages := helper.ImagesValidation(fileImage)
	if errImages != nil {
		return nil, errImages
	}

	var imageUrls []string
	for _, image := range validImages {
		imageUrl, err := helper.UploadToCloudinary(image, "task_images_update+"+userTaskId)
		if err != nil {
			return nil, pkg.ErrUploadCloudinary
		}
		imageUrls = append(imageUrls, imageUrl)
	}

	data := &user_task.UserTaskChallenge{
		DescriptionImage: request.Description,
		StatusAccept:     "need_rivew",
		ImageTask:        []user_task.UserTaskImage{},
	}

	for _, image := range imageUrls {
		data.ImageTask = append(data.ImageTask, user_task.UserTaskImage{
			ImageUrl: image,
		})
	}

	userTask, err := usecase.UserTaskRepository.UpdateUserTask(data, userTaskId)
	if err != nil {
		return nil, err
	}
	return userTask, nil
}

func (usecase *UserTaskUsecaseImpl) GetUserTaskDetailsUsecase(userTaskId string, userId string) (*user_task.UserTaskChallenge, []*user_task.UserTaskImage, error) {
	userTask, imageTask, err := usecase.UserTaskRepository.GetUserTaskDetails(userTaskId, userId)
	if err != nil {
		return nil, nil, pkg.ErrUserTaskNotFound
	}
	return userTask, imageTask, nil
}

func (usecase *UserTaskUsecaseImpl) GetHistoryPointByUserIdUsecase(userId string) ([]user_task.UserTaskChallenge, int, error) {

	userTask, err := usecase.UserTaskRepository.GetHistoryPointByUserId(userId)
	if err != nil {
		return nil, 0, err
	}
	if len(userTask) == 0 {
		return nil, 0, pkg.ErrUserNoHasTask
	}

	var totalPoint int
	for _, task := range userTask {
		totalPoint += task.Point
	}
	return userTask, totalPoint, nil
}

func (usecase *UserTaskUsecaseImpl) UpdateTaskStepUsecase(request *dto.UpdateTaskStepRequest, userId string) (*user_task.UserTaskChallenge, error) {
	userTask, errUserTask := usecase.UserTaskRepository.FindUserTask(userId, request.UserTaskId)
	if errUserTask != nil {
		return nil, pkg.ErrUserTaskNotFound
	}

	if userTask.StatusProgress != "in_progress" {
		return nil, pkg.ErrUserTaskDone
	}

	if userTask.StatusAccept != "need_rivew" {
		return nil, pkg.ErrUserTaskAlreadyApprove
	}

	taskStep, errStep := usecase.UserTaskRepository.FindTaskStep(request.TaskStepId, userTask.TaskChallengeId)
	if errStep != nil {
		if errStep == gorm.ErrRecordNotFound {
			return nil, pkg.ErrTaskStepNotFound
		}
		return nil, errStep
	}

	userTaskStep, errUserTaskStep := usecase.UserTaskRepository.FindUserTaskStep(userTask.ID, taskStep.ID)
	if errUserTaskStep != nil {
		return nil, pkg.ErrUserTaskStepNotFound
	}

	if userTaskStep.Completed {
		return nil, pkg.ErrUserTaskStepAlreadyCompleted
	}

	// Find completed user task steps
	completedSteps, err := usecase.UserTaskRepository.FindCompletedUserSteps(userTask.ID)
	if err != nil {
		return nil, err
	}

	if len(completedSteps) > 0 {
		nextStepID := completedSteps[len(completedSteps)-1].TaskStepID + 1
		if request.TaskStepId != nextStepID {
			return nil, pkg.ErrStepNotInOrder
		}
	} else {
		firstStep := userTask.TaskChallenge.TaskSteps[0]
		if request.TaskStepId != firstStep.ID {
			return nil, pkg.ErrStepNotInOrder
		}
	}

	userTaskStep.Completed = true
	if err := usecase.UserTaskRepository.UpdateUserTaskStep(userTaskStep); err != nil {
		return nil, err
	}

	updatedUserTask, err := usecase.UserTaskRepository.FindUserTask(userId, request.UserTaskId)
	if err != nil {
		return nil, err
	}

	return updatedUserTask, nil
}

func (usecase *UserTaskUsecaseImpl) GetUserTaskByUserTaskId(userId string, userTaskId string) (*user_task.UserTaskChallenge, error) {

	userTask, err := usecase.UserTaskRepository.FindUserTask(userId, userTaskId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrUserTaskNotFound
		}
		return nil, err
	}
	return userTask, nil
}

func (usecase *UserTaskUsecaseImpl) GetUserTaskRejectedByUserId(userId string, userTaskId string) (*user_task.UserTaskChallenge, error) {
	userTask, errUserTask := usecase.UserTaskRepository.GetUserTaskRejectedByUserId(userId, userTaskId)

	if errUserTask != nil {
		if errors.Is(errUserTask, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrUserTaskNotFound
		}
		return nil, errUserTask
	}
	return userTask, nil
}
