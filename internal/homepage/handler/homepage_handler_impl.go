package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/homepage/usecase"
	"github.com/sawalreverr/recything/pkg"
)

type HomePageHandlerImpl struct {
	homepageUsecase usecase.HomepageUsecase
}

func NewHomePageHandler(homepageUsecase usecase.HomepageUsecase) HomePageHandler {
	return &HomePageHandlerImpl{homepageUsecase: homepageUsecase}
}

func (handler *HomePageHandlerImpl) GetHomepageHandler(c echo.Context) error {
	userId := c.Get("user").(*helper.JwtCustomClaims).UserID
	homepageResponse, err := handler.homepageUsecase.GetHomepageUsecase(userId)
	if err != nil {
		if errors.Is(err, pkg.ErrUserNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}

	responseData := helper.ResponseData(http.StatusOK, "ok", &homepageResponse)
	return c.JSON(http.StatusOK, responseData)

}
