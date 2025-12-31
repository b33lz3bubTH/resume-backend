package dto

type ChatRequest struct {
	Message   string `json:"message" validate:"required,min=1"`
	SessionID string `json:"session_id" validate:"required,min=1"`
}

type ChatResponse struct {
	Answer      string  `json:"answer"`
	SessionID   string  `json:"session_id"`
	MessageID   string  `json:"message_id"`
}

