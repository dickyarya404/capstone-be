package dto

type RejectUserTaskRequest struct {
	Reason string `json:"reason" validate:"required"`
}
