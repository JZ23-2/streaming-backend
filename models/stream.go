package models

import "time"

type Stream struct {
	StreamID  string    `gorm:"primaryKey;type:varchar(255)"`
	HostID    string    `gorm:"type:varchar(255)"`
	Title     string    `gorm:"type:varchar(255)"`
	Messages  []Message `gorm:"foreignKey:StreamID;references:StreamID"`
	CreatedAt time.Time

	Host User `gorm:"foreignKey:HostID;references:UserID"`
}
