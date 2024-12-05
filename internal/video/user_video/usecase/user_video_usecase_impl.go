package usecase

import (
	"log"

	"github.com/sawalreverr/recything/internal/helper"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
	"github.com/sawalreverr/recything/internal/video/user_video/dto"
	"github.com/sawalreverr/recything/internal/video/user_video/repository"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type UserVideoUsecaseImpl struct {
	Repository repository.UserVideoRepository
}

func NewUserVideoUsecase(repository repository.UserVideoRepository) *UserVideoUsecaseImpl {
	return &UserVideoUsecaseImpl{Repository: repository}
}

func (usecase *UserVideoUsecaseImpl) GetAllVideoUsecase(limit int) (*[]video.Video, error) {
	videos, err := usecase.Repository.GetAllVideo(limit)
	if err != nil {
		return nil, err
	}

	for i := range *videos {
		view, errGetView := helper.GetVideoViewCount((*videos)[i].Link)
		if errGetView != nil {
			log.Printf("failed to get view count for video %d: %v", (*videos)[i].ID, errGetView)
			continue
		}
		(*videos)[i].Viewer = int(view)
		if errUpdate := usecase.Repository.UpdateViewer(int(view), (*videos)[i].ID); errUpdate != nil {
			log.Printf("failed to update viewer count for video %d: %v", (*videos)[i].ID, errUpdate)
			continue
		}
	}

	updatedVideos, err := usecase.Repository.GetAllVideo(limit)
	if err != nil {
		return nil, err
	}

	return updatedVideos, nil
}

func (usecase *UserVideoUsecaseImpl) SearchVideoByKeywordUsecase(keyword string) (*[]video.Video, error) {
	videos, err := usecase.Repository.SearchVideoByKeyword(keyword)
	if err != nil {
		return nil, err
	}
	if videos == nil {
		return nil, pkg.ErrVideoNotFound
	}
	return videos, nil
}

func (usecase *UserVideoUsecaseImpl) SearchVideoByCategoryUsecase(categoryType string, name string) (*[]video.Video, error) {
	videos, err := usecase.Repository.SearchVideoByCategory(categoryType, name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, pkg.ErrVideoNotFound
		}
		return nil, err
	}
	if len(*videos) == 0 {
		return nil, pkg.ErrVideoNotFound
	}
	return videos, nil
}

func (usecase *UserVideoUsecaseImpl) GetVideoDetailUsecase(id int) (*video.Video, *[]video.Comment, int, error) {
	video, comments, totalComment, err := usecase.Repository.GetVideoDetail(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, 0, pkg.ErrVideoNotFound
		}
		return nil, nil, 0, err
	}
	return video, comments, totalComment, nil
}

func (usecase *UserVideoUsecaseImpl) AddCommentUsecase(request *dto.AddCommentRequest, userId string) error {
	if _, _, _, err := usecase.Repository.GetVideoDetail(request.VideoID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return pkg.ErrVideoNotFound
		}
		return err
	}

	comment := video.Comment{
		Comment: request.Comment,
		UserID:  userId,
		VideoID: request.VideoID,
	}
	if err := usecase.Repository.AddComment(&comment); err != nil {
		return err
	}
	return nil
}
