package services

import (
	"main/database"
	"main/dtos"
	"main/helper"
	"main/models"
)

func CreateViewerHistory(req dtos.CreateViewerHistoryRequest) (*dtos.CreateViewerHistoryResponse, error) {
	viewerHistoryID := helper.GenerateID()

	viewerHistory := models.ViewerHistory{
		ViewerHistoryID:          viewerHistoryID,
		ViewerHistoryStreamID:    req.ViewerHistoryStreamID,
		ViewerHistoryPrincipalID: req.ViewerHistoryPrincipalID,
	}

	if err := database.DB.Create(&viewerHistory).Error; err != nil {
		return nil, err
	}

	resp := dtos.CreateViewerHistoryResponse{
		ViewerHistoryID:          viewerHistory.ViewerHistoryID,
		ViewerHistoryStreamID:    viewerHistory.ViewerHistoryStreamID,
		ViewerHistoryPrincipalID: viewerHistory.ViewerHistoryPrincipalID,
	}

	return &resp, nil
}
