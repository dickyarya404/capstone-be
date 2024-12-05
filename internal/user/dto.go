package user

import "time"

type UserDetail struct {
	Name  string `json:"name"`
	Email string `json:"email" validate:"email"`
	// PhoneNumber     string    `json:"phone_number" validate:"min=10"`
	Gender          string    `json:"gender"`
	BirthDate       string    `json:"birth_date"`
	ParsedBirthDate time.Time `json:"-"`
	Address         string    `json:"address"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// PhoneNumber string    `json:"phone_number"`
	Point      uint      `json:"point"`
	Badge      string    `json:"badge"`
	Gender     string    `json:"gender"`
	BirthDate  time.Time `json:"birth_date"`
	Address    string    `json:"address"`
	PictureURL string    `json:"picture_url"`
	CreatedAt  time.Time `json:"created_at"`
}

type UserPaginationResponse struct {
	TotalUser int            `json:"total_user"`
	Page      int            `json:"page"`
	Limit     int            `json:"limit"`
	Users     []UserResponse `json:"users"`
}
