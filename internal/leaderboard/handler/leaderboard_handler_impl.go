package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/leaderboard/usecase"
)

type LeaderboardHandlerImpl struct {
	LeaderboardUsecase usecase.LeaderboardUsecase
}

func NewLeaderboardHandler(leaderboardUsecase usecase.LeaderboardUsecase) LeaderboardHandler {
	return LeaderboardHandlerImpl{LeaderboardUsecase: leaderboardUsecase}
}

func (handler LeaderboardHandlerImpl) GetLeaderboardHandler(c echo.Context) error {
	leaderboard, err := handler.LeaderboardUsecase.GetLeaderboardUsecase()
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}

	responseData := helper.ResponseData(http.StatusOK, "data successfully retrieved", leaderboard.DataLeaderboard)
	return c.JSON(http.StatusOK, responseData)
}
