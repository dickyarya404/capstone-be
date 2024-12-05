package article

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	art "github.com/sawalreverr/recything/internal/article"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type articleHandler struct {
	usecase art.ArticleUsecase
}

func NewArticleHandler(uc art.ArticleUsecase) art.ArticleHandler {
	return &articleHandler{usecase: uc}
}

func (h *articleHandler) NewArticle(c echo.Context) error {
	var request art.ArticleInput

	authorID := c.Get("user").(*helper.JwtCustomClaims).UserID

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	response, err := h.usecase.NewArticle(request, authorID)
	if err != nil {
		if errors.Is(pkg.ErrCategoryArticleNotFound, err) {
			return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusCreated, "article created!", response)
}

func (h *articleHandler) UpdateArticle(c echo.Context) error {
	var request art.ArticleInput
	articleID := c.Param("articleId")

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.Update(articleID, request); err != nil {
		if errors.Is(pkg.ErrArticleNotFound, err) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		} else if errors.Is(pkg.ErrCategoryArticleNotFound, err) {
			return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "article updated!", nil)
}

func (h *articleHandler) DeleteArticle(c echo.Context) error {
	articleID := c.Param("articleId")

	if err := h.usecase.Delete(articleID); err != nil {
		if errors.Is(pkg.ErrArticleNotFound, err) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "article deleted!", nil)
}

func (h *articleHandler) GetAllArticle(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}

	sortBy := c.QueryParam("sort_by")
	sortType := c.QueryParam("sort_type")

	if sortBy == "" {
		sortBy = "created_at"
		sortType = "asc"
	}

	if sortType == "" {
		sortType = "asc"
	}

	response, err := h.usecase.GetAllArticle(page, limit, sortBy, sortType)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", response)
}

func (h *articleHandler) GetArticleByKeyword(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	response, err := h.usecase.GetArticleByKeyword(keyword)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", response)
}

func (h *articleHandler) GetArticleByCategory(c echo.Context) error {
	categoryType := c.QueryParam("type")
	categoryName := c.QueryParam("name")

	response, err := h.usecase.GetArticleByCategory(categoryName, categoryType)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", response)
}

func (h *articleHandler) GetArticleByID(c echo.Context) error {
	articleId := c.Param("articleId")

	articleFound, err := h.usecase.GetArticleByID(articleId)
	if err != nil {
		if errors.Is(pkg.ErrArticleNotFound, err) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", articleFound)
}

func (h *articleHandler) NewArticleComment(c echo.Context) error {
	var request art.CommentInput
	articleID := c.Param("articleId")
	userID := c.Get("user").(*helper.JwtCustomClaims).UserID

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	request.ArticleID = articleID
	request.UserID = userID

	if err := h.usecase.NewArticleComment(request); err != nil {
		if errors.Is(pkg.ErrArticleNotFound, err) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}

		if errors.Is(pkg.ErrUserNotFound, err) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusCreated, "comment added!", nil)
}

func (h *articleHandler) ArticleUploadImage(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "please upload your image!")
	}

	if file.Size > 2*1024*1024 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "upload image size must less than 2MB!")
	}

	fileType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(fileType, "image/") {
		return helper.ErrorHandler(c, http.StatusBadRequest, "only image allowed!")
	}

	src, _ := file.Open()
	defer src.Close()

	resp, err := helper.UploadToCloudinary(src, "recything/article/")
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "upload failed, cloudinary server error!")
	}

	return helper.ResponseHandler(c, http.StatusOK, "upload successfully!", echo.Map{
		"image_url": resp,
	})
}

func (h *articleHandler) GetAllCategories(c echo.Context) error {
	response, err := h.usecase.GetAllCategories()
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", response)
}
