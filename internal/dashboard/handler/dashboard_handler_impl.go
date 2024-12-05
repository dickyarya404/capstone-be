package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/dashboard/usecase"
	"github.com/sawalreverr/recything/internal/helper"
)

type DashboardHandlerImpl struct {
	DashboardUsecase usecase.DashboardUsecase
}

func NewDashboardHandler(dashboardUsecase usecase.DashboardUsecase) DashboardHandler {
	return &DashboardHandlerImpl{DashboardUsecase: dashboardUsecase}
}

func (h *DashboardHandlerImpl) GetDashboardHandler(c echo.Context) error {
	dashboard, err := h.DashboardUsecase.GetDashboardUsecase()
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	responseData := helper.ResponseData(http.StatusOK, "success", dashboard)
	return c.JSON(http.StatusOK, responseData)

}
