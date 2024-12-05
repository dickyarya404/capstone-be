package repository

import (
	art "github.com/sawalreverr/recything/internal/article"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type ManageVideoRepository interface {
	CreateVideoAndCategories(video *video.Video) (*video.Video, error)
	CreateVideoCategories(videoCategories []video.VideoCategory) error
	FindTitleVideo(title string) error
	FindNameCategoryVideo(name string) (*art.ContentCategory, error)
	FindNamaTrashCategory(name string) (*art.WasteCategory, error)
	GetAllCategoryVideo() ([]string, error)
	GetAllTrashCategoryVideo() ([]string, error)
	GetCategoryVideoById(id int) (*video.VideoCategory, error)
	GetAllDataVideoPagination(limit int, page int) ([]video.Video, int, error)
	GetDetailsDataVideoById(id int) (*video.Video, error)
	UpdateDataVideo(video *video.Video, id int) error
	DeleteDataVideo(id int) error
}
