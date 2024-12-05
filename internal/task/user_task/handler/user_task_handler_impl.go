package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/task/user_task/dto"
	"github.com/sawalreverr/recything/internal/task/user_task/usecase"
	"github.com/sawalreverr/recything/pkg"
)

type UserTaskHandlerImpl struct {
	Usecase usecase.UserTaskUsecase
}

func NewUserTaskHandler(usecase usecase.UserTaskUsecase) UserTaskHandler {
	return &UserTaskHandlerImpl{Usecase: usecase}
}

func (handler *UserTaskHandlerImpl) GetAllTasksHandler(c echo.Context) error {
	userTask, err := handler.Usecase.GetAllTasksUsecase()
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var data []dto.DataUserTask

	for _, task := range userTask {
		var taskStep []dto.TaskSteps

		for _, step := range task.TaskSteps {
			taskStep = append(taskStep, dto.TaskSteps{
				Id:          step.ID,
				Title:       step.Title,
				Description: step.Description,
			})
		}
		data = append(data, dto.DataUserTask{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Thumbnail:   task.Thumbnail,
			StartDate:   task.StartDate,
			EndDate:     task.EndDate,
			Point:       task.Point,
			Status:      task.Status,
			TaskSteps:   taskStep,
		})
	}
	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *UserTaskHandlerImpl) GetTaskByIdHandler(c echo.Context) error {
	id := c.Param("taskId")
	task, err := handler.Usecase.GetTaskByIdUsecase(id)
	if err != nil {
		if errors.Is(err, pkg.ErrTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrTaskNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	var taskStep []dto.TaskSteps
	for _, step := range task.TaskSteps {
		taskStep = append(taskStep, dto.TaskSteps{
			Id:          step.ID,
			Title:       step.Title,
			Description: step.Description,
		})
	}
	data := dto.DataUserTask{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Thumbnail:   task.Thumbnail,
		StartDate:   task.StartDate,
		EndDate:     task.EndDate,
		Point:       task.Point,
		Status:      task.Status,
		TaskSteps:   taskStep,
	}
	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *UserTaskHandlerImpl) CreateUserTaskHandler(c echo.Context) error {
	taskChallengeId := c.Param("taskChallengeId")
	claims := c.Get("user").(*helper.JwtCustomClaims)
	userTask, err := handler.Usecase.CreateUserTaskUsecase(taskChallengeId, claims.UserID)
	if err != nil {
		if errors.Is(err, pkg.ErrTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrTaskNotFound.Error())
		}
		if errors.Is(err, pkg.ErrUserTaskExist) {
			return helper.ErrorHandler(c, http.StatusConflict, pkg.ErrUserTaskExist.Error())
		}
		if errors.Is(err, pkg.ErrTaskCannotBeFollowed) {
			return helper.ErrorHandler(c, http.StatusConflict, pkg.ErrTaskCannotBeFollowed.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var taskStep []dto.TaskSteps
	var userSteps []dto.DataUserSteps

	for _, step := range userTask.TaskChallenge.TaskSteps {
		taskStep = append(taskStep, dto.TaskSteps{
			Id:          step.ID,
			Title:       step.Title,
			Description: step.Description,
		})
	}

	for _, step := range userTask.UserTaskSteps {
		userSteps = append(userSteps, dto.DataUserSteps{
			Id:                  step.ID,
			UserTaskChallengeID: step.UserTaskChallengeID,
			TaskStepID:          step.TaskStepID,
			Completed:           step.Completed,
		})
	}
	data := dto.TaskChallengeResponseCreate{
		Id:          userTask.TaskChallenge.ID,
		Title:       userTask.TaskChallenge.Title,
		Description: userTask.TaskChallenge.Description,
		Thumbnail:   userTask.TaskChallenge.Thumbnail,
		StartDate:   userTask.TaskChallenge.StartDate,
		EndDate:     userTask.TaskChallenge.EndDate,
		Point:       userTask.TaskChallenge.Point,
		StatusTask:  userTask.TaskChallenge.Status,
		TaskSteps:   taskStep,
		UserSteps:   userSteps,
	}
	dataUsertask := dto.UserTaskResponseCreate{
		Id:             userTask.ID,
		StatusProgress: userTask.StatusProgress,
		TaskChalenge:   data,
	}
	responseData := helper.ResponseData(http.StatusCreated, "success", dataUsertask)
	return c.JSON(http.StatusCreated, responseData)
}

func (handler *UserTaskHandlerImpl) UploadImageTaskHandler(c echo.Context) error {
	var request dto.UploadImageTask
	claims := c.Get("user").(*helper.JwtCustomClaims)

	userTaskId := c.Param("userTaskId")

	jsonData := c.FormValue("json_data")

	if err := json.Unmarshal([]byte(jsonData), &request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	form, errForm := c.MultipartForm()
	if errForm != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, errForm.Error())
	}
	images := form.File["images"]
	if len(images) == 0 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "image is required")
	}

	userTask, err := handler.Usecase.UploadImageTaskUsecase(&request, images, claims.UserID, userTaskId)
	if err != nil {
		if errors.Is(err, pkg.ErrUserTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserTaskNotFound.Error())
		}
		if errors.Is(err, pkg.ErrUserTaskDone) {
			return helper.ErrorHandler(c, http.StatusConflict, pkg.ErrUserTaskDone.Error())
		}
		if errors.Is(err, errors.New("upload image size must less than 2MB")) {
			return helper.ErrorHandler(c, http.StatusBadRequest, "upload image size must less than 2MB")
		}
		if errors.Is(err, errors.New("only image allowed")) {
			return helper.ErrorHandler(c, http.StatusBadRequest, "only image allowed")
		}
		if errors.Is(err, pkg.ErrUploadCloudinary) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrUploadCloudinary.Error())
		}
		if errors.Is(err, pkg.ErrTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrTaskNotFound.Error())
		}
		if errors.Is(err, pkg.ErrImagesExceed) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrImagesExceed.Error())
		}
		if errors.Is(err, pkg.ErrUserTaskNotCompleted) {
			return helper.ErrorHandler(c, http.StatusConflict, pkg.ErrUserTaskNotCompleted.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	var taskStep []dto.TaskSteps
	var urlImages []dto.Images
	var userSteps []dto.DataUserSteps
	data := dto.UserTaskUploadImageResponse{
		Id:             userTask.ID,
		StatusProgress: userTask.StatusProgress,
		StatusAccept:   userTask.StatusAccept,
		Point:          userTask.Point,
		TaskChallenge: dto.DataTaskChallenges{
			Id:          userTask.TaskChallenge.ID,
			Title:       userTask.TaskChallenge.Title,
			Description: userTask.TaskChallenge.Description,
			Thumbnail:   userTask.TaskChallenge.Thumbnail,
			StartDate:   userTask.TaskChallenge.StartDate,
			EndDate:     userTask.TaskChallenge.EndDate,
			StatusTask:  userTask.TaskChallenge.Status,
			TaskSteps:   taskStep,
		},
		Images:    urlImages,
		UserSteps: userSteps,
	}

	for _, step := range userTask.UserTaskSteps {
		userSteps = append(userSteps, dto.DataUserSteps{
			Id:                  step.ID,
			UserTaskChallengeID: step.UserTaskChallengeID,
			TaskStepID:          step.TaskStepID,
			Completed:           step.Completed,
		})
	}

	for _, step := range userTask.TaskChallenge.TaskSteps {
		taskStep = append(taskStep, dto.TaskSteps{
			Id:          step.ID,
			Title:       step.Title,
			Description: step.Description,
		})
	}
	for _, image := range userTask.ImageTask {
		urlImages = append(urlImages, dto.Images{
			Images: image.ImageUrl,
		})
	}
	data.UserSteps = userSteps
	data.TaskChallenge.TaskSteps = taskStep
	data.Images = urlImages
	responseData := helper.ResponseData(http.StatusCreated, "success", data)
	return c.JSON(http.StatusOK, responseData)

}

func (handler *UserTaskHandlerImpl) GetUserTaskByUserIdHandler(c echo.Context) error {
	userId := c.Get("user").(*helper.JwtCustomClaims).UserID
	userTasks, err := handler.Usecase.GetUserTaskByUserIdUsecase(userId)
	if err != nil {
		if errors.Is(err, pkg.ErrUserNoHasTask) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrUserNoHasTask.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var data []dto.UserTaskGetByIdUserResponse
	for _, userTask := range userTasks {
		data = append(data, dto.UserTaskGetByIdUserResponse{
			Id:             userTask.ID,
			StatusProgress: userTask.StatusProgress,
			TaskChallenge: dto.TaskChallengeResponseCreate{
				Id:          userTask.TaskChallenge.ID,
				Title:       userTask.TaskChallenge.Title,
				Description: userTask.TaskChallenge.Description,
				Thumbnail:   userTask.TaskChallenge.Thumbnail,
				StartDate:   userTask.TaskChallenge.StartDate,
				EndDate:     userTask.TaskChallenge.EndDate,
				Point:       userTask.TaskChallenge.Point,
				StatusTask:  userTask.TaskChallenge.Status,
				TaskSteps:   []dto.TaskSteps{},
				UserSteps:   []dto.DataUserSteps{},
			},
		})

		for _, step := range userTask.TaskChallenge.TaskSteps {
			data[len(data)-1].TaskChallenge.TaskSteps = append(data[len(data)-1].TaskChallenge.TaskSteps, dto.TaskSteps{
				Id:          step.ID,
				Title:       step.Title,
				Description: step.Description,
			})
		}

		for _, step := range userTask.UserTaskSteps {
			data[len(data)-1].TaskChallenge.UserSteps = append(data[len(data)-1].TaskChallenge.UserSteps, dto.DataUserSteps{
				Id:                  step.ID,
				UserTaskChallengeID: step.UserTaskChallengeID,
				TaskStepID:          step.TaskStepID,
				Completed:           step.Completed,
			})
		}
	}

	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *UserTaskHandlerImpl) GetUserTaskDoneByUserIdHandler(c echo.Context) error {
	userId := c.Get("user").(*helper.JwtCustomClaims).UserID
	userTasks, err := handler.Usecase.GetUserTaskDoneByUserIdUsecase(userId)
	if err != nil {
		if errors.Is(err, pkg.ErrUserNoHasTask) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrUserNoHasTask.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var data []dto.GetUserTaskDoneByIdUserResponse
	for _, userTask := range userTasks {
		data = append(data, dto.GetUserTaskDoneByIdUserResponse{
			Id:             userTask.ID,
			StatusProgress: userTask.StatusProgress,
			StatusAccept:   userTask.StatusAccept,
			Point:          userTask.Point,
			ReasonReject:   userTask.Reason,
			TaskChallenge: dto.DataTaskChanlengesDone{
				Id:          userTask.TaskChallenge.ID,
				Title:       userTask.TaskChallenge.Title,
				Description: userTask.TaskChallenge.Description,
				Thumbnail:   userTask.TaskChallenge.Thumbnail,
				StartDate:   userTask.TaskChallenge.StartDate,
				EndDate:     userTask.TaskChallenge.EndDate,
				StatusTask:  userTask.TaskChallenge.Status,
				TaskSteps:   []dto.TaskSteps{},
				UserSteps:   []dto.DataUserSteps{},
			},
		})

		for _, step := range userTask.TaskChallenge.TaskSteps {
			data[len(data)-1].TaskChallenge.TaskSteps = append(data[len(data)-1].TaskChallenge.TaskSteps, dto.TaskSteps{
				Id:          step.ID,
				Title:       step.Title,
				Description: step.Description,
			})
		}

		for _, step := range userTask.UserTaskSteps {
			data[len(data)-1].TaskChallenge.UserSteps = append(data[len(data)-1].TaskChallenge.UserSteps, dto.DataUserSteps{
				Id:                  step.ID,
				UserTaskChallengeID: step.UserTaskChallengeID,
				TaskStepID:          step.TaskStepID,
				Completed:           step.Completed,
			})
		}
	}

	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *UserTaskHandlerImpl) UpdateUserTaskHandler(c echo.Context) error {
	var request dto.UpdateUserTaskRequest
	claims := c.Get("user").(*helper.JwtCustomClaims)

	userTaskId := c.Param("userTaskId")

	jsonData := c.FormValue("json_data")

	if err := json.Unmarshal([]byte(jsonData), &request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	form, errForm := c.MultipartForm()
	if errForm != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, errForm.Error())
	}
	images := form.File["images"]
	if len(images) == 0 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "image is required")
	}

	userTask, err := handler.Usecase.UpdateUserTaskUsecase(&request, images, claims.UserID, userTaskId)
	if err != nil {
		if errors.Is(err, pkg.ErrUserTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserTaskNotFound.Error())
		}
		if errors.Is(err, pkg.ErrUserTaskDone) {
			return helper.ErrorHandler(c, http.StatusConflict, pkg.ErrUserTaskDone.Error())
		}
		if errors.Is(err, errors.New("upload image size must less than 2MB")) {
			return helper.ErrorHandler(c, http.StatusBadRequest, "upload image size must less than 2MB")
		}
		if errors.Is(err, errors.New("only image allowed")) {
			return helper.ErrorHandler(c, http.StatusBadRequest, "only image allowed")
		}
		if errors.Is(err, pkg.ErrUploadCloudinary) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrUploadCloudinary.Error())
		}
		if errors.Is(err, pkg.ErrTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrTaskNotFound.Error())
		}
		if errors.Is(err, pkg.ErrImagesExceed) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrImagesExceed.Error())
		}
		if errors.Is(err, pkg.ErrUserTaskNotReject) {
			return helper.ErrorHandler(c, http.StatusConflict, pkg.ErrUserTaskNotReject.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	var taskStep []dto.TaskSteps
	var urlImages []dto.Images
	var userSteps []dto.DataUserSteps
	data := dto.UserTaskUploadImageResponse{
		Id:             userTask.ID,
		StatusProgress: userTask.StatusProgress,
		StatusAccept:   userTask.StatusAccept,
		Point:          userTask.Point,
		TaskChallenge: dto.DataTaskChallenges{
			Id:          userTask.TaskChallenge.ID,
			Title:       userTask.TaskChallenge.Title,
			Description: userTask.TaskChallenge.Description,
			Thumbnail:   userTask.TaskChallenge.Thumbnail,
			StartDate:   userTask.TaskChallenge.StartDate,
			EndDate:     userTask.TaskChallenge.EndDate,
			StatusTask:  userTask.TaskChallenge.Status,
			TaskSteps:   taskStep,
		},
		Images:    urlImages,
		UserSteps: userSteps,
	}

	for _, step := range userTask.UserTaskSteps {
		userSteps = append(userSteps, dto.DataUserSteps{
			Id:                  step.ID,
			UserTaskChallengeID: step.UserTaskChallengeID,
			TaskStepID:          step.TaskStepID,
			Completed:           step.Completed,
		})
	}

	for _, step := range userTask.TaskChallenge.TaskSteps {
		taskStep = append(taskStep, dto.TaskSteps{
			Id:          step.ID,
			Title:       step.Title,
			Description: step.Description,
		})
	}
	for _, image := range userTask.ImageTask {
		urlImages = append(urlImages, dto.Images{
			Images: image.ImageUrl,
		})
	}

	data.UserSteps = userSteps
	data.TaskChallenge.TaskSteps = taskStep
	data.Images = urlImages
	responseData := helper.ResponseData(http.StatusCreated, "success", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *UserTaskHandlerImpl) GetUserTaskDetailsHandler(c echo.Context) error {
	userId := c.Get("user").(*helper.JwtCustomClaims).UserID

	userTaskId := c.Param("userTaskId")

	userTask, images, err := handler.Usecase.GetUserTaskDetailsUsecase(userTaskId, userId)
	if err != nil {
		if errors.Is(err, pkg.ErrUserTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserTaskNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var dataImages []*dto.DataImages
	data := dto.GetUserTaskDetailsResponse{
		Id:          userTask.ID,
		TitleTask:   userTask.TaskChallenge.Title,
		UserName:    userTask.User.Name,
		Images:      []*dto.DataImages{},
		Description: userTask.DescriptionImage,
	}

	for _, image := range images {
		dataImages = append(dataImages, &dto.DataImages{
			Id:       image.ID,
			ImageUrl: image.ImageUrl,
		})
	}
	data.Images = dataImages
	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *UserTaskHandlerImpl) GetHistoryPointByUserIdHandler(c echo.Context) error {
	userId := c.Get("user").(*helper.JwtCustomClaims).UserID
	userTask, totalPoint, err := handler.Usecase.GetHistoryPointByUserIdUsecase(userId)
	if err != nil {
		if errors.Is(err, pkg.ErrUserNoHasTask) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserNoHasTask.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var dataHistoryPoints []*dto.DataHistoryPoint
	for _, task := range userTask {
		dataHistoryPoints = append(dataHistoryPoints, &dto.DataHistoryPoint{
			Id:         task.ID,
			TitleTask:  task.TaskChallenge.Title,
			Point:      task.Point,
			AcceptedAt: task.AcceptedAt,
		})
	}

	data := dto.HistoryPointResponse{
		TotalPoint: totalPoint,
		Data:       dataHistoryPoints,
	}

	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *UserTaskHandlerImpl) UpdateTaskStepHandler(c echo.Context) error {
	userId := c.Get("user").(*helper.JwtCustomClaims).UserID
	request := new(dto.UpdateTaskStepRequest)
	if err := c.Bind(request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body")
	}

	userTask, err := handler.Usecase.UpdateTaskStepUsecase(request, userId)
	if err != nil {
		if errors.Is(err, pkg.ErrUserTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserTaskNotFound.Error())
		}
		if errors.Is(err, pkg.ErrUserTaskDone) {
			return helper.ErrorHandler(c, http.StatusConflict, pkg.ErrUserTaskDone.Error())
		}
		if errors.Is(err, pkg.ErrTaskStepNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrTaskStepNotFound.Error())
		}
		if errors.Is(err, pkg.ErrStepNotInOrder) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrStepNotInOrder.Error())
		}
		if errors.Is(err, pkg.ErrUserTaskAlreadyApprove) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrUserTaskAlreadyApprove.Error())
		}
		if errors.Is(err, pkg.ErrUserTaskStepAlreadyCompleted) {
			return helper.ErrorHandler(c, http.StatusConflict, pkg.ErrUserTaskStepAlreadyCompleted.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var taskStep []dto.TaskSteps
	var userSteps []dto.DataUserSteps

	for _, step := range userTask.TaskChallenge.TaskSteps {
		taskStep = append(taskStep, dto.TaskSteps{
			Id:          step.ID,
			Title:       step.Title,
			Description: step.Description,
		})
	}

	for _, step := range userTask.UserTaskSteps {
		userSteps = append(userSteps, dto.DataUserSteps{
			Id:                  step.ID,
			UserTaskChallengeID: step.UserTaskChallengeID,
			TaskStepID:          step.TaskStepID,
			Completed:           step.Completed,
		})
	}
	data := dto.TaskChallengeResponseCreate{
		Id:          userTask.TaskChallenge.ID,
		Title:       userTask.TaskChallenge.Title,
		Description: userTask.TaskChallenge.Description,
		Thumbnail:   userTask.TaskChallenge.Thumbnail,
		StartDate:   userTask.TaskChallenge.StartDate,
		EndDate:     userTask.TaskChallenge.EndDate,
		Point:       userTask.TaskChallenge.Point,
		StatusTask:  userTask.TaskChallenge.Status,
		TaskSteps:   taskStep,
		UserSteps:   userSteps,
	}
	dataUsertask := dto.UserTaskResponseCreate{
		Id:             userTask.ID,
		StatusProgress: userTask.StatusProgress,
		TaskChalenge:   data,
	}
	responseData := helper.ResponseData(http.StatusCreated, "success", dataUsertask)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *UserTaskHandlerImpl) GetUserTaskByUserTaskIdHandler(c echo.Context) error {
	userId := c.Get("user").(*helper.JwtCustomClaims).UserID
	userTaskId := c.Param("userTaskId")
	userTask, err := handler.Usecase.GetUserTaskByUserTaskId(userId, userTaskId)
	if err != nil {
		if errors.Is(err, pkg.ErrUserTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserTaskNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	var taskStep []dto.TaskSteps
	var userSteps []dto.DataUserSteps

	for _, step := range userTask.TaskChallenge.TaskSteps {
		taskStep = append(taskStep, dto.TaskSteps{
			Id:          step.ID,
			Title:       step.Title,
			Description: step.Description,
		})
	}

	for _, step := range userTask.UserTaskSteps {
		userSteps = append(userSteps, dto.DataUserSteps{
			Id:                  step.ID,
			UserTaskChallengeID: step.UserTaskChallengeID,
			TaskStepID:          step.TaskStepID,
			Completed:           step.Completed,
		})
	}
	data := dto.DataGetUserTaskByUserTaskId{
		Id:          userTask.TaskChallenge.ID,
		Title:       userTask.TaskChallenge.Title,
		Description: userTask.TaskChallenge.Description,
		Thumbnail:   userTask.TaskChallenge.Thumbnail,
		StartDate:   userTask.TaskChallenge.StartDate,
		EndDate:     userTask.TaskChallenge.EndDate,
		Point:       userTask.TaskChallenge.Point,
		StatusTask:  userTask.TaskChallenge.Status,
		TaskSteps:   taskStep,
		UserSteps:   userSteps,
	}
	dataUsertask := dto.GetUserTaskByUserTaskIdResponse{
		Id:             userTask.ID,
		StatusProgress: userTask.StatusProgress,
		StatusAccept:   userTask.StatusAccept,
		Reason:         userTask.Reason,
		TaskChalenge:   data,
	}
	responseData := helper.ResponseData(http.StatusCreated, "success", dataUsertask)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *UserTaskHandlerImpl) GetUserTaskRejectedByUserIdHandler(c echo.Context) error {
	userId := c.Get("user").(*helper.JwtCustomClaims).UserID
	userTaskId := c.Param("userTaskId")
	userTask, err := handler.Usecase.GetUserTaskRejectedByUserId(userId, userTaskId)
	if err != nil {
		if errors.Is(err, pkg.ErrUserTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserTaskNotFound.Error())
		}
		if errors.Is(err, pkg.ErrUserTaskDone) {
			return helper.ErrorHandler(c, http.StatusConflict, pkg.ErrUserTaskDone.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail: "+err.Error())
	}
	var taskStep []dto.TaskSteps
	var userSteps []dto.DataUserSteps

	for _, step := range userTask.TaskChallenge.TaskSteps {
		taskStep = append(taskStep, dto.TaskSteps{
			Id:          step.ID,
			Title:       step.Title,
			Description: step.Description,
		})
	}

	for _, step := range userTask.UserTaskSteps {
		userSteps = append(userSteps, dto.DataUserSteps{
			Id:                  step.ID,
			UserTaskChallengeID: step.UserTaskChallengeID,
			TaskStepID:          step.TaskStepID,
			Completed:           step.Completed,
		})
	}
	dataTaskChallenges := dto.DataTaskChallenges{
		Id:          userTask.TaskChallenge.ID,
		Title:       userTask.TaskChallenge.Title,
		Description: userTask.TaskChallenge.Description,
		Thumbnail:   userTask.TaskChallenge.Thumbnail,
		StartDate:   userTask.TaskChallenge.StartDate,
		EndDate:     userTask.TaskChallenge.EndDate,
		StatusTask:  userTask.TaskChallenge.Status,
		TaskSteps:   taskStep,
	}

	dataUsertask := dto.UserTaskRejectResponse{
		Id:             userTask.ID,
		StatusProgress: userTask.StatusProgress,
		StatusAccept:   userTask.StatusAccept,
		Reason:         userTask.Reason,
		Point:          userTask.Point,
		TaskChalenge:   dataTaskChallenges,
		UserSteps:      userSteps,
	}
	responseData := helper.ResponseData(http.StatusCreated, "success", dataUsertask)
	return c.JSON(http.StatusOK, responseData)
}
