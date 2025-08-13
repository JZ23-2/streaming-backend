package models

type Category struct {
	CategoryID   string `gorm:"PrimaryKey;type:varchar(255)"`
	CategoryName string `gorm:"type:varchar(255)"`
}
