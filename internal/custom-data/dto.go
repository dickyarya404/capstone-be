package customdata

type CustomDataInput struct {
	Topic       string `json:"topic" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CustomDataResponse struct {
	ID          string `json:"id"`
	Topic       string `json:"topic"`
	Description string `json:"description"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CustomDataPaginationResponse struct {
	Total      int64                `json:"total"`
	Page       int                  `json:"page"`
	Limit      int                  `json:"limit"`
	CustomData []CustomDataResponse `json:"custom_datas"`
}
