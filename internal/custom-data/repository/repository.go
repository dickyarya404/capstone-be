package customdata

import (
	"fmt"

	cdt "github.com/sawalreverr/recything/internal/custom-data"
	"github.com/sawalreverr/recything/internal/database"
)

type customDataRepository struct {
	DB database.Database
}

func NewCustomDataRepository(db database.Database) cdt.CustomDataRepository {
	return &customDataRepository{DB: db}
}

func (r *customDataRepository) Create(data cdt.CustomData) (*cdt.CustomData, error) {
	if err := r.DB.GetDB().Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *customDataRepository) FindByID(dataID string) (*cdt.CustomData, error) {
	var customData cdt.CustomData
	if err := r.DB.GetDB().Where("id = ?", dataID).First(&customData).Error; err != nil {
		return nil, err
	}

	return &customData, nil
}

func (r *customDataRepository) FindAll(page int, limit int, sortBy string, sortType string) (*[]cdt.CustomData, int64, error) {
	var customDatas []cdt.CustomData
	var total int64

	db := r.DB.GetDB().Model(&cdt.CustomData{})

	if sortBy != "" {
		sort := fmt.Sprintf("%s %s", sortBy, sortType)
		db = db.Order(sort)
	}

	db.Count(&total)

	offset := (page - 1) * limit
	if err := db.Offset(offset).Limit(limit).Find(&customDatas).Error; err != nil {
		return nil, 0, err
	}

	return &customDatas, total, nil
}

func (r *customDataRepository) FindLastID() (string, error) {
	var customData cdt.CustomData
	if err := r.DB.GetDB().Unscoped().Order("id DESC").First(&customData).Error; err != nil {
		return "CDT0000", err
	}

	return customData.ID, nil
}

func (r *customDataRepository) Update(data cdt.CustomData) error {
	if err := r.DB.GetDB().Save(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r *customDataRepository) Delete(dataID string) error {
	var customData cdt.CustomData
	if err := r.DB.GetDB().Where("id = ?", dataID).Delete(&customData).Error; err != nil {
		return err
	}

	return nil
}
