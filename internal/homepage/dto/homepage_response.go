package dto

type HomepageResponse struct {
	User        *DataUser          `json:"user"`
	Articles    []*DataArtcicle    `json:"articles"`
	Videos      []*DataVideo       `json:"videos"`
	Leaderboard []*DataLeaderboard `json:"leaderboard"`
}

type DataUser struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	PictureURL string `json:"picture_url"`
	Point      int    `json:"point"`
	Badge      string `json:"badge"`
}

type DataArtcicle struct {
	Id             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Thumbnail      string `json:"thumbnail"`
	AuthorName     string `json:"author_name"`
	Author_Profile string `json:"author_profile"`
	CreatedAt      string `json:"created_at"`
}

type DataVideo struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	UrlThumbnail string `json:"url_thumbnail"`
	LinkVideo    string `json:"link_video"`
	Viewer       int    `json:"viewer"`
}

type DataLeaderboard struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	PictureURL string `json:"picture_url"`
	Point      int    `json:"point"`
	Badge      string `json:"badge"`
}
