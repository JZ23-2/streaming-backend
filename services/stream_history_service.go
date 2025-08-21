package services

import (
	"main/database"
	"main/dtos"
	"main/helper"
	"main/models"
)

func CreateStreamHistory(req dtos.CreateStreamHistoryRequest) (*dtos.CreateStreamHistoryResponse, error) {
	streamHistoryID := helper.GenerateID()

	stream, err := GetStreamByID(req.StreamID)

	if err != nil {
		return nil, err
	}

	duration, err := helper.GetVideoDurationFromURL(req.VideoUrl)

	if err != nil {
		return nil, err
	}

	streamHistory := models.StreamHistory{
		StreamHistoryID:       streamHistoryID,
		StreamHistoryStreamID: stream.StreamID,
		HostPrincipalID:       req.HostPrincipalID,
		VideoUrl:              req.VideoUrl,
		Duration:              int(duration),
	}

	if stream.StreamInfoID != nil {
		streamHistory.Title = stream.StreamInfo.Title
		if stream.StreamInfo.StreamCategoryID != nil {
			streamHistory.StreamCategoryID = stream.StreamInfo.StreamCategoryID
		}
	}

	if err := database.DB.Create(&streamHistory).Error; err != nil {
		return nil, err
	}

	resp := dtos.CreateStreamHistoryResponse{
		StreamHistoryID:       streamHistory.StreamHistoryID,
		StreamHistoryStreamID: streamHistory.StreamHistoryStreamID,
		HostPrincipalID:       streamHistory.HostPrincipalID,
		VideoUrl:              streamHistory.VideoUrl,
		Duration:              streamHistory.Duration,
	}

	return &resp, nil
}

func GetAllStreamHistoryByStreamerID(hostPrincipalID string) ([]dtos.StreamHistoryResponse, error) {
	var streamHistories []models.StreamHistory

	if err := database.DB.
		Preload("Category").
		Preload("Stream.Messages").
		Preload("Stream.StreamInfo").
		Where("host_principal_id = ?", hostPrincipalID).
		Find(&streamHistories).Error; err != nil {
		return nil, err
	}

	var responses []dtos.StreamHistoryResponse
	for _, s := range streamHistories {

		var messages []dtos.MessageResponse
		for _, m := range s.Stream.Messages {
			messages = append(messages, dtos.MessageResponse{
				MessageID: m.MessageID,
				SenderID:  m.MessagePrincipalID,
				Content:   m.Content,
				CreatedAt: m.CreatedAt,
			})
		}

		var totalViews int64
		database.DB.Model(&models.ViewerHistory{}).
			Where("viewer_history_stream_history_id = ?", s.StreamHistoryID).
			Count(&totalViews)

		resp := dtos.StreamHistoryResponse{
			StreamHistoryID:       s.StreamHistoryID,
			StreamHistoryStreamID: s.StreamHistoryStreamID,
			HostPrincipalID:       s.HostPrincipalID,
			VideoUrl:              s.VideoUrl,
			Duration:              s.Duration,
			Thumbnail:             s.Stream.ThumbnailURL,
			MessageResponse:       messages,
			TotalView:             int(totalViews),
			Title:                 s.Title,
			CreatedAt:             s.CreatedAt,
		}

		if s.StreamCategoryID != nil {
			resp.CategoryName = &s.Category.CategoryName
		}

		responses = append(responses, resp)
	}

	return responses, nil
}

func GetStreamHistoryByID(streamHistoryID string) (*dtos.StreamHistoryResponse, error) {
	var streamHistory models.StreamHistory

	if err := database.DB.
		Preload("Category").
		Preload("Stream.Messages").
		Where("stream_history_id = ?", streamHistoryID).
		First(&streamHistory).Error; err != nil {
		return nil, err
	}

	var messages []dtos.MessageResponse
	for _, m := range streamHistory.Stream.Messages {
		messages = append(messages, dtos.MessageResponse{
			MessageID: m.MessageID,
			SenderID:  m.MessagePrincipalID,
			Content:   m.Content,
			CreatedAt: m.CreatedAt,
		})
	}

	var totalViews int64
	database.DB.Model(&models.ViewerHistory{}).
		Where("viewer_history_stream_history_id = ?", streamHistory.StreamHistoryID).
		Count(&totalViews)

	resp := &dtos.StreamHistoryResponse{
		StreamHistoryID:       streamHistory.StreamHistoryID,
		StreamHistoryStreamID: streamHistory.StreamHistoryStreamID,
		HostPrincipalID:       streamHistory.HostPrincipalID,
		VideoUrl:              streamHistory.VideoUrl,
		Duration:              streamHistory.Duration,
		Thumbnail:             streamHistory.Stream.ThumbnailURL,
		MessageResponse:       messages,
		TotalView:             int(totalViews),
		Title:                 streamHistory.Title,
		CategoryName:          &streamHistory.Category.CategoryName,
	}

	return resp, nil
}
