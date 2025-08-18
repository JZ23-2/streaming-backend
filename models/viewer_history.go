package models

type ViewerHistory struct {
	ViewerHistoryID              string `gorm:"PrimaryKey;type:varchar(255)"`
	ViewerHistoryPrincipalID     string `gorm:"type:varchar(255)"`
	ViewerHistoryStreamHistoryID string `gorm:"type:varchar(255)"`

	StreamHistory StreamHistory `gorm:"foreignKey:ViewerHistoryStreamHistoryID;references:StreamHistoryID"`
}
