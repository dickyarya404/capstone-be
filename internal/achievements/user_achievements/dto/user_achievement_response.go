package dto

import "time"

type DataAchievement struct {
	Id          int    `json:"id"`
	Level       string `json:"level"`
	TargetPoint int    `json:"target_point"`
	BadgeUrl    string `json:"badge_url"`
}

type DataUser struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Point int    `json:"point"`
	Badge string `json:"badge"`
}

type HistoryUserPoint struct {
	Point      int       `json:"point"`
	AcceptedAt time.Time `json:"accepted_at"`
}

type GetAchievementByUserResponse struct {
	DataAchievement  []*DataAchievement  `json:"data_achievement"`
	DataUser         *DataUser           `json:"data_user"`
	HistoryUserPoint []*HistoryUserPoint `json:"history_user_point"`
}
