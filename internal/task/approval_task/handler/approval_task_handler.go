package handler

import "github.com/labstack/echo/v4"

type ApprovalTaskHandler interface {
	GetAllApprovalTaskPaginationHandler(c echo.Context) error
	ApproveUserTaskHandler(c echo.Context) error
	RejectUserTaskHandler(c echo.Context) error
	GetUserTaskDetailsHandler(c echo.Context) error
}
