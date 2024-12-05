package customdata

import (
	"time"

	cdt "github.com/sawalreverr/recything/internal/custom-data"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type customDataUsecase struct {
	customDataRepository cdt.CustomDataRepository
}

func NewCustomDataUsecase(repo cdt.CustomDataRepository) cdt.CustomDataUsecase {
	return &customDataUsecase{customDataRepository: repo}
}

func (uc *customDataUsecase) NewCustomData(data cdt.CustomDataInput) (*cdt.CustomDataResponse, error) {
	lastID, _ := uc.customDataRepository.FindLastID()
	newID := helper.GenerateCustomID(lastID, "CDT")

	newCustomData := cdt.CustomData{
		ID:          newID,
		Topic:       data.Topic,
		Description: data.Description,
	}

	dataCreated, err := uc.customDataRepository.Create(newCustomData)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	response := cdt.CustomDataResponse{
		ID:          dataCreated.ID,
		Topic:       dataCreated.Topic,
		Description: dataCreated.Description,
		CreatedAt:   dataCreated.CreatedAt.String(),
		UpdatedAt:   dataCreated.UpdatedAt.String(),
	}

	return &response, nil
}

func (uc *customDataUsecase) FindByID(dataID string) (*cdt.CustomDataResponse, error) {
	dataFound, err := uc.customDataRepository.FindByID(dataID)
	if err != nil {
		return nil, pkg.ErrCustomDataNotFound
	}

	response := cdt.CustomDataResponse{
		ID:          dataFound.ID,
		Topic:       dataFound.Topic,
		Description: dataFound.Description,
		CreatedAt:   dataFound.CreatedAt.String(),
		UpdatedAt:   dataFound.UpdatedAt.String(),
	}

	return &response, nil
}

func (uc *customDataUsecase) FindAll(page int, limit int, sortBy string, sortType string) (*cdt.CustomDataPaginationResponse, error) {
	var customDataResponses []cdt.CustomDataResponse
	customDatas, total, err := uc.customDataRepository.FindAll(page, limit, sortBy, sortType)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	for _, data := range *customDatas {
		response := cdt.CustomDataResponse{
			ID:          data.ID,
			Topic:       data.Topic,
			Description: data.Description,
			CreatedAt:   data.CreatedAt.String(),
			UpdatedAt:   data.UpdatedAt.String(),
		}
		customDataResponses = append(customDataResponses, response)
	}

	response := cdt.CustomDataPaginationResponse{
		Total:      total,
		Page:       page,
		Limit:      limit,
		CustomData: customDataResponses,
	}

	return &response, nil
}

func (uc *customDataUsecase) UpdateData(dataID string, data cdt.CustomDataInput) error {
	dataFound, err := uc.customDataRepository.FindByID(dataID)
	if err != nil {
		return pkg.ErrCustomDataNotFound
	}

	dataFound.Topic = data.Topic
	dataFound.Description = data.Description
	dataFound.UpdatedAt = time.Now()

	if err := uc.customDataRepository.Update(*dataFound); err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}

func (uc *customDataUsecase) DeleteData(dataID string) error {
	dataFound, err := uc.customDataRepository.FindByID(dataID)
	if err != nil {
		return pkg.ErrCustomDataNotFound
	}

	if err := uc.customDataRepository.Delete(dataFound.ID); err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}
