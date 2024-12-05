package reminai

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	rai "github.com/sawalreverr/recything/internal/remin-ai"
)

type reminAIHandler struct {
	reminAIUsecase rai.ReminAIUsecase
}

func NewReminAIHandler(uc rai.ReminAIUsecase) rai.ReminAIHandler {
	return &reminAIHandler{reminAIUsecase: uc}
}

func (h *reminAIHandler) AskGPT(c echo.Context) error {
	var request rai.RequestInput

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	answer, err := h.reminAIUsecase.AskGPT(request)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	resp := rai.RequestOutput{
		Question: request.Question,
		AnswerAI: answer,
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", resp)
}
