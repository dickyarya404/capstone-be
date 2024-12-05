package usecase

import (
	"github.com/sawalreverr/recything/internal/homepage/dto"
	"github.com/sawalreverr/recything/internal/homepage/repository"
	"github.com/sawalreverr/recything/pkg"
)

type HomepageUsecaseImpl struct {
	HomepageRepository repository.HomepageRepository
}

func NewHomepageUsecase(homepageRepository repository.HomepageRepository) HomepageUsecase {
	return &HomepageUsecaseImpl{HomepageRepository: homepageRepository}
}

func (usecase *HomepageUsecaseImpl) GetHomepageUsecase(userId string) (*dto.HomepageResponse, error) {
	videos, err := usecase.HomepageRepository.GetVideo()
	if err != nil {
		return nil, err
	}
	var dataVideo []*dto.DataVideo
	for _, video := range *videos {
		dataVideo = append(dataVideo, &dto.DataVideo{
			Id:           video.ID,
			Title:        video.Title,
			Description:  video.Description,
			UrlThumbnail: video.Thumbnail,
			LinkVideo:    video.Link,
			Viewer:       video.Viewer,
		})
	}

	leaderboard, err := usecase.HomepageRepository.GetLeaderboard()
	if err != nil {
		return nil, err
	}
	var dataLeaderboard []*dto.DataLeaderboard
	for _, user := range *leaderboard {
		dataLeaderboard = append(dataLeaderboard, &dto.DataLeaderboard{
			Id:         user.ID,
			Name:       user.Name,
			PictureURL: user.PictureURL,
			Point:      int(user.Point),
			Badge:      user.Badge,
		})
	}
	user, err := usecase.HomepageRepository.GetUserData(userId)
	if err != nil {
		return nil, pkg.ErrUserNotFound
	}
	articles, errArticle := usecase.HomepageRepository.GetArcticle()
	if errArticle != nil {
		return nil, err
	}
	var dataArticle []*dto.DataArtcicle
	for _, article := range *articles {
		admin, erradmin := usecase.HomepageRepository.FindAdminByID(article.AuthorID)
		if erradmin != nil {
			return nil, erradmin
		}
		dataArticle = append(dataArticle, &dto.DataArtcicle{
			Id:             article.ID,
			Title:          article.Title,
			Description:    article.Description,
			Thumbnail:      article.ThumbnailURL,
			CreatedAt:      article.CreatedAt.String(),
			AuthorName:     admin.Name,
			Author_Profile: admin.ImageUrl,
		})
	}
	return &dto.HomepageResponse{
		User:        &dto.DataUser{Id: user.ID, Name: user.Name, Point: int(user.Point), Badge: user.Badge, PictureURL: user.PictureURL},
		Articles:    dataArticle,
		Videos:      dataVideo,
		Leaderboard: dataLeaderboard,
	}, nil
}
