package handler

import (
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/task/manage_task/dto"
	"github.com/sawalreverr/recything/internal/task/manage_task/usecase"
	"github.com/sawalreverr/recything/pkg"
)

type ManageTaskHandlerImpl struct {
	Usecase usecase.ManageTaskUsecase
}

func NewManageTaskHandler(usecase usecase.ManageTaskUsecase) ManageTaskHandler {
	return &ManageTaskHandlerImpl{Usecase: usecase}
}

func (handler *ManageTaskHandlerImpl) CreateTaskHandler(c echo.Context) error {
	claims := c.Get("user").(*helper.JwtCustomClaims)
	var request dto.CreateTaskResquest
	json_data := c.FormValue("json_data")
	if err := json.Unmarshal([]byte(json_data), &request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	form, errForm := c.MultipartForm()
	if errForm != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, errForm.Error())
	}
	thumbnail := form.File["thumbnail"]

	taskChallange, err := handler.Usecase.CreateTaskUsecase(&request, thumbnail, claims.UserID)

	if err != nil {
		if errors.Is(err, pkg.ErrTaskStepsNull) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrTaskStepsNull.Error())
		}
		if errors.Is(err, pkg.ErrParsedTime) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrParsedTime.Error())
		}
		if errors.Is(err, pkg.ErrThumbnail) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrThumbnail.Error())
		}
		if errors.Is(err, pkg.ErrThumbnailMaximum) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrThumbnailMaximum.Error())
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
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	taskStep := []dto.TaskSteps{}

	data := dto.CreateTaskResponse{
		Id:          taskChallange.ID,
		Title:       taskChallange.Title,
		Description: taskChallange.Description,
		Thumbnail:   taskChallange.Thumbnail,
		StartDate:   taskChallange.StartDate,
		EndDate:     taskChallange.EndDate,
		Point:       taskChallange.Point,
		Status:      taskChallange.Status,
		Steps:       taskStep,
	}
	for _, step := range taskChallange.TaskSteps {
		taskSteps := dto.TaskSteps{
			Id:          step.ID,
			Title:       step.Title,
			Description: step.Description,
		}
		taskStep = append(taskStep, taskSteps)
	}
	data.Steps = taskStep

	responseData := helper.ResponseData(http.StatusCreated, "success", data)
	return c.JSON(http.StatusOK, responseData)

}

func (handler *ManageTaskHandlerImpl) GetTaskChallengePaginationHandler(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	status := c.QueryParam("status")
	endDate := c.QueryParam("end-date")
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}
	limitInt, errLimit := strconv.Atoi(limit)
	if errLimit != nil || limitInt <= 0 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid limit parameter")
	}
	pageInt, errPage := strconv.Atoi(page)
	if errPage != nil || pageInt <= 0 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid page parameter")
	}
	if status != "" {
		if status != "true" && status != "false" {
			return helper.ErrorHandler(c, http.StatusBadRequest, "invalid status parameter")
		}
	}

	if endDate != "" {
		if endDate != "asc" && endDate != "desc" {
			return helper.ErrorHandler(c, http.StatusBadRequest, "invalid end date parameter")
		}
	}

	tasks, totalData, err := handler.Usecase.GetTaskChallengePagination(pageInt, limitInt, status, endDate)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var data []dto.DataTasks
	for _, task := range tasks {
		var taskSteps []dto.TaskSteps
		for _, step := range task.TaskSteps {
			taskSteps = append(taskSteps, dto.TaskSteps{
				Id:          step.ID,
				Title:       step.Title,
				Description: step.Description,
			})
		}
		data = append(data, dto.DataTasks{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Thumbnail:   task.Thumbnail,
			StartDate:   task.StartDate,
			EndDate:     task.EndDate,
			Point:       task.Point,
			Status:      task.Status,
			Steps:       taskSteps,
			TaskCreator: dto.TaskCreatorAdmin{
				Id:   task.AdminId,
				Name: task.Admin.Name,
			},
		})
	}

	totalPage := totalData / limitInt
	if totalData%limitInt != 0 {
		totalPage++
	}

	responseDataPagination := dto.GetTaskPagination{
		Code:      http.StatusOK,
		Message:   "success",
		Data:      data,
		Page:      pageInt,
		Limit:     limitInt,
		TotalData: totalData,
		TotalPage: totalPage,
	}

	return c.JSON(http.StatusOK, responseDataPagination)
}

func (handler *ManageTaskHandlerImpl) GetTaskByIdHandler(c echo.Context) error {
	id := c.Param("taskId")
	task, err := handler.Usecase.GetTaskByIdUsecase(id)
	if err != nil {
		if errors.Is(err, pkg.ErrTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrTaskNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var taskSteps []dto.TaskSteps
	for _, step := range task.TaskSteps {
		taskSteps = append(taskSteps, dto.TaskSteps{
			Id:          step.ID,
			Title:       step.Title,
			Description: step.Description,
		})
	}
	data := dto.TaskGetByIdResponse{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Thumbnail:   task.Thumbnail,
		StartDate:   task.StartDate,
		EndDate:     task.EndDate,
		Point:       task.Point,
		Status:      task.Status,
		Steps:       taskSteps,
		TaskCreator: dto.TaskCreatorAdmin{
			Id:   task.AdminId,
			Name: task.Admin.Name,
		},
	}
	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *ManageTaskHandlerImpl) UpdateTaskHandler(c echo.Context) error {
	var request dto.UpdateTaskRequest
	id := c.Param("taskId")
	json_data := c.FormValue("json_data")
	if json_data != "" {
		if err := json.Unmarshal([]byte(json_data), &request); err != nil {
			return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
		}
	}
	form, errForm := c.MultipartForm()
	if errForm != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, errForm.Error())
	}
	var thumbnail []*multipart.FileHeader
	if form != nil {
		thumbnail = form.File["thumbnail"]
	}

	task, err := handler.Usecase.UpdateTaskChallengeUsecase(&request, thumbnail, id)
	if err != nil {
		if errors.Is(err, pkg.ErrTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrTaskNotFound.Error())
		}
		if errors.Is(err, pkg.ErrTaskStepsNull) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrTaskStepsNull.Error())
		}
		if errors.Is(err, pkg.ErrParsedTime) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrParsedTime.Error())
		}
		if errors.Is(err, pkg.ErrThumbnail) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrThumbnail.Error())
		}
		if errors.Is(err, pkg.ErrThumbnailMaximum) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrThumbnailMaximum.Error())
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
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	var taskSteps []dto.TaskSteps
	for _, step := range task.TaskSteps {
		taskSteps = append(taskSteps, dto.TaskSteps{
			Id:          step.ID,
			Title:       step.Title,
			Description: step.Description,
		})
	}
	data := dto.UpdateTaskResponse{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Thumbnail:   task.Thumbnail,
		StartDate:   task.StartDate,
		EndDate:     task.EndDate,
		Point:       task.Point,
		Steps:       taskSteps,
	}
	responseData := helper.ResponseData(http.StatusOK, "data updated successfully", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *ManageTaskHandlerImpl) DeleteTaskHandler(c echo.Context) error {
	id := c.Param("taskId")
	err := handler.Usecase.DeleteTaskChallengeUsecase(id)
	if err != nil {
		if errors.Is(err, pkg.ErrTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrTaskNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	responseData := helper.ResponseData(http.StatusOK, "data deleted successfully", nil)
	return c.JSON(http.StatusOK, responseData)
}
