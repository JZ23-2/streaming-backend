package models

type StreamHistory struct {
	StreamHistoryID       string `gorm:"PrimaryKey;type:varchar(255)"`
	StreamHistoryStreamID string `gorm:"type:varchar(255)"`
	HostPrincipalID       string `gorm:"type:varchar(255)"`
	VideoUrl              string `gorm:"type:varchar(255)"`
	Duration              int    `gorm:"type:int"`

	Stream Stream `gorm:"foreignKey:StreamHistoryStreamID;references:StreamID"`
}
