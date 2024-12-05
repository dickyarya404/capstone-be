package dto

type UpdateAchievementRequest struct {
	Level       string `json:"level" form:"level"`
	TargetPoint int    `json:"target_point" form:"target_point"`
}
