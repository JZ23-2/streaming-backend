package models

import "time"

type User struct {
	UserID    string `gorm:"PrimaryKey;type:varchar(255)"`
	Username  string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
}
