package repository

import (
	"log"

	"github.com/sawalreverr/recything/internal/database"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
	"gorm.io/gorm"

	art "github.com/sawalreverr/recything/internal/article"
)

type ManageVideoRepositoryImpl struct {
	DB database.Database
}

func NewManageVideoRepository(db database.Database) *ManageVideoRepositoryImpl {
	return &ManageVideoRepositoryImpl{DB: db}
}

func (repository *ManageVideoRepositoryImpl) CreateVideoAndCategories(video *video.Video) (*video.Video, error) {
	if err := repository.DB.GetDB().Create(&video).Error; err != nil {
		return nil, err
	}
	return video, nil
}

func (repository *ManageVideoRepositoryImpl) CreateVideoCategories(videoCategories []video.VideoCategory) error {
	if err := repository.DB.GetDB().Create(&videoCategories).Error; err != nil {
		return err
	}
	return nil
}

func (repository *ManageVideoRepositoryImpl) FindTitleVideo(title string) error {
	var video video.Video
	if err := repository.DB.GetDB().Where("title = ?", title).First(&video).Error; err != nil {
		return err
	}
	return nil
}

func (repository *ManageVideoRepositoryImpl) FindNameCategoryVideo(name string) (*art.ContentCategory, error) {
	var category art.ContentCategory
	if err := repository.DB.GetDB().Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (repository *ManageVideoRepositoryImpl) FindNamaTrashCategory(name string) (*art.WasteCategory, error) {
	var category art.WasteCategory
	if err := repository.DB.GetDB().Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (repository *ManageVideoRepositoryImpl) GetAllCategoryVideo() ([]string, error) {
	var categories []string
	if err := repository.DB.GetDB().Model(&video.VideoCategory{}).Distinct("name").Pluck("name", &categories).
		Error; err != nil {

	}
	return categories, nil
}

func (repository *ManageVideoRepositoryImpl) GetAllTrashCategoryVideo() ([]string, error) {
	var categories []string
	if err := repository.DB.GetDB().Model(&video.Video{}).Distinct("name").Pluck("name", &categories).
		Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (repository *ManageVideoRepositoryImpl) GetCategoryVideoById(id int) (*video.VideoCategory, error) {
	var category video.VideoCategory
	if err := repository.DB.GetDB().Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (repository *ManageVideoRepositoryImpl) GetAllDataVideoPagination(limit int, page int) ([]video.Video, int, error) {
	var videos []video.Video
	var total int64
	offset := (page - 1) * limit
	if err := repository.DB.GetDB().Model(&video.Video{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := repository.DB.GetDB().Limit(limit).Offset(offset).Order("id desc").Find(&videos).Error; err != nil {
		return nil, 0, err
	}
	return videos, int(total), nil

}

func (repository *ManageVideoRepositoryImpl) GetDetailsDataVideoById(id int) (*video.Video, error) {
	var video video.Video
	if err := repository.DB.GetDB().
		Preload("Categories.ContentCategory").
		Preload("Categories.WasteCategory").
		Where("id = ?", id).
		First(&video).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

func (repository *ManageVideoRepositoryImpl) UpdateDataVideo(videos *video.Video, id int) error {
	tx := repository.DB.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("Transaction rollback due to panic:", r)
		}
	}()

	if len(videos.Categories) > 0 {
		if err := tx.Where("video_id = ?", id).Delete(&video.VideoCategory{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Update video details along with associations
	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(&videos).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repository *ManageVideoRepositoryImpl) DeleteDataVideo(id int) error {
	if err := repository.DB.GetDB().Delete(&video.Video{}, id).Error; err != nil {
		return err
	}
	return nil
}
