package repository

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type UserVideoRepository interface {
	GetAllVideo(limit int) (*[]video.Video, error)
	SearchVideoByKeyword(keyword string) (*[]video.Video, error)
	SearchVideoByCategory(categoryType string, name string) (*[]video.Video, error)
	GetVideoDetail(id int) (*video.Video, *[]video.Comment, int, error)
	AddComment(comment *video.Comment) error
	UpdateViewer(view int, id int) error
}
