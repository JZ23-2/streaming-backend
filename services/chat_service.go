package services

import (
	"main/database"
	"main/helper"
	"main/models"
	"time"
)

func SaveMessage(streamID, userID, username, content string) error {
	message := models.Message{
		MessageID:          helper.GenerateID(),
		StreamID:           streamID,
		MessagePrincipalID: userID,
		Content:            content,
		Username:           username,
		CreatedAt:          time.Now(),
	}

	return database.DB.Create(&message).Error
}
