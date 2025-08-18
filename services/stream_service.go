package services

import (
	"fmt"
	"main/controllers"
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

	var createdStream models.Stream

	if err := database.DB.Preload("Category").Where("stream_id = ?", stream.StreamID).First(&createdStream).Error; err != nil {
		return nil, err
	}

	resp := &dtos.CreateStreamingResponse{
		StreamID:           createdStream.StreamID,
		HostPrincipalID:    createdStream.HostPrincipalID,
		Title:              createdStream.Title,
		Thumbnail:          createdStream.Thumbnail,
		StreamCategoryName: createdStream.Category.CategoryName,
		CreatedAt:          createdStream.CreatedAt,
	}

	return resp, nil
}

func GetAllActiveStream() ([]dtos.GetActiveAllStreamResponse, error) {
	var stream []models.Stream

	if err := database.DB.
		Preload("Category").
		Preload("Messages").
		Where("is_active = ?", true).
		Find(&stream).Error; err != nil {
		return nil, err
	}

	var responses []dtos.GetActiveAllStreamResponse

	for _, s := range stream {

		var messages []dtos.MessageAllStreamResponse
		for _, m := range s.Messages {
			messages = append(messages, dtos.MessageAllStreamResponse{
				MessageID: m.MessageID,
				SenderID:  m.MessagePrincipalID,
				Content:   m.Content,
				CreatedAt: m.CreatedAt,
			})
		}

		resp := dtos.GetActiveAllStreamResponse{
			StreamID:                 s.StreamID,
			HostPrincipalID:          s.HostPrincipalID,
			Title:                    s.Title,
			Thumbnail:                s.Thumbnail,
			CategoryName:             s.Category.CategoryName,
			IsActive:                 s.IsActive,
			CreatedAt:                s.CreatedAt,
			MessageAllStreamResponse: messages,
			ViewerCount:              controllers.GetViewerCount(s.HostPrincipalID),
		}
		responses = append(responses, resp)
	}

	return responses, nil
}

func GetActiveStreamByStreamID(streamID string) (*dtos.GetActiveAllStreamResponse, error) {
	var stream models.Stream

	if err := database.DB.
		Preload("Category").
		Preload("Messages").
		Where("is_active = ?", true).
		First(&stream).Error; err != nil {
		return nil, err
	}

	var messages []dtos.MessageAllStreamResponse
	for _, m := range stream.Messages {
		messages = append(messages, dtos.MessageAllStreamResponse{
			MessageID: m.MessageID,
			SenderID:  m.MessagePrincipalID,
			Content:   m.Content,
			CreatedAt: m.CreatedAt,
		})
	}

	resp := dtos.GetActiveAllStreamResponse{
		StreamID:                 stream.StreamID,
		HostPrincipalID:          stream.HostPrincipalID,
		Title:                    stream.Title,
		Thumbnail:                stream.Thumbnail,
		CategoryName:             stream.Category.CategoryName,
		IsActive:                 stream.IsActive,
		CreatedAt:                stream.CreatedAt,
		MessageAllStreamResponse: messages,
	}

	return &resp, nil
}

func UpdateStreamActiveStatus(req dtos.UpdateStreamActiveStatusRequest) (*dtos.GetActiveAllStreamResponse, error) {

	var stream models.Stream

	if err := database.DB.
		Where("stream_id = ?", req.StreamID).
		First(&stream).Error; err != nil {
		return nil, err
	}

	if err := database.DB.
		Model(&stream).
		Update("is_active", req.IsActive).Error; err != nil {
		return nil, err
	}

	if err := database.DB.
		Preload("Category").
		Preload("Messages").
		Where("stream_id = ?", req.StreamID).
		First(&stream).Error; err != nil {
		return nil, err
	}

	var messages []dtos.MessageAllStreamResponse
	for _, m := range stream.Messages {
		messages = append(messages, dtos.MessageAllStreamResponse{
			MessageID: m.MessageID,
			SenderID:  m.MessagePrincipalID,
			Content:   m.Content,
			CreatedAt: m.CreatedAt,
		})
	}

	resp := dtos.GetActiveAllStreamResponse{
		StreamID:                 stream.StreamID,
		HostPrincipalID:          stream.HostPrincipalID,
		Title:                    stream.Title,
		Thumbnail:                stream.Thumbnail,
		CategoryName:             stream.Category.CategoryName,
		IsActive:                 stream.IsActive,
		CreatedAt:                stream.CreatedAt,
		MessageAllStreamResponse: messages,
	}

	return &resp, nil
}

func UpdateStream(req dtos.UpdateStreamingRequest) (*dtos.UpdateStreamingResponse, error) {
	var stream models.Stream

	if err := database.DB.
		Where("stream_id = ?", req.StreamID).
		Preload("Category").
		First(&stream).Error; err != nil {
		return nil, err
	}

	publicUrl := stream.Thumbnail
	if req.Thumbnail != nil {
		uploadedUrl, err := UploadThumbnail(req.Thumbnail, stream.StreamID)
		if err != nil {
			return nil, err
		}
		publicUrl = uploadedUrl
	}

	updateData := map[string]interface{}{
		"title":              req.Title,
		"thumbnail":          publicUrl,
		"stream_category_id": req.StreamCategoryID,
	}

	if err := database.DB.Model(&stream).Updates(updateData).Error; err != nil {
		return nil, err
	}

	if err := database.DB.Preload("Category").
		First(&stream, "stream_id = ?", req.StreamID).Error; err != nil {
		return nil, err
	}

	resp := &dtos.UpdateStreamingResponse{
		StreamID:           stream.StreamID,
		HostPrincipalID:    stream.HostPrincipalID,
		Title:              stream.Title,
		Thumbnail:          stream.Thumbnail,
		StreamCategoryName: stream.Category.CategoryName,
		CreatedAt:          stream.CreatedAt,
	}

	return resp, nil
}
