package handler

import "github.com/labstack/echo/v4"

type ManageAchievementHandler interface {
	GetAllAchievementHandler(c echo.Context) error
	GetAchievementByIdHandler(c echo.Context) error
	UpdateAchievementHandler(c echo.Context) error
	DeleteAchievementHandler(c echo.Context) error
}
