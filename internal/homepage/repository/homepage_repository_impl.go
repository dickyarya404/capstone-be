package repository

import (
	admin "github.com/sawalreverr/recything/internal/admin/entity"
	"github.com/sawalreverr/recything/internal/article"
	"github.com/sawalreverr/recything/internal/database"
	"github.com/sawalreverr/recything/internal/user"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type HomepageRepositoryImpl struct {
	DB database.Database
}

func NewHomepageRepository(db database.Database) HomepageRepository {
	return &HomepageRepositoryImpl{DB: db}
}

func (repository *HomepageRepositoryImpl) GetArcticle() (*[]article.Article, error) {
	var articles []article.Article
	if err := repository.DB.GetDB().
		Order("created_at desc").
		Limit(2).
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return &articles, nil
}

func (repository *HomepageRepositoryImpl) GetVideo() (*[]video.Video, error) {
	var videos []video.Video
	if err := repository.DB.GetDB().
		Order("created_at desc").
		Limit(3).
		Find(&videos).Error; err != nil {
		return nil, err
	}
	return &videos, nil
}

func (repository *HomepageRepositoryImpl) GetLeaderboard() (*[]user.User, error) {
	var users []user.User
	if err := repository.DB.GetDB().
		Order("point desc").
		Limit(3).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (repository *HomepageRepositoryImpl) GetUserData(userId string) (*user.User, error) {
	var user user.User
	if err := repository.DB.GetDB().First(&user, "id = ?", userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *HomepageRepositoryImpl) FindAdminByID(adminId string) (*admin.Admin, error) {
	var admin *admin.Admin
	if err := repository.DB.GetDB().First(&admin, "id = ?", adminId).Error; err != nil {
		return nil, err
	}
	return admin, nil
}
