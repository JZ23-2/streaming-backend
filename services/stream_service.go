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

	thumbnailURL, err := UploadThumbnail(req.Thumbnail, streamID)

	stream := models.Stream{
		StreamID:        streamID,
		HostPrincipalID: req.HostPrincipalID,
		CreatedAt:       time.Now(),
		ThumbnailURL:    thumbnailURL,
		IsActive:        true,
	}

	if err == nil {
		stream.ThumbnailURL = thumbnailURL
	}

	streamInfo, err := GetStreamInfoByUserID(req.HostPrincipalID)
	if err == nil {
		stream.StreamInfoID = &streamInfo.HostPrincipalID
	}
	if err := database.DB.Create(&stream).Error; err != nil {
		return nil, err
	}

	var createdStream models.Stream

	if err := database.DB.Preload("StreamInfo").Preload("StreamInfo.Category").Where("stream_id = ?", stream.StreamID).First(&createdStream).Error; err != nil {
		return nil, err
	}

	resp := &dtos.CreateStreamingResponse{
		StreamID:        createdStream.StreamID,
		HostPrincipalID: createdStream.HostPrincipalID,
		CreatedAt:       createdStream.CreatedAt,
		IsActive:        createdStream.IsActive,
	}

	return resp, nil
}

func GetAllActiveStream() ([]dtos.StreamResponse, error) {
	var stream []models.Stream

	if err := database.DB.
		Preload("StreamInfo").
		Preload("StreamInfo.Category").
		Preload("Messages").
		Where("is_active = ?", true).
		Find(&stream).Error; err != nil {
		return nil, err
	}

	var responses []dtos.StreamResponse

	for _, s := range stream {
		var messages []dtos.MessageResponse
		for _, m := range s.Messages {
			messages = append(messages, dtos.MessageResponse{
				MessageID: m.MessageID,
				SenderID:  m.MessagePrincipalID,
				Content:   m.Content,
				CreatedAt: m.CreatedAt,
			})
		}

		resp := dtos.StreamResponse{
			StreamID:        s.StreamID,
			HostPrincipalID: s.HostPrincipalID,
			ThumbnailURL:    s.ThumbnailURL,
			IsActive:        s.IsActive,
			CreatedAt:       s.CreatedAt,
			Messages:        messages,
		}

		streamInfo, err := GetStreamInfoByUserID(s.HostPrincipalID)
		if err == nil {
			resp.Title = streamInfo.Title
			resp.CategoryName = streamInfo.Category.CategoryName
		}
		responses = append(responses, resp)
	}
	return responses, nil
}

func GetActiveStreamByStreamID(streamID string) (*dtos.StreamResponse, error) {
	var stream models.Stream

	if err := database.DB.
		Preload("Messages").
		Preload("StreamInfo.Category").
		Where("stream_id = ? and is_active = ? ", streamID, true).
		First(&stream).Error; err != nil {
		return nil, err
	}

	var messages []dtos.MessageResponse
	for _, m := range stream.Messages {
		messages = append(messages, dtos.MessageResponse{
			MessageID: m.MessageID,
			SenderID:  m.MessagePrincipalID,
			Content:   m.Content,
			CreatedAt: m.CreatedAt,
		})
	}

	resp := dtos.StreamResponse{
		StreamID:        stream.StreamID,
		HostPrincipalID: stream.HostPrincipalID,
		ThumbnailURL:    stream.ThumbnailURL,
		IsActive:        stream.IsActive,
		CreatedAt:       stream.CreatedAt,
		Messages:        messages,
	}

	if stream.StreamInfo != nil {
		resp.Title = stream.StreamInfo.Title
		if stream.StreamInfo.Category != nil {
			resp.CategoryName = stream.StreamInfo.Category.CategoryName
		}
	}

	return &resp, nil
}

func StopStream(streamerID string) (*dtos.StreamResponse, error) {
	var stream models.Stream

	if err := database.DB.Where("host_principal_id = ? AND is_active = ?", streamerID, true).First(&stream).Error; err != nil {
		return nil, err
	}

	stream.IsActive = false

	if err := database.DB.Save(&stream).Error; err != nil {
		return nil, err
	}

	if err := database.DB.
		Preload("StreamInfo").
		Preload("StreamInfo.Category").
		Where("stream_id = ?", stream.StreamID).
		First(&stream).Error; err != nil {
		return nil, err
	}

	resp := &dtos.StreamResponse{
		StreamID:        stream.StreamID,
		HostPrincipalID: stream.HostPrincipalID,
		ThumbnailURL:    stream.ThumbnailURL,
		IsActive:        stream.IsActive,
		CreatedAt:       stream.CreatedAt,
	}

	if stream.StreamInfoID != nil {
		resp.Title = stream.StreamInfo.Title
		if stream.StreamInfo.StreamCategoryID != nil {
			resp.CategoryName = stream.StreamInfo.Category.CategoryName

		}
	}

	return resp, nil
}

func GetActiveStreamByStreamerID(streamerID string) (*dtos.StreamResponse, error) {
	var stream models.Stream

	if err := database.DB.
		Preload("Messages").
		Preload("StreamInfo").
		Preload("StreamInfo.Category").
		Where("is_active = ? AND host_principal_id = ?", true, streamerID).
		First(&stream).Error; err != nil {
		return nil, err
	}

	var messages []dtos.MessageResponse
	for _, m := range stream.Messages {
		messages = append(messages, dtos.MessageResponse{
			MessageID: m.MessageID,
			SenderID:  m.MessagePrincipalID,
			Content:   m.Content,
			CreatedAt: m.CreatedAt,
		})
	}

	resp := dtos.StreamResponse{
		StreamID:        stream.StreamID,
		HostPrincipalID: stream.HostPrincipalID,
		ThumbnailURL:    stream.ThumbnailURL,
		IsActive:        stream.IsActive,
		CreatedAt:       stream.CreatedAt,
		Messages:        messages,
	}

	if stream.StreamInfoID != nil {
		resp.Title = stream.StreamInfo.Title
		if stream.StreamInfo.StreamCategoryID != nil {
			resp.CategoryName = stream.StreamInfo.Category.CategoryName
		}
	}

	return &resp, nil
}
