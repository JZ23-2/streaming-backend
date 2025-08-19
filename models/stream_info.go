package models

type StreamInfo struct {
	HostPrincipalID  string  `gorm:"primaryKey;type:varchar(255)"`
	Title            string  `gorm:"type:varchar(255)"`
	StreamCategoryID *string `gorm:"type:varchar(255)"`

	Category *Category `gorm:"foreignKey:StreamCategoryID;references:CategoryID"`
}
