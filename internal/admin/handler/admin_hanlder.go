package handler

import "github.com/labstack/echo/v4"

type AdminHandler interface {
	AddAdminHandler(c echo.Context) error
	UpdateAdminHandler(c echo.Context) error
	GetDataAllAdminHandler(c echo.Context) error
	GetDataAdminByIdHandler(c echo.Context) error
	DeleteAdminHandler(c echo.Context) error
	GetProfileAdminHandler(c echo.Context) error
}
