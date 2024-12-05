package handler

import "github.com/labstack/echo/v4"

type UserAchievementHandler interface {
	GetAvhievementsByUserhandler(c *echo.Context) error
}
