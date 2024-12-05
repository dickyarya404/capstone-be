package handler

import (
	"errors"
	"net/http"
	"sort"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/video/user_video/dto"
	"github.com/sawalreverr/recything/internal/video/user_video/usecase"
	"github.com/sawalreverr/recything/pkg"
)

type UserVideoHandlerImpl struct {
	Usecase usecase.UserVideoUsecase
}

func NewUserVideoHandler(usecase usecase.UserVideoUsecase) UserVideoHandler {
	return &UserVideoHandlerImpl{Usecase: usecase}
}

func (handler *UserVideoHandlerImpl) GetAllVideoHandler(c echo.Context) error {
	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "10"
	}
	limitInt, errLimit := strconv.Atoi(limit)
	if errLimit != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid limit")
	}
	videos, err := handler.Usecase.GetAllVideoUsecase(limitInt)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var dataVideo []*dto.DataVideoSearchByCategory
	for _, video := range *videos {
		uniqueContentCategories := make(map[uint]*dto.DataCategoryVideo)
		uniqueTrashCategories := make(map[uint]*dto.DataTrashCategoryVideo)

		for _, vc := range video.Categories {
			if _, exists := uniqueContentCategories[vc.ContentCategoryID]; !exists {
				uniqueContentCategories[vc.ContentCategoryID] = &dto.DataCategoryVideo{
					Id:   int(vc.ContentCategory.ID),
					Name: vc.ContentCategory.Name,
				}
			}
			if _, exists := uniqueTrashCategories[vc.WasteCategoryID]; !exists {
				uniqueTrashCategories[vc.WasteCategoryID] = &dto.DataTrashCategoryVideo{
					Id:   int(vc.WasteCategory.ID),
					Name: vc.WasteCategory.Name,
				}
			}
		}

		// Convert maps back to slices
		var videoCategories []*dto.DataCategoryVideo
		for _, vc := range uniqueContentCategories {
			videoCategories = append(videoCategories, vc)
		}
		var trashCategories []*dto.DataTrashCategoryVideo
		for _, tc := range uniqueTrashCategories {
			trashCategories = append(trashCategories, tc)
		}

		dataVideo = append(dataVideo, &dto.DataVideoSearchByCategory{
			Id:            video.ID,
			Title:         video.Title,
			Description:   video.Description,
			UrlThumbnail:  video.Thumbnail,
			LinkVideo:     video.Link,
			Viewer:        video.Viewer,
			VideoCategory: videoCategories,
			TrashCategory: trashCategories,
		})
	}

	responseData := dto.SearchVideoByCategoryVideoResponse{
		DataVideo: dataVideo,
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", responseData.DataVideo))
}

func (handler *UserVideoHandlerImpl) SearchVideoByKeywordHandler(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	videos, err := handler.Usecase.SearchVideoByKeywordUsecase(keyword)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var dataVideo []*dto.DataVideoSearchByCategory
	for _, video := range *videos {
		uniqueContentCategories := make(map[uint]*dto.DataCategoryVideo)
		uniqueTrashCategories := make(map[uint]*dto.DataTrashCategoryVideo)

		for _, vc := range video.Categories {
			if _, exists := uniqueContentCategories[vc.ContentCategoryID]; !exists {
				uniqueContentCategories[vc.ContentCategoryID] = &dto.DataCategoryVideo{
					Id:   int(vc.ContentCategory.ID),
					Name: vc.ContentCategory.Name,
				}
			}
			if _, exists := uniqueTrashCategories[vc.WasteCategoryID]; !exists {
				uniqueTrashCategories[vc.WasteCategoryID] = &dto.DataTrashCategoryVideo{
					Id:   int(vc.WasteCategory.ID),
					Name: vc.WasteCategory.Name,
				}
			}
		}

		// Convert maps back to slices
		var videoCategories []*dto.DataCategoryVideo
		for _, vc := range uniqueContentCategories {
			videoCategories = append(videoCategories, vc)
		}
		var trashCategories []*dto.DataTrashCategoryVideo
		for _, tc := range uniqueTrashCategories {
			trashCategories = append(trashCategories, tc)
		}

		dataVideo = append(dataVideo, &dto.DataVideoSearchByCategory{
			Id:            video.ID,
			Title:         video.Title,
			Description:   video.Description,
			UrlThumbnail:  video.Thumbnail,
			LinkVideo:     video.Link,
			Viewer:        video.Viewer,
			VideoCategory: videoCategories,
			TrashCategory: trashCategories,
		})
	}

	responseData := dto.SearchVideoByCategoryVideoResponse{
		DataVideo: dataVideo,
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", responseData.DataVideo))
}

func (handler *UserVideoHandlerImpl) SearchVideoByCategoryHandler(c echo.Context) error {
	categoryType := c.QueryParam("type")
	categoryName := c.QueryParam("name")
	videos, err := handler.Usecase.SearchVideoByCategoryUsecase(categoryType, categoryName)
	if err != nil {
		if errors.Is(err, pkg.ErrVideoNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrVideoNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	var dataVideo []*dto.DataVideoSearchByCategory
	for _, video := range *videos {
		uniqueContentCategories := make(map[uint]*dto.DataCategoryVideo)
		uniqueTrashCategories := make(map[uint]*dto.DataTrashCategoryVideo)

		for _, vc := range video.Categories {
			if _, exists := uniqueContentCategories[vc.ContentCategoryID]; !exists {
				uniqueContentCategories[vc.ContentCategoryID] = &dto.DataCategoryVideo{
					Id:   int(vc.ContentCategory.ID),
					Name: vc.ContentCategory.Name,
				}
			}
			if _, exists := uniqueTrashCategories[vc.WasteCategoryID]; !exists {
				uniqueTrashCategories[vc.WasteCategoryID] = &dto.DataTrashCategoryVideo{
					Id:   int(vc.WasteCategory.ID),
					Name: vc.WasteCategory.Name,
				}
			}
		}

		// Convert maps back to slices
		var videoCategories []*dto.DataCategoryVideo
		for _, vc := range uniqueContentCategories {
			videoCategories = append(videoCategories, vc)
		}
		var trashCategories []*dto.DataTrashCategoryVideo
		for _, tc := range uniqueTrashCategories {
			trashCategories = append(trashCategories, tc)
		}

		dataVideo = append(dataVideo, &dto.DataVideoSearchByCategory{
			Id:            video.ID,
			Title:         video.Title,
			Description:   video.Description,
			UrlThumbnail:  video.Thumbnail,
			LinkVideo:     video.Link,
			Viewer:        video.Viewer,
			VideoCategory: videoCategories,
			TrashCategory: trashCategories,
		})
	}

	responseData := dto.SearchVideoByCategoryVideoResponse{
		DataVideo: dataVideo,
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", responseData.DataVideo))
}

func (handler *UserVideoHandlerImpl) GetVideoDetailHandler(c echo.Context) error {
	id := c.Param("videoId")
	intId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid id parameter")
	}

	video, comments, totalComment, err := handler.Usecase.GetVideoDetailUsecase(intId)
	if err != nil {
		if errors.Is(err, pkg.ErrVideoNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrVideoNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	sortComments := c.QueryParam("sort-comments")

	if sortComments == "" {
		sortComments = "false"
	}
	if sortComments == "true" {
		sort.SliceStable(*comments, func(i, j int) bool {
			return (*comments)[i].CreatedAt.After((*comments)[j].CreatedAt)
		})
	}

	var dataComments []dto.DataComment
	data := dto.GetDetailsDataVideoByIdResponse{
		DataVideo: &dto.DataVideo{
			Id:           video.ID,
			Title:        video.Title,
			Description:  video.Description,
			UrlThumbnail: video.Thumbnail,
			LinkVideo:    video.Link,
			Viewer:       video.Viewer,
		},
		TotalComment: totalComment,
		Comments:     &dataComments,
	}

	for _, comment := range *comments {
		dataComments = append(dataComments, dto.DataComment{
			Id:          comment.ID,
			Comment:     comment.Comment,
			UserID:      comment.UserID,
			UserName:    comment.User.Name,
			UserProfile: comment.User.PictureURL,
			CreatedAt:   comment.CreatedAt,
		})
	}
	data.Comments = &dataComments
	responseData := helper.ResponseData(http.StatusOK, "success", data)

	return c.JSON(http.StatusOK, responseData)
}

func (handler *UserVideoHandlerImpl) AddCommentHandler(c echo.Context) error {
	var request dto.AddCommentRequest
	userId := c.Get("user").(*helper.JwtCustomClaims).UserID
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	if err := handler.Usecase.AddCommentUsecase(&request, userId); err != nil {
		if errors.Is(err, pkg.ErrVideoNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrVideoNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}
	responseData := helper.ResponseData(http.StatusOK, "success add comment", nil)
	return c.JSON(http.StatusCreated, responseData)
}
