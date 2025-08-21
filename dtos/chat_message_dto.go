package dtos

type ChatMessage struct {
	StreamID string `json:"streamId"`
	UserID   string `json:"userId"`
	Content  string `json:"content"`
	Username string `json:"username"`
}
