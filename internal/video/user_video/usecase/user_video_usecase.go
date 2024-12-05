package usecase

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
	"github.com/sawalreverr/recything/internal/video/user_video/dto"
)

type UserVideoUsecase interface {
	GetAllVideoUsecase(limit int) (*[]video.Video, error)
	SearchVideoByKeywordUsecase(keyword string) (*[]video.Video, error)
	SearchVideoByCategoryUsecase(categoryType string, name string) (*[]video.Video, error)
	GetVideoDetailUsecase(id int) (*video.Video, *[]video.Comment, int, error)
	AddCommentUsecase(request *dto.AddCommentRequest, userId string) error
}
