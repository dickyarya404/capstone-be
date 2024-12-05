package aboutus

import (
	"net/http"

	"github.com/labstack/echo/v4"
	au "github.com/sawalreverr/recything/internal/about-us"
	"github.com/sawalreverr/recything/internal/helper"
)

type aboutUsHandler struct {
	AboutUsUsecase au.AboutUsUsecase
}

func NewAboutUsHandler(uc au.AboutUsUsecase) au.AboutUsHandler {
	return &aboutUsHandler{AboutUsUsecase: uc}
}

func (h *aboutUsHandler) GetAboutUsByCategory(c echo.Context) error {
	categoryName := c.QueryParam("name")

	response, err := h.AboutUsUsecase.GetAboutUsByCategory(categoryName)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", response)
}
