package dto

import (
	"time"
)

type UserTaskGetAllResponse struct {
	Data []DataUserTask `json:"data"`
}

type DataUserTask struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Thumbnail   string      `json:"thumbnail"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     time.Time   `json:"end_date"`
	Point       int         `json:"point"`
	Status      bool        `json:"status"`
	TaskSteps   []TaskSteps `json:"task_steps"`
}

type TaskSteps struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserTaskResponseCreate struct {
	Id             string                      `json:"id"`
	StatusProgress string                      `json:"status_progress"`
	TaskChalenge   TaskChallengeResponseCreate `json:"task_challenge"`
}

type UserTaskRejectResponse struct {
	Id             string             `json:"id"`
	StatusProgress string             `json:"status_progress"`
	StatusAccept   string             `json:"status_accepted"`
	Reason         string             `json:"reason"`
	Point          int                `json:"point"`
	TaskChalenge   DataTaskChallenges `json:"task_challenge"`
	UserSteps      []DataUserSteps    `json:"user_steps"`
}

type TaskChallengeResponseCreate struct {
	Id          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Thumbnail   string          `json:"thumbnail"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
	Point       int             `json:"point"`
	StatusTask  bool            `json:"status_task"`
	TaskSteps   []TaskSteps     `json:"task_steps"`
	UserSteps   []DataUserSteps `json:"user_steps"`
}

type DataUserSteps struct {
	Id                  int    `json:"id"`
	UserTaskChallengeID string `json:"user_task_challenge_id"`
	TaskStepID          int    `json:"task_step_id"`
	Completed           bool   `json:"completed"`
}

type UserTaskUploadImageResponse struct {
	Id             string             `json:"id"`
	StatusProgress string             `json:"status_progress"`
	StatusAccept   string             `json:"status_accepted"`
	Point          int                `json:"point"`
	TaskChallenge  DataTaskChallenges `json:"task_challenge"`
	Images         []Images           `json:"images"`
	UserSteps      []DataUserSteps    `json:"user_steps"`
}

type Images struct {
	Images string `json:"images"`
}

type UserTaskGetByIdUserResponse struct {
	Id             string                      `json:"id"`
	StatusProgress string                      `json:"status_progress"`
	TaskChallenge  TaskChallengeResponseCreate `json:"task_challenge"`
}

type GetUserTaskDoneByIdUserResponse struct {
	Id             string                 `json:"id"`
	StatusProgress string                 `json:"status_progress"`
	StatusAccept   string                 `json:"status_accepted"`
	Point          int                    `json:"point"`
	ReasonReject   string                 `json:"reason_reject"`
	TaskChallenge  DataTaskChanlengesDone `json:"task_challenge"`
}

type DataTaskChanlengesDone struct {
	Id          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Thumbnail   string          `json:"thumbnail"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
	StatusTask  bool            `json:"status_task"`
	TaskSteps   []TaskSteps     `json:"task_steps"`
	UserSteps   []DataUserSteps `json:"user_steps"`
}

type DataTaskChallenges struct {
	Id          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Thumbnail   string          `json:"thumbnail"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
	StatusTask  bool            `json:"status_task"`
	TaskSteps   []TaskSteps     `json:"task_steps"`
	UserSteps   []DataUserSteps `json:"user_steps"`
}

type GetUserTaskDetailsResponse struct {
	Id          string        `json:"id"`
	TitleTask   string        `json:"title_task"`
	UserName    string        `json:"user_name"`
	Images      []*DataImages `json:"images"`
	Description string        `json:"description"`
}

type DataImages struct {
	Id       int    `json:"id"`
	ImageUrl string `json:"image_url"`
}

type DataHistoryPoint struct {
	Id         string    `json:"id"`
	TitleTask  string    `json:"title_task"`
	Point      int       `json:"point"`
	AcceptedAt time.Time `json:"accepted_at"`
}

type HistoryPointResponse struct {
	TotalPoint int                 `json:"total_point"`
	Data       []*DataHistoryPoint `json:"data_history_point"`
}

type UpdateTaskStep struct {
	Id             string                      `json:"id"`
	StatusProgress string                      `json:"status_progress"`
	TaskChalenge   TaskChallengeResponseCreate `json:"task_challenge"`
}

type DataGetUserTaskByUserTaskId struct {
	Id          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Thumbnail   string          `json:"thumbnail"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
	Point       int             `json:"point"`
	StatusTask  bool            `json:"status_task"`
	TaskSteps   []TaskSteps     `json:"task_steps"`
	UserSteps   []DataUserSteps `json:"user_steps"`
}

type GetUserTaskByUserTaskIdResponse struct {
	Id             string                      `json:"id"`
	StatusProgress string                      `json:"status_progress"`
	StatusAccept   string                      `json:"status_accepted"`
	Reason         string                      `json:"reason"`
	TaskChalenge   DataGetUserTaskByUserTaskId `json:"task_challenge"`
}
