package helper

import (
	"context"
	"net/url"

	"github.com/sawalreverr/recything/config"
	"github.com/sawalreverr/recything/pkg"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func GetVideoViewCount(videoURL string) (uint64, error) {

	conf := config.GetConfig()
	apiKey := conf.YouTube.APIKey
	parsedURL, err := url.Parse(videoURL)
	if err != nil {
		return 0, pkg.ErrParsingUrl
	}

	queryParams := parsedURL.Query()

	videoID := queryParams.Get("v")
	if videoID == "" {
		return 0, pkg.ErrNoVideoIdFoundOnUrl
	}

	ctx := context.Background()

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return 0, pkg.ErrVideoService
	}

	call := service.Videos.List([]string{"statistics"}).Id(videoID)
	response, err := call.Do()
	if err != nil {
		return 0, pkg.ErrApiYouTube
	}

	if len(response.Items) == 0 {
		return 0, pkg.ErrVideoNotFound
	}
	video := response.Items[0]
	viewCount := video.Statistics.ViewCount

	return viewCount, nil
}
