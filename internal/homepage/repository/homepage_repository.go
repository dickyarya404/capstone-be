package repository

import (
	admin "github.com/sawalreverr/recything/internal/admin/entity"
	"github.com/sawalreverr/recything/internal/article"
	user "github.com/sawalreverr/recything/internal/user"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type HomepageRepository interface {
	GetArcticle() (*[]article.Article, error)
	GetVideo() (*[]video.Video, error)
	GetLeaderboard() (*[]user.User, error)
	GetUserData(userId string) (*user.User, error)
	FindAdminByID(adminId string) (*admin.Admin, error)
}
