package handler

import "github.com/labstack/echo/v4"

type ManageVideoHandler interface {
	CreateDataVideoHandler(c echo.Context) error
	GetAllCategoryVideoHandler(c echo.Context) error
	GetAllDataVideoPaginationHandler(c echo.Context) error
	GetDetailsDataVideoByIdHandler(c echo.Context) error
	UpdateDataVideoHandler(c echo.Context) error
	DeleteDataVideoHandler(c echo.Context) error
}
