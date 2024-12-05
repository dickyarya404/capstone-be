package repository

import (
	"github.com/sawalreverr/recything/internal/database"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
	"github.com/sawalreverr/recything/pkg"
)

type UserVideoRepositoryImpl struct {
	DB database.Database
}

func NewUserVideoRepository(db database.Database) *UserVideoRepositoryImpl {
	return &UserVideoRepositoryImpl{DB: db}
}

func (repository *UserVideoRepositoryImpl) GetAllVideo(limit int) (*[]video.Video, error) {
	var videos []video.Video
	if err := repository.DB.GetDB().
		Limit(limit).
		Order("created_at desc").
		Preload("Categories.ContentCategory").
		Preload("Categories.WasteCategory").
		Find(&videos).
		Error; err != nil {
		return nil, err
	}

	return &videos, nil
}

func (repository *UserVideoRepositoryImpl) SearchVideoByKeyword(keyword string) (*[]video.Video, error) {
	var videos []video.Video
	if err := repository.DB.GetDB().
		Order("created_at desc").
		Preload("Categories.ContentCategory").
		Preload("Categories.WasteCategory").
		Joins("JOIN video_categories ON video_categories.video_id = videos.id").
		Joins("JOIN content_categories ON content_categories.id = video_categories.content_category_id").
		Joins("JOIN waste_categories ON waste_categories.id = video_categories.waste_category_id").
		Where("videos.title LIKE ? OR videos.description LIKE ? OR content_categories.name LIKE ? OR waste_categories.name LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").
		Find(&videos).Error; err != nil {
		return nil, err
	}

	videoMap := make(map[int]video.Video)
	for _, v := range videos {
		videoMap[v.ID] = v
	}

	uniqueVideos := make([]video.Video, 0, len(videoMap))
	for _, v := range videoMap {
		uniqueVideos = append(uniqueVideos, v)
	}

	return &uniqueVideos, nil
}

func (repository *UserVideoRepositoryImpl) SearchVideoByCategory(categoryType string, name string) (*[]video.Video, error) {
	var videos []video.Video
	if categoryType == "content" {
		if err := repository.DB.GetDB().
			Order("created_at desc").
			Joins("JOIN video_categories ON video_categories.video_id = videos.id").
			Joins("JOIN content_categories ON content_categories.id = video_categories.content_category_id").
			Where("content_categories.name LIKE ?", "%"+name+"%").
			Preload("Categories.ContentCategory").
			Preload("Categories.WasteCategory").
			Find(&videos).Error; err != nil {
			return nil, err
		}

	} else if categoryType == "waste" {
		if err := repository.DB.GetDB().
			Order("created_at desc").
			Joins("JOIN video_categories ON video_categories.video_id = videos.id").
			Joins("JOIN waste_categories ON waste_categories.id = video_categories.waste_category_id").
			Where("waste_categories.name LIKE ?", "%"+name+"%").
			Preload("Categories.ContentCategory").
			Preload("Categories.WasteCategory").
			Find(&videos).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, pkg.ErrVideoNotFound
	}

	videoMap := make(map[int]video.Video)
	for _, v := range videos {
		videoMap[v.ID] = v
	}

	uniqueVideos := make([]video.Video, 0, len(videoMap))
	for _, v := range videoMap {
		uniqueVideos = append(uniqueVideos, v)
	}

	return &uniqueVideos, nil
}

func (repository *UserVideoRepositoryImpl) GetVideoDetail(id int) (*video.Video, *[]video.Comment, int, error) {
	var videos video.Video
	var comments []video.Comment
	var totalComment int64

	db := repository.DB.GetDB()

	if err := db.Model(&video.Video{}).Where("id = ?", id).First(&videos).Error; err != nil {
		return nil, nil, 0, err
	}

	if err := db.Model(&video.Comment{}).Preload("User").Where("video_id = ?", id).Find(&comments).Error; err != nil {
		return nil, nil, 0, err
	}

	if err := db.Model(&video.Comment{}).Where("video_id = ?", id).Count(&totalComment).Error; err != nil {
		return nil, nil, 0, err
	}

	return &videos, &comments, int(totalComment), nil
}

func (repository *UserVideoRepositoryImpl) AddComment(comment *video.Comment) error {
	if err := repository.DB.GetDB().Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (repository *UserVideoRepositoryImpl) UpdateViewer(view int, id int) error {
	if err := repository.DB.GetDB().Model(&video.Video{}).Where("id = ?", id).Update("viewer", view).Error; err != nil {
		return err
	}
	return nil
}
