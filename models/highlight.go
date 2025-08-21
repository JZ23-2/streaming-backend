package models

type Highlight struct {
	HighlightID              string `gorm:"PrimaryKey;type:varchar(255)"`
	HighlightStreamHistoryID string `gorm:"type:varchar(255)"`
	HighlightUrl             string `gorm:"varchar(255)"`
	StartHighlight           string `gorm:"varchar(20)"`
	EndHighlight             string `gorm:"varchar(20)"`
	HighlightDescription     string `gorm:"varchar(255)"`

	StreamHistory StreamHistory `gorm:"foreignKey:HighlightStreamHistoryID;references:StreamHistoryID"`
}
