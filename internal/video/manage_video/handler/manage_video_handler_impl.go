package handler

import (
	"encoding/json"
	"errors"
	"math"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/video/manage_video/dto"
	"github.com/sawalreverr/recything/internal/video/manage_video/usecase"
	"github.com/sawalreverr/recything/pkg"
)

type ManageVideoHandlerImpl struct {
	ManageVideoUsecase usecase.ManageVideoUsecase
}

func NewManageVideoHandlerImpl(manageVideoUsecase usecase.ManageVideoUsecase) ManageVideoHandler {
	return &ManageVideoHandlerImpl{ManageVideoUsecase: manageVideoUsecase}
}

func (handler *ManageVideoHandlerImpl) CreateDataVideoHandler(c echo.Context) error {
	var request dto.CreateDataVideoRequest

	json_data := c.FormValue("json_data")
	if err := json.Unmarshal([]byte(json_data), &request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	form, errForm := c.MultipartForm()
	if errForm != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, errForm.Error())
	}
	thumbnail := form.File["thumbnail"]
	if err := handler.ManageVideoUsecase.CreateDataVideoUseCase(&request, thumbnail); err != nil {
		if errors.Is(err, pkg.ErrVideoTitleAlreadyExist) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrVideoTitleAlreadyExist.Error())
		}
		if errors.Is(err, pkg.ErrVideoCategory) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrVideoCategory.Error())
		}
		if errors.Is(err, pkg.ErrVideoTrashCategory) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrVideoTrashCategory.Error())
		}
		if errors.Is(err, pkg.ErrNameCategoryVideoNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrNameCategoryVideoNotFound.Error())
		}
		if errors.Is(err, pkg.ErrNameTrashCategoryNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrNameTrashCategoryNotFound.Error())
		}
		if errors.Is(err, pkg.ErrNoVideoIdFoundOnUrl) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrNoVideoIdFoundOnUrl.Error())
		}
		if errors.Is(err, pkg.ErrVideoNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrVideoNotFound.Error())
		}
		if errors.Is(err, pkg.ErrVideoService) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrVideoService.Error())
		}
		if errors.Is(err, pkg.ErrApiYouTube) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrApiYouTube.Error())
		}
		if errors.Is(err, pkg.ErrParsingUrl) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrParsingUrl.Error())
		}
		if errors.Is(err, pkg.ErrThumbnail) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrThumbnail.Error())
		}
		if errors.Is(err, pkg.ErrThumbnailMaximum) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrThumbnailMaximum.Error())
		}
		if errors.Is(err, errors.New("upload image size must less than 2MB")) {
			return helper.ErrorHandler(c, http.StatusBadRequest, "upload image size must less than 2MB")
		}
		if errors.Is(err, errors.New("only image allowed")) {
			return helper.ErrorHandler(c, http.StatusBadRequest, "only image allowed")
		}
		if errors.Is(err, pkg.ErrUploadCloudinary) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrUploadCloudinary.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	return helper.ResponseHandler(c, http.StatusCreated, "success create data video", nil)
}

func (handler *ManageVideoHandlerImpl) GetAllCategoryVideoHandler(c echo.Context) error {
	videoCategories, trashCategories, err := handler.ManageVideoUsecase.GetAllCategoryVideoUseCase()
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	var dataVideoCategories []*dto.DataCategoryVideo
	var dataTrashCategories []*dto.DataTrashCategory
	data := &dto.GetAllCategoryVideoResponse{
		VideoCategory: dataVideoCategories,
		TrashCategory: dataTrashCategories,
	}

	for _, category := range videoCategories {
		dataVideoCategories = append(dataVideoCategories, &dto.DataCategoryVideo{
			Name: category,
		})
	}
	for _, category := range trashCategories {
		dataTrashCategories = append(dataTrashCategories, &dto.DataTrashCategory{
			Name: category,
		})
	}

	data.VideoCategory = dataVideoCategories
	data.TrashCategory = dataTrashCategories
	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *ManageVideoHandlerImpl) GetAllDataVideoPaginationHandler(c echo.Context) error {
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")

	if limit == "" {
		limit = "10"
	}
	if page == "" {
		page = "1"
	}

	limitInt, errLimit := strconv.Atoi(limit)
	if errLimit != nil || limitInt <= 0 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid limit parameter")
	}
	pageInt, errPage := strconv.Atoi(page)
	if errPage != nil || pageInt <= 0 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid page parameter")
	}
	videos, totalData, err := handler.ManageVideoUsecase.GetAllDataVideoPaginationUseCase(limitInt, pageInt)

	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	var dataVideos []*dto.DataVideo
	for _, video := range videos {
		dataVideos = append(dataVideos, &dto.DataVideo{
			Id:           video.ID,
			Title:        video.Title,
			Description:  video.Description,
			UrlThumbnail: video.Thumbnail,
		})
	}

	data := &dto.GetAllDataVideoPaginationResponse{
		Code:      http.StatusOK,
		Message:   "success",
		Data:      dataVideos,
		Page:      pageInt,
		Limit:     limitInt,
		TotalData: totalData,
		TotalPage: int(math.Ceil(float64(totalData) / float64(limitInt))),
	}

	return c.JSON(http.StatusOK, data)
}

func (handler *ManageVideoHandlerImpl) GetDetailsDataVideoByIdHandler(c echo.Context) error {
	id := c.Param("videoId")
	idInt, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid id parameter")
	}

	video, err := handler.ManageVideoUsecase.GetDetailsDataVideoByIdUseCase(idInt)
	if err != nil {
		if errors.Is(err, pkg.ErrVideoNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrVideoNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	uniqueContentCategories := make(map[uint]*dto.DataCategoryVideoResponse)
	uniqueWasteCategories := make(map[uint]*dto.DataTrashCategoryResponse)

	for _, category := range video.Categories {
		if category.ContentCategoryID != 0 {
			if _, exists := uniqueContentCategories[category.ContentCategoryID]; !exists {
				uniqueContentCategories[category.ContentCategoryID] = &dto.DataCategoryVideoResponse{
					Id:   int(category.ContentCategory.ID),
					Name: category.ContentCategory.Name,
				}
			}
		}
		if category.WasteCategoryID != 0 {
			if _, exists := uniqueWasteCategories[category.WasteCategoryID]; !exists {
				uniqueWasteCategories[category.WasteCategoryID] = &dto.DataTrashCategoryResponse{
					Id:   int(category.WasteCategory.ID),
					Name: category.WasteCategory.Name,
				}
			}
		}
	}

	var dataContentCategories []*dto.DataCategoryVideoResponse
	for _, vc := range uniqueContentCategories {
		dataContentCategories = append(dataContentCategories, vc)
	}
	var dataWasteCategories []*dto.DataTrashCategoryResponse
	for _, wc := range uniqueWasteCategories {
		dataWasteCategories = append(dataWasteCategories, wc)
	}

	dataVideo := &dto.GetDetailsDataVideoByIdResponse{
		Id:            video.ID,
		Title:         video.Title,
		Description:   video.Description,
		UrlThumbnail:  video.Thumbnail,
		LinkVideo:     video.Link,
		Viewer:        video.Viewer,
		VideoCategory: dataContentCategories,
		TrashCategory: dataWasteCategories,
	}

	responseData := helper.ResponseData(http.StatusOK, "success", dataVideo)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *ManageVideoHandlerImpl) UpdateDataVideoHandler(c echo.Context) error {
	id := c.Param("videoId")
	idInt, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid id parameter")
	}
	var request dto.UpdateDataVideoRequest
	json_data := c.FormValue("json_data")
	if json_data != "" {
		if err := json.Unmarshal([]byte(json_data), &request); err != nil {
			return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
		}
	}

	form, errForm := c.MultipartForm()
	if errForm != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, errForm.Error())
	}
	var thumbnail []*multipart.FileHeader
	if form != nil {
		thumbnail = form.File["thumbnail"]
	}

	if err := handler.ManageVideoUsecase.UpdateDataVideoUseCase(&request, thumbnail, idInt); err != nil {
		if errors.Is(err, pkg.ErrNameCategoryVideoNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrNameCategoryVideoNotFound.Error())
		}
		if errors.Is(err, pkg.ErrNameTrashCategoryNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrNameTrashCategoryNotFound.Error())
		}
		if errors.Is(err, pkg.ErrNoVideoIdFoundOnUrl) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrNoVideoIdFoundOnUrl.Error())
		}
		if errors.Is(err, pkg.ErrVideoNotFound) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrVideoNotFound.Error())
		}
		if errors.Is(err, pkg.ErrVideoService) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrVideoService.Error())
		}
		if errors.Is(err, pkg.ErrApiYouTube) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrApiYouTube.Error())
		}
		if errors.Is(err, pkg.ErrParsingUrl) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrParsingUrl.Error())
		}
		if errors.Is(err, pkg.ErrThumbnail) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrThumbnail.Error())
		}
		if errors.Is(err, pkg.ErrThumbnailMaximum) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrThumbnailMaximum.Error())
		}
		if errors.Is(err, errors.New("upload image size must less than 2MB")) {
			return helper.ErrorHandler(c, http.StatusBadRequest, "upload image size must less than 2MB")
		}
		if errors.Is(err, errors.New("only image allowed")) {
			return helper.ErrorHandler(c, http.StatusBadRequest, "only image allowed")
		}
		if errors.Is(err, pkg.ErrUploadCloudinary) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrUploadCloudinary.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	return helper.ResponseHandler(c, http.StatusOK, "success update data video", nil)
}

func (handler *ManageVideoHandlerImpl) DeleteDataVideoHandler(c echo.Context) error {
	id := c.Param("videoId")
	idInt, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid id parameter")
	}
	if err := handler.ManageVideoUsecase.DeleteDataVideoUseCase(idInt); err != nil {
		if errors.Is(err, pkg.ErrVideoNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrVideoNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	return helper.ResponseHandler(c, http.StatusOK, "success delete data video", nil)
}
