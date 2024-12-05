package report

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	rpt "github.com/sawalreverr/recything/internal/report"
	"github.com/sawalreverr/recything/pkg"
)

type reportHandler struct {
	reportUsecase rpt.ReportUsecase
}

func NewReportHandler(usecase rpt.ReportUsecase) rpt.ReportHandler {
	return &reportHandler{reportUsecase: usecase}
}

// for user
func (h *reportHandler) NewReport(c echo.Context) error {
	var request rpt.ReportInput

	authorID := c.Get("user").(*helper.JwtCustomClaims).UserID

	jsonData := c.FormValue("json_data")
	if err := json.Unmarshal([]byte(jsonData), &request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	form, _ := c.MultipartForm()
	imageFiles := form.File["images"]

	validImages, err := helper.ImagesValidation(imageFiles)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	var imageURLs []string
	for _, file := range validImages {
		resultURL, err := helper.UploadToCloudinary(file, "recything/reports")
		if err != nil {
			helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrUploadCloudinary.Error())
		}
		imageURLs = append(imageURLs, resultURL)
	}

	newReport, err := h.reportUsecase.CreateReport(request, authorID, imageURLs)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusCreated, "report created!", newReport)
}

func (h *reportHandler) GetHistoryUserReports(c echo.Context) error {
	authorID := c.Get("user").(*helper.JwtCustomClaims).UserID

	reports, err := h.reportUsecase.FindHistoryUserReports(authorID)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", reports)
}

// for admin
func (h *reportHandler) UpdateStatus(c echo.Context) error {
	var request rpt.UpdateStatus

	reportID := c.Param("reportId")

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := h.reportUsecase.UpdateStatusReport(request, reportID); err != nil {
		if errors.Is(err, pkg.ErrReportNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "report status updated!", nil)
}

func (h *reportHandler) GetAllReports(c echo.Context) error {
	var date time.Time
	var err error

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}

	reportType := c.QueryParam("report_type")
	status := c.QueryParam("status")
	dateParam := c.QueryParam("date")

	if dateParam != "" {
		date, err = time.Parse("2006-01-02", dateParam)
		if err != nil {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrDateFormat.Error())
		}
	} else {
		date = time.Time{}
	}

	reportDetails, total, err := h.reportUsecase.FindAllReports(page, limit, reportType, status, date)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	response := rpt.ReportResponsePagination{
		Total:  total,
		Page:   page,
		Limit:  limit,
		Report: *reportDetails,
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", response)
}
