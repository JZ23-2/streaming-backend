package dtos

import (
	"mime/multipart"
	"time"
)

type CreateStreamingRequest struct {
	HostPrincipalID  string                `json:"hostPrincipalId" example:"nigger 123"`
	Title            string                `json:"title" example:"nigger show"`
	Thumbnail        *multipart.FileHeader `form:"thumbnail" swaggerignore:"false"`
	StreamCategoryID string                `json:"streamCategoryId" example:"nigger category"`
}

type CreateStreamingResponse struct {
	StreamID         string    `json:"streamId"`
	HostPrincipalID  string    `json:"hostPrincipalId"`
	Title            string    `json:"title"`
	Thumbnail        string    `form:"thumbnail"`
	StreamCategoryID string    `json:"streamCategoryId"`
	CreatedAt        time.Time `json:"createAt"`
}
