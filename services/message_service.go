package services

import (
	"main/database"
	"main/models"
)

func GetMessagesByStreamID(streamID string) ([]models.Message, error) {
	var messages []models.Message
	err := database.DB.Where("stream_id = ?", streamID).
		Preload("User").
		Find(&messages).Error
	return messages, err
}
