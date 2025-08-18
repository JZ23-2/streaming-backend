package services

import (
	"main/database"
	"main/dtos"
	"main/helper"
	"main/models"
)

func CreateStreamHistory(req dtos.CreateStreamHistoryRequest) (*dtos.CreateStreamHistoryResponse, error) {
	streamHistoryID := helper.GenerateID()

	duration, err := helper.GetVideoDurationFromURL(req.VideoUrl)

	if err != nil {
		return nil, err
	}

	streamHistory := models.StreamHistory{
		StreamHistoryID:       streamHistoryID,
		StreamHistoryStreamID: req.StreamHistoryStreamID,
		HostPrincipalID:       req.HostPrincipalID,
		VideoUrl:              req.VideoUrl,
		Duration:              int(duration),
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

func GetAllStreamHistoryByStreamerID(hostPrincipalID string) ([]dtos.GetAllStreamHistoryResponse, error) {
	var streamHistories []models.StreamHistory

	if err := database.DB.
		Preload("Stream.Category").
		Preload("Stream.Messages").
		Where("host_principal_id = ?", hostPrincipalID).
		Find(&streamHistories).Error; err != nil {
		return nil, err
	}

	var responses []dtos.GetAllStreamHistoryResponse
	for _, s := range streamHistories {

		var messages []dtos.MessageAllStreamResponse
		for _, m := range s.Stream.Messages {
			messages = append(messages, dtos.MessageAllStreamResponse{
				MessageID: m.MessageID,
				SenderID:  m.MessagePrincipalID,
				Content:   m.Content,
				CreatedAt: m.CreatedAt,
			})
		}

		var totalViews int64
		database.DB.Model(&models.ViewerHistory{}).
			Where("viewer_history_stream_id = ?", s.StreamHistoryStreamID).
			Count(&totalViews)

		resp := dtos.GetAllStreamHistoryResponse{
			StreamHistoryID:          s.StreamHistoryID,
			StreamHistoryStreamID:    s.StreamHistoryStreamID,
			HostPrincipalID:          s.HostPrincipalID,
			VideoUrl:                 s.VideoUrl,
			Duration:                 s.Duration,
			Title:                    s.Stream.Title,
			Thumbnail:                s.Stream.Thumbnail,
			CategoryName:             s.Stream.Category.CategoryName,
			MessageAllStreamResponse: messages,
			TotalView:                int(totalViews),
		}
		responses = append(responses, resp)
	}

	return responses, nil
}

func GetAllStreamHistoryByID(streamHistoryID string) (*dtos.GetAllStreamHistoryResponse, error) {
	var streamHistory models.StreamHistory

	if err := database.DB.
		Preload("Stream.Category").
		Preload("Stream.Messages").
		Where("stream_history_id = ?", streamHistoryID).
		First(&streamHistory).Error; err != nil {
		return nil, err
	}

	var messages []dtos.MessageAllStreamResponse
	for _, m := range streamHistory.Stream.Messages {
		messages = append(messages, dtos.MessageAllStreamResponse{
			MessageID: m.MessageID,
			SenderID:  m.MessagePrincipalID,
			Content:   m.Content,
			CreatedAt: m.CreatedAt,
		})
	}

	var totalViews int64
	database.DB.Model(&models.ViewerHistory{}).
		Where("viewer_history_stream_id = ?", streamHistory.StreamHistoryStreamID).
		Count(&totalViews)

	resp := &dtos.GetAllStreamHistoryResponse{
		StreamHistoryID:          streamHistory.StreamHistoryID,
		StreamHistoryStreamID:    streamHistory.StreamHistoryStreamID,
		HostPrincipalID:          streamHistory.HostPrincipalID,
		VideoUrl:                 streamHistory.VideoUrl,
		Duration:                 streamHistory.Duration,
		Title:                    streamHistory.Stream.Title,
		Thumbnail:                streamHistory.Stream.Thumbnail,
		CategoryName:             streamHistory.Stream.Category.CategoryName,
		MessageAllStreamResponse: messages,
		TotalView:                int(totalViews),
	}

	return resp, nil
}
