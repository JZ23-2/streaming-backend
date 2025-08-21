package services

import (
	"main/database"
	"main/dtos"
	"main/helper"
	"main/models"
)

func CreateHighlight(req dtos.CreateHighlightRequest) (*dtos.CreateHighlightResponse, error) {
	var highlights []dtos.ClipResponse

	for _, clip := range req.Clips {
		highlightID := helper.GenerateID()

		highlight := models.Highlight{
			HighlightID:              highlightID,
			HighlightStreamHistoryID: clip.HighlightStreamHistoryID,
			HighlightUrl:             clip.HighlightUrl,
			StartHighlight:           clip.StartHighlight,
			EndHighlight:             clip.EndHighlight,
			HighlightDescription:     clip.HighlightDescription,
		}

		if err := database.DB.Create(&highlight).Error; err != nil {
			return nil, err
		}

		highlights = append(highlights, dtos.ClipResponse{
			HighlightID:              highlight.HighlightID,
			HighlightStreamHistoryID: highlight.HighlightStreamHistoryID,
			HighlightUrl:             highlight.HighlightUrl,
			StartHighlight:           highlight.StartHighlight,
			EndHighlight:             highlight.EndHighlight,
			HighlightDescription:     highlight.HighlightDescription,
		})
	}

	resp := &dtos.CreateHighlightResponse{
		Highlights: highlights,
	}

	return resp, nil
}

func GetAllHighlightByStreamerID(streamerID string) (*dtos.CreateHighlightResponse, error) {
	var highlights []models.Highlight

	err := database.DB.
		Joins("JOIN stream_histories ON stream_histories.stream_history_id = highlights.highlight_stream_history_id").
		Where("stream_histories.host_principal_id = ?", streamerID).
		Preload("StreamHistory").
		Find(&highlights).Error

	if err != nil {
		return nil, err
	}

	var responses []dtos.ClipResponse
	for _, h := range highlights {
		responses = append(responses, dtos.ClipResponse{
			HighlightID:              h.HighlightID,
			HighlightStreamHistoryID: h.HighlightStreamHistoryID,
			HighlightUrl:             h.HighlightUrl,
			StartHighlight:           h.StartHighlight,
			EndHighlight:             h.EndHighlight,
			HighlightDescription:     h.HighlightDescription,
		})
	}

	resp := &dtos.CreateHighlightResponse{
		Highlights: responses,
	}

	return resp, nil
}
