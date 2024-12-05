package user

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// struct
type User struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// PhoneNumber string    `json:"phone_number"`
	Password   string    `json:"-"`
	Point      uint      `json:"point" gorm:"default:0"`
	Gender     string    `json:"gender" gorm:"type:enum('laki-laki', 'perempuan', '-');default:-"`
	BirthDate  time.Time `json:"birth_date"`
	Address    string    `json:"address"`
	PictureURL string    `json:"picture_url"`
	OTP        uint      `json:"otp"`
	IsVerified bool      `json:"is_verified" gorm:"default:false"`
	Badge      string    `json:"badge" gorm:"default:'https://res.cloudinary.com/dymhvau8n/image/upload/v1718189121/user_badge/htaemsjtlhfof7ww01ss.png'"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// interface
type UserRepository interface {
	Create(user User) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByPhoneNumber(phoneNumber string) (*User, error)
	FindByID(userID string) (*User, error)
	FindAll(page int, limit int, sortBy string, sortType string) (*[]User, error)
	FindLastID() (string, error)
	Update(user User) error
	Delete(userID string) error
	CountAllUser() (int, error)
}

type UserUsecase interface {
	UpdateUserDetail(userID string, user UserDetail) error
	UpdateUserPicture(userID string, picture_url string) error
	UpdatePointAndBadge(userID string, point uint) error

	FindUserByID(userID string) (*UserResponse, error)
	FindAllUser(page int, limit int, sortBy string, sortType string) (*UserPaginationResponse, error)
	DeleteUser(userID string) error
}

type UserHandler interface {
	Profile(c echo.Context) error
	UpdateDetail(c echo.Context) error
	UploadAvatar(c echo.Context) error

	FindAllUser(c echo.Context) error
	DeleteUser(c echo.Context) error

	FindUser(c echo.Context) error
}
