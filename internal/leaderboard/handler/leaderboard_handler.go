package handler

import "github.com/labstack/echo/v4"

type LeaderboardHandler interface {
	GetLeaderboardHandler(c echo.Context) error
}
