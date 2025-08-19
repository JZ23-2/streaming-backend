package dtos

import "time"

type MessageResponse struct {
	MessageID string    `json:"messageID"`
	SenderID  string    `json:"senderID"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
