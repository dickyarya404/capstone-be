package dto

type AddCommentRequest struct {
	VideoID int    `json:"video_id"`
	Comment string `json:"comment"`
}
