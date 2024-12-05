package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/achievements/user_achievements/usecase"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type UserAchievementHandlerImpl struct {
	UserAchievementUsecase usecase.UserAchievementUsecase
}

func NewUserAchievementHandler(userAchievementUsecase usecase.UserAchievementUsecase) *UserAchievementHandlerImpl {
	return &UserAchievementHandlerImpl{UserAchievementUsecase: userAchievementUsecase}
}

func (handler *UserAchievementHandlerImpl) GetAvhievementsByUserhandler(c echo.Context) error {
	userId := c.Get("user").(*helper.JwtCustomClaims).UserID

	data, err := handler.UserAchievementUsecase.GetAvhievementsByUser(userId)
	if err != nil {
		if errors.Is(err, pkg.ErrUserNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserNotFound.Error())
		}
		if errors.Is(err, pkg.ErrUserNotHasHistoryPoint) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserNotHasHistoryPoint.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail: "+err.Error())
	}

	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}
