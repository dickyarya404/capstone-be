package usecase

import (
	au "github.com/sawalreverr/recything/internal/about-us"
	"github.com/sawalreverr/recything/pkg"
)

type aboutUsUsecase struct {
	AboutUsRepository au.AboutUsRepository
}

func NewAboutUsUsecase(repo au.AboutUsRepository) au.AboutUsUsecase {
	return &aboutUsUsecase{AboutUsRepository: repo}
}

func (uc *aboutUsUsecase) GetAboutUsByCategory(categoryName string) (*[]au.AboutUsResponse, error) {
	var response []au.AboutUsResponse
	categoryAbouts, err := uc.AboutUsRepository.FindByCategory(categoryName)
	if err != nil {
		return nil, pkg.ErrAboutUsCategoryNotFound
	}

	for _, category := range *categoryAbouts {
		var respImages []au.AboutUsImageResponse
		images, _ := uc.AboutUsRepository.FindAllImageByID(category.ID)

		if len(*images) != 0 {
			for _, image := range *images {
				respImage := au.AboutUsImageResponse{
					AboutUsID: image.AboutUsID,
					Name:      image.Name,
					ImageURL:  image.ImageURL,
				}

				respImages = append(respImages, respImage)
			}
		}

		resp := au.AboutUsResponse{
			ID:          category.ID,
			Category:    category.Category,
			Title:       category.Title,
			Description: category.Description,
			Images:      respImages,
		}

		response = append(response, resp)
	}

	return &response, nil
}
