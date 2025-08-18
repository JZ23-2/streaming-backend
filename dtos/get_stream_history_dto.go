package dtos

import "time"

type GetAllStreamHistoryResponse struct {
	StreamHistoryID          string                     `json:"streamHistoryID"`
	StreamHistoryStreamID    string                     `json:"streamHistoryStreamID"`
	HostPrincipalID          string                     `json:"hostPrincipalID"`
	VideoUrl                 string                     `json:"videoUrl"`
	Duration                 int                        `json:"duration"`
	Title                    string                     `json:"title"`
	Thumbnail                string                     `json:"thumbnail"`
	CategoryName             string                     `json:"categoryName"`
	MessageAllStreamResponse []MessageAllStreamResponse `json:"messages"`
	TotalView                int                        `json:"totalView"`
}

type MessageAllStreamResponse struct {
	MessageID string    `json:"messageID"`
	SenderID  string    `json:"senderID"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
