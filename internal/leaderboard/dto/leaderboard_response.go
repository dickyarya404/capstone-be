package dto

type DataLeaderboard struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	PictureURL string `json:"picture_url"`
	Point      int    `json:"point"`
	Badge      string `json:"badge"`
	Address    string `json:"address"`
}

type LeaderboardResponse struct {
	DataLeaderboard []*DataLeaderboard `json:"data_leaderboard"`
}
