package usecase

import (
	"mime/multipart"
	"strings"

	art "github.com/sawalreverr/recything/internal/article"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/video/manage_video/dto"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
	repository "github.com/sawalreverr/recything/internal/video/manage_video/repository"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type ManageVideoUsecaseImpl struct {
	manageVideoRepository repository.ManageVideoRepository
}

func NewManageVideoUsecaseImpl(manageVideoRepository repository.ManageVideoRepository) *ManageVideoUsecaseImpl {
	return &ManageVideoUsecaseImpl{
		manageVideoRepository: manageVideoRepository,
	}
}

func (usecase *ManageVideoUsecaseImpl) CreateDataVideoUseCase(request *dto.CreateDataVideoRequest, thumbnail []*multipart.FileHeader) error {
	if len(thumbnail) == 0 {
		return pkg.ErrThumbnail
	}
	if len(request.ContentCategory) == 0 {
		return pkg.ErrVideoCategory
	}
	if len(request.WasteCategory) == 0 {
		return pkg.ErrVideoTrashCategory
	}
	if len(thumbnail) > 1 {
		return pkg.ErrThumbnailMaximum
	}
	validImages, errImages := helper.ImagesValidation(thumbnail)
	if errImages != nil {
		return errImages
	}

	if err := usecase.manageVideoRepository.FindTitleVideo(request.Title); err == nil {
		return pkg.ErrVideoTitleAlreadyExist
	}

	var contentCategories []*art.ContentCategory
	var wasteCategories []*art.WasteCategory

	for _, category := range request.ContentCategory {
		name := strings.ToLower(category.Name)
		content, err := usecase.manageVideoRepository.FindNameCategoryVideo(name)
		if err != nil {
			return pkg.ErrVideoCategory
		}
		contentCategories = append(contentCategories, content)
	}

	for _, category := range request.WasteCategory {
		name := strings.ToLower(category.Name)
		waste, err := usecase.manageVideoRepository.FindNamaTrashCategory(name)
		if err != nil {
			return pkg.ErrVideoTrashCategory
		}
		wasteCategories = append(wasteCategories, waste)
	}

	var videoCategories []video.VideoCategory
	for _, content := range contentCategories {
		for _, waste := range wasteCategories {
			videoCategories = append(videoCategories, video.VideoCategory{
				ContentCategoryID: content.ID,
				WasteCategoryID:   waste.ID,
			})
		}
	}

	view, errGetView := helper.GetVideoViewCount(request.LinkVideo)
	if errGetView != nil {
		return errGetView
	}
	urlThumbnail, errUpload := helper.UploadToCloudinary(validImages[0], "video_thumbnail")
	if errUpload != nil {
		return pkg.ErrUploadCloudinary
	}
	intView := int(view)
	videos := video.Video{
		Title:       request.Title,
		Description: request.Description,
		Thumbnail:   urlThumbnail,
		Link:        request.LinkVideo,
		Viewer:      intView,
		Categories:  videoCategories,
	}

	_, errVideo := usecase.manageVideoRepository.CreateVideoAndCategories(&videos)
	if errVideo != nil {
		return errVideo
	}

	return nil
}

func (usecase *ManageVideoUsecaseImpl) GetAllCategoryVideoUseCase() ([]string, []string, error) {
	videoCategories, errvidCategory := usecase.manageVideoRepository.GetAllCategoryVideo()
	if errvidCategory != nil {
		return nil, nil, errvidCategory
	}
	contentCatgory, errTrashCategory := usecase.manageVideoRepository.GetAllTrashCategoryVideo()
	if errTrashCategory != nil {
		return nil, nil, errTrashCategory
	}
	return videoCategories, contentCatgory, nil
}

func (usecase *ManageVideoUsecaseImpl) GetAllDataVideoPaginationUseCase(limit int, page int) ([]video.Video, int, error) {
	videos, count, err := usecase.manageVideoRepository.GetAllDataVideoPagination(limit, page)
	if err != nil {
		return nil, 0, err
	}
	return videos, count, nil
}

func (usecase *ManageVideoUsecaseImpl) GetDetailsDataVideoByIdUseCase(id int) (*video.Video, error) {
	video, err := usecase.manageVideoRepository.GetDetailsDataVideoById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, pkg.ErrVideoNotFound
		}
		return nil, err
	}
	return video, nil
}

func (usecase *ManageVideoUsecaseImpl) UpdateDataVideoUseCase(request *dto.UpdateDataVideoRequest, thumbnail []*multipart.FileHeader, id int) error {
	if len(thumbnail) > 1 {
		return pkg.ErrThumbnailMaximum
	}

	dataVideo, err := usecase.manageVideoRepository.GetDetailsDataVideoById(id)
	if err != nil {
		return pkg.ErrVideoNotFound
	}

	var contentCategories []*art.ContentCategory
	var wasteCategories []*art.WasteCategory

	if request.ContentCategories != nil {
		for _, category := range request.ContentCategories {
			name := strings.ToLower(category.Name)
			content, err := usecase.manageVideoRepository.FindNameCategoryVideo(name)
			if err != nil {
				return pkg.ErrVideoCategory
			}
			contentCategories = append(contentCategories, content)
		}
	}

	if request.WasteCategories != nil {
		for _, category := range request.WasteCategories {
			name := strings.ToLower(category.Name)
			waste, err := usecase.manageVideoRepository.FindNamaTrashCategory(name)
			if err != nil {
				return pkg.ErrVideoTrashCategory
			}
			wasteCategories = append(wasteCategories, waste)
		}

	}

	var videoCategories []video.VideoCategory
	for _, content := range contentCategories {
		for _, waste := range wasteCategories {
			videoCategories = append(videoCategories, video.VideoCategory{
				ContentCategoryID: content.ID,
				WasteCategoryID:   waste.ID,
			})
		}
	}

	dataVideo.Categories = videoCategories

	var urlThumbnail string
	if len(thumbnail) == 1 {
		validImages, errImages := helper.ImagesValidation(thumbnail)
		if errImages != nil {
			return errImages
		}
		urlThumbnailUpload, errUpload := helper.UploadToCloudinary(validImages[0], "video_thumbnail_update")
		if errUpload != nil {
			return pkg.ErrUploadCloudinary
		}
		urlThumbnail = urlThumbnailUpload
	}

	if request.Title != "" {
		dataVideo.Title = request.Title
	}
	if request.Description != "" {
		dataVideo.Description = request.Description
	}
	if urlThumbnail != "" {
		dataVideo.Thumbnail = urlThumbnail
	}
	if request.LinkVideo != "" {
		view, errGetView := helper.GetVideoViewCount(request.LinkVideo)
		if errGetView != nil {
			return errGetView
		}
		if view != 0 {
			intView := int(view)
			dataVideo.Viewer = intView
		}
		dataVideo.Link = request.LinkVideo
	}

	if err := usecase.manageVideoRepository.UpdateDataVideo(dataVideo, id); err != nil {
		return err
	}
	return nil
}

func (usecase *ManageVideoUsecaseImpl) DeleteDataVideoUseCase(id int) error {
	if _, err := usecase.manageVideoRepository.GetDetailsDataVideoById(id); err != nil {
		return pkg.ErrVideoNotFound
	}
	if err := usecase.manageVideoRepository.DeleteDataVideo(id); err != nil {
		return err
	}
	return nil
}
