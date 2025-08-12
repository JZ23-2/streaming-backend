package models

import "time"

type Message struct {
	MessageID     string `gorm:"primaryKey;type:varchar(255)"`
	StreamID      string `gorm:"type:varchar(255)"`
	UserMessageID string `gorm:"type:varchar(255)"`
	Content       string `gorm:"type:text"`
	CreatedAt     time.Time

	User User `gorm:"foreignKey:UserMessageID;references:UserID"`
}
