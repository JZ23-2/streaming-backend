package models

import "time"

type Stream struct {
	StreamID        string  `gorm:"primaryKey;type:varchar(255)"`
	HostPrincipalID string  `gorm:"type:varchar(255)"`
	IsActive        bool    `gorm:"type:bool;default:false"`
	ThumbnailURL    string  `gorm:"thumbnailURL"`
	StreamInfoID    *string `gorm:"type:varchar(255);"` // Nullable FK

	Messages   []Message   `gorm:"foreignKey:StreamID;references:StreamID"`
	StreamInfo *StreamInfo `gorm:"foreignKey:StreamInfoID;references:HostPrincipalID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Nullable relation
	CreatedAt  time.Time
}
