package models

import "time"

type Stream struct {
	StreamID         string `gorm:"primaryKey;type:varchar(255)"`
	HostPrincipalID  string `gorm:"type:varchar(255)"`
	Title            string `gorm:"type:varchar(255)"`
	Thumbnail        string `gorm:"type:varchar(255)"`
	StreamCategoryID string `gorm:"type:varchar(255)"`

	Category  Category  `gorm:"foreignKey:StreamCategoryID;references:CategoryID"`
	Messages  []Message `gorm:"foreignKey:StreamID;references:StreamID"`
	CreatedAt time.Time
}
