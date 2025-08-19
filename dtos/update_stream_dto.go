package dtos

import (
	"mime/multipart"
	"time"
)

type UpdateStreamActiveStatusRequest struct {
	HostPrincipalID string `json:"hostPrincipalId" example:"stream123"`
}

type UpdateStreamingRequest struct {
	StreamID         string                `json:"streamId" example:"stream 123"`
	Title            string                `json:"title" example:"nigger show"`
	Thumbnail        *multipart.FileHeader `form:"thumbnail" swaggerignore:"false"`
	StreamCategoryID string                `json:"streamCategoryId" example:"nigger category"`
}

type UpdateStreamingResponse struct {
	StreamID           string    `json:"streamId"`
	HostPrincipalID    string    `json:"hostPrincipalId"`
	Title              string    `json:"title"`
	Thumbnail          string    `form:"thumbnail"`
	StreamCategoryName string    `json:"streamCategoryName"`
	CreatedAt          time.Time `json:"createAt"`
}
