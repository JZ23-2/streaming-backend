package models

import "time"

type Message struct {
	MessageID          string `gorm:"primaryKey;type:varchar(255)"`
	StreamID           string `gorm:"type:varchar(255)"`
	MessagePrincipalID string `gorm:"type:varchar(255)"`
	Username           string `gorm:"type:varchar(255)"`
	Content            string `gorm:"type:text"`
	CreatedAt          time.Time
}
