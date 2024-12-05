package handler

import "github.com/labstack/echo/v4"

type ManageTaskHandler interface {
	CreateTaskHandler(c echo.Context) error
	GetTaskChallengePaginationHandler(c echo.Context) error
	GetTaskByIdHandler(c echo.Context) error
	UpdateTaskHandler(c echo.Context) error
	DeleteTaskHandler(c echo.Context) error
}
