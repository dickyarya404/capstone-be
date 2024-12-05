package reminai

import "github.com/labstack/echo/v4"

type ReminAIUsecase interface {
	AskGPT(question RequestInput) (string, error)
}

type ReminAIHandler interface {
	AskGPT(c echo.Context) error
}
