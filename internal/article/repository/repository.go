package article

import (
	"errors"
	"fmt"

	art "github.com/sawalreverr/recything/internal/article"
	"github.com/sawalreverr/recything/internal/database"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type articleRepository struct {
	DB database.Database
}

func NewArticleRepository(db database.Database) art.ArticleRepository {
	return &articleRepository{DB: db}
}

func (r *articleRepository) Create(article art.Article) (*art.Article, error) {
	if err := r.DB.GetDB().Create(&article).Error; err != nil {
		return nil, err
	}

	return &article, nil
}

func (r *articleRepository) FindByID(articleID string) (*art.Article, error) {
	var article art.Article
	if err := r.DB.GetDB().Preload("Categories").Preload("Sections").Preload("Comments").First(&article, "id = ?", articleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrArticleNotFound
		}
		return nil, err
	}
	return &article, nil
}

func (r *articleRepository) FindAll(page, limit uint, sortBy string, sortType string) (*[]art.Article, int64, error) {
	var articles []art.Article
	var total int64

	db := r.DB.GetDB().Model(&art.Article{})

	offset := (page - 1) * limit

	if sortBy != "" {
		sort := fmt.Sprintf("%s %s", sortBy, sortType)
		db = db.Order(sort)
	}

	db.Count(&total)

	if err := db.Preload("Categories").Limit(int(limit)).Offset(int(offset)).Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return &articles, total, nil
}

func (r *articleRepository) FindLastID() (string, error) {
	var article art.Article
	if err := r.DB.GetDB().Unscoped().Order("id DESC").First(&article).Error; err != nil {
		return "ART0000", err
	}

	return article.ID, nil
}

func (r *articleRepository) FindByKeyword(keyword string) (*[]art.Article, error) {
	var articles []art.Article
	query := "%" + keyword + "%"

	if err := r.DB.GetDB().
		Preload("Categories").
		Joins("LEFT JOIN article_categories ON articles.id = article_categories.article_id").
		Joins("LEFT JOIN waste_categories ON article_categories.waste_category_id = waste_categories.id").
		Joins("LEFT JOIN content_categories ON article_categories.content_category_id = content_categories.id").
		Where("articles.title LIKE ? OR articles.description LIKE ? OR waste_categories.name LIKE ? OR content_categories.name LIKE ?", query, query, query, query).
		Find(&articles).Error; err != nil {
		return nil, err
	}

	return &articles, nil
}

func (r *articleRepository) FindByCategory(categoryName string, categoryType string) (*[]art.Article, error) {
	var articles []art.Article

	if categoryType == "waste" {
		if err := r.DB.GetDB().Preload("Categories").
			Joins("JOIN article_categories ON articles.id = article_categories.article_id").
			Joins("JOIN waste_categories ON article_categories.waste_category_id = waste_categories.id").
			Where("waste_categories.name = ?", categoryName).
			Find(&articles).Error; err != nil {
			return nil, err
		}
	} else if categoryType == "content" {
		if err := r.DB.GetDB().Preload("Categories").
			Joins("JOIN article_categories ON articles.id = article_categories.article_id").
			Joins("JOIN content_categories ON article_categories.content_category_id = content_categories.id").
			Where("content_categories.name = ?", categoryName).
			Find(&articles).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, pkg.ErrCategoryArticleNotFound
	}

	return &articles, nil
}

func (r *articleRepository) Update(article art.Article) error {
	if err := r.DB.GetDB().Save(&article).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) Delete(articleID string) error {
	var article art.Article
	if err := r.DB.GetDB().Delete(&article, "id = ?", articleID).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) FindCategories(articleID string) (*[]art.WasteCategory, *[]art.ContentCategory, error) {
	var articleCategories []art.ArticleCategories
	var wasteCategories []art.WasteCategory
	var contentCategories []art.ContentCategory

	if err := r.DB.GetDB().Where("article_id = ?", articleID).Find(&articleCategories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, pkg.ErrArticleNotFound
		}
		return nil, nil, err
	}

	var wasteCategoryIDs []uint
	var contentCategoryIDs []uint
	for _, ac := range articleCategories {
		if ac.WasteCategoryID != 0 {
			wasteCategoryIDs = append(wasteCategoryIDs, ac.WasteCategoryID)
		}
		if ac.ContentCategoryID != 0 {
			contentCategoryIDs = append(contentCategoryIDs, uint(ac.ContentCategoryID))
		}
	}

	if len(wasteCategoryIDs) > 0 {
		if err := r.DB.GetDB().Where("id IN (?)", wasteCategoryIDs).Find(&wasteCategories).Error; err != nil {
			return nil, nil, err
		}
	}

	if len(contentCategoryIDs) > 0 {
		if err := r.DB.GetDB().Where("id IN (?)", contentCategoryIDs).Find(&contentCategories).Error; err != nil {
			return nil, nil, err
		}
	}

	return &wasteCategories, &contentCategories, nil
}

func (r *articleRepository) CreateSection(section art.ArticleSection) error {
	if err := r.DB.GetDB().Create(&section).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) UpdateSection(section art.ArticleSection) error {
	if err := r.DB.GetDB().Save(&section).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) DeleteSection(sectionID uint) error {
	var section art.ArticleSection
	if err := r.DB.GetDB().Delete(&section, "id = ?", sectionID).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) DeleteAllSection(articleID string) error {
	if err := r.DB.GetDB().Where("article_id = ?", articleID).Delete(&art.ArticleSection{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) CreateArticleCategory(categories art.ArticleCategories) error {
	if err := r.DB.GetDB().Create(&categories).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) UpdateArticleCategory(categories art.ArticleCategories) error {
	if err := r.DB.GetDB().Save(&categories).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) DeleteAllArticleCategory(articleID string) error {
	if err := r.DB.GetDB().Where("article_id = ?", articleID).Delete(&art.ArticleCategories{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) FindCategoryByName(categoryName, categoryType string) (uint, error) {
	if categoryType == "waste" {
		var wasteCategory art.WasteCategory
		if err := r.DB.GetDB().Where("name = ?", categoryName).First(&wasteCategory).Error; err != nil {
			return 0, err
		}

		return wasteCategory.ID, nil
	}

	if categoryType == "content" {
		var contentCategory art.ContentCategory
		if err := r.DB.GetDB().Where("name = ?", categoryName).First(&contentCategory).Error; err != nil {
			return 0, err
		}

		return contentCategory.ID, nil
	}

	return 0, pkg.ErrCategoryArticleNotFound
}

func (r *articleRepository) CreateArticleComment(comment art.ArticleComment) error {
	if err := r.DB.GetDB().Create(&comment).Error; err != nil {
		return err
	}

	return nil
}

func (r *articleRepository) DeleteAllArticleComment(articleID string) error {
	if err := r.DB.GetDB().Where("article_id = ?", articleID).Delete(&art.ArticleComment{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *articleRepository) FindAllCategories() (*[]art.WasteCategory, *[]art.ContentCategory, error) {
	var wasteCategories []art.WasteCategory
	var contentCategories []art.ContentCategory

	if err := r.DB.GetDB().Model(&art.WasteCategory{}).Find(&wasteCategories).Error; err != nil {
		return nil, nil, err
	}

	if err := r.DB.GetDB().Model(&art.ContentCategory{}).Find(&contentCategories).Error; err != nil {
		return nil, nil, err
	}

	return &wasteCategories, &contentCategories, nil
}
