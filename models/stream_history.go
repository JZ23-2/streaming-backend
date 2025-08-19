package models

type StreamHistory struct {
	StreamHistoryID       string  `gorm:"PrimaryKey;type:varchar(255)"`
	StreamHistoryStreamID string  `gorm:"type:varchar(255)"`
	HostPrincipalID       string  `gorm:"type:varchar(255)"`
	Title                 string  `gorm:"type:varchar(255)"`
	StreamCategoryID      *string `gorm:"type:varchar(255)"`
	VideoUrl              string  `gorm:"type:varchar(255)"`
	Duration              int     `gorm:"type:int"`

	Category *Category `gorm:"foreignKey:StreamCategoryID;references:CategoryID"`
	Stream   Stream    `gorm:"foreignKey:StreamHistoryStreamID;references:StreamID"`
}
