package dtos

import "time"

type GetActiveAllStreamResponse struct {
	StreamID                 string                     `json:"streamId"`
	HostPrincipalID          string                     `json:"hostPrincipalID"`
	Title                    string                     `json:"title"`
	Thumbnail                string                     `json:"thumbnail"`
	CategoryName             string                     `json:"categoryName"`
	IsActive                 bool                       `json:"isActive"`
	CreatedAt                time.Time                  `json:"createdAt"`
	MessageAllStreamResponse []MessageAllStreamResponse `json:"messages"`
}
