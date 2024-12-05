package handler

import "github.com/labstack/echo/v4"

type UserTaskHandler interface {
	GetAllTasksHandler(c echo.Context) error
	GetTaskByIdHandler(c echo.Context) error
	CreateUserTaskHandler(c echo.Context) error
	UploadImageTaskHandler(c echo.Context) error
	GetUserTaskByUserIdHandler(c echo.Context) error
	GetUserTaskDoneByUserIdHandler(c echo.Context) error
	UpdateUserTaskHandler(c echo.Context) error
	GetUserTaskDetailsHandler(c echo.Context) error
	GetHistoryPointByUserIdHandler(c echo.Context) error
	UpdateTaskStepHandler(c echo.Context) error
	GetUserTaskByUserTaskIdHandler(c echo.Context) error
	GetUserTaskRejectedByUserIdHandler(c echo.Context) error
}
