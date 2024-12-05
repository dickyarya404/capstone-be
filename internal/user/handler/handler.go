package user

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	u "github.com/sawalreverr/recything/internal/user"
	"github.com/sawalreverr/recything/pkg"
)

type userHandler struct {
	userUsecase u.UserUsecase
}

func NewUserHandler(uc u.UserUsecase) u.UserHandler {
	return &userHandler{userUsecase: uc}
}

func (h *userHandler) Profile(c echo.Context) error {
	claims := c.Get("user").(*helper.JwtCustomClaims)

	user, err := h.userUsecase.FindUserByID(claims.UserID)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrStatusInternalError.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", user)
}

func (h *userHandler) UpdateDetail(c echo.Context) error {
	var user u.UserDetail
	claims := c.Get("user").(*helper.JwtCustomClaims)

	if err := c.Bind(&user); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if user.BirthDate != "" {
		parsedDate, err := time.Parse("2006-01-02", user.BirthDate)
		if err != nil {
			return helper.ErrorHandler(c, http.StatusBadRequest, fmt.Sprintf("invalid birth_date format: %v", err))
		}
		user.ParsedBirthDate = parsedDate
	}

	if err := c.Validate(&user); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := h.userUsecase.UpdateUserDetail(claims.UserID, user); err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrStatusInternalError.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "user detail updated!", nil)
}

func (h *userHandler) UploadAvatar(c echo.Context) error {
	claims := c.Get("user").(*helper.JwtCustomClaims)

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

	resp, err := helper.UploadToCloudinary(src, "recything/avatar/")
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "upload failed, cloudinary server error!")
	}

	if err := h.userUsecase.UpdateUserPicture(claims.UserID, resp); err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "update database error!")
	}

	return helper.ResponseHandler(c, http.StatusOK, "upload successfully!", echo.Map{
		"avatar_url": resp,
	})
}

func (h *userHandler) FindAllUser(c echo.Context) error {
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
		sortType = "desc"
	}

	if sortType == "" {
		sortType = "asc"
	}

	users, err := h.userUsecase.FindAllUser(page, limit, sortBy, sortType)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrStatusInternalError.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", users)
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	userID := c.Param("userId")

	if err := h.userUsecase.DeleteUser(userID); err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrStatusInternalError.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "user successfully deleted!", nil)
}

func (h *userHandler) FindUser(c echo.Context) error {
	userID := c.Param("userId")

	resp, err := h.userUsecase.FindUserByID(userID)
	if err != nil {
		if errors.Is(err, pkg.ErrUserNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrUserNotFound.Error())
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrStatusInternalError.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "ok", resp)
}
