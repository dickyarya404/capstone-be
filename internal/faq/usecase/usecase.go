package faq

import (
	"github.com/sawalreverr/recything/internal/faq"
	"github.com/sawalreverr/recything/pkg"
)

type faqUsecase struct {
	faqRepo faq.FaqRepository
}

func NewFaqUsecase(faqRepo faq.FaqRepository) faq.FaqUsecase {
	return &faqUsecase{faqRepo: faqRepo}
}

func (u *faqUsecase) GetAllFaqs() (*[]faq.FaqResponse, error) {
	var faqs []faq.FaqResponse

	faqsFound, err := u.faqRepo.FindAll()
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	for _, fq := range *faqsFound {
		resp := faq.FaqResponse{
			ID:       fq.ID,
			Category: fq.Category,
			Question: fq.Question,
			Answer:   fq.Answer,
		}

		faqs = append(faqs, resp)
	}

	return &faqs, nil
}

func (u *faqUsecase) GetFaqsByCategory(category string) (*[]faq.FaqResponse, error) {
	var faqs []faq.FaqResponse

	faqsFound, err := u.faqRepo.FindByCategory(category)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	for _, fq := range *faqsFound {
		resp := faq.FaqResponse{
			ID:       fq.ID,
			Category: fq.Category,
			Question: fq.Question,
			Answer:   fq.Answer,
		}

		faqs = append(faqs, resp)
	}

	return &faqs, nil
}

func (u *faqUsecase) GetFaqsByKeyword(keyword string) (*[]faq.FaqResponse, error) {
	var faqs []faq.FaqResponse

	faqsFound, err := u.faqRepo.FindByKeyword(keyword)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	for _, fq := range *faqsFound {
		resp := faq.FaqResponse{
			ID:       fq.ID,
			Category: fq.Category,
			Question: fq.Question,
			Answer:   fq.Answer,
		}

		faqs = append(faqs, resp)
	}

	return &faqs, nil
}
