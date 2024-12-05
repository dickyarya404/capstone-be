package dto

type DashboardResponse struct {
	User                           TotalUser               `json:"user"`
	Report                         TotalReport             `json:"report"`
	Challenge                      TotalChallenge          `json:"challenge"`
	Content                        TotalContent            `json:"content"`
	UserAchievement                UserAchievement         `json:"user_achievement"`
	DataReportByCity               []DataReportByCity      `json:"data_report_by_city"`
	DataUserByGender               []DataUserByGender      `json:"data_user_by_gender"`
	TotalLittering                 int                     `json:"total_report_littering"`
	TotalRubbish                   int                     `json:"total_report_rubbish"`
	DataReportByWasteTypeRubbish   []DataReportByWasteType `json:"data_report_by_waste_rubbish"`
	DataReportByWasteTypeLittering []DataReportByWasteType `json:"data_report_by_waste_littering"`
	DataReportStatistic            DataReportStatistic     `json:"data_report_statistic"`
}

type TotalUser struct {
	TotalUser                  int `json:"total_user"`
	AdditionUserSinceYesterday int `json:"addition_user_since_yesterday"`
}

type UserAchievement struct {
	TotalUser int `json:"total_user"`
	Classic   int `json:"classic"`
	Silver    int `json:"silver"`
	Gold      int `json:"gold"`
	Platinum  int `json:"platinum"`
}

type TotalReport struct {
	TotalReport                  int `json:"total_report"`
	AdditionReportSinceYesterday int `json:"addition_report_since_yesterday"`
}

type TotalChallenge struct {
	TotalChallenge                 int `json:"total_challenge"`
	AdditionChallengeSinceLastWeek int `json:"addition_challenge_since_last_week"`
}

type TotalContent struct {
	TotalContent         int `json:"total_content"`
	AdditionContentToday int `json:"addition_content_today"`
}

type DailyReportStats struct {
	Day          int64 `json:"day"`
	TotalReports int64 `json:"total_reports"`
}

type MonthlyReportStats struct {
	Month        string             `json:"month"`
	DailyStats   []DailyReportStats `json:"daily_statistic"`
	TotalReports int64              `json:"total_reports"`
}

type DataReportStatistic struct {
	ReportLittering []MonthlyReportStats `json:"report_littering"`
	ReportRubbish   []MonthlyReportStats `json:"report_rubbish"`
}

type DataReportByCity struct {
	City        string `json:"city"`
	TotalReport int    `json:"total_report"`
}

type DataUserByGender struct {
	Gender     string  `json:"gender"`
	TotalUser  int64   `json:"total_user"`
	Percentage float64 `json:"percentage"`
}

type DataReportByWasteType struct {
	ReportType   string  `json:"report_type"`
	WasteType    string  `json:"waste_type"`
	TotalReports int64   `json:"total_reports"`
	Percentage   float64 `json:"percentage"`
}
