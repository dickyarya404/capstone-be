package faq

import (
	"github.com/sawalreverr/recything/internal/database"
	faq "github.com/sawalreverr/recything/internal/faq"
)

type faqRepository struct {
	DB database.Database
}

func NewFaqRepository(db database.Database) faq.FaqRepository {
	return &faqRepository{DB: db}
}

func (r *faqRepository) FindAll() (*[]faq.FAQ, error) {
	var faqs []faq.FAQ
	if err := r.DB.GetDB().Find(&faqs).Error; err != nil {
		return nil, err
	}

	return &faqs, nil
}

func (r *faqRepository) FindByCategory(category string) (*[]faq.FAQ, error) {
	var faqs []faq.FAQ
	if err := r.DB.GetDB().Where("category = ?", category).Find(&faqs).Error; err != nil {
		return nil, err
	}

	return &faqs, nil
}

func (r *faqRepository) FindByKeyword(keyword string) (*[]faq.FAQ, error) {
	var faqs []faq.FAQ
	query := "%" + keyword + "%"
	if err := r.DB.GetDB().Where("question LIKE ?", query).Find(&faqs).Error; err != nil {
		return nil, err
	}

	return &faqs, nil
}
