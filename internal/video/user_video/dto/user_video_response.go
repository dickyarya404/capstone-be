package dto

import "time"

type DataVideo struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	UrlThumbnail string `json:"url_thumbnail"`
	LinkVideo    string `json:"link_video"`
	Viewer       int    `json:"viewer"`
}

type SearchVideoByKeywordResponse struct {
	DataVideo *[]DataVideoSearchByCategory `json:"data_video"`
}

type GetAllVideoResponse struct {
	DataVideo []DataVideo `json:"data_video"`
}

type DataComment struct {
	Id          int       `json:"id"`
	Comment     string    `json:"comment"`
	UserID      string    `json:"user_id"`
	UserName    string    `json:"user_name"`
	UserProfile string    `json:"user_profile"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetDetailsDataVideoByIdResponse struct {
	DataVideo    *DataVideo     `json:"data_video"`
	TotalComment int            `json:"total_comment"`
	Comments     *[]DataComment `json:"comments"`
}

type DataCategoryVideo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DataTrashCategoryVideo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type SearchVideoByCategoryVideoResponse struct {
	DataVideo []*DataVideoSearchByCategory `json:"data_video"`
}

type DataVideoSearchByCategory struct {
	Id            int                       `json:"id"`
	Title         string                    `json:"title"`
	Description   string                    `json:"description"`
	UrlThumbnail  string                    `json:"url_thumbnail"`
	LinkVideo     string                    `json:"link_video"`
	Viewer        int                       `json:"viewer"`
	VideoCategory []*DataCategoryVideo      `json:"content_categories"`
	TrashCategory []*DataTrashCategoryVideo `json:"waste_categories"`
}
