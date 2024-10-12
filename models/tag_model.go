package models

type TagModel struct {
	MODEL
	Title string `gorm:"size:16;comment:タグの名" json:"title"`
}
