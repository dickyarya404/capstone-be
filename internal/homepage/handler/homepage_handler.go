package handler

import "github.com/labstack/echo/v4"

type HomePageHandler interface {
	GetHomepageHandler(c echo.Context) error
}
