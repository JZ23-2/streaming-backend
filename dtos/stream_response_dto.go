package dtos

import "time"

type StreamResponse struct {
	StreamID        string            `json:"streamId"`
	HostPrincipalID string            `json:"hostPrincipalID"`
	Title           string            `json:"title"`
	ThumbnailURL    string            `json:"thumbnailURL"`
	CategoryName    string            `json:"categoryName"`
	IsActive        bool              `json:"isActive"`
	CreatedAt       time.Time         `json:"createdAt"`
	Messages        []MessageResponse `json:"messages"`
	ViewerCount     int               `json:"viewerCount"`
}
