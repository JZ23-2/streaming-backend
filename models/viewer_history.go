package models

type ViewerHistory struct {
	ViewerHistoryID          string `gorm:"PrimaryKey;type:varchar(255)"`
	ViewerHistoryPrincipalID string `gorm:"type:varchar(255)"`
	ViewerHistoryStreamID    string `gorm:"type:varchar(255)"`

	Stream Stream `gorm:"foreignKey:ViewerHistoryStreamID;references:StreamID"`
}
