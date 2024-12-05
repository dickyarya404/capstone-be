package faq

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/faq"
	"github.com/sawalreverr/recything/internal/helper"
)

type faqHandler struct {
	faqUsecase faq.FaqUsecase
}

func NewFaqHandler(faqUsecase faq.FaqUsecase) faq.FaqHandler {
	return &faqHandler{faqUsecase: faqUsecase}
}

func (h *faqHandler) GetAllFaqs(c echo.Context) error {
	faqs, err := h.faqUsecase.GetAllFaqs()
	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", faqs)
}

func (h *faqHandler) GetFaqsByCategory(c echo.Context) error {
	category := c.QueryParam("name")

	faqs, err := h.faqUsecase.GetFaqsByCategory(category)
	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", faqs)
}

func (h *faqHandler) GetFaqsByKeyword(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	faqs, err := h.faqUsecase.GetFaqsByKeyword(keyword)
	if err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", faqs)
}
