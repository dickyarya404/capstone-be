package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/task/approval_task/dto"
	"github.com/sawalreverr/recything/internal/task/approval_task/usecase"
	"github.com/sawalreverr/recything/pkg"
)

type ApprovalTaskHandlerImpl struct {
	usecase usecase.ApprovalTaskUsecase
}

func NewApprovalTaskHandler(usecase usecase.ApprovalTaskUsecase) *ApprovalTaskHandlerImpl {
	return &ApprovalTaskHandlerImpl{usecase: usecase}
}

func (handler *ApprovalTaskHandlerImpl) GetAllApprovalTaskPaginationHandler(c echo.Context) error {
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	if limit == "" {
		limit = "10"
	}
	if page == "" {
		page = "1"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return err
	}

	userTask, total, err := handler.usecase.GetAllApprovalTaskPaginationUseCase(limitInt, pageInt)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var data []dto.DataUserTask
	for _, task := range userTask {
		data = append(data, dto.DataUserTask{
			Id:           task.ID,
			StatusAccept: task.StatusAccept,
			Point:        task.Point,
			TaskChallenge: dto.DataTasks{
				Id:        task.TaskChallenge.ID,
				Title:     task.TaskChallenge.Title,
				StartDate: task.TaskChallenge.StartDate,
				EndDate:   task.TaskChallenge.EndDate,
			},
			User: dto.DataUser{
				Id:      task.User.ID,
				Name:    task.User.Name,
				Profile: task.User.PictureURL,
			},
		})
	}

	totalPage := total / limitInt
	if total%limitInt != 0 {
		totalPage++
	}
	responseDataPagination := dto.GetUserTaskPagination{
		Code:      http.StatusOK,
		Message:   "success get all task",
		Data:      data,
		Page:      pageInt,
		Limit:     limitInt,
		TotalData: total,
		TotalPage: totalPage,
	}
	responseData := helper.ResponseData(http.StatusOK, "success get all user task", responseDataPagination)

	return c.JSON(http.StatusOK, responseData)

}

func (handler *ApprovalTaskHandlerImpl) ApproveUserTaskHandler(c echo.Context) error {
	userTaskId := c.Param("userTaskId")
	if err := handler.usecase.ApproveUserTaskUseCase(userTaskId); err != nil {
		if errors.Is(err, pkg.ErrUserTaskAlreadyApprove) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrUserTaskAlreadyApprove.Error())
		}
		if errors.Is(err, pkg.ErrUserTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserTaskNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	responseData := helper.ResponseData(http.StatusOK, "success approve user task", nil)

	return c.JSON(http.StatusOK, responseData)
}

func (handler *ApprovalTaskHandlerImpl) RejectUserTaskHandler(c echo.Context) error {
	userTaskId := c.Param("userTaskId")
	var request dto.RejectUserTaskRequest
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body")
	}
	if err := handler.usecase.RejectUserTaskUseCase(&request, userTaskId); err != nil {
		if errors.Is(err, pkg.ErrUserTaskAlreadyReject) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrUserTaskAlreadyReject.Error())
		}
		if errors.Is(err, pkg.ErrUserTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserTaskNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	responseData := helper.ResponseData(http.StatusOK, "success reject user task", nil)

	return c.JSON(http.StatusOK, responseData)
}

func (handler *ApprovalTaskHandlerImpl) GetUserTaskDetailsHandler(c echo.Context) error {
	userTaskId := c.Param("userTaskId")
	task, images, err := handler.usecase.GetUserTaskDetailsUseCase(userTaskId)
	if err != nil {
		if errors.Is(err, pkg.ErrUserTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserTaskNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var dataImages []*dto.DataImages
	data := dto.GetUserTaskDetailsResponse{
		Id:          task.ID,
		TitleTask:   task.TaskChallenge.Title,
		StartDate:   task.TaskChallenge.StartDate,
		EndDate:     task.TaskChallenge.EndDate,
		UserName:    task.User.Name,
		Images:      []*dto.DataImages{},
		Description: task.DescriptionImage,
	}

	for _, image := range images {
		dataImages = append(dataImages, &dto.DataImages{
			Id:         image.ID,
			ImageUrl:   image.ImageUrl,
			UploadedAt: image.CreatedAt,
		})
	}

	data.Images = dataImages

	responseData := helper.ResponseData(http.StatusOK, "success get user task details", &data)

	return c.JSON(http.StatusOK, responseData)
}
