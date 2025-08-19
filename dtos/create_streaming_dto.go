package dtos

import (
	"mime/multipart"
	"time"
)

type CreateStreamingRequest struct {
	HostPrincipalID string                `json:"hostPrincipalId" example:"123"`
	Thumbnail       *multipart.FileHeader `json:"thumbnail" swaggerignore:"false"`
}

type CreateStreamingResponse struct {
	StreamID           string    `json:"streamId"`
	HostPrincipalID    string    `json:"hostPrincipalId"`
	Title              string    `json:"title"`
	Thumbnail          string    `form:"thumbnail"`
	StreamCategoryName string    `json:"streamCategoryName"`
	CreatedAt          time.Time `json:"createAt"`
	IsActive           bool      `json:"isActive"`
}
