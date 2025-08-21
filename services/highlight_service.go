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
