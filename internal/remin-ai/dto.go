package reminai

type RequestInput struct {
	Question string `json:"question" validate:"required"`
}

type RequestOutput struct {
	Question string `json:"question"`
	AnswerAI string `json:"answer_ai"`
}
