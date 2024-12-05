package usecase

import (
	"mime/multipart"

	"github.com/sawalreverr/recything/internal/video/manage_video/dto"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type ManageVideoUsecase interface {
	CreateDataVideoUseCase(request *dto.CreateDataVideoRequest, thumbnail []*multipart.FileHeader) error
	GetAllCategoryVideoUseCase() ([]string, []string, error)
	GetAllDataVideoPaginationUseCase(limit int, page int) ([]video.Video, int, error)
	GetDetailsDataVideoByIdUseCase(id int) (*video.Video, error)
	UpdateDataVideoUseCase(request *dto.UpdateDataVideoRequest, thumbnail []*multipart.FileHeader, id int) error
	DeleteDataVideoUseCase(id int) error
}
