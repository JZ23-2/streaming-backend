package dtos

type ChatMessage struct {
	StreamID string `json:"stream_id"`
	UserID   string `json:"user_id"`
	Content  string `json:"content"`
}
