package customdata

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	cdt "github.com/sawalreverr/recything/internal/custom-data"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type customDataHandler struct {
	customDataUsecase cdt.CustomDataUsecase
}

func NewCustomDataHandler(uc cdt.CustomDataUsecase) cdt.CustomDataHandler {
	return &customDataHandler{customDataUsecase: uc}
}

func (h *customDataHandler) NewCustomData(c echo.Context) error {
	var request cdt.CustomDataInput

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	newData, err := h.customDataUsecase.NewCustomData(request)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusCreated, "custom data created!", newData)
}

func (h *customDataHandler) GetDataByID(c echo.Context) error {
	dataID := c.Param("dataId")

	dataFound, err := h.customDataUsecase.FindByID(dataID)
	if err != nil {
		if errors.Is(err, pkg.ErrCustomDataNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", dataFound)
}

func (h *customDataHandler) GetAllData(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}
	sortBy := c.QueryParam("sort_by")
	sortType := c.QueryParam("sort_type")

	if sortBy == "" {
		sortBy = "created_at"
		sortType = "desc"
	}

	if sortType == "" {
		sortType = "asc"
	}

	datas, err := h.customDataUsecase.FindAll(page, limit, sortBy, sortType)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", datas)
}

func (h *customDataHandler) UpdateData(c echo.Context) error {
	var request cdt.CustomDataInput

	dataID := c.Param("dataId")

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := h.customDataUsecase.UpdateData(dataID, request); err != nil {
		if errors.Is(err, pkg.ErrCustomDataNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "custom data updated!", nil)
}

func (h *customDataHandler) DeleteData(c echo.Context) error {
	dataID := c.Param("dataId")

	if err := h.customDataUsecase.DeleteData(dataID); err != nil {
		if errors.Is(err, pkg.ErrCustomDataNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "custom data deleted!", nil)
}
