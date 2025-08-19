package services

import (
	"main/database"
	"main/dtos"
	"main/models"
)

func CreateOrUpdateStreamInfo(req dtos.CreateStreamInfoRequest) error {
	streamInfo := models.StreamInfo{
		HostPrincipalID:  req.HostPrincipalID,
		Title:            req.Title,
		StreamCategoryID: &req.CategoryID,
	}
	if err := database.DB.Save(&streamInfo).Error; err != nil {
		return err
	}

	return nil
}

func GetStreamInfoByUserID(hostPrincipalID string) (*models.StreamInfo, error) {
	var streamInfo models.StreamInfo

	if err := database.DB.Preload("Category").Where("host_principal_id = ?", hostPrincipalID).First(&streamInfo).Error; err != nil {
		return nil, err
	}

	return &streamInfo, nil
}
