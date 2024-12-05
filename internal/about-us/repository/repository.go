package aboutus

import (
	au "github.com/sawalreverr/recything/internal/about-us"
	"github.com/sawalreverr/recything/internal/database"
)

type aboutUsRepository struct {
	DB database.Database
}

func NewAboutUsRepository(db database.Database) au.AboutUsRepository {
	return &aboutUsRepository{DB: db}
}

func (r *aboutUsRepository) FindByCategory(categoryName string) (*[]au.AboutUs, error) {
	var abouts []au.AboutUs
	if err := r.DB.GetDB().Where("category = ?", categoryName).Find(&abouts).Error; err != nil {
		return nil, err
	}

	return &abouts, nil
}

func (r *aboutUsRepository) FindAllImageByID(aboutUsID string) (*[]au.AboutUsImage, error) {
	var images []au.AboutUsImage
	if err := r.DB.GetDB().Where("about_us_id = ?", aboutUsID).Find(&images).Error; err != nil {
		return nil, err
	}

	return &images, nil
}
