package handler

import "github.com/labstack/echo/v4"

type DashboardHandler interface {
	GetDashboardHandler(c echo.Context) error
}
