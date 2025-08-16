package services

import (
	"fmt"
	"main/database"
	"main/dtos"
	"main/helper"
	"main/models"
	"main/storage"
	"mime/multipart"
	"time"
)

func UploadThumbnail(fileHeader *multipart.FileHeader, streamID string) (string, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	destPath := fmt.Sprintf("thumbnails/%s_%s", streamID, fileHeader.Filename)

	publicURL, err := storage.UploadFileFromReader(src, destPath, fileHeader.Header.Get("Content-Type"))
	if err != nil {
		return "", err
	}

	return publicURL, nil
}

func CreateStream(req dtos.CreateStreamingRequest) (*dtos.CreateStreamingResponse, error) {
	streamID := helper.GenerateID()

	publicUrl, err := UploadThumbnail(req.Thumbnail, streamID)
	if err != nil {
		return nil, err
	}

	stream := models.Stream{
		StreamID:         streamID,
		HostPrincipalID:  req.HostPrincipalID,
		Title:            req.Title,
		Thumbnail:        publicUrl,
		StreamCategoryID: req.StreamCategoryID,
		CreatedAt:        time.Now(),
	}

	if err := database.DB.Create(&stream).Error; err != nil {
		return nil, err
	}

	resp := &dtos.CreateStreamingResponse{
		StreamID:         stream.StreamID,
		HostPrincipalID:  stream.HostPrincipalID,
		Title:            stream.Title,
		Thumbnail:        stream.Thumbnail,
		StreamCategoryID: stream.StreamCategoryID,
		CreatedAt:        stream.CreatedAt,
	}

	return resp, nil
}
