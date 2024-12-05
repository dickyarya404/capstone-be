package faq

type FaqResponse struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
